package api

import (
	"MengGoods/app/payment/usecase"
	payment "MengGoods/kitex_gen/payment"
	"MengGoods/pkg/base"
	"context"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{
	usecase *usecase.PaymentUsecase
}

func NewPaymentServiceImpl(usecase *usecase.PaymentUsecase) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		usecase: usecase,
	}
}

// GetPaymentToken implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) GetPaymentToken(ctx context.Context, req *payment.GetPaymentTokenReq) (resp *payment.GetPaymentTokenResp, err error) {
	resp = payment.NewGetPaymentTokenResp()
	token, expiredAt, err := s.usecase.GetPaymentToken(ctx, req.OrderId, int(req.PaymentMethod))
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	resp.PaymentToken = token
	resp.ExpiredAt = expiredAt
	return
}

// Payment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Payment(ctx context.Context, req *payment.PaymentReq) (resp *payment.PaymentResp, err error) {
	// TODO: Your code here...
	resp = payment.NewPaymentResp()
	err = s.usecase.Payment(ctx, req.OrderId, req.PaymentToken)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return
}

// PaymentRefund implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) PaymentRefund(ctx context.Context, req *payment.PaymentRefundReq) (resp *payment.PaymentRefundResp, err error) {
	// TODO: Your code here...
	resp = payment.NewPaymentRefundResp()
	err = s.usecase.PaymentRefund(ctx, req.OrderItemId, req.RefundReason)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return
}

// Refund implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) ReviewRefund(ctx context.Context, req *payment.ReviewRefundReq) (resp *payment.ReviewRefundResp, err error) {
	// TODO: Your code here...
	resp = payment.NewReviewRefundResp()
	err = s.usecase.ReviewRefund(ctx, req.OrderItemId, req.Approve)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = base.BuildBaseResp(nil)
	return
}
