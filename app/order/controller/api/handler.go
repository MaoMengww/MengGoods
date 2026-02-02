package api

import (
	"MengGoods/app/order/controller/api/pack"
	"MengGoods/app/order/usecase"
	order "MengGoods/kitex_gen/order"
	"MengGoods/pkg/base"
	"context"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct {
	usecase *usecase.Usecase
}

func NewOrderServiceImpl(usecase *usecase.Usecase) *OrderServiceImpl {
	return &OrderServiceImpl{usecase: usecase}
}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	resp = new(order.CreateOrderResp)
	Items := pack.ToDomainOrderItemList(req.Items)
	orderId, err := s.usecase.CreateOrder(ctx, req.AddressId, req.CouponId, Items)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	resp.OrderId = orderId
	return resp, nil
}

// ViewOrderById implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ViewOrderById(ctx context.Context, req *order.ViewOrderByIdReq) (resp *order.ViewOrderByIdResp, err error) {
	resp = new(order.ViewOrderByIdResp)
	orderWithItems, err := s.usecase.ViewOrderById(ctx, req.OrderId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Order = pack.ToRpcOrderWithItems(orderWithItems)
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// ViewOrderList implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ViewOrderList(ctx context.Context, req *order.ViewOrderListReq) (resp *order.ViewOrderListResp, err error) {
	// TODO: Your code here...
	resp = new(order.ViewOrderListResp)
	orderList, err := s.usecase.ViewOrderList(ctx, int(req.Status), int(req.PageNum), int(req.PageSize))
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.OrderList = pack.ToRpcOrderList(orderList)
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// CancelOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, req *order.CancelOrderReq) (resp *order.CancelOrderResp, err error) {
	resp = new(order.CancelOrderResp)
	err = s.usecase.CancelOrder(ctx, req.OrderId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// ConfirmReceiptOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ConfirmReceiptOrder(ctx context.Context, req *order.ConfirmReceiptOrderReq) (resp *order.ConfirmReceiptOrderResp, err error) {
	// TODO: Your code here...
	resp = new(order.ConfirmReceiptOrderResp)
	err = s.usecase.ConfirmReceipt(ctx, req.OrderId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// TODO: Your code here...
	resp = new(order.MarkOrderPaidResp)
	err = s.usecase.MarkOrderPaid(ctx, req.OrderId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// GetPayAmount implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetPayAmount(ctx context.Context, req *order.GetPayAmountReq) (resp *order.GetPayAmountResp, err error) {
	resp = new(order.GetPayAmountResp)
	amount, err := s.usecase.GetPayAmount(ctx, req.OrderId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Amount = amount
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

func (s *OrderServiceImpl) GetOrderItem(ctx context.Context, req *order.GetOrderItemReq) (resp *order.GetOrderItemResp, err error) {
	resp = new(order.GetOrderItemResp)
	orderItem, err := s.usecase.GetOrderItem(ctx, req.OrderItemId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.OrderItem = pack.ToRpcOrderItem(orderItem)
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

func (s *OrderServiceImpl) GetOrderInfo(ctx context.Context, req *order.GetOrderInfoReq) (resp *order.GetOrderInfoResp, err error) {
	resp = new(order.GetOrderInfoResp)
	orderInfo, err := s.usecase.GetOrderInfo(ctx, req.OrderId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.OrderInfo = pack.ToRpcOrder(orderInfo)
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

func (s *OrderServiceImpl) IsOrderExist(ctx context.Context, req *order.IsOrderExistReq) (resp *order.IsOrderExistResp, err error) {
	resp = new(order.IsOrderExistResp)
	exist, expiredAt, err := s.usecase.IsOrderExist(ctx, req.OrderId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Exist = exist
	resp.ExpiredAt = expiredAt
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}
