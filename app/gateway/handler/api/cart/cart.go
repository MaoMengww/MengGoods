package cart

import (
	"MengGoods/app/gateway/rpc"
	"MengGoods/kitex_gen/cart"
	"MengGoods/pkg/base"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var req cart.AddCartItemReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.AddCartItem(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func GetCartItem(ctx context.Context, c *app.RequestContext) {
	var req cart.GetCartItemReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.GetCartItem(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func UpdateCartItem(ctx context.Context, c *app.RequestContext) {
	var req cart.UpdateCartItemReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.UpdateCartItem(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func DeleteCartItem(ctx context.Context, c *app.RequestContext) {
	var req cart.DeleteCartItemReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.DeleteCartItem(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}