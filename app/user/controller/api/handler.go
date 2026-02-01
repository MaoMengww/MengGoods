package api

import (
	"MengGoods/app/user/controller/api/pack"
	"MengGoods/app/user/domain/model"
	"MengGoods/app/user/usecase"
	user "MengGoods/kitex_gen/user"
	"MengGoods/pkg/base"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	usecase usecase.UserUsecase
}

func NewUserServiceImpl(usecase usecase.UserUsecase) *UserServiceImpl {
	return &UserServiceImpl{usecase: usecase}
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	r := new(user.RegisterResp)
	var uid int64
	uid, err = s.usecase.Register(ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	r.UserId = uid
	resp = r
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	r := new(user.LoginResp)
	var user *model.User
	user, err = s.usecase.Login(ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	r.UserInfo = pack.BuildUserInfo(user)
	resp = r
	return
}

// AddAddress implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddAddress(ctx context.Context, req *user.AddAddressReq) (resp *user.AddAddressResp, err error) {
	r := new(user.AddAddressResp)
	var addressId int64
	addressId, err = s.usecase.AddAddress(ctx, &model.Address{
		Province: req.Province,
		City:     req.City,
		Detail:   req.Detail,
	})
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	r.AddressId = addressId
	resp = r
	return
}

// GetAddress implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetAddress(ctx context.Context, req *user.GetAddressReq) (resp *user.GetAddressResp, err error) {
	r := new(user.GetAddressResp)
	address, err := s.usecase.GetAddress(ctx, req.AddressId)
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	r.Address = pack.BuildAddress(address)
	resp = r
	return r, nil
}

func (s *UserServiceImpl) GetAddresses(ctx context.Context, req *user.GetAddressesReq) (resp *user.GetAddressesResp, err error) {
	r := new(user.GetAddressesResp)
	addressList, err := s.usecase.GetAddressList(ctx)
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	r.Address = pack.BuildAddressList(addressList)
	resp = r
	return r, nil
}

// BanUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) BanUser(ctx context.Context, req *user.BanUserReq) (resp *user.BanUserResp, err error) {
	r := new(user.BanUserResp)
	if err = s.usecase.BanUser(ctx, req.UserId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	resp = r
	return r, nil
}

// UnBanUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UnBanUser(ctx context.Context, req *user.UnBanUserReq) (resp *user.UnBanUserResp, err error) {
	r := new(user.UnBanUserResp)
	if err = s.usecase.UnBanUser(ctx, req.UserId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	resp = r
	return r, nil
}

// SetAdmin implements the UserServiceImpl interface.
func (s *UserServiceImpl) SetAdmin(ctx context.Context, req *user.SetAdminReq) (resp *user.SetAdminResp, err error) {
	r := new(user.SetAdminResp)
	if err = s.usecase.SetAdmin(ctx, req.Password, req.UserId); err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	resp = r
	return r, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	r := new(user.GetUserInfoResp)
	var userInfo *model.User
	userInfo, err = s.usecase.GetUserInfo(ctx, req.UserId)
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	r.UserInfo = pack.BuildUserInfo(userInfo)
	resp = r
	return r, nil
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	r := new(user.LogoutResp)
	if err = s.usecase.LogOut(ctx); err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	resp = r
	return r, nil
}

// SendCode implements the UserServiceImpl interface.
func (s *UserServiceImpl) SendCode(ctx context.Context, req *user.SendCodeReq) (resp *user.SendCodeResp, err error) {
	r := new(user.SendCodeResp)
	if err = s.usecase.SendCode(ctx, req.Email); err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	resp = r
	return r, nil
}

// ResetPwd implements the UserServiceImpl interface.
func (s *UserServiceImpl) ResetPwd(ctx context.Context, req *user.ResetPwdReq) (resp *user.ResetPwdResp, err error) {
	// TODO: Your code here...
	r := new(user.ResetPwdResp)
	if err = s.usecase.UpdatePassword(ctx, req.Code, req.Password); err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.Base = base.BuildBaseResp(nil)
	resp = r
	return r, nil
}

func (s *UserServiceImpl) UploadAvatar(ctx context.Context, req *user.UploadAvatarReq) (resp *user.UploadAvatarResp, err error) {
	// TODO: Your code here...
	r := new(user.UploadAvatarResp)
	avatarURL, err := s.usecase.UploadAvatar(ctx, req.AvatarData, req.AvatarName)
	if err != nil {
		r.Base = base.BuildBaseResp(err)
		return r, nil
	}
	r.AvatarURL = avatarURL
	r.Base = base.BuildBaseResp(nil)
	resp = r
	return r, nil
}

