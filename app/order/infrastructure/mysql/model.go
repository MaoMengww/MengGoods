package mysql

import "time"

type Orders struct {
	OrderId          int64      `gorm:"primaryKey"`
	UserId           int64      `gorm:"use_id"`
	OrderStatus      int        `gorm:"order_status"`
	TotalPrice       int64      `gorm:"total_price"`
	PayPrice         int64      `gorm:"pay_price"`
	ReceiverName     string     `gorm:"receiver_name"`
	ReceiverEmail    string     `gorm:"receiver_email"`
	ReceiverProvince string     `gorm:"receiver_province"`
	ReceiverCity     string     `gorm:"receiver_city"`
	ReceiverDetail   string     `gorm:"receiver_detail"`
	CreateTime       time.Time  `gorm:"create_time"`
	UpdateTime       time.Time  `gorm:"update_time"`
	ExpireTime       time.Time  `gorm:"expire_time"`
	CancelTime       *time.Time `gorm:"cancel_time"`
	CancelReason     string     `gorm:"cancel_reason"`
}

type OrderItem struct {
	OrderItemId       int64  `gorm:"primaryKey"`
	OrderId           int64  `gorm:"order_id"`
	ProductId         int64  `gorm:"product_id"`
	ProductName       string `gorm:"product_name"`
	ProductImage      string `gorm:"product_image"`
	ProductPrice      int64  `gorm:"product_price"`
	ProductNum        int    `gorm:"product_num"`
	ProductTotalPrice int64  `gorm:"product_total_price"`
	ProductProperties string `gorm:"product_properties"`
}

type MqMsg struct {
	MsgId      int64     `gorm:"primaryKey"`
	OrderId    int64     `gorm:"order_id"`
	CouponId   int64     `gorm:"coupon_id"`
	CreateTime time.Time `gorm:"create_time"`
	MsgStatus  int       `gorm:"msg_status"`
}
