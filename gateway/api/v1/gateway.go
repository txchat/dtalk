package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/gateway/api/v1/internal/config"
	http "github.com/txchat/dtalk/gateway/api/v1/internal/handler"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/im-pkg/trace"
)

const srvName = "gateway"

var (
	// projectVersion 项目版本
	projectVersion = "0.1.2"
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

// @Title 聊天网关
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
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Str("service", srvName).Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().
		Interface("Modules", config.Conf.Modules).
		Interface("Server", config.Conf.Server).
		Msg("config info")

	//trace init
	tracer, tracerCloser := trace.Init(srvName, config.Conf.Trace)
	//不然后续不会有Jaeger实例
	opentracing.SetGlobalTracer(tracer)

	// service init
	ctx := svc.NewServiceContext(*config.Conf)
	httpSrv := http.Init(ctx)

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
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error().Err(err).Msg("server shutdown")
			}
			if err := tracerCloser.Close(); err != nil {
				log.Error().Err(err).Msg("tracer close failed")
			}
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
