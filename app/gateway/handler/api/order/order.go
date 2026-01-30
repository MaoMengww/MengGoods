package ordergo

import (
	"MengGoods/app/gateway/rpc"
	"MengGoods/kitex_gen/order"
	"MengGoods/pkg/base"
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var req order.CreateOrderReq
	if err := c.Bind(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.CreateOrder(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func ViewOrderById(ctx context.Context, c *app.RequestContext) {
	id := c.Param("orderId")
	orderId, err := strconv.ParseInt(id, 10, 64)
	resp, err := rpc.ViewOrderById(ctx, &order.ViewOrderByIdReq{
		OrderId: orderId,
	})
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func ViewOrderList(ctx context.Context, c *app.RequestContext) {
	var req order.ViewOrderListReq
	if err := c.Bind(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.ViewOrderList(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func ConfirmReceiptOrder(ctx context.Context, c *app.RequestContext) {
	id := c.Param("orderId")
	orderId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.ConfirmReceiptOrder(ctx, &order.ConfirmReceiptOrderReq{
		OrderId: orderId,
	})
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func CancelOrder(ctx context.Context, c *app.RequestContext) {
	var req order.CancelOrderReq
	if err := c.Bind(&req); err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.CancelOrder(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}
