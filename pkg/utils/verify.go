package utils

import (
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"regexp"
)

//邮箱格式
var emailRe = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

//验证选项, 可插拔
type VerifyOps func() error

//验证多个选项
func Verify(ops ...VerifyOps) error {
	for _, op := range ops {
		if err := op(); err != nil {
			return err
		}
	}
	return nil
}

//验证密码(长度应处于8-16之间)
func VerifyPassword(password string) VerifyOps {
	return func() error {
		if len(password) < 8 {
			return merror.NewMerror(
				merror.ParamPasswordTooShort,
				"password too short",
			)
		}
		if len(password) > 16 {
			return merror.NewMerror(
				merror.ParamPasswordTooLong,	
				"password too long",
			)
		}
		return nil
	}
}


//验证邮箱格式
func VerifyEmail(email string) VerifyOps {
	return func() error {
		if !emailRe.MatchString(email) {
			return merror.NewMerror(
				merror.ParamEmailInvalid,
				"email invalid",
			)
		}
		return nil
	}
}

//验证用户名(最多 10 个中文字符或等长英文字符)
func VerifyUsername(username string) VerifyOps {
	return func() error {
		if len(username) > 30 {
			return merror.NewMerror(
				merror.ParamUsernameTooLong,
				"username too long",
			)
		}
		return nil
	}
}

//验证Spu以及下属sku



func VerifySpuName(name string) VerifyOps {
	return func() error {
		if len(name) > 60 {
			return merror.NewMerror(
				merror.ParamSpuNameTooLong,
				"spu name too long",
			)
		}
		return nil
	}
}

func VerifySpuDescription(description string) VerifyOps {
	return func() error {
		if len(description) > 200 {
			return merror.NewMerror(
				merror.ParamSpuDescriptionTooLong,
				"spu description too long",
			)
		}
		return nil
	}
}

func VerifySkutPrice(price float64) VerifyOps {
	return func() error {
		if price <= 0 {
			return merror.NewMerror(
				merror.ParamSkuPriceInvalid,
				"sku price invalid",
			)
		}
		return nil
	}
}

func VerifySkuName(name string) VerifyOps {
	return func() error {
		if len(name) > 255 {
			return merror.NewMerror(
				merror.ParamSkuNameTooLong,
				"sku name too long",
			)
		}
		return nil
	}
}

func VerifySkuPrice(price int64) VerifyOps {
	return func() error {
		if price <= 0 {
			return merror.NewMerror(
				merror.ParamSkuPriceInvalid,
				"sku price invalid",
			)
		}
		return nil
	}
}

func VerifySkuDescription(description string) VerifyOps {
	return func() error {
		if len(description) > 200 {
			return merror.NewMerror(
				merror.ParamSkuDescriptionTooLong,
				"sku description too long",
			)
		}
		return nil
	}
}

func VerifyCategoryName(name string) VerifyOps {
	return func() error {
		if len(name) > 30 {
			return merror.NewMerror(
				merror.ParamCategoryNameTooLong,
				"category name too long",
			)
		}
		return nil
	}
}

func VerifyCount(count int32) VerifyOps {
	return func() error {
		if count < 0 {
			return merror.NewMerror(
				merror.ParamCountInvalid,
				"count invalid",
			)
		}
		return nil
	}
}

func VerifyCouponName(name string) VerifyOps {
	return func() error {
		if len(name) > 30 {
			return merror.NewMerror(
				merror.ParamCouponNameTooLong,
				"coupon name too long",
			)
		}
		return nil
	}
}

func VerifyCouponRemark(remark string) VerifyOps {
	return func() error {
		if len(remark) > 200 {
			return merror.NewMerror(
				merror.ParamCouponRemarkTooLong,
				"coupon remark too long",
			)
		}
		return nil
	}
}

func VerifyCouponType(t int) VerifyOps {
	return func() error {
		if t != constants.CouponTypeDiscount && t != constants.CouponTypePercent {
			return merror.NewMerror(
				merror.ParamCouponTypeInvalid,
				"coupon type invalid",
			)
		}
		return nil
	}
}

func VerifyCouponStatus(s int) VerifyOps {
	return func() error {
		if s != constants.CouponStatusUnused && s != constants.CouponStatusLocked && s != constants.CouponStatusUsed && s != constants.CouponStatusExpired {
			return merror.NewMerror(
				merror.ParamCouponStatusInvalid,
				"coupon status invalid",
			)
		}
		return nil
	}
}

func VerifyCouponThreshold(t int64) VerifyOps {
	return func() error {
		if t < 0 {
			return merror.NewMerror(
				merror.ParamCouponThresholdInvalid,
				"coupon threshold invalid",
			)
		}
		return nil
	}
}

func VerifyCouponDuration(d int) VerifyOps {
	return func() error {
		if d < 0 {
			return merror.NewMerror(
				merror.ParamCouponDurationInvalid,
				"coupon duration invalid",
			)
		}
		return nil
	}
}

func VerifyDiscountPercent(p int) VerifyOps {
	return func() error {
		if p < 0 || p > 100 {
			return merror.NewMerror(
				merror.ParamCouponDiscountPercentInvalid,
				"coupon discount percent invalid",
			)
		}
		return nil
	}
}

func VerifyDiscountAmount(a int64) VerifyOps {
	return func() error {
		if a < 0 {
			return merror.NewMerror(
				merror.ParamCouponDiscountAmountInvalid,
				"coupon discount amount invalid",
			)
		}
		return nil
	}
}

func VerifyTotalNum(n int64) VerifyOps {
	return func() error {
		if n < 0 {
			return merror.NewMerror(
				merror.ParamCouponTotalNumInvalid,
				"coupon total num invalid",
			)
		}
		return nil
	}
}


