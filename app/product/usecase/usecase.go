package usecase

import (
	"MengGoods/app/product/domain/repository"
	"MengGoods/app/product/domain/service"
)

type ProductUsecase struct {
	service *service.ProductUsecase
	db      repository.ProductDB
	cache   repository.ProductCache
	mq      repository.ProductMq
	es      repository.ProductEs
	rpc     repository.ProductRpc
}

func NewProductUsecase(db repository.ProductDB, cache repository.ProductCache, mq repository.ProductMq, es repository.ProductEs, rpc repository.ProductRpc) *ProductUsecase {
	return &ProductUsecase{
		service: service.NewProductUsecase(db, cache, mq, es, rpc),
		db:      db,
		cache:   cache,
		mq:      mq,
		es:      es,
		rpc:     rpc,
	}
}
