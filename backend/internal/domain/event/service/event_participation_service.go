package service

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/repository"
	"ghoona-camp-backend/internal/domain/event/vo"
)

// EventParticipationService イベント参加管理のビジネスロジックを提供するドメインサービス
type EventParticipationService struct {
	eventRepo           repository.EventRepository
	eventParticipantRepo repository.EventParticipantRepository
}

// NewEventParticipationService コンストラクタ
func NewEventParticipationService(
	eventRepo repository.EventRepository,
	eventParticipantRepo repository.EventParticipantRepository,
) *EventParticipationService {
	return &EventParticipationService{
		eventRepo:           eventRepo,
		eventParticipantRepo: eventParticipantRepo,
	}
}

// CanJoinEvent イベント参加が可能かチェックする
func (s *EventParticipationService) CanJoinEvent(ctx context.Context, eventID, userID uuid.UUID) error {
	// 1. イベントを取得
	eventEntity, err := s.eventRepo.GetByID(ctx, eventID)
	if err != nil {
		return err
	}

	// 2. 既に参加しているかチェック（重複参加防止）
	existingParticipant, err := s.eventParticipantRepo.GetByEventAndUser(ctx, eventID, userID)
	if err == nil && existingParticipant != nil {
		// キャンセル状態でない場合は重複参加とみなす
		if existingParticipant.Status() == vo.ParticipantStatusRegistered {
			return event.ErrEventParticipantAlreadyRegistered
		}
	}

	// 3. 定員チェック（max_participantsが設定されている場合）
	maxParticipants := eventEntity.MaxParticipants()
	if maxParticipants != nil {
		// 現在の登録済み参加者数を取得
		currentCount, err := s.eventParticipantRepo.CountByEventIDAndStatus(ctx, eventID, vo.ParticipantStatusRegistered)
		if err != nil {
			return err
		}

		// 定員オーバーチェック
		if currentCount >= *maxParticipants {
			return event.ErrEventCapacityFull
		}
	}

	return nil
}

// JoinEvent イベントに参加する（ビジネスルールチェック付き）
func (s *EventParticipationService) JoinEvent(ctx context.Context, eventID, userID uuid.UUID) (*entity.EventParticipant, error) {
	// ビジネスルールチェック
	if err := s.CanJoinEvent(ctx, eventID, userID); err != nil {
		return nil, err
	}

	// 既存のキャンセル済み参加者がいる場合は更新、ない場合は新規作成
	existingParticipant, err := s.eventParticipantRepo.GetByEventAndUser(ctx, eventID, userID)
	if err == nil && existingParticipant != nil && existingParticipant.Status() == vo.ParticipantStatusCancelled {
		// キャンセル済みの参加者を登録済みに変更
		// ここではシンプルに新しいエンティティを作成して返す
	}

	// 参加者エンティティ作成
	participant := entity.NewEventParticipant(eventID, userID, vo.ParticipantStatusRegistered)

	// リポジトリに保存
	if err := s.eventParticipantRepo.Create(ctx, participant); err != nil {
		return nil, err
	}

	return participant, nil
}