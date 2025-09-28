package entity

import (
	"time"

	"ghoona-camp-backend/internal/domain/user/value"

	"github.com/google/uuid"
)

// UserSocialLink はユーザーのソーシャルメディアリンクを表す
type UserSocialLink struct {
	ID        uuid.UUID        // リンクID
	UserID    uuid.UUID        // ユーザーID（外部キー）
	Platform  value.Platform   // プラットフォーム
	URL       string           // リンクURL
	Title     *string          // リンクタイトル（オプショナル）
	IsPublic  value.PublicFlag // 公開設定
	CreatedAt time.Time        // 作成日時
	UpdatedAt time.Time        // 更新日時
}

// NewUserSocialLink は新しいUserSocialLinkエンティティを作成する
func NewUserSocialLink(userID uuid.UUID, platform value.Platform, url string) *UserSocialLink {
	return &UserSocialLink{
		ID:        uuid.New(),
		UserID:    userID,
		Platform:  platform,
		URL:       url,
		IsPublic:  value.PublicTrue,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
