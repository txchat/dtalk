package main

import (
	"flag"
	"fmt"

	"github.com/txchat/dtalk/app/services/transfer/internal/config"
	"github.com/txchat/dtalk/app/services/transfer/internal/server"
	"github.com/txchat/dtalk/app/services/transfer/internal/svc"
	"github.com/txchat/dtalk/app/services/transfer/transfer"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/transfer.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		transfer.RegisterTransferServer(grpcServer, server.NewTransferServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
