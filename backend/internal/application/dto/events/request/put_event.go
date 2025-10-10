package request

// PutEventRequest PUT /events/{eventId} のリクエスト形式
type PutEventRequest struct {
	Title           *string `json:"title,omitempty" validate:"omitempty,max=200"`          // イベント名
	Description     *string `json:"description,omitempty"`                                // イベント詳細
	MaxParticipants *int    `json:"max_participants,omitempty" validate:"omitempty,min=1"` // 最大参加者数
	IsActive        *bool   `json:"is_active,omitempty"`                                  // イベント有効状態
}