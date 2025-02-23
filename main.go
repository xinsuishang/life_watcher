package main

import (
	"context"
	"fmt"
	"lonely-monitor/biz/service/monitor"
	"time"

	"lonely-monitor/biz/router"
	"lonely-monitor/pkg/config"
	"lonely-monitor/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func main() {
	// 初始化 Hertz 服务实例
	h := initServer()

	// 初始化系统组件
	initComponents(h)

	// 启动监控服务
	monitorSvc := initMonitorService(h)
	monitorSvc.Start()

	hlog.Info("服务启动完成")
	h.Spin()

	hlog.Info("正在关闭服务...")
	hlog.Info("服务已关闭")
}

// initServer 初始化服务实例
func initServer() *server.Hertz {
	return server.New(
		server.WithHostPorts(fmt.Sprintf("%s:%d", config.GetConfig().Server.Host, config.GetConfig().Server.Port)),
		server.WithExitWaitTime(time.Second*10),
	)
}

// initComponents 初始化系统组件
func initComponents(h *server.Hertz) {
	// 初始化日志
	utils.InitLogger(h)

	// 初始化数据库
	// dal.Init()

	// 注册路由
	router.Register(h)
}

// initMonitorService 初始化监控服务
func initMonitorService(h *server.Hertz) *monitor.MonitorService {
	monitorSvc := monitor.NewMonitorService()

	// 注册服务关闭回调
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		hlog.Info("服务关闭，停止监控服务")
		monitorSvc.Stop()
	})

	return monitorSvc
}
