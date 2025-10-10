package notification

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Notification 通知テーブル
type Notification struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Type        string         `gorm:"type:varchar(50);not null" json:"type"`
	Title       string         `gorm:"type:varchar(100);not null" json:"title"`
	Message     string         `gorm:"type:text;not null" json:"message"`
	Data        datatypes.JSON `gorm:"type:jsonb" json:"data"`
	IsRead      bool           `gorm:"default:false" json:"is_read"`
	ScheduledAt *time.Time     `json:"scheduled_at"`
	SentAt      *time.Time     `json:"sent_at"`
	CreatedAt   time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"default:now()" json:"updated_at"`
}

// BeforeCreate GORM hook - IDを自動生成
func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (Notification) TableName() string {
	return "notifications"
}