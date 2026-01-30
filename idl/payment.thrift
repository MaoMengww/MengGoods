namespace go payment

include "model.thrift"

struct GetPaymentTokenReq {
    1: i64 orderId;
    2: i32 paymentMethod;
}

struct GetPaymentTokenResp {
    1: model.BaseResp base;
    2: string paymentToken;
    3: i64 expiredAt;
}

struct PaymentReq {
    1: i64 orderId;
    2: string paymentToken;
}

struct PaymentResp {
    1: model.BaseResp base;
}

struct PaymentRefundReq {
    1: i64 orderItemId;
    2: string refundReason;
}

struct PaymentRefundResp {
    1: model.BaseResp base;
}

struct ReviewRefundReq {
    1: i64 orderItemId;
    2: bool approve;
}

struct ReviewRefundResp {
    1: model.BaseResp base;
}

service PaymentService {
    GetPaymentTokenResp GetPaymentToken(1: GetPaymentTokenReq req);
    PaymentResp Payment(1: PaymentReq req);
    PaymentRefundResp PaymentRefund(1: PaymentRefundReq req);
    ReviewRefundResp ReviewRefund(1: ReviewRefundReq req);
}
