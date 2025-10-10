package response

import "ghoona-camp-backend/internal/application/dto/common"

// GetEventsResponse GET /events のレスポンス形式
type GetEventsResponse struct {
	Data struct {
		Events     []GetEventsEvent      `json:"events"`
		Pagination *common.Pagination    `json:"pagination,omitempty"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetEventsEvent GET /events の個別イベント情報
type GetEventsEvent struct {
	ID                string               `json:"id"`
	CreatorID         string               `json:"creator_id"`
	Title             string               `json:"title"`
	Description       string               `json:"description"`
	EventType         string               `json:"event_type"`
	ScheduledDate     string               `json:"scheduled_date"`
	StartTime         string               `json:"start_time"`
	EndTime           string               `json:"end_time"`
	MaxParticipants   int                  `json:"max_participants"`
	IsRecurring       bool                 `json:"is_recurring"`
	RecurrencePattern *string              `json:"recurrence_pattern"`
	DiscordChannelID  string               `json:"discord_channel_id"`
	IsActive          bool                 `json:"is_active"`
	CreatedAt         string               `json:"created_at"`
	UpdatedAt         string               `json:"updated_at"`
	Creator           GetEventsCreator     `json:"creator"`
}

// GetEventsCreator GET /events のイベント作成者情報
type GetEventsCreator struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url,omitempty"`
}