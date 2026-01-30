package rpc

import (
	//	mresp "MengGoods/app/gateway/model/resp"
	mresp "MengGoods/app/gateway/model/resp"
	"MengGoods/kitex_gen/stock/stockservice"
	"MengGoods/kitex_gen/stock"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"context"

	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
)

var StockClient stockservice.Client

func StockInit() {
	r, err := etcd.NewEtcdResolver(viper.GetStringSlice("etcd.endpoints"))
	if err != nil {
		logger.Fatalf("stock rpc Init Falied: err: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := stockservice.NewClient(
		"stock",
		client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithCircuitBreaker(cbSuite),
		client.WithSuite(tracing.NewClientSuite()),
		//
	)

	if err != nil {
		logger.Fatalf("init client failed: err:%v", err)
	}

	StockClient = c
}

func CreateStock(ctx context.Context, req *stock.CreateStockReq) (resp *mresp.CreateStockResp, err error) {
	r, err := StockClient.CreateStock(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "stock rpc CreateStock failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.CreateStockResp{}, nil
}

func AddStock(ctx context.Context, req *stock.AddStockReq) (resp *mresp.AddStockResp, err error) {
	r, err := StockClient.AddStock(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "stock rpc AddStock failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.AddStockResp{}, nil
}

func GetStock(ctx context.Context, req *stock.GetStockReq) (resp *mresp.GetStockResp, err error) {
	r, err := StockClient.GetStock(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "stock rpc GetStock failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.GetStockResp{Stock: int64(r.Stock.Count)}, nil
}

func GetStocks(ctx context.Context, req *stock.GetStocksReq) (resp *mresp.GetStocksResp, err error) {
	r, err := StockClient.GetStocks(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "stock rpc GetStocks failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.GetStocksResp{
		Stocks: make([]*mresp.GetStockResp, 0, len(r.Stocks)),
	}, nil
}