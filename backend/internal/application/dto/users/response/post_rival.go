package response

// PostRivalResponse POST /users/{userId}/rivals のレスポンス形式
type PostRivalResponse struct {
	Data struct {
		Rival PostRivalRival `json:"rival"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PostRivalRival POST /users/{userId}/rivals のライバル情報
type PostRivalRival struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	RivalUserID string `json:"rival_user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}