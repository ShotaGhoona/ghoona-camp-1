package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserSocialLinkRepository はユーザーソーシャルリンクのデータアクセスインターフェースを定義する
type UserSocialLinkRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserSocialLink, error) // ユーザーIDでソーシャルリンク一覧を取得
	GetByID(ctx context.Context, id uuid.UUID) (*entity.UserSocialLink, error)            // IDでソーシャルリンクを取得
	Create(ctx context.Context, link *entity.UserSocialLink) error                        // ソーシャルリンクを作成
	Update(ctx context.Context, link *entity.UserSocialLink) error                        // ソーシャルリンクを更新
	Delete(ctx context.Context, id uuid.UUID) error                                       // ソーシャルリンクを削除
}