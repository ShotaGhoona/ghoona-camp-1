package event

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/repository"
	eventModel "ghoona-camp-backend/internal/infrastructure/gorm/model/event"
)

type eventParticipantRepository struct {
	db *gorm.DB
}

// NewEventParticipantRepository コンストラクタ
func NewEventParticipantRepository(db *gorm.DB) repository.EventParticipantRepository {
	return &eventParticipantRepository{db: db}
}

// Create イベント参加者を作成する
func (r *eventParticipantRepository) Create(ctx context.Context, participant *entity.EventParticipant) error {
	gormParticipant := r.toGORMEventParticipant(participant)
	return r.db.WithContext(ctx).Create(gormParticipant).Error
}

// GetByID IDで参加者を取得する
func (r *eventParticipantRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.EventParticipant, error) {
	var gormParticipant eventModel.EventParticipant
	err := r.db.WithContext(ctx).First(&gormParticipant, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMEventParticipant(&gormParticipant)
}

// GetByEventID イベントIDで参加者一覧を取得する
func (r *eventParticipantRepository) GetByEventID(ctx context.Context, eventID uuid.UUID) ([]*entity.EventParticipant, error) {
	var gormParticipants []eventModel.EventParticipant
	err := r.db.WithContext(ctx).Where("event_id = ?", eventID).Find(&gormParticipants).Error
	if err != nil {
		return nil, err
	}

	participants := make([]*entity.EventParticipant, len(gormParticipants))
	for i, gormParticipant := range gormParticipants {
		participant, err := r.fromGORMEventParticipant(&gormParticipant)
		if err != nil {
			return nil, err
		}
		participants[i] = participant
	}
	return participants, nil
}

// GetByUserID ユーザーIDで参加イベント一覧を取得する
func (r *eventParticipantRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.EventParticipant, error) {
	var gormParticipants []eventModel.EventParticipant
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&gormParticipants).Error
	if err != nil {
		return nil, err
	}

	participants := make([]*entity.EventParticipant, len(gormParticipants))
	for i, gormParticipant := range gormParticipants {
		participant, err := r.fromGORMEventParticipant(&gormParticipant)
		if err != nil {
			return nil, err
		}
		participants[i] = participant
	}
	return participants, nil
}

// GetByEventAndUser イベントIDとユーザーIDで参加者を取得する
func (r *eventParticipantRepository) GetByEventAndUser(ctx context.Context, eventID, userID uuid.UUID) (*entity.EventParticipant, error) {
	var gormParticipant eventModel.EventParticipant
	err := r.db.WithContext(ctx).Where("event_id = ? AND user_id = ?", eventID, userID).First(&gormParticipant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMEventParticipant(&gormParticipant)
}

// CountByEventID イベントの参加者数を取得する
func (r *eventParticipantRepository) CountByEventID(ctx context.Context, eventID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&eventModel.EventParticipant{}).Where("event_id = ?", eventID).Count(&count).Error
	return int(count), err
}

// Update イベント参加者を更新する
func (r *eventParticipantRepository) Update(ctx context.Context, participant *entity.EventParticipant) error {
	gormParticipant := r.toGORMEventParticipant(participant)
	return r.db.WithContext(ctx).Save(gormParticipant).Error
}

// Delete イベント参加者を削除する
func (r *eventParticipantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&eventModel.EventParticipant{}, "id = ?", id).Error
}

// DeleteByEventAndUser イベントIDとユーザーIDで参加者を削除する
func (r *eventParticipantRepository) DeleteByEventAndUser(ctx context.Context, eventID, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&eventModel.EventParticipant{}, "event_id = ? AND user_id = ?", eventID, userID).Error
}

// 型変換: Domain Entity → GORM Model
func (r *eventParticipantRepository) toGORMEventParticipant(participant *entity.EventParticipant) *eventModel.EventParticipant {
	return &eventModel.EventParticipant{
		ID:        participant.ID(),
		EventID:   participant.EventID(),
		UserID:    participant.UserID(),
		Status:    participant.Status(),
		JoinedAt:  participant.JoinedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *eventParticipantRepository) fromGORMEventParticipant(gormParticipant *eventModel.EventParticipant) (*entity.EventParticipant, error) {
	// TODO: entity.NewEventParticipant の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}