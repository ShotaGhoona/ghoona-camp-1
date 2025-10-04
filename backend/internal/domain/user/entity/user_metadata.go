package entity

import (
	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/user/value"
)

// UserMetadata はユーザーの詳細プロフィール情報を表す
type UserMetadata struct {
	common.BaseEntity                // 共通フィールド (ID, CreatedAt, UpdatedAt)
	UserID          common.UUID      // ユーザーID（外部キー）
	DisplayName     *string          // 表示名（オプショナル）
	ProfileImageURL *string          // プロフィール画像URL（オプショナル）
	Tagline         *string          // 一言プロフィール（オプショナル）
	Bio             *string          // 自己紹介（オプショナル）
	Vision          *string          // ビジョン（オプショナル）
	VisionPublic    value.PublicFlag // ビジョンの公開設定
	Timezone        string           // タイムゾーン
	Skills          []string         // スキル一覧
	Interests       []string         // 興味・関心一覧
}

// NewUserMetadata は新しいUserMetadataエンティティを作成する
func NewUserMetadata(userID common.UUID) *UserMetadata {
	return &UserMetadata{
		BaseEntity:   common.NewBaseEntity(),
		UserID:       userID,
		VisionPublic: value.PublicFalse,
		Timezone:     "Asia/Tokyo",
		Skills:       make([]string, 0),
		Interests:    make([]string, 0),
	}
}
