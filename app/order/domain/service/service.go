package service

import (
	"MengGoods/app/order/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"time"
)

func (s *OrderService) CreateOrder(ctx context.Context, addressId int64, couponId int64, orderItems []*model.OrderItem) (int64, error) {
	userId, err := mcontext.GetUserIDFromContext(ctx)
	orderId := utils.GenerateID()
	if err != nil {
		return 0, err
	}
	userInfo, err := s.OrderRpc.GetUserInfo(ctx, userId)
	if err != nil {
		return 0, err
	}
	totalPrice := int64(0)
	if len(orderItems) == 0 {
		cartItems, err := s.OrderRpc.GetCartItems(ctx)
		if err != nil {
			return 0, err
		}
		for _, cartItem := range cartItems {
			skuInfo, err := s.OrderRpc.GetSkuInfo(ctx, cartItem.SkuId)
			if err != nil {
				return 0, err
			}
			totalPrice += int64(cartItem.Count) * skuInfo.Price
			orderItems = append(orderItems, &model.OrderItem{
				ProductId:         cartItem.SkuId,
				OrderId:           orderId,
				ProductName:       skuInfo.Name,
				ProductImage:      skuInfo.SkuImageURL,
				ProductPrice:      skuInfo.Price,
				ProductNum:        int(cartItem.Count),
				ProductTotalPrice: int64(cartItem.Count) * skuInfo.Price,
				ProductProperties: skuInfo.Properties,
			})
		}
	} else {
		totalPrice := int64(0)
		for _, orderItem := range orderItems {
			totalPrice += orderItem.ProductTotalPrice
		}
	}

	// 检查优惠券是否可用
	couponInfo, err := s.OrderRpc.GetCouponInfo(ctx, couponId)
	if err != nil {
		return 0, err
	}
	if couponInfo.Status != constants.CouponStatusUnused {
		return 0, merror.NewMerror(merror.CouponIsUsed, "Coupon have been used")
	}
	if couponInfo.ExpiredAt < time.Now().Unix() {
		if err := s.OrderRpc.LetCouponExpire(ctx, couponId); err != nil {
			logger.CtxErrorf(ctx, "Let coupon expire error: %v", err)
		}
		return 0, merror.NewMerror(merror.CouponExpired, "Coupon expired")
	}
	if couponInfo.Threshold > int64(totalPrice) {
		return 0, merror.NewMerror(merror.CouponTotalPriceLessThanCouponThreshold, "Coupon total price less than coupon threshold")
	}
	if couponInfo.Type == constants.CouponTypeDiscount {
		totalPrice = totalPrice - couponInfo.Amount
	} else if couponInfo.Type == constants.CouponTypePercent {
		totalPrice = totalPrice - couponInfo.Threshold + couponInfo.Rate/100
	} else {
		return 0, merror.NewMerror(merror.CouponTypeNotExist, "Coupon type not exist")
	}

	addressInfo, err := s.OrderRpc.GetAddressInfo(ctx, addressId)
	if err != nil {
		return 0, err
	}
	order := &model.Order{
		OrderId:          orderId,
		UserId:           userInfo.Id,
		OrderStatus:      constants.OrderWaitPay,
		TotalPrice:       int64(totalPrice),
		PayPrice:         int64(totalPrice),
		ReceiverName:     userInfo.Username,
		ReceiverEmail:    userInfo.Email,
		ReceiverProvince: addressInfo.Province,
		ReceiverCity:     addressInfo.City,
		ReceiverDetail:   addressInfo.Detail,
	}
	err = s.OrderRpc.LockCoupon(ctx, couponId)
	if err != nil {
		return 0, err
	}
	err = s.OrderDB.CreateOrder(ctx, order, orderItems)
	if err != nil {
		return 0, err
	}
	go func() {
		err := s.OrderMq.SendOrderMessage(ctx, orderId, couponId)
		if err == nil {
			_ = s.OrderDB.MarkMsg(ctx, orderId)
		}
	}()
	return orderId, nil
}

func (s *OrderService) HandleMqMsg(ctx context.Context, orderId int64, couponId int64) error {
	orderStatus, err := s.OrderDB.GetOrderStatus(ctx, orderId)
	if err != nil {
		return err
	}
	if orderStatus == constants.OrderWaitPay {
		err = s.OrderRpc.ReleaseCoupon(ctx, couponId)
		if err != nil {
			return err
		}
		err = s.OrderDB.UpdateOrderStatus(ctx, orderId, constants.OrderCanceled)
		if err != nil {
			return err
		}
	}
	if err := s.OrderDB.MarkMsg(ctx, orderId); err != nil {
		return err
	}
	return nil
}

func (s *OrderService) ConsumeOrderMessage(ctx context.Context) error {
	return s.OrderMq.ConsumeOrderMessage(ctx, s.HandleMqMsg)
}

func (s *OrderService) SendMsg(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := s.ConsumeOrderMessage(ctx); err != nil {
				logger.CtxErrorf(ctx, "Consume order message error: %v", err)
			}
		}
	}
}
