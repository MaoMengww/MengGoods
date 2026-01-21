package repository

import (
	"MengGoods/app/stock/domain/model"
	"context"
)

type StockDB interface {
	CreateStock(ctx context.Context, stockItem *model.StockItem) error
	AddStock(ctx context.Context, stockItem *model.StockItem) error
	GetStock(ctx context.Context, skuId int64) (int32, error)
	LockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error
	UnlockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error
	DeductStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error
}

type StockCache interface {
	GetStockKey(ctx context.Context, skuId int64) string
	SetStock(ctx context.Context, key string, count int32) error
	GetStock(ctx context.Context, key string) (int32, error)
	AddStock(ctx context.Context, key string, count int32) error
	RedStock(ctx context.Context, stockItems map[string]int32) error
}

type StockMq interface {
	SendLockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error
	SendUnlockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error
	SendDeductStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error
	ConsumeLockStock(ctx context.Context, fn func(ctx context.Context, orderId int64, stockItems []*model.StockItem) error) error
	ConsumeUnlockStock(ctx context.Context, fn func(ctx context.Context, orderId int64, stockItems []*model.StockItem) error) error
	ConsumeDeductStock(ctx context.Context, fn func(ctx context.Context, orderId int64, stockItems []*model.StockItem) error) error
}
