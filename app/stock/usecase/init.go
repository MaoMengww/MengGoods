package usecase

import (
	"MengGoods/app/stock/domain/repository"
	"MengGoods/app/stock/domain/service"
)

type StockUsecase struct {
	service *service.StockService
	db      repository.StockDB
	cache   repository.StockCache
	mq      repository.StockMq
}

func NewStockUsecase(db repository.StockDB, cache repository.StockCache, mq repository.StockMq, service *service.StockService) *StockUsecase {
	return &StockUsecase{
		service: service,
		db:      db,
		cache:   cache,
		mq:      mq,
	}
}
