package router

import (
	"context"
	"lonely-monitor/biz/router/contact"
	"lonely-monitor/biz/router/user"
	"lonely-monitor/pkg/config"
	"lonely-monitor/pkg/errno"
	"lonely-monitor/pkg/utils"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/google/uuid"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/requestid"
)

// Register 注册所有路由
func Register(r *server.Hertz) {
	// 添加全局恢复中间件
	r.Use(recovery.Recovery())

	// 添加 request id 中间件
	r.Use(requestid.New(
		requestid.WithGenerator(func(c context.Context, ctx *app.RequestContext) string {
			// 为空则重新生成
			traceId := utils.GetTraceId(ctx)
			if traceId == "" {
				traceId = uuid.New().String()
				ctx.Request.Header.Set(utils.GetRequestKey(), traceId)
			}

			return traceId
		}),
		requestid.WithCustomHeaderStrKey(requestid.HeaderStrKey(utils.GetRequestKey())),
	))

	// 添加全局日志中间件
	r.Use(accesslog.New())

	// 限流，1000次/秒 302
	// r.Use(ratelimit.NewRateLimiter(ratelimit.WithLimit(1000), ratelimit.WithInterval(time.Second)))

	version := config.GetConfig().PongTime.Version
	if version == "" {
		version = time.Now().Format("20060102150405")
	}
	r.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(errno.HttpSuccess, utils.Success(ctx, version))
	})

	// 注册用户模块路由
	user.Register(r)
	// 注册联系人模块路由
	contact.Register(r)
}
