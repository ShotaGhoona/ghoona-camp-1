// Package model GORMで使用するデータベーステーブルモデルを定義する
// ドメインエンティティとDBテーブル間のO/Rマッピングを担当
package model

import (
	"time"

	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/value"

	"github.com/google/uuid"
)

// Event イベントテーブルのGORMモデル
type Event struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatorID         uuid.UUID `gorm:"type:uuid;not null;index" json:"creator_id"`
	Title             string    `gorm:"type:varchar(200);not null" json:"title"`
	Description       string    `gorm:"type:text" json:"description"`
	EventType         string    `gorm:"type:varchar(20);not null;index" json:"event_type"`
	ScheduledDate     time.Time `gorm:"type:date;not null;index" json:"scheduled_date"`
	StartTime         time.Time `gorm:"type:time;not null" json:"start_time"`
	EndTime           time.Time `gorm:"type:time;not null" json:"end_time"`
	MaxParticipants   int       `gorm:"not null;default:10" json:"max_participants"`
	IsRecurring       bool      `gorm:"not null;default:false" json:"is_recurring"`
	RecurrencePattern string    `gorm:"type:varchar(20);default:'none'" json:"recurrence_pattern"`
	DiscordChannelID  string    `gorm:"type:varchar(255)" json:"discord_channel_id"`
	IsActive          bool      `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// EventParticipant イベント参加者テーブルのGORMモデル
type EventParticipant struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	EventID   uuid.UUID `gorm:"type:uuid;not null;index" json:"event_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Status    string    `gorm:"type:varchar(20);not null;default:'registered';index" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName テーブル名を指定する
func (Event) TableName() string {
	return "events"
}

// TableName テーブル名を指定する
func (EventParticipant) TableName() string {
	return "event_participants"
}

// ToEntity GORMモデルからドメインエンティティへ変換する
func (e *Event) ToEntity() *entity.Event {
	return &entity.Event{
		ID:                e.ID,
		CreatorID:         e.CreatorID,
		Title:             e.Title,
		Description:       e.Description,
		EventType:         value.EventType(e.EventType),
		ScheduledDate:     e.ScheduledDate,
		StartTime:         e.StartTime,
		EndTime:           e.EndTime,
		MaxParticipants:   e.MaxParticipants,
		IsRecurring:       e.IsRecurring,
		RecurrencePattern: value.RecurrencePattern(e.RecurrencePattern),
		DiscordChannelID:  e.DiscordChannelID,
		IsActive:          e.IsActive,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
	}
}

// ToEntity GORMモデルからドメインエンティティへ変換する
func (ep *EventParticipant) ToEntity() *entity.EventParticipant {
	return &entity.EventParticipant{
		ID:        ep.ID,
		EventID:   ep.EventID,
		UserID:    ep.UserID,
		Status:    value.ParticipantStatus(ep.Status),
		CreatedAt: ep.CreatedAt,
		UpdatedAt: ep.UpdatedAt,
	}
}

// FromEventEntity ドメインエンティティからGORMモデルへ変換する
func FromEventEntity(e *entity.Event) *Event {
	return &Event{
		ID:                e.ID,
		CreatorID:         e.CreatorID,
		Title:             e.Title,
		Description:       e.Description,
		EventType:         string(e.EventType),
		ScheduledDate:     e.ScheduledDate,
		StartTime:         e.StartTime,
		EndTime:           e.EndTime,
		MaxParticipants:   e.MaxParticipants,
		IsRecurring:       e.IsRecurring,
		RecurrencePattern: string(e.RecurrencePattern),
		DiscordChannelID:  e.DiscordChannelID,
		IsActive:          e.IsActive,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
	}
}

// FromEventParticipantEntity ドメインエンティティからGORMモデルへ変換する
func FromEventParticipantEntity(ep *entity.EventParticipant) *EventParticipant {
	return &EventParticipant{
		ID:        ep.ID,
		EventID:   ep.EventID,
		UserID:    ep.UserID,
		Status:    string(ep.Status),
		CreatedAt: ep.CreatedAt,
		UpdatedAt: ep.UpdatedAt,
	}
}