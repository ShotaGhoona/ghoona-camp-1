package response

import "ghoona-camp-backend/internal/application/dto/common"

// GetUsersResponse GET /users のレスポンス形式
type GetUsersResponse struct {
	Data struct {
		Users      []GetUsersUser     `json:"users"`
		Pagination *common.Pagination `json:"pagination,omitempty"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetUsersUser GET /users の個別ユーザー情報
type GetUsersUser struct {
	ID          string `json:"id"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Tagline     string `json:"tagline,omitempty"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
}
