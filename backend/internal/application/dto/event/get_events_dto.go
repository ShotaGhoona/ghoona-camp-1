// Package dto GET /events API用のDTO定義
package dto

import "ghoona-camp-backend/internal/application/dto/event/shared"

// GetEventsResponse イベント一覧取得レスポンス
type GetEventsResponse struct {
	Data      []GetEventsDataItem `json:"data"`
	Message   string              `json:"message"`
	Timestamp string              `json:"timestamp"`
}

// GetEventsDataItem イベント一覧の個別アイテム
type GetEventsDataItem struct {
	ID                string           `json:"id"`
	Creator           shared.CreatorDTO `json:"creator"`
	Title             string           `json:"title"`
	Description       string           `json:"description"`
	EventType         string           `json:"event_type"`
	ScheduledDate     string           `json:"scheduled_date"`
	StartTime         string           `json:"start_time"`
	EndTime           string           `json:"end_time"`
	MaxParticipants   int              `json:"max_participants"`
	IsRecurring       bool             `json:"is_recurring"`
	RecurrencePattern *string          `json:"recurrence_pattern,omitempty"`
	DiscordChannelID  *string          `json:"discord_channel_id,omitempty"`
	ParticipantCount  int              `json:"participant_count"`
	IsUserRegistered  bool             `json:"is_user_registered"`
	IsActive          bool             `json:"is_active"`
	CreatedAt         string           `json:"created_at"`
	UpdatedAt         string           `json:"updated_at"`
}