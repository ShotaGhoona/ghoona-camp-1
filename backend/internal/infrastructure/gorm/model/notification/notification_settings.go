package notification

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NotificationSettings 通知設定テーブル
type NotificationSettings struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID               uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	AchievementEnabled   bool      `gorm:"default:true" json:"achievement_enabled"`
	ReminderEnabled      bool      `gorm:"default:true" json:"reminder_enabled"`
	RivalUpdateEnabled   bool      `gorm:"default:true" json:"rival_update_enabled"`
	EventReminderEnabled bool      `gorm:"default:true" json:"event_reminder_enabled"`
	ReminderTime         time.Time `gorm:"type:time;default:'21:00'" json:"reminder_time"`
	CreatedAt            time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt            time.Time `gorm:"default:now()" json:"updated_at"`
}

// BeforeCreate GORM hook - IDを自動生成
func (ns *NotificationSettings) BeforeCreate(tx *gorm.DB) error {
	if ns.ID == uuid.Nil {
		ns.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (NotificationSettings) TableName() string {
	return "notification_settings"
}