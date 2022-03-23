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

	"github.com/txchat/dtalk/pkg/logger"

	"github.com/Terry-Mao/goim/pkg/ip"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/service/group/config"
	"github.com/txchat/dtalk/service/group/server/grpc"
	"github.com/txchat/dtalk/service/group/service"
)

const srvName = "group"

var (
	// projectVersion 项目版本
	projectVersion = "2.3.1"
	// goVersion go版本
	goVersion = ""
	// gitCommit git提交commit id
	gitCommit = ""
	// buildTime 编译时间
	buildTime = ""

	isShowVersion = flag.Bool("v", false, "show project version")
	isMaintain    = flag.Bool("maintain", false, "maintain data base")
)

// showVersion 显示项目版本信息
func showVersion(isShow bool) {
	if isShow {
		fmt.Printf("Project: %s\n", srvName)
		fmt.Printf(" Version: %s\n", projectVersion)
		fmt.Printf(" Go Version: %s\n", goVersion)
		fmt.Printf(" Git Commit: %s\n", gitCommit)
		fmt.Printf(" Build Time: %s\n", buildTime)
		os.Exit(0)
	}
}

func maintain(isMaintain bool, svc *service.Service) {
	if isMaintain {
		//err := svc.MaintainGroupAESKey()
		//if err != nil {
		//	panic(err)
		//}
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	showVersion(*isShowVersion)

	if err := config.Init(); err != nil {
		panic(err)
	}
	//log init
	log := logger.New(config.Conf.Env, srvName)
	log.Info().Interface("Config", config.Conf).
		Msg("config info")

	// service init
	svc := service.New(config.Conf)
	maintain(*isMaintain, svc)

	//
	//httpSrv := http.Init(svc)

	// rpc init
	rpc := grpc.New(config.Conf.GRPCServer, svc)

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
			// 退出时从 etcd 中删除
			if err := naming.UnRegister(config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema); err != nil {
				log.Error().Err(err).Msg("naming.UnRegister")
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			//if err := httpSrv.Shutdown(ctx); err != nil {
			//	log.Error().Err(err).Msg("server shutdown")
			//}
			//time.Sleep(time.Second * 2)
			if err := rpc.Shutdown(ctx); err != nil {
				log.Error().Err(err).Msg("rpc Shutdown")
			}

			log.Info().Msg(srvName + " server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
