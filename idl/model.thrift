namespace go model

struct BaseResp {
    1: i64 code;
    2: string message;
}

struct UserInfo {
    1: i64 id;
    2: string username;
    3: i64 role;
}

//userID放在context中隐式透传
struct AddressInfo {
    1: string province;
    2: string city;
    3: string detail;
    4: i64 addressID;
}

//商品
struct SpuInfo {
    1: i64 id;
    2: i64 creatorId;
    3: string name;
    4: string description;
    5: i64 categoryId;
    6: string mainImageURL;
    7: string sliderImageURLs;
    8: SpuStatus status;
    9: list<SkuInfo> sku;
    10: i64 createTime;
    11: i64 updateTime;
    12: i64 deleteTime;
}

struct SkuInfo {
    1: i64 id;
    2: i64 spuId;
    3: string name;
    4: string description;
    5: i64 price;
    6: string skuImageURL;
    7: string properties;
    8: i64 createTime;
    9: i64 updateTime;
    10: i64 deleteTime;
}

struct CreateSkuItem {
    1: string name;
    2: string description;
    3: i64 stock;
    4: i64 price;
    5: string skuImageURL;
    6: string properties;
}

enum SpuStatus {
    UNKNOWN = 0;
    ONLINE = 1;
    OFFLINE = 2;
    DELETED = 3;
}

struct CategoryInfo {
    1: i64 id;
    2: string name;
    3: i64 createTime
    4: i64 updateTime;
    5: i64 deleteTime;
}

struct StockItem {
    1: i64 skuId;
    2: i32 count;
}

struct CartItem {
    1: i64 skuId;
    2: i32 count;
    3: i64 updateTime
}

struct CouponBatchInfo {
    1: i64 id; //优惠券批次ID
    2: string batchName;
    3: string remark;
    4: i32 type;
    5: i64 threshold; //优惠券使用门槛
    6: i64 amount;
    7: i64 rate; //优惠券折扣率
    8: i64 total;
    9: i64 usedNum; //优惠券已发放数量
    10: i64 startTime;
    11: i64 endTime;
    12: i64 duration; //优惠券持续时间(天)
    13: i64 createTime;
    14: i64 updateTime;
}

struct CouponInfo {
    1: i64 id; //优惠券ID
    2: i64 batchId; //优惠券批次ID
    3: i64 userId;
    4: i64 orderId;
    5: i64 status;
    6: i32 type;
    7: i64 threshold; //优惠券使用门槛
    8: i64 amount;
    9: i64 rate; //优惠券折扣率
    10: i64 createTime;
    11: i64 expiredAt; //过期时间
    12: i64 usedAt; 
}



