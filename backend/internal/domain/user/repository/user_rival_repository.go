package repository

import (
	"context"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserRivalRepository defines the interface for user rival data access
type UserRivalRepository interface {
	GetByUserID(ctx context.Context, userID common.UUID) ([]*entity.UserRival, error)
	Create(ctx context.Context, rival *entity.UserRival) error
	Delete(ctx context.Context, id common.UUID) error
	CountByUserID(ctx context.Context, userID common.UUID) (int, error)
}