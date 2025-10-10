package request

// PostEventRequest POST /events のリクエスト形式
type PostEventRequest struct {
	Title               string  `json:"title" validate:"required,max=200"`                  // イベント名
	Description         string  `json:"description" validate:"required"`                    // イベント詳細
	EventType           string  `json:"event_type" validate:"required,max=50"`              // イベントタイプ
	ScheduledDate       string  `json:"scheduled_date" validate:"required"`                 // 開催日（YYYY-MM-DD形式）
	StartTime           string  `json:"start_time" validate:"required"`                     // 開始時間（HH:MM:SS形式）
	EndTime             string  `json:"end_time" validate:"required"`                       // 終了時間（HH:MM:SS形式）
	MaxParticipants     int     `json:"max_participants" validate:"required,min=1"`         // 最大参加者数
	IsRecurring         bool    `json:"is_recurring"`                                       // 定期開催かどうか
	RecurrencePattern   *string `json:"recurrence_pattern,omitempty"`                       // 繰り返しパターン
	DiscordChannelID    string  `json:"discord_channel_id" validate:"required,max=255"`     // 対応するDiscordチャンネルID
}