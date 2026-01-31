package payment

import (
	"MengGoods/app/payment/controller/api"
	"MengGoods/app/payment/domain/service"
	"MengGoods/app/payment/infrastructrue/cache"
	"MengGoods/app/payment/infrastructrue/mysql"
	"MengGoods/app/payment/infrastructrue/prpc"
	"MengGoods/app/payment/usecase"
	"MengGoods/config"
	"MengGoods/pkg/base/client"
)

func InjectPaymentServiceImpl() *api.PaymentServiceImpl {
	gormDB, err := client.NewMySQLClient(config.Conf.MySQL.PaymentDB)
	if err != nil {
		panic(err)
	}
	paymentDB := mysql.NewPaymentDB(gormDB)
	redisClient, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}
	paymentRedis := cache.NewPaymentRedis(redisClient)
	orderClient := prpc.NewOrderClient()
	paymentRpc := prpc.NewPaymentRpc(orderClient)
	service := service.NewPaymentService(paymentDB, paymentRedis, paymentRpc)
	usecase := usecase.NewPaymentUsecase(service, paymentDB, paymentRedis, paymentRpc)
	impl := api.NewPaymentServiceImpl(usecase)
	return impl
}
