package mysql

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PrductDB struct {
	db *gorm.DB
}

func NewProductDB(db *gorm.DB) *PrductDB {
	return &PrductDB{
		db: db,
	}
}

func (p *PrductDB) CreateSpu(ctx context.Context, spu *model.Spu) (int64, error) {
	spu_id := utils.GenerateID()
	spuDB := &Spu{
		ID:              spu_id,
		Name:            spu.Name,
		Price:           spu.Price,
		Description:     spu.Description,
		MainImageURL:    spu.MainImageURL,
		SliderImageURLs: spu.SliderImageURLs,
		Creator:         spu.UserId,
		Category:        int(spu.CategoryId),
		Status:          int32(spu.Status),
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
	}
	err := p.db.WithContext(ctx).Create(spuDB).Error
	if err != nil {
		return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, "创建spu失败")
	}

	for _, sku := range spu.Skus {
		skuDB := &Sku{
			ID:          utils.GenerateID(),
			Name:        sku.Name,
			Price:       sku.Price,
			Description: sku.Description,
			ImageURL:    sku.ImageURL,
			Properties:  sku.Properties,
			Sale:        sku.Sale,
			SpuID:       spu_id,
			UpdateAt:    time.Now(),
			CreateAt:    time.Now(),
		}
		err := p.db.WithContext(ctx).Create(skuDB).Error
		if err != nil {
			return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("创建sku失败: %v", err))
		}
	}
	return spuDB.ID, nil
}

func (p *PrductDB) CreateCategory(ctx context.Context, category *model.Category) (int64, error) {
	categoryDB := &Category{
		ID:       category.Id,
		Name:     category.Name,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err := p.db.WithContext(ctx).Create(categoryDB).Error
	if err != nil {
		return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db create category error: %v", err))
	}
	return categoryDB.ID, nil
}

func (p *PrductDB) DeleteSpu(ctx context.Context, spuId int64) error {
	// 删除spu - 使用正确的字段名
	err := p.db.WithContext(ctx).Where("spu_id = ?", spuId).Delete(&Spu{}).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("删除spu失败: %v", err))
	}
	// 删除sku
	err = p.db.WithContext(ctx).Where("spu_id = ?", spuId).Delete(&Sku{}).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("删除sku失败: %v", err))
	}
	return nil
}

func (p *PrductDB) DeleteSku(ctx context.Context, skuId int64) error {
	// 删除sku
	err := p.db.WithContext(ctx).Delete(&Sku{}, skuId).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("删除sku失败: %v", err))
	}
	return nil
}

func (p *PrductDB) DeleteCategory(ctx context.Context, categoryId int64) error {
	// 删除category
	err := p.db.WithContext(ctx).Delete(&Category{}, categoryId).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("删除category失败: %v", err))
	}
	return nil
}

func (p *PrductDB) UpdateSpu(ctx context.Context, spu *model.Spu) error {
	// 更新spu - 使用正确的字段名
	err := p.db.WithContext(ctx).Model(&Spu{}).Where("spu_id = ?", spu.Id).Updates(map[string]interface{}{
		"name":              spu.Name,
		"price":             spu.Price,
		"description":       spu.Description,
		"main_image_url":    spu.MainImageURL,
		"slider_image_urls": spu.SliderImageURLs,
		"category":          int(spu.CategoryId),
		"status":            int32(spu.Status),
		"updated_at":        time.Now(),
	}).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("更新spu失败: %v", err))
	}
	return nil
}

func (p *PrductDB) UpdateSku(ctx context.Context, sku *model.Sku) error {
	// 更新sku
	err := p.db.WithContext(ctx).Model(&Sku{}).Where("id = ?", sku.Id).Updates(map[string]interface{}{
		"name":        sku.Name,
		"price":       sku.Price,
		"description": sku.Description,
		"image_url":   sku.ImageURL,
		"properties":  sku.Properties,
		"update_at":   time.Now().Unix(),
	}).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("更新sku失败: %v", err))
	}
	return nil
}
func (p *PrductDB) UpdateCategory(ctx context.Context, category *model.Category) error {
	// 删除category
	err := p.db.WithContext(ctx).Model(&Category{}).Where("id = ?", category.Id).Updates(map[string]interface{}{
		"name":      category.Name,
		"update_at": time.Now().Unix(),
	}).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("更新category失败: %v", err))
	}
	return nil
}
func (p *PrductDB) GetSpuById(ctx context.Context, spuId int64) (*model.Spu, error) {
	// 查询spu
	var spu Spu
	err := p.db.WithContext(ctx).Where("spu_id = ?", spuId).First(&spu).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("查询spu失败: %v", err))
	}
	return &model.Spu{
		Id:              spu.ID,
		UserId:          spu.Creator,
		Name:            spu.Name,
		Description:     spu.Description,
		CategoryId:      int64(spu.Category),
		MainImageURL:    spu.MainImageURL,
		SliderImageURLs: spu.SliderImageURLs,
		CreateTime:      spu.CreateAt,
		UpdateTime:      spu.UpdateAt,
		Status:          int32(spu.Status),
		Price:           spu.Price,
	}, nil
}

func (p *PrductDB) GetSkuById(ctx context.Context, skuId int64) (*model.Sku, error) {
	// 获取sku
	var sku model.Sku
	err := p.db.WithContext(ctx).Where("id = ?", skuId).First(&sku).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("查询sku失败: %v", err))
	}
	return &sku, nil
}

func (p *PrductDB) GetSpusByIds(ctx context.Context, spuIds []int64) ([]*model.Spu, error) {
	// 批量查询spu
	var spus []*model.Spu
	err := p.db.WithContext(ctx).Where("spu_id in ?", spuIds).Find(&spus).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("批量查询spu失败: %v", err))
	}
	return spus, nil
}

func (p *PrductDB) GetSkusByIds(ctx context.Context, skuIds []int64) ([]*model.Sku, error) {
	// 批量查询sku
	var skus []*model.Sku
	err := p.db.WithContext(ctx).Where("id in ?", skuIds).Find(&skus).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("批量查询sku失败: %v", err))
	}
	return skus, nil
}

func (p *PrductDB) GetSkusBySpuId(ctx context.Context, spuId int64) ([]*model.Sku, error) {
	// 查询sku
	var skus []*model.Sku
	err := p.db.WithContext(ctx).Where("spu_id = ?", spuId).Find(&skus).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("查询sku失败: %v", err))
	}
	return skus, nil
}
func (p *PrductDB) IsSpuExist(ctx context.Context, spuId int64) (bool, error) {
	// 判断spu是否存在
	var count int64
	err := p.db.WithContext(ctx).Model(&Spu{}).Where("spu_id = ?", spuId).Count(&count).Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("判断spu是否存在失败: %v", err))
	}
	return count > 0, nil
}

func (p *PrductDB) IsSpuOwer(ctx context.Context, spuId int64) (bool, error) {
	//查询spu的userid
	// 查询spu的userid
	var spu Spu
	err := p.db.WithContext(ctx).Where("spu_id = ?", spuId).First(&spu).Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("查询spu失败: %v", err))
	}
	uid, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	return spu.Creator == uid, nil
}

func (p *PrductDB) IsSkuOwer(ctx context.Context, skuId int64) (bool, error) {
	// 查询sku的userid
	var creatorId int64

	err := p.db.WithContext(ctx).
		Table("sku").                                       // 1. 指定主表
		Select("spu.creator").                              // 2. 指定只想查这一个字段 (对应 SQL SELECT p.creator)
		Joins("INNER JOIN spu ON sku.spu_id = spu.spu_id"). // 3. 写连表逻辑 (对应 INNER JOIN)
		Where("sku.id = ?", skuId).                         // 4. 条件 (WHERE s.id = ?)
		Scan(&creatorId).                                   // 5. 结果赋值给变量 (不要用 First(&sku))
		Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("查询sku失败: %v", err))
	}
	uid, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	return creatorId == uid, nil
}

func (p *PrductDB) IsSkuExist(ctx context.Context, skuId int64) (bool, error) {
	// 判断sku是否存在
	var count int64
	err := p.db.WithContext(ctx).Model(&Sku{}).Where("id = ?", skuId).Count(&count).Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("判断sku是否存在失败: %v", err))
	}
	return count > 0, nil
}

func (p *PrductDB) IsCategoryExist(ctx context.Context, categoryId int64) (bool, error) {
	// 删除category
	var count int64
	err := p.db.WithContext(ctx).Model(&Category{}).Where("id = ?", categoryId).Count(&count).Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("判断category是否存在失败: %v", err))
	}
	return count > 0, nil
}
