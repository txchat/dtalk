package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/service/offline-push/config"
	_ "github.com/txchat/dtalk/service/offline-push/pusher/android"
	_ "github.com/txchat/dtalk/service/offline-push/pusher/ios"
	"github.com/txchat/dtalk/service/offline-push/server/http"
	"github.com/txchat/dtalk/service/offline-push/service"
)

const srvName = "offline-push"

var log = log15.New("cmd", srvName)

var (
	// projectVersion 项目版本
	projectVersion = "0.0.2"
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
	log.Info("config info:",
		"AppId", config.Conf.AppId,
		"Pushers Android", config.Conf.Pushers["Android"],
		"Pushers iOS", config.Conf.Pushers["iOS"],
		"MQSub", config.Conf.MQSub)
	// service init
	svc := service.New(config.Conf)
	svc.ListenMQ()

	httpSrv := http.Init()
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("server Shutdown", "err", err)
			}
			time.Sleep(time.Second * 2)
			log.Info(srvName + " server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
