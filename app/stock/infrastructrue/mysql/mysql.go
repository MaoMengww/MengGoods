package mysql

import (
	"MengGoods/app/stock/domain/model"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"context"
	"time"

	"gorm.io/gorm"
)

type StockDB struct {
	*gorm.DB
}

func NewStockDB(db *gorm.DB) *StockDB {
	return &StockDB{DB: db}
}

func (p *StockDB) CreateStock(ctx context.Context, item *model.StockItem) error {
	newStock := Stock{
		ID:          item.SkuId,
		Stock:       item.Count,
		LockedStock: 0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := p.DB.WithContext(ctx).Create(&newStock).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, "创建库存失败")
	}
	newStockJournal := StockJournal{
		SkuID:     item.SkuId,
		OrderID:   -1,
		Count:     item.Count,
		Type:      constants.CreateType,
		CreatedAt: time.Now(),
	}
	err = p.DB.WithContext(ctx).Create(&newStockJournal).Error
	if err != nil {
		logger.CtxErrorf(ctx, "创建库存日志失败: %v", err)
	}
	return nil
}

func (p *StockDB) AddStock(ctx context.Context, item *model.StockItem) error {
	// 检查库存是否存在
	var stock Stock
	err := p.DB.WithContext(ctx).First(&stock, item.SkuId).Error
	if err != nil {
		return merror.NewMerror(merror.StockNotExist, "库存不存在")
	}
	stock.Stock += item.Count
	stock.UpdatedAt = time.Now()
	err = p.DB.WithContext(ctx).Save(&stock).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, "更新库存失败")
	}
	newStockJournal := StockJournal{
		SkuID:     item.SkuId,
		OrderID:   -1,
		Count:     item.Count,
		Type:      constants.AddType,
		CreatedAt: time.Now(),
	}
	err = p.DB.WithContext(ctx).Create(&newStockJournal).Error
	if err != nil {
		logger.CtxErrorf(ctx, "创建库存日志失败: %v", err)
	}
	return nil
}

func (p *StockDB) GetStock(ctx context.Context, skuId int64) (int32, error) {
	var stock Stock
	err := p.DB.WithContext(ctx).First(&stock, skuId).Error
	if err != nil {
		return 0, merror.NewMerror(merror.StockNotExist, "库存不存在")
	}
	return stock.Stock, nil
}

func (p *StockDB) LockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error {
	for _, item := range stockItems {
		// 检查库存是否存在
		var stock Stock
		err := p.DB.WithContext(ctx).First(&stock, item.SkuId).Error
		if err != nil {
			return merror.NewMerror(merror.StockNotExist, "库存不存在")
		}
		// 检查库存是否充足
		if stock.Stock < item.Count {
			return merror.NewMerror(merror.StockNotEnough, "库存不足")
		}
		stock.LockedStock += item.Count
		stock.UpdatedAt = time.Now()
		err = p.DB.WithContext(ctx).Save(&stock).Error
		if err != nil {
			return merror.NewMerror(merror.InternalDatabaseErrorCode, "更新库存失败")
		}
		newStockJournal := StockJournal{
			SkuID:     item.SkuId,
			OrderID:   orderId,
			Count:     item.Count,
			Type:      constants.LockType,
			CreatedAt: time.Now(),
		}
		err = p.DB.WithContext(ctx).Create(&newStockJournal).Error
		if err != nil {
			logger.CtxErrorf(ctx, "创建库存日志失败: %v", err)
		}
	}
	return nil
}

func (p *StockDB) UnlockStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error {
	for _, item := range stockItems {
		// 检查库存是否存在
		var stock Stock
		err := p.DB.WithContext(ctx).First(&stock, item.SkuId).Error
		if err != nil {
			return merror.NewMerror(merror.StockNotExist, "库存不存在")
		}
		stock.LockedStock -= item.Count
		stock.UpdatedAt = time.Now()
		err = p.DB.WithContext(ctx).Save(&stock).Error
		if err != nil {
			return merror.NewMerror(merror.InternalDatabaseErrorCode, "更新库存失败")
		}
		newStockJournal := StockJournal{
			SkuID:     item.SkuId,
			OrderID:   orderId,
			Count:     item.Count,
			Type:      constants.UnlockType,
			CreatedAt: time.Now(),
		}
		err = p.DB.WithContext(ctx).Create(&newStockJournal).Error
		if err != nil {
			logger.CtxErrorf(ctx, "创建库存日志失败: %v", err)
		}
	}
	return nil
}

func (p *StockDB) DeductStock(ctx context.Context, orderId int64, stockItems []*model.StockItem) error {
	for _, item := range stockItems {
		// 检查库存是否存在
		var stock Stock
		err := p.DB.WithContext(ctx).First(&stock, item.SkuId).Error
		if err != nil {
			return merror.NewMerror(merror.StockNotExist, "库存不存在")
		}
		// 检查库存是否充足
		if stock.Stock < item.Count {
			return merror.NewMerror(merror.StockNotEnough, "库存不足")
		}
		stock.Stock -= item.Count
		stock.UpdatedAt = time.Now()
		err = p.DB.WithContext(ctx).Save(&stock).Error
		if err != nil {
			return merror.NewMerror(merror.InternalDatabaseErrorCode, "更新库存失败")
		}
		newStockJournal := StockJournal{
			SkuID:     item.SkuId,
			OrderID:   orderId,
			Count:     item.Count,
			Type:      constants.DeductType,
			CreatedAt: time.Now(),
		}
		err = p.DB.WithContext(ctx).Create(&newStockJournal).Error
		if err != nil {
			logger.CtxErrorf(ctx, "创建库存日志失败: %v", err)
		}
	}
	return nil
}
