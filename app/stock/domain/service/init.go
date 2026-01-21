package service

import (
	"MengGoods/app/stock/domain/repository"
	"context"
)

type StockService struct {
	StockDB repository.StockDB
	StockCache repository.StockCache
	StockMq repository.StockMq
}

func NewStockService(db repository.StockDB, cache repository.StockCache, mq repository.StockMq) *StockService {
  return &StockService{
    StockDB: db,
    StockCache: cache,
    StockMq: mq,
  }
}

func (s *StockService) Init() {
  s.ConsumeLockStock(context.Background())
  s.ConsumeUnlockStock(context.Background())
  s.ConsumeDeductStock(context.Background())
}
