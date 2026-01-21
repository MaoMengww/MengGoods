namespace go product

include "model.thrift"

struct CreateSpuReq {
    1: string name;
    2: string description;
    3: i64 categoryId;
    4: string mainSpuImageURL;
    5: string sliderSpuImageURLs;
    6: list<model.CreateSkuItem> sku;
}

struct CreateSpuResp {
    1: required model.BaseResp base
    2: required i64 spuId;
}


struct UpdateSpuReq {
    1: required i64 spuId;
    2: optional string name;
    3: optional string description;
    4: optional i64 categoryId;
    5: optional string mainSpuImageURL;
    6: optional string sliderSpuImageURLs;
}

struct UpdateSpuResp {
    1: required model.BaseResp base;
}

struct UpdateSkuReq {
    1: required i64 skuId;
    2: optional string name;
    3: optional string description;
    4: optional i64 price;
    5: optional string skuImageURL;
    6: optional string properties;
}

struct UpdateSkuResp {
    1: required model.BaseResp base;
}

struct DeleteSpuReq {
    1: required i64 spuId;
}

struct DeleteSpuResp {
    1: required model.BaseResp base;
}

struct DeleteSkuReq {
    1: required i64 skuId;
}

struct DeleteSkuResp {
    1: required model.BaseResp base;
}

struct GetSpuByIdReq {
    1: i64 spuId;
}

struct GetSpuByIdResp {
    1: required model.BaseResp base;
    2: required model.SpuInfo spuInfo;
}

//主要用于购物车
struct GetSkuReq {
    1: i64 skuId;
}

struct GetSkuResp {
    1: required model.BaseResp base;
    2: required model.SkuInfo skuInfo;
}



struct CreateCategoryReq {
    1: string name;
}

struct CreateCategoryResp {
    1: required model.BaseResp base;
    2: required i64 categoryId;
}

struct UpdateCategoryReq {
    1: required i64 categoryId;
    2: optional string name;
}

struct UpdateCategoryResp {
    1: required model.BaseResp base;
}

struct DeleteCategoryReq {
    1: required i64 categoryId;
}

struct DeleteCategoryResp {
    1: required model.BaseResp base;
}

//用于es查询spu列表
struct GetSpuReq {
    1: string name;
    2: i64 spuId
    3: i64 categoryId;
    4: i64 pageNum;
    5: i64 pageSize;
    6: i64 minPrice;
    7: i64 maxPrice;
}

struct GetSpuResp {
    1: required model.BaseResp base;
    2: required list<model.SpuInfo> spuList;
    3: required i64 total;
}



service ProductService {
    CreateSpuResp CreateSpu(1: CreateSpuReq req);
    UpdateSpuResp UpdateSpu(1: UpdateSpuReq req);
    UpdateSkuResp UpdateSku(1: UpdateSkuReq req);
    DeleteSpuResp DeleteSpu(1: DeleteSpuReq req);
    DeleteSkuResp DeleteSku(1: DeleteSkuReq req);
    GetSpuByIdResp GetSpuById(1: GetSpuByIdReq req);
    GetSkuResp GetSku(1: GetSkuReq req);
    CreateCategoryResp CreateCategory(1: CreateCategoryReq req);
    UpdateCategoryResp UpdateCategory(1: UpdateCategoryReq req);
    DeleteCategoryResp DeleteCategory(1: DeleteCategoryReq req);
    GetSpuResp GetSpuList(1: GetSpuReq req);
}

