package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserRepository ユーザーリポジトリインターフェース
type UserRepository interface {
	// Create ユーザーを作成する
	Create(ctx context.Context, user *entity.User) error

	// GetByID IDでユーザーを取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)

	// GetByClerkID Clerk IDでユーザーを取得する
	GetByClerkID(ctx context.Context, clerkID string) (*entity.User, error)

	// GetByEmail メールアドレスでユーザーを取得する
	GetByEmail(ctx context.Context, email string) (*entity.User, error)

	// Update ユーザーを更新する
	Update(ctx context.Context, user *entity.User) error

	// Delete ユーザーを削除する
	Delete(ctx context.Context, id uuid.UUID) error

	// List ユーザー一覧を取得する
	List(ctx context.Context, limit, offset int) ([]*entity.User, error)

	// Search ユーザーを検索する
	Search(ctx context.Context, query string, limit, offset int) ([]*entity.User, error)
}