package request

// PostSocialLinkRequest POST /users/{userId}/social-links のリクエストボディ
type PostSocialLinkRequest struct {
	Platform string `json:"platform" validate:"required,min=1,max=50"`
	URL      string `json:"url" validate:"required,url,max=2048"`
	Title    string `json:"title" validate:"required,min=1,max=100"`
	IsPublic bool   `json:"is_public"`
}