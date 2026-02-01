package usecase

import (
	"MengGoods/app/user/domain/model"
	"MengGoods/app/user/domain/repository"
	"MengGoods/app/user/domain/service"
	"context"
)

type UserUsecase interface {
	Register(ctx context.Context, user *model.User) (int64, error)
	Login(ctx context.Context, user *model.User) (*model.User, error)
	GetUserInfo(ctx context.Context, uid int64) (*model.User, error)
	LogOut(ctx context.Context) error
	AddAddress(ctx context.Context, address *model.Address) (int64, error)
	GetAddressList(ctx context.Context) ([]*model.Address, error)
	GetAddress(ctx context.Context, uid int64) (*model.Address, error)
	BanUser(ctx context.Context, uid int64) error
	UnBanUser(ctx context.Context, uid int64) error
	SetAdmin(ctx context.Context, password string, uid int64) error
	SendCode(ctx context.Context, email string) error
	UpdatePassword(ctx context.Context, code string, password string) error
	UploadAvatar(ctx context.Context, avatarData []byte, fileName string) (string, error)
}

type userUsecase struct {
	db      repository.UserDB
	cache   repository.UserCache
	service *service.UserService
	cos     repository.UserCos
}

func NewUserUsecase(db repository.UserDB, cache repository.UserCache, svc *service.UserService, cos repository.UserCos) *userUsecase {
	if db == nil || cache == nil || svc == nil || cos == nil {
		panic("db or cache or svc or cos is nil")
	}
	return &userUsecase{
		db:      db,
		cache:   cache,
		service: svc,
		cos:     cos,
	}
}
