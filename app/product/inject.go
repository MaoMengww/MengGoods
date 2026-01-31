package product

import (
	"MengGoods/app/product/controller/api"
	"MengGoods/app/product/domain/service"
	"MengGoods/app/product/infrastructure/cache"
	"MengGoods/app/product/infrastructure/es"
	"MengGoods/app/product/infrastructure/mq"
	"MengGoods/app/product/infrastructure/mysql"
	"MengGoods/app/product/infrastructure/prpc"
	"MengGoods/app/product/usecase"
	"MengGoods/config"
	"MengGoods/kitex_gen/product"
	"MengGoods/pkg/base/client"
)

func InjectProductServiceImpl() product.ProductService {
	gormDB, err := client.NewMySQLClient(config.Conf.MySQL.ProductDB)
	if err != nil {
		panic(err)
	}
	productDB := mysql.NewProductDB(gormDB)
	cacheClient, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}
	productCache := cache.NewProductCache(cacheClient)
	kafkaClient := client.NewKafka()
	productMq := mq.NewProductMq(kafkaClient)
	esClient, err := client.NewEsClient()
	if err != nil {
		panic(err)
	}
	productEs := es.NewProductEs(esClient)
	userRpc := prpc.NewProductClient()
	productRpc := prpc.NewProductRpc(userRpc)
	ProductService := service.NewProductUsecase(productDB, productCache, productMq, productEs, productRpc)
	ProductService.Init()
	productUsecase := usecase.NewProductUsecase(productDB, productCache, productMq, productEs, productRpc)
	productServiceImpl := api.NewProductServiceImpl(productUsecase)
	return productServiceImpl
}
