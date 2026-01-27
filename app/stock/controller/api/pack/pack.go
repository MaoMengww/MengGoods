package pack

import (
	mModel "MengGoods/app/stock/domain/model"
	"MengGoods/kitex_gen/model"
)

func ToRpcStock(item *mModel.StockItem) *model.StockItem {
	return &model.StockItem{
		SkuId: item.SkuId,
		Count: item.Count,
	}
}

func ToRpcStocks(items []*mModel.StockItem) []*model.StockItem {
	var stocks []*model.StockItem
	for _, item := range items {
		stocks = append(stocks, ToRpcStock(item))
	}
	return stocks
}

func ToDomainStock(item *model.StockItem) *mModel.StockItem {
	return &mModel.StockItem{
		SkuId: item.SkuId,
		Count: item.Count,
	}
}

func ToDomainStocks(items []*model.StockItem) []*mModel.StockItem {
	var stocks []*mModel.StockItem
	for _, item := range items {
		stocks = append(stocks, ToDomainStock(item))
	}
	return stocks
}
