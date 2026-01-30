package resp

import "MengGoods/kitex_gen/model"

type AddCartItemResp struct {
}

type GetCartItemResp struct {
	CartItem []*model.CartItem `json:"cartItems"`
}

type UpdateCartItemResp struct {
}

type DeleteCartItemResp struct {
}
