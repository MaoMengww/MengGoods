package cache

import (
	"MengGoods/app/cart/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type CartCache struct {
	*redis.Client
}

func NewCartCache(redisClient *redis.Client) *CartCache {
	return &CartCache{
		Client: redisClient,
	}
}

func (p *CartCache) GetCartKey(ctx context.Context) (string, error) {
	userID, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil || userID == 0 {
		return "", merror.NewMerror(merror.ParamFromContextFailed, "failed to get user ID")
	}
	key := "MengGoods:cart:" + strconv.FormatInt(userID, 10)
	return key, nil
}

func (p *CartCache) AddCartItem(ctx context.Context, cartItem *model.CartItem) error {
	key, err := p.GetCartKey(ctx)
	if err != nil {
		return err
	}

	field := strconv.FormatInt(cartItem.SkuID, 10)
	val, err := p.Client.HGet(ctx, key, field).Result()

	// 最终要存入的对象
	var finalItem *model.CartItem

	if err == redis.Nil {
		finalItem = cartItem
	} else if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("redis error: %v", err))
	} else {
		var oldCartItem model.CartItem
		if err := json.Unmarshal([]byte(val), &oldCartItem); err != nil {
			finalItem = cartItem
		} else {
			oldCartItem.Count += cartItem.Count
			finalItem = &oldCartItem
		}
	}
	data, err := json.Marshal(finalItem)
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, "marshal cart item failed")
	}
	pipe := p.Client.Pipeline()
	pipe.HSet(ctx, key, field, data)
	pipe.Expire(ctx, key, time.Duration(constants.CartExpireTime)*time.Second)

	if _, err = pipe.Exec(ctx); err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("execute pipeline failed: %v", err))
	}

	return nil
}

func (p *CartCache) GetCart(ctx context.Context) ([]*model.CartItem, error) {
	key, err := p.GetCartKey(ctx)
	if err != nil {
		return nil, err
	}
	val, err := p.Client.HGetAll(ctx, key).Result()
	if err == redis.Nil {
		return nil, merror.NewMerror(merror.CartIsEmptyCode, "cart is empty")
	} else if err != nil {
		return nil, merror.NewMerror(merror.InternalCacheErrorCode, "failed to get cart")
	} else {
		var cartItems []*model.CartItem
		for _, v := range val {
			var cartItem model.CartItem
			err := json.Unmarshal([]byte(v), &cartItem)
			if err != nil {
				return nil, merror.NewMerror(merror.InternalCacheErrorCode, "failed to get cart")
			}
			cartItems = append(cartItems, &cartItem)
		}
		return cartItems, nil
	}
}

func (p *CartCache) DeleteCart(ctx context.Context) error {
	key, err := p.GetCartKey(ctx)
	if err != nil {
		return err
	}
	err = p.Client.Del(ctx, key).Err()
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, "failed to delete cart")
	}
	return nil
}

func (p *CartCache) UpdateCartItem(ctx context.Context, cartItem *model.CartItem) error {
    key, err := p.GetCartKey(ctx)
    if err != nil {
        return err
    }
    data, err := json.Marshal(cartItem)
    if err != nil {
        return merror.NewMerror(merror.InternalCacheErrorCode, "marshal cart item failed")
    }
    pipe := p.Client.Pipeline()
    pipe.HSet(ctx, key, strconv.FormatInt(cartItem.SkuID, 10), data) 
    pipe.Expire(ctx, key, time.Duration(constants.CartExpireTime)*time.Second) 
    _, err = pipe.Exec(ctx)
    if err != nil {
        return merror.NewMerror(merror.InternalCacheErrorCode, fmt.Sprintf("update cart redis failed: %v", err))
    }
    return nil
}
