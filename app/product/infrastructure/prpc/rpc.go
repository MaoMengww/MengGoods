package prpc

import (
	"MengGoods/config"
	"MengGoods/kitex_gen/stock/stockservice"
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

type ProductRpc struct {
	userClient  userservice.Client
	stockClient stockservice.Client
}

func NewProductRpc(userClient userservice.Client, stockClient stockservice.Client) *ProductRpc {
	return &ProductRpc{
		userClient:  userClient,
		stockClient: stockClient,
	}
}

func NewUserClient() userservice.Client {
	r, err := etcd.NewEtcdResolver(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("calc rpc Init Falied: err: %v", err)
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

func NewStockClient() stockservice.Client {
	r, err := etcd.NewEtcdResolver(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("calc rpc Init Falied: err: %v", err)
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
	)

	if err != nil {
		logger.Fatalf("init client failed: err:%v", err)
	}

	return c
}
