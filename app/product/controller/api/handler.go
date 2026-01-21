package api

import (
	"MengGoods/app/product/controller/api/pack"
	"MengGoods/app/product/domain/model"
	"MengGoods/app/product/usecase"
	product "MengGoods/kitex_gen/product"
	"MengGoods/pkg/base"
	"context"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct {
	usecase *usecase.ProductUsecase
}

func NewProductServiceImpl(usecase *usecase.ProductUsecase) *ProductServiceImpl {
	return &ProductServiceImpl{
		usecase: usecase,
	}
}

// CreateSpu implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateSpu(ctx context.Context, req *product.CreateSpuReq) (resp *product.CreateSpuResp, err error) {
	resp = new(product.CreateSpuResp)
	resp.SpuId, err = s.usecase.CreateSpu(ctx, &model.Spu{
		Name:            req.Name,
		Description:     req.Description,
		CategoryId:      req.CategoryId,
		MainImageURL:    req.MainSpuImageURL,
		SliderImageURLs: req.SliderSpuImageURLs,
		Skus: pack.BuildSkus(req.Sku),
	})
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// UpdateSpu implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateSpu(ctx context.Context, req *product.UpdateSpuReq) (resp *product.UpdateSpuResp, err error) {
	resp = new(product.UpdateSpuResp)
	err = s.usecase.UpdateSpu(ctx, &model.Spu{
		Id:              req.SpuId,
		Name:            *req.Name,
		Description:     *req.Description,
		CategoryId:      *req.CategoryId,
		MainImageURL:    *req.MainSpuImageURL,
		SliderImageURLs: *req.SliderSpuImageURLs,
	})
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// UpdateSku implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateSku(ctx context.Context, req *product.UpdateSkuReq) (resp *product.UpdateSkuResp, err error) {
	resp = new(product.UpdateSkuResp)
	err = s.usecase.UpdateSku(ctx, &model.Sku{
		Id:          req.SkuId,
		Name:        *req.Name,
		Description: *req.Description,
		Properties:  *req.Properties,
		ImageURL:    *req.SkuImageURL,
		Price:       *req.Price,
	})
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// DeleteSpu implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteSpu(ctx context.Context, req *product.DeleteSpuReq) (resp *product.DeleteSpuResp, err error) {
	resp = new(product.DeleteSpuResp)
	err = s.usecase.DeleteSpu(ctx, req.SpuId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return  resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return  resp, nil
}

// DeleteSku implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteSku(ctx context.Context, req *product.DeleteSkuReq) (resp *product.DeleteSkuResp, err error) {
	resp = new(product.DeleteSkuResp)
	err = s.usecase.DeleteSku(ctx, req.SkuId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// GetSpuById implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetSpuById(ctx context.Context, req *product.GetSpuByIdReq) (resp *product.GetSpuByIdResp, err error) {
	resp = new(product.GetSpuByIdResp)
	spu, err := s.usecase.GetSpuById(ctx, req.SpuId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	resp.SpuInfo = pack.BuildSpuInfo(spu)
	return resp, nil
}

// GetSku implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetSku(ctx context.Context, req *product.GetSkuReq) (resp *product.GetSkuResp, err error) {
	resp = new(product.GetSkuResp)
	sku, err := s.usecase.GetSkuById(ctx, req.SkuId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	resp.SkuInfo = pack.BuildSkuInfo(sku)
	return resp, nil
}

// CreateCategory implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateCategory(ctx context.Context, req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {
	resp = new(product.CreateCategoryResp)
	resp.CategoryId, err = s.usecase.CreateCategory(ctx, &model.Category{
		Name: req.Name,
	})
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// UpdateCategory implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateCategory(ctx context.Context, req *product.UpdateCategoryReq) (resp *product.UpdateCategoryResp, err error) {
	// TODO: Your code here...
	resp = new(product.UpdateCategoryResp)
	err = s.usecase.UpdateCategory(ctx, &model.Category{
		Id:   req.CategoryId,
		Name: *req.Name,
	})
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return resp, nil
}

// DeleteCategory implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq) (resp *product.DeleteCategoryResp, err error) {
	resp = new(product.DeleteCategoryResp)
	err = s.usecase.DeleteCategory(ctx, req.CategoryId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return  resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return  resp, nil
}

// GetSpuList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetSpuList(ctx context.Context, req *product.GetSpuReq) (resp *product.GetSpuResp, err error) {
	resp = new(product.GetSpuResp)
	spus, total, err := s.usecase.GetSpuList(ctx, req)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return  resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	resp.SpuList = pack.BuildSpuInfoList(spus)
	resp.Total = total
	return   resp, nil
}
