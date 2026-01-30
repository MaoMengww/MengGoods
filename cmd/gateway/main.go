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
	rpc.CouponInit()
	rpc.CartInit()
	rpc.PaymentInit()
	rpc.OrderInit()
	rpc.StockInit()
	h := server.Default(server.WithHostPorts(":10000"))
	router.InitUserRouter(h)
	router.InitProductRouter(h)
	router.InitCouponRouter(h)
	router.InitCartRouter(h)
	router.InitPaymentRouter(h)
	router.InitOrderRouter(h)
	router.InitStockRouter(h)
	h.Spin()
}
