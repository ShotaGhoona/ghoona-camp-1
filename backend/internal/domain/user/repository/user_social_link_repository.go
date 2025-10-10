package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserSocialLinkRepository ユーザーソーシャルリンクリポジトリインターフェース
type UserSocialLinkRepository interface {
	// Create ソーシャルリンクを作成する
	Create(ctx context.Context, link *entity.UserSocialLink) error

	// GetByID IDでソーシャルリンクを取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.UserSocialLink, error)

	// GetByUserID ユーザーIDでソーシャルリンク一覧を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserSocialLink, error)

	// Update ソーシャルリンクを更新する
	Update(ctx context.Context, link *entity.UserSocialLink) error

	// Delete ソーシャルリンクを削除する
	Delete(ctx context.Context, id uuid.UUID) error
}