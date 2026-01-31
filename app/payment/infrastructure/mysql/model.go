package mysql

type PaymentOrder struct {
	PaymentID int64  `gorm:"primaryKey"`
	PaymentNo string `gorm:"column:payment_no"`
	OrderId   int64  `gorm:"column:order_id"`
	UserId    int64  `gorm:"column:user_id"`
	Amount    int64  `gorm:"column:amount"`
	Method    int    `gorm:"column:method"`
	Status    int    `gorm:"column:status"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}

type PaymentRefund struct {
	RefundID    int64  `gorm:"primaryKey"`
	RefundNo    string `gorm:"column:refund_no"`
	OrderItemID int64  `gorm:"column:order_item_id"`
	SellerID    int64  `gorm:"column:seller_id"`
	UserId      int64  `gorm:"column:user_id"`
	Amount      int64  `gorm:"column:amount"`
	Reason      string `gorm:"column:reason"`
	Status      int    `gorm:"column:status"`
	CreatedAt   int64  `gorm:"column:created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at"`
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
