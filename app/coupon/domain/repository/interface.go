package repository

import (
	"MengGoods/app/coupon/domain/model"
	"context"
	"time"
)

type CouponDB interface {
	IsUnusedCoupon(ctx context.Context, couponId int64) (bool, error)
	GetCouponInfo(ctx context.Context, couponId int64) (*model.Coupon, error)
	GetCouponBatchByID(ctx context.Context, batchId int64) (*model.CouponBatch, error)
	CreateCouponBatch(ctx context.Context, batch *model.CouponBatch) (int64, error)
	CreateCoupon(ctx context.Context, coupon *model.Coupon) error
	UpdateCouponStatus(ctx context.Context, couponId int64, status int) error
	GetCouponList(ctx context.Context, status int) ([]*model.Coupon, error)
	LockCoupon(ctx context.Context, couponId int64) error
	ReleaseCoupon(ctx context.Context, couponId int64) error
	RedeemCoupon(ctx context.Context, couponId int64, orderId int64) error
	LetCouponExpire(ctx context.Context, couponId int64) error
}

type CouponCache interface {
	GetCouponBatchKey(ctx context.Context, batchId int64) string
	SetCoupon(ctx context.Context, key string, totalNum int64, duration time.Duration) error
	ClaimCoupon(ctx context.Context, key string) error
	AddCoupon(ctx context.Context, key string) error
}

type CouponMq interface {
	SendClaimCoupon(ctx context.Context, batchId int64) error
	ConsumeClaimCoupon(ctx context.Context, fn func(ctx context.Context, batchId int64) error) error
}
