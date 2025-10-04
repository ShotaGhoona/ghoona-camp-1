package entity

import (
	"time"
	
	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title/value"
)

// TitleAchievement はユーザーが獲得した称号を表す
type TitleAchievement struct {
	common.BaseEntity            // 共通フィールド (ID, CreatedAt, UpdatedAt)
	UserID     common.UUID       // ユーザーID（外部キー）
	TitleID    common.UUID       // 称号ID（外部キー）
	AchievedAt time.Time         // 獲得日時
	IsCurrent  value.CurrentFlag // 現在設定中の称号かどうか
}

// NewTitleAchievement は新しいTitleAchievementエンティティを作成する
func NewTitleAchievement(userID, titleID common.UUID) *TitleAchievement {
	return &TitleAchievement{
		BaseEntity: common.NewBaseEntity(),
		UserID:     userID,
		TitleID:    titleID,
		AchievedAt: time.Now(),
		IsCurrent:  value.CurrentFalse,
	}
}

// SetAsCurrent はこの称号を現在表示中に設定する
func (ta *TitleAchievement) SetAsCurrent() {
	ta.IsCurrent = value.CurrentTrue
	ta.UpdateTimestamp()
}

// UnsetCurrent はこの称号の現在表示を解除する
func (ta *TitleAchievement) UnsetCurrent() {
	ta.IsCurrent = value.CurrentFalse
	ta.UpdateTimestamp()
}

// IsCurrentTitle は現在表示中の称号かどうかを確認する
func (ta *TitleAchievement) IsCurrentTitle() bool {
	return ta.IsCurrent.Bool()
}