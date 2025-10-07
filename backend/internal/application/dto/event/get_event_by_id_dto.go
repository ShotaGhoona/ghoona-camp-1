// Package dto GET /events/{eventId} API用のDTO定義
package dto

// GetEventByIDResponse イベント詳細取得レスポンス
type GetEventByIDResponse struct {
	Data      GetEventByIDData `json:"data"`
	Message   string           `json:"message"`
	Timestamp string           `json:"timestamp"`
}

// GetEventByIDData イベント詳細データ
type GetEventByIDData struct {
	ID                string     `json:"id"`
	Creator           CreatorDTO `json:"creator"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	EventType         string     `json:"event_type"`
	ScheduledDate     string     `json:"scheduled_date"`
	StartTime         string     `json:"start_time"`
	EndTime           string     `json:"end_time"`
	MaxParticipants   int        `json:"max_participants"`
	IsRecurring       bool       `json:"is_recurring"`
	RecurrencePattern *string    `json:"recurrence_pattern,omitempty"`
	DiscordChannelID  *string    `json:"discord_channel_id,omitempty"`
	ParticipantCount  int        `json:"participant_count"`
	IsUserRegistered  bool       `json:"is_user_registered"`
	IsActive          bool       `json:"is_active"`
	CreatedAt         string     `json:"created_at"`
	UpdatedAt         string     `json:"updated_at"`
}
