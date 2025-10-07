// Package dto POST /events API用のDTO定義
package dto

// PostEventsRequest イベント作成リクエスト
type PostEventsRequest struct {
	Title             string `json:"title" binding:"required,min=1,max=200"`
	Description       string `json:"description" binding:"max=2000"`
	EventType         string `json:"event_type" binding:"required,oneof=general study exercise meditation creative business"`
	ScheduledDate     string `json:"scheduled_date" binding:"required,datetime=2006-01-02"`
	StartTime         string `json:"start_time" binding:"required,datetime=15:04:05"`
	EndTime           string `json:"end_time" binding:"required,datetime=15:04:05"`
	MaxParticipants   int    `json:"max_participants" binding:"omitempty,min=1,max=100" default:"10"`
	IsRecurring       bool   `json:"is_recurring"`
	RecurrencePattern string `json:"recurrence_pattern" binding:"omitempty,oneof=daily weekly monthly"`
	DiscordChannelID  string `json:"discord_channel_id"`
}

// PostEventsResponse イベント作成レスポンス
type PostEventsResponse struct {
	Data      PostEventsData `json:"data"`
	Message   string         `json:"message"`
	Timestamp string         `json:"timestamp"`
}

// PostEventsData イベント作成データ
type PostEventsData struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	EventType         string  `json:"event_type"`
	ScheduledDate     string  `json:"scheduled_date"`
	StartTime         string  `json:"start_time"`
	EndTime           string  `json:"end_time"`
	MaxParticipants   int     `json:"max_participants"`
	IsRecurring       bool    `json:"is_recurring"`
	RecurrencePattern *string `json:"recurrence_pattern,omitempty"`
	DiscordChannelID  *string `json:"discord_channel_id,omitempty"`
	ParticipantCount  int     `json:"participant_count"`
	IsActive          bool    `json:"is_active"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}