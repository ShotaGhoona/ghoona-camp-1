// Package repository イベントドメインのリポジトリ実装
// GORMを使用してドメインリポジトリインターフェースを実装
package repository

import (
	"context"
	"fmt"

	"ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/repository"
	"ghoona-camp-backend/internal/domain/event/value"
	"ghoona-camp-backend/internal/infrastructure/gorm/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// eventRepository EventRepositoryの実装
type eventRepository struct {
	db *gorm.DB
}

// NewEventRepository EventRepositoryの新しいインスタンスを作成する
func NewEventRepository(db *gorm.DB) repository.EventRepository {
	return &eventRepository{db: db}
}

// FindByID IDでイベントを取得する
func (r *eventRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Event, error) {
	var eventModel model.Event
	if err := r.db.WithContext(ctx).First(&eventModel, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, event.ErrEventNotFound
		}
		return nil, fmt.Errorf("failed to find event by id: %w", err)
	}
	return eventModel.ToEntity(), nil
}

// FindAll フィルタ条件でイベント一覧を取得する
func (r *eventRepository) FindAll(ctx context.Context, filter repository.EventFilter) ([]*entity.Event, error) {
	var eventModels []model.Event
	query := r.db.WithContext(ctx).Where("is_active = ?", true)

	// フィルタ条件を適用
	if filter.DateFrom != nil {
		query = query.Where("scheduled_date >= ?", *filter.DateFrom)
	}
	if filter.DateTo != nil {
		query = query.Where("scheduled_date <= ?", *filter.DateTo)
	}
	if filter.EventType != nil {
		query = query.Where("event_type = ?", string(*filter.EventType))
	}
	if filter.Search != nil && *filter.Search != "" {
		searchTerm := "%" + *filter.Search + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ?", searchTerm, searchTerm)
	}

	// ページネーション
	offset := (filter.Page - 1) * filter.Limit
	query = query.Offset(offset).Limit(filter.Limit)

	// ソート（開始日時順）
	query = query.Order("scheduled_date ASC, start_time ASC")

	if err := query.Find(&eventModels).Error; err != nil {
		return nil, fmt.Errorf("failed to find events: %w", err)
	}

	events := make([]*entity.Event, len(eventModels))
	for i, eventModel := range eventModels {
		events[i] = eventModel.ToEntity()
	}

	return events, nil
}

// Create イベントを作成する
func (r *eventRepository) Create(ctx context.Context, e *entity.Event) error {
	eventModel := model.FromEventEntity(e)
	if err := r.db.WithContext(ctx).Create(eventModel).Error; err != nil {
		return fmt.Errorf("failed to create event: %w", err)
	}
	return nil
}

// Update イベントを更新する
func (r *eventRepository) Update(ctx context.Context, e *entity.Event) error {
	eventModel := model.FromEventEntity(e)
	if err := r.db.WithContext(ctx).Save(eventModel).Error; err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}
	return nil
}

// Delete イベントを削除する（論理削除）
func (r *eventRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Model(&model.Event{}).Where("id = ?", id).Update("is_active", false)
	if result.Error != nil {
		return fmt.Errorf("failed to delete event: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return event.ErrEventNotFound
	}
	return nil
}

// FindByCreatorID 作成者IDでイベントを取得する
func (r *eventRepository) FindByCreatorID(ctx context.Context, creatorID uuid.UUID) ([]*entity.Event, error) {
	var eventModels []model.Event
	if err := r.db.WithContext(ctx).Where("creator_id = ? AND is_active = ?", creatorID, true).
		Order("scheduled_date DESC").Find(&eventModels).Error; err != nil {
		return nil, fmt.Errorf("failed to find events by creator: %w", err)
	}

	events := make([]*entity.Event, len(eventModels))
	for i, eventModel := range eventModels {
		events[i] = eventModel.ToEntity()
	}

	return events, nil
}

// FindByDateRange 日付範囲でイベントを取得する
func (r *eventRepository) FindByDateRange(ctx context.Context, startDate, endDate string) ([]*entity.Event, error) {
	var eventModels []model.Event
	if err := r.db.WithContext(ctx).Where("scheduled_date BETWEEN ? AND ? AND is_active = ?", startDate, endDate, true).
		Order("scheduled_date ASC, start_time ASC").Find(&eventModels).Error; err != nil {
		return nil, fmt.Errorf("failed to find events by date range: %w", err)
	}

	events := make([]*entity.Event, len(eventModels))
	for i, eventModel := range eventModels {
		events[i] = eventModel.ToEntity()
	}

	return events, nil
}

// CountParticipants イベントの参加者数をカウントする
func (r *eventRepository) CountParticipants(ctx context.Context, eventID uuid.UUID) (int, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.EventParticipant{}).
		Where("event_id = ? AND status = ?", eventID, string(value.ParticipantStatusRegistered)).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count participants: %w", err)
	}
	return int(count), nil
}

// ------------------------------------------------------------

// eventParticipantRepository EventParticipantRepositoryの実装
type eventParticipantRepository struct {
	db *gorm.DB
}

// NewEventParticipantRepository EventParticipantRepositoryの新しいインスタンスを作成する
func NewEventParticipantRepository(db *gorm.DB) repository.EventParticipantRepository {
	return &eventParticipantRepository{db: db}
}

// FindByEventID イベントIDで参加者一覧を取得する
func (r *eventParticipantRepository) FindByEventID(ctx context.Context, eventID uuid.UUID) ([]*entity.EventParticipant, error) {
	var participantModels []model.EventParticipant
	if err := r.db.WithContext(ctx).Where("event_id = ?", eventID).
		Order("created_at ASC").Find(&participantModels).Error; err != nil {
		return nil, fmt.Errorf("failed to find participants by event: %w", err)
	}

	participants := make([]*entity.EventParticipant, len(participantModels))
	for i, participantModel := range participantModels {
		participants[i] = participantModel.ToEntity()
	}

	return participants, nil
}

// FindByEventAndUser イベントIDとユーザーIDで参加者を取得する
func (r *eventParticipantRepository) FindByEventAndUser(ctx context.Context, eventID, userID uuid.UUID) (*entity.EventParticipant, error) {
	var participantModel model.EventParticipant
	if err := r.db.WithContext(ctx).Where("event_id = ? AND user_id = ?", eventID, userID).
		First(&participantModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, event.ErrParticipantNotFound
		}
		return nil, fmt.Errorf("failed to find participant: %w", err)
	}
	return participantModel.ToEntity(), nil
}

// Create 参加者を作成する
func (r *eventParticipantRepository) Create(ctx context.Context, participant *entity.EventParticipant) error {
	participantModel := model.FromEventParticipantEntity(participant)
	if err := r.db.WithContext(ctx).Create(participantModel).Error; err != nil {
		return fmt.Errorf("failed to create participant: %w", err)
	}
	return nil
}

// Update 参加者を更新する
func (r *eventParticipantRepository) Update(ctx context.Context, participant *entity.EventParticipant) error {
	participantModel := model.FromEventParticipantEntity(participant)
	if err := r.db.WithContext(ctx).Save(participantModel).Error; err != nil {
		return fmt.Errorf("failed to update participant: %w", err)
	}
	return nil
}

// CountRegisteredParticipants 登録済み参加者数をカウントする
func (r *eventParticipantRepository) CountRegisteredParticipants(ctx context.Context, eventID uuid.UUID) (int, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.EventParticipant{}).
		Where("event_id = ? AND status = ?", eventID, string(value.ParticipantStatusRegistered)).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count registered participants: %w", err)
	}
	return int(count), nil
}

// ExistsByEventAndUser イベントとユーザーの組み合わせで参加者が存在するかチェックする
func (r *eventParticipantRepository) ExistsByEventAndUser(ctx context.Context, eventID, userID uuid.UUID) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.EventParticipant{}).
		Where("event_id = ? AND user_id = ?", eventID, userID).
		Count(&count).Error; err != nil {
		return false, fmt.Errorf("failed to check participant existence: %w", err)
	}
	return count > 0, nil
}
