package mq

import (
	"MengGoods/pkg/base/client"
	"MengGoods/pkg/merror"
	"context"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type CouponMq struct {
	kafka *client.Kafka
}

func NewCouponMq(kafka *client.Kafka) *CouponMq {
	return &CouponMq{
		kafka: kafka,
	}
}

func (p *CouponMq) SendClaimCoupon(ctx context.Context, couponId int64) error {
	return p.kafka.Publish(ctx, "coupon_claim", &kafka.Message{
		Key:   []byte(strconv.FormatInt(couponId, 10)),
		Value: nil,
	})
}

func (p *CouponMq) ConsumeClaimCoupon(ctx context.Context, fn func(ctx context.Context, couponId int64) error) error {
	if err := p.kafka.Consumer("coupon_claim", "coupon_claim_group", func(ctx context.Context, msg *kafka.Message) error {
		couponId, err := strconv.ParseInt(string(msg.Key), 10, 64)
		if err != nil {
			return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("parse coupon id error: %v", err))
		}
		return fn(ctx, couponId)
	}, ctx); err != nil {
		return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("consumer error: %v", err))
	}
	return nil
}
