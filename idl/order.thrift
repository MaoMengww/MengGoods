namespace go order

include "model.thrift"

struct CreateOrderReq {
    1: list<model.OrderItem> items;
    2: i64 couponId;
    3: i64 addressId;
}

struct CreateOrderResp {
    1: model.BaseResp base;
    2: i64 orderId;
}


struct ViewOrderByIdReq {
    1: i64 orderId;
}

struct ViewOrderByIdResp {
    1: model.BaseResp base;
    2: model.OrderWithItems order;
}

struct ViewOrderListReq {
    1: i64 status;
    2: i64 pageNum;
    3: i64 pageSize;
}

struct ViewOrderListResp {
    1: model.BaseResp base;
    2: list<model.OrderInfo> orderList;
    3: i64 total;
}

struct ConfirmReceiptOrderReq {
    1: i64 orderId;
}

struct ConfirmReceiptOrderResp {
    1: model.BaseResp base;
}

struct CancelOrderReq {
    1: i64 orderId;
    2: string cancelReason;
}

struct CancelOrderResp {
    1: model.BaseResp base;
}

struct MarkOrderPaidReq {
    1: i64 orderId;
}

struct MarkOrderPaidResp {
    1: model.BaseResp base;
}

struct GetPayAmountReq {
    1: i64 orderId;
}

struct GetPayAmountResp {
    1: model.BaseResp base;
    2: i64 amount;
}

struct IsOrderExistReq {
    1: i64 orderId;
}

struct IsOrderExistResp {
    1: model.BaseResp base;
    2: bool exist;
    3: i64 expiredAt;
}

struct GetOrderItemReq {
    1: i64 orderItemId;
}

struct GetOrderItemResp {
    1: model.BaseResp base;
    2: model.OrderItem orderItem;
}

struct GetOrderInfoReq {
    1: i64 orderId;
}

struct GetOrderInfoResp {
    1: model.BaseResp base;
    2: model.OrderInfo orderInfo;
}



service OrderService {
    CreateOrderResp CreateOrder(1: CreateOrderReq req);
    ViewOrderByIdResp ViewOrderById(1: ViewOrderByIdReq req);
    ViewOrderListResp ViewOrderList(1: ViewOrderListReq req);
    CancelOrderResp CancelOrder(1: CancelOrderReq req);
    ConfirmReceiptOrderResp ConfirmReceiptOrder(1: ConfirmReceiptOrderReq req);
    MarkOrderPaidResp MarkOrderPaid(1: MarkOrderPaidReq req);
    GetPayAmountResp GetPayAmount(1: GetPayAmountReq req);
    GetOrderInfoResp GetOrderInfo(1: GetOrderInfoReq req);
    GetOrderItemResp GetOrderItem(1: GetOrderItemReq req);
    IsOrderExistResp IsOrderExist(1: IsOrderExistReq req);
}

