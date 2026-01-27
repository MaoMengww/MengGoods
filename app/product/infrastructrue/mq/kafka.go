package mq

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/pkg/base/client"
	"MengGoods/pkg/merror"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type ProductMq struct {
	*client.Kafka
}

func NewProductMq(client *client.Kafka) *ProductMq {
	return &ProductMq{
		Kafka: client,
	}
}

func (p *ProductMq) SendCreateSpuInfo(ctx context.Context, spu *model.SpuEs) error {
	msg, err := json.Marshal(spu)
	if err != nil {
		return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("json marshal error: %v", err))
	}
	return p.Publish(ctx, "product_create_spu", &kafka.Message{
		Key:   []byte(strconv.FormatInt(spu.Id, 10)),
		Value: msg,
	})
}

func (p *ProductMq) SendUpdateSpuInfo(ctx context.Context, spu *model.SpuEs) error {
	msg, err := json.Marshal(spu)
	if err != nil {
		return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("json marshal error: %v", err))
	}
	return p.Publish(ctx, "product_update_spu", &kafka.Message{
		Key:   []byte(strconv.FormatInt(spu.Id, 10)),
		Value: msg,
	})
}

func (p *ProductMq) SendDeleteSpuInfo(ctx context.Context, id int64) error {
	return p.Publish(ctx, "product_delete_spu", &kafka.Message{
		Key: []byte(strconv.FormatInt(id, 10)),
	})
}

func (p *ProductMq) ConsumeCreateSpuInfo(ctx context.Context, fn func(ctx context.Context, spu *model.SpuEs) error) error {
	if err := p.Consumer("product_create_spu", "product_create_spu_group", func(ctx context.Context, msg *kafka.Message) error {
		var spu model.SpuEs
		if err := json.Unmarshal(msg.Value, &spu); err != nil {
			return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("json unmarshal error: %v", err))
		}
		return fn(ctx, &spu)
	}, ctx); err != nil {
		return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("consumer error: %v", err))
	}
	return nil
}

func (p *ProductMq) ConsumeUpdateSpuInfo(ctx context.Context, fn func(ctx context.Context, spu *model.SpuEs) error) error {
	if err := p.Consumer("product_update_spu", "product_update_spu_group", func(ctx context.Context, msg *kafka.Message) error {
		var spu model.SpuEs
		if err := json.Unmarshal(msg.Value, &spu); err != nil {
			return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("json unmarshal error: %v", err))
		}
		return fn(ctx, &spu)
	}, ctx); err != nil {
		return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("consumer error: %v", err))
	}
	return nil
}

func (p *ProductMq) ConsumeDeleteSpuInfo(ctx context.Context, fn func(ctx context.Context, id int64) error) error {
	if err := p.Consumer("product_delete_spu", "product_delete_spu_group", func(ctx context.Context, msg *kafka.Message) error {
		id, err := strconv.ParseInt(string(msg.Key), 10, 64)
		if err != nil {
			return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("parse id error: %v", err))
		}
		return fn(ctx, id)
	}, ctx); err != nil {
		return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("consumer error: %v", err))
	}
	return nil
}
