package service

import (
	"MengGoods/app/product/domain/repository"
	"context"
)

type ProductUsecase struct {
	db repository.ProductDB
	cache repository.ProductCache
	mq repository.ProductMq
	es repository.ProductEs
	rpc repository.ProductRpc
}

func NewProductUsecase(db repository.ProductDB, cache repository.ProductCache, mq repository.ProductMq, es repository.ProductEs, rpc repository.ProductRpc) *ProductUsecase {
	return &ProductUsecase{
		db: db,
		cache: cache,
		es: es,
		mq: mq,
		rpc: rpc,

	}
}

func (s *ProductUsecase) Init(){
	s.ConsumeCreateSpuInfo(context.Background())
	s.ConsumeUpdateSpuInfo(context.Background())
	s.ConsumeDeleteSpuInfo(context.Background())
}

