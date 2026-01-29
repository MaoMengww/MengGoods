package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type paymentRedis struct {
	client *redis.Client
}

func NewPaymentRedis(client *redis.Client) *paymentRedis {
	cli := &paymentRedis{client: client}
	if err := cli.loadScript(); err != nil {
		panic(err)
	}
	return cli
}

func (p *paymentRedis) GetTTLAndDelPaymentToken(ctx context.Context, key string, token string) (exists bool, ttl int64, err error) {
	values, err := p.client.Eval(ctx, scripts[GetTTLAndDelPaymentTokenScriptKey].Hash, []string{key}, token).Result()
	if err != nil {
		return false, 0, err
	}
	ttl, exists = values.([]int64)[0], values.([]int64)[1] == 1
	return exists, ttl, nil
}

func (p *paymentRedis) SetOrIncrRefundKey(ctx context.Context, key string) (ttl int64, err error) {
	values, err := p.client.Eval(ctx, scripts[SetOrIncrRefundKeyScriptKey].Hash, []string{key}).Result()
	if err != nil {
		return 0, err
	}
	ttl = values.([]int64)[1]
	return ttl, nil
}

func (p *paymentRedis) GetPaymentKey(ctx context.Context, orderId int64) string {
	return fmt.Sprintf("MengGoods:Payment:Order:%d", orderId)
}

func (p *paymentRedis) GetRefundKey(ctx context.Context, orderItemId int64) string {
	return fmt.Sprintf("MengGoods:Payment:Refund:%d", orderItemId)
}

func (p *paymentRedis) GetExpiredDuration(ctx context.Context, expiredTime int64) (int64, error) {
	duration := expiredTime - time.Now().Unix()
	if duration < 0 {
		return 0, fmt.Errorf("expired time is invalid")
	}
	return duration, nil
}

func (p *paymentRedis) SetPaymentToken(ctx context.Context, key string, token string, expire int64) error {
	_, err := p.client.Set(ctx, key, token, time.Duration(expire)*time.Second).Result()
	return err
}

func (p *paymentRedis) SetDailyRefund(ctx context.Context, key string) error {
	_, err := p.client.Set(ctx, key, 1, 24*time.Hour).Result()
	return err
}

func (p *paymentRedis) CheckDailyRefundCount(ctx context.Context, key string) (bool, error) {
	exists, err := p.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}