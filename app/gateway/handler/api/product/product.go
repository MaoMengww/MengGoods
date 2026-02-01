package product

import (
	"MengGoods/app/gateway/rpc"
	"MengGoods/kitex_gen/product"
	"MengGoods/pkg/base"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"

	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func CreateSpu(ctx context.Context, app *app.RequestContext) {
	var req product.CreateSpuReq
	err := app.BindAndValidate(&req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	resp, err := rpc.CreateSpu(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
}

func UpdateSpu(ctx context.Context, app *app.RequestContext) {
	var req product.UpdateSpuReq
	err := app.Bind(&req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	idStr := app.Param("spu_id")
	req.SpuId, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		base.ResErr(app, merror.NewMerror(merror.ParamIDInvalid, idStr + " is invalid"))
		return
	}
	_, err = rpc.UpdateSpu(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResSuccess(app)
}

func DeleteSpu(ctx context.Context, app *app.RequestContext) {
	var req product.DeleteSpuReq
	err := app.Bind(&req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	idStr := app.Param("spu_id")
	req.SpuId, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		base.ResErr(app, merror.NewMerror(merror.ParamIDInvalid, err.Error()))
		return
	}
	_, err = rpc.DeleteSpu(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResSuccess(app)
}

func UpdateSku(ctx context.Context, app *app.RequestContext) {
	var req product.UpdateSkuReq
	err := app.Bind(&req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	idStr := app.Param("sku_id")
	req.SkuId, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		base.ResErr(app, merror.NewMerror(merror.ParamIDInvalid, err.Error()))
		return
	}
	_, err = rpc.UpdateSku(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResSuccess(app)
}

func DeleteSku(ctx context.Context, app *app.RequestContext) {
	var req product.DeleteSkuReq
	err := app.Bind(&req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	idStr := app.Param("sku_id")
	req.SkuId, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		base.ResErr(app, merror.NewMerror(merror.ParamIDInvalid, err.Error()))
		return
	}
	_, err = rpc.DeleteSku(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResSuccess(app)
}

func GetSpuById(ctx context.Context, app *app.RequestContext) {
	var req product.GetSpuByIdReq
	idStr := app.Param("spu_id")
	spuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		base.ResErr(app, merror.NewMerror(merror.ParamIDInvalid, err.Error()))
		return
	}
	req.SpuId = spuId
	resp, err := rpc.GetSpuById(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
}

func GetSku(ctx context.Context, app *app.RequestContext) {
	var req product.GetSkuReq
	idStr := app.Param("sku_id")
	skuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		base.ResErr(app, merror.NewMerror(merror.ParamIDInvalid, err.Error()))
		return
	}
	req.SkuId = skuId
	resp, err := rpc.GetSku(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
}

func CreateCategory(ctx context.Context, app *app.RequestContext) {
	var req product.CreateCategoryReq
	err := app.BindAndValidate(&req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	resp, err := rpc.CreateCategory(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
}

func DeleteCategory(ctx context.Context, app *app.RequestContext) {
	var req product.DeleteCategoryReq
	err := app.BindAndValidate(&req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	_, err = rpc.DeleteCategory(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResSuccess(app)
}

func GetSpu(ctx context.Context, app *app.RequestContext) {
	var req product.GetSpuReq
	err := app.BindAndValidate(&req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	resp, err := rpc.GetSpuList(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
}

func UpdateCategory(ctx context.Context, app *app.RequestContext) {
	var req product.UpdateCategoryReq
	err := app.BindAndValidate(&req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	_, err = rpc.UpdateCategory(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResSuccess(app)
}

func UploadSkuImage(ctx context.Context, app *app.RequestContext) {
	var req product.UploadSkuImageReq
	pathSkuId := app.Param("sku_Id")
	skuId, err := strconv.ParseInt(pathSkuId, 10, 64)
	if err != nil {
		base.ResErr(app, merror.NewMerror(merror.ParamSkuIdInvalid, "invalid sku id"))
		return
	}
	req.SkuId = skuId
	file, err := app.FormFile("image")
	if err != nil {
		base.ResErr(app, merror.NewMerror(merror.InvalidImageFileType, "image file type not supported"))
		return
	}
	if !utils.CheckImageFileType(file) {
		base.ResErr(app, merror.NewMerror(merror.InvalidImageFileType, "image file type not supported"))
		return
	}
	datas, err := utils.FileToBytes(file)
	if err != nil {
		base.ResErr(app, merror.NewMerror(merror.InvalidImageFileType, "image file type not supported"))
		return
	}
	req.ImageData = datas
	req.ImageName = file.Filename
	_, err = rpc.UploadSkuImage(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResSuccess(app)
}

func UploadSpuImage(ctx context.Context, c *app.RequestContext) {
	var req product.UploadSpuImageReq
	pathSpuId := c.Param("spu_Id")
	spuId, err := strconv.ParseInt(pathSpuId, 10, 64)
	if err != nil {
		base.ResErr(c, merror.NewMerror(merror.ParamSpuIdInvalid, "invalid spu id"))
		return
	}
	req.SpuId = spuId
	file, err := c.FormFile("image")
	if err != nil {
		base.ResErr(c, merror.NewMerror(merror.InvalidImageFileType, "image file type not supported"))
		return
	}
	if !utils.CheckImageFileType(file) {
		base.ResErr(c, merror.NewMerror(merror.InvalidImageFileType, "image file type not supported"))
		return
	}
	datas, err := utils.FileToBytes(file)
	if err != nil {
		base.ResErr(c, merror.NewMerror(merror.InvalidImageFileType, "image file type not supported"))
		return
	}
	req.ImageData = datas
	req.ImageName = file.Filename
	_, err = rpc.UploadSpuImage(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}
