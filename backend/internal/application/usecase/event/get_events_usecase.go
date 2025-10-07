// Package usecase GET /events API用のUseCase
package event

import (
	"context"
	"strings"

	"ghoona-camp-backend/internal/application/dto/event"
	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/repository"
	"ghoona-camp-backend/internal/domain/event/value"
)

// GetEventsUseCase イベント一覧取得
type GetEventsUseCase interface {
	Execute(ctx context.Context, dateFrom, dateTo, eventType, search *string, page, limit int, userID string) ([]dto.GetEventsDataItem, error)
}

type getEventsUseCase struct {
	eventRepo       repository.EventRepository
	participantRepo repository.EventParticipantRepository
}

// NewGetEventsUseCase コンストラクタ
func NewGetEventsUseCase(
	eventRepo repository.EventRepository,
	participantRepo repository.EventParticipantRepository,
) GetEventsUseCase {
	return &getEventsUseCase{
		eventRepo:       eventRepo,
		participantRepo: participantRepo,
	}
}

// Execute イベント一覧取得を実行
func (u *getEventsUseCase) Execute(ctx context.Context, dateFrom, dateTo, eventType, search *string, page, limit int, userID string) ([]dto.GetEventsDataItem, error) {
	// パラメータ正規化
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 20
	}

	// フィルタ構築
	filter := repository.EventFilter{
		DateFrom:  dateFrom,
		DateTo:    dateTo,
		EventType: parseEventType(eventType),
		Search:    parseSearchTerm(search),
		Page:      page,
		Limit:     limit,
	}

	// リポジトリから取得
	events, err := u.eventRepo.FindAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	// DTO変換
	result := make([]dto.GetEventsDataItem, len(events))
	for i, event := range events {
		participantCount, _ := u.participantRepo.CountRegisteredParticipants(ctx, event.ID)
		isUserRegistered := checkUserRegistration(ctx, u.participantRepo, event.ID, userID)
		result[i] = eventToListItem(event, participantCount, isUserRegistered)
	}

	return result, nil
}

// Helper functions
func parseEventType(eventType *string) *value.EventType {
	if eventType != nil && *eventType != "" {
		et := value.EventType(*eventType)
		return &et
	}
	return nil
}

func parseSearchTerm(search *string) *string {
	if search != nil {
		trimmed := strings.TrimSpace(*search)
		if trimmed != "" {
			return &trimmed
		}
	}
	return nil
}

func eventToListItem(event *entity.Event, participantCount int, isUserRegistered bool) dto.GetEventsDataItem {
	return dto.GetEventsDataItem{
		ID: event.ID.String(),
		Creator: dto.CreatorDTO{
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
		CreatedAt:         event.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:         event.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}