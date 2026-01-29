package usecase

import (
	"MengGoods/app/payment/domain/repository"
	"MengGoods/app/payment/domain/service"
)

type PaymentUsecase struct {
	PaymentService *service.PaymentService
	PaymentDB      repository.PaymentDB
	PaymentCache   repository.PaymentCache
	PaymentRpc     repository.PaymentRpc
}

func NewPaymentUsecase(svc *service.PaymentService, db repository.PaymentDB, cache repository.PaymentCache, rpc repository.PaymentRpc) *PaymentUsecase {
	return &PaymentUsecase{
		PaymentService: svc,
		PaymentDB:      db,
		PaymentCache:   cache,
		PaymentRpc:     rpc,
	}
}
