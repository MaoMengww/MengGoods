package service

import (
	"MengGoods/app/payment/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"MengGoods/pkg/utils"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func (s *PaymentService) CheckAndSetRefundLimit(ctx context.Context, orderItemId int64) error {
	key := s.PaymentCache.GetRefundKey(ctx, orderItemId)
	ok, err := s.PaymentCache.CheckDailyRefundCount(ctx, key)
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, err.Error())
	}
	if ok {
		return merror.NewMerror(merror.PaymentRefundLimitExceededErrorCode, "本单已在二十四小时内申请过退款，无法再次申请")
	}
	count, err := s.PaymentCache.SetOrIncrRefundKey(ctx, key)
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, err.Error())
	}
	if count > constants.RefundTryMaxCountInMinute {
		return merror.NewMerror(merror.PaymentRefundLimitExceededErrorCode, "退款尝试次数超过最大限制")
	}
	err = s.PaymentCache.SetDailyRefund(ctx, key)
	if err != nil {
		return merror.NewMerror(merror.InternalCacheErrorCode, err.Error())
	}
	return nil
}

func (s *PaymentService) SetPaymentToken(ctx context.Context, orderId int64, token string, expiredTime int64) error {
	key := s.PaymentCache.GetPaymentKey(ctx, orderId)
	expiredDuration, err := s.PaymentCache.GetExpiredDuration(ctx, expiredTime)
	if err != nil {
		return err
	}
	err = s.PaymentCache.SetPaymentToken(ctx, key, token, expiredDuration)
	if err != nil {
		return err
	}
	return nil
}

func (s *PaymentService) CreatePaymentOrder(ctx context.Context, orderId int64, method int) error {
	orderInfo, err := s.PaymentRpc.GetOrderInfoById(ctx, orderId)
	if err != nil {
		return merror.NewMerror(merror.InternalRpcErrorCode, "get order info failed: "+err.Error())
	}
	paymentOrder := &model.PaymentOrder{
		PaymentNo: utils.GenerateIDStr(),
		OrderId:   orderId,
		UserId:    orderInfo.UserId,
		Amount:    orderInfo.TotalPrice,
		Method:    method,
		Status:    constants.PaymentOrderStatusWaiting,
	}
	err = s.PaymentDB.CreatePaymentOrder(ctx, paymentOrder)
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}

func (s *PaymentService) CreateRefundOrder(ctx context.Context, orderItemId int64, reason string) error {
	userId, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	orderItemInfo, err := s.PaymentRpc.GetOrderItemInfoById(ctx, orderItemId)
	if err != nil {
		return merror.NewMerror(merror.InternalRpcErrorCode, "get order item info failed: "+err.Error())
	}
	payment, err := s.PaymentDB.GetPaymentOrderByOrderId(ctx, orderItemInfo.OrderId)
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	if payment == nil {
		return merror.NewMerror(merror.OrderNotExistErrorCode, "payment order not exist")
	}
	if payment.UserId != userId {
		return merror.NewMerror(merror.PaymentOrderNotBelongToUserErrorCode, "payment order not belong to user")
	}
	if payment.Status != constants.PaymentOrderStatusPaid {
		return merror.NewMerror(merror.PaymentOrderNotPaidErrorCode, "payment order not paid")
	}
	paymentRefund := &model.PaymentRefund{
		RefundNo:    utils.GenerateIDStr(),
		OrderItemId: orderItemInfo.Id,
		PaymentNo:   payment.PaymentNo,
		SellerId:    orderItemInfo.SellerId,
		UserId:      userId,
		Amount:      orderItemInfo.ProductTotalPrice,
		Reason:      reason,
	}
	err = s.PaymentDB.CreateRefundOrder(ctx, paymentRefund)
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}

// HMAC生成支付token
func (s *PaymentService) CreateToken(ctx context.Context, orderId int64) (token string, expiredTime int64, err error) {
	expiredTime = time.Now().Unix() + constants.PaymentExpireTime
	secret := []byte(viper.GetString("secret.paymentSecret"))
	//计算 HMAC-SHA256 哈希
	h := hmac.New(sha256.New, secret)
	_, err = h.Write([]byte(fmt.Sprintf("%d%d", orderId, expiredTime)))
	if err != nil {
		return "", 0, merror.NewMerror(merror.InternalServerErrorCode, "create hmac failed: "+err.Error())
	}
	token = hex.EncodeToString(h.Sum(nil))
	if token == "" {
		return "", 0, merror.NewMerror(merror.InternalServerErrorCode, "create token failed")
	}
	return token, expiredTime, nil
}

func (s *PaymentService) GetExpiredAtAndDelPaymentToken(ctx context.Context, key string, token string) (bool, int64, error) {
	ok, ttl, err := s.PaymentCache.GetTTLAndDelPaymentToken(ctx, key, token)
	if err != nil {
		return false, 0, merror.NewMerror(merror.InternalCacheErrorCode, err.Error())
	}
	return ok, time.Now().Unix() + ttl, nil
}

func (s *PaymentService) GetOrderStatus(ctx context.Context, orderId int64) error {
	ok, expiredTime, err := s.PaymentRpc.IsOrderExist(ctx, orderId)
	if err != nil {
		return merror.NewMerror(merror.InternalRpcErrorCode, "get order exist failed: "+err.Error())
	}
	if !ok {
		return merror.NewMerror(merror.OrderNotExistErrorCode, "order not exist")
	}
	if expiredTime > time.Now().Unix() {
		return merror.NewMerror(merror.OrderExpiredErrorCode, "payment order expired")
	}
	return nil
}

func (s *PaymentService) ConfirmPaymentOrder(ctx context.Context, paymentOrder *model.PaymentOrder) error {
	if paymentOrder == nil {
		return merror.NewMerror(merror.OrderNotExistErrorCode, "payment order not exist")
	}
	err := s.PaymentDB.ConfirmPaymentOrder(ctx, paymentOrder)
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}

func (s *PaymentService) ConfirmRefundOrder(ctx context.Context, refundOrder *model.PaymentRefund) error {
	if refundOrder == nil {
		return merror.NewMerror(merror.OrderNotExistErrorCode, "refund order not exist")
	}
	err := s.PaymentDB.ConfirmRefundOrder(ctx, refundOrder)
	if err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, err.Error())
	}
	return nil
}
