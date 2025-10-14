package notification

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/notification/entity"
	"ghoona-camp-backend/internal/domain/notification/repository"
	notificationModel "ghoona-camp-backend/internal/infrastructure/gorm/model/notification"
)

type notificationSettingsRepository struct {
	db *gorm.DB
}

// NewNotificationSettingsRepository コンストラクタ
func NewNotificationSettingsRepository(db *gorm.DB) repository.NotificationSettingsRepository {
	return &notificationSettingsRepository{db: db}
}

// Create 通知設定を作成する
func (r *notificationSettingsRepository) Create(ctx context.Context, settings *entity.NotificationSettings) error {
	gormSettings := r.toGORMNotificationSettings(settings)
	return r.db.WithContext(ctx).Create(gormSettings).Error
}

// GetByUserID ユーザーIDで通知設定を取得する
func (r *notificationSettingsRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.NotificationSettings, error) {
	var gormSettings notificationModel.NotificationSettings
	err := r.db.WithContext(ctx).First(&gormSettings, "user_id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMNotificationSettings(&gormSettings)
}

// Update 通知設定を更新する
func (r *notificationSettingsRepository) Update(ctx context.Context, settings *entity.NotificationSettings) error {
	gormSettings := r.toGORMNotificationSettings(settings)
	return r.db.WithContext(ctx).Save(gormSettings).Error
}

// Delete 通知設定を削除する
func (r *notificationSettingsRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&notificationModel.NotificationSettings{}, "user_id = ?", userID).Error
}

// Upsert 通知設定を作成または更新する
func (r *notificationSettingsRepository) Upsert(ctx context.Context, settings *entity.NotificationSettings) error {
	gormSettings := r.toGORMNotificationSettings(settings)
	return r.db.WithContext(ctx).Save(gormSettings).Error
}

// 型変換: Domain Entity → GORM Model
func (r *notificationSettingsRepository) toGORMNotificationSettings(settings *entity.NotificationSettings) *notificationModel.NotificationSettings {
	return &notificationModel.NotificationSettings{
		UserID:         settings.UserID(),
		EmailEnabled:   settings.EmailEnabled(),
		PushEnabled:    settings.PushEnabled(),
		SMSEnabled:     settings.SMSEnabled(),
		EventReminders: settings.EventReminders(),
		GoalUpdates:    settings.GoalUpdates(),
		TitleAchievements: settings.TitleAchievements(),
		RivalUpdates:   settings.RivalUpdates(),
		CreatedAt:      settings.CreatedAt(),
		UpdatedAt:      settings.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *notificationSettingsRepository) fromGORMNotificationSettings(gormSettings *notificationModel.NotificationSettings) (*entity.NotificationSettings, error) {
	// TODO: entity.NewNotificationSettings の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}