package response

// GetEventDetailResponse GET /events/{eventId} のレスポンス形式
type GetEventDetailResponse struct {
	Data struct {
		Event GetEventDetailEvent `json:"event"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetEventDetailEvent GET /events/{eventId} のイベント詳細情報
type GetEventDetailEvent struct {
	ID                string                        `json:"id"`
	CreatorID         string                        `json:"creator_id"`
	Title             string                        `json:"title"`
	Description       string                        `json:"description"`
	EventType         string                        `json:"event_type"`
	ScheduledDate     string                        `json:"scheduled_date"`
	StartTime         string                        `json:"start_time"`
	EndTime           string                        `json:"end_time"`
	MaxParticipants   int                           `json:"max_participants"`
	IsRecurring       bool                          `json:"is_recurring"`
	RecurrencePattern *string                       `json:"recurrence_pattern"`
	DiscordChannelID  string                        `json:"discord_channel_id"`
	IsActive          bool                          `json:"is_active"`
	CreatedAt         string                        `json:"created_at"`
	UpdatedAt         string                        `json:"updated_at"`
	Creator           GetEventDetailCreator         `json:"creator"`
	Participants      []GetEventDetailParticipant   `json:"participants"`
}

// GetEventDetailCreator GET /events/{eventId} のイベント作成者情報
type GetEventDetailCreator struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

// GetEventDetailParticipant GET /events/{eventId} の参加者情報
type GetEventDetailParticipant struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Status    string `json:"status"`
}