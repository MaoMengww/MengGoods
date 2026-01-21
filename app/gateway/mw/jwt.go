package mw

import (
	"MengGoods/pkg/base"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/utils"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := string(c.GetHeader(constants.AuthHeader))
		claims, err := utils.CheckToken(token)
		if err != nil {
			logger.CtxError(ctx, err)
			c.Abort()
			return
		}
		access, refresh, err := utils.CreateGatewayToken(claims.Uid)
		if err != nil {
			logger.CtxError(ctx, err)
			base.ResErr(c, err)
			c.Abort()
			return
		}

		// 实现规范化服务透传，不需要中间进行编解码
		ctx = mcontext.WithUserIDInContext(ctx, claims.Uid)
		ctx = mcontext.WithStreamUserIDInContext(ctx, claims.Uid)
		c.Header(constants.AccessTokenHeader, access)
		c.Header(constants.RefreshTokenHeader, refresh)
		c.Next(ctx)
	}
}
