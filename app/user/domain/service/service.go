package service

import (
	"MengGoods/app/user/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"fmt"
	"time"
)

func (s *UserService) EncryptPassword(password string) (string, error) {
	return utils.EncryptPassword(password)
}

func (s *UserService) ComparePassword(hashedPassword, password string) error {
	return utils.ComparePassword(hashedPassword, password)
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	uid, err := s.db.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	user.Uid = uid
	return uid, nil
}

func (s *UserService) GetUserInfo(ctx context.Context, uid int64) (*model.User, error) {
	user, err := s.db.GetUserByID(ctx, uid)
	if err != nil {
		return nil, merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("failed to get user info, err:%v", err),
		)
	}
	return user, nil
}

func (s *UserService) GetAddress(ctx context.Context, uid int64) ([]*model.Address, error) {
	addresses, err := s.db.GetAddress(ctx, uid)
	if err != nil {
		return nil, merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("failed to get address, err:%v", err),
		)
	}
	return addresses, nil
}

func (s *UserService) GetAddressByID(ctx context.Context, addressId int64) (*model.Address, error) {
	address, err := s.db.GetAddressByID(ctx, addressId)
	if err != nil {
		return nil, merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("failed to get address, err:%v", err),
		)
	}
	return address, nil
}

func (s *UserService) AddAddress(ctx context.Context, address *model.Address) (int64, error) {
	addrID, err := s.db.AddAddress(ctx, address)
	if err != nil {
		return 0, merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("failed to add address, err:%v", err),
		)
	}
	return addrID, nil
}

func (s *UserService) IsBanned(ctx context.Context, uid int64) (bool, error) {
	sKey := s.cache.GetBanKey(ctx, uid)
	isBan, err := s.cache.IsBanned(ctx, sKey)
	if err != nil {
		return false, merror.NewMerror(
			merror.InternalCacheErrorCode,
			fmt.Sprintf("failed to check if user is banned, err:%v", err),
		)
	}
	return isBan, nil
}

func (s *UserService) UserLogin(ctx context.Context, uid int64) error {
	key := s.cache.GetInKey(ctx, uid)
	exist, err := s.cache.IsExist(ctx, key) //查询用户会话是否存在
	if err != nil {
		return merror.NewMerror(
			merror.InternalCacheErrorCode,
			fmt.Sprintf("failed to check if user login, err:%v", err),
		)
	}
	var token string
	if exist {
		oldToken, err := s.cache.GetToken(ctx, key)
		if err != nil {
			return merror.NewMerror(
				merror.InternalCacheErrorCode,
				fmt.Sprintf("failed to get token, err1:%v", err),
			)
		}
		//检查token是否过期
		claims, err := utils.CheckToken(oldToken)
		if err != nil {
			return merror.NewMerror(
				merror.TokenExpired,
				fmt.Sprintf("failed to check token, err2:%v", err),
			)
		}
		//检查token是否过期
		if claims.ExpiresAt < time.Now().Unix() {
			return merror.NewMerror(
				merror.TokenExpired,
				"token过期",
			)
		}
		return nil
	} else {
		token, err = utils.CreateToken(constants.TypeLogin, uid)
		if err != nil {
			return merror.NewMerror(
				merror.InternalCacheErrorCode,
				fmt.Sprintf("failed to create token, err3:%v", err),
			)
		}
	}
	if err := s.cache.SetLogin(ctx, key, token); err != nil {
		return merror.NewMerror(
			merror.InternalCacheErrorCode,
			fmt.Sprintf("failed to set login, err4:%v", err),
		)
	}
	return nil
}

func (s *UserService) BanUser(ctx context.Context, uid int64) error {
	me, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return merror.NewMerror(
			merror.ParamFromContextFailed,
			fmt.Sprintf("failed to get user info, err:%v", err),
		)
	}
	MyInfo, err := s.db.GetUserByID(ctx, me)
	if err != nil {
		return merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("failed to get user info, err:%v", err),
		)
	}
	//检查是否为管理员
	if MyInfo.Role != constants.Admin {
		return merror.NewMerror(
			merror.AuthNoOperatePermissionCode,
			"only admin can ban user",
		)
	}
	//不能删除自己
	if uid == me {
		return merror.NewMerror(
			merror.AuthNoOperatePermissionCode,
			"can not ban self",
		)
	}
	//检查用户是否存在
	user, err := s.db.GetUserByID(ctx, uid)
	if err != nil {
		return merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("failed to get user info, err:%v", err),
		)
	}
	//检查用户是否被封禁
	isBan, err := s.IsBanned(ctx, uid)
	if err != nil {
		return merror.NewMerror(
			merror.InternalCacheErrorCode,
			fmt.Sprintf("failed to check if user is banned, err:%v", err),
		)
	}
	if isBan {
		return merror.NewMerror(
			merror.UserIsBanned,
			fmt.Sprintf("user %d is banned", uid),
		)
	}
	//判断用户是否为管理员
	if user.Role == constants.Admin {
		return merror.NewMerror(
			merror.AuthNoOperatePermissionCode,
			"can not ban admin",
		)
	}
	//封禁用户
	key := s.cache.GetBanKey(ctx, uid)
	if err := s.cache.SetUserBan(ctx, key); err != nil {
		return merror.NewMerror(
			merror.InternalCacheErrorCode,
			fmt.Sprintf("failed to set ban, err:%v", err),
		)
	}
	return nil
}

func (s *UserService) UnBanUser(ctx context.Context, uid int64) error {
	me, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return merror.NewMerror(
			merror.ParamFromContextFailed,
			fmt.Sprintf("failed to get user info, err:%v", err),
		)
	}
	MyInfo, err := s.db.GetUserByID(ctx, me)
	if err != nil {
		return merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("failed to get user info, err:%v", err),
		)
	}
	//检查是否为管理员
	if MyInfo.Role != constants.Admin {
		return merror.NewMerror(
			merror.AuthNoOperatePermissionCode,
			"only admin can unban user",
		)
	}
	//检查用户是否存在
	user, err := s.db.GetUserByID(ctx, uid)
	if err != nil {
		return merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("failed to get user info, err:%v", err),
		)
	}
	if user.Role == constants.Admin {
		return merror.NewMerror(
			merror.AuthNoOperatePermissionCode,
			"can not unban admin",
		)
	}
	//检查用户是否被封禁
	isBan, err := s.IsBanned(ctx, uid)
	if err != nil {
		return merror.NewMerror(
			merror.InternalCacheErrorCode,
			fmt.Sprintf("failed to check if user is banned, err:%v", err),
		)
	}
	if !isBan {
		return merror.NewMerror(
			merror.UserNotBanned,
			fmt.Sprintf("user %d not banned", uid),
		)
	}
	//解封用户
	key := s.cache.GetBanKey(ctx, uid)
	if err := s.cache.DeleteUserBan(ctx, key); err != nil {
		return merror.NewMerror(
			merror.InternalCacheErrorCode,
			fmt.Sprintf("failed to delete ban, err:%v", err),
		)
	}
	return nil
}

func (s *UserService) UserLogOut(ctx context.Context, uid int64) error {
	key := s.cache.GetInKey(ctx, uid)
	if err := s.cache.DeleteLogIn(ctx, key); err != nil {
		return merror.NewMerror(
			merror.InternalCacheErrorCode,
			fmt.Sprintf("failed to delete login, err:%v", err),
		)
	}
	return nil
}
