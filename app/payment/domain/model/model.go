package model

type PaymentOrder struct {
	PaymentId int64
	PaymentNo string
	OrderId   int64
	UserId    int64
	Amount    int64
	Method    int
	Status    int
}

type PaymentRefund struct {
	RefundId    int64
	OrderItemId int64
	PaymentNo   string
	RefundNo    string
	SellerId    int64
	UserId      int64
	Amount      int64
	Reason      string
	Status      int
}

type PaymentTransaction struct {
	TransactionId     int64
	PaymentNo         int64
	OrderId           int64
	SellerId          int64
	UserId            int64
	TransactionAmount int64
	TransactionType   int
}
