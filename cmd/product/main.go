package main

import (
	"MengGoods/app/product"
	"MengGoods/config"
	"MengGoods/kitex_gen/product/productservice"
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
	"github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
)


func Init() {
	config.Init()
	logger.InitLogger()
	utils.InitSnowflake(0)
}

func main(){
	Init()
	shutdown := base.InitTracing("product")
	defer shutdown(context.Background())
	register, err := etcd.NewEtcdRegistry(viper.GetStringSlice("etcd.endpoints"))
	if err != nil {
		logger.Fatalf("Error creating etcd registry: %s", err)
	}
	adder, err := net.ResolveTCPAddr("tcp", viper.GetString("server.Product"))
	if err != nil {
		logger.Fatalf("Error resolving TCP address: %s", err)
	}
	svr := productservice.NewServer(
		product.InjectProductUsecaseImpl(),
		server.WithRegistry(register),
		server.WithServiceAddr(adder),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "product",
		}),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnection,
			MaxQPS: constants.MaxQPS,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		//server.WithMiddleware(middleware.PackResp()),
	)
	svr.Run()
}
