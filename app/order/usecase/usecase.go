package usecase

import (
	"MengGoods/app/order/domain/repository"
	"MengGoods/app/order/domain/service"
)

type Usecase struct {
	service  service.OrderService
	OrderDB  repository.OrderDB
	OrderMq  repository.OrderMq
	OrderRpc repository.OrderRpc
}

func NewUsecase(service service.OrderService, mysql repository.OrderDB, mq repository.OrderMq, rpc repository.OrderRpc) *Usecase {
	return &Usecase{
		service:  service,
		OrderDB:  mysql,
		OrderMq:  mq,
		OrderRpc: rpc,
	}
}
