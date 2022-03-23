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
	"github.com/txchat/dtalk/service/generator/config"
	"github.com/txchat/dtalk/service/generator/server/grpc"
	"github.com/txchat/dtalk/service/generator/server/http"
	"github.com/txchat/dtalk/service/generator/service"
)

const srvName = "generator"

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

func main() {
	flag.Parse()
	showVersion(*isShowVersion)

	if err := config.Init(); err != nil {
		panic(err)
	}
	log.Info("config info:",
		"Node", config.Conf.Node)
	// service init
	svc := service.New(config.Conf)
	rpc := grpc.New(config.Conf.GRPCServer, svc)

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
		log.Info("service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			time.Sleep(time.Second * 2)
			if err := naming.UnRegister(config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema); err != nil {
				log.Error("naming.UnRegister", "err", err)
			}
			if err := rpc.Shutdown(ctx); err != nil {
				log.Error("rpc.Shutdown", "err", err)
			}
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("server Shutdown", "err", err)
			}
			log.Info(srvName + " server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
