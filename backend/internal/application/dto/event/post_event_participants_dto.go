// Package dto POST /events/{eventId}/participants API用のDTO定義
package dto

// PostEventParticipantsResponse イベント参加申込レスポンス
type PostEventParticipantsResponse struct {
	Data      PostEventParticipantsData `json:"data"`
	Message   string                    `json:"message"`
	Timestamp string                    `json:"timestamp"`
}

// PostEventParticipantsData 参加申込データ
type PostEventParticipantsData struct {
	ID        string `json:"id"`
	EventID   string `json:"event_id"`
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}