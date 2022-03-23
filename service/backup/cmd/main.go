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
	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/service/backup/config"
	"github.com/txchat/dtalk/service/backup/server/grpc"
	"github.com/txchat/dtalk/service/backup/server/http"
	"github.com/txchat/dtalk/service/backup/service"
)

const srvName = "backup"

var log = log15.New("cmd", srvName)

var (
	// projectVersion 项目版本
	projectVersion = "1.1.1"
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
	log.Info("config info:",
		"Env", config.Conf.Env,
		"SMS", config.Conf.SMS,
		"Email", config.Conf.Email,
		"Mysql", *config.Conf.MySQL,
		"Server", *config.Conf.Server,
	)
	// service init
	svc := service.New(config.Conf)
	httpSrv := http.Init(svc)

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
		log.Info("service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// 退出时从 etcd 中删除
			if err := naming.UnRegister(config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema); err != nil {
				log.Error("naming.UnRegister", "err", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("server shutdown:", "err", err)
			}
			if err := rpc.Shutdown(ctx); err != nil {
				log.Error("rpc Shutdown", "err", err)
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
