package base

import (
	"context"

	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/spf13/viper"
)

func InitTracing(serviceName string) func(context.Context) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(viper.GetString("otel.address")),
		provider.WithInsecure(),
	)

	return  func(ctx context.Context) {
		p.Shutdown(ctx)
	}
}