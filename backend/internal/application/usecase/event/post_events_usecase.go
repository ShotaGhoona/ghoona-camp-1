// Package usecase POST /events API用のUseCase
package event

import (
	"context"
	"time"

	"ghoona-camp-backend/internal/application/dto/event"
	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/event/entity"
	domainEvent "ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/repository"
	"ghoona-camp-backend/internal/domain/event/value"
)

// PostEventsUseCase イベント作成
type PostEventsUseCase interface {
	Execute(ctx context.Context, userID string, req *dto.PostEventsRequest) (*dto.PostEventsData, error)
}

type postEventsUseCase struct {
	eventRepo repository.EventRepository
	txManager transaction.Manager
}

// NewPostEventsUseCase コンストラクタ
func NewPostEventsUseCase(
	eventRepo repository.EventRepository,
	txManager transaction.Manager,
) PostEventsUseCase {
	return &postEventsUseCase{
		eventRepo: eventRepo,
		txManager: txManager,
	}
}

// Execute イベント作成を実行
func (u *postEventsUseCase) Execute(ctx context.Context, userID string, req *dto.PostEventsRequest) (*dto.PostEventsData, error) {
	var response *dto.PostEventsData

	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		creatorID, err := common.ParseUUID(userID)
		if err != nil {
			return err
		}

		// 日付・時刻パース
		scheduledDate, err := time.Parse("2006-01-02", req.ScheduledDate)
		if err != nil {
			return err
		}
		startTime, err := time.Parse("15:04:05", req.StartTime)
		if err != nil {
			return err
		}
		endTime, err := time.Parse("15:04:05", req.EndTime)
		if err != nil {
			return err
		}

		// バリデーション
		if !startTime.Before(endTime) {
			return domainEvent.ErrInvalidTimeSlot
		}

		// デフォルト値設定
		maxParticipants := req.MaxParticipants
		if maxParticipants == 0 {
			maxParticipants = 10
		}

		recurrencePattern := value.RecurrencePatternNone
		if req.IsRecurring && req.RecurrencePattern != "" {
			recurrencePattern = value.RecurrencePattern(req.RecurrencePattern)
		}

		// エンティティ作成
		event := entity.NewEvent(
			creatorID,
			req.Title,
			req.Description,
			value.EventType(req.EventType),
			scheduledDate,
			startTime,
			endTime,
			maxParticipants,
			req.IsRecurring,
			recurrencePattern,
			req.DiscordChannelID,
		)

		// 保存
		err = u.eventRepo.Create(txCtx, event)
		if err != nil {
			return err
		}

		// レスポンス作成
		response = &dto.PostEventsData{
			ID:                event.ID.String(),
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
			ParticipantCount:  0,
			IsActive:          event.IsActive,
			CreatedAt:         event.CreatedAt.Format(time.RFC3339),
			UpdatedAt:         event.UpdatedAt.Format(time.RFC3339),
		}

		return nil
	})

	return response, err
}

