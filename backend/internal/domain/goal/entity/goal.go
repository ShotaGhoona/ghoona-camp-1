package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/goal/vo"
)

// Goal 目標エンティティ
type Goal struct {
	id          uuid.UUID       // 目標ID
	userID      uuid.UUID       // ユーザーID
	title       string          // 目標タイトル
	description string          // 目標詳細
	startedAt   time.Time       // 目標開始日
	endedAt     *time.Time      // 目標終了日
	isActive    vo.ActiveFlag   // 目標の有効状態
	isPublic    vo.PublicFlag   // 外部公開設定
	createdAt   time.Time       // 作成日時
	updatedAt   time.Time       // 更新日時
}

// NewGoal Goalエンティティを作成する
func NewGoal(
	userID uuid.UUID,
	title string,
	description string,
	startedAt time.Time,
	endedAt *time.Time,
	isActive vo.ActiveFlag,
	isPublic vo.PublicFlag,
) (*Goal, error) {
	// バリデーション
	if title == "" {
		return nil, errors.New("目標タイトルは必須です")
	}
	
	// 期間妥当性チェック (ended_at > started_at)
	if endedAt != nil && endedAt.Before(startedAt) {
		return nil, errors.New("終了日は開始日より後に設定してください")
	}

	now := time.Now()
	return &Goal{
		id:          uuid.New(),
		userID:      userID,
		title:       title,
		description: description,
		startedAt:   startedAt,
		endedAt:     endedAt,
		isActive:    isActive,
		isPublic:    isPublic,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// Getters
func (g *Goal) ID() uuid.UUID            { return g.id }
func (g *Goal) UserID() uuid.UUID        { return g.userID }
func (g *Goal) Title() string            { return g.title }
func (g *Goal) Description() string      { return g.description }
func (g *Goal) StartedAt() time.Time     { return g.startedAt }
func (g *Goal) EndedAt() *time.Time      { return g.endedAt }
func (g *Goal) IsActive() vo.ActiveFlag  { return g.isActive }
func (g *Goal) IsPublic() vo.PublicFlag  { return g.isPublic }
func (g *Goal) CreatedAt() time.Time     { return g.createdAt }
func (g *Goal) UpdatedAt() time.Time     { return g.updatedAt }