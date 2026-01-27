package order

import (
	"MengGoods/app/order/controller/api"
	"MengGoods/app/order/domain/service"
	"MengGoods/app/order/infrastructrue/mq"
	"MengGoods/app/order/infrastructrue/mysql"
	"MengGoods/app/order/infrastructrue/prpc"
	"MengGoods/app/order/usecase"
	"MengGoods/pkg/base/client"

	"github.com/spf13/viper"
)

func InjectOrderServiceImpl() *api.OrderServiceImpl {
	gormDB, err := client.NewMySQLClient(viper.GetString("mysql.orderdb"))
	if err != nil {
		panic(err)
	}
	orderDB := mysql.NewOrderDB(gormDB)
	rabbitMq := client.NewRabbitMq(
		viper.GetString("rabbitmq.exchange"),
		viper.GetString("rabbitmq.delayqueue"),
		viper.GetString("rabbitmq.processqueue"),
		viper.GetString("rabbitmq.routingkey"),
	)
	orderMq := mq.NewRabbitMQ(rabbitMq)
	userClient := prpc.NewUserClient()
	cartClient := prpc.NewCartClient()
	couponClient := prpc.NewCouponClient()
	productClient := prpc.NewProductClient()
	orderRpc := prpc.NewOrderRpc(cartClient, couponClient, productClient, userClient)
	orderService := service.NewOrderService(orderDB, orderMq, orderRpc)
	orderUsecase := usecase.NewUsecase(*orderService, orderDB, orderMq, orderRpc)
	orderServiceImpl := api.NewOrderServiceImpl(orderUsecase)
	return orderServiceImpl
}
