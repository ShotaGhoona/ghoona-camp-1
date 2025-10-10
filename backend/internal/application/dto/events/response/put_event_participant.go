package response

// PutEventParticipantResponse PUT /events/{eventId}/participants/{userId} のレスポンス形式
type PutEventParticipantResponse struct {
	Data struct {
		Participation PutEventParticipation `json:"participation"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PutEventParticipation PUT /events/{eventId}/participants/{userId} の参加情報
type PutEventParticipation struct {
	ID        string `json:"id"`
	EventID   string `json:"event_id"`
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}