package rpc

import (
	mresp "MengGoods/app/gateway/model/resp"
	mModel "MengGoods/app/product/domain/model"
	"MengGoods/kitex_gen/model"
	"MengGoods/kitex_gen/product"
	"MengGoods/kitex_gen/product/productservice"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	//"MengGoods/pkg/middleware"
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
)

var ProductClient productservice.Client

func ProductInit() {
	r, err := etcd.NewEtcdResolver(viper.GetStringSlice("etcd.endpoints"))
	if err != nil {
		logger.Fatalf("product rpc Init Falied: err: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := productservice.NewClient(
		"product",
		client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithCircuitBreaker(cbSuite),
		client.WithSuite(tracing.NewClientSuite()),
		//client.WithMiddleware(middleware.ErrorLog()),
	)

	if err != nil {
		logger.Fatalf("init client failed: err:%v", err)
	}

	ProductClient = c
}

func GetSpuById(ctx context.Context, req *product.GetSpuByIdReq) (resp *mresp.GetSpuByIdResp, err error) {
	r, err := ProductClient.GetSpuById(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	resp = &mresp.GetSpuByIdResp{
		Spu: *r.SpuInfo,
	}
	return
}

func CreateSpu (ctx context.Context, req *product.CreateSpuReq) (resp *mresp.CreateSpuResp, err error) {
	r, err := ProductClient.CreateSpu(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	resp = &mresp.CreateSpuResp{
		SpuId: r.SpuId,
	}
	return
}

func UpdateSpu(ctx context.Context, req *product.UpdateSpuReq) (resp *mresp.UpdateSpuResp, err error) {
	r, err := ProductClient.UpdateSpu(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}

func UpdateSku(ctx context.Context, req *product.UpdateSkuReq) (resp *mresp.UpdateSkuResp, err error) {
	r, err := ProductClient.UpdateSku(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}

func DeleteSpu(ctx context.Context, req *product.DeleteSpuReq) (resp *mresp.DeleteSpuResp, err error) {
	r, err := ProductClient.DeleteSpu(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}

func DeleteSku(ctx context.Context, req *product.DeleteSkuReq) (resp *mresp.DeleteSkuResp, err error) {
	r, err := ProductClient.DeleteSku(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}



func GetSku(ctx context.Context, req *product.GetSkuReq) (resp *mresp.GetSkuResp, err error) {
	r, err := ProductClient.GetSku(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	resp = &mresp.GetSkuResp{
		Sku: *r.SkuInfo,
	}
	return resp, nil
}

func CreateCategory(ctx context.Context, req *product.CreateCategoryReq) (resp *mresp.CreateCategoryResp, err error) {
	r, err := ProductClient.CreateCategory(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}

func UpdateCategory(ctx context.Context, req *product.UpdateCategoryReq) (resp *mresp.UpdateCategoryResp, err error) {
	r, err := ProductClient.UpdateCategory(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}

func DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq) (resp *mresp.DeleteCategoryResp, err error) {
	r, err := ProductClient.DeleteCategory(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return
}

func GetSpuList(ctx context.Context, req *product.GetSpuReq) (resp *mresp.GetSpuResp, err error) {
	r, err := ProductClient.GetSpuList(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	resp = &mresp.GetSpuResp{
		Spus: SpuList(r.SpuList),
		Total: r.Total,
	}
	return
}

func SpuList(spus []*model.SpuInfo) []*mModel.SpuEs {
	var list []*mModel.SpuEs
	for _, spu := range spus {
		item := &mModel.SpuEs{
			Id: spu.Id,
			Name: spu.Name,
			CategoryId: spu.CategoryId,
			Description: spu.Description,
			MainImageURL: spu.MainImageURL,
		}
			list = append(list, item)
		
	}
	return list
}


