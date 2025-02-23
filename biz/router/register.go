package router

import (
	"context"
	"lonely-monitor/biz/router/contact"
	"lonely-monitor/biz/router/user"
	"lonely-monitor/pkg/config"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/logger/accesslog"
)

// Register 注册所有路由
func Register(r *server.Hertz) {
	// 添加全局日志中间件
	r.Use(accesslog.New())

	version := config.GetConfig().PongTime.Version
	if version == "" {
		version = time.Now().Format("20060102150405")
	}
	r.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, map[string]string{"message": version})
	})

	// 注册用户模块路由
	user.Register(r)
	// 注册联系人模块路由
	contact.Register(r)
}
