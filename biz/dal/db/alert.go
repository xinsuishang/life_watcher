// AIGC START
package db

import "time"

// AlertRecord 预警记录
type AlertRecord struct {
	BaseModel
	UserID          int64     `json:"user_id" gorm:"column:user_id;type:bigint"`                          // 用户ID
	LastCheckInTime time.Time `json:"last_check_in_time" gorm:"column:last_check_in_time;type:timestamp"` // 最后签到时间
	AlertTime       time.Time `json:"alert_time" gorm:"column:alert_time;type:timestamp"`                 // 预警时间
	Status          int       `json:"status" gorm:"column:status;type:int"`                               // 状态：0-待处理，1-已通知，2-已处理
}

// CreateAlertRecord 创建预警记录
func CreateAlertRecord(record *AlertRecord) error {
	return GetDB().Create(record).Error
}

// UpdateAlertRecord 更新预警记录
func UpdateAlertRecord(id int64, updates map[string]interface{}) error {
	return GetDB().Model(&AlertRecord{}).Where("id = ?", id).Updates(updates).Error
}

// GetPendingAlerts 获取待处理的预警记录
func GetPendingAlerts(lastID int64) ([]AlertRecord, error) {
	var records []AlertRecord
	query := GetDB().Where("status = ? AND alert_time <= ?", 0, time.Now())
	if lastID > 0 {
		query = query.Where("id > ?", lastID)
	}
	query.Order("id asc")
	query.Limit(100)
	err := query.Find(&records).Error
	return records, err
}

// AIGC END
