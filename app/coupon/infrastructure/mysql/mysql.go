package mysql

import (
	"MengGoods/app/coupon/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"context"

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
	var coupon model.Coupon
	err := c.db.WithContext(ctx).Where("coupon_id = ?", couponId).First(&coupon).Error
	if err != nil {
		return nil, err
	}
	return &coupon, nil
}

func (c *CouponDB) GetCouponBatchByID(ctx context.Context, batchId int64) (*model.CouponBatch, error) {
	var couponBatch model.CouponBatch
	err := c.db.WithContext(ctx).Where("batch_id = ?", batchId).First(&couponBatch).Error
	if err != nil {
		return nil, err
	}
	return &couponBatch, nil
}

func (c *CouponDB) CreateCouponBatch(ctx context.Context, batch *model.CouponBatch) (int64, error) {
	CouponBatch := &model.CouponBatch{
		BatchName:       batch.BatchName,
		Remark:          batch.Remark,
		Type:            batch.Type,
		Threshold:       batch.Threshold,
		DiscountAmount:  batch.DiscountAmount,
		DiscountPercent: batch.DiscountPercent,
		TotalNum:        batch.TotalNum,
		StartTime:       batch.StartTime,
		EndTime:         batch.EndTime,
		Duration:        batch.Duration,
	}
	err := c.db.WithContext(ctx).Create(CouponBatch).Error
	if err != nil {
		return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return CouponBatch.BatchId, nil
}

func (c *CouponDB) IsUnusedCoupon(ctx context.Context, couponId int64) (bool, error) {
	var coupon model.Coupon
	err := c.db.WithContext(ctx).Where("coupon_id = ?", couponId).First(&coupon).Error
	if err != nil {
		return false, err
	}
	return coupon.Status == constants.CouponStatusUnused, nil
}

func (c *CouponDB) CreateCoupon(ctx context.Context, coupon *model.Coupon) error {
	err := c.db.WithContext(ctx).Create(coupon).Error
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
	var coupons []*model.Coupon
	err = c.db.WithContext(ctx).Model(&model.Coupon{}).Where("user_id = ? AND status = ?", userId, status).Find(&coupons).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return coupons, nil
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
