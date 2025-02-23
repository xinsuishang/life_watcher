// AIGC START
package db

import "time"

// NotifyRecord 通知记录
type NotifyRecord struct {
	BaseModel
	UserID     int64     `json:"user_id" gorm:"column:user_id;type:bigint"`              // 用户ID
	AlertID    int64     `json:"alert_id" gorm:"column:alert_id;type:bigint"`            // 预警记录ID
	NotifyType string    `json:"notify_type" gorm:"column:notify_type;type:varchar(50)"` // 通知类型：email/sms/letter
	NotifyTime time.Time `json:"notify_time" gorm:"column:notify_time;type:timestamp"`   // 通知时间
	Status     int       `json:"status" gorm:"column:status;type:int"`                   // 状态：0-发送中，1-发送成功，2-发送失败
	RetryCount int       `json:"retry_count" gorm:"column:retry_count;type:int"`         // 重试次数
	LastError  string    `json:"last_error" gorm:"column:last_error;type:varchar(255)"`  // 最后一次错误信息
	ContactID  int64     `json:"contact_id" gorm:"column:contact_id;type:bigint"`        // 联系方式ID
}

// CreateNotifyRecord 创建通知记录
func CreateNotifyRecord(record *NotifyRecord) error {
	return GetDB().Create(record).Error
}

// UpdateNotifyRecord 更新通知记录
func UpdateNotifyRecord(id int64, updates map[string]interface{}) error {
	return GetDB().Model(&NotifyRecord{}).Where("id = ?", id).Updates(updates).Error
}

// GetPendingNotifyRecords 获取待处理的通知记录
func GetPendingNotifyRecords() ([]NotifyRecord, error) {
	var records []NotifyRecord
	err := GetDB().Where("status = ? AND retry_count < ?", 0, 3).Find(&records).Error
	return records, err
}

// AIGC END
