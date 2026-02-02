package stock

import (
	"MengGoods/app/gateway/rpc"
	"MengGoods/kitex_gen/stock"
	"MengGoods/pkg/base"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetStock(ctx context.Context, c *app.RequestContext) {
	var req stock.GetStockReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.GetStock(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func GetStocks(ctx context.Context, c *app.RequestContext) {
	var req stock.GetStocksReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.GetStocks(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func AddStock(ctx context.Context, c *app.RequestContext) {
	var req stock.AddStockReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	_, err = rpc.AddStock(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}

func CreateStock(ctx context.Context, c *app.RequestContext) {
	var req stock.CreateStockReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	_, err = rpc.CreateStock(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}
