package model

type CouponBatch struct {
	BatchId int64
	BatchName string 
	Remark string
	Type int
	Threshold int64
	DiscountAmount int64
	DiscountPercent int
	TotalNum int64
	StartTime int64
	EndTime int64
	Duration int
	CreatedAt int64
	UpdatedAt int64
}

type Coupon struct {
	CouponId int64
	BatchId int64
	OrderId int64
	UserId int64
	Status int
	CreatedAt int64
	ExpiredAt int64
	UsedAt int64
}


