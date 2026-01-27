namespace go payment

include "model.thrift"

struct GetPaymentTokenReq {
    1: i64 orderId;
    2: string paymentMethod;
}

struct GetPaymentTokenResp {
    1: model.BaseResp base;
    2: string paymentToken;
}

struct PaymentReq {
    1: i64 orderId;
    2: string paymentToken;
}

struct PaymentResp {
    1: model.BaseResp base;
}

struct GetRefundTokenReq {
    1: i64 orderId;
    2: string paymentMethod;
}

struct GetRefundTokenResp {
    1: model.BaseResp base;
    2: string refundToken;
}

struct RefundReq {
    1: i64 orderId;
    2: string refundReason;
    3: string refundToken;
}

struct RefundResp {
    1: model.BaseResp base;
}

service PaymentService {
    GetPaymentTokenResp GetPaymentToken(1: GetPaymentTokenReq req);
    PaymentResp Payment(1: PaymentReq req);
    GetRefundTokenResp GetRefundToken(1: GetRefundTokenReq req);
    RefundResp Refund(1: RefundReq req);
}
