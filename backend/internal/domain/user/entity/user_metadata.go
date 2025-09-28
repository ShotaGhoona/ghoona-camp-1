package entity

import (
	"time"

	"ghoona-camp-backend/internal/domain/user/value"

	"github.com/google/uuid"
)

// UserMetadata represents detailed user profile information
type UserMetadata struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	DisplayName     *string
	ProfileImageURL *string
	Tagline         *string
	Bio             *string
	Vision          *string
	VisionPublic    value.PublicFlag
	Timezone        string
	Skills          []string
	Interests       []string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// NewUserMetadata creates a new UserMetadata entity
func NewUserMetadata(userID uuid.UUID) *UserMetadata {
	return &UserMetadata{
		ID:           uuid.New(),
		UserID:       userID,
		VisionPublic: value.PublicFalse,
		Timezone:     "Asia/Tokyo",
		Skills:       make([]string, 0),
		Interests:    make([]string, 0),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
