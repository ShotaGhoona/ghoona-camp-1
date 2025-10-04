package title

import (
	"time"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title/entity"
)

// TitleResponse は称号のレスポンスDTO
type TitleResponse struct {
	ID           common.UUID `json:"id"`
	Level        int         `json:"level"`
	NameJP       string      `json:"nameJp"`
	NameEN       string      `json:"nameEn"`
	Description  string      `json:"description"`
	RequiredDays int         `json:"requiredDays"`
	ImageURL     *string     `json:"imageUrl"`
	ColorTheme   *string     `json:"colorTheme"`
	IsActive     bool        `json:"isActive"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

// TitleListResponse は称号一覧のレスポンスDTO
type TitleListResponse struct {
	Titles []TitleResponse `json:"titles"`
	Total  int             `json:"total"`
}

// TitleResponseFromEntity はエンティティからレスポンスDTOに変換する
func TitleResponseFromEntity(title *entity.Title) *TitleResponse {
	if title == nil {
		return nil
	}

	return &TitleResponse{
		ID:           title.ID,
		Level:        title.Level,
		NameJP:       title.NameJP,
		NameEN:       title.NameEN,
		Description:  title.Description,
		RequiredDays: title.RequiredDays,
		ImageURL:     title.ImageURL,
		ColorTheme:   title.ColorTheme,
		IsActive:     title.IsActive.Bool(),
		CreatedAt:    title.CreatedAt,
		UpdatedAt:    title.UpdatedAt,
	}
}

// TitleListFromEntities はエンティティリストからレスポンスDTOリストに変換する
func TitleListFromEntities(titles []*entity.Title) []TitleResponse {
	if titles == nil {
		return []TitleResponse{}
	}

	responses := make([]TitleResponse, len(titles))
	for i, title := range titles {
		if response := TitleResponseFromEntity(title); response != nil {
			responses[i] = *response
		}
	}
	return responses
}