package request

// PostGoalRequest POST /goals のリクエストボディ
type PostGoalRequest struct {
	UserID      string `json:"user_id" validate:"required,uuid"`
	Title       string `json:"title" validate:"required,min=1,max=200"`
	Description string `json:"description" validate:"required,min=1,max=1000"`
	StartedAt   string `json:"started_at" validate:"required,date"`
	EndedAt     string `json:"ended_at" validate:"omitempty,date"`
	IsPublic    bool   `json:"is_public"`
}