package rpc

import (
	//	mresp "MengGoods/app/gateway/model/resp"
	"MengGoods/config"
	"MengGoods/kitex_gen/payment"
	"MengGoods/kitex_gen/payment/paymentservice"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"context"

	"time"

	mresp "MengGoods/app/gateway/model/resp"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var PaymentClient paymentservice.Client

func PaymentInit() {
	r, err := etcd.NewEtcdResolver(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("payment rpc Init Falied: err: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := paymentservice.NewClient(
		"payment",
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

	PaymentClient = c
}

func GetPaymentToken(ctx context.Context, req *payment.GetPaymentTokenReq) (resp *mresp.GetPaymentTokenResp, err error) {
	r, err := PaymentClient.GetPaymentToken(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "payment rpc GetPaymentToken failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.GetPaymentTokenResp{
		PaymentToken: r.PaymentToken,
		ExpiredAt:    r.ExpiredAt,
	}, nil
}

func Payment(ctx context.Context, req *payment.PaymentReq) (resp *mresp.PaymentResp, err error) {
	r, err := PaymentClient.Payment(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "payment rpc Payment failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.PaymentResp{}, nil
}

func PaymentRefund(ctx context.Context, req *payment.PaymentRefundReq) (resp *mresp.PaymentRefundResp, err error) {
	r, err := PaymentClient.PaymentRefund(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "payment rpc PaymentRefund failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.PaymentRefundResp{}, nil
}

func ReviewRefund(ctx context.Context, req *payment.ReviewRefundReq) (resp *mresp.ReviewRefundResp, err error) {
	r, err := PaymentClient.ReviewRefund(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "payment rpc ReviewRefund failed: err: %v", err)
		return nil, err
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return &mresp.ReviewRefundResp{}, nil
}
