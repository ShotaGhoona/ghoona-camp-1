package request

// PutGoalRequest PUT /goals/{goalId} のリクエストボディ
type PutGoalRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=1,max=200"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=1,max=1000"`
	EndedAt     *string `json:"ended_at,omitempty" validate:"omitempty,date"`
	IsPublic    *bool   `json:"is_public,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}