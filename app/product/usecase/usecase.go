package usecase

import (
	"MengGoods/app/product/domain/repository"
	"MengGoods/app/product/domain/service"
)

type ProductUsecase struct {
	service *service.ProductService
	db      repository.ProductDB
	cache   repository.ProductCache
	mq      repository.ProductMq
	es      repository.ProductEs
	rpc     repository.ProductRpc
	cos     repository.ProductCos
}

func NewProductUsecase(service *service.ProductService, db repository.ProductDB, cache repository.ProductCache, mq repository.ProductMq, es repository.ProductEs, rpc repository.ProductRpc, cos repository.ProductCos) *ProductUsecase {
	return &ProductUsecase{
		service: service,	
		db:      db,
		cache:   cache,
		mq:      mq,
		es:      es,
		rpc:     rpc,
		cos:     cos,
	}
}
