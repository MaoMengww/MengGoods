package router

import (
	api "MengGoods/app/gateway/handler/api/cart"
	"MengGoods/app/gateway/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitCartRouter(h *server.Hertz) {
	h.Use(mw.Sentinel())
	cart := h.Group("/api/v1/cart")
	cart.Use(mw.AuthMiddleware())
	{
		cart.POST("/add", api.AddCartItem)
		cart.POST("/delete", api.DeleteCartItem)
		cart.POST("/update", api.UpdateCartItem)
		cart.GET("/list", api.GetCartItem)
	}
}