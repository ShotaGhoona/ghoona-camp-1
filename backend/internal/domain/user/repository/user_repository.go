package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
)

// UserRepository はユーザーデータアクセスのインターフェースを定義する
type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)         // IDでユーザーを取得
	GetByClerkID(ctx context.Context, clerkID string) (*entity.User, error)  // Clerk IDでユーザーを取得
	GetByEmail(ctx context.Context, email string) (*entity.User, error)      // メールアドレスでユーザーを取得
	Create(ctx context.Context, user *entity.User) error                     // ユーザーを作成
	Update(ctx context.Context, user *entity.User) error                     // ユーザーを更新
	Delete(ctx context.Context, id uuid.UUID) error                          // ユーザーを削除
}