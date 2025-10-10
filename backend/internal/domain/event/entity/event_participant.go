package entity

import (
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/event/vo"
)

// EventParticipant イベント参加者エンティティ
type EventParticipant struct {
	id        uuid.UUID             // 参加ID
	eventID   uuid.UUID             // イベントID
	userID    uuid.UUID             // 参加者ID
	status    vo.ParticipantStatus  // 参加状態
	createdAt time.Time             // 登録日時
	updatedAt time.Time             // 更新日時
}

// NewEventParticipant EventParticipantエンティティを作成する
func NewEventParticipant(
	eventID uuid.UUID,
	userID uuid.UUID,
	status vo.ParticipantStatus,
) *EventParticipant {
	now := time.Now()
	return &EventParticipant{
		id:        uuid.New(),
		eventID:   eventID,
		userID:    userID,
		status:    status,
		createdAt: now,
		updatedAt: now,
	}
}

// Getters
func (p *EventParticipant) ID() uuid.UUID                    { return p.id }
func (p *EventParticipant) EventID() uuid.UUID               { return p.eventID }
func (p *EventParticipant) UserID() uuid.UUID                { return p.userID }
func (p *EventParticipant) Status() vo.ParticipantStatus     { return p.status }
func (p *EventParticipant) CreatedAt() time.Time             { return p.createdAt }
func (p *EventParticipant) UpdatedAt() time.Time             { return p.updatedAt }