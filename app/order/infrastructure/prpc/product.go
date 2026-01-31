package prpc

import (
	"MengGoods/kitex_gen/model"
	"MengGoods/kitex_gen/product"
	"MengGoods/pkg/merror"
	"context"
)

func (c *OrderRpc) GetSkuInfo(ctx context.Context, skuId int64) (*model.SkuInfo, error) {
	resp, err := c.ProductClient.GetSku(ctx, &product.GetSkuReq{
		SkuId: skuId,
	})
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return resp.SkuInfo, nil
}
