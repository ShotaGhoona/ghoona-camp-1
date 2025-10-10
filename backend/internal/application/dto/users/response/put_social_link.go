package response

// PutSocialLinkResponse PUT /users/{userId}/social-links/{linkId} のレスポンス形式
type PutSocialLinkResponse struct {
	Data struct {
		SocialLink PutSocialLinkSocialLink `json:"social_link"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// PutSocialLinkSocialLink PUT /users/{userId}/social-links/{linkId} のSNSリンク情報
type PutSocialLinkSocialLink struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Platform  string `json:"platform"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	IsPublic  bool   `json:"is_public"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}