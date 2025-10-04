package entity

import (
	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/user/value"
)

// UserSocialLink はユーザーのソーシャルメディアリンクを表す
type UserSocialLink struct {
	common.BaseEntity                // 共通フィールド (ID, CreatedAt, UpdatedAt)
	UserID    common.UUID            // ユーザーID（外部キー）
	Platform  value.Platform         // プラットフォーム
	URL       string                 // リンクURL
	Title     *string                // リンクタイトル（オプショナル）
	IsPublic  value.PublicFlag       // 公開設定
}

// NewUserSocialLink は新しいUserSocialLinkエンティティを作成する
func NewUserSocialLink(userID common.UUID, platform value.Platform, url string) *UserSocialLink {
	return &UserSocialLink{
		BaseEntity: common.NewBaseEntity(),
		UserID:     userID,
		Platform:   platform,
		URL:        url,
		IsPublic:   value.PublicTrue,
	}
}
