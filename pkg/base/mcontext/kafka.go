package mcontext

/* import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type KafkaHeaderAdapter []kafka.Header

func (k KafkaHeaderAdapter) Get(key string) string {
	for _, h := range k {
		if h.Key == key {
			return string(h.Value)
		}
	}
	return ""
}

func (k KafkaHeaderAdapter) Set(key string, value string) {
}

func (k KafkaHeaderAdapter) Keys() []string {
	keys := make([]string, len(k))
	for i, h := range k {
		keys[i] = h.Key
	}
	return keys
}

func GetCtxFromKafka(msg *kafka.Message) context.Context {
	var traceID string
	if msg.Headers != nil {
		for _, h := range msg.Headers {
			if h.Key == "trace_id" {
				traceID = string(h.Value)
				break
			}
		}
	}
	if traceID == "" {
		return context.Background()
	}
	ctx := metainfo.WithPersistentValue(context.Background(), "trace_id", traceID)
	mctx := otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier{msg.Headers})
	return mctx
} */