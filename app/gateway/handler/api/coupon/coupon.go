package coupon

import (
	"MengGoods/app/gateway/rpc"
	"MengGoods/kitex_gen/coupon"
	"MengGoods/pkg/base"
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetCouponInfo(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.ParseInt(c.Param("couponId"), 10, 64)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.GetCouponInfo(ctx, &coupon.GetCouponInfoReq{
		CouponId: id,
	})
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp.Coupon)
}

func CreateCouponBatch(ctx context.Context, c *app.RequestContext) {
	var req coupon.CreateCouponBatchReq
	if err := c.BindAndValidate(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.CreateCouponBatch(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func GetCoupon(ctx context.Context, c *app.RequestContext) {
	var req coupon.GetCouponReq
	if err := c.BindAndValidate(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	_, err := rpc.GetCoupon(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}

func GetCouponList(ctx context.Context, c *app.RequestContext) {
	var req coupon.GetCouponListReq
	if err := c.BindAndValidate(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.GetCouponList(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}
