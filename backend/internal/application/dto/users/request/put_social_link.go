package request

// PutSocialLinkRequest PUT /users/{userId}/social-links/{linkId} のリクエストボディ
type PutSocialLinkRequest struct {
	URL      *string `json:"url,omitempty" validate:"omitempty,url,max=2048"`
	Title    *string `json:"title,omitempty" validate:"omitempty,min=1,max=100"`
	IsPublic *bool   `json:"is_public,omitempty"`
}