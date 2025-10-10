package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/title/entity"
	"ghoona-camp-backend/internal/domain/title/vo"
)

// TitleRepository 称号リポジトリインターフェース
type TitleRepository interface {
	// Create 称号を作成する
	Create(ctx context.Context, title *entity.Title) error

	// GetByID IDで称号を取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Title, error)

	// GetByLevel レベルで称号を取得する
	GetByLevel(ctx context.Context, level vo.TitleLevel) (*entity.Title, error)

	// GetAll すべての称号を取得する
	GetAll(ctx context.Context) ([]*entity.Title, error)

	// GetActive 有効な称号一覧を取得する
	GetActive(ctx context.Context) ([]*entity.Title, error)

	// Update 称号を更新する
	Update(ctx context.Context, title *entity.Title) error

	// Delete 称号を削除する
	Delete(ctx context.Context, id uuid.UUID) error
}