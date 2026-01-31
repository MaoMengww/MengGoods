package main

import (
	"MengGoods/app/user"
	"MengGoods/config"
	"MengGoods/kitex_gen/user/userservice"
	"MengGoods/pkg/base"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/middleware"
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	config.Init()
	logger.InitLogger()
}

func main() {
	Init()
	shutdown := base.InitTracing("user")
	defer shutdown(context.Background())
	register, err := etcd.NewEtcdRegistry(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("Error creating etcd registry: %s", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", config.Conf.Server.User)
	if err != nil {
		logger.Fatalf("Error resolving TCP address: %s", err)
	}
	svr := userservice.NewServer(
		user.InjectUserServiceImpl(),
		server.WithRegistry(register),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "user",
		}),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnection,
			MaxQPS:         constants.MaxQPS,
		}),
		server.WithSuite(tracing.NewServerSuite()),   // 添加PackResp中间件
		server.WithMiddleware(middleware.ErrorLog()), // 保持原有的ErrorLog中间件
	)
	svr.Run()
}
