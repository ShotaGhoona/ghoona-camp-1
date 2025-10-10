package response

// PutGoalResponse PUT /goals/{goalId} のレスポンス形式
type PutGoalResponse struct {
	Data struct {
		Goal PutGoalGoal `json:"goal"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PutGoalGoal PUT /goals/{goalId} の目標情報
type PutGoalGoal struct {
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