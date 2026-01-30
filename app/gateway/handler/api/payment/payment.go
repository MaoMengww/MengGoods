package payment

import (
	"MengGoods/app/gateway/rpc"
	"MengGoods/kitex_gen/payment"
	"MengGoods/pkg/base"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetPaymentToken(ctx context.Context, c *app.RequestContext) {
	var req payment.GetPaymentTokenReq
	if err := c.Bind(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.GetPaymentToken(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func Payment(ctx context.Context, c *app.RequestContext) {
	var req payment.PaymentReq
	if err := c.Bind(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.Payment(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func PaymentRefund(ctx context.Context, c *app.RequestContext) {
	var req payment.PaymentRefundReq
	if err := c.Bind(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.PaymentRefund(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func ReviewRefund(ctx context.Context, c *app.RequestContext) {
	var req payment.ReviewRefundReq
	if err := c.Bind(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.ReviewRefund(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}