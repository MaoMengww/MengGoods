package product

import (
	"MengGoods/app/gateway/rpc"
	"MengGoods/kitex_gen/product"
	"MengGoods/pkg/base"
	"MengGoods/pkg/merror"
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
		base.ResErr(app, merror.NewMerror(merror.ParamIDInvalid, err.Error()))
		return
	}
	resp, err := rpc.UpdateSpu(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
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
	resp, err := rpc.DeleteSpu(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
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
	resp, err := rpc.UpdateSku(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
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
	resp, err := rpc.DeleteSku(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
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
	resp, err := rpc.DeleteCategory(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
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
	resp, err := rpc.UpdateCategory(ctx, &req)
	if err != nil {
		base.ResErr(app, err)
		return
	}
	base.ResData(app, resp)
}
