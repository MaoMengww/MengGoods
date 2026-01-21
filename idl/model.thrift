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




