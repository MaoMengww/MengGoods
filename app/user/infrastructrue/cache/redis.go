package cache

import (
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
	return u.redis.Set(ctx, key, true, 0).Err()
}

func (u *UserCache) DeleteUserBan(ctx context.Context, key string) error {
	return u.redis.Del(ctx, key).Err()
}

func (u *UserCache) SetLogin(ctx context.Context, key string, token string) error {
	return u.redis.Set(ctx, key, token, 0).Err()
}

func (u *UserCache) DeleteLogIn(ctx context.Context, key string) error {
	return u.redis.Del(ctx, key).Err()
}

func (u *UserCache) GetToken(ctx context.Context, key string) (string, error) {
	return u.redis.Get(ctx, key).Result()
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




