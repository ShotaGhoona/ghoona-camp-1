package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserMetadataRepository ユーザーメタデータリポジトリインターフェース
type UserMetadataRepository interface {
	// Create ユーザーメタデータを作成する
	Create(ctx context.Context, metadata *entity.UserMetadata) error

	// GetByUserID ユーザーIDでメタデータを取得する
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserMetadata, error)

	// Update ユーザーメタデータを更新する
	Update(ctx context.Context, metadata *entity.UserMetadata) error

	// Delete ユーザーメタデータを削除する
	Delete(ctx context.Context, userID uuid.UUID) error
}