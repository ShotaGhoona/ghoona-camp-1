package entity

import (
	"time"

	"github.com/google/uuid"
)

// UserMetadata ユーザーメタデータエンティティ
type UserMetadata struct {
	id          uuid.UUID // メタデータID
	userID      uuid.UUID // ユーザーID
	displayName string    // 表示名
	tagline     string    // 一言プロフィール
	bio         string    // 自己紹介文章
	skills      []string  // スキル
	interests   []string  // 興味・関心
	createdAt   time.Time // 作成日時
	updatedAt   time.Time // 更新日時
}

// NewUserMetadata UserMetadataエンティティを作成する
func NewUserMetadata(
	userID uuid.UUID,
	displayName string,
	tagline string,
	bio string,
	skills []string,
	interests []string,
) *UserMetadata {
	now := time.Now()
	return &UserMetadata{
		id:          uuid.New(),
		userID:      userID,
		displayName: displayName,
		tagline:     tagline,
		bio:         bio,
		skills:      skills,
		interests:   interests,
		createdAt:   now,
		updatedAt:   now,
	}
}

// Getters
func (m *UserMetadata) ID() uuid.UUID          { return m.id }
func (m *UserMetadata) UserID() uuid.UUID      { return m.userID }
func (m *UserMetadata) DisplayName() string    { return m.displayName }
func (m *UserMetadata) Tagline() string        { return m.tagline }
func (m *UserMetadata) Bio() string            { return m.bio }
func (m *UserMetadata) Skills() []string       { return m.skills }
func (m *UserMetadata) Interests() []string    { return m.interests }
func (m *UserMetadata) CreatedAt() time.Time   { return m.createdAt }
func (m *UserMetadata) UpdatedAt() time.Time   { return m.updatedAt }