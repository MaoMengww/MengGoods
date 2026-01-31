package mysql

import "time"

type CouponBatch struct {
	BatchId        int64     `gorm:"primaryKey;column:batch_id"`
	BatchName      string    `gorm:"column:batch_name"`
	Remark         string    `gorm:"column:remark"`
	Type           int       `gorm:"column:type"`
	Threshold      int64     `gorm:"column:threshold"`
	DiscountAmount int64     `gorm:"column:discount_amount"`
	DiscountRate   int       `gorm:"column:discount_rate"`
	TotalNum       int       `gorm:"column:total_num"`
	UsedNum        int       `gorm:"column:used_num"`
	StartTime      time.Time `gorm:"column:start_time"`
	EndTime        time.Time `gorm:"column:end_time"`
	Duration       int       `gorm:"column:duration"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}

type Coupon struct {
	CouponId       int64     `gorm:"primaryKey;column:coupon_id"`
	BatchId        int64     `gorm:"column:batch_id"`
	OrderId        int64     `gorm:"column:user_order_id"`
	UserId         int64     `gorm:"column:user_id"`
	Type           int       `gorm:"column:type"`
	Threshold      int64     `gorm:"column:threshold"`
	DiscountAmount int64     `gorm:"column:discount_amount"`
	DiscountRate   int       `gorm:"column:discount_rate"`
	Status         int       `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	ExpiredAt      time.Time `gorm:"column:expired_at"`
	UsedAt         time.Time `gorm:"column:used_at"`
}
