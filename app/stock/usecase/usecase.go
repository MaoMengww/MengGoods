package usecase

import (
	"MengGoods/app/stock/domain/model"
	"context"
)

func (u *StockUsecase) CreateStock(ctx context.Context, model *model.StockItem) error {
	return u.service.CreateStock(ctx, model)
}

func (u *StockUsecase) AddStock(ctx context.Context, model *model.StockItem) error {
	return u.service.AddStock(ctx, model)
}

func (u *StockUsecase) GetStock(ctx context.Context, skuId int64) (*model.StockItem, error) {
	return u.service.GetStock(ctx, skuId)
}

func (u *StockUsecase) GetStocks(ctx context.Context, skuIds []int64) ([]*model.StockItem, error) {
	return u.service.GetStocks(ctx, skuIds)
}

func (u *StockUsecase) LockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error {
	return u.service.LockStock(ctx, orderId, stockItems)
}

func (u *StockUsecase) UnlockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error {
	return u.service.UnlockStock(ctx, orderId, stockItems)
}

func (u *StockUsecase) DeductStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error {
	return u.service.DeductStock(ctx, orderId, stockItems)
}
