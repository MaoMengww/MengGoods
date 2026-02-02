package mw

import (
	"MengGoods/pkg/logger"
	"context"
	"sync"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/opensergo/sentinel/adapter"
)

var once sync.Once

func Sentinel() app.HandlerFunc {
	once.Do(initSentinel)
	return adapter.SentinelServerMiddleware(
		adapter.WithServerResourceExtractor(func(ctx context.Context, rc *app.RequestContext) string {
			return "gateway"
		}),
		adapter.WithServerBlockFallback(func(ctx context.Context, rc *app.RequestContext) {
			logger.Warnf("request has been rejected by the getway, client.ip:%v\n", rc.ClientIP())
			rc.AbortWithStatusJSON(200, map[string]interface{}{
				"code":    200,
				"message": "服务器当前繁忙,请稍后尝试",
			})
		}),
	)
}

func initSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		logger.Fatalf("sentinel init failed: err: %v", err)
	}

	//limit qps to 100
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "gateway",
			Threshold:              100,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			StatIntervalInMs:       1000,
		},
	})

	if err != nil {
		logger.Fatalf("sentinel init failed: err: %v", err)
	}

}
