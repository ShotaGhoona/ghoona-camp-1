package entity

import (
	"errors"
	"net/url"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/vo"
)

// UserSocialLink ユーザーソーシャルリンクエンティティ
type UserSocialLink struct {
	id        uuid.UUID      // リンクID
	userID    uuid.UUID      // ユーザーID
	platform  vo.Platform    // プラットフォーム
	url       string         // リンクURL
	title     string         // リンクのタイトル・説明
	isPublic  vo.PublicFlag  // 公開設定
	createdAt time.Time      // 作成日時
	updatedAt time.Time      // 更新日時
}

// NewUserSocialLink UserSocialLinkエンティティを作成する
func NewUserSocialLink(
	userID uuid.UUID,
	platform vo.Platform,
	linkURL string,
	title string,
	isPublic vo.PublicFlag,
) (*UserSocialLink, error) {
	// URL妥当性チェック
	if linkURL == "" {
		return nil, errors.New("URLは必須です")
	}
	
	// URL形式チェック
	if _, err := url.ParseRequestURI(linkURL); err != nil {
		return nil, errors.New("不正なURL形式です")
	}

	now := time.Now()
	return &UserSocialLink{
		id:        uuid.New(),
		userID:    userID,
		platform:  platform,
		url:       linkURL,
		title:     title,
		isPublic:  isPublic,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// Getters
func (l *UserSocialLink) ID() uuid.UUID       { return l.id }
func (l *UserSocialLink) UserID() uuid.UUID   { return l.userID }
func (l *UserSocialLink) Platform() vo.Platform { return l.platform }
func (l *UserSocialLink) URL() string         { return l.url }
func (l *UserSocialLink) Title() string       { return l.title }
func (l *UserSocialLink) IsPublic() vo.PublicFlag { return l.isPublic }
func (l *UserSocialLink) CreatedAt() time.Time { return l.createdAt }
func (l *UserSocialLink) UpdatedAt() time.Time { return l.updatedAt }