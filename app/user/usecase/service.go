package usecase

import (
	"MengGoods/app/user/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("MengGoods/app/user/usecase")

// 注册用户
func (u *userUsecase) Register(ctx context.Context, user *model.User) (int64, error) {
	//校验输入合法性
	if err := utils.Verify(utils.VerifyEmail(user.Email), utils.VerifyPassword(user.Password), utils.VerifyUsername(user.Username)); err != nil {
		return 0, err
	}
	exist, err := u.db.IsUserExist(ctx, user.Username)
	//db错误
	if err != nil {
		return 0, merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("check user exist failed: %v", err),
		)
	}
	//用户已存在
	if exist {
		return 0, merror.NewMerror(merror.UserAlreadyExist, "用户已存在")
	}
	//加密密码
	user.Password, err = utils.EncryptPassword(user.Password)
	if err != nil {
		return 0, merror.NewMerror(
			merror.InternalCacheErrorCode,
			fmt.Sprintf("encrypt password failed: %v", err),
		)
	}
	//创建用户
	uid, err := u.service.CreateUser(ctx, user)
	if err != nil {
		return 0, merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("create user failed: %v", err),
		)
	}
	return uid, nil
}

// 登录用户
func (u *userUsecase) Login(ctx context.Context, user *model.User) (*model.User, error) {
	if err := utils.Verify(utils.VerifyUsername(user.Username), utils.VerifyPassword(user.Password)); err != nil {
		return nil, err
	}
	Dbuser, err := u.db.GetUserByName(ctx, user.Username)
	if err != nil {
		return nil, merror.NewMerror(
			merror.UserNotExist,
			"用户不存在",
		)
	}
	exist, err := u.service.IsBanned(ctx, Dbuser.Uid)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, merror.NewMerror(
			merror.UserIsBanned,
			"用户已被封禁",
		)
	}
	err = u.service.ComparePassword(Dbuser.Password, user.Password)
	if err != nil {
		return nil, merror.NewMerror(
			merror.PasswordNotMatch,
			"密码错误",
		)
	}
	err = u.service.UserLogin(ctx, Dbuser.Uid)
	if err != nil {
		return nil, err
	}
	return Dbuser, nil
}

// 封禁用户
func (u *userUsecase) BanUser(ctx context.Context, uid int64) error {
	if err := u.service.BanUser(ctx, uid); err != nil {
		return err
	}
	return nil
}

// 解封用户
func (u *userUsecase) UnBanUser(ctx context.Context, uid int64) error {
	if err := u.service.UnBanUser(ctx, uid); err != nil {
		return err
	}
	return nil
}

// 添加地址
func (u *userUsecase) AddAddress(ctx context.Context, address *model.Address) (int64, error) {
	addressId, err := u.service.AddAddress(ctx, address)
	if err != nil {
		return 0, err
	}
	return addressId, nil
}

// 获取用户地址列表
func (u *userUsecase) GetAddressList(ctx context.Context) ([]*model.Address, error) {
	uid, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	addressList, err := u.service.GetAddress(ctx, uid)
	if err != nil {
		return nil, err
	}
	return addressList, nil
}

func (u *userUsecase) GetAddress(ctx context.Context, addressId int64) (*model.Address, error) {
	address, err := u.service.GetAddressByID(ctx, addressId)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (u *userUsecase) SetAdmin(ctx context.Context, password string, uid int64) error {
	if err := utils.Verify(utils.VerifyPassword(password)); err != nil {
		return err
	}
	me, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	MyInfo, err := u.db.GetUserByID(ctx, me)
	if err != nil {
		return merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("get user info failed: %v", err),
		)
	}
	if MyInfo.Role != constants.Admin {
		return merror.NewMerror(
			merror.PermissionDenied,
			"权限不足",
		)
	}

	if password != viper.GetString("TopSecret") {
		return merror.NewMerror(
			merror.PasswordNotMatch,
			"密码错误",
		)
	}
	if err := u.db.SetUserAdmin(ctx, uid); err != nil {
		return merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("set user admin failed: %v", err),
		)
	}
	return nil
}

func (u *userUsecase) GetUserInfo(ctx context.Context, uid int64) (*model.User, error) {
	UserInfo, err := u.db.GetUserByID(ctx, uid)
	if err != nil {
		return nil, merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("get user info failed: %v", err),
		)
	}
	return UserInfo, nil
}

func (u *userUsecase) LogOut(ctx context.Context) error {
	uid, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	if err := u.service.UserLogOut(ctx, uid); err != nil {
		return err
	}
	return nil
}
