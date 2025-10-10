package response

import "ghoona-camp-backend/internal/application/dto/common"

// GetGoalsPublicResponse GET /goals/public のレスポンス形式
type GetGoalsPublicResponse struct {
	Data struct {
		Goals      []GetGoalsPublicGoal `json:"goals"`
		Pagination *common.Pagination   `json:"pagination,omitempty"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetGoalsPublicGoal GET /goals/public の個別目標情報
type GetGoalsPublicGoal struct {
	ID          string                  `json:"id"`
	UserID      string                  `json:"user_id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	StartedAt   string                  `json:"started_at"`
	EndedAt     string                  `json:"ended_at,omitempty"`
	IsActive    bool                    `json:"is_active"`
	IsPublic    bool                    `json:"is_public"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
	User        GetGoalsPublicUser      `json:"user"`
}

// GetGoalsPublicUser GET /goals/public のユーザー情報
type GetGoalsPublicUser struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url,omitempty"`
}