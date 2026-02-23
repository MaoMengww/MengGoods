package api

import (
	"MengGoods/app/gateway/rpc"
	"MengGoods/kitex_gen/user"
	"MengGoods/pkg/base"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var req user.RegisterReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.Register(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func Login(ctx context.Context, c *app.RequestContext) {
	var req user.LoginReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.Login(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	accessToken, refresh, err := utils.CreateGatewayToken(resp.User.Id)
	if err != nil {
		logger.CtxError(ctx, err)
		base.ResErr(c, err)
		return
	}
	c.Header(constants.AccessTokenHeader, accessToken)
	c.Header(constants.RefreshTokenHeader, refresh)
	base.ResData(c, resp)
}

func RefreshToken(ctx context.Context, c *app.RequestContext) {
	refreshToken := c.GetHeader(constants.RefreshTokenHeader)
	if len(refreshToken) == 0 {
		base.ResErr(c, merror.NewMerror(merror.InvalidToken, "refresh token is required"))
		return
	}
	claims, err := utils.CheckToken(string(refreshToken))
	if err != nil {
		logger.CtxError(ctx, err)
		base.ResErr(c, merror.NewMerror(merror.InvalidToken, err.Error()))
		return
	}
	newAccessToken, newRefresh, err := utils.CreateGatewayToken(claims.Uid)
	if err != nil {
		logger.CtxError(ctx, err)
		base.ResErr(c, merror.NewMerror(merror.InvalidToken, err.Error()))
		return
	}

	c.Header(constants.AccessTokenHeader, newAccessToken)
	c.Header(constants.RefreshTokenHeader, newRefresh)

	base.ResSuccess(c)
}

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	uid, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		base.ResErr(c, err)
		return
	}

	// 使用从context获取的用户ID构建请求
	req := user.GetUserInfoReq{
		UserId: uid,
	}

	err = c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.GetUserInfo(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func AddAddress(ctx context.Context, c *app.RequestContext) {
	var req user.AddAddressReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.AddAddress(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func GetAddressList(ctx context.Context, c *app.RequestContext) {
	var req user.GetAddressesReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err := rpc.GetAddressList(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}

func BanUser(ctx context.Context, c *app.RequestContext) {
	var req user.BanUserReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	_, err = rpc.BanUser(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}

func UnBanUser(ctx context.Context, c *app.RequestContext) {
	var req user.UnBanUserReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	_, err = rpc.UnBanUser(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}

func SetAdmin(ctx context.Context, c *app.RequestContext) {
	var req user.SetAdminReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	_, err = rpc.SetAdmin(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}

func Logout(ctx context.Context, c *app.RequestContext) {
	var req user.LogoutReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	_, err = rpc.Logout(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}

func SendCode(ctx context.Context, c *app.RequestContext) {
	var req user.SendCodeReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	_, err = rpc.SendCode(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}

func ResetPassword(ctx context.Context, c *app.RequestContext) {
	var req user.ResetPwdReq
	err := c.BindAndValidate(&req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	_, err = rpc.ResetPwd(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResSuccess(c)
}

func LoadAvatar(ctx context.Context, c *app.RequestContext) {
	var req user.UploadAvatarReq
	file, err := c.FormFile("file")
	if err != nil {
		base.ResErr(c, merror.NewMerror(merror.InvalidFile, "upload file failed"))
		return
	}
	ok := utils.CheckImageFileType(file)
	if !ok {
		base.ResErr(c, merror.NewMerror(merror.InvalidImageFileType, "file type is not image"))
		return
	}
	avatarData, err := utils.FileToBytes(file)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	req.AvatarData = avatarData
	req.AvatarName = file.Filename
	resp, err := rpc.UploadAvatar(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	resp, err = rpc.UploadAvatar(ctx, &req)
	if err != nil {
		base.ResErr(c, err)
		return
	}
	base.ResData(c, resp)
}
