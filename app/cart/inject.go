package cart

import (
	"MengGoods/app/cart/controller/api"
	"MengGoods/app/cart/infrastructrue/cache"
	"MengGoods/app/cart/usecase"
	"MengGoods/pkg/base/client"
)

func InjectCartUsecaseImpl() *api.CartServiceImpl {
	redisClient, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}
	cartCache := cache.NewCartCache(redisClient)
	cartUsecase := usecase.NewCartUsecase(cartCache)
	CartUsecaseImpl := api.NewCartServiceImpl(cartUsecase)
	return CartUsecaseImpl
}
