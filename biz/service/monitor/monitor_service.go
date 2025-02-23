package monitor

import (
	"context"
	"lonely-monitor/biz/dal/db"
	"lonely-monitor/pkg/consts"
	"lonely-monitor/pkg/notice"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type MonitorService struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewMonitorService() *MonitorService {
	ctx, cancel := context.WithCancel(context.Background())
	return &MonitorService{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *MonitorService) Start() {
	hlog.Info("开始监控服务")
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.checkAlerts()
			case <-s.ctx.Done():
				hlog.Info("监控服务停止")
				ticker.Stop()
				return
			}
		}
	}()
}

func (s *MonitorService) Stop() {
	hlog.Info("停止监控服务")
	s.cancel()
}

func (s *MonitorService) checkAlerts() {

	// 查找需要处理的预警记录
	lastID := int64(0) // 假设这是上次处理的最大ID
	for {
		alerts, err := db.GetPendingAlerts(lastID)
		if err != nil {
			hlog.Errorf("获取预警记录失败: %v", err)
			return
		}
		if len(alerts) == 0 {
			break
		}
		lastID = alerts[len(alerts)-1].ID

		for _, alert := range alerts {
			// 获取用户信息
			user, err := db.GetUserByID(alert.UserID)
			if err != nil {
				hlog.Errorf("获取用户信息失败, user_id: %d, alert_id: %d, error: %v",
					alert.UserID, alert.ID, err)
				continue
			}

			// 处理预警记录
			if err := s.processAlertRecord(&alert, user); err != nil {
				hlog.Errorf("处理预警记录失败, alert_id: %d, error: %v", alert.ID, err)
				continue
			}

			// 如果预警状态已更新为完成，发送通知
			if alert.Status == 1 {
				level := checkAlertLevel(alert.AlertTime)
				alertLevel, ok := consts.GetAlertLevel(level)
				if !ok {
					hlog.Errorf("获取预警等级失败, alert_id: %d, alert_level: %d", alert.ID, level)
					continue
				}
				if notifier, ok := notice.GetNotifier(alertLevel.Notifier); ok {
					if err := notifier.Notify(alert.UserID); err != nil {
						hlog.Errorf("发送通知失败: %v", err)
					}
				}
			}
		}
	}
}

func checkAlertLevel(lastCheckInTime time.Time) int {
	duration := time.Since(lastCheckInTime)
	return int(duration / (24 * time.Hour))
}

// processAlertRecord 处理预警记录
func (s *MonitorService) processAlertRecord(alert *db.AlertRecord, user *db.User) error {
	// 计算当前预警等级
	level := checkAlertLevel(alert.LastCheckInTime)

	// 仅在预警等级大于3时更新通知状态为完成
	if level > 3 {
		alert.Status = 1 // 更新为已通知
	}

	// 更新预警时间（延后24小时）
	alert.AlertTime = alert.AlertTime.Add(24 * time.Hour)

	// 保存更新
	if err := db.UpdateAlertRecord(alert.ID, map[string]any{
		"alert_time": alert.AlertTime,
		"status":     alert.Status,
	}); err != nil {
		return err
	}

	// 发送通知
	if alert.Status == 1 {
		level := checkAlertLevel(alert.AlertTime)
		alertLevel, ok := consts.GetAlertLevel(level)
		if !ok {
			hlog.Errorf("获取预警等级失败, alert_id: %d, alert_level: %d", alert.ID, level)
			return nil
		}
		if notifier, ok := notice.GetNotifier(alertLevel.Notifier); ok {
			if err := notifier.Notify(user.ID); err != nil {
				hlog.Errorf("发送通知失败: %v", err)
			}
		}
	}
	return nil
}
