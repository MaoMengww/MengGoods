package resp

import (
	mModel "MengGoods/app/product/domain/model"
	"MengGoods/kitex_gen/model"
)


type CreateSpuResp struct {
    SpuId int64 `json:"spu_id"`
}

type UpdateSpuResp struct {
}

type UpdateSkuResp struct {
}

type DeleteSpuResp struct {
}

type DeleteSkuResp struct {
}

type GetSpuByIdResp struct {
    Spu model.SpuInfo `json:"spu"`
}

type GetSkuResp struct {
    Sku model.SkuInfo `json:"sku"`
}

type CreateCategoryResp struct {
    CategoryId int64 `json:"category_id"`
}

type UpdateCategoryResp struct {
}

type DeleteCategoryResp struct {
}

type GetSpuResp struct {
    Spus []*mModel.SpuEs `json:"spus"`
    Total int64 `json:"total"`
}

