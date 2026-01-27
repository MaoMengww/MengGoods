package service

import (
	"MengGoods/app/user/domain/repository"
)

type UserService struct {
	db    repository.UserDB
	cache repository.UserCache
}

func NewUserService(db repository.UserDB, cache repository.UserCache) *UserService {
	if db == nil || cache == nil {
		panic("db or cache is nil")
	}
	return &UserService{
		db:    db,
		cache: cache,
	}
}
