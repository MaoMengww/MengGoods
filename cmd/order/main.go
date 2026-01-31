package main

import (
	"MengGoods/app/order"
	"MengGoods/config"
	"MengGoods/kitex_gen/order/orderservice"
	"MengGoods/pkg/base"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/utils"
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
	utils.InitSnowflake(0)
}

func main() {
	Init()
	shutdown := base.InitTracing("order")
	defer shutdown(context.Background())
	register, err := etcd.NewEtcdRegistry(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("Error creating etcd registry: %s", err)
	}
	adder, err := net.ResolveTCPAddr("tcp", config.Conf.Server.Order)
	if err != nil {
		logger.Fatalf("Error resolving TCP address: %s", err)
	}
	svr := orderservice.NewServer(
		order.InjectOrderServiceImpl(),
		server.WithRegistry(register),
		server.WithServiceAddr(adder),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "order",
		}),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnection,
			MaxQPS:         constants.MaxQPS,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		//server.WithMiddleware(middleware.PackResp()),
	)
	svr.Run()
}
