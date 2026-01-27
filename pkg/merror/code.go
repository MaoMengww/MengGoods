package merror

const (
	SuccessCode = 10000
)

// 参数错误
const (
	ParamPasswordTooShort             = 20001 + iota //密码过短
	ParamPasswordTooLong                             //密码过长
	ParamEmailInvalid                                //邮箱格式错误
	ParamUsernameTooLong                             //用户名过长
	ParamFromContextFailed                           //从context中获取用户ID失败
	ParamSpuNameTooLong                              //商品名称过长
	ParamSpuDescriptionTooLong                       //商品描述过长
	ParamSkuPriceInvalid                             //商品价格无效
	ParamCategoryNameTooLong                         //分类名称过长
	ParamSkuNameTooLong                              //sku名称过长
	ParamSkuDescriptionTooLong                       //sku描述过长
	ParamIDInvalid                                   //商品ID无效
	ParamCountInvalid                                //商品数量无效
	ParamCouponDiscountAmountInvalid                 //优惠券折扣金额无效
	ParamCouponThresholdInvalid                      //优惠券门槛无效
	ParamCouponDiscountPercentInvalid                //优惠券折扣比例无效
	ParamCouponNameTooLong                           //优惠券名称过长
	ParamCouponRemarkTooLong                         //优惠券备注过长
	ParamCouponDurationInvalid                       //优惠券有效期无效
	ParamCouponTypeInvalid                           //优惠券类型无效
	ParamCouponStatusInvalid                         //优惠券状态无效
	ParamCouponTotalNumInvalid                       //优惠券总数量无效
)

// 业务相关
const (
	UserAlreadyExist                        = 30001 + iota //用户已存在
	UserNotExist                                           //用户不存在
	PasswordIncorrect                                      //密码错误
	UserIsBanned                                           //用户已被ban
	UserNotBanned                                          //用户未被ban
	PasswordNotMatch                                       //密码错误
	TokenExpired                                           //token过期
	AuthNoOperatePermissionCode                            //无操作权限
	PermissionDenied                                       //权限不足
	QuerySettingInvalid                                    //查询设置无效
	RedisNotFound                                          //redis未找到
	StockNotEnough                                         //库存不足
	StockNotExist                                          //库存不存在
	CartIsEmptyCode                                        //购物车为空
	CouponIsUsed                                           //优惠券已被使用
	CouponExpired                                          //优惠券已过期
	CouponTypeNotExist                                     //优惠券类型不存在
	CouponTotalPriceLessThanCouponThreshold                //未达优惠卷使用门槛

)

// 系统内部错误
const (
	InternalDatabaseErrorCode = 40001 + iota //数据库错误
	InternalCacheErrorCode                   //缓存错误
	InternalNetworkErrorCode                 //网络错误
	InternalRpcErrorCode                     //rpc错误
	InternalESErrorCode                      //es错误
	InternalKafkaErrorCode                   //kafka错误
	InternalRabbitMqErrorCode                //rabbitmq错误
)

const (
	InternalServerErrorCode = 50001 + iota //服务器内部错误
)
