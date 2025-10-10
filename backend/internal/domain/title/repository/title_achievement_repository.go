package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/title/entity"
)

// TitleAchievementRepository 称号実績リポジトリインターフェース
type TitleAchievementRepository interface {
	// Create 称号実績を作成する
	Create(ctx context.Context, achievement *entity.TitleAchievement) error

	// GetByID IDで実績を取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.TitleAchievement, error)

	// GetByUserID ユーザーIDで実績一覧を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.TitleAchievement, error)

	// GetCurrentByUserID ユーザーの現在設定中の称号を取得する
	GetCurrentByUserID(ctx context.Context, userID uuid.UUID) (*entity.TitleAchievement, error)

	// GetByUserAndTitle ユーザーIDと称号IDで実績を取得する
	GetByUserAndTitle(ctx context.Context, userID, titleID uuid.UUID) (*entity.TitleAchievement, error)

	// Update 実績を更新する
	Update(ctx context.Context, achievement *entity.TitleAchievement) error

	// SetCurrent 指定した実績を現在の称号に設定する
	SetCurrent(ctx context.Context, userID, achievementID uuid.UUID) error

	// Delete 実績を削除する
	Delete(ctx context.Context, id uuid.UUID) error
}