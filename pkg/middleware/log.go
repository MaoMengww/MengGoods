package middleware

import (
	"MengGoods/kitex_gen/model" // 引用你的 BaseResp 定义
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"context"
	"reflect"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

type BaseGetter interface {
	GetBase() *model.BaseResp
}

func ErrorLog() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp any) (err error) {
			err = next(ctx, req, resp)

			//系统错误
			if err != nil {
				logger.CtxErrorf(ctx, "SystemError: %v", err)
				return err
			}

			//反射获取resp
			respVal := reflect.ValueOf(resp)
			if respVal.Kind() == reflect.Pointer {
				respVal = respVal.Elem()
			}
			successField := respVal.FieldByName("Success")
			if !successField.IsValid() || successField.IsNil() {
				return nil
			}

			realResp := successField.Interface()

		
			
			// 使用类型断言检查 resp 是否包含 BaseResp
			if respGetter, ok := realResp.(BaseGetter); ok {
				
				baseResp := respGetter.GetBase()
				if baseResp != nil {
					// 获取业务状态码
					code := baseResp.Code
					msg := baseResp.Message

					// 根据 Code 分级打印日志
					switch {
					case code == merror.SuccessCode:
						logger.CtxInfof(ctx, "Success: %s", msg)

					case code >= merror.InternalDatabaseErrorCode:
						logger.CtxErrorf(ctx, "InternalServerError: [%d] %s", code, msg)

					case code >= merror.UserAlreadyExist:
						logger.CtxWarnf(ctx, "LogicError: [%d] %s", code, msg)

					case code >= merror.ParamPasswordTooShort:
						logger.CtxWarnf(ctx, "ParamError: [%d] %s", code, msg)

					default:
						logger.CtxErrorf(ctx, "Unknown: [%d] %s", code, msg)
					}
				}
			}

			return nil
		}
	}
}
