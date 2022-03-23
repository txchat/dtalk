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

	"github.com/Terry-Mao/goim/pkg/ip"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/service/record/store/config"
	"github.com/txchat/dtalk/service/record/store/server/grpc"
	"github.com/txchat/dtalk/service/record/store/server/http"
	"github.com/txchat/dtalk/service/record/store/service"
	xlog "github.com/txchat/im-pkg/log"
	"github.com/txchat/im-pkg/trace"
)

const srvName = "store"

var (
	// projectVersion 项目版本
	projectVersion = "0.7.6"
	// goVersion go版本
	goVersion = ""
	// gitCommit git提交commit id
	gitCommit = ""
	// buildTime 编译时间
	buildTime = ""

	isShowVersion = flag.Bool("v", false, "show project version")
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

// @Title 聊天单模块集成测试
// @Version 0.1
// @Description
// @SecurityDefinitions.ApiKey ApiKeyAuth
// @In header
// @Name Authorization
// @BasePath /
func main() {
	flag.Parse()
	showVersion(*isShowVersion)

	if err := config.Init(); err != nil {
		panic(err)
	}
	//log init
	logger, err := xlog.Init(config.Conf.Log)
	if err != nil {
		panic(err)
	}
	// set global log instance
	log.Logger = logger.With().Str("service", srvName).Logger()
	log.Info().
		Str("AppId", config.Conf.AppId).
		Str("Env", config.Conf.Env).
		Interface("Reg", config.Conf.Reg).
		Interface("MySQL", config.Conf.MySQL).
		Interface("Redis", config.Conf.Redis).
		Interface("PusherRPCClient", config.Conf.PusherRPCClient).
		Interface("RevSub", config.Conf.RevSub).
		Interface("StoreSub", config.Conf.StoreSub).
		Msg("config info")

	// trace init
	tracer, tracerCloser := trace.Init(srvName, config.Conf.Trace)
	//不然后续不会有Jaeger实例
	opentracing.SetGlobalTracer(tracer)

	// service init
	svc := service.New(config.Conf)
	rpc := grpc.New(config.Conf.GRPCServer, svc)
	svc.ListenMQ()

	httpSrv := http.Init()

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
			time.Sleep(time.Second * 2)
			log.Info().Str("name", srvName).Msg("server exit")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := rpc.Shutdown(ctx); err != nil {
				log.Error().Err(err).Msg("rpc.Shutdown")
			}
			if err := naming.UnRegister(config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema); err != nil {
				log.Error().Err(err).Msg("naming.UnRegister")
			}
			if err := tracerCloser.Close(); err != nil {
				log.Error().Err(err).Msg("tracer close failed")
			}
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error().Err(err).Msg("server shutdown")
			}
			svc.Shutdown(ctx)
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
