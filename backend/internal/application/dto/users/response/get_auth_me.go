package response

// GetAuthMeResponse GET /auth/me のレスポンス形式
type GetAuthMeResponse struct {
	Data struct {
		User GetAuthMeUser `json:"user"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetAuthMeUser GET /auth/me のユーザー情報
type GetAuthMeUser struct {
	ID          string `json:"id"`
	ClerkID     string `json:"clerk_id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	DiscordID   string `json:"discord_id,omitempty"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}