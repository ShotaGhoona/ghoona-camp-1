// Package usecase POST /events/{eventId}/participants API用のUseCase
package event

import (
	"context"
	"time"

	"ghoona-camp-backend/internal/application/dto/event"
	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/domain/common"
	domainEvent "ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/repository"
	"ghoona-camp-backend/internal/domain/event/value"
)

// PostEventParticipantsUseCase イベント参加申込
type PostEventParticipantsUseCase interface {
	Execute(ctx context.Context, eventID common.UUID, userID string) (*dto.PostEventParticipantsData, error)
}

type postEventParticipantsUseCase struct {
	eventRepo       repository.EventRepository
	participantRepo repository.EventParticipantRepository
	txManager       transaction.Manager
}

// NewPostEventParticipantsUseCase コンストラクタ
func NewPostEventParticipantsUseCase(
	eventRepo repository.EventRepository,
	participantRepo repository.EventParticipantRepository,
	txManager transaction.Manager,
) PostEventParticipantsUseCase {
	return &postEventParticipantsUseCase{
		eventRepo:       eventRepo,
		participantRepo: participantRepo,
		txManager:       txManager,
	}
}

// Execute イベント参加申込を実行
func (u *postEventParticipantsUseCase) Execute(ctx context.Context, eventID common.UUID, userID string) (*dto.PostEventParticipantsData, error) {
	var response *dto.PostEventParticipantsData

	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		userUUID, err := common.ParseUUID(userID)
		if err != nil {
			return err
		}

		// イベント存在確認
		event, err := u.eventRepo.FindByID(txCtx, eventID)
		if err != nil {
			return err
		}
		if event == nil || !event.IsActive {
			return domainEvent.ErrEventNotFound
		}

		// 重複チェック
		existingParticipant, err := u.participantRepo.FindByEventAndUser(txCtx, eventID, userUUID)
		if err != nil {
			return err
		}
		if existingParticipant != nil && existingParticipant.Status == value.ParticipantStatusRegistered {
			return domainEvent.ErrParticipantAlreadyExists
		}

		// 定員チェック
		if event.MaxParticipants > 0 {
			currentRegistered, err := u.participantRepo.CountRegisteredParticipants(txCtx, eventID)
			if err != nil {
				return err
			}
			if currentRegistered >= event.MaxParticipants {
				return domainEvent.ErrEventFull
			}
		}

		// 参加者作成
		var participant *entity.EventParticipant
		if existingParticipant != nil && existingParticipant.Status == value.ParticipantStatusCancelled {
			// 再登録の場合はステータス更新
			existingParticipant.Status = value.ParticipantStatusRegistered
			existingParticipant.UpdatedAt = time.Now()
			err = u.participantRepo.Update(txCtx, existingParticipant)
			if err != nil {
				return err
			}
			participant = existingParticipant
		} else {
			// 新規登録
			participant = entity.NewEventParticipant(eventID, userUUID)
			err = u.participantRepo.Create(txCtx, participant)
			if err != nil {
				return err
			}
		}

		// レスポンス作成
		response = &dto.PostEventParticipantsData{
			ID:        participant.ID.String(),
			EventID:   participant.EventID.String(),
			UserID:    participant.UserID.String(),
			Status:    string(participant.Status),
			CreatedAt: participant.CreatedAt.Format(time.RFC3339),
		}

		return nil
	})

	return response, err
}