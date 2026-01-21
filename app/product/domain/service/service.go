package service

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"fmt"
)


func (s *ProductUsecase) CreateSpu(ctx context.Context, spu *model.Spu) (int64, error) {
	// 创建商品
	spuId, err := s.db.CreateSpu(ctx, spu)
	if err != nil {
		return 0, err
	}
	// 发送创建商品消息
	if err := s.mq.SendCreateSpuInfo(ctx, &model.SpuEs{
		Id: spuId,
		UserId: spu.UserId,
		Name: spu.Name,
		Description: spu.Description,
		CategoryId: spu.CategoryId,
		MainImageURL: spu.MainImageURL,
		Price: spu.Price,
	}); err != nil {
		logger.Error(ctx, "send create spu info error: %v", err)
	}
	return spuId, nil
}

func (s *ProductUsecase) UpdateSpu(ctx context.Context, spu *model.Spu) error {
	// 更新商品
	if err := s.db.UpdateSpu(ctx, spu); err != nil {
		return err
	}

	if err := s.mq.SendUpdateSpuInfo(ctx, &model.SpuEs{
		Id: spu.Id,
		UserId: spu.UserId,
		Name: spu.Name,
		Description: spu.Description,
		CategoryId: spu.CategoryId,
		MainImageURL: spu.MainImageURL,
		Price: spu.Price,
	}); err != nil {
		logger.Error(ctx, "send update spu info error: %v", err)
	}
	return nil
}

func (s *ProductUsecase) UpdateSku(ctx context.Context, sku *model.Sku) error {
	// 更新商品sku
	if err := s.db.UpdateSku(ctx, sku); err != nil {
		return err
	}
	return nil
}

func (s *ProductUsecase) DeleteSpu(ctx context.Context, spuId int64) error {
	// 删除商品
	if err := s.db.DeleteSpu(ctx, spuId); err != nil {
		return err
	}

	if err := s.mq.SendDeleteSpuInfo(ctx, spuId); err != nil {
		logger.Error(ctx, "send delete spu info error: %v", err)
	}
	return nil
}

func(s *ProductUsecase) VerifySpu (spu *model.Spu) error {
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

func(s *ProductUsecase) CreateCategory(ctx context.Context, category *model.Category) (int64, error) {
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

func (s *ProductUsecase) DeleteCategory(ctx context.Context, categoryId int64) error { 
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

func (s *ProductUsecase) UpdateCategory(ctx context.Context, category *model.Category) error {
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



func (s *ProductUsecase) ConsumeCreateSpuInfo(ctx context.Context) error { 
	return s.mq.ConsumeCreateSpuInfo(ctx, s.es.AddSpuItem)
}

func (s *ProductUsecase) ConsumeUpdateSpuInfo(ctx context.Context) error { 
	return s.mq.ConsumeUpdateSpuInfo(ctx, s.es.UptateSpuItem)
}

func (s *ProductUsecase) ConsumeDeleteSpuInfo(ctx context.Context) error { 
	return s.mq.ConsumeDeleteSpuInfo(ctx, s.es.DeleteSpuItem)
}








