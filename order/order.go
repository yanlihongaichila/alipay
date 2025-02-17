package main

import (
	"flag"
	"fmt"
	"order/internal/pkg"

	"pkg/app"

	"order/internal/config"
	"order/internal/server"
	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()
	//连接mysql
	err := app.Init("mysql")
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建表
	err = pkg.Migrate()
	if err != nil {
		fmt.Println(err)
		return
	}
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrderServer(grpcServer, server.NewOrderServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
