package prpc

import (
	"MengGoods/kitex_gen/user"
	"MengGoods/kitex_gen/user/userservice"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
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

type ProductRpc struct {
	userRpc userservice.Client
}

func NewProductRpc(userRpc userservice.Client) *ProductRpc {
	return &ProductRpc{
		userRpc: userRpc,
	}
}

func NewProductClient() userservice.Client {
	r, err := etcd.NewEtcdResolver(viper.GetStringSlice("etcd.endpoints"))
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

func (p *ProductRpc) IsAdmin(ctx context.Context) (bool, error) {
	uid, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	var req user.GetUserInfoReq
	req.UserId = uid
	resp, err := p.userRpc.GetUserInfo(ctx, &req)
	if err != nil {
		return false, err
	}
	return resp.UserInfo.Role == constants.Admin, nil
}
