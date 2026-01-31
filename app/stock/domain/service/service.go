package service

import (
	"MengGoods/app/stock/domain/model"
	"MengGoods/pkg/merror"
	"context"
)

func (s *StockService) CreateStock(ctx context.Context, Item *model.StockItem) error {
	if Item.Count <= 0 {
		return merror.NewMerror(merror.ParamCountInvalid, "Stock count must be greater than 0")
	}
	if err := s.StockDB.CreateStock(ctx, Item); err != nil {
		return err
	}
	return s.StockCache.SetStock(ctx, s.StockCache.GetStockKey(ctx, Item.SkuId), Item.Count)
}

func (s *StockService) AddStock(ctx context.Context, Item *model.StockItem) error {
	if Item.Count <= 0 {
		return merror.NewMerror(merror.ParamCountInvalid, "Stock count must be greater than 0")
	}
	if err := s.StockDB.AddStock(ctx, Item); err != nil {
		return err
	}
	return s.StockCache.AddStock(ctx, s.StockCache.GetStockKey(ctx, Item.SkuId), Item.Count)
}

func (s *StockService) GetStock(ctx context.Context, id int64) (*model.StockItem, error) {
	count, err := s.StockCache.GetStock(ctx, s.StockCache.GetStockKey(ctx, id))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return &model.StockItem{
			SkuId: id,
			Count: count,
		}, nil
	}
	count, err = s.StockDB.GetStock(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.StockItem{
		SkuId: id,
		Count: count,
	}, nil
}

func (s *StockService) GetStocks(ctx context.Context, ids []int64) ([]*model.StockItem, error) {
	stockItems := make([]*model.StockItem, 0)
	for _, id := range ids {
		item, err := s.GetStock(ctx, id)
		if err != nil {
			return nil, err
		}
		stockItems = append(stockItems, item)
	}
	return stockItems, nil
}

func (s *StockService) LockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error {
	stockMap := make(map[string]int32)
	for _, item := range stockItems {
		if item.Count <= 0 {
			return merror.NewMerror(merror.ParamCountInvalid, "Stock count must be greater than 0")
		}
		key := s.StockCache.GetStockKey(ctx, item.SkuId)
		stockMap[key] = item.Count
	}
	if err := s.StockCache.RedStock(ctx, stockMap); err != nil {
		return err
	}
	s.StockMq.SendLockStock(ctx, orderId, stockItems)
	return nil
}

func (s *StockService) UnlockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error {
	for _, item := range stockItems {
		if item.Count <= 0 {
			return merror.NewMerror(merror.ParamCountInvalid, "Stock count must be greater than 0")
		}
		key := s.StockCache.GetStockKey(ctx, item.SkuId)
		s.StockCache.AddStock(ctx, key, item.Count)
	}
	if err := s.StockMq.SendUnlockStock(ctx, orderId, stockItems); err != nil {
		return err
	}
	return nil
}

// 扣减库存
func (s *StockService) DeductStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error {
	if err := s.StockMq.SendDeductStock(ctx, orderId, stockItems); err != nil {
		return err
	}
	return nil
}

func (s *StockService) ConsumeLockStock(ctx context.Context) error {
	return s.StockMq.ConsumeLockStock(ctx, s.StockDB.LockStock)
}

func (s *StockService) ConsumeUnlockStock(ctx context.Context) error {
	return s.StockMq.ConsumeUnlockStock(ctx, s.StockDB.UnlockStock)
}

func (s *StockService) ConsumeDeductStock(ctx context.Context) error {
	return s.StockMq.ConsumeDeductStock(ctx, s.StockDB.DeductStock)
}
