package model

import "time"

type Order struct {
	OrderId          int64
	UserId           int64
	OrderStatus      int
	TotalPrice       int64
	PayPrice         int64
	ReceiverName     string
	ReceiverEmail    string
	ReceiverProvince string
	ReceiverCity     string
	ReceiverDetail   string
	CreateTime       time.Time
	UpdateTime       time.Time
	ExpireTime       time.Time
	CancelTime       *time.Time
	CancelReason     string
}

type OrderItem struct {
	OrderItemId       int64
	OrderId           int64
	SellerID          int64
	UserId            int64
	ProductId         int64
	ProductName       string
	ProductImage      string
	ProductPrice      int64
	ProductNum        int
	ProductTotalPrice int64
	ProductProperties string
}

type OrderWithItems struct {
	Order      *Order
	OrderItems []*OrderItem
}

type MqMsg struct {
	MsgId      int64
	OrderId    int64
	CouponId   int64
	CreateTime time.Time
	MsgStatus  int
}
