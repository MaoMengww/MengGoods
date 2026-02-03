package repository

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/kitex_gen/product"
	"context"
)

// 负责写和精确查找spu
type ProductDB interface {
	CreateSpu(ctx context.Context, spu *model.Spu) (int64, error)
	CreateCategory(ctx context.Context, category *model.Category) (int64, error)
	DeleteSpu(ctx context.Context, spuId int64) error
	DeleteSku(ctx context.Context, skuId int64) error
	DeleteCategory(ctx context.Context, categoryId int64) error
	UpdateSpu(ctx context.Context, spu *model.Spu) error
	UpdateSku(ctx context.Context, sku *model.Sku) error
	UpdateCategory(ctx context.Context, category *model.Category) error

	GetSpuById(ctx context.Context, spuId int64) (*model.Spu, error)
	GetSkuById(ctx context.Context, skuId int64) (*model.Sku, error)
	GetSpusByIds(ctx context.Context, spuIds []int64) ([]*model.Spu, error)
	GetSkusByIds(ctx context.Context, skuIds []int64) ([]*model.Sku, error) //用于购物车相关
	GetSkusBySpuId(ctx context.Context, spuId int64) ([]*model.Sku, error)

	IsSpuOwer(ctx context.Context, spuId int64) (bool, error)
	IsSkuOwer(ctx context.Context, skuId int64) (bool, error)

	IsSpuExist(ctx context.Context, spuId int64) (bool, error)
	IsSkuExist(ctx context.Context, skuId int64) (bool, error)
	IsCategoryExist(ctx context.Context, categoryId int64) (bool, error)

	UploadSpuImage(ctx context.Context, spuId int64, imageUrl string) error
	UploadSkuImage(ctx context.Context, skuId int64, imageUrl string) error

	GetAllSpuIdAndSkuId(ctx context.Context)(spuIds []int64, skuIds []int64, err error)
}

type ProductCache interface {
	GetSpuKey(ctx context.Context, spuId int64) string
	GetSkuKey(ctx context.Context, skuId int64) string
	SetSpu(ctx context.Context, spu *model.SpuEs) error
	SetSku(ctx context.Context, sku *model.SkuEs) error
	GetSpu(ctx context.Context, spuId int64) (string, error)
	GetSku(ctx context.Context, skuId int64) (string, error)
	DeleteSpu(ctx context.Context, key string) error
	DeleteSku(ctx context.Context, key string) error
	LoadBloomFilter(ctx context.Context, spuIds []int64, skuIds []int64) error
}

type ProductMq interface {
	SendCreateSpuInfo(ctx context.Context, spu *model.SpuEs) error
	SendUpdateSpuInfo(ctx context.Context, spu *model.SpuEs) error
	SendDeleteSpuInfo(ctx context.Context, id int64) error
	ConsumeCreateSpuInfo(ctx context.Context, fn func(ctx context.Context, spu *model.SpuEs) error) error
	ConsumeUpdateSpuInfo(ctx context.Context, fn func(ctx context.Context, spu *model.SpuEs) error) error
	ConsumeDeleteSpuInfo(ctx context.Context, fn func(ctx context.Context, spuId int64) error) error
}


type ProductEs interface {
	AddSpuItem(ctx context.Context, spu *model.SpuEs) error
	UptateSpuItem(ctx context.Context, spu *model.SpuEs) error
	DeleteSpuItem(ctx context.Context, spuId int64) error
	SearchSpu(ctx context.Context, req *product.GetSpuReq) ([]*model.SpuEs, int64, error)
}

type ProductRpc interface {
	IsAdmin(ctx context.Context) (bool, error)
}

type ProductCos interface {
	UploadSpuImage(ctx context.Context, spuImageData []byte, fileName string) (string, error)
	UploadSkuImage(ctx context.Context, skuImageData []byte, fileName string) (string, error)
}
