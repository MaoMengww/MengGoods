package service

import (
	"MengGoods/app/order/domain/repository"
	"MengGoods/pkg/logger"
	"context"
)

type OrderService struct {
	OrderDB  repository.OrderDB
	OrderMq  repository.OrderMq
	OrderRpc repository.OrderRpc
}

func NewOrderService(orderDB repository.OrderDB, orderMq repository.OrderMq, orderRpc repository.OrderRpc) *OrderService {
	service := &OrderService{
		OrderDB:  orderDB,
		OrderMq:  orderMq,
		OrderRpc: orderRpc,
	}
	service.Init()
	return service
}

func (s *OrderService) Init() {
	if err := s.ConsumeOrderMessage(context.Background()); err != nil {
		logger.CtxFatalf(context.Background(), "Consume order message error: %v", err)
	}
	go s.SendMsg(context.Background())
}
