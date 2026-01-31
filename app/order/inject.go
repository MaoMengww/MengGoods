package order

import (
	"MengGoods/app/order/controller/api"
	"MengGoods/app/order/domain/service"
	"MengGoods/app/order/infrastructure/mq"
	"MengGoods/app/order/infrastructure/mysql"
	"MengGoods/app/order/infrastructure/prpc"
	"MengGoods/app/order/usecase"
	"MengGoods/config"
	"MengGoods/pkg/base/client"
)

func InjectOrderServiceImpl() *api.OrderServiceImpl {
	gormDB, err := client.NewMySQLClient(config.Conf.MySQL.OrderDB)
	if err != nil {
		panic(err)
	}
	orderDB := mysql.NewOrderDB(gormDB)
	rabbitMq := client.NewRabbitMq(
		config.Conf.RabbitMQ.Exchange,
		config.Conf.RabbitMQ.DelayQueue,
		config.Conf.RabbitMQ.ProcessQueue,
		config.Conf.RabbitMQ.RoutingKey,
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
