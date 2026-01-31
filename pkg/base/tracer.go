package base

import (
	"MengGoods/config"
	"context"

	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) func(context.Context) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(config.Conf.Otel.Address),
		provider.WithInsecure(),
	)

	return func(ctx context.Context) {
		p.Shutdown(ctx)
	}
}
