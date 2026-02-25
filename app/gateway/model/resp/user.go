package resp

import "MengGoods/kitex_gen/model"

type RegisterResp struct {
	UserId int64 `json:"userId"`
}

type LoginResp struct {
	User model.UserInfo `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AddAddressResp struct {
	AddressId int64 `json:"address_id"`
}

type GetAddressResp struct {
	AddressList []*model.AddressInfo `json:"address_list"`
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

type UploadAvatarResp struct {
	Url string `json:"url"`
}
