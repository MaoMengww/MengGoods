package resp

import "MengGoods/kitex_gen/model"

type GetCouponInfoResp struct {
    Coupon model.CouponInfo `json:"coupon"`
}

type CreateCouponBatchResp struct {
    BatchId int64 `json:"batchId"`
}

type GetCouponResp struct {
}

type GetCouponListResp struct {
	CouponList []*model.CouponInfo `json:"couponList"`
	Total int64 `json:"total"`
}



