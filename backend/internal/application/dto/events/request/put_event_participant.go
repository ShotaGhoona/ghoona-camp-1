package request

// PutEventParticipantRequest PUT /events/{eventId}/participants/{userId} のリクエスト形式
type PutEventParticipantRequest struct {
	Status string `json:"status" validate:"required,oneof=registered cancelled"` // 参加ステータス
}