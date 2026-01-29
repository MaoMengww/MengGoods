package usecase

import (
	"MengGoods/app/payment/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"context"
)

func(u *PaymentUsecase) GetPaymentToken(ctx context.Context, orderId int64, method int) (string, int64, error) {
	err := u.PaymentService.GetOrderStatus(ctx, orderId)
	if err != nil {
		return "", 0, err
	}
	err = u.PaymentService.CreatePaymentOrder(ctx, orderId, method)
	if err != nil {
		return "", 0, err
	}
	token, expiredTime, err := u.PaymentService.CreateToken(ctx, orderId)
	if err != nil {
		return "", 0, err
	}
	err = u.PaymentService.SetPaymentToken(ctx, orderId, token, expiredTime)
	if err != nil {
		return "", 0, err
	}
	if err := u.PaymentService.SetPaymentToken(ctx, orderId, token, expiredTime); err != nil {
		return "", 0, err
	}
	return token, expiredTime, nil
}

func(u *PaymentUsecase) Payment(ctx context.Context, orderId int64, token string) error {
	err := u.PaymentService.GetOrderStatus(ctx, orderId)
	if err != nil {
		return err
	}
	key := u.PaymentCache.GetPaymentKey(ctx, orderId)
	ok, expiredTime, err := u.PaymentService.GetExpiredAtAndDelPaymentToken(ctx, key, token)
	if !ok && err == nil {
		return merror.NewMerror(merror.PaymentTokenNotExistOrExpired, "payment token not exist or expired")
	}
	
	//redis错误, 尝试查询数据库
	var paymentOrder *model.PaymentOrder
	if err != nil {
		paymentOrder, err = u.PaymentDB.GetPaymentOrderByOrderId(ctx, orderId)
		if err != nil {
			return err
		}
		if paymentOrder.Status == constants.PaymentOrderStatusPaid {
			return nil
		}
		if paymentOrder.Status != constants.PaymentOrderStatusWaiting {
			return merror.NewMerror(merror.PaymentOrderNotPaidErrorCode, "payment token not exist or expired")
		}
	}
	if paymentOrder == nil {
		paymentOrder, err = u.PaymentDB.GetPaymentOrderByOrderId(ctx, orderId) 
		if err != nil {
			return err
		}
	}
	//锁定，方便后台对账系统及时发现异常支付订单
	if err := u.PaymentDB.UpdatePaymentOrderStatus(ctx, paymentOrder.PaymentNo, constants.PaymentOrderStatusProcessing); err != nil {
		return err
	}

	//模拟银行返回结果
	MockBankResult := true
	if MockBankResult{
		err = u.PaymentService.ConfirmPaymentOrder(ctx, paymentOrder)
		if err != nil {
			//支付成功库没更新, 打日志依靠后台对账系统
			logger.CtxErrorf(ctx, "confirm payment order failed, err: %v", err)
		}
		if err := u.PaymentRpc.MarkOrderPaid(ctx, orderId); err != nil {
			logger.CtxErrorf(ctx, "mark order paid failed, err: %v", err)
		}
		return nil
	} else {
        _ = u.PaymentDB.UpdatePaymentOrderStatus(ctx, paymentOrder.PaymentNo, constants.PaymentOrderStatusWaiting)
        // 2. 归还 Token (允许用户重试)
        if err == nil { // 只有 Redis 正常时才归还
             _ = u.PaymentService.SetPaymentToken(ctx, orderId, token, expiredTime)
        }
        return merror.NewMerror(merror.InternalServerErrorCode, "payment failed by bank")
	}
}

func (u *PaymentUsecase) PaymentRefund(ctx context.Context, orderItemId int64, reason string) error {
	userId, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	orderItem, err := u.PaymentRpc.GetOrderItemInfoById(ctx, orderItemId)
	if err != nil {
		return err
	}
	if userId != orderItem.UserId {
		return merror.NewMerror(merror.PaymentOrderNotBelongToUserErrorCode, "payment order not belong to user")
	}
	err = u.PaymentService.CheckAndSetRefundLimit(ctx, orderItemId)
	if err != nil {
		return err
	}
	err = u.PaymentService.CreateRefundOrder(ctx, orderItemId, reason)
	if err != nil {
		return err
	}
	return nil
}

func (u *PaymentUsecase) ReviewRefund(ctx context.Context, orderItemId int64, approve bool) error {
	userId, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	paymentRefund , err := u.PaymentDB.GetRefundOrderByOrderItemId(ctx, orderItemId)
	if err != nil {
		return err
	}
	if userId != paymentRefund.SellerId{
		return merror.NewMerror(merror.PaymentOrderNotBelongToSellerErrorCode, "payment order not belong to seller")
	}
	if paymentRefund.Status == constants.RefundStatusProcessing {
		return merror.NewMerror(merror.RefundOrderProcessingErrorCode, "payment refund order not processing")
	}
	paymentRefund.Status = constants.RefundStatusProcessing
	if err := u.PaymentDB.UpdateRefundOrderStatus(ctx, orderItemId, paymentRefund.Status); err != nil {
		return err
	}
	if approve {
		if err := u.PaymentService.ConfirmRefundOrder(ctx, paymentRefund); err != nil {
			//退款成功库没更新, 打日志依靠后台对账系统
			logger.CtxErrorf(ctx, "confirm refund order failed, err: %v", err)
		}
	} else {
		if err := u.PaymentDB.UpdateRefundOrderStatus(ctx, orderItemId, constants.RefundStatusFailed); err != nil {
			return err
		}
	}
	return nil
}


