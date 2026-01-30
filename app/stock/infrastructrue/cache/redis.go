package cache

import (
	"MengGoods/pkg/merror"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// lua脚本, 防止超卖
const LuaReduceStock = `
	local failed_status = -1
	for i = 1, #KEYS do
		local stock = tonumber(redis.call('get', KEYS[i]))
		local count = tonumber(ARGV[i])
		if stock < count then
			failed_status = i
			break
		end
	end
	if failed_status ~= -1 then
		return -1
	end
	for i = 1, #KEYS do
		redis.call('decrby', KEYS[i], ARGV[i])
	end
	return 1
`

type StockCache struct {
	*redis.Client
}

func NewStockCache(redisClient *redis.Client) *StockCache {
	return &StockCache{
		Client: redisClient,
	}
}

func (p *StockCache) GetStockKey(ctx context.Context, skuId int64) string {
	key := fmt.Sprintf("MengGoods:Stock:Sku:%d", skuId)
	return key
}

func (p *StockCache) SetStock(ctx context.Context, key string, count int32) error {
	if err := p.Client.Set(ctx, key, count, 0).Err(); err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("set stock cache error: %v", err))
	}
	return nil
}

func (p *StockCache) GetStock(ctx context.Context, key string) (int32, error) {
	value, err := p.Client.Get(ctx, key).Int()
	if err != nil {
		return 0, merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("get stock cache error: %v", err))
	}
	return int32(value), nil
}

func (p *StockCache) AddStock(ctx context.Context, key string, count int32) error {
	err := p.Client.IncrBy(ctx, key, int64(count)).Err()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("add stock cache error: %v", err))
	}
	return nil
}

func (p *StockCache) RedStock(ctx context.Context, stockItems map[string]int32) error {
	//准备lua脚本的参数
	keys := make([]string, 0, len(stockItems))
	args := make([]interface{}, 0, len(stockItems))
	for key, count := range stockItems {
		keys = append(keys, key)
		args = append(args, count)
	}
	//执行lua脚本
	result, err := p.Client.EvalSha(ctx, LuaReduceStock, keys, args...).Int()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("reduce stock cache error: %v", err))
	}
	if result != 1 {
		return merror.NewMerror(merror.StockNotEnough, "库存不足")
	}
	return nil
}
