package entity

import (
	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/user/value"
)

// User はシステム内のユーザーアカウントを表す
type User struct {
	common.BaseEntity          // 共通フィールド (ID, CreatedAt, UpdatedAt)
	ClerkID   string           // Clerk認証ID
	Email     string           // メールアドレス
	Username  *string          // ユーザー名（オプショナル）
	AvatarURL *string          // アバター画像URL（オプショナル）
	DiscordID *string          // Discord ID（オプショナル）
	Status    value.UserStatus // ユーザーステータス
}

// NewUser は新しいUserエンティティを作成する
func NewUser(clerkID, email string) *User {
	return &User{
		BaseEntity: common.NewBaseEntity(),
		ClerkID:    clerkID,
		Email:      email,
		Status:     value.UserStatusActive,
	}
}

// IsActive はユーザーがアクティブ状態かどうかを確認する
func (u *User) IsActive() bool {
	return u.Status.IsActiveState()
}
