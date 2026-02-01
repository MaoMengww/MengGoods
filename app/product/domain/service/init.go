package service

import (
	"MengGoods/app/product/domain/repository"
	"MengGoods/pkg/logger"
	"context"
)

type ProductService struct {
	db    repository.ProductDB
	cache repository.ProductCache
	mq    repository.ProductMq
	es    repository.ProductEs
	rpc   repository.ProductRpc
	cos   repository.ProductCos
}

func NewProductService(db repository.ProductDB, cache repository.ProductCache, mq repository.ProductMq, es repository.ProductEs, rpc repository.ProductRpc, cos repository.ProductCos) *ProductService {
	return &ProductService{
		db:    db,
		cache: cache,
		es:    es,
		mq:    mq,
		rpc:   rpc,
		cos:   cos,
	}
}

func (s *ProductService) Init() {
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
