// Package usecase GET /events/{eventId}/participants API用のUseCase
package event

import (
	"context"
	"time"

	"ghoona-camp-backend/internal/application/dto/event"
	"ghoona-camp-backend/internal/application/dto/event/shared"
	"ghoona-camp-backend/internal/domain/common"
	domainEvent "ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/repository"
	"ghoona-camp-backend/internal/domain/event/value"
)

// GetEventParticipantsUseCase イベント参加者一覧取得
type GetEventParticipantsUseCase interface {
	Execute(ctx context.Context, eventID common.UUID, status *string) ([]dto.GetEventParticipantsDataItem, error)
}

type getEventParticipantsUseCase struct {
	eventRepo       repository.EventRepository
	participantRepo repository.EventParticipantRepository
}

// NewGetEventParticipantsUseCase コンストラクタ
func NewGetEventParticipantsUseCase(
	eventRepo repository.EventRepository,
	participantRepo repository.EventParticipantRepository,
) GetEventParticipantsUseCase {
	return &getEventParticipantsUseCase{
		eventRepo:       eventRepo,
		participantRepo: participantRepo,
	}
}

// Execute イベント参加者一覧取得を実行
func (u *getEventParticipantsUseCase) Execute(ctx context.Context, eventID common.UUID, status *string) ([]dto.GetEventParticipantsDataItem, error) {
	// イベント存在確認
	event, err := u.eventRepo.FindByID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	if event == nil || !event.IsActive {
		return nil, domainEvent.ErrEventNotFound
	}

	// 参加者取得
	participants, err := u.participantRepo.FindByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	// ステータスフィルタ適用
	targetStatus := value.ParticipantStatusRegistered // デフォルト
	if status != nil && *status != "" {
		targetStatus = value.ParticipantStatus(*status)
	}

	var filteredParticipants []*entity.EventParticipant
	for _, p := range participants {
		if p.Status == targetStatus {
			filteredParticipants = append(filteredParticipants, p)
		}
	}

	// DTO変換
	result := make([]dto.GetEventParticipantsDataItem, len(filteredParticipants))
	for i, participant := range filteredParticipants {
		result[i] = dto.GetEventParticipantsDataItem{
			ID: participant.ID.String(),
			User: shared.UserDTO{
				ID:          participant.UserID.String(),
				DisplayName: "User " + participant.UserID.String()[:8],
				Username:    "user_" + participant.UserID.String()[:8],
				AvatarURL:   "https://img.clerk.com/preview.png",
			},
			Status:    string(participant.Status),
			CreatedAt: participant.CreatedAt.Format(time.RFC3339),
			UpdatedAt: participant.UpdatedAt.Format(time.RFC3339),
		}
	}

	return result, nil
}