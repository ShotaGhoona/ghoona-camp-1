package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/vo"
)

// User ユーザーエンティティ
type User struct {
	id        uuid.UUID   // ユーザーID
	clerkID   string      // Clerk User ID
	email     string      // メールアドレス
	username  string      // ユーザー名
	avatarURL string      // アバター画像URL
	discordID string      // Discord User ID
	status    vo.UserStatus // アカウント状態
	createdAt time.Time   // 作成日時
	updatedAt time.Time   // 更新日時
}

// NewUser Userエンティティを作成する
func NewUser(
	clerkID string,
	email string,
	username string,
	avatarURL string,
	discordID string,
	status vo.UserStatus,
) (*User, error) {
	// バリデーション
	if clerkID == "" {
		return nil, errors.New("Clerk IDは必須です")
	}
	if email == "" {
		return nil, errors.New("メールアドレスは必須です")
	}

	now := time.Now()
	return &User{
		id:        uuid.New(),
		clerkID:   clerkID,
		email:     email,
		username:  username,
		avatarURL: avatarURL,
		discordID: discordID,
		status:    status,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// Getters
func (u *User) ID() uuid.UUID         { return u.id }
func (u *User) ClerkID() string       { return u.clerkID }
func (u *User) Email() string         { return u.email }
func (u *User) Username() string      { return u.username }
func (u *User) AvatarURL() string     { return u.avatarURL }
func (u *User) DiscordID() string     { return u.discordID }
func (u *User) Status() vo.UserStatus { return u.status }
func (u *User) CreatedAt() time.Time  { return u.createdAt }
func (u *User) UpdatedAt() time.Time  { return u.updatedAt }