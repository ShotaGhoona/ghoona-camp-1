package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserRivalRepository defines the interface for user rival data access
type UserRivalRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserRival, error)
	Create(ctx context.Context, rival *entity.UserRival) error
	Delete(ctx context.Context, id uuid.UUID) error
	CountByUserID(ctx context.Context, userID uuid.UUID) (int, error)
}