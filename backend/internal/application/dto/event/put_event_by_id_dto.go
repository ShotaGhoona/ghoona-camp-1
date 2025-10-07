// Package dto PUT /events/{eventId} API用のDTO定義
package dto

// PutEventByIDRequest イベント更新リクエスト
type PutEventByIDRequest struct {
	Title            string `json:"title" binding:"omitempty,min=1,max=200"`
	Description      string `json:"description" binding:"omitempty,max=2000"`
	MaxParticipants  int    `json:"max_participants" binding:"omitempty,min=1,max=100"`
	DiscordChannelID string `json:"discord_channel_id"`
}

// PutEventByIDResponse イベント更新レスポンス
type PutEventByIDResponse struct {
	Data      PutEventByIDData `json:"data"`
	Message   string           `json:"message"`
	Timestamp string           `json:"timestamp"`
}

// PutEventByIDData イベント更新データ
type PutEventByIDData struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	MaxParticipants  int    `json:"max_participants"`
	DiscordChannelID string `json:"discord_channel_id"`
	UpdatedAt        string `json:"updated_at"`
}