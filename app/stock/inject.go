package stock

import (
	"MengGoods/app/stock/controller/api"
	"MengGoods/app/stock/domain/service"
	"MengGoods/app/stock/infrastructrue/cache"
	"MengGoods/app/stock/infrastructrue/mq"
	"MengGoods/app/stock/infrastructrue/mysql"
	"MengGoods/app/stock/usecase"
	"MengGoods/kitex_gen/stock"
	"MengGoods/pkg/base/client"

	"github.com/spf13/viper"
)

func InjectStockUsecaseImpl() stock.StockService {
    gormDB, err := client.NewMySQLClient(viper.GetString("mysql.stockdb"))
	if err != nil {
		panic(err)
	}
	stockDB := mysql.NewStockDB(gormDB)
	cacheClient, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}
	stockCache := cache.NewStockCache(cacheClient)
	mqClient := client.NewKafka()
	stockMq := mq.NewStockMq(mqClient)
	stockService := service.NewStockService(stockDB, stockCache, stockMq)
	stockService.Init()
	stockUsecase := usecase.NewStockUsecase(stockDB, stockCache, stockMq, stockService)
	stockServiceImpl := api.NewStockServiceImpl(stockUsecase)
	return stockServiceImpl
}
