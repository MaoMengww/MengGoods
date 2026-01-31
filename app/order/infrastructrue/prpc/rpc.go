package prpc

import (
	"MengGoods/config"
	"MengGoods/kitex_gen/cart/cartservice"
	"MengGoods/kitex_gen/coupon/couponservice"
	"MengGoods/kitex_gen/product/productservice"
	"MengGoods/kitex_gen/user/userservice"
	"MengGoods/pkg/logger"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type OrderRpc struct {
	cartClient    cartservice.Client
	CouponClient  couponservice.Client
	ProductClient productservice.Client
	UserClient    userservice.Client
}

func NewOrderRpc(cartClient cartservice.Client, CouponClient couponservice.Client, ProductClient productservice.Client, UserClient userservice.Client) *OrderRpc {
	return &OrderRpc{
		cartClient:    cartClient,
		CouponClient:  CouponClient,
		ProductClient: ProductClient,
		UserClient:    UserClient,
	}
}

func NewCartClient() cartservice.Client {
	r, err := etcd.NewEtcdResolver(config.Conf.Etcd.Endpoints)
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
	)

	if err != nil {
		logger.Fatalf("init client failed: err:%v", err)
	}

	return c
}

func NewCouponClient() couponservice.Client {
	r, err := etcd.NewEtcdResolver(config.Conf.Etcd.Endpoints)
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
	)

	if err != nil {
		logger.Fatalf("init client failed: err:%v", err)
	}

	return c
}

func NewProductClient() productservice.Client {
	r, err := etcd.NewEtcdResolver(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("product rpc Init Falied: err: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := productservice.NewClient(
		"product",
		client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithCircuitBreaker(cbSuite),
		client.WithSuite(tracing.NewClientSuite()),
	)

	if err != nil {
		logger.Fatalf("init client failed: err:%v", err)
	}

	return c
}

func NewUserClient() userservice.Client {
	r, err := etcd.NewEtcdResolver(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("user rpc Init Falied: err: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := userservice.NewClient(
		"user",
		client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithCircuitBreaker(cbSuite),
		client.WithSuite(tracing.NewClientSuite()),
	)

	if err != nil {
		logger.Fatalf("init client failed: err:%v", err)
	}

	return c
}
