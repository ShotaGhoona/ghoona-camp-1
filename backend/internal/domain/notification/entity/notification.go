package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/notification/vo"
)

// Notification 通知エンティティ
type Notification struct {
	id           uuid.UUID                  // 通知ID
	userID       uuid.UUID                  // 通知対象ユーザーID
	notificationType vo.NotificationType    // 通知タイプ
	title        string                     // 通知タイトル
	message      string                     // 通知メッセージ
	data         map[string]interface{}     // 関連データ
	isRead       vo.ReadFlag                // 既読状態
	scheduledAt  *time.Time                 // 送信予定時刻
	sentAt       *time.Time                 // 実際の送信時刻
	createdAt    time.Time                  // 作成日時
	updatedAt    time.Time                  // 更新日時
}

// NewNotification Notificationエンティティを作成する
func NewNotification(
	userID uuid.UUID,
	notificationType vo.NotificationType,
	title string,
	message string,
	data map[string]interface{},
	isRead vo.ReadFlag,
	scheduledAt *time.Time,
	sentAt *time.Time,
) (*Notification, error) {
	// バリデーション
	if title == "" {
		return nil, errors.New("通知タイトルは必須です")
	}
	if message == "" {
		return nil, errors.New("通知メッセージは必須です")
	}

	now := time.Now()
	return &Notification{
		id:               uuid.New(),
		userID:           userID,
		notificationType: notificationType,
		title:            title,
		message:          message,
		data:             data,
		isRead:           isRead,
		scheduledAt:      scheduledAt,
		sentAt:           sentAt,
		createdAt:        now,
		updatedAt:        now,
	}, nil
}

// Getters
func (n *Notification) ID() uuid.UUID                         { return n.id }
func (n *Notification) UserID() uuid.UUID                     { return n.userID }
func (n *Notification) NotificationType() vo.NotificationType { return n.notificationType }
func (n *Notification) Title() string                         { return n.title }
func (n *Notification) Message() string                       { return n.message }
func (n *Notification) Data() map[string]interface{}          { return n.data }
func (n *Notification) IsRead() vo.ReadFlag                   { return n.isRead }
func (n *Notification) ScheduledAt() *time.Time               { return n.scheduledAt }
func (n *Notification) SentAt() *time.Time                    { return n.sentAt }
func (n *Notification) CreatedAt() time.Time                  { return n.createdAt }
func (n *Notification) UpdatedAt() time.Time                  { return n.updatedAt }