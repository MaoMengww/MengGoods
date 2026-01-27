package pack

import (
	mModel "MengGoods/app/order/domain/model"
	"MengGoods/kitex_gen/model"
	"time"
)

func ToDomainOrder(order *model.OrderInfo) *mModel.Order {
	cancelTime := time.Unix(order.CancelTime, 0)
	return &mModel.Order{
		OrderId:          order.Id,
		UserId:           order.UserId,
		OrderStatus:      int(order.Status),
		TotalPrice:       order.TotalPrice,
		PayPrice:         order.PaymentPrice,
		ReceiverName:     order.ReceiverName,
		ReceiverEmail:    order.ReceiverEmail,
		ReceiverProvince: order.ReceiverProvince,
		ReceiverCity:     order.ReceiverCity,
		ReceiverDetail:   order.ReceiverDetail,
		CreateTime:       time.Unix(order.CreateTime, 0),
		UpdateTime:       time.Unix(order.UpdateTime, 0),
		ExpireTime:       time.Unix(order.ExpireTime, 0),
		CancelTime:       &cancelTime,
		CancelReason:     order.CancelReason,
	}
}

func ToRpcOrder(order *mModel.Order) *model.OrderInfo {
	return &model.OrderInfo{
		Id:               order.OrderId,
		UserId:           order.UserId,
		Status:           int32(order.OrderStatus),
		TotalPrice:       order.TotalPrice,
		PaymentPrice:     order.PayPrice,
		ReceiverName:     order.ReceiverName,
		ReceiverEmail:    order.ReceiverEmail,
		ReceiverProvince: order.ReceiverProvince,
		ReceiverCity:     order.ReceiverCity,
		ReceiverDetail:   order.ReceiverDetail,
		CreateTime:       order.CreateTime.Unix(),
		UpdateTime:       order.UpdateTime.Unix(),
		ExpireTime:       order.ExpireTime.Unix(),
		CancelTime:       order.CancelTime.Unix(),
		CancelReason:     order.CancelReason,
	}
}

func ToDomainOrderItem(orderItem *model.OrderItem) *mModel.OrderItem {
	return &mModel.OrderItem{
		OrderId:           orderItem.OrderId,
		OrderItemId:       orderItem.Id,
		ProductId:         orderItem.ProductId,
		ProductImage:      orderItem.ProductImage,
		ProductName:       orderItem.ProductName,
		ProductNum:        int(orderItem.ProductNum),
		ProductPrice:      orderItem.ProductPrice,
		ProductTotalPrice: orderItem.ProductTotalPrice,
		ProductProperties: orderItem.ProductProperties,
	}
}

func ToRpcOrderItem(orderItem *mModel.OrderItem) *model.OrderItem {
	return &model.OrderItem{
		OrderId:           orderItem.OrderId,
		Id:                orderItem.OrderItemId,
		ProductId:         orderItem.ProductId,
		ProductImage:      orderItem.ProductImage,
		ProductName:       orderItem.ProductName,
		ProductNum:        int64(orderItem.ProductNum),
		ProductPrice:      orderItem.ProductPrice,
		ProductTotalPrice: orderItem.ProductTotalPrice,
		ProductProperties: orderItem.ProductProperties,
	}
}

func ToRpcOrderItemList(orderItemList []*mModel.OrderItem) []*model.OrderItem {
	var orderItemModelList []*model.OrderItem
	for _, orderItem := range orderItemList {
		orderItemModelList = append(orderItemModelList, ToRpcOrderItem(orderItem))
	}
	return orderItemModelList
}

func ToDomainOrderItemList(orderItemList []*model.OrderItem) []*mModel.OrderItem {
	var orderItemModelList []*mModel.OrderItem
	for _, orderItem := range orderItemList {
		orderItemModelList = append(orderItemModelList, ToDomainOrderItem(orderItem))
	}
	return orderItemModelList
}

func ToRpcOrderWithItems(orderWithItems *mModel.OrderWithItems) *model.OrderWithItems {
	return &model.OrderWithItems{
		OrderInfo:  ToRpcOrder(orderWithItems.Order),
		OrderItems: ToRpcOrderItemList(orderWithItems.OrderItems),
	}
}

func ToRpcOrderList(orderList []*mModel.Order) []*model.OrderInfo {
	var orderModelList []*model.OrderInfo
	for _, order := range orderList {
		orderModelList = append(orderModelList, ToRpcOrder(order))
	}
	return orderModelList
}
