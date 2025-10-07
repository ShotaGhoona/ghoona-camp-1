// Package dto DELETE /events/{eventId} API用のDTO定義
package dto

// DeleteEventByIDResponse イベント削除レスポンス
type DeleteEventByIDResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
