package prpc

import (
	"MengGoods/kitex_gen/model"
	"MengGoods/kitex_gen/user"
	"MengGoods/pkg/merror"
	"context"
)

func (c *OrderRpc) GetUserInfo(ctx context.Context, userId int64) (*model.UserInfo, error) {
	userInfo, err := c.UserClient.GetUserInfo(ctx, &user.GetUserInfoReq{
		UserId: userId,
	})
	if err != nil {
		return &model.UserInfo{}, err
	}
	if userInfo.Base.Code != merror.SuccessCode {
		return &model.UserInfo{}, merror.NewMerror(userInfo.Base.Code, userInfo.Base.Message)
	}
	return &model.UserInfo{
		Id:       userInfo.UserInfo.Id,
		Username: userInfo.UserInfo.Username,
	}, nil
}

func (c *OrderRpc) GetAddressInfo(ctx context.Context, addressId int64) (*model.AddressInfo, error) {
	addressInfo, err := c.UserClient.GetAddress(ctx, &user.GetAddressReq{
		AddressId: addressId,
	})
	if err != nil {
		return &model.AddressInfo{}, err
	}
	if addressInfo.Base.Code != merror.SuccessCode {
		return &model.AddressInfo{}, merror.NewMerror(addressInfo.Base.Code, addressInfo.Base.Message)
	}
	return &model.AddressInfo{
		Province:  addressInfo.Address.Province,
		City:      addressInfo.Address.City,
		Detail:    addressInfo.Address.Detail,
		AddressID: addressInfo.Address.AddressID,
	}, nil
}
