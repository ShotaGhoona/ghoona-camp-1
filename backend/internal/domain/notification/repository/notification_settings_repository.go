package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/notification/entity"
)

// NotificationSettingsRepository 通知設定リポジトリインターフェース
type NotificationSettingsRepository interface {
	// Create 通知設定を作成する
	Create(ctx context.Context, settings *entity.NotificationSettings) error

	// GetByUserID ユーザーIDで通知設定を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.NotificationSettings, error)

	// Update 通知設定を更新する
	Update(ctx context.Context, settings *entity.NotificationSettings) error

	// Delete 通知設定を削除する
	Delete(ctx context.Context, userID uuid.UUID) error

	// Upsert 通知設定を作成または更新する
	Upsert(ctx context.Context, settings *entity.NotificationSettings) error
}