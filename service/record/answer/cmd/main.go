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
	"github.com/txchat/dtalk/service/record/answer/config"
	"github.com/txchat/dtalk/service/record/answer/server/grpc"
	"github.com/txchat/dtalk/service/record/answer/server/http"
	"github.com/txchat/dtalk/service/record/answer/service"
	xlog "github.com/txchat/im-pkg/log"
	"github.com/txchat/im-pkg/trace"
)

const srvName = "answer"

var (
	// projectVersion 项目版本
	projectVersion = "0.6.6"
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
	// log init
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
		Interface("IdGenerator", config.Conf.IdGenRPCClient).
		Interface("GRPCServer", config.Conf.GRPCServer).
		Interface("LogicAddr", config.Conf.LogicRPCClient).
		Interface("Redis", config.Conf.Redis).
		Interface("MQSub", config.Conf.MQSub).
		Interface("MQPub", config.Conf.MQPub).
		Msg("config info")

	// trace init
	tracer, tracerCloser := trace.Init(srvName, config.Conf.Trace)
	//不然后续不会有Jaeger实例
	opentracing.SetGlobalTracer(tracer)

	// service init
	svc := service.New(config.Conf)
	rpc := grpc.New(config.Conf.GRPCServer, svc, log.Logger)
	svc.ListenMQ()

	httpSrv := http.Init()
	// register server
	_, port, _ := net.SplitHostPort(config.Conf.GRPCServer.Addr)
	addr := fmt.Sprintf("%s:%s", ip.InternalIP(), port)
	if err := naming.Register(config.Conf.Reg.RegAddrs, config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema, 15); err != nil {
		panic(err)
	}
	log.Info().Msg("register ok")

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
			if err := rpc.Shutdown(ctx); err != nil {
				log.Error().Err(err).Msg("rpc.Shutdown")
			}
			if err := naming.UnRegister(config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema); err != nil {
				log.Error().Err(err).Msg("naming.UnRegister")
			}
			if err := tracerCloser.Close(); err != nil {
				log.Error().Err(err).Msg("tracer close failed")
			}
			svc.Shutdown(ctx)
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error().Err(err).Msg("server shutdown")
			}
			time.Sleep(time.Second * 2)
			log.Info().Msg("server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
