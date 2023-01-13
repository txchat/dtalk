package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/txchat/dtalk/app/services/oss/internal/config"
	"github.com/txchat/dtalk/app/services/oss/internal/server"
	"github.com/txchat/dtalk/app/services/oss/internal/svc"
	"github.com/txchat/dtalk/app/services/oss/oss"
	xerror "github.com/txchat/dtalk/pkg/error"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	// projectName 项目名称
	projectName = "oss"
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
	configFile = flag.String("f", "etc/oss.yaml", "the config file")
)

func main() {
	flag.Parse()
	showVersion(*isShowVersion)

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		oss.RegisterOssServer(grpcServer, server.NewOssServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(xerror.ErrInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
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
