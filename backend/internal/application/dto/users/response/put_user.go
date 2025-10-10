package response

// PutUserResponse PUT /users/{userId} のレスポンス形式
type PutUserResponse struct {
	Data struct {
		User PutUserUser `json:"user"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PutUserUser PUT /users/{userId} のユーザー情報
type PutUserUser struct {
	ID          string              `json:"id"`
	ClerkID     string              `json:"clerk_id"`
	Email       string              `json:"email"`
	Username    string              `json:"username"`
	DisplayName string              `json:"display_name,omitempty"`
	AvatarURL   string              `json:"avatar_url,omitempty"`
	DiscordID   string              `json:"discord_id,omitempty"`
	IsActive    bool                `json:"is_active"`
	CreatedAt   string              `json:"created_at"`
	UpdatedAt   string              `json:"updated_at"`
	Metadata    *PutUserMetadata    `json:"metadata,omitempty"`
}

// PutUserMetadata PUT /users/{userId} のメタデータ情報
type PutUserMetadata struct {
	ID          string   `json:"id"`
	DisplayName string   `json:"display_name,omitempty"`
	Tagline     string   `json:"tagline,omitempty"`
	Bio         string   `json:"bio,omitempty"`
	Skills      []string `json:"skills,omitempty"`
	Interests   []string `json:"interests,omitempty"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}