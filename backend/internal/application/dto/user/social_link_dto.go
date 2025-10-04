package user

import (
	"time"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// CreateSocialLinkRequest はソーシャルリンク作成のリクエストDTO
type CreateSocialLinkRequest struct {
	Platform string  `json:"platform" binding:"required"`
	URL      string  `json:"url" binding:"required,url"`
	Title    *string `json:"title" binding:"omitempty,max=100"`
	IsPublic *bool   `json:"isPublic"`
}

// UpdateSocialLinkRequest はソーシャルリンク更新のリクエストDTO
type UpdateSocialLinkRequest struct {
	URL      *string `json:"url" binding:"omitempty,url"`
	Title    *string `json:"title" binding:"omitempty,max=100"`
	IsPublic *bool   `json:"isPublic"`
}

// SocialLinkResponse はソーシャルリンクのレスポンスDTO
type SocialLinkResponse struct {
	ID        common.UUID `json:"id"`
	UserID    common.UUID `json:"userId"`
	Platform  string      `json:"platform"`
	URL       string      `json:"url"`
	Title     *string     `json:"title"`
	IsPublic  bool        `json:"isPublic"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

// SocialLinkListResponse はソーシャルリンク一覧のレスポンスDTO
type SocialLinkListResponse struct {
	SocialLinks []SocialLinkResponse `json:"socialLinks"`
	Total       int                  `json:"total"`
}

// SocialLinkResponseFromEntity はエンティティからレスポンスDTOに変換する
func SocialLinkResponseFromEntity(link *entity.UserSocialLink) *SocialLinkResponse {
	if link == nil {
		return nil
	}

	return &SocialLinkResponse{
		ID:        link.ID,
		UserID:    link.UserID,
		Platform:  link.Platform.String(),
		URL:       link.URL,
		Title:     link.Title,
		IsPublic:  link.IsPublic.Bool(),
		CreatedAt: link.CreatedAt,
		UpdatedAt: link.UpdatedAt,
	}
}

// SocialLinkListFromEntities はエンティティリストからレスポンスDTOリストに変換する
func SocialLinkListFromEntities(links []*entity.UserSocialLink) []SocialLinkResponse {
	if links == nil {
		return []SocialLinkResponse{}
	}

	responses := make([]SocialLinkResponse, len(links))
	for i, link := range links {
		if response := SocialLinkResponseFromEntity(link); response != nil {
			responses[i] = *response
		}
	}
	return responses
}