namespace go stock

include "model.thrift"

struct GetStockReq {
    1: i64 skuId;
}

struct GetStockResp {
    1: model.BaseResp base;
    2: model.StockItem stock;
}

struct GetStocksReq {
    1: list<i64> skuIds;
}

struct GetStocksResp {
    1: model.BaseResp base;
    2: list<model.StockItem> stocks;
}

struct CreateStockReq {
    1: i64 skuId;
    2: i32 count;
}

struct CreateStockResp {
    1: model.BaseResp base;
}

struct AddStockReq {
    1: i64 skuId;
    2: i32 count;
}

struct AddStockResp {
    1: model.BaseResp base;
}


# 锁定库存
struct LockStockReq {
    1: i64 orderId;
    2: list<model.StockItem> stockItems;
}

struct LockStockResp {
    1:  model.BaseResp base;
}

# 解锁库存(订单过期, 支付失败)
struct UnlockStockReq {
    1: i64 orderId;
    2: list<model.StockItem> stockItems;
}

struct UnlockStockResp {
    1: model.BaseResp base;
}

# 扣减库存(订单支付成功)
struct DeductStockReq {
    1: i64 orderId;
    2: list<model.StockItem> stockItems;
}

struct DeductStockResp {
    1: model.BaseResp base;
}

service StockService {
    CreateStockResp CreateStock(1: CreateStockReq req);
    AddStockResp AddStock(1: AddStockReq req);
    GetStockResp GetStock(1: GetStockReq req);
    GetStocksResp GetStocks(1: GetStocksReq req);
    LockStockResp LockStock(1: LockStockReq req);
    UnlockStockResp UnlockStock(1: UnlockStockReq req);
    DeductStockResp DeductStock(1: DeductStockReq req);
}







