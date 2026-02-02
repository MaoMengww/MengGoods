package stock

import (
	"MengGoods/app/stock/controller/api"
	"MengGoods/app/stock/domain/service"
	"MengGoods/app/stock/infrastructure/cache"
	"MengGoods/app/stock/infrastructure/mq"
	"MengGoods/app/stock/infrastructure/mysql"
	"MengGoods/app/stock/usecase"
	"MengGoods/config"
	"MengGoods/kitex_gen/stock"
	"MengGoods/pkg/base/client"
)

func InjectStockUsecaseImpl() stock.StockService {
	gormDB, err := client.NewMySQLClient(config.Conf.MySQL.StockDB)
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
	stockUsecase := usecase.NewStockUsecase(stockDB, stockCache, stockMq, stockService)
	stockServiceImpl := api.NewStockServiceImpl(stockUsecase)
	return stockServiceImpl
}
