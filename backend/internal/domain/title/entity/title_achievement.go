package entity

import (
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/title/vo"
)

// TitleAchievement 称号実績エンティティ
type TitleAchievement struct {
	id         uuid.UUID        // 実績ID
	userID     uuid.UUID        // ユーザーID
	titleID    uuid.UUID        // 称号ID
	achievedAt time.Time        // 獲得日時
	isCurrent  vo.CurrentFlag   // 現在設定中の称号かどうか
	createdAt  time.Time        // 作成日時
	updatedAt  time.Time        // 更新日時
}

// NewTitleAchievement TitleAchievementエンティティを作成する
func NewTitleAchievement(
	userID uuid.UUID,
	titleID uuid.UUID,
	achievedAt time.Time,
	isCurrent vo.CurrentFlag,
) *TitleAchievement {
	now := time.Now()
	return &TitleAchievement{
		id:         uuid.New(),
		userID:     userID,
		titleID:    titleID,
		achievedAt: achievedAt,
		isCurrent:  isCurrent,
		createdAt:  now,
		updatedAt:  now,
	}
}

// Getters
func (a *TitleAchievement) ID() uuid.UUID             { return a.id }
func (a *TitleAchievement) UserID() uuid.UUID         { return a.userID }
func (a *TitleAchievement) TitleID() uuid.UUID        { return a.titleID }
func (a *TitleAchievement) AchievedAt() time.Time     { return a.achievedAt }
func (a *TitleAchievement) IsCurrent() vo.CurrentFlag { return a.isCurrent }
func (a *TitleAchievement) CreatedAt() time.Time      { return a.createdAt }
func (a *TitleAchievement) UpdatedAt() time.Time      { return a.updatedAt }