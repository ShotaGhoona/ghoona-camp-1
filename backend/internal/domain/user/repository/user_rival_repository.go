package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserRivalRepository ユーザーライバルリポジトリインターフェース
type UserRivalRepository interface {
	// Create ライバル関係を作成する
	Create(ctx context.Context, rival *entity.UserRival) error

	// GetByID IDでライバル関係を取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.UserRival, error)

	// GetByUserID ユーザーIDでライバル一覧を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserRival, error)

	// Delete ライバル関係を削除する
	Delete(ctx context.Context, id uuid.UUID) error

	// DeleteByUserAndRival ユーザーIDとライバルIDで関係を削除する
	DeleteByUserAndRival(ctx context.Context, userID, rivalUserID uuid.UUID) error

	// CountByUserID ユーザーのライバル数を取得する
	CountByUserID(ctx context.Context, userID uuid.UUID) (int, error)
}