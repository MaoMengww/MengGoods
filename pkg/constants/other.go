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
