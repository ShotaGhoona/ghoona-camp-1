package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserVisionRepository ユーザービジョンリポジトリインターフェース
type UserVisionRepository interface {
	// Create ユーザービジョンを作成する
	Create(ctx context.Context, vision *entity.UserVision) error

	// GetByUserID ユーザーIDでビジョンを取得する
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserVision, error)

	// Update ユーザービジョンを更新する
	Update(ctx context.Context, vision *entity.UserVision) error

	// Delete ユーザービジョンを削除する
	Delete(ctx context.Context, userID uuid.UUID) error
}