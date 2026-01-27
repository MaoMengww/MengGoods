package usecase

import (
	"MengGoods/app/coupon/domain/repository"
	"MengGoods/app/coupon/domain/service"
)

type CouponUsecase struct {
	service     *service.CouponService
	couponCache repository.CouponCache
	couponDB    repository.CouponDB
	couponMq    repository.CouponMq
}

func NewCouponUsecase(service *service.CouponService, couponCache repository.CouponCache, couponDB repository.CouponDB, couponMq repository.CouponMq) *CouponUsecase {
	return &CouponUsecase{
		service:     service,
		couponCache: couponCache,
		couponDB:    couponDB,
		couponMq:    couponMq,
	}
}
