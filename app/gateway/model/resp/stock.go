package resp

type GetStockResp struct {
	Stock int64 `json:"stock"`
}

type GetStocksResp struct {
	Stocks []*GetStockResp `json:"stocks"`
}

type CreateStockResp struct {
}

type AddStockResp struct {
}
