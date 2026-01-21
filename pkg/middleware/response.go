package middleware

//
//import (
//	"MengGoods/pkg/merror"
//	"context"
//	"errors"
//
//	"github.com/cloudwego/kitex/pkg/endpoint"
//)
//
////本中间件主要用于判断返回的错误是否可控, 能否加入熔断判断
//func PackResp() endpoint.Middleware {
//	return func(next endpoint.Endpoint) endpoint.Endpoint {
//		return func(ctx context.Context, req, resp any) (err error) {
//			err = next(ctx, req, resp)
//			if err == nil {
//				return
//			}
//			// 正确识别*merror.Merror指针类型
//			var mErr *merror.Merror
//			if errors.As(err, &mErr) && mErr != nil {
//				// 只有系统级错误才保持为error返回
//				if mErr.Code >= merror.InternalDatabaseErrorCode {
//					return err
//				}
//			}
//
//			// 其他类型的错误保持原样
//			return err
//		}
//	}
//}

/* mport (
	"context"
	"errors"
	"fmt"

	"MengGoods/kitex_gen/model"
	"MengGoods/pkg/merror"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

type baser interface {
	IsSetBase() bool
	GetBase() *model.BaseResp
	SetBase(base *model.BaseResp)
}

// PackResp 会对所有的响应和 error 进行拦截
// 主要用于为 response 中的 BaseResp 加上 code 和 msg (仅当 response 未设置且 error 为 merror)
// 让业务层不需要处理相关的操作, 只需要返回 error 即可.
//
// PackResp 流程如下:
//  1. 获取到 response 后尝试从 response 中提取出来 model.BaseResp
//  2. 尝试判断 err 是否为 merror, 如果是的话说明这是一个可控的 error, 我们对外部返回 nil 即可
//  3. 自动为 BaseResp 设置错误码和消息
func PackResp() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) error {
			err := next(ctx, req, resp)

			// 检查resp是否实现了baser接口
			res, ok := resp.(baser)
	
			if !ok {
				// 如果响应对象不包含Base字段，直接返回原始错误
				return err
			}

			// 如果Base未设置，创建一个新的Base
			if !res.IsSetBase() {
				res.SetBase(&model.BaseResp{})
			}

			base := res.GetBase()

			// 处理业务错误
			var merr *merror.Merror
			if errors.As(err, &merr) {
				// 业务错误：设置Base字段，返回nil错误
				if base.Code == 0 {
					base.Code = int64(merr.Code)
					base.Message = merr.Msg
				}
				return nil
			}

			// 系统错误：保持原始错误
			if err != nil && base.Code == 0 {
				// 如果是系统错误但Base未设置，设置为系统错误
				base.Code = int64(merror.InternalServerErrorCode)
				base.Message = "Internal server error"
			}

			// 成功的情况
			if err == nil && base.Code == 0 {
				base.Code = int64(merror.SuccessCode)
				base.Message = "Success"
			}

			return err
		}
	}
} */