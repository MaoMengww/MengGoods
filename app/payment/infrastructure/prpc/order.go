package prpc

import (
	"MengGoods/kitex_gen/model"
	"MengGoods/kitex_gen/order"
	"MengGoods/pkg/merror"
	"context"
)

func (p *PaymentRpc) IsOrderExist(ctx context.Context, orderId int64) (bool, int64, error) {
	resp, err := p.orderClient.IsOrderExist(ctx, &order.IsOrderExistReq{
		OrderId: orderId,
	})
	if err != nil {
		return false, 0, err
	}
	if resp.Base.Code != merror.SuccessCode {
		return false, 0, merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return resp.Exist, resp.ExpiredAt, nil
}

func (p *PaymentRpc) MarkOrderPaid(ctx context.Context, orderId int64) error {
	resp, err := p.orderClient.MarkOrderPaid(ctx, &order.MarkOrderPaidReq{
		OrderId: orderId,
	})
	if err != nil {
		return err
	}
	if resp.Base.Code != merror.SuccessCode {
		return merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return nil
}

func (p *PaymentRpc) GetOrderInfoById(ctx context.Context, orderId int64) (*model.OrderInfo, error) {
	resp, err := p.orderClient.GetOrderInfo(ctx, &order.GetOrderInfoReq{
		OrderId: orderId,
	})
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return resp.OrderInfo, nil
}

func (p *PaymentRpc) GetOrderItemInfoById(ctx context.Context, orderItemId int64) (*model.OrderItem, error) {
	resp, err := p.orderClient.GetOrderItem(ctx, &order.GetOrderItemReq{
		OrderItemId: orderItemId,
	})
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return resp.OrderItem, nil
}
