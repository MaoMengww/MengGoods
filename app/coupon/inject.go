package coupon

import (
	"MengGoods/app/coupon/controller/api"
	"MengGoods/app/coupon/domain/service"
	"MengGoods/app/coupon/infrastructrue/cache"
	"MengGoods/app/coupon/infrastructrue/mq"
	"MengGoods/app/coupon/infrastructrue/mysql"
	"MengGoods/app/coupon/usecase"
	"MengGoods/pkg/base/client"

	"github.com/spf13/viper"
)

func InjectCouponServiceImpl() *api.CouponServiceImpl {
    gormDB, err := client.NewMySQLClient(viper.GetString("mysql.coupondb"))
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