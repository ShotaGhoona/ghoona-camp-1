package entity

import (
	"ghoona-camp-backend/internal/domain/common"
)

// UserRival はユーザー間のライバル関係を表す
type UserRival struct {
	common.BaseEntity               // 共通フィールド (ID, CreatedAt, UpdatedAt)
	UserID      common.UUID         // ユーザーID（外部キー）
	RivalUserID common.UUID         // ライバルのユーザーID（外部キー）
}

// NewUserRival は新しいUserRivalエンティティを作成する
func NewUserRival(userID, rivalUserID common.UUID) *UserRival {
	return &UserRival{
		BaseEntity:  common.NewBaseEntity(),
		UserID:      userID,
		RivalUserID: rivalUserID,
	}
}
