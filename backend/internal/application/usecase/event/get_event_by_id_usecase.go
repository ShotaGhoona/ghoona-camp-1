// Package usecase GET /events/{eventId} API用のUseCase
package event

import (
	"context"
	"time"

	"ghoona-camp-backend/internal/application/dto/event"
	"ghoona-camp-backend/internal/application/dto/event/shared"
	"ghoona-camp-backend/internal/domain/common"
	domainEvent "ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/repository"
)

// GetEventByIDUseCase イベント詳細取得
type GetEventByIDUseCase interface {
	Execute(ctx context.Context, eventID common.UUID, userID string) (*dto.GetEventByIDData, error)
}

type getEventByIDUseCase struct {
	eventRepo       repository.EventRepository
	participantRepo repository.EventParticipantRepository
}

// NewGetEventByIDUseCase コンストラクタ
func NewGetEventByIDUseCase(
	eventRepo repository.EventRepository,
	participantRepo repository.EventParticipantRepository,
) GetEventByIDUseCase {
	return &getEventByIDUseCase{
		eventRepo:       eventRepo,
		participantRepo: participantRepo,
	}
}

// Execute イベント詳細取得を実行
func (u *getEventByIDUseCase) Execute(ctx context.Context, eventID common.UUID, userID string) (*dto.GetEventByIDData, error) {
	// イベント取得
	event, err := u.eventRepo.FindByID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	if event == nil || !event.IsActive {
		return nil, domainEvent.ErrEventNotFound
	}

	// 参加者数取得
	participantCount, _ := u.participantRepo.CountRegisteredParticipants(ctx, event.ID)

	// ユーザー登録状況確認
	isUserRegistered := checkUserRegistration(ctx, u.participantRepo, event.ID, userID)

	// レスポンス作成
	return &dto.GetEventByIDData{
		ID: event.ID.String(),
		Creator: shared.CreatorDTO{
			ID:          event.CreatorID.String(),
			DisplayName: "User " + event.CreatorID.String()[:8],
			Username:    "user_" + event.CreatorID.String()[:8],
			AvatarURL:   "https://img.clerk.com/preview.png",
		},
		Title:             event.Title,
		Description:       event.Description,
		EventType:         string(event.EventType),
		ScheduledDate:     event.ScheduledDate.Format("2006-01-02"),
		StartTime:         event.StartTime.Format("15:04:05"),
		EndTime:           event.EndTime.Format("15:04:05"),
		MaxParticipants:   event.MaxParticipants,
		IsRecurring:       event.IsRecurring,
		RecurrencePattern: getRecurrencePattern(event.RecurrencePattern),
		DiscordChannelID:  getDiscordChannel(event.DiscordChannelID),
		ParticipantCount:  participantCount,
		IsUserRegistered:  isUserRegistered,
		IsActive:          event.IsActive,
		CreatedAt:         event.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         event.UpdatedAt.Format(time.RFC3339),
	}, nil
}