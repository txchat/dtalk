package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/txchat/dtalk/app/services/pusher/rmq/rdmq"

	"github.com/txchat/dtalk/app/services/pusher/internal/config"
	"github.com/txchat/dtalk/app/services/pusher/internal/svc"
	"github.com/txchat/dtalk/app/services/pusher/rmq/connmq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	// projectName 项目名称
	projectName = "pusher-rmq"
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
	// isShowVersion 是否显示项目版本信息
	isShowVersion = flag.Bool("v", false, "show project version")
	// configFile 配置文件路径
	configFile = flag.String("f", "etc/pusher-rmq.yaml", "the config file")
)

func main() {
	flag.Parse()
	showVersion(*isShowVersion)

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	connMQSvc := connmq.NewService(c, ctx)
	connMQSvc.Serve()

	rdMQSvc := rdmq.NewService(c, ctx)
	rdMQSvc.Serve()

	// init signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sc
		logx.Info("service get a signal", "signal", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			wg := sync.WaitGroup{}
			go func() {
				defer wg.Done()
				wg.Add(1)
				connMQSvc.Shutdown(context.Background())
			}()
			rdMQSvc.Shutdown(context.Background())
			wg.Wait()
			logx.Info("server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
			logx.Info("server hangup")
		default:
			return
		}
	}
}

// showVersion 显示项目版本信息
func showVersion(isShow bool) {
	if isShow {
		fmt.Printf("Project: %s\n", projectName)
		fmt.Printf(" Version: %s\n", projectVersion)
		fmt.Printf(" Go Version: %s\n", goVersion)
		fmt.Printf(" Git Commit: %s\n", gitCommit)
		fmt.Printf(" Built: %s\n", buildTime)
		fmt.Printf(" OS/Arch: %s\n", osArch)
		os.Exit(0)
	}
}
