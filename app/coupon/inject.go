package coupon

import (
	"MengGoods/app/coupon/controller/api"
	"MengGoods/app/coupon/domain/service"
	"MengGoods/app/coupon/infrastructure/cache"
	"MengGoods/app/coupon/infrastructure/mq"
	"MengGoods/app/coupon/infrastructure/mysql"
	"MengGoods/app/coupon/usecase"
	"MengGoods/config"
	"MengGoods/pkg/base/client"
)

func InjectCouponServiceImpl() *api.CouponServiceImpl {
	gormDB, err := client.NewMySQLClient(config.Conf.MySQL.CouponDB)
	if err != nil {
		panic(err)
	}
	CouponDB := mysql.NewCouponDB(gormDB)
	cacheClient, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}
	CouponCache := cache.NewCouponCache(cacheClient)
	kafkaClient := client.NewKafka()
	CouponMq := mq.NewCouponMq(kafkaClient)
	CouponService := service.NewCouponService(CouponDB, CouponCache, CouponMq)
	CouponUsecase := usecase.NewCouponUsecase(CouponService, CouponCache, CouponDB, CouponMq)
	CouponServiceImpl := api.NewCouponServiceImpl(CouponUsecase)
	return CouponServiceImpl
}
