package response

// PostEventParticipantResponse POST /events/{eventId}/participants のレスポンス形式
type PostEventParticipantResponse struct {
	Data struct {
		Participation PostEventParticipation `json:"participation"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PostEventParticipation POST /events/{eventId}/participants の参加情報
type PostEventParticipation struct {
	ID        string `json:"id"`
	EventID   string `json:"event_id"`
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}