package service

import (
	"MengGoods/app/product/domain/repository"
	"MengGoods/pkg/logger"
	"context"
)

type ProductUsecase struct {
	db    repository.ProductDB
	cache repository.ProductCache
	mq    repository.ProductMq
	es    repository.ProductEs
	rpc   repository.ProductRpc
}

func NewProductUsecase(db repository.ProductDB, cache repository.ProductCache, mq repository.ProductMq, es repository.ProductEs, rpc repository.ProductRpc) *ProductUsecase {
	return &ProductUsecase{
		db:    db,
		cache: cache,
		es:    es,
		mq:    mq,
		rpc:   rpc,
	}
}

func (s *ProductUsecase) Init() {
	if err := s.ConsumeCreateSpuInfo(context.Background()); err != nil {
		logger.CtxFatalf(context.Background(), "Consume create spu info error: %v", err)
	}
	if err := s.ConsumeUpdateSpuInfo(context.Background()); err != nil {
		logger.CtxFatalf(context.Background(), "Consume update spu info error: %v", err)
	}
	if err := s.ConsumeDeleteSpuInfo(context.Background()); err != nil {
		logger.CtxFatalf(context.Background(), "Consume delete spu info error: %v", err)
	}
}
