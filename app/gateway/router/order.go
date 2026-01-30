package router

import (
	api "MengGoods/app/gateway/handler/api/order"
	"MengGoods/app/gateway/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitOrderRouter(h *server.Hertz) {
	h.Use(mw.Sentinel())
	order := h.Group("/api/v1/order")
	order.Use(mw.AuthMiddleware())
	{
		order.POST("/create", api.CreateOrder)
		order.GET("/list", api.ViewOrderList)
		order.GET("/:orderId", api.ViewOrderById)
		order.PUT("/:orderId/confirm", api.ConfirmReceiptOrder)
	}
}