package service

import (
	"MengGoods/app/coupon/domain/model"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"time"
)

func (s *CouponService) CreateCoupon(ctx context.Context, batch *model.CouponBatch) error {
	batchId, err := s.CouponDB.CreateCouponBatch(ctx, batch)
	if err != nil {
		return err
	}
	key := s.CouponCache.GetCouponBatchKey(ctx, batchId)
	duration := time.Duration(batch.EndTime-batch.StartTime) * time.Second
	err = s.CouponCache.SetCoupon(ctx, key, batch.TotalNum, duration)
	if err != nil {
		return err
	}
	return nil
}

func (s *CouponService) LockCoupon(ctx context.Context, couponId int64) error {
	isUnused, err := s.CouponDB.IsUnusedCoupon(ctx, couponId)
	if err != nil {
		return err
	}
	if !isUnused {
		return merror.NewMerror(merror.CouponIsUsed, "coupon is used")
	}
	return s.CouponDB.LockCoupon(ctx, couponId)
}

func (s *CouponService) ClaimCoupon(ctx context.Context, batchId int64) error {
	key := s.CouponCache.GetCouponBatchKey(ctx, batchId)
	err := s.CouponCache.ClaimCoupon(ctx, key)
	if err != nil {
		return err
	}
	err = s.CouponMq.SendClaimCoupon(ctx, batchId)
	if err != nil {
		return err
	}
	return nil
}

func (s *CouponService) HandleClaimCoupon(ctx context.Context, userId int64, batchId int64) error {
	batch, err := s.CouponDB.GetCouponBatchByID(ctx, batchId)
	if err != nil {
		return err
	}
	expireTime := time.Now().Unix() + int64(batch.Duration)
	coupon := &model.Coupon{
		CouponId:       utils.GenerateID(),
		BatchId:        batch.BatchId,
		Type:           batch.Type,
		Threshold:      batch.Threshold,
		DiscountAmount: batch.DiscountAmount,
		DiscountRate:   batch.DiscountPercent,
		Status:         constants.CouponStatusUnused,
		CreatedAt:      batch.CreatedAt,
		ExpiredAt:      expireTime,
		UserId:         userId,
	}
	if time.Now().Unix() > expireTime{
		coupon.Status = constants.CouponStatusExpired
	}
	if err := s.CouponDB.CreateCoupon(ctx, coupon); err != nil {
		return err
	}
	return nil
}

func (s *CouponService) ConsumeClaimCoupon(ctx context.Context) error {
	return s.CouponMq.ConsumeClaimCoupon(ctx, s.HandleClaimCoupon)
}
