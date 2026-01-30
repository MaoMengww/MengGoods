package resp

import "MengGoods/kitex_gen/model"

type CreateOrderResp struct {
	OrderID int64 `json:"orderId"`
}

type ViewOrderById struct {
	Order model.OrderWithItems `json:"order"`
}

type ViewOrderListResp struct {
	Orders []*model.OrderInfo `json:"orders"`
	Total  int64              `json:"total"`
}

type ConfirmReceiptOrderResp struct {
}

type CancelOrderResp struct {
}
