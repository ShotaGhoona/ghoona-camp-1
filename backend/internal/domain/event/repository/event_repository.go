package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/vo"
)

// EventRepository イベントリポジトリインターフェース
type EventRepository interface {
	// Create イベントを作成する
	Create(ctx context.Context, event *entity.Event) error

	// GetByID IDでイベントを取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Event, error)

	// GetAll イベント一覧を取得する
	GetAll(ctx context.Context, limit, offset int) ([]*entity.Event, error)

	// GetByCreatorID 作成者IDでイベント一覧を取得する
	GetByCreatorID(ctx context.Context, creatorID uuid.UUID, limit, offset int) ([]*entity.Event, error)

	// GetByDateRange 日付範囲でイベントを取得する
	GetByDateRange(ctx context.Context, dateFrom, dateTo *time.Time, limit, offset int) ([]*entity.Event, error)

	// GetByEventType イベントタイプでイベントを取得する
	GetByEventType(ctx context.Context, eventType vo.EventType, limit, offset int) ([]*entity.Event, error)

	// Update イベントを更新する
	Update(ctx context.Context, event *entity.Event) error

	// Delete イベントを削除する
	Delete(ctx context.Context, id uuid.UUID) error
}