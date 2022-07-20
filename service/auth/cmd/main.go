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
	"github.com/txchat/dtalk/pkg/naming"
	"github.com/txchat/dtalk/service/auth/server/grpc"

	"github.com/inconshreveable/log15"
	"github.com/txchat/dtalk/service/auth/config"
	"github.com/txchat/dtalk/service/auth/server/http"
	"github.com/txchat/dtalk/service/auth/service"
)

const srvName = "auth"

var log = log15.New("cmd", srvName)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	log.Info("config info:",
		"Server", *config.Conf.Server,
		"reg", *config.Conf.Reg,
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
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			// etcd 删除 kv
			log.Info("exit", "", "etcd 删除 kv")
			if err := naming.UnRegister(config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema); err != nil {
				log.Error("naming.UnRegister:", "err", err)
			}
			// 关闭 grpc
			log.Info("exit", "", "关闭 grpc")
			if err := rpc.Shutdown(ctx); err != nil {
				log.Error("rpc Shutdown:", "err", err)
			}
			// 关闭 http
			log.Info("exit", "", "关闭 http")
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("server shutdown:", "err", err)
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
