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

const (
	OrderWaitPay = 0 //待支付
	OrderPaid = 1 //已支付
	OrderDelivered = 2 //已发货
	OrderConfirmed = 3 //已确认
	OrderCanceled = 4 //已取消
)

const (
	MsgWaited = 0
	MsgSended = 1
)
