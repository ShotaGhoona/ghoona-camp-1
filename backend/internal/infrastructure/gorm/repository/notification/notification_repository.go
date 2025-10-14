package notification

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/notification/entity"
	"ghoona-camp-backend/internal/domain/notification/repository"
	"ghoona-camp-backend/internal/domain/notification/vo"
	notificationModel "ghoona-camp-backend/internal/infrastructure/gorm/model/notification"
)

type notificationRepository struct {
	db *gorm.DB
}

// NewNotificationRepository コンストラクタ
func NewNotificationRepository(db *gorm.DB) repository.NotificationRepository {
	return &notificationRepository{db: db}
}

// Create 通知を作成する
func (r *notificationRepository) Create(ctx context.Context, notification *entity.Notification) error {
	gormNotification := r.toGORMNotification(notification)
	return r.db.WithContext(ctx).Create(gormNotification).Error
}

// GetByID IDで通知を取得する
func (r *notificationRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Notification, error) {
	var gormNotification notificationModel.Notification
	err := r.db.WithContext(ctx).First(&gormNotification, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMNotification(&gormNotification)
}

// GetByUserID ユーザーIDで通知一覧を取得する
func (r *notificationRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.Notification, error) {
	var gormNotifications []notificationModel.Notification
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&gormNotifications).Error
	if err != nil {
		return nil, err
	}

	notifications := make([]*entity.Notification, len(gormNotifications))
	for i, gormNotification := range gormNotifications {
		notification, err := r.fromGORMNotification(&gormNotification)
		if err != nil {
			return nil, err
		}
		notifications[i] = notification
	}
	return notifications, nil
}

// GetUnreadByUserID ユーザーの未読通知一覧を取得する
func (r *notificationRepository) GetUnreadByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.Notification, error) {
	var gormNotifications []notificationModel.Notification
	err := r.db.WithContext(ctx).Where("user_id = ? AND is_read = ?", userID, false).Limit(limit).Offset(offset).Find(&gormNotifications).Error
	if err != nil {
		return nil, err
	}

	notifications := make([]*entity.Notification, len(gormNotifications))
	for i, gormNotification := range gormNotifications {
		notification, err := r.fromGORMNotification(&gormNotification)
		if err != nil {
			return nil, err
		}
		notifications[i] = notification
	}
	return notifications, nil
}

// GetByUserIDAndType ユーザーIDとタイプで通知を取得する
func (r *notificationRepository) GetByUserIDAndType(ctx context.Context, userID uuid.UUID, notificationType vo.NotificationType, limit, offset int) ([]*entity.Notification, error) {
	var gormNotifications []notificationModel.Notification
	err := r.db.WithContext(ctx).Where("user_id = ? AND type = ?", userID, string(notificationType)).Limit(limit).Offset(offset).Find(&gormNotifications).Error
	if err != nil {
		return nil, err
	}

	notifications := make([]*entity.Notification, len(gormNotifications))
	for i, gormNotification := range gormNotifications {
		notification, err := r.fromGORMNotification(&gormNotification)
		if err != nil {
			return nil, err
		}
		notifications[i] = notification
	}
	return notifications, nil
}

// Update 通知を更新する
func (r *notificationRepository) Update(ctx context.Context, notification *entity.Notification) error {
	gormNotification := r.toGORMNotification(notification)
	return r.db.WithContext(ctx).Save(gormNotification).Error
}

// MarkAsRead 通知を既読にマークする
func (r *notificationRepository) MarkAsRead(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&notificationModel.Notification{}).Where("id = ?", id).Update("is_read", true).Error
}

// MarkAllAsReadByUserID ユーザーのすべての通知を既読にマークする
func (r *notificationRepository) MarkAllAsReadByUserID(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&notificationModel.Notification{}).Where("user_id = ?", userID).Update("is_read", true).Error
}

// Delete 通知を削除する
func (r *notificationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&notificationModel.Notification{}, "id = ?", id).Error
}

// CountUnreadByUserID ユーザーの未読通知数を取得する
func (r *notificationRepository) CountUnreadByUserID(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&notificationModel.Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Count(&count).Error
	return int(count), err
}

// 型変換: Domain Entity → GORM Model
func (r *notificationRepository) toGORMNotification(notification *entity.Notification) *notificationModel.Notification {
	return &notificationModel.Notification{
		ID:        notification.ID(),
		UserID:    notification.UserID(),
		Type:      string(notification.Type()),
		Title:     notification.Title(),
		Message:   notification.Message(),
		Data:      notification.Data(),
		IsRead:    notification.IsRead(),
		ReadAt:    notification.ReadAt(),
		CreatedAt: notification.CreatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *notificationRepository) fromGORMNotification(gormNotification *notificationModel.Notification) (*entity.Notification, error) {
	// TODO: entity.NewNotification の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}