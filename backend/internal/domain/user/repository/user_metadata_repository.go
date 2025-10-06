package repository

import (
	"context"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserMetadataRepository はユーザーメタデータのデータアクセスインターフェースを定義する
type UserMetadataRepository interface {
	GetByUserID(ctx context.Context, userID common.UUID) (*entity.UserMetadata, error) // ユーザーIDでメタデータを取得
	Create(ctx context.Context, metadata *entity.UserMetadata) error                  // メタデータを作成
	Update(ctx context.Context, metadata *entity.UserMetadata) error                  // メタデータを更新
	Delete(ctx context.Context, userID common.UUID) error                               // メタデータを削除
}