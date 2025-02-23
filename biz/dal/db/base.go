package db

import "time"

type BaseModel struct {
	ID int64 `json:"id" gorm:"primaryKey"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`                             // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间
	DeletedAt time.Time `json:"deleted_at" gorm:"autoUpdateTime;column:deleted_at;type:timestamp"`
}
