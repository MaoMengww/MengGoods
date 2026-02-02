package mysql

import "time"

type PaymentOrder struct {
	PaymentID int64     `gorm:"primaryKey"`
	PaymentNo string    `gorm:"column:payment_no"`
	OrderId   int64     `gorm:"column:order_id"`
	UserId    int64     `gorm:"column:user_id"`
	Amount    int64     `gorm:"column:payment_amount"`
	Method    int       `gorm:"column:payment_method"`
	Status    int       `gorm:"column:payment_status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type PaymentRefund struct {
	RefundID    int64     `gorm:"primaryKey"`
	RefundNo    string    `gorm:"column:refund_no"`
	OrderItemID int64     `gorm:"column:order_item_id"`
	SellerID    int64     `gorm:"column:seller_id"`
	UserId      int64     `gorm:"column:user_id"`
	Amount      int64     `gorm:"column:refund_amount"`
	Reason      string    `gorm:"column:refund_reason"`
	Status      int       `gorm:"column:refund_status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type PaymentTransaction struct {
	TransactionID     int64  `gorm:"primaryKey"`
	TransactionNo     string `gorm:"column:transaction_no"`
	OrderId           int64  `gorm:"column:order_id"`
	UserId            int64  `gorm:"column:user_id"`
	TransactionAmount int64  `gorm:"column:transaction_amount"`
	Type              int    `gorm:"column:transaction_type"`
	CreatedAt         int64  `gorm:"column:created_at"`
}
