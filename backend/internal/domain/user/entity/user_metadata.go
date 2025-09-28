package entity

import (
	"time"

	"ghoona-camp-backend/internal/domain/user/value"

	"github.com/google/uuid"
)

// UserMetadata はユーザーの詳細プロフィール情報を表す
type UserMetadata struct {
	ID              uuid.UUID        // メタデータID
	UserID          uuid.UUID        // ユーザーID（外部キー）
	DisplayName     *string          // 表示名（オプショナル）
	ProfileImageURL *string          // プロフィール画像URL（オプショナル）
	Tagline         *string          // 一言プロフィール（オプショナル）
	Bio             *string          // 自己紹介（オプショナル）
	Vision          *string          // ビジョン（オプショナル）
	VisionPublic    value.PublicFlag // ビジョンの公開設定
	Timezone        string           // タイムゾーン
	Skills          []string         // スキル一覧
	Interests       []string         // 興味・関心一覧
	CreatedAt       time.Time        // 作成日時
	UpdatedAt       time.Time        // 更新日時
}

// NewUserMetadata は新しいUserMetadataエンティティを作成する
func NewUserMetadata(userID uuid.UUID) *UserMetadata {
	return &UserMetadata{
		ID:           uuid.New(),
		UserID:       userID,
		VisionPublic: value.PublicFalse,
		Timezone:     "Asia/Tokyo",
		Skills:       make([]string, 0),
		Interests:    make([]string, 0),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
