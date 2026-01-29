package prpc

import (
	"MengGoods/kitex_gen/order/orderservice"
	"MengGoods/pkg/logger"
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



type PaymentRpc struct {
	orderClient orderservice.Client
}

func NewPaymentRpc(orderClient orderservice.Client) *PaymentRpc {
	return &PaymentRpc{
		orderClient: orderClient,
	}
}

func NewOrderClient() orderservice.Client {
	r, err := etcd.NewEtcdResolver(viper.GetStringSlice("etcd.endpoints"))
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
	)

	if err != nil {
		logger.Fatalf("init client failed: err:%v", err)
	}

	return c
}

