package rpc

import (
	mresp "MengGoods/app/gateway/model/resp"
	"MengGoods/config"
	"MengGoods/kitex_gen/user"
	"MengGoods/kitex_gen/user/userservice"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"

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
)

var UserClient userservice.Client

func UserInit() {
	r, err := etcd.NewEtcdResolver(config.Conf.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("user rpc Init Falied: err: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := userservice.NewClient(
		"user",
		client.WithResolver(r),
		client.WithRPCTimeout(30*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithCircuitBreaker(cbSuite),
		client.WithSuite(tracing.NewClientSuite()),
		//
	)

	if err != nil {
		logger.Fatalf("init client failed: err:%v", err)
	}

	UserClient = c
}

func Register(ctx context.Context, req *user.RegisterReq) (resp *mresp.RegisterResp, err error) {
	r, err := UserClient.Register(ctx, req)
	if err != nil {
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	//base是服务器内部字段, 不需要返回给前端
	resp = &mresp.RegisterResp{
		UserId: r.UserId,
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq) (resp *mresp.LoginResp, err error) {
	r, err := UserClient.Login(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	accessToken, refresh, err := utils.CreateGatewayToken(r.UserInfo.Id)
	if err != nil {
		logger.CtxError(ctx, err)
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	//base是服务器内部字段, 不需要返回给前端
	resp = &mresp.LoginResp{
		User:         *r.UserInfo,
		AccessToken:  accessToken,
		RefreshToken: refresh,
	}
	return resp, nil
}


func AddAddress(ctx context.Context, req *user.AddAddressReq) (resp *mresp.AddAddressResp, err error) {
	r, err := UserClient.AddAddress(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	//base是服务器内部字段, 不需要返回给前端
	resp = &mresp.AddAddressResp{
		AddressId: r.AddressId,
	}
	return resp, nil
}

func GetAddressList(ctx context.Context, req *user.GetAddressesReq) (resp *mresp.GetAddressResp, err error) {
	r, err := UserClient.GetAddresses(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	//base是服务器内部字段, 不需要返回给前端
	resp = &mresp.GetAddressResp{
		AddressList: r.Address,
	}
	return resp, nil
}

func BanUser(ctx context.Context, c *user.BanUserReq) (resp *mresp.BanUserResp, err error) {
	r, err := UserClient.BanUser(ctx, c)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	//base是服务器内部字段, 不需要返回给前端
	return resp, nil
}

func UnBanUser(ctx context.Context, c *user.UnBanUserReq) (resp *mresp.UnBanUserResp, err error) {
	r, err := UserClient.UnBanUser(ctx, c)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return resp, nil
}

func SetAdmin(ctx context.Context, c *user.SetAdminReq) (resp *mresp.SetAdminResp, err error) {
	r, err := UserClient.SetAdmin(ctx, c)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return resp, nil
}

func GetUserInfo(ctx context.Context, c *user.GetUserInfoReq) (resp *mresp.GetUserInfoResp, err error) {
	ctx = mcontext.WithUserIDInContext(ctx, c.UserId)
	r, err := UserClient.GetUserInfo(ctx, c)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	//base是服务器内部字段, 不需要返回给前端
	resp = &mresp.GetUserInfoResp{
		User: *r.UserInfo,
	}
	return resp, nil
}

func Logout(ctx context.Context, c *user.LogoutReq) (resp *mresp.LogoutResp, err error) {
	r, err := UserClient.Logout(ctx, c)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return resp, nil
}

func SendCode(ctx context.Context, c *user.SendCodeReq) (resp *mresp.SendCodeResp, err error) {
	r, err := UserClient.SendCode(ctx, c)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return resp, nil
}

func ResetPwd(ctx context.Context, c *user.ResetPwdReq) (resp *mresp.ResetPwdResp, err error) {
	r, err := UserClient.ResetPwd(ctx, c)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	return resp, nil
}

func UploadAvatar(ctx context.Context, c *user.UploadAvatarReq) (resp *mresp.UploadAvatarResp, err error) {
	r, err := UserClient.UploadAvatar(ctx, c)
	if err != nil {
		logger.CtxErrorf(ctx, err.Error())
		return nil, merror.NewMerror(merror.InternalServerErrorCode, err.Error())
	}
	if r.Base.Code != merror.SuccessCode {
		return nil, merror.NewMerror(r.Base.Code, r.Base.Message)
	}
	resp = &mresp.UploadAvatarResp{
		Url: r.AvatarURL,
	}
	return resp, nil
}
