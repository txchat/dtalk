package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/txchat/dtalk/service/vip/dao/mock"

	"github.com/Terry-Mao/goim/pkg/ip"
	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/service/vip/config"
	"github.com/txchat/dtalk/service/vip/server/grpc"
	"github.com/txchat/dtalk/service/vip/service"
	xlog "github.com/txchat/im-pkg/log"
)

const srvName = "vip"

var (
	// projectVersion 项目版本
	projectVersion = "0.0.1"
	// goVersion go版本
	goVersion = ""
	// gitCommit git提交commit id
	gitCommit = ""
	// buildTime 编译时间
	buildTime = ""
	// osArch 目标主机架构
	osArch = ""

	isShowVersion = flag.Bool("v", false, "show project version")
)

// showVersion 显示项目版本信息
func showVersion(isShow bool) {
	if isShow {
		fmt.Printf("Project: %s\n", srvName)
		fmt.Printf(" Version: %s\n", projectVersion)
		fmt.Printf(" Go Version: %s\n", goVersion)
		fmt.Printf(" Git Commit: %s\n", gitCommit)
		fmt.Printf(" Built: %s\n", buildTime)
		fmt.Printf(" OS/Arch: %s\n", osArch)
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	showVersion(*isShowVersion)

	if err := config.Init(); err != nil {
		panic(err)
	}
	// log init
	logger, err := xlog.Init(config.Conf.Log)
	if err != nil {
		panic(err)
	}
	// set global log instance
	log.Logger = logger.With().Str("service", srvName).Logger()
	log.Info().
		Interface("xLog", config.Conf.Log).
		Interface("Reg", config.Conf.Reg).
		Interface("GRPCServer", config.Conf.GRPCServer).
		Interface("MySQL", config.Conf.MySQL).
		Msg("config info")

	// repository init
	//repo := dao.NewVIPRepositoryMySQL(config.Conf.Env, config.Conf.MySQL)
	repo := mock.NewAllowMockUsers(config.Conf.Whitelists)
	// service init
	svc := service.New(repo)
	rpc := grpc.New(config.Conf.GRPCServer, svc, log.Logger)

	// register server
	_, port, _ := net.SplitHostPort(config.Conf.GRPCServer.Addr)
	addr := fmt.Sprintf("%s:%s", ip.InternalIP(), port)
	if err := naming.Register(config.Conf.Reg.RegAddrs, config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema, 15); err != nil {
		panic(err)
	}
	fmt.Println("register ok")

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info().Str("signal", s.String()).Msg("service get a signal")
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := naming.UnRegister(config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema); err != nil {
				log.Error().Err(err).Msg("naming.UnRegister")
			}
			if err := rpc.Shutdown(ctx); err != nil {
				log.Error().Err(err).Msg("rpc.Shutdown")
			}
			svc.Shutdown(ctx)
			time.Sleep(time.Second * 2)
			log.Info().Str("name", srvName).Msg("server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
