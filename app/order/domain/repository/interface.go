package repository

import (
	domain "MengGoods/app/order/domain/model"
	"MengGoods/kitex_gen/model"
	"context"
)

type OrderDB interface {
	GetOrderStatus(ctx context.Context, orderId int64) (int, error)
	UpdateOrderStatus(ctx context.Context, orderId int64, status int) error
	//注意CreateOrder是一个事务，包含创建订单和存储待发送的mqMsg(解决双写不一致)
	CreateOrder(ctx context.Context, order *domain.Order, orderItem []*domain.OrderItem) error
	ViewOrderById(ctx context.Context, orderId int64) (*domain.OrderWithItems, error)
	ViewOrderList(ctx context.Context, status int, pageNum int, pageSize int) ([]*domain.Order, error)
	GetPendingMsgs(ctx context.Context) ([]*domain.MqMsg, error)
	MarkMsg(ctx context.Context, orderId int64) error
	GetPayAmount(ctx context.Context, orderId int64) (int64, error)
}

type OrderMq interface {
	SendOrderMessage(ctx context.Context, orderId int64, couponId int64) error
	ConsumeOrderMessage(ctx context.Context, fn func(ctx context.Context, orderId int64, couponId int64) error) error
}

type OrderRpc interface {
	GetUserInfo(ctx context.Context, userId int64) (*model.UserInfo, error)
	GetAddressInfo(ctx context.Context, addressId int64) (*model.AddressInfo, error)
	GetSkuInfo(ctx context.Context, skuId int64) (*model.SkuInfo, error)
	GetCouponInfo(ctx context.Context, couponId int64) (*model.CouponInfo, error)
	LockCoupon(ctx context.Context, couponId int64) error
	ReleaseCoupon(ctx context.Context, couponId int64) error
	RedeemCoupon(ctx context.Context, couponId int64) error
	LetCouponExpire(ctx context.Context, couponId int64) error
	GetCartItems(ctx context.Context) ([]*model.CartItem, error)
}
