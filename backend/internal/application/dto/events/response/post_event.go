package response

// PostEventResponse POST /events のレスポンス形式
type PostEventResponse struct {
	Data struct {
		Event PostEventEvent `json:"event"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PostEventEvent POST /events のイベント情報
type PostEventEvent struct {
	ID                string  `json:"id"`
	CreatorID         string  `json:"creator_id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	EventType         string  `json:"event_type"`
	ScheduledDate     string  `json:"scheduled_date"`
	StartTime         string  `json:"start_time"`
	EndTime           string  `json:"end_time"`
	MaxParticipants   int     `json:"max_participants"`
	IsRecurring       bool    `json:"is_recurring"`
	RecurrencePattern *string `json:"recurrence_pattern"`
	DiscordChannelID  string  `json:"discord_channel_id"`
	IsActive          bool    `json:"is_active"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}