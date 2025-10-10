package response

// PostGoalResponse POST /goals のレスポンス形式
type PostGoalResponse struct {
	Data struct {
		Goal PostGoalGoal `json:"goal"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PostGoalGoal POST /goals の目標情報
type PostGoalGoal struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartedAt   string `json:"started_at"`
	EndedAt     string `json:"ended_at,omitempty"`
	IsActive    bool   `json:"is_active"`
	IsPublic    bool   `json:"is_public"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}