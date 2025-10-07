// Package repository イベント参加者ドメインのリポジトリインターフェースを定義する
package repository

import (
	"context"

	"ghoona-camp-backend/internal/domain/event/entity"

	"github.com/google/uuid"
)

// EventParticipantRepository イベント参加者の永続化を抽象化するインターフェース
type EventParticipantRepository interface {
	// イベント別参加者一覧 - 参加者一覧表示APIで使用
	FindByEventID(ctx context.Context, eventID uuid.UUID) ([]*entity.EventParticipant, error)

	// イベント×ユーザー検索 - 重複参加チェック、ステータス更新で使用
	FindByEventAndUser(ctx context.Context, eventID, userID uuid.UUID) (*entity.EventParticipant, error)

	// 作成 - イベント参加登録APIで使用
	Create(ctx context.Context, participant *entity.EventParticipant) error

	// 更新 - 参加ステータス変更APIで使用
	Update(ctx context.Context, participant *entity.EventParticipant) error

	// 登録者数カウント - 定員チェック、一覧表示のparticipant_countで使用
	CountRegisteredParticipants(ctx context.Context, eventID uuid.UUID) (int, error)

	// 存在確認 - 重複登録防止の簡易チェックで使用
	ExistsByEventAndUser(ctx context.Context, eventID, userID uuid.UUID) (bool, error)
}
