namespace go cart

include "model.thrift"

struct AddCartItemReq {
    1: i64 userId;
    2: model.CartItem cartItem;
}

struct AddCartItemResp {
    1: model.BaseResp base;
}

struct GetCartItemReq {
    1: i64 userId;
}

struct GetCartItemResp {
    1: model.BaseResp base;
    2: list<model.CartItem> cartItems;
}

struct DeleteCartItemReq {
    1: i64 userId;
}

struct DeleteCartItemResp {
    1: model.BaseResp base;
}

struct UpdateCartItemReq {
    1: i64 userId;
    2: model.CartItem cartItem;
}

struct UpdateCartItemResp {
    1: model.BaseResp base;
}

service CartService {
    AddCartItemResp AddCartItem(1: AddCartItemReq req);
    GetCartItemResp GetCartItem(1: GetCartItemReq req);
    DeleteCartItemResp DeleteCartItem(1: DeleteCartItemReq req);
    UpdateCartItemResp UpdateCartItem(1: UpdateCartItemReq req);
}
