package entity

import (
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/value"
)

// UserSocialLink represents a user's social media link
type UserSocialLink struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Platform  value.Platform
	URL       string
	Title     *string
	IsPublic  value.PublicFlag
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUserSocialLink creates a new UserSocialLink entity
func NewUserSocialLink(userID uuid.UUID, platform value.Platform, url string) *UserSocialLink {
	return &UserSocialLink{
		ID:        uuid.New(),
		UserID:    userID,
		Platform:  platform,
		URL:       url,
		IsPublic:  value.PublicTrue,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}