package rpc

import (
	//	mresp "MengGoods/app/gateway/model/resp"
	mresp "MengGoods/app/gateway/model/resp"
	"MengGoods/kitex_gen/cart"
	"MengGoods/kitex_gen/cart/cartservice"
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

var CartClient cartservice.Client

func CartInit() {
	r, err := etcd.NewEtcdResolver(viper.GetStringSlice("etcd.endpoints"))
	if err != nil {
		logger.Fatalf("cart rpc Init Falied: err: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := cartservice.NewClient(
		"cart",
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

	CartClient = c
}

func AddCartItem(ctx context.Context, req *cart.AddCartItemReq) (resp *mresp.AddCartItemResp, err error) {
	r, err := CartClient.AddCartItem(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "cart rpc AddCartItem failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.AddCartItemResp{}, nil
}

func GetCartItem(ctx context.Context, req *cart.GetCartItemReq) (resp *mresp.GetCartItemResp, err error) {
	r, err := CartClient.GetCartItem(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "cart rpc GetCartItem failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.GetCartItemResp{
		CartItem: r.CartItems,
	}, nil
}

func DeleteCartItem(ctx context.Context, req *cart.DeleteCartItemReq) (resp *mresp.DeleteCartItemResp, err error) {
	r, err := CartClient.DeleteCartItem(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "cart rpc DeleteCartItem failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.DeleteCartItemResp{}, nil
}

func UpdateCartItem(ctx context.Context, req *cart.UpdateCartItemReq) (resp *mresp.UpdateCartItemResp, err error) {
	r, err := CartClient.UpdateCartItem(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "cart rpc UpdateCartItem failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.UpdateCartItemResp{}, nil
}
