namespace go coupon

include "model.thrift"

struct CreateCouponBatchReq {
    1: string batchName;
    2: string remark;
    3: i32 type;
    4: i64 threshold;
    5: i64 amount;
    6: i64 rate;
    7: i64 total;
    8: i64 startTime;
    9: i64 endTime;
    10: i64 duration; //优惠券持续时间(天)
}

struct CreateCouponBatchResp {
    1: model.BaseResp base;
    2: i64 batchId;
}

// 领取优惠券
struct GetCouponReq {
    1: i64 batchId;
}

struct GetCouponResp {
    1: model.BaseResp base;
}

struct GetCouponListReq {
    1: i64 status;
}

struct GetCouponListResp {
    1: model.BaseResp base;
    2: list<model.CouponInfo> couponList;
    3: i64 total;
}

struct LockCouponReq {
    1: i64 couponId;
    2: i64 orderId;
}

struct LockCouponResp {
    1: model.BaseResp base;
}

struct ReleaseCouponReq {
    1: i64 couponId;
    2: i64 orderId;
}

struct ReleaseCouponResp {
    1: model.BaseResp base;
}

struct RedeemCouponReq {
    1: i64 couponId;
    2: i64 orderId;
}

struct RedeemCouponResp {
    1: model.BaseResp base;
}

service CouponService {
    CreateCouponBatchResp CreateCouponBatch(1: CreateCouponBatchReq req);
    GetCouponResp GetCoupon(1: GetCouponReq req);
    GetCouponListResp GetCouponList(1: GetCouponListReq req);
    LockCouponResp LockCoupon(1: LockCouponReq req);
    ReleaseCouponResp ReleaseCoupon(1: ReleaseCouponReq req);
    RedeemCouponResp RedeemCoupon(1: RedeemCouponReq req);
}


