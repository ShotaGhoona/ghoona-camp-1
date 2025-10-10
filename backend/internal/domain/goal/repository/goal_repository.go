package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/goal/entity"
)

// GoalRepository 目標リポジトリインターフェース
type GoalRepository interface {
	// Create 目標を作成する
	Create(ctx context.Context, goal *entity.Goal) error

	// GetByID IDで目標を取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Goal, error)

	// GetByUserID ユーザーIDで目標一覧を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.Goal, error)

	// GetPublicGoals 公開目標一覧を取得する
	GetPublicGoals(ctx context.Context, limit, offset int) ([]*entity.Goal, error)

	// Update 目標を更新する
	Update(ctx context.Context, goal *entity.Goal) error

	// Delete 目標を削除する
	Delete(ctx context.Context, id uuid.UUID) error

	// SearchByUserID ユーザーの目標を検索する
	SearchByUserID(ctx context.Context, userID uuid.UUID, query string, limit, offset int) ([]*entity.Goal, error)

	// SearchPublic 公開目標を検索する
	SearchPublic(ctx context.Context, query string, limit, offset int) ([]*entity.Goal, error)

	// GetByDateRange 期間で目標を絞り込み取得する
	GetByDateRange(ctx context.Context, userID uuid.UUID, startFrom, startTo, endFrom, endTo *time.Time, limit, offset int) ([]*entity.Goal, error)
}