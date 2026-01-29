package pack

import (
	domain "MengGoods/app/payment/domain/model"
	"MengGoods/kitex_gen/model"
)

func ToDomainPaymentOrder(paymentOrder *model.PaymentOrderInfo) *domain.PaymentOrder {
	return &domain.PaymentOrder{
		PaymentNo: paymentOrder.PaymentNo,
		OrderId:   paymentOrder.OrderId,
		UserId:    paymentOrder.UserId,
		Amount:    paymentOrder.Amount,
		Method:    int(paymentOrder.PaymentMethod),
		Status:    int(paymentOrder.Status),
	}
}
func ToDomainPaymentRefund(paymentRefund *model.PaymentRefundItem) *domain.PaymentRefund {
    return &domain.PaymentRefund{
		OrderItemId: paymentRefund.OrderItemId,
		PaymentNo: paymentRefund.PaymentNo,
		RefundNo: paymentRefund.RefundNo,
		SellerId: paymentRefund.SellerId,
		UserId: paymentRefund.UserId,
		Reason: paymentRefund.RefundReason,
		Amount: paymentRefund.RefundAmount,
		Status: int(paymentRefund.Status),
    }
}