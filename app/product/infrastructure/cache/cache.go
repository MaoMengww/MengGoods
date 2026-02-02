package cache

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/pkg/merror"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ProductCache struct {
	redisClient *redis.Client
}

func NewProductCache(redisClient *redis.Client) *ProductCache {
	return&ProductCache{
		redisClient: redisClient,
	}
}

func (p *ProductCache) GetSpuKey(ctx context.Context, spuId int64) string {
	key := fmt.Sprintf("MengGoods:Product:Spu:%d", spuId)
	return key
}

func (p *ProductCache) GetSkuKey(ctx context.Context, skuId int64) string {
	key := fmt.Sprintf("MengGoods:Product:Sku:%d", skuId)
	return key
}

func (p *ProductCache) SetSpu(ctx context.Context, spu *model.SpuEs) error {
	key := p.GetSpuKey(ctx, spu.Id)
	value, err := json.Marshal(spu)
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("marshal spu error: %v", err))
	}
	err = p.redisClient.Set(ctx, key, value, 2*time.Minute).Err()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("set spu cache error: %v", err))
	}
	return nil
}

func (p *ProductCache) SetSku(ctx context.Context, sku *model.SkuEs) error {
	key := p.GetSkuKey(ctx, sku.Id)
	value, err := json.Marshal(sku)
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("marshal sku error: %v", err))
	}
	err = p.redisClient.Set(ctx, key, value, 2*time.Minute).Err()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("set sku cache error: %v", err))
	}
	return nil
}

func (p *ProductCache) GetSpu(ctx context.Context, key string) (string, error) {
	res, err := p.redisClient.Get(ctx, key).Result()
	if err == nil {
		return res, nil
	} else if err == redis.Nil {
		return "", err
	} else {
		return "", merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("get spu cache error: %v", err))
	}
}

func (p *ProductCache) GetSku(ctx context.Context, key string) (string, error) {
	res, err := p.redisClient.Get(ctx, key).Result()
	if err == nil {
		return res, nil
	} else if err == redis.Nil {
		return "", err
	} else {
		return "", merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("get sku cache error: %v", err))
	}
}

func (p *ProductCache) DeleteSpu(ctx context.Context, key string) error {
	return p.redisClient.Del(ctx, key).Err()
}

func (p *ProductCache) DeleteSku(ctx context.Context, key string) error {
	return p.redisClient.Del(ctx, key).Err()
}
