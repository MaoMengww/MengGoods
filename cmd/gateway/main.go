package main

import (
	"MengGoods/app/gateway/router"
	"MengGoods/app/gateway/rpc"
	"MengGoods/config"
//	"MengGoods/pkg/base"
	"MengGoods/pkg/logger"
//	"context"

	"github.com/cloudwego/hertz/pkg/app/server"
)
func main() {
	config.Init()
	logger.InitLogger()
	/* p := base.InitTracing("gateway")
	defer p(context.Background()) */
	rpc.UserInit()
	rpc.ProductInit()
	h := server.Default(server.WithHostPorts(":10000"))
	router.InitUserRouter(h)
	router.InitProductRouter(h)
	h.Spin()
}