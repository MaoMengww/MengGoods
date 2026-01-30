package resp

type GetPaymentTokenResp struct {
    PaymentToken string `json:"paymentToken"`
	ExpiredAt int64 `json:"expiredAt"`
}

type PaymentResp struct {
}

type PaymentRefundResp struct {
}

type ReviewRefundResp struct {
}