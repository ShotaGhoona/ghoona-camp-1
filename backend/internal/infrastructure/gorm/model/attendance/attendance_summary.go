package attendance

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AttendanceSummary 参加サマリーテーブル
type AttendanceSummary struct {
	ID                   uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID               uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Date                 time.Time  `gorm:"type:date;not null" json:"date"`
	TotalDurationMinutes int        `gorm:"default:0" json:"total_duration_minutes"`
	SessionCount         int        `gorm:"default:0" json:"session_count"`
	FirstJoinTime        *time.Time `gorm:"type:time" json:"first_join_time"`
	LastLeaveTime        *time.Time `gorm:"type:time" json:"last_leave_time"`
	IsMorningActive      bool       `gorm:"default:false" json:"is_morning_active"`
	CreatedAt            time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"default:now()" json:"updated_at"`
}

// BeforeCreate GORM hook - IDを自動生成
func (as *AttendanceSummary) BeforeCreate(tx *gorm.DB) error {
	if as.ID == uuid.Nil {
		as.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (AttendanceSummary) TableName() string {
	return "attendance_summaries"
}