package api

import (
	"MengGoods/app/coupon/controller/api/pack"
	"MengGoods/app/coupon/domain/model"
	"MengGoods/app/coupon/usecase"
	coupon "MengGoods/kitex_gen/coupon"
	"MengGoods/pkg/base"
	"context"
)

// CouponServiceImpl implements the last service interface defined in the IDL.
type CouponServiceImpl struct {
	Usecase *usecase.CouponUsecase
}

func NewCouponServiceImpl(usecase *usecase.CouponUsecase) *CouponServiceImpl {
	return &CouponServiceImpl{
		Usecase: usecase,
	}
}

// CreateCouponBatch implements the CouponServiceImpl interface.
func (s *CouponServiceImpl) CreateCouponBatch(ctx context.Context, req *coupon.CreateCouponBatchReq) (resp *coupon.CreateCouponBatchResp, err error) {
	resp = new(coupon.CreateCouponBatchResp)
	couponBatch := &model.CouponBatch{
		BatchName:       req.BatchName,
		Remark:          req.Remark,
		Type:            int(req.Type),
		Threshold:       req.Threshold,
		DiscountAmount:  req.Amount,
		DiscountPercent: int(req.Rate),
		TotalNum:        req.Total,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		Duration:        int(req.Duration),
	}
	id, err := s.Usecase.CreateCouponBatch(ctx, couponBatch)
	if err != nil {
		resp.Base = base.BuildBaseResp(nil)
		return resp, nil
	}
	resp.BatchId = id
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil

}

// GetCoupon implements the CouponServiceImpl interface.
func (s *CouponServiceImpl) GetCoupon(ctx context.Context, req *coupon.GetCouponReq) (resp *coupon.GetCouponResp, err error) {
	resp = new(coupon.GetCouponResp)
	err = s.Usecase.GetCoupon(ctx, req.BatchId)
	if err != nil {
		resp.Base = base.BuildBaseResp(nil)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// GetCouponList implements the CouponServiceImpl interface.
func (s *CouponServiceImpl) GetCouponList(ctx context.Context, req *coupon.GetCouponListReq) (resp *coupon.GetCouponListResp, err error) {
	resp = new(coupon.GetCouponListResp)
	coupons, err := s.Usecase.GetCouponList(ctx, int(req.Status))
	if err != nil {
		resp.Base = base.BuildBaseResp(nil)
		return resp, nil
	}
	resp.CouponList = pack.ToRpcCoupons(coupons)
	resp.Total = int64(len(coupons))
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// LockCoupon implements the CouponServiceImpl interface.
func (s *CouponServiceImpl) LockCoupon(ctx context.Context, req *coupon.LockCouponReq) (resp *coupon.LockCouponResp, err error) {
	resp = new(coupon.LockCouponResp)
	err = s.Usecase.LockCoupon(ctx, req.CouponId)
	if err != nil {
		resp.Base = base.BuildBaseResp(nil)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// ReleaseCoupon implements the CouponServiceImpl interface.
func (s *CouponServiceImpl) ReleaseCoupon(ctx context.Context, req *coupon.ReleaseCouponReq) (resp *coupon.ReleaseCouponResp, err error) {
	resp = new(coupon.ReleaseCouponResp)
	err = s.Usecase.ReleaseCoupon(ctx, req.CouponId)
	if err != nil {
		resp.Base = base.BuildBaseResp(nil)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// RedeemCoupon implements the CouponServiceImpl interface.
func (s *CouponServiceImpl) RedeemCoupon(ctx context.Context, req *coupon.RedeemCouponReq) (resp *coupon.RedeemCouponResp, err error) {
	resp = new(coupon.RedeemCouponResp)
	err = s.Usecase.RedeemCoupon(ctx, req.CouponId, req.OrderId)
	if err != nil {
		resp.Base = base.BuildBaseResp(nil)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}
