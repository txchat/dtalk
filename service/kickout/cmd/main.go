package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/service/kickout/config"
	"github.com/txchat/dtalk/service/kickout/service"
	xlog "github.com/txchat/im-pkg/log"
)

const srvName = "kickout"

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
		Str("env", config.Conf.Env).
		Interface("xLog", config.Conf.Log).
		Str("TaskSpec", config.Conf.TaskSpec).
		Interface("GroupRPCClient", config.Conf.GroupRPCClient).
		Interface("SlgHTTPClient", config.Conf.SlgHTTPClient).
		Msg("config info")

	// service init
	svc := service.New(config.Conf)
	svc.Run(context.Background(), config.Conf.TaskSpec)

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
			svc.Shutdown(ctx)
			log.Info().Str("name", srvName).Msg("server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
