package attendance

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AttendanceStatistics 参加統計テーブル
type AttendanceStatistics struct {
	ID                   uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID               uuid.UUID  `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	TotalAttendanceDays  int        `gorm:"default:0" json:"total_attendance_days"`
	CurrentStreakDays    int        `gorm:"default:0" json:"current_streak_days"`
	MaxStreakDays        int        `gorm:"default:0" json:"max_streak_days"`
	LastAttendanceDate   *time.Time `gorm:"type:date" json:"last_attendance_date"`
	FirstAttendanceDate  *time.Time `gorm:"type:date" json:"first_attendance_date"`
	TotalDurationMinutes int        `gorm:"default:0" json:"total_duration_minutes"`
	CreatedAt            time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"default:now()" json:"updated_at"`
}

// BeforeCreate GORM hook - IDを自動生成
func (as *AttendanceStatistics) BeforeCreate(tx *gorm.DB) error {
	if as.ID == uuid.Nil {
		as.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (AttendanceStatistics) TableName() string {
	return "attendance_statistics"
}