// Package usecase PUT /events/{eventId}/participants/{userId} API用のUseCase
package event

import (
	"context"
	"time"

	"ghoona-camp-backend/internal/application/dto/event"
	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/domain/common"
	domainEvent "ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/repository"
	"ghoona-camp-backend/internal/domain/event/value"
)

// PutEventParticipantUseCase イベント参加ステータス更新
type PutEventParticipantUseCase interface {
	Execute(ctx context.Context, eventID common.UUID, participantUserID, authUserID string, req *dto.PutEventParticipantRequest) (*dto.PutEventParticipantData, error)
}

type putEventParticipantUseCase struct {
	eventRepo       repository.EventRepository
	participantRepo repository.EventParticipantRepository
	txManager       transaction.Manager
}

// NewPutEventParticipantUseCase コンストラクタ
func NewPutEventParticipantUseCase(
	eventRepo repository.EventRepository,
	participantRepo repository.EventParticipantRepository,
	txManager transaction.Manager,
) PutEventParticipantUseCase {
	return &putEventParticipantUseCase{
		eventRepo:       eventRepo,
		participantRepo: participantRepo,
		txManager:       txManager,
	}
}

// Execute イベント参加ステータス更新を実行
func (u *putEventParticipantUseCase) Execute(ctx context.Context, eventID common.UUID, participantUserID, authUserID string, req *dto.PutEventParticipantRequest) (*dto.PutEventParticipantData, error) {
	var response *dto.PutEventParticipantData

	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// 権限チェック
		if participantUserID != authUserID {
			return domainEvent.ErrUnauthorized
		}

		userUUID, err := common.ParseUUID(participantUserID)
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

		// 参加者取得
		participant, err := u.participantRepo.FindByEventAndUser(txCtx, eventID, userUUID)
		if err != nil {
			return err
		}
		if participant == nil {
			return domainEvent.ErrParticipantNotFound
		}

		newStatus := value.ParticipantStatus(req.Status)

		// 同じステータスの場合は何もしない
		if participant.Status == newStatus {
			response = &dto.PutEventParticipantData{
				ID:        participant.ID.String(),
				EventID:   participant.EventID.String(),
				UserID:    participant.UserID.String(),
				Status:    string(participant.Status),
				UpdatedAt: participant.UpdatedAt.Format(time.RFC3339),
			}
			return nil
		}

		// cancelled → registered の場合は定員チェック
		if participant.Status == value.ParticipantStatusCancelled && newStatus == value.ParticipantStatusRegistered {
			if event.MaxParticipants > 0 {
				currentRegistered, err := u.participantRepo.CountRegisteredParticipants(txCtx, eventID)
				if err != nil {
					return err
				}
				if currentRegistered >= event.MaxParticipants {
					return domainEvent.ErrEventFull
				}
			}
		}

		// ステータス更新
		participant.Status = newStatus
		participant.UpdatedAt = time.Now()
		err = u.participantRepo.Update(txCtx, participant)
		if err != nil {
			return err
		}

		// レスポンス作成
		response = &dto.PutEventParticipantData{
			ID:        participant.ID.String(),
			EventID:   participant.EventID.String(),
			UserID:    participant.UserID.String(),
			Status:    string(participant.Status),
			UpdatedAt: participant.UpdatedAt.Format(time.RFC3339),
		}

		return nil
	})

	return response, err
}