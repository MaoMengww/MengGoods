package usecase

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/kitex_gen/product"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func (u *ProductUsecase) CreateSpu(ctx context.Context, spu *model.Spu) (int64, error) {
	//校验输入合法性
	if err := u.service.VerifySpu(spu); err != nil {
		return 0, err
	}
	//创建商品
	spuId, err := u.service.CreateSpu(ctx, spu)
	if err != nil {
		return 0, err
	}
	return spuId, nil
}

func (u *ProductUsecase) UpdateSpu(ctx context.Context, spu *model.Spu) error {
	//校验输入合法性
	if err := utils.Verify(utils.VerifySpuName(spu.Name), utils.VerifySpuDescription(spu.Description)); err != nil {
		return err
	}
	//验证权限
	isOwner, err := u.db.IsSpuOwer(ctx, spu.SpuId)
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	if !isOwner {
		return merror.NewMerror(merror.PermissionDenied, "Permission denied: not the SPU owner")
	}
	//更新商品
	return u.service.UpdateSpu(ctx, spu)
}

func (u *ProductUsecase) UpdateSku(ctx context.Context, sku *model.Sku) error {
	//校验输入合法性
	if err := utils.Verify(utils.VerifySkuName(sku.Name), utils.VerifySkuDescription(sku.Description)); err != nil {
		return err
	}
	//验证权限
	isOwner, err := u.db.IsSkuOwer(ctx, sku.SkuId)
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	if !isOwner {
		return merror.NewMerror(merror.PermissionDenied, "Permission denied: not the SKU owner")
	}
	//更新商品sku
	return u.service.UpdateSku(ctx, sku)
}

func (u *ProductUsecase) UpdateCategory(ctx context.Context, category *model.Category) error {
	//校验输入合法性
	if err := utils.Verify(utils.VerifyCategoryName(category.Name)); err != nil {
		return err
	}
	//验证权限
	isAdmin, err := u.rpc.IsAdmin(ctx)
	if err != nil {
		return merror.NewMerror(merror.InternalRpcErrorCode, err.Error())
	}
	if !isAdmin {
		return merror.NewMerror(merror.PermissionDenied, "Permission denied: not an admin")
	}
	//更新商品分类
	return u.db.UpdateCategory(ctx, category)
}

func (u *ProductUsecase) DeleteSpu(ctx context.Context, spuId int64) error {
	//验证权限
	isOwner, err := u.db.IsSpuOwer(ctx, spuId)
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	if !isOwner {
		return merror.NewMerror(merror.PermissionDenied, "Permission denied: not the SPU owner")
	}
	//删除商品
	return u.service.DeleteSpu(ctx, spuId)
}

func (u *ProductUsecase) DeleteSku(ctx context.Context, skuId int64) error {
	//验证权限
	isOwner, err := u.db.IsSkuOwer(ctx, skuId)
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	if !isOwner {
		return merror.NewMerror(merror.PermissionDenied, "Permission denied: not the SKU owner")
	}
	//删除商品
	return u.db.DeleteSku(ctx, skuId)
}

func (u *ProductUsecase) DeleteCategory(ctx context.Context, categoryId int64) error {
	//验证权限
	isAdmin, err := u.rpc.IsAdmin(ctx)
	if err != nil {
		return merror.NewMerror(merror.InternalRpcErrorCode, err.Error())
	}
	if !isAdmin {
		return merror.NewMerror(merror.PermissionDenied, "Permission denied: not an admin")
	}
	//删除商品
	return u.db.DeleteCategory(ctx, categoryId)
}

func (u *ProductUsecase) GetSpuById(ctx context.Context, spuId int64) (*model.Spu, error) {
	key := u.cache.GetSpuKey(ctx, spuId)
	var spu *model.Spu
	spuStr, err := u.cache.GetSpu(ctx, key)
	if err == nil {
		json.Unmarshal([]byte(spuStr), &spu)
		return spu, nil
	} else if err != redis.Nil {
		return nil, merror.NewMerror(merror.RedisNotFound, "spu not found")
	}
	spu, err = u.db.GetSpuById(ctx, spuId)
	if err != nil {
		return nil, err
	}
	skus, err := u.db.GetSkusBySpuId(ctx, spuId)
	if err != nil {
		return nil, err
	}
	spu.Skus = skus
	return spu, nil
}

func (s *ProductUsecase) CreateCategory(ctx context.Context, category *model.Category) (int64, error) {
	utils.Verify(utils.VerifyCategoryName(category.Name))
	return s.db.CreateCategory(ctx, category)
}

func (s *ProductUsecase) GetSkuById(ctx context.Context, skuId int64) (*model.Sku, error) {
	return s.db.GetSkuById(ctx, skuId)
}

func (s *ProductUsecase) GetSpusByIds(ctx context.Context, ids []int64) ([]*model.Spu, error) {
	return s.db.GetSpusByIds(ctx, ids)
}

func (s *ProductUsecase) GetSkusByIds(ctx context.Context, ids []int64) ([]*model.Sku, error) {
	return s.db.GetSkusByIds(ctx, ids)
}

func (s *ProductUsecase) GetSpuList(ctx context.Context, req *product.GetSpuReq) ([]*model.SpuEs, int64, error) {
	return s.es.SearchSpu(ctx, req)
}

func (s *ProductUsecase) UploadSkuImage(ctx context.Context, skuId int64, skuImageData []byte, fileName string) (string, error) {
	return s.service.UploadSkuImage(ctx, skuId, skuImageData, fileName)
}

func (s *ProductUsecase) UploadSpuImage(ctx context.Context, spuId int64, spuImageData []byte, fileName string) (string, error) {
	return s.service.UploadSpuImage(ctx, spuId, spuImageData, fileName)
}
