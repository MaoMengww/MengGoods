package prpc

import (
	"MengGoods/kitex_gen/user"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"context"
)

func (p *ProductRpc) IsAdmin(ctx context.Context) (bool, error) {
	uid, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	var req user.GetUserInfoReq
	req.UserId = uid
	resp, err := p.userClient.GetUserInfo(ctx, &req)
	if err != nil {
		return false, err
	}
	if resp.Base.Code != merror.SuccessCode {
		return false, merror.NewMerror(resp.Base.Code, resp.Base.Message)
	}
	return resp.UserInfo.Role == constants.Admin, nil
}
