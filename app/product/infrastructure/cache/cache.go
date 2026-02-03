package cache

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/pkg/constants"
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
	ctx := context.Background()
	//初始化布隆过滤器
	redisClient.Do(ctx, "BF.RESERVE", constants.BloomFilterSpuKey, 1000000, 0.01)
	redisClient.Do(ctx, "BF.RESERVE", constants.BloomFilterSkuKey, 1000000, 0.01)
	return &ProductCache{
		redisClient: redisClient,
	}
}

func (p *ProductCache) LoadBloomFilter(ctx context.Context, spuIds []int64, skuIds []int64) error {
	if len(spuIds) == 0 && len(skuIds) == 0 {
		return nil
	}
	//加载spuIds到布隆过滤器
	pipr := p.redisClient.Pipeline()
	for _, spuId := range spuIds {
		pipr.BFAdd(ctx, constants.BloomFilterSpuKey, spuId)
	}
	for _, skuId := range skuIds {
		pipr.BFAdd(ctx, constants.BloomFilterSkuKey, skuId)
	}
	_, err := pipr.Exec(ctx)
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("load bloom filter error: %v", err))
	}
	return nil
}

func (p *ProductCache) AddSpuToBloomFilter(ctx context.Context, spuId int64) error {
	return p.redisClient.BFAdd(ctx, constants.BloomFilterSpuKey, spuId).Err()
}

func (p *ProductCache) AddSkuToBloomFilter(ctx context.Context, skuId int64) error {
	return p.redisClient.BFAdd(ctx, constants.BloomFilterSkuKey, skuId).Err()
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
	err = p.AddSpuToBloomFilter(ctx, spu.Id)
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("add spu to bloom filter error: %v", err))
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
	err = p.AddSkuToBloomFilter(ctx, sku.Id)
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("add sku to bloom filter error: %v", err))
	}
	return nil
}

func (p *ProductCache) GetSpu(ctx context.Context, spuId int64) (string, error) {
	exist, err := p.redisClient.BFExists(ctx, constants.BloomFilterSpuKey, spuId).Result()
	if err != nil {
		return "", merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("check spu in bloom filter error: %v", err))
	}
	if !exist {
		return "", merror.NewMerror(merror.RedisNotFound, "spu not found")
	}
	key := p.GetSpuKey(ctx, spuId)
	res, err := p.redisClient.Get(ctx, key).Result()
	if err == nil {
		return res, nil
	} else if err == redis.Nil {
		return "", err
	} else {
		return "", merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("get spu cache error: %v", err))
	}
}

func (p *ProductCache) GetSku(ctx context.Context, skuId int64) (string, error) {
	exist, err := p.redisClient.BFExists(ctx, constants.BloomFilterSkuKey, skuId).Result()
	if err != nil {
		return "", merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("check sku in bloom filter error: %v", err))
	}
	if !exist {
		return "", merror.NewMerror(merror.RedisNotFound, "sku not found")
	}
	key := p.GetSkuKey(ctx, skuId)
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
