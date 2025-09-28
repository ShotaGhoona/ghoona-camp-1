package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserSocialLinkRepository defines the interface for user social link data access
type UserSocialLinkRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserSocialLink, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.UserSocialLink, error)
	Create(ctx context.Context, link *entity.UserSocialLink) error
	Update(ctx context.Context, link *entity.UserSocialLink) error
	Delete(ctx context.Context, id uuid.UUID) error
}