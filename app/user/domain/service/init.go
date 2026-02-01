package service

import (
	"MengGoods/app/user/domain/repository"
)

type UserService struct {
	db    repository.UserDB
	cache repository.UserCache
	cos   repository.UserCos
}

func NewUserService(db repository.UserDB, cache repository.UserCache, cos repository.UserCos) *UserService {
	if db == nil || cache == nil || cos == nil {
		panic("db or cache or cos is nil")
	}
	return &UserService{
		db:    db,
		cache: cache,
		cos:   cos,
	}
}
