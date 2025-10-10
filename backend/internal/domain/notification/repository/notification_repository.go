package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/notification/entity"
	"ghoona-camp-backend/internal/domain/notification/vo"
)

// NotificationRepository 通知リポジトリインターフェース
type NotificationRepository interface {
	// Create 通知を作成する
	Create(ctx context.Context, notification *entity.Notification) error

	// GetByID IDで通知を取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Notification, error)

	// GetByUserID ユーザーIDで通知一覧を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.Notification, error)

	// GetUnreadByUserID ユーザーの未読通知一覧を取得する
	GetUnreadByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.Notification, error)

	// GetByUserIDAndType ユーザーIDとタイプで通知を取得する
	GetByUserIDAndType(ctx context.Context, userID uuid.UUID, notificationType vo.NotificationType, limit, offset int) ([]*entity.Notification, error)

	// Update 通知を更新する
	Update(ctx context.Context, notification *entity.Notification) error

	// MarkAsRead 通知を既読にマークする
	MarkAsRead(ctx context.Context, id uuid.UUID) error

	// MarkAllAsReadByUserID ユーザーのすべての通知を既読にマークする
	MarkAllAsReadByUserID(ctx context.Context, userID uuid.UUID) error

	// Delete 通知を削除する
	Delete(ctx context.Context, id uuid.UUID) error

	// CountUnreadByUserID ユーザーの未読通知数を取得する
	CountUnreadByUserID(ctx context.Context, userID uuid.UUID) (int, error)
}