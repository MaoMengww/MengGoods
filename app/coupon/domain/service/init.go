package service

import (
	"MengGoods/app/coupon/domain/repository"
	"context"
)

type CouponService struct {
	CouponDB    repository.CouponDB
	CouponCache repository.CouponCache
	CouponMq    repository.CouponMq
}

func NewCouponService(couponDB repository.CouponDB, couponCache repository.CouponCache, couponMq repository.CouponMq) *CouponService {
	service := &CouponService{
		CouponDB:    couponDB,
		CouponCache: couponCache,
		CouponMq:    couponMq,
	}
	service.Init()
	return service
}

func (s *CouponService) Init() {
	s.ConsumeClaimCoupon(context.Background())
}
