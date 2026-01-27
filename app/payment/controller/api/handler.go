package main

import (
	payment "MengGoods/kitex_gen/payment"
	"context"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// GetPaymentToken implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) GetPaymentToken(ctx context.Context, req *payment.GetPaymentTokenReq) (resp *payment.GetPaymentTokenResp, err error) {
	// TODO: Your code here...
	return
}

// Payment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Payment(ctx context.Context, req *payment.PaymentReq) (resp *payment.PaymentResp, err error) {
	// TODO: Your code here...
	return
}

// GetRefundToken implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) GetRefundToken(ctx context.Context, req *payment.GetRefundTokenReq) (resp *payment.GetRefundTokenResp, err error) {
	// TODO: Your code here...
	return
}

// Refund implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Refund(ctx context.Context, req *payment.RefundReq) (resp *payment.RefundResp, err error) {
	// TODO: Your code here...
	return
}
