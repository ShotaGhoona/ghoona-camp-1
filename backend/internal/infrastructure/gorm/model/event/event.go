package event

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Event 朝活イベントテーブル
type Event struct {
	ID                uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatorID         uuid.UUID  `gorm:"type:uuid;not null;index" json:"creator_id"`
	Title             string     `gorm:"type:varchar(200);not null" json:"title"`
	Description       *string    `gorm:"type:text" json:"description"`
	EventType         string     `gorm:"type:varchar(50);default:'general'" json:"event_type"`
	ScheduledDate     time.Time  `gorm:"type:date;not null" json:"scheduled_date"`
	StartTime         time.Time  `gorm:"type:time;not null" json:"start_time"`
	EndTime           time.Time  `gorm:"type:time;not null" json:"end_time"`
	MaxParticipants   *int       `gorm:"type:integer" json:"max_participants"`
	IsRecurring       bool       `gorm:"default:false" json:"is_recurring"`
	RecurrencePattern *string    `gorm:"type:varchar(50)" json:"recurrence_pattern"`
	DiscordChannelID  *string    `gorm:"type:varchar(255)" json:"discord_channel_id"`
	IsActive          bool       `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"default:now()" json:"updated_at"`

	// Relationships
	Participants []EventParticipant `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"participants,omitempty"`
}

// BeforeCreate GORM hook - IDを自動生成
func (e *Event) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (Event) TableName() string {
	return "events"
}