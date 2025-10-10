package event

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EventParticipant イベント参加者テーブル
type EventParticipant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	EventID   uuid.UUID `gorm:"type:uuid;not null;index" json:"event_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Status    string    `gorm:"type:varchar(20);default:'registered'" json:"status"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`

	// Relationships
	Event Event `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"event,omitempty"`
}

// BeforeCreate GORM hook - IDを自動生成
func (ep *EventParticipant) BeforeCreate(tx *gorm.DB) error {
	if ep.ID == uuid.Nil {
		ep.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (EventParticipant) TableName() string {
	return "event_participants"
}