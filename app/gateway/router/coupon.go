package router

import (
	api "MengGoods/app/gateway/handler/api/coupon"
	"MengGoods/app/gateway/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitCouponRouter(h *server.Hertz) {
	h.Use(mw.Sentinel())
	coupon := h.Group("/api/v1/coupon")
	coupon.Use(mw.AuthMiddleware())
	{
		coupon.POST("/create", api.CreateCouponBatch)
		coupon.GET("/list", api.GetCouponList)
		coupon.GET("/:couponId", api.GetCouponInfo)
		coupon.GET("/get", api.GetCoupon)
	}
}