package constants

const (
	Normal = 0 //普通用户
	Admin  = 1 //管理员
)

// 限流
const (
	MaxConnection = 1000
	MaxQPS        = 100
)

// 优惠券状态
const (
	CouponStatusAll = -1 //所有
	CouponStatusUnused = 0 //未使用
	CouponStatusLocked = 1 //已锁定
	CouponStatusUsed   = 2 //已使用
	CouponStatusExpired = 3 //已过期
)

//优惠卷类型
const (
	CouponTypeDiscount = 0 //折扣优惠券
	CouponTypePercent  = 1 //满减优惠券
)

//库存操作类型
const (
	CreateType = 1 + iota  
	AddType
	LockType
	UnlockType
	DeductType
)

//订单状态
const (
	OrderWaitPay = 0 //待支付
	OrderPaid = 1 //已支付
	OrderDelivered = 2 //已发货
	OrderConfirmed = 3 //已确认
	OrderCanceled = 4 //已取消
)

//订单消息状态
const (
	MsgWaited = 0
	MsgSended = 1
)

//支付方式
const (
	PayMethodAliPay = 0 //支付宝
	PayMethodWeChatPay = 1 //微信支付
)

//支付状态
const (
	PaymentOrderStatusWaiting = 0 //未支付
	PaymentOrderStatusProcessing = 1 //处理中
	PaymentOrderStatusPaid = 2 //已支付
	PaymentOrderStatusRefunded = 3 //支付失败
)

//退款状态
const (
	RefundStatusWaiting = 0 //待处理
	RefundStatusProcessing = 1 //处理中
	RefundStatusSuccess = 2 //成功
	RefundStatusFailed = 3 //失败
)

//流水类型
const (
	PaymentTransactionTypePay = 0 //支付
	PaymentTransactionTypeRefund = 1 //退款
)
