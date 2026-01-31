package main

import (
	"MengGoods/app/payment"
	"MengGoods/config"
	"MengGoods/kitex_gen/payment/paymentservice"
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
	shutdown := base.InitTracing("payment")
	defer shutdown(context.Background())
	register, err := etcd.NewEtcdRegistry(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("Error creating etcd registry: %s", err)
	}
	adder, err := net.ResolveTCPAddr("tcp", config.Conf.Server.Payment)
	if err != nil {
		logger.Fatalf("Error resolving TCP address: %s", err)
	}
	svr := paymentservice.NewServer(
		payment.InjectPaymentServiceImpl(),
		server.WithRegistry(register),
		server.WithServiceAddr(adder),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "payment",
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
