package repository

import (
	"context"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title/entity"
)

// TitleAchievementRepository は称号獲得実績データアクセスのインターフェースを定義する
type TitleAchievementRepository interface {
	GetByUserID(ctx context.Context, userID common.UUID) ([]*entity.TitleAchievement, error)        // ユーザーIDで称号獲得実績一覧を取得
	GetCurrentByUserID(ctx context.Context, userID common.UUID) (*entity.TitleAchievement, error)   // ユーザーの現在表示中の称号を取得
	GetByUserIDAndTitleID(ctx context.Context, userID, titleID common.UUID) (*entity.TitleAchievement, error) // ユーザーIDと称号IDで実績を取得
	Create(ctx context.Context, achievement *entity.TitleAchievement) error                        // 称号獲得実績を作成
	Update(ctx context.Context, achievement *entity.TitleAchievement) error                        // 称号獲得実績を更新
	SetCurrent(ctx context.Context, userID, titleID common.UUID) error                               // 指定称号を現在表示中に設定（他の称号は解除）
}