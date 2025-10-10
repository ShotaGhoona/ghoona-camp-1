package response

// PutUserAchievementResponse PUT /users/{userId}/achievements/{titleId} のレスポンス形式
type PutUserAchievementResponse struct {
	Data struct {
		Achievement PutUserAchievementAchievement `json:"achievement"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PutUserAchievementAchievement PUT /users/{userId}/achievements/{titleId} の実績情報
type PutUserAchievementAchievement struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	TitleID    string `json:"title_id"`
	AchievedAt string `json:"achieved_at"`
	IsCurrent  bool   `json:"is_current"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}