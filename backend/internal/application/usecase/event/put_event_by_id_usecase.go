// Package usecase PUT /events/{eventId} API用のUseCase
package event

import (
	"context"
	"time"

	"ghoona-camp-backend/internal/application/dto/event"
	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/domain/common"
	domainEvent "ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/repository"
)

// PutEventByIDUseCase イベント更新
type PutEventByIDUseCase interface {
	Execute(ctx context.Context, eventID common.UUID, userID string, req *dto.PutEventByIDRequest) (*dto.PutEventByIDData, error)
}

type putEventByIDUseCase struct {
	eventRepo       repository.EventRepository
	participantRepo repository.EventParticipantRepository
	txManager       transaction.Manager
}

// NewPutEventByIDUseCase コンストラクタ
func NewPutEventByIDUseCase(
	eventRepo repository.EventRepository,
	participantRepo repository.EventParticipantRepository,
	txManager transaction.Manager,
) PutEventByIDUseCase {
	return &putEventByIDUseCase{
		eventRepo:       eventRepo,
		participantRepo: participantRepo,
		txManager:       txManager,
	}
}

// Execute イベント更新を実行
func (u *putEventByIDUseCase) Execute(ctx context.Context, eventID common.UUID, userID string, req *dto.PutEventByIDRequest) (*dto.PutEventByIDData, error) {
	var response *dto.PutEventByIDData

	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// イベント取得
		event, err := u.eventRepo.FindByID(txCtx, eventID)
		if err != nil {
			return err
		}
		if event == nil || !event.IsActive {
			return domainEvent.ErrEventNotFound
		}

		// 権限チェック
		if event.CreatorID.String() != userID {
			return domainEvent.ErrUnauthorized
		}

		// フィールド更新
		if req.Title != "" {
			event.Title = req.Title
		}
		if req.Description != "" {
			event.Description = req.Description
		}
		if req.MaxParticipants > 0 {
			// 定員減少チェック
			currentRegistered, err := u.participantRepo.CountRegisteredParticipants(txCtx, eventID)
			if err != nil {
				return err
			}
			if req.MaxParticipants < currentRegistered {
				return domainEvent.ErrCannotReduceCapacity
			}
			event.MaxParticipants = req.MaxParticipants
		}
		if req.DiscordChannelID != "" {
			event.DiscordChannelID = req.DiscordChannelID
		}

		// 更新実行
		err = u.eventRepo.Update(txCtx, event)
		if err != nil {
			return err
		}

		// レスポンス作成
		response = &dto.PutEventByIDData{
			ID:               event.ID.String(),
			Title:            event.Title,
			Description:      event.Description,
			MaxParticipants:  event.MaxParticipants,
			DiscordChannelID: event.DiscordChannelID,
			UpdatedAt:        event.UpdatedAt.Format(time.RFC3339),
		}

		return nil
	})

	return response, err
}