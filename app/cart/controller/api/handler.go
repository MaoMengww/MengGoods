package api

import (
	"MengGoods/app/cart/controller/api/pack"
	"MengGoods/app/cart/usecase"
	cart "MengGoods/kitex_gen/cart"
	"MengGoods/pkg/base"
	"context"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct {
	CartUsecase *usecase.CartUsecase
}

func NewCartServiceImpl(cartUsecase *usecase.CartUsecase) *CartServiceImpl {
	return &CartServiceImpl{
		CartUsecase: cartUsecase,
	}
}

// AddCartItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddCartItem(ctx context.Context, req *cart.AddCartItemReq) (resp *cart.AddCartItemResp, err error) {
	resp = new(cart.AddCartItemResp)
	reqCartItem := pack.ToDomainCartItem(req.CartItem)
	err = s.CartUsecase.AddCartItem(ctx, reqCartItem)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return
}

// GetCartItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCartItem(ctx context.Context, req *cart.GetCartItemReq) (resp *cart.GetCartItemResp, err error) {
	resp = new(cart.GetCartItemResp)
	cartItems, err := s.CartUsecase.GetCart(ctx)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}
	resp.CartItems = pack.ToRpcCartItems(cartItems)
	resp.Base = base.BuildBaseResp(nil)
	return
}

// DeleteCartItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) DeleteCartItem(ctx context.Context, req *cart.DeleteCartItemReq) (resp *cart.DeleteCartItemResp, err error) {
	resp = new(cart.DeleteCartItemResp)
	err = s.CartUsecase.DeleteCart(ctx)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return
}

// UpdateCartItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) UpdateCartItem(ctx context.Context, req *cart.UpdateCartItemReq) (resp *cart.UpdateCartItemResp, err error) {
	// TODO: Your code here...
	resp = new(cart.UpdateCartItemResp)
	reqCartItem := pack.ToDomainCartItem(req.CartItem)
	err = s.CartUsecase.UpdateCartItem(ctx, reqCartItem)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return
}
