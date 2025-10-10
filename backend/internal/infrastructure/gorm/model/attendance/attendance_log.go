package attendance

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AttendanceLog 参加ログテーブル
type AttendanceLog struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID           uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	EventID          *uuid.UUID `gorm:"type:uuid;index" json:"event_id"`
	DiscordChannelID string     `gorm:"type:varchar(255);not null" json:"discord_channel_id"`
	JoinedAt         time.Time  `gorm:"not null" json:"joined_at"`
	LeftAt           *time.Time `json:"left_at"`
	DurationMinutes  *int       `gorm:"type:integer" json:"duration_minutes"`
	IsValid          bool       `gorm:"default:true" json:"is_valid"`
	CreatedAt        time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"default:now()" json:"updated_at"`
}

// BeforeCreate GORM hook - IDを自動生成
func (al *AttendanceLog) BeforeCreate(tx *gorm.DB) error {
	if al.ID == uuid.Nil {
		al.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (AttendanceLog) TableName() string {
	return "attendance_logs"
}