package utils

import (
	"context"
	"lonely-monitor/pkg/config"
	"os"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(h *server.Hertz) {
	// config 获取不到使用 default log path
	hlog.Infof("init logger, config log path: %s, env log path: %s", config.GetConfig().Log.Path, os.Getenv("APP_CONFIG_LOG_PATH"))
	logPath := config.GetConfig().Log.Path
	if logPath == "" {
		logPath = "logs"
	}
	if err := os.MkdirAll(logPath, 0777); err != nil {
		panic(err)
	}

	// 配置日志轮转
	writer := &lumberjack.Logger{
		Filename:   path.Join(logPath, "app.log"),
		MaxSize:    50, // MB
		MaxBackups: 7,
		MaxAge:     30,    // 天
		Compress:   false, // 是否压缩
		// 使用本地时间
		LocalTime: true,
	}
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(writer),
		FlushInterval: time.Second,
	}
	logger := hertzlogrus.NewLogger(
		hertzlogrus.WithLogger(logrus.StandardLogger()),
	)

	// 配置 logrus
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
			logrus.FieldKeyFunc:  "caller",
		},
	})

	// 配置 hlog
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelInfo)
	hlog.SetOutput(asyncWriter)

	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})
}
