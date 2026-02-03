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
	service := &ProductService{
		db:    db,
		cache: cache,
		es:    es,
		mq:    mq,
		rpc:   rpc,
		cos:   cos,
	}
	service.Init()
	return service
}

func (s *ProductService) Init() {
	spuIds, skuIds, err := s.db.GetAllSpuIdAndSkuId(context.Background())
	if err != nil {
		logger.CtxFatalf(context.Background(), "Get all spu id and sku id error: %v", err)
	}
	if err := s.cache.LoadBloomFilter(context.Background(), spuIds, skuIds); err != nil {
		logger.CtxFatalf(context.Background(), "Load bloom filter error: %v", err)
	}
	if err := s.cache.LoadBloomFilter(context.Background(), spuIds, skuIds); err != nil {
		logger.CtxFatalf(context.Background(), "Load bloom filter error: %v", err)
	}
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
