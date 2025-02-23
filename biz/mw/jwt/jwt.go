package jwt

import (
	"context"
	"lonely-monitor/pkg/utils"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func JWTAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		authHeader := string(ctx.Request.Header.Peek("Authorization"))
		if authHeader == "" {
			ctx.JSON(consts.StatusUnauthorized, utils.Error(401, "未授权访问"))
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(consts.StatusUnauthorized, utils.Error(401, "无效的认证格式"))
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(consts.StatusUnauthorized, utils.Error(401, "无效的 token"))
			ctx.Abort()
			return
		}

		// 将用户信息存储到上下文中
		ctx.Set("userID", claims.UserID)
		ctx.Next(c)
	}
}
