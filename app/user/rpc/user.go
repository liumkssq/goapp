package main

import (
	"flag"
	"fmt"

	"github.com/liumkssq/goapp/app/user/rpc/internal/config"
	"github.com/liumkssq/goapp/app/user/rpc/internal/server"
	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 配置文件路径，默认为etc/user.yaml
var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	// 解析命令行参数
	flag.Parse()

	// 加载配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 创建服务上下文
	ctx := svc.NewServiceContext(c)

	// 创建RPC服务器
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// 注册用户服务
		user.RegisterUserServiceServer(grpcServer, server.NewUserServiceServer(ctx))

		// 在开发或测试模式下注册反射服务，便于调试
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 打印启动信息并启动服务
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
