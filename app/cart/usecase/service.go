package usecase

import (
	"MengGoods/app/cart/domain/model"
	"MengGoods/pkg/utils"
	"context"
)

func (s *CartUsecase) AddCartItem(ctx context.Context, cartItem *model.CartItem) error {
    if err := utils.Verify(utils.VerifyCount(cartItem.Count)); err != nil {
        return err
    }
    return s.cache.AddCartItem(ctx, cartItem)
}

func (s *CartUsecase) GetCart(ctx context.Context) ([]*model.CartItem, error) {
    return s.cache.GetCart(ctx)
}

func (s *CartUsecase) DeleteCart(ctx context.Context) error {
    return s.cache.DeleteCart(ctx)
}

func (s *CartUsecase) UpdateCartItem(ctx context.Context, cartItem *model.CartItem) error {
    if err := utils.Verify(utils.VerifyCount(cartItem.Count)); err != nil {
        return err
    }
    return s.cache.UpdateCartItem(ctx, cartItem)
}
