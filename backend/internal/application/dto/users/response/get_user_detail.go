package response

// GetUserDetailResponse GET /users/{userId} のレスポンス形式
type GetUserDetailResponse struct {
	Data struct {
		User GetUserDetailUser `json:"user"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetUserDetailUser GET /users/{userId} のユーザー詳細情報
type GetUserDetailUser struct {
	ID          string                       `json:"id"`
	Username    string                       `json:"username"`
	DisplayName string                       `json:"display_name,omitempty"`
	AvatarURL   string                       `json:"avatar_url,omitempty"`
	IsActive    bool                         `json:"is_active"`
	CreatedAt   string                       `json:"created_at"`
	UpdatedAt   string                       `json:"updated_at"`
	Metadata    *GetUserDetailMetadata       `json:"metadata,omitempty"`
	SocialLinks []GetUserDetailSocialLink    `json:"social_links,omitempty"`
}

// GetUserDetailMetadata GET /users/{userId} のメタデータ情報
type GetUserDetailMetadata struct {
	ID          string   `json:"id"`
	DisplayName string   `json:"display_name"`
	Tagline     string   `json:"tagline,omitempty"`
	Bio         string   `json:"bio,omitempty"`
	Skills      []string `json:"skills,omitempty"`
	Interests   []string `json:"interests,omitempty"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

// GetUserDetailSocialLink GET /users/{userId} のSNSリンク情報
type GetUserDetailSocialLink struct {
	ID        string `json:"id"`
	Platform  string `json:"platform"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	IsPublic  bool   `json:"is_public"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}