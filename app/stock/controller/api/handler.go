package api

import (
	"MengGoods/app/stock/controller/api/pack"
	"MengGoods/app/stock/domain/model"
	"MengGoods/app/stock/usecase"
	stock "MengGoods/kitex_gen/stock"
	"MengGoods/pkg/base"
	"context"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{
	usecase *usecase.StockUsecase
}

func NewStockServiceImpl(usecase *usecase.StockUsecase) *StockServiceImpl {
	return &StockServiceImpl{
		usecase: usecase,
	}
}

func (s *StockServiceImpl) CreateStock(ctx context.Context, req *stock.CreateStockReq) (resp *stock.CreateStockResp, err error) {
	resp = &stock.CreateStockResp{}
	item := &model.StockItem{
		SkuId: req.SkuId,
		Count: req.Count,
	}
	if err := s.usecase.CreateStock(ctx, item); err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// AddStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) AddStock(ctx context.Context, req *stock.AddStockReq) (resp *stock.AddStockResp, err error) {
	resp = &stock.AddStockResp{}
	item := &model.StockItem{
		SkuId: req.SkuId,
		Count: req.Count,
	}
	if err := s.usecase.AddStock(ctx, item); err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// GetStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) GetStock(ctx context.Context, req *stock.GetStockReq) (resp *stock.GetStockResp, err error) {
	resp = &stock.GetStockResp{}
	stockItem, err := s.usecase.GetStock(ctx, req.SkuId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	resp.Stock = pack.ToRpcStock(stockItem)
	return resp, nil
}

// GetStocks implements the StockServiceImpl interface.
func (s *StockServiceImpl) GetStocks(ctx context.Context, req *stock.GetStocksReq) (resp *stock.GetStocksResp, err error) {
	resp = &stock.GetStocksResp{}
	stockItems, err := s.usecase.GetStocks(ctx, req.SkuIds)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	resp.Stocks = pack.ToRpcStocks(stockItems)
	return resp, nil
}

// LockStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) LockStock(ctx context.Context, req *stock.LockStockReq) (resp *stock.LockStockResp, err error) {
	resp = &stock.LockStockResp{}
	stockItems := pack.ToDomainStocks(req.StockItems)
	if err := s.usecase.LockStock(ctx, req.OrderId, stockItems); err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// UnlockStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) UnlockStock(ctx context.Context, req *stock.UnlockStockReq) (resp *stock.UnlockStockResp, err error) {
	// TODO: Your code here...
	resp = &stock.UnlockStockResp{}
	stockItems := pack.ToDomainStocks(req.StockItems)
	if err := s.usecase.UnlockStock(ctx, req.OrderId, stockItems); err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// DeductStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) DeductStock(ctx context.Context, req *stock.DeductStockReq) (resp *stock.DeductStockResp, err error) {
	// TODO: Your code here...
	resp = &stock.DeductStockResp{}
	stockItems := pack.ToDomainStocks(req.StockItems)
	if err := s.usecase.DeductStock(ctx, req.OrderId, stockItems); err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}
