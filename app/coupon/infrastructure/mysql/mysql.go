package mysql

import (
	"MengGoods/app/coupon/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"time"

	"gorm.io/gorm"
)

type CouponDB struct {
	db *gorm.DB
}

func NewCouponDB(db *gorm.DB) *CouponDB {
	return &CouponDB{
		db: db,
	}
}

func (c *CouponDB) GetCouponInfo(ctx context.Context, couponId int64) (*model.Coupon, error) {
	var couponDB Coupon
	err := c.db.WithContext(ctx).Where("coupon_id = ?", couponId).First(&couponDB).Error
	if err != nil {
		return nil, err
	}
	coupon := model.Coupon{
		CouponId:       couponDB.CouponId,
		BatchId:        couponDB.BatchId,
		OrderId:        couponDB.OrderId,
		UserId:         couponDB.UserId,
		Type:           couponDB.Type,
		Threshold:      couponDB.Threshold,
		DiscountAmount: couponDB.DiscountAmount,
		DiscountRate:   couponDB.DiscountRate,
		Status:         couponDB.Status,
		CreatedAt:      couponDB.CreatedAt.Unix(),
		ExpiredAt:      couponDB.ExpiredAt.Unix(),
		UsedAt:         couponDB.UsedAt.Unix(),
	}
	return &coupon, nil
}

func (c *CouponDB) GetCouponBatchByID(ctx context.Context, batchId int64) (*model.CouponBatch, error) {
	var couponBatch CouponBatch
	err := c.db.WithContext(ctx).Where("batch_id = ?", batchId).First(&couponBatch).Error
	if err != nil {
		return nil, err
	}
	coupon := model.CouponBatch{
		BatchId:         couponBatch.BatchId,
		BatchName:       couponBatch.BatchName,
		CreatedAt:       couponBatch.CreatedAt.Unix(),
		DiscountAmount:  couponBatch.DiscountAmount,
		DiscountPercent: couponBatch.DiscountRate,
		Duration:        couponBatch.Duration,
		EndTime:         couponBatch.EndTime.Unix(),
		Remark:          couponBatch.Remark,
		Type:            couponBatch.Type,
		Threshold:       couponBatch.Threshold,
	}
	return &coupon, nil
}

func (c *CouponDB) CreateCouponBatch(ctx context.Context, batch *model.CouponBatch) (int64, error) {
	CouponBatch := &CouponBatch{
		BatchName:      batch.BatchName,
		Remark:         batch.Remark,
		Type:           batch.Type,
		Threshold:      batch.Threshold,
		DiscountAmount: batch.DiscountAmount,
		DiscountRate:   batch.DiscountPercent,
		TotalNum:       int(batch.TotalNum),
		StartTime:      time.Unix(batch.StartTime, 0),
		EndTime:        time.Unix(batch.EndTime, 0),
		Duration:       batch.Duration,
	}
	err := c.db.WithContext(ctx).Create(CouponBatch).Error
	if err != nil {
		return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return CouponBatch.BatchId, nil
}

func (c *CouponDB) IsUnusedCoupon(ctx context.Context, couponId int64) (bool, error) {
	var coupon Coupon
	err := c.db.WithContext(ctx).Where("coupon_id = ?", couponId).First(&coupon).Error
	if err != nil {
		return false, err
	}
	return coupon.Status == constants.CouponStatusUnused, nil
}

func (c *CouponDB) CreateCoupon(ctx context.Context, coupon *model.Coupon) error {
	couponId := utils.GenerateID()
	couponDB := &Coupon{
		BatchId:        coupon.BatchId,
		CouponId:       couponId,
		CreatedAt:      time.Unix(coupon.CreatedAt, 0),
		UserId:         coupon.UserId,
		Type:           coupon.Type,
		Threshold:      coupon.Threshold,
		DiscountAmount: coupon.DiscountAmount,
		DiscountRate:   coupon.DiscountRate,
		ExpiredAt:      time.Unix(coupon.ExpiredAt, 0),
		UsedAt:         time.Unix(coupon.UsedAt, 0),
		OrderId:        coupon.OrderId,
		Status:         constants.CouponStatusUnused,
	}
	err := c.db.WithContext(ctx).Create(couponDB).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}

func (c *CouponDB) UpdateCouponStatus(ctx context.Context, couponId int64, status int) error {
	err := c.db.WithContext(ctx).Model(&model.Coupon{}).Where("coupon_id = ?", couponId).Update("status", status).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}

func (c *CouponDB) GetCouponList(ctx context.Context, status int) ([]*model.Coupon, error) {
	userId, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var coupons []*Coupon
	err = c.db.WithContext(ctx).Model(&Coupon{}).Where("user_id = ? AND status = ?", userId, status).Find(&coupons).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	var couponsRes []*model.Coupon
	for _, coupon := range coupons {
		couponsRes = append(couponsRes, &model.Coupon{
			BatchId:        coupon.BatchId,
			CouponId:       coupon.CouponId,
			UserId:         coupon.UserId,
			Type:           coupon.Type,
			Threshold:      coupon.Threshold,
			CreatedAt:      coupon.CreatedAt.Unix(),
			OrderId:        coupon.OrderId,
			Status:         coupon.Status,
			DiscountAmount: coupon.DiscountAmount,
			DiscountRate:   coupon.DiscountRate,
			ExpiredAt:      coupon.ExpiredAt.Unix(),
			UsedAt:         coupon.UsedAt.Unix(),
		})
	}
	return couponsRes, nil
}

func (c *CouponDB) LockCoupon(ctx context.Context, couponId int64) error {
	err := c.db.WithContext(ctx).Model(&model.Coupon{}).Where("coupon_id = ?", couponId).Update("status", constants.CouponStatusLocked).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}

func (c *CouponDB) ReleaseCoupon(ctx context.Context, couponId int64) error {
	err := c.db.WithContext(ctx).Model(&model.Coupon{}).Where("coupon_id = ?", couponId).Update("status", constants.CouponStatusUnused).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}

func (c *CouponDB) RedeemCoupon(ctx context.Context, couponId int64, orderId int64) error {
	err := c.db.WithContext(ctx).Where("coupon_id = ?", couponId).Update("status", constants.CouponStatusUsed).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	//使优惠卷已使用加一
	err = c.db.WithContext(ctx).Model(&model.CouponBatch{}).Where("batch_id = ?", couponId).Update("used_num", gorm.Expr("used_num + ?", 1)).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}

func (c *CouponDB) LetCouponExpire(ctx context.Context, couponId int64) error {
	err := c.db.WithContext(ctx).Model(&model.Coupon{}).Where("coupon_id = ?", couponId).Update("status", constants.CouponStatusExpired).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}
