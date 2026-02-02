package usecase

import (
	"MengGoods/app/coupon/domain/model"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"time"
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
	batchId, err := u.couponDB.CreateCouponBatch(ctx, batch)
	if err != nil {
		return 0, err
	}
	durantion := batch.EndTime - time.Now().Unix()
	if durantion < 0 {
		return 0, merror.NewMerror(merror.CouponExpired, "coupon is expired")
	}
	if int64(batch.Duration) < durantion {
		durantion = int64(batch.Duration)
	}
	key := u.couponCache.GetCouponBatchKey(ctx, batchId)
	if err := u.couponCache.SetCoupon(ctx, key, batch.TotalNum, time.Duration(durantion)*time.Second); err != nil {
		return 0, err
	}
	return batchId, nil
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
