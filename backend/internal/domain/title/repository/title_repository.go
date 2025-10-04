package repository

import (
	"context"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title/entity"
)

// TitleRepository は称号データアクセスのインターフェースを定義する
type TitleRepository interface {
	GetAll(ctx context.Context) ([]*entity.Title, error)               // 全称号を取得
	GetByID(ctx context.Context, id common.UUID) (*entity.Title, error) // IDで称号を取得
	GetByLevel(ctx context.Context, level int) (*entity.Title, error) // レベルで称号を取得
	GetActiveTitles(ctx context.Context) ([]*entity.Title, error)      // アクティブな称号一覧を取得
}