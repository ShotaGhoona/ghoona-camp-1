package response

// PostSocialLinkResponse POST /users/{userId}/social-links のレスポンス形式
type PostSocialLinkResponse struct {
	Data struct {
		SocialLink PostSocialLinkSocialLink `json:"social_link"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PostSocialLinkSocialLink POST /users/{userId}/social-links のSNSリンク情報
type PostSocialLinkSocialLink struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Platform  string `json:"platform"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	IsPublic  bool   `json:"is_public"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}