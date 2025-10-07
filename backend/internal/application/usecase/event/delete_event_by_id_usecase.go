// Package usecase DELETE /events/{eventId} API用のUseCase
package event

import (
	"context"

	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/domain/common"
	domainEvent "ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/repository"
)

// DeleteEventByIDUseCase イベント削除（論理削除）
type DeleteEventByIDUseCase interface {
	Execute(ctx context.Context, eventID common.UUID, userID string) error
}

type deleteEventByIDUseCase struct {
	eventRepo repository.EventRepository
	txManager transaction.Manager
}

// NewDeleteEventByIDUseCase コンストラクタ
func NewDeleteEventByIDUseCase(
	eventRepo repository.EventRepository,
	txManager transaction.Manager,
) DeleteEventByIDUseCase {
	return &deleteEventByIDUseCase{
		eventRepo: eventRepo,
		txManager: txManager,
	}
}

// Execute イベント削除を実行
func (u *deleteEventByIDUseCase) Execute(ctx context.Context, eventID common.UUID, userID string) error {
	return u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// イベント取得
		event, err := u.eventRepo.FindByID(txCtx, eventID)
		if err != nil {
			return err
		}
		if event == nil {
			return domainEvent.ErrEventNotFound
		}

		// 権限チェック
		if event.CreatorID.String() != userID {
			return domainEvent.ErrUnauthorized
		}

		// 既に削除済みでも成功（冪等性）
		if !event.IsActive {
			return nil
		}

		// 論理削除実行
		event.IsActive = false
		return u.eventRepo.Update(txCtx, event)
	})
}