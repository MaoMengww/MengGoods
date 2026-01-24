package pack

import (
	mModel "MengGoods/app/coupon/domain/model"
	"MengGoods/kitex_gen/model"
)

func ToDomainCoupon(coupon *model.CouponInfo) *mModel.Coupon {
    return &mModel.Coupon{
        CouponId: coupon.Id,
		BatchId: coupon.BatchId,
		OrderId: coupon.OrderId,
		UserId: coupon.UserId,
		Status: int(coupon.Status),	
		CreatedAt: coupon.CreateTime,
		ExpiredAt: coupon.ExpiredAt,
		UsedAt: coupon.UsedAt,
    }
}

func ToRpcCoupon(coupon *mModel.Coupon) *model.CouponInfo {
    return &model.CouponInfo{
        Id: coupon.CouponId,
		BatchId: coupon.BatchId,
		OrderId: coupon.OrderId,
		UserId: coupon.UserId,
		Status: int64(coupon.Status),	
		CreateTime: coupon.CreatedAt,
		ExpiredAt: coupon.ExpiredAt,
		UsedAt: coupon.UsedAt,
    }
}

func ToDomainCouponBatch(batch *model.CouponBatchInfo) *mModel.CouponBatch {
    return &mModel.CouponBatch{
        BatchId: batch.Id,
		BatchName: batch.BatchName,
		Remark: batch.Remark,
		Type: int(batch.Type),
		Threshold: batch.Threshold,
		DiscountAmount: batch.Amount,
		DiscountPercent: int(batch.Rate),
		TotalNum: batch.Total,
		StartTime: batch.StartTime,
		EndTime: batch.EndTime,
		Duration: int(batch.Duration),
		CreatedAt: batch.CreateTime,
		UpdatedAt: batch.UpdateTime,
    }
}

func ToRpcCouponBatch(batch *mModel.CouponBatch) *model.CouponBatchInfo {
    return &model.CouponBatchInfo{
        Id: batch.BatchId,
		BatchName: batch.BatchName,
		Remark: batch.Remark,
		Type: int32(batch.Type),
		Threshold: batch.Threshold,
		Amount: batch.DiscountAmount,
		Rate: int64(batch.DiscountPercent),
		Total: batch.TotalNum,
		StartTime: batch.StartTime,
		EndTime: batch.EndTime,
		Duration: int64(batch.Duration),
		CreateTime: batch.CreatedAt,
		UpdateTime: batch.UpdatedAt,
    }
}

func ToDomainCoupons(coupons []*model.CouponInfo) []*mModel.Coupon {
    var domainCoupons []*mModel.Coupon
    for _, coupon := range coupons {
        domainCoupons = append(domainCoupons, ToDomainCoupon(coupon))
    }
    return domainCoupons
}

func ToRpcCoupons(coupons []*mModel.Coupon) []*model.CouponInfo {
    var rpcCoupons []*model.CouponInfo
    for _, coupon := range coupons {
        rpcCoupons = append(rpcCoupons, ToRpcCoupon(coupon))
    }
    return rpcCoupons
}