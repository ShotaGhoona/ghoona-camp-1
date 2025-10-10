package entity

import (
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/vo"
)

// UserVision ユーザービジョンエンティティ
type UserVision struct {
	id        uuid.UUID     // ビジョンID
	userID    uuid.UUID     // ユーザーID
	vision    string        // ビジョン・将来の目標
	isPublic  vo.PublicFlag // ビジョンの公開設定
	createdAt time.Time     // 作成日時
	updatedAt time.Time     // 更新日時
}

// NewUserVision UserVisionエンティティを作成する
func NewUserVision(
	userID uuid.UUID,
	vision string,
	isPublic vo.PublicFlag,
) *UserVision {
	now := time.Now()
	return &UserVision{
		id:        uuid.New(),
		userID:    userID,
		vision:    vision,
		isPublic:  isPublic,
		createdAt: now,
		updatedAt: now,
	}
}

// Getters
func (v *UserVision) ID() uuid.UUID        { return v.id }
func (v *UserVision) UserID() uuid.UUID    { return v.userID }
func (v *UserVision) Vision() string       { return v.vision }
func (v *UserVision) IsPublic() vo.PublicFlag { return v.isPublic }
func (v *UserVision) CreatedAt() time.Time { return v.createdAt }
func (v *UserVision) UpdatedAt() time.Time { return v.updatedAt }