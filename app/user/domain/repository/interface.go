package repository

import (
	"MengGoods/app/user/domain/model"
	"context"
)

type UserDB interface {
	IsUserExist(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	GetUserByID(ctx context.Context, uid int64) (*model.User, error)
	GetUserByName(ctx context.Context, username string) (*model.User, error)
	GetAddress(ctx context.Context, uid int64) ([]*model.Address, error)
	AddAddress(ctx context.Context, addr *model.Address) (int64, error)
	GetAddressByID(ctx context.Context, addressId int64) (*model.Address, error)
	SetUserAdmin(ctx context.Context, uid int64) error
}

// 缓存
type UserCache interface {
	IsBanned(ctx context.Context, key string) (bool, error)
	SetUserBan(ctx context.Context, key string) error             //设置用户被ban
	DeleteUserBan(ctx context.Context, key string) error          //删除用户被ban
	SetLogin(ctx context.Context, key string, token string) error //设置用户登录
	DeleteLogIn(ctx context.Context, key string) error            //删除用户登录
	GetBanKey(ctx context.Context, uid int64) string              //获取用户被ban的key
	GetInKey(ctx context.Context, uid int64) string               //获取用户会话存在的key
	GetToken(ctx context.Context, key string) (string, error)     //获取key对应的token
	IsExist(ctx context.Context, key string) (bool, error)        //判断key是否存在
}
