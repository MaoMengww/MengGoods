package router

import (
	api "MengGoods/app/gateway/handler/api/payment"
	"MengGoods/app/gateway/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitPaymentRouter(h *server.Hertz) {
	h.Use(mw.Sentinel())
	payment := h.Group("/api/v1/payment")
	payment.Use(mw.AuthMiddleware())
	{
		payment.POST("/paymentToken", api.GetPaymentToken)
		payment.POST("/payment", api.Payment)
		payment.POST("/refund", api.PaymentRefund)
		payment.POST("/reviewRefund", api.ReviewRefund)
	}
}