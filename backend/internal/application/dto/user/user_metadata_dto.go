package user

import (
	"time"

	"github.com/google/uuid"

	"ghoona-camp-backend/internal/domain/user/entity"
)

// UpdateUserMetadataRequest はユーザーメタデータ更新のリクエストDTO
type UpdateUserMetadataRequest struct {
	DisplayName      *string  `json:"displayName" binding:"omitempty,max=100"`
	ProfileImageURL  *string  `json:"profileImageUrl" binding:"omitempty"`
	Tagline          *string  `json:"tagline" binding:"omitempty,max=150"`
	Bio              *string  `json:"bio" binding:"omitempty,max=1000"`
	Vision           *string  `json:"vision" binding:"omitempty,max=2000"`
	VisionPublic     *bool    `json:"visionPublic"`
	Timezone         *string  `json:"timezone" binding:"omitempty"`
	Skills           []string `json:"skills" binding:"max=20,dive,max=50"`
	Interests        []string `json:"interests" binding:"max=20,dive,max=50"`
}

// CreateUserMetadataRequest はユーザーメタデータ作成のリクエストDTO
type CreateUserMetadataRequest struct {
	DisplayName      *string  `json:"displayName" binding:"omitempty,max=100"`
	ProfileImageURL  *string  `json:"profileImageUrl" binding:"omitempty"`
	Tagline          *string  `json:"tagline" binding:"omitempty,max=150"`
	Bio              *string  `json:"bio" binding:"omitempty,max=1000"`
	Vision           *string  `json:"vision" binding:"omitempty,max=2000"`
	VisionPublic     *bool    `json:"visionPublic"`
	Timezone         string   `json:"timezone" binding:"omitempty"`
	Skills           []string `json:"skills" binding:"max=20,dive,max=50"`
	Interests        []string `json:"interests" binding:"max=20,dive,max=50"`
}

// UserMetadataResponse はユーザーメタデータのレスポンスDTO
type UserMetadataResponse struct {
	ID               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"userId"`
	DisplayName      *string   `json:"displayName"`
	ProfileImageURL  *string   `json:"profileImageUrl"`
	Tagline          *string   `json:"tagline"`
	Bio              *string   `json:"bio"`
	Vision           *string   `json:"vision"`
	VisionPublic     bool      `json:"visionPublic"`
	Timezone         string    `json:"timezone"`
	Skills           []string  `json:"skills"`
	Interests        []string  `json:"interests"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// UserMetadataResponseFromEntity はエンティティからレスポンスDTOに変換する
func UserMetadataResponseFromEntity(metadata *entity.UserMetadata) *UserMetadataResponse {
	if metadata == nil {
		return nil
	}

	skills := metadata.Skills
	if skills == nil {
		skills = []string{}
	}

	interests := metadata.Interests
	if interests == nil {
		interests = []string{}
	}

	return &UserMetadataResponse{
		ID:               metadata.ID,
		UserID:           metadata.UserID,
		DisplayName:      metadata.DisplayName,
		ProfileImageURL:  metadata.ProfileImageURL,
		Tagline:          metadata.Tagline,
		Bio:              metadata.Bio,
		Vision:           metadata.Vision,
		VisionPublic:     metadata.VisionPublic.Bool(),
		Timezone:         metadata.Timezone,
		Skills:           skills,
		Interests:        interests,
		CreatedAt:        metadata.CreatedAt,
		UpdatedAt:        metadata.UpdatedAt,
	}
}