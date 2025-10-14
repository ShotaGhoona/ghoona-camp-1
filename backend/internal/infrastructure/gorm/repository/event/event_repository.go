package event

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/repository"
	"ghoona-camp-backend/internal/domain/event/vo"
	eventModel "ghoona-camp-backend/internal/infrastructure/gorm/model/event"
)

type eventRepository struct {
	db *gorm.DB
}

// NewEventRepository コンストラクタ
func NewEventRepository(db *gorm.DB) repository.EventRepository {
	return &eventRepository{db: db}
}

// Create イベントを作成する
func (r *eventRepository) Create(ctx context.Context, event *entity.Event) error {
	gormEvent := r.toGORMEvent(event)
	return r.db.WithContext(ctx).Create(gormEvent).Error
}

// GetByID IDでイベントを取得する
func (r *eventRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Event, error) {
	var gormEvent eventModel.Event
	err := r.db.WithContext(ctx).First(&gormEvent, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMEvent(&gormEvent)
}

// List イベント一覧を取得する
func (r *eventRepository) List(ctx context.Context, limit, offset int) ([]*entity.Event, error) {
	var gormEvents []eventModel.Event
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&gormEvents).Error
	if err != nil {
		return nil, err
	}

	events := make([]*entity.Event, len(gormEvents))
	for i, gormEvent := range gormEvents {
		event, err := r.fromGORMEvent(&gormEvent)
		if err != nil {
			return nil, err
		}
		events[i] = event
	}
	return events, nil
}

// GetByDateRange 日付範囲でイベントを取得する
func (r *eventRepository) GetByDateRange(ctx context.Context, dateFrom, dateTo time.Time, limit, offset int) ([]*entity.Event, error) {
	var gormEvents []eventModel.Event
	err := r.db.WithContext(ctx).Where("start_time >= ? AND start_time <= ?", dateFrom, dateTo).Limit(limit).Offset(offset).Find(&gormEvents).Error
	if err != nil {
		return nil, err
	}

	events := make([]*entity.Event, len(gormEvents))
	for i, gormEvent := range gormEvents {
		event, err := r.fromGORMEvent(&gormEvent)
		if err != nil {
			return nil, err
		}
		events[i] = event
	}
	return events, nil
}

// GetByStatus ステータスでイベントを取得する
func (r *eventRepository) GetByStatus(ctx context.Context, status vo.EventStatus, limit, offset int) ([]*entity.Event, error) {
	var gormEvents []eventModel.Event
	err := r.db.WithContext(ctx).Where("status = ?", string(status)).Limit(limit).Offset(offset).Find(&gormEvents).Error
	if err != nil {
		return nil, err
	}

	events := make([]*entity.Event, len(gormEvents))
	for i, gormEvent := range gormEvents {
		event, err := r.fromGORMEvent(&gormEvent)
		if err != nil {
			return nil, err
		}
		events[i] = event
	}
	return events, nil
}

// GetByCreatorID 作成者IDでイベントを取得する
func (r *eventRepository) GetByCreatorID(ctx context.Context, creatorID uuid.UUID, limit, offset int) ([]*entity.Event, error) {
	var gormEvents []eventModel.Event
	err := r.db.WithContext(ctx).Where("creator_id = ?", creatorID).Limit(limit).Offset(offset).Find(&gormEvents).Error
	if err != nil {
		return nil, err
	}

	events := make([]*entity.Event, len(gormEvents))
	for i, gormEvent := range gormEvents {
		event, err := r.fromGORMEvent(&gormEvent)
		if err != nil {
			return nil, err
		}
		events[i] = event
	}
	return events, nil
}

// Update イベントを更新する
func (r *eventRepository) Update(ctx context.Context, event *entity.Event) error {
	gormEvent := r.toGORMEvent(event)
	return r.db.WithContext(ctx).Save(gormEvent).Error
}

// Delete イベントを削除する
func (r *eventRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&eventModel.Event{}, "id = ?", id).Error
}

// Search イベントを検索する
func (r *eventRepository) Search(ctx context.Context, query string, limit, offset int) ([]*entity.Event, error) {
	// TODO: 検索仕様確定後実装（検索対象フィールド：title? description?）
	return nil, errors.New("TODO: 要件確定後実装")
}

// 型変換: Domain Entity → GORM Model
func (r *eventRepository) toGORMEvent(event *entity.Event) *eventModel.Event {
	return &eventModel.Event{
		ID:             event.ID(),
		CreatorID:      event.CreatorID(),
		Title:          event.Title(),
		Description:    event.Description(),
		StartTime:      event.StartTime(),
		EndTime:        event.EndTime(),
		MaxParticipants: event.MaxParticipants(),
		Location:       event.Location(),
		Status:         string(event.Status()),
		ImageURL:       event.ImageURL(),
		Tags:           event.Tags(),
		CreatedAt:      event.CreatedAt(),
		UpdatedAt:      event.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *eventRepository) fromGORMEvent(gormEvent *eventModel.Event) (*entity.Event, error) {
	// TODO: entity.NewEvent の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}