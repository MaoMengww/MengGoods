package usecase

import (
	"MengGoods/app/coupon/domain/model"
	"MengGoods/pkg/utils"
	"context"
)

func (u *CouponUsecase) CreateCouponBatch(ctx context.Context, batch *model.CouponBatch) (int64, error) {
	err := utils.Verify(utils.VerifyCouponName(batch.BatchName),
		utils.VerifyCouponType(batch.Type),
		utils.VerifyCouponThreshold(batch.Threshold),
		utils.VerifyDiscountAmount(batch.DiscountAmount),
		utils.VerifyDiscountPercent(batch.DiscountPercent),
		utils.VerifyTotalNum(batch.TotalNum),
		utils.VerifyCouponDuration(batch.Duration))
	if err != nil {
		return 0, err
	}
	return u.couponDB.CreateCouponBatch(ctx, batch)
}

func (u *CouponUsecase) GetCoupon(ctx context.Context, batchId int64) error {
	return u.service.ClaimCoupon(ctx, batchId)
}

func (u *CouponUsecase) GetCouponList(ctx context.Context, statu int) ([]*model.Coupon, error) {
	return u.couponDB.GetCouponList(ctx, statu)
}

func (u *CouponUsecase) LockCoupon(ctx context.Context, couponId int64) error {
	return u.couponDB.LockCoupon(ctx, couponId)
}

func (u *CouponUsecase) ReleaseCoupon(ctx context.Context, couponId int64) error {
	return u.couponDB.ReleaseCoupon(ctx, couponId)
}

func (u *CouponUsecase) RedeemCoupon(ctx context.Context, couponId int64, orderId int64) error {
	return u.couponDB.RedeemCoupon(ctx, couponId, orderId)
}

func (u *CouponUsecase) GetCouponInfo(ctx context.Context, couponId int64) (*model.Coupon, error) {
	return u.couponDB.GetCouponInfo(ctx, couponId)
}

func (u *CouponUsecase) LetCouponExpire(ctx context.Context, couponId int64) error {
	return u.couponDB.LetCouponExpire(ctx, couponId)
}
