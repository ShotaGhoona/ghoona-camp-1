// Package dto GET /events/{eventId}/participants API用のDTO定義
package dto


// GetEventParticipantsResponse イベント参加者一覧取得レスポンス
type GetEventParticipantsResponse struct {
	Data      []GetEventParticipantsDataItem `json:"data"`
	Message   string                         `json:"message"`
	Timestamp string                         `json:"timestamp"`
}

// GetEventParticipantsDataItem 参加者データアイテム
type GetEventParticipantsDataItem struct {
	ID        string        `json:"id"`
	User      UserDTO        `json:"user"`
	Status    string        `json:"status"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
}