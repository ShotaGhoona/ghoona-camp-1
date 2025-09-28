package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserMetadataRepository defines the interface for user metadata data access
type UserMetadataRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserMetadata, error)
	Create(ctx context.Context, metadata *entity.UserMetadata) error
	Update(ctx context.Context, metadata *entity.UserMetadata) error
	Delete(ctx context.Context, userID uuid.UUID) error
}