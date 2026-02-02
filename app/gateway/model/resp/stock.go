package resp

type GetStockResp struct {
	Stock int64 `json:"stock"`
}

type GetStocksResp struct {
	Stocks []*StockItem `json:"stocks"`
}

type StockItem struct {
	SkuId int64 `json:"skuId"`
	Stock int64 `json:"stock"`
}

type CreateStockResp struct {
}

type AddStockResp struct {
}
