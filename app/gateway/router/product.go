package router

import (
	api "MengGoods/app/gateway/handler/api/product"
	"MengGoods/app/gateway/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
)


func InitProductRouter(h *server.Hertz) {
	h.Use(mw.Sentinel())
	product := h.Group("/api/v1/product")
	{
		product.POST("/SearchSpu", api.GetSpu)
		product.GET("/spu/:spu_id", api.GetSpuById)
		product.GET("/sku/:sku_id", api.GetSku)
	}
	product.Use(mw.AuthMiddleware())
	{
		product.POST("/spu/:spu_Id/update", api.UpdateSpu)
		product.POST("/sku/:sku_Id/update", api.UpdateSku)	
		product.POST("/spu/create", api.CreateSpu)
		product.POST("/spu/:spu_Id/delete", api.DeleteSpu)
		product.POST("/sku/:sku_Id/delete", api.DeleteSku)
		product.POST("/category/create", api.CreateCategory)
		product.POST("/category/delete", api.DeleteCategory)
		product.POST("/category/update", api.UpdateCategory)
	}
}