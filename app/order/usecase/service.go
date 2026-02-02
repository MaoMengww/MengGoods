package usecase

import (
	"MengGoods/app/order/domain/model"
	"MengGoods/pkg/constants"
	"context"
)

func (u *Usecase) CreateOrder(ctx context.Context, addressId int64, couponId int64, orderItems []*model.OrderItem) (int64, error) {
	return u.service.CreateOrder(ctx, addressId, couponId, orderItems)
}

func (u *Usecase) ViewOrderById(ctx context.Context, orderId int64) (*model.OrderWithItems, error) {
	return u.OrderDB.ViewOrderById(ctx, orderId)
}

func (u *Usecase) ViewOrderList(ctx context.Context, status int, pageNum int, pageSize int) ([]*model.Order, error) {
	return u.OrderDB.ViewOrderList(ctx, status, pageNum, pageSize)
}

func (u *Usecase) CancelOrder(ctx context.Context, orderId int64) error {
	return u.OrderDB.UpdateOrderStatus(ctx, orderId, constants.OrderCanceled)
}

func (u *Usecase) ConfirmReceipt(ctx context.Context, orderId int64) error {
	return u.OrderDB.UpdateOrderStatus(ctx, orderId, constants.OrderConfirmed)
}

func (u *Usecase) MarkOrderPaid(ctx context.Context, orderId int64) error {
	return u.OrderDB.UpdateOrderStatus(ctx, orderId, constants.OrderPaid)
}

func (u *Usecase) GetPayAmount(ctx context.Context, orderId int64) (int64, error) {
	return u.OrderDB.GetPayAmount(ctx, orderId)
}

func (u *Usecase) IsOrderExist(ctx context.Context, orderId int64) (bool, int64, error) {
	return u.service.IsOrderExist(ctx, orderId)
}

func (u *Usecase) GetOrderItem(ctx context.Context, orderItemId int64) (*model.OrderItem, error) {
	item, err := u.OrderDB.GetOrderItem(ctx, orderItemId)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (u *Usecase) GetOrderInfo(ctx context.Context, orderId int64) (*model.Order, error) {
	return u.OrderDB.GetOrderInfo(ctx, orderId)
}
