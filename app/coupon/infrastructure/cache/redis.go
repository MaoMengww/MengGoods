package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// lua脚本
const verifyDiscountScript = `
local couponKey = KEYS[1]

local currentNum = tonumber(redis.call('GET', couponKey) or 0)
if currentNum < 1 then
    return 0
end

redis.call('DECRBY', couponKey, 1)
return 1
`

type CouponCache struct {
	redisClient *redis.Client
}

func NewCouponCache(redisClient *redis.Client) *CouponCache {
	return &CouponCache{
		redisClient: redisClient,
	}
}

func (c *CouponCache) GetCouponBatchKey(ctx context.Context, batchId int64) string {
	return fmt.Sprintf("MengGoods:coupon:%v", batchId)
}
func (c *CouponCache) SetCoupon(ctx context.Context, key string, totalNum int64, duration time.Duration) error {
	return c.redisClient.Set(ctx, key, totalNum, duration).Err()
}

func (c *CouponCache) ClaimCoupon(ctx context.Context, key string) error {
	//使用lua脚本验证并领取优惠券
	script := redis.NewScript(verifyDiscountScript)
	result, err := script.Run(ctx, c.redisClient, []string{key}).Result()
	if err != nil {
		return err
	}
	if result.(int64) == 0 {
		return fmt.Errorf("coupon %s is not available", key)
	}
	return nil
}

func (c *CouponCache) AddCoupon(ctx context.Context, key string) error {
	return c.redisClient.SAdd(ctx, key, 1).Err()
}
