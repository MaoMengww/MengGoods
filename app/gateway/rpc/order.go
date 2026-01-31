package rpc

import (
	//	mresp "MengGoods/app/gateway/model/resp"
	mresp "MengGoods/app/gateway/model/resp"
	"MengGoods/kitex_gen/order"
	"MengGoods/kitex_gen/order/orderservice"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"context"

	"time"

	"MengGoods/config"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var OrderClient orderservice.Client

func OrderInit() {
	r, err := etcd.NewEtcdResolver(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("order rpc Init Falied: err: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := orderservice.NewClient(
		"order",
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

	OrderClient = c
}

func CreateOrder(ctx context.Context, req *order.CreateOrderReq) (resp *mresp.CreateOrderResp, err error) {
	r, err := OrderClient.CreateOrder(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "order rpc CreateOrder failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.CreateOrderResp{OrderID: r.OrderId}, nil
}

func ViewOrderList(ctx context.Context, req *order.ViewOrderListReq) (resp *mresp.ViewOrderListResp, err error) {
	r, err := OrderClient.ViewOrderList(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "order rpc ViewOrderList failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.ViewOrderListResp{Orders: r.OrderList, Total: r.Total}, nil
}

func ViewOrderById(ctx context.Context, req *order.ViewOrderByIdReq) (resp *mresp.ViewOrderById, err error) {
	r, err := OrderClient.ViewOrderById(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "order rpc ViewOrderById failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.ViewOrderById{Order: *r.Order}, nil
}

func ConfirmReceiptOrder(ctx context.Context, req *order.ConfirmReceiptOrderReq) (resp *mresp.ConfirmReceiptOrderResp, err error) {
	r, err := OrderClient.ConfirmReceiptOrder(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "order rpc ConfirmReceiptOrder failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.ConfirmReceiptOrderResp{}, nil
}

func CancelOrder(ctx context.Context, req *order.CancelOrderReq) (resp *mresp.CancelOrderResp, err error) {
	r, err := OrderClient.CancelOrder(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "order rpc CancelOrder failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.CancelOrderResp{}, nil
}
