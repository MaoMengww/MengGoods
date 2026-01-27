package prpc

import (
	"MengGoods/kitex_gen/coupon"
	"MengGoods/kitex_gen/model"
	"MengGoods/pkg/merror"
	"context"
)

func (c *OrderRpc) LockCoupon(ctx context.Context, couponId int64) error {
	resp, err := c.CouponClient.LockCoupon(ctx, &coupon.LockCouponReq{
		CouponId: couponId,
	})
	if err != nil {
		return err
	}
	if resp.Base.Code != merror.SuccessCode {
		return merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return nil
}

func (c *OrderRpc) ReleaseCoupon(ctx context.Context, couponId int64) error {
	resp, err := c.CouponClient.ReleaseCoupon(ctx, &coupon.ReleaseCouponReq{
		CouponId: couponId,
	})
	if err != nil {
		return err
	}
	if resp.Base.Code != merror.SuccessCode {
		return merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return nil
}

func (c *OrderRpc) GetCouponInfo(ctx context.Context, couponId int64) (*model.CouponInfo, error) {
	resp, err := c.CouponClient.GetCouponInfo(ctx, &coupon.GetCouponInfoReq{
		CouponId: couponId,
	})
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return resp.Coupon, nil
}

func (c *OrderRpc) RedeemCoupon(ctx context.Context, couponId int64) error {
	resp, err := c.CouponClient.RedeemCoupon(ctx, &coupon.RedeemCouponReq{
		CouponId: couponId,
	})
	if err != nil {
		return err
	}
	if resp.Base.Code != merror.SuccessCode {
		return merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return nil
}

func (c *OrderRpc) LetCouponExpire(ctx context.Context, couponId int64) error {
	resp, err := c.CouponClient.LetCouponExpire(ctx, &coupon.LetCouponExpireReq{
		CouponId: couponId,
	})
	if err != nil {
		return err
	}
	if resp.Base.Code != merror.SuccessCode {
		return merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return nil
}
