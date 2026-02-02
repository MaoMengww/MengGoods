package mysql

import "time"

type Orders struct {
	OrderId          int64      `gorm:"primaryKey"`
	UserId           int64      `gorm:"column:user_id"`
	OrderStatus      int        `gorm:"column:order_status"`
	TotalPrice       int64      `gorm:"column:total_price"`
	PayPrice         int64      `gorm:"column:pay_price"`
	ReceiverName     string     `gorm:"column:receiver_name"`
	ReceiverEmail    string     `gorm:"column:receiver_email"`
	ReceiverProvince string     `gorm:"column:receiver_province"`
	ReceiverCity     string     `gorm:"column:receiver_city"`
	ReceiverDetail   string     `gorm:"column:receiver_detail"`
	CreateTime       time.Time  `gorm:"column:create_time"`
	UpdateTime       time.Time  `gorm:"column:update_time"`
	ExpireTime       time.Time  `gorm:"column:expire_time"`
	CancelTime       *time.Time `gorm:"column:cancel_time"`
	CancelReason     string     `gorm:"column:cancel_reason"`
}

type OrderItem struct {
	OrderItemId       int64  `gorm:"primaryKey;column:order_item_id"`
	OrderId           int64  `gorm:"column:order_id"`
	SellerId          int64  `gorm:"column:seller_id"`
	ProductId         int64  `gorm:"column:product_id"`
	ProductName       string `gorm:"column:product_name"`
	ProductImage      string `gorm:"column:product_img"`
	ProductPrice      int64  `gorm:"column:product_price"`
	ProductNum        int    `gorm:"column:product_num"`
	ProductTotalPrice int64  `gorm:"column:product_total_price"`
	ProductProperties string `gorm:"column:product_properties"`
}

type MqMsg struct {
	MsgId      int64     `gorm:"primaryKey;column:msg_id"`
	OrderId    int64     `gorm:"column:order_id"`
	CouponId   int64     `gorm:"column:coupon_id"`
	CreateTime time.Time `gorm:"column:create_time"`
	MsgStatus  int       `gorm:"column:msg_status"`
}
