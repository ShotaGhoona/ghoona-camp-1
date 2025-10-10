package response

// GetUserRivalsResponse GET /users/{userId}/rivals のレスポンス形式
type GetUserRivalsResponse struct {
	Data struct {
		Rivals []GetUserRivalsRival `json:"rivals"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetUserRivalsRival GET /users/{userId}/rivals のライバル情報
type GetUserRivalsRival struct {
	ID          string                    `json:"id"`
	UserID      string                    `json:"user_id"`
	RivalUserID string                    `json:"rival_user_id"`
	CreatedAt   string                    `json:"created_at"`
	UpdatedAt   string                    `json:"updated_at"`
	RivalUser   GetUserRivalsRivalUser    `json:"rival_user"`
}

// GetUserRivalsRivalUser GET /users/{userId}/rivals のライバルユーザー基本情報
type GetUserRivalsRivalUser struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url,omitempty"`
}