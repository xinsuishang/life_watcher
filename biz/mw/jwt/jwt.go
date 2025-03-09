package jwt

import (
	"context"
	"lonely-monitor/pkg/consts"
	"lonely-monitor/pkg/errno"
	"lonely-monitor/pkg/utils"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

func JWTAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		authHeader := string(ctx.Request.Header.Peek("Authorization"))
		if authHeader == "" {
			ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.AuthErrCode, "未授权访问"))
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.AuthErrCode, "无效的认证格式"))
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.AuthErrCode, "无效的 token"))
			ctx.Abort()
			return
		}

		// 将用户信息存储到上下文中
		ctx.Set(consts.UserId, claims.UserID)
		ctx.Next(c)
	}
}
