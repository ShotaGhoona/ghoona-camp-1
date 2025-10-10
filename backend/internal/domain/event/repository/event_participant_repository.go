package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/vo"
)

// EventParticipantRepository イベント参加者リポジトリインターフェース
type EventParticipantRepository interface {
	// Create イベント参加者を作成する
	Create(ctx context.Context, participant *entity.EventParticipant) error

	// GetByID IDで参加者を取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.EventParticipant, error)

	// GetByEventID イベントIDで参加者一覧を取得する
	GetByEventID(ctx context.Context, eventID uuid.UUID) ([]*entity.EventParticipant, error)

	// GetByUserID ユーザーIDで参加イベント一覧を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.EventParticipant, error)

	// GetByEventAndUser イベントIDとユーザーIDで参加者を取得する
	GetByEventAndUser(ctx context.Context, eventID, userID uuid.UUID) (*entity.EventParticipant, error)

	// Update 参加者情報を更新する
	Update(ctx context.Context, participant *entity.EventParticipant) error

	// Delete 参加者を削除する
	Delete(ctx context.Context, id uuid.UUID) error

	// CountByEventID イベントの参加者数を取得する
	CountByEventID(ctx context.Context, eventID uuid.UUID) (int, error)

	// CountByEventIDAndStatus イベントの特定ステータスの参加者数を取得する
	CountByEventIDAndStatus(ctx context.Context, eventID uuid.UUID, status vo.ParticipantStatus) (int, error)
}