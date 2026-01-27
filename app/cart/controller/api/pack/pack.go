package pack

import (
	mModel "MengGoods/app/cart/domain/model"
	"MengGoods/kitex_gen/model"
	"time"
)

func ToDomainCartItem(cartItem *model.CartItem) *mModel.CartItem {
	return &mModel.CartItem{
		Count:      cartItem.Count,
		SkuID:      cartItem.SkuId,
		UpdateTime: time.Now().Unix(),
	}
}

func ToRpcCartItem(cartItem *mModel.CartItem) *model.CartItem {
	return &model.CartItem{
		Count:      cartItem.Count,
		SkuId:      cartItem.SkuID,
		UpdateTime: cartItem.UpdateTime,
	}
}

func ToRpcCartItems(cartItems []*mModel.CartItem) []*model.CartItem {
	rpcCartItems := make([]*model.CartItem, 0, len(cartItems))
	for _, cartItem := range cartItems {
		rpcCartItems = append(rpcCartItems, ToRpcCartItem(cartItem))
	}
	return rpcCartItems
}

func ToDomainCartItems(cartItems []*model.CartItem) []*mModel.CartItem {
	domainCartItems := make([]*mModel.CartItem, 0, len(cartItems))
	for _, cartItem := range cartItems {
		domainCartItems = append(domainCartItems, ToDomainCartItem(cartItem))
	}
	return domainCartItems
}
