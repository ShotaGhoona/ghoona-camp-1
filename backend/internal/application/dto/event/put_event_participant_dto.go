// Package dto PUT /events/{eventId}/participants/{userId} API用のDTO定義
package dto

// PutEventParticipantRequest イベント参加ステータス更新リクエスト
type PutEventParticipantRequest struct {
	Status string `json:"status" binding:"required,oneof=registered cancelled"`
}

// PutEventParticipantResponse イベント参加ステータス更新レスポンス
type PutEventParticipantResponse struct {
	Data      PutEventParticipantData `json:"data"`
	Message   string                  `json:"message"`
	Timestamp string                  `json:"timestamp"`
}

// PutEventParticipantData 参加ステータス更新データ
type PutEventParticipantData struct {
	ID        string `json:"id"`
	EventID   string `json:"event_id"`
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}