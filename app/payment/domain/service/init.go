package service

import "MengGoods/app/payment/domain/repository"

type PaymentService struct {
	PaymentDB repository.PaymentDB
	PaymentCache repository.PaymentCache
	PaymentRpc repository.PaymentRpc
}

func NewPaymentService(db repository.PaymentDB, cache repository.PaymentCache, rpc repository.PaymentRpc) *PaymentService {
	return &PaymentService{
		PaymentDB:    db,
		PaymentCache: cache,
		PaymentRpc:   rpc,
	}
}
