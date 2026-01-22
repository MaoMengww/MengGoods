package usecase

import (
	"MengGoods/app/cart/domain/repository"
)

type CartUsecase struct {
	cache repository.CartCache
}

func NewCartUsecase(cache repository.CartCache) *CartUsecase {
	return &CartUsecase{
		cache: cache,
	}
}

