package cache

import (
	"MengGoods/pkg/merror"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	redis *redis.Client
}

func NewUserCache(redis *redis.Client) *UserCache {
	return &UserCache{
		redis: redis,
	}
}

func (u *UserCache) IsBanned(ctx context.Context, key string) (bool, error) {
	exists, err := u.redis.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return exists != "", nil
}

func (u *UserCache) SetUserBan(ctx context.Context, key string) error {
	err := u.redis.Set(ctx, key, true, 0).Err()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("set user ban failed,err:%v", err))
	}
	return nil
}

func (u *UserCache) DeleteUserBan(ctx context.Context, key string) error {
	err := u.redis.Del(ctx, key).Err()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("delete user ban failed,err:%v", err))
	}
	return nil
}

func (u *UserCache) SetLogin(ctx context.Context, key string, token string) error {
	err := u.redis.Set(ctx, key, token, 0).Err()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("set user login failed,err:%v", err))
	}
	return nil
}

func (u *UserCache) SetCode(ctx context.Context, key string, code string) error {
	err := u.redis.Set(ctx, key, code, 0).Err()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("set user code failed,err:%v", err))
	}
	return nil
}


func (u *UserCache) DeleteLogIn(ctx context.Context, key string) error {
	err := u.redis.Del(ctx, key).Err()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("delete user login failed,err:%v", err))
	}
	return nil
}

func (u *UserCache) GetToken(ctx context.Context, key string) (string, error) {
	token, err := u.redis.Get(ctx, key).Result()
	if err != nil {
		return "", merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("get user token failed,err:%v", err))
	}
	return token, nil
}

func (u *UserCache) GetCode(ctx context.Context, key string) (string, error) {
	code, err := u.redis.Get(ctx, key).Result()
	if err != nil {
		return "", merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("get user code failed,err:%v", err))
	}
	return code, nil
}


func (u *UserCache) IsExist(ctx context.Context, key string) (bool, error) {
	exists, err := u.redis.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists >= 1, nil
}

func (u *UserCache) GetBanKey(ctx context.Context, uid int64) string {
	return fmt.Sprintf("MengGoods:User:Ban:%v", uid)
}

func (u *UserCache) GetInKey(ctx context.Context, uid int64) string {
	return fmt.Sprintf("MengGoods:User:In:%v", uid)
}

func (u *UserCache) GetCodeKey(ctx context.Context, uid int64) string {
	return fmt.Sprintf("MengGoods:User:Code:%v", uid)
}
