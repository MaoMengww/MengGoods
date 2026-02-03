package mysql

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
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

var g singleflight.Group

func (p *PrductDB) GetAllSpuIdAndSkuId(ctx context.Context)(spuIds []int64, skuIds []int64, err error) {
	err = p.db.WithContext(ctx).Model(&Spu{}).Pluck("spu_id", &spuIds).Error
	if err != nil {
		return nil, nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get all spu id failed: %v", err))
	}
	err = p.db.WithContext(ctx).Model(&Sku{}).Pluck("sku_id", &skuIds).Error
	if err != nil {
		return nil, nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get all sku id failed: %v", err))
	}
	return spuIds, skuIds, nil
}

func (p *PrductDB) CreateSpu(ctx context.Context, spu *model.Spu) (int64, error) {
	spu_id := utils.GenerateID()
	creatorId, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return 0, merror.NewMerror(merror.InternalServerErrorCode, fmt.Sprintf("get user id failed: %v", err))
	}
	spuDB := &Spu{
		SpuId:           spu_id,
		Name:            spu.Name,
		Price:           spu.Price,
		Description:     spu.Description,
		MainImageURL:    spu.MainImageURL,
		SliderImageURLs: spu.SliderImageURLs,
		Creator:         creatorId,
		Category:        int(spu.CategoryId),
		Status:          int32(spu.Status),
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
	}
	tx := p.db.WithContext(ctx).Begin()
	err = tx.WithContext(ctx).Create(spuDB).Error
	if err != nil {
		tx.Rollback()
		return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("create spu failed: %v", err))
	}
	for _, sku := range spu.Skus {
		skuId := utils.GenerateID()
		skuDB := &Sku{
			SkuId:       skuId,
			Name:        sku.Name,
			Price:       sku.Price,
			Description: sku.Description,
			ImageURL:    sku.ImageURL,
			Properties:  sku.Properties,
			SpuID:       spu_id,
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		}
		err := tx.WithContext(ctx).Create(skuDB).Error
		if err != nil {
			tx.Rollback()
			return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("create sku failed: %v", err))
		}
	}
	tx.Commit()
	return spuDB.SpuId, nil
}

func (p *PrductDB) CreateCategory(ctx context.Context, category *model.Category) (int64, error) {
	categoryId := utils.GenerateID()
	categoryDB := &Category{
		CategoryId: categoryId,
		Name:       category.Name,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}
	err := p.db.WithContext(ctx).Create(categoryDB).Error
	if err != nil {
		return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("db create category error: %v", err))
	}
	return categoryId, nil
}

func (p *PrductDB) DeleteSpu(ctx context.Context, spuId int64) error {
	// 删除spu - 使用正确的字段名
	err := p.db.WithContext(ctx).Where("spu_id = ?", spuId).Delete(&Spu{}).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("delete spu failed: %v", err))
	}
	// 删除sku
	err = p.db.WithContext(ctx).Where("spu_id = ?", spuId).Delete(&Sku{}).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("delete sku failed: %v", err))
	}
	return nil
}

func (p *PrductDB) DeleteSku(ctx context.Context, skuId int64) error {
	// 删除sku
	err := p.db.WithContext(ctx).Delete(&Sku{}, skuId).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("delete sku failed: %v", err))
	}
	return nil
}

func (p *PrductDB) DeleteCategory(ctx context.Context, categoryId int64) error {
	// 删除category
	err := p.db.WithContext(ctx).Delete(&Category{}, categoryId).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("delete category failed: %v", err))
	}
	return nil
}

func (p *PrductDB) UpdateSpu(ctx context.Context, spu *model.Spu) error {
	// 更新spu - 使用正确的字段名
	err := p.db.WithContext(ctx).Model(&Spu{}).Where("spu_id = ?", spu.SpuId).Updates(map[string]interface{}{
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
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("update spu failed: %v", err))
	}
	return nil
}

func (p *PrductDB) UpdateSku(ctx context.Context, sku *model.Sku) error {
	// 更新sku
	err := p.db.WithContext(ctx).Model(&Sku{}).Where("sku_id = ?", sku.SkuId).Updates(map[string]interface{}{
		"name":        sku.Name,
		"price":       sku.Price,
		"description": sku.Description,
		"image_url":   sku.ImageURL,
		"properties":  sku.Properties,
		"updated_at":   time.Now(),
	}).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("update sku failed: %v", err))
	}
	return nil
}
func (p *PrductDB) UpdateCategory(ctx context.Context, category *model.Category) error {
	// 更新category
	err := p.db.WithContext(ctx).Model(&Category{}).Where("category_id = ?", category.Id).Updates(map[string]interface{}{
		"name":      category.Name,
		"updated_at": time.Now(),
	}).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("update category failed: %v", err))
	}
	return nil
}
func (p *PrductDB) GetSpuById(ctx context.Context, spuId int64) (*model.Spu, error) {
	// 查询spu
	var spu Spu
	groupKey := fmt.Sprintf("get_spu_%d", spuId)
	v, err, shared := g.Do(groupKey, func() (interface{}, error) {
		err := p.db.WithContext(ctx).Where("spu_id = ?", spuId).First(&spu).Error
		if err != nil {
			return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get spu failed: %v", err))
		}
		return &model.Spu{
			SpuId:           spu.SpuId,
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
	})
	if err != nil {
		return nil, err
	}
	if shared {
		fmt.Printf("Cache hit for spu_id %d\n", spuId)
	}
	return v.(*model.Spu), nil
}

func (p *PrductDB) GetSkuById(ctx context.Context, skuId int64) (*model.Sku, error) {
	// 获取sku
	var sku Sku
	err := p.db.WithContext(ctx).Where("sku_id = ?", skuId).First(&sku).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get sku failed: %v", err))
	}
	return &model.Sku{
		SkuId:       sku.SkuId,
		SpuId:       sku.SpuID,
		Name:        sku.Name,
		Description: sku.Description,
		Properties:  sku.Properties,
		ImageURL:    sku.ImageURL,
		Price:       sku.Price,
		CreateTime:  sku.CreateAt.Unix(),
		UpdateTime:  sku.UpdateAt.Unix(),
	}, nil
}

func (p *PrductDB) GetSpusByIds(ctx context.Context, spuIds []int64) ([]*model.Spu, error) {
	// 批量查询spu
	var spus []*model.Spu
	err := p.db.WithContext(ctx).Where("spu_id in ?", spuIds).Find(&spus).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get spu failed: %v", err))
	}
	return spus, nil
}

func (p *PrductDB) GetSkusByIds(ctx context.Context, skuIds []int64) ([]*model.Sku, error) {
	// 批量查询sku
	var skuModels []*Sku
	err := p.db.WithContext(ctx).Where("sku_id in ?", skuIds).Find(&skuModels).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get sku failed: %v", err))
	}

	// 转换数据库模型到域模型
	var skus []*model.Sku
	for _, skuModel := range skuModels {
		sku := &model.Sku{
			SkuId:       skuModel.SkuId,
			SpuId:       skuModel.SpuID,
			Name:        skuModel.Name,
			Description: skuModel.Description,
			Properties:  skuModel.Properties,
			ImageURL:    skuModel.ImageURL,
			Price:       skuModel.Price,
			CreateTime:  skuModel.CreateAt.Unix(),
			UpdateTime:  skuModel.UpdateAt.Unix(),
		}
		skus = append(skus, sku)
	}
	return skus, nil
}

func (p *PrductDB) GetSkusBySpuId(ctx context.Context, spuId int64) ([]*model.Sku, error) {
	// 查询sku
	var skus []*Sku
	err := p.db.WithContext(ctx).Where("spu_id = ?", spuId).Find(&skus).Error
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get sku failed: %v", err))
	}
	// 转换数据库模型到域模型
	var skusResp []*model.Sku
	for _, skuModel := range skus {
		sku := &model.Sku{
			SkuId:       skuModel.SkuId,
			SpuId:       skuModel.SpuID,
			Name:        skuModel.Name,
			Description: skuModel.Description,
			Properties:  skuModel.Properties,
			ImageURL:    skuModel.ImageURL,
			Price:       skuModel.Price,
			CreateTime:  skuModel.CreateAt.Unix(),
			UpdateTime:  skuModel.UpdateAt.Unix(),
		}
		skusResp = append(skusResp, sku)
	}
	return skusResp, nil
}
func (p *PrductDB) IsSpuExist(ctx context.Context, spuId int64) (bool, error) {
	// 判断spu是否存在
	var count int64
	err := p.db.WithContext(ctx).Model(&Spu{}).Where("spu_id = ?", spuId).Count(&count).Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get spu failed: %v", err))
	}
	return count > 0, nil
}

func (p *PrductDB) IsSpuOwer(ctx context.Context, spuId int64) (bool, error) {
	//查询spu的userid
	// 查询spu的userid
	var spu Spu
	err := p.db.WithContext(ctx).Where("spu_id = ?", spuId).First(&spu).Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get spu failed: %v", err))
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
		Where("sku.sku_id = ?", skuId).                     // 4. 条件 (WHERE s_id = ?)
		Scan(&creatorId).                                   // 5. 结果赋值给变量
		Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get sku failed: %v", err))
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
	err := p.db.WithContext(ctx).Model(&Sku{}).Where("sku_id = ?", skuId).Count(&count).Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get sku failed: %v", err))
	}
	return count > 0, nil
}

func (p *PrductDB) IsCategoryExist(ctx context.Context, categoryId int64) (bool, error) {
	// 删除category
	var count int64
	err := p.db.WithContext(ctx).Model(&Category{}).Where("category_id = ?", categoryId).Count(&count).Error
	if err != nil {
		return false, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("get category failed: %v", err))
	}
	return count > 0, nil
}

func (p *PrductDB) UploadSpuImage(ctx context.Context, spuId int64, imageUrl string) error {
	// 删除spu
	err := p.db.WithContext(ctx).Model(&Spu{}).Where("spu_id = ?", spuId).Update("main_image_url", imageUrl).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("update spu failed: %v", err))
	}
	return nil
}

func (p *PrductDB) UploadSkuImage(ctx context.Context, skuId int64, imageUrl string) error {
	// 删除spu
	err := p.db.WithContext(ctx).Model(&Sku{}).Where("sku_id = ?", skuId).Update("image_url", imageUrl).Error
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("update sku failed: %v", err))
	}
	return nil
}
