package prpc

import (
	"MengGoods/kitex_gen/cart"
	"MengGoods/kitex_gen/model"
	"MengGoods/pkg/merror"
	"context"
)

func (c *OrderRpc) GetCartItems(ctx context.Context) ([]*model.CartItem, error) {
	resp, err := c.cartClient.GetCartItem(ctx, &cart.GetCartItemReq{})
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return resp.CartItems, nil
}
