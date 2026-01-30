package router

import (
	api "MengGoods/app/gateway/handler/api/stock"
	"MengGoods/app/gateway/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitStockRouter(h *server.Hertz) {
	h.Use(mw.Sentinel())
	stock := h.Group("/api/v1/stock")
	stock.Use(mw.AuthMiddleware())
	{
		stock.POST("/create", api.CreateStock)
		stock.POST("/add", api.AddStock)
		stock.GET("/list", api.GetStocks)
		stock.GET("/get", api.GetStock)
	}

}