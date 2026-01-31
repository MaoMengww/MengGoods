package mq

import (
	"MengGoods/pkg/base/client"
	"MengGoods/pkg/merror"
	"context"
	"encoding/json"
)

type RabbitMQ struct {
	*client.RabbitMq
}

func NewRabbitMQ(client *client.RabbitMq) *RabbitMQ {
	return &RabbitMQ{
		RabbitMq: client,
	}
}

type OrderMessage struct {
	OrderId  int64 `json:"order_id"`
	CouponId int64 `json:"coupon_id"`
}

func (r *RabbitMQ) SendOrderMessage(ctx context.Context, orderId int64, couponId int64) error {
	msg := &OrderMessage{
		OrderId:  orderId,
		CouponId: couponId,
	}
	orderMsg, err := json.Marshal(msg)
	if err != nil {
		return merror.NewMerror(merror.InternalRabbitMqErrorCode, "json : "+err.Error())
	}
	return r.Publish(orderMsg)
}

func (r *RabbitMQ) ConsumeOrderMessage(ctx context.Context, fn func(ctx context.Context, orderId int64, couponId int64) error) error {
	return r.Consume(func(msg []byte) error {
		var orderMsg OrderMessage
		err := json.Unmarshal(msg, &orderMsg)
		if err != nil {
			return merror.NewMerror(merror.InternalRabbitMqErrorCode, "json解析错误: "+err.Error())
		}
		return fn(ctx, orderMsg.OrderId, orderMsg.CouponId)
	})
}
