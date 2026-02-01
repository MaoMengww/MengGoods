package pack

import (
	mModel "MengGoods/app/product/domain/model"
	"MengGoods/kitex_gen/model"
)

func BuildSkus(skus []*model.CreateSkuItem) []*mModel.Sku {
	var skusResp []*mModel.Sku
	for _, sku := range skus {
		skusResp = append(skusResp, &mModel.Sku{
			Name:        sku.Name,
			Description: sku.Description,
			Properties:  sku.Properties,
			Price:       sku.Price,
		})
	}
	return skusResp
}

func BuildSpuInfoList(spus []*mModel.SpuEs) []*model.SpuInfo {
	var spusResp []*model.SpuInfo
	for _, spu := range spus {
		spusResp = append(spusResp, BuildSpuEsInfo(spu))
	}
	return spusResp
}

func BuildSpuInfo(spu *mModel.Spu) *model.SpuInfo {
	deleteTime := int64(0)
	if spu.DeleteTime != nil {
		deleteTime = spu.DeleteTime.Unix()
	}
	return &model.SpuInfo{
		Id:              spu.SpuId,
		CreatorId:       spu.UserId,
		Name:            spu.Name,
		Description:     spu.Description,
		CategoryId:      spu.CategoryId,
		Sku:             BuildSkusResp(spu.Skus),
		MainImageURL:    spu.MainImageURL,
		SliderImageURLs: spu.SliderImageURLs,
		Status:          model.SpuStatus(spu.Status),
		CreateTime:      spu.CreateTime.Unix(),
		UpdateTime:      spu.UpdateTime.Unix(),
		DeleteTime:      deleteTime,
	}
}

func BuildSpuEsInfo(spu *mModel.SpuEs) *model.SpuInfo {
	return &model.SpuInfo{
		Id:           spu.Id,
		CreatorId:    spu.UserId,
		Name:         spu.Name,
		Description:  spu.Description,
		CategoryId:   spu.CategoryId,
		MainImageURL: spu.MainImageURL,
	}

}

func BuildSkuInfo(sku *mModel.Sku) *model.SkuInfo {
	return &model.SkuInfo{
		Id:          sku.SkuId,
		Name:        sku.Name,
		Description: sku.Description,
		Properties:  sku.Properties,
		SkuImageURL: sku.ImageURL,
		Price:       sku.Price,
		CreateTime:  sku.CreateTime,
		UpdateTime:  sku.UpdateTime,
		DeleteTime:  sku.DeleteTime,
	}
}

func BuildSkusResp(skus []*mModel.Sku) []*model.SkuInfo {
	var skusResp []*model.SkuInfo
	for _, sku := range skus {
		skusResp = append(skusResp, BuildSkuInfo(sku))
	}
	return skusResp
}
