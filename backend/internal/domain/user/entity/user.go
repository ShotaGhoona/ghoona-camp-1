package entity

import (
	"time"

	"ghoona-camp-backend/internal/domain/user/value"

	"github.com/google/uuid"
)

// User はシステム内のユーザーアカウントを表す
type User struct {
	ID        uuid.UUID        // ユーザーID
	ClerkID   string           // Clerk認証ID
	Email     string           // メールアドレス
	Username  *string          // ユーザー名（オプショナル）
	AvatarURL *string          // アバター画像URL（オプショナル）
	DiscordID *string          // Discord ID（オプショナル）
	Status    value.UserStatus // ユーザーステータス
	CreatedAt time.Time        // 作成日時
	UpdatedAt time.Time        // 更新日時
}

// NewUser は新しいUserエンティティを作成する
func NewUser(clerkID, email string) *User {
	return &User{
		ID:        uuid.New(),
		ClerkID:   clerkID,
		Email:     email,
		Status:    value.UserStatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// IsActive はユーザーがアクティブ状態かどうかを確認する
func (u *User) IsActive() bool {
	return u.Status.IsActiveState()
}
