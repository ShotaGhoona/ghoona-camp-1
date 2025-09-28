package entity

import (
	"time"

	"github.com/google/uuid"
)

// UserRival はユーザー間のライバル関係を表す
type UserRival struct {
	ID          uuid.UUID // 関係ID
	UserID      uuid.UUID // ユーザーID（外部キー）
	RivalUserID uuid.UUID // ライバルのユーザーID（外部キー）
	CreatedAt   time.Time // 作成日時
	UpdatedAt   time.Time // 更新日時
}

// NewUserRival は新しいUserRivalエンティティを作成する
func NewUserRival(userID, rivalUserID uuid.UUID) *UserRival {
	return &UserRival{
		ID:          uuid.New(),
		UserID:      userID,
		RivalUserID: rivalUserID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
