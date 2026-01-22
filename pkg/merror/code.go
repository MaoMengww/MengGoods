package merror

const (
	SuccessCode = 10000
)

// 参数错误
const (
	ParamPasswordTooShort      = 20001 + iota //密码过短
	ParamPasswordTooLong                      //密码过长
	ParamEmailInvalid                         //邮箱格式错误
	ParamUsernameTooLong                      //用户名过长
	ParamFromContextFailed                    //从context中获取用户ID失败
	ParamSpuNameTooLong                       //商品名称过长
	ParamSpuDescriptionTooLong                //商品描述过长
	ParamSkuPriceInvalid                      //商品价格无效
	ParamCategoryNameTooLong                  //分类名称过长
	ParamSkuNameTooLong                       //sku名称过长
	ParamSkuDescriptionTooLong                //sku描述过长
	ParamIDInvalid                            //商品ID无效
	ParamCountInvalid                         //商品数量无效
)

// 业务相关
const (
	UserAlreadyExist            = 30001 + iota //用户已存在
	UserNotExist                               //用户不存在
	PasswordIncorrect                          //密码错误
	UserIsBanned                               //用户已被ban
	UserNotBanned                              //用户未被ban
	PasswordNotMatch                           //密码错误
	TokenExpired                               //token过期
	AuthNoOperatePermissionCode                //无操作权限
	PermissionDenied                           //权限不足
	QuerySettingInvalid                        //查询设置无效
	RedisNotFound                              //redis未找到
	StockNotEnough                             //库存不足
	StockNotExist                              //库存不存在
	CartIsEmptyCode                            //购物车为空
)

// 系统内部错误
const (
	InternalDatabaseErrorCode = 40001 + iota //数据库错误
	InternalCacheErrorCode                   //缓存错误
	InternalNetworkErrorCode                 //网络错误
	InternalRpcErrorCode                     //rpc错误
	InternalESErrorCode                      //es错误
	InternalMqErrorCode                      //mq错误
)

const (
	InternalServerErrorCode = 50001 + iota //服务器内部错误
)
