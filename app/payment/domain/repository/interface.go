package repository

import (
	domain"MengGoods/app/payment/domain/model"
	"MengGoods/kitex_gen/model"
	"context"
)

type PaymentDB interface {
	CreatePaymentOrder(ctx context.Context, paymentOrder *domain.PaymentOrder) error
	CreateRefundOrder(ctx context.Context, refundOrder *domain.PaymentRefund) error
	GetPaymentOrderByOrderId(ctx context.Context, orderId int64) (*domain.PaymentOrder, error)
	ConfirmPaymentOrder(ctx context.Context, paymentOrder *domain.PaymentOrder) error //事务操作，更新支付订单状态以及创建流水记录
	ConfirmRefundOrder(ctx context.Context, refundOrder *domain.PaymentRefund) error //事务操作，更新退款订单状态以及创建流水记录
	GetRefundOrderByOrderItemId(ctx context.Context, orderItemId int64) (*domain.PaymentRefund, error)
	UpdatePaymentOrderStatus(ctx context.Context, paymentNo string, status int) error
	UpdateRefundOrderStatus(ctx context.Context, orderItemId int64, status int) error
}	

type PaymentCache interface {
	GetPaymentKey(ctx context.Context, orderId int64) string
	GetRefundKey(ctx context.Context, orderItemId int64) string
	GetExpiredDuration(ctx context.Context, expiredTime int64) (int64, error)
	SetPaymentToken(ctx context.Context, key string, token string, expire int64) error
	GetTTLAndDelPaymentToken(ctx context.Context, key string, token string) (bool, int64, error)
	SetOrIncrRefundKey(ctx context.Context, key string) (int64, error) 
	SetDailyRefund(ctx context.Context, key string) error
	CheckDailyRefundCount(ctx context.Context, key string) (bool, error)
}

type PaymentRpc interface {
	IsOrderExist(ctx context.Context, orderId int64) (bool, int64, error)
	MarkOrderPaid(ctx context.Context, orderId int64) error
	GetOrderInfoById(ctx context.Context, orderId int64) (*model.OrderInfo, error)
	GetOrderItemInfoById(ctx context.Context, orderItemId int64) (*model.OrderItem, error)
}
