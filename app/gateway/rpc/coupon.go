package rpc

import (
	//	mresp "MengGoods/app/gateway/model/resp"
	mresp "MengGoods/app/gateway/model/resp"
	"MengGoods/kitex_gen/coupon"
	"MengGoods/kitex_gen/coupon/couponservice"
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

var CouponClient couponservice.Client

func CouponInit() {
	r, err := etcd.NewEtcdResolver(viper.GetStringSlice("etcd.endpoints"))
	if err != nil {
		logger.Fatalf("coupon rpc Init Falied: err: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := couponservice.NewClient(
		"coupon",
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

	CouponClient = c
}

func GetCouponInfo(ctx context.Context, req *coupon.GetCouponInfoReq) (resp *mresp.GetCouponInfoResp, err error) {
	r, err := CouponClient.GetCouponInfo(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "coupon rpc GetCouponInfo failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}

func CreateCouponBatch(ctx context.Context, req *coupon.CreateCouponBatchReq) (resp *mresp.CreateCouponBatchResp, err error) {
	r, err := CouponClient.CreateCouponBatch(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "coupon rpc CreateCouponBatch failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}

func GetCoupon(ctx context.Context, req *coupon.GetCouponReq) (resp *mresp.GetCouponResp, err error) {
	r, err := CouponClient.GetCoupon(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "coupon rpc GetCoupon failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}

func GetCouponList(ctx context.Context, req *coupon.GetCouponListReq) (resp *mresp.GetCouponListResp, err error) {
	r, err := CouponClient.GetCouponList(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "coupon rpc GetCouponList failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}
