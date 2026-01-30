package resp

import "MengGoods/kitex_gen/model"

type RegisterResp struct {
	UserId int64 `json:"uid"`
}

type LoginResp struct {
	User model.UserInfo `json:"user"`
}

type AddAddressResp struct {
	AddressId int64 `json:"addressid"`
}

type GetAddressResp struct {
	AddressList []*model.AddressInfo `json:"addressList"`
}

type BanUserResp struct {
}

type UnBanUserResp struct {
}

type SetAdminResp struct {
}

type GetUserInfoResp struct {
	User model.UserInfo `json:"user"`
}

type LogoutResp struct {
}

type SendCodeResp struct {
}

type ResetPwdResp struct {
}
