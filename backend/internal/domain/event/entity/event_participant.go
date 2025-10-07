// Package entity イベントドメインのエンティティを定義する
package entity

import (
	"time"

	"ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/value"

	"github.com/google/uuid"
)

// EventParticipant イベント参加者を表すエンティティ
type EventParticipant struct {
	ID        uuid.UUID
	EventID   uuid.UUID
	UserID    uuid.UUID
	Status    value.ParticipantStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewEventParticipant 新しいイベント参加者を作成するファクトリメソッド
func NewEventParticipant(eventID, userID uuid.UUID) *EventParticipant {
	now := time.Now()
	return &EventParticipant{
		ID:        uuid.New(),
		EventID:   eventID,
		UserID:    userID,
		Status:    value.ParticipantStatusRegistered,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Register 参加登録する（キャンセルからの復帰も含む）
func (ep *EventParticipant) Register() error {
	return ep.UpdateStatus(value.ParticipantStatusRegistered)
}

// Cancel 参加をキャンセルする
func (ep *EventParticipant) Cancel() error {
	return ep.UpdateStatus(value.ParticipantStatusCancelled)
}

// UpdateStatus 参加ステータスを更新する（ゲート一本化）
func (ep *EventParticipant) UpdateStatus(status value.ParticipantStatus) error {
	if !status.IsValid() {
		return event.ErrInvalidParticipantStatus
	}
	ep.Status = status
	ep.UpdatedAt = time.Now()
	return nil
}

// IsRegistered 現在登録状態かどうかを判定する
func (ep *EventParticipant) IsRegistered() bool {
	return ep.Status == value.ParticipantStatusRegistered
}

// IsCancelled 現在キャンセル状態かどうかを判定する
func (ep *EventParticipant) IsCancelled() bool {
	return ep.Status == value.ParticipantStatusCancelled
}
