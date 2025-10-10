package response

// GetUserAchievementsResponse GET /users/{userId}/achievements のレスポンス形式
type GetUserAchievementsResponse struct {
	Data struct {
		Achievements []GetUserAchievementsAchievement `json:"achievements"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetUserAchievementsAchievement GET /users/{userId}/achievements の個別実績情報
type GetUserAchievementsAchievement struct {
	ID          string                         `json:"id"`
	UserID      string                         `json:"user_id"`
	TitleID     string                         `json:"title_id"`
	AchievedAt  string                         `json:"achieved_at"`
	IsCurrent   bool                           `json:"is_current"`
	CreatedAt   string                         `json:"created_at"`
	UpdatedAt   string                         `json:"updated_at"`
	Title       GetUserAchievementsTitle       `json:"title"`
}

// GetUserAchievementsTitle GET /users/{userId}/achievements の称号情報
type GetUserAchievementsTitle struct {
	ID          string `json:"id"`
	Level       int    `json:"level"`
	NameJP      string `json:"name_jp"`
	NameEN      string `json:"name_en"`
	Description string `json:"description"`
	RequiredDays int   `json:"required_days"`
	ImageURL    string `json:"image_url,omitempty"`
	ColorTheme  string `json:"color_theme,omitempty"`
}