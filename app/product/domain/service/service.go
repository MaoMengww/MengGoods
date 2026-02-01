package service

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"fmt"
)

func (s *ProductService) CreateSpu(ctx context.Context, spu *model.Spu) (int64, error) {
	userId, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return 0, err
	}
	// 创建商品
	spuId, err := s.db.CreateSpu(ctx, spu)
	if err != nil {
		return 0, err
	}
	spuEs := &model.SpuEs{
		Id:           spuId,
		UserId:       userId,
		Name:         spu.Name,
		Description:  spu.Description,
		CategoryId:   spu.CategoryId,
		MainImageURL: spu.MainImageURL,
		Price:        spu.Price,
	}
	err = s.cache.SetSpu(ctx, spuEs)
	if err != nil {
		return 0, err
	}
	// 发送创建商品消息
	go func() {
		if err := s.mq.SendCreateSpuInfo(ctx, spuEs); err != nil {
			logger.Error(ctx, "send create spu info error: %v", err)
		}
	}()
	return spuId, nil
}

func (s *ProductService) UpdateSpu(ctx context.Context, spu *model.Spu) error {
	// 更新商品
	if err := s.db.UpdateSpu(ctx, spu); err != nil {
		return err
	}
	go func() {
		// 创建商品索引
		if err := s.es.AddSpuItem(ctx, &model.SpuEs{
			Id:           spu.SpuId,
			UserId:       spu.UserId,
			Name:         spu.Name,
			Description:  spu.Description,
			CategoryId:   spu.CategoryId,
			MainImageURL: spu.MainImageURL,
			Price:        spu.Price,
		}); err != nil {
			logger.Error(ctx, "add spu item error: %v", err)
		}
	}()
	return nil
}

func (s *ProductService) UpdateSku(ctx context.Context, sku *model.Sku) error {
	// 更新商品sku
	if err := s.db.UpdateSku(ctx, sku); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) DeleteSpu(ctx context.Context, spuId int64) error {
	// 删除商品
	if err := s.db.DeleteSpu(ctx, spuId); err != nil {
		return err
	}

	go func() {
		if err := s.mq.SendDeleteSpuInfo(ctx, spuId); err != nil {
			logger.Error(ctx, "send delete spu info error: %v", err)
		}
	}()
	return nil
}

func (s *ProductService) VerifySpu(spu *model.Spu) error {
	if err := utils.Verify(utils.VerifySpuName(spu.Name), utils.VerifySpuDescription(spu.Description)); err != nil {
		return err
	}
	for _, sku := range spu.Skus {
		if err := utils.Verify(utils.VerifySkuName(sku.Name), utils.VerifySkuPrice(sku.Price), utils.VerifySkuDescription(sku.Description)); err != nil {
			return err
		}
	}
	return nil
}

func (s *ProductService) CreateCategory(ctx context.Context, category *model.Category) (int64, error) {
	// 检查是否为管理员
	res, err := s.rpc.IsAdmin(ctx)
	if err != nil {
		return 0, merror.NewMerror(merror.InternalRpcErrorCode, fmt.Sprintf("rpc is admin error: %v", err))
	}
	if !res {
		return 0, merror.NewMerror(merror.PermissionDenied, "only admin can create category")
	}
	// 创建商品分类
	categoryId, err := s.db.CreateCategory(ctx, category)
	if err != nil {
		return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db create category error: %v", err))
	}
	return categoryId, nil
}

func (s *ProductService) DeleteCategory(ctx context.Context, categoryId int64) error {
	// 检查是否为管理员
	res, err := s.rpc.IsAdmin(ctx)
	if err != nil {
		return merror.NewMerror(merror.InternalRpcErrorCode, fmt.Sprintf("rpc is admin error: %v", err))
	}
	if !res {
		return merror.NewMerror(merror.PermissionDenied, "only admin can delete category")
	}
	// 删除商品分类
	if err := s.db.DeleteCategory(ctx, categoryId); err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db delete category error: %v", err))
	}
	return nil
}

func (s *ProductService) UpdateCategory(ctx context.Context, category *model.Category) error {
	// 检查是否为管理员
	res, err := s.rpc.IsAdmin(ctx)
	if err != nil {
		return merror.NewMerror(merror.InternalRpcErrorCode, fmt.Sprintf("rpc is admin error: %v", err))
	}
	if !res {
		return merror.NewMerror(merror.PermissionDenied, "only admin can update category")
	}
	// 更新商品分类
	if err := s.db.UpdateCategory(ctx, category); err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db update category error: %v", err))
	}
	return nil
}

func (s *ProductService) UploadSpuImage(ctx context.Context, spuId int64, spuImageData []byte, fileName string) (string, error) {
	ok, err := s.db.IsSpuOwer(ctx, spuId)
	if !ok {
		return "", merror.NewMerror(merror.PermissionDenied, "permission denied")
	}
	if err != nil {
		return "", merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db check spu owner error: %v", err))
	}
	imageUrl, err := s.cos.UploadSpuImage(ctx, spuImageData, fileName)
	if err != nil {
		return "", merror.NewMerror(merror.InternalCosErrorCode, fmt.Sprintf("cos upload spu image error: %v", err))
	}
	if err := s.db.UploadSpuImage(ctx, spuId, imageUrl); err != nil {
		return "", merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db upload spu image error: %v", err))
	}
	spu, err := s.db.GetSpuById(ctx, spuId)
	if err != nil {
		return "", merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db get spu error: %v", err))
	}
	spuEs := &model.SpuEs{
		Id:           spu.SpuId,
		UserId:       spu.UserId,
		Name:         spu.Name,
		Description:  spu.Description,
		CategoryId:   spu.CategoryId,
		MainImageURL: spu.MainImageURL,
		Price:        spu.Price,
	}
	if err := s.cache.SetSpu(ctx, spuEs); err != nil {
		return "", merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("cache set spu key error: %v", err))
	}
	if err := s.es.UptateSpuItem(ctx, spuEs); err != nil {
		return "", merror.NewMerror(merror.InternalESErrorCode, fmt.Sprintf("es update spu item error: %v", err))
	}
	return imageUrl, nil
}

func (s *ProductService) UploadSkuImage(ctx context.Context, skuId int64, skuImageData []byte, fileName string) (string, error) {
	ok, err := s.db.IsSkuOwer(ctx, skuId)
	if !ok {
		return "", merror.NewMerror(merror.PermissionDenied, "permission denied")
	}
	if err != nil {
		return "", merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db check sku owner error: %v", err))
	}
	imageUrl, err := s.cos.UploadSkuImage(ctx, skuImageData, fileName)
	if err != nil {
		return "", merror.NewMerror(merror.InternalCosErrorCode, fmt.Sprintf("cos upload sku image error: %v", err))
	}
	if err := s.db.UploadSkuImage(ctx, skuId, imageUrl); err != nil {
		return "", merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db upload sku image error: %v", err))
	}
	sku, err := s.db.GetSkuById(ctx, skuId)
	if err != nil {
		return "", merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db get sku error: %v", err))
	}
	skuEs := &model.SkuEs{
		Id:          sku.SkuId,
		SpuId:       sku.SpuId,
		Name:        sku.Name,
		Description: sku.Description,
		Properties:  sku.Properties,
		Price:       sku.Price,
		ImageURL:    sku.ImageURL,
	}
	if err := s.cache.SetSku(ctx, skuEs); err != nil {
		return "", merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("cache set sku key error: %v", err))
	}
	return imageUrl, nil
}

func (s *ProductService) ConsumeCreateSpuInfo(ctx context.Context) error {
	return s.mq.ConsumeCreateSpuInfo(ctx, s.es.AddSpuItem)
}

func (s *ProductService) ConsumeUpdateSpuInfo(ctx context.Context) error {
	return s.mq.ConsumeUpdateSpuInfo(ctx, s.es.UptateSpuItem)
}

func (s *ProductService) ConsumeDeleteSpuInfo(ctx context.Context) error {
	return s.mq.ConsumeDeleteSpuInfo(ctx, s.es.DeleteSpuItem)
}
