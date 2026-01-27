package repository

import (
	"MengGoods/app/cart/domain/model"
	"context"
)

type CartCache interface {
	GetCartKey(ctx context.Context) (string, error)
	AddCartItem(ctx context.Context, cartItem *model.CartItem) error
	GetCart(ctx context.Context) ([]*model.CartItem, error)
	DeleteCart(ctx context.Context) error
	UpdateCartItem(ctx context.Context, cartItem *model.CartItem) error
}
