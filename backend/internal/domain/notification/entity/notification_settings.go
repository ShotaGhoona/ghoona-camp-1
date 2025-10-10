package entity

import (
	"time"

	"github.com/google/uuid"
)

// NotificationSettings 通知設定エンティティ
type NotificationSettings struct {
	id                      uuid.UUID // 設定ID
	userID                  uuid.UUID // ユーザーID
	achievementEnabled      bool      // 称号獲得通知の有効/無効
	reminderEnabled         bool      // リマインダー通知の有効/無効
	rivalUpdateEnabled      bool      // ライバル更新通知の有効/無効
	eventReminderEnabled    bool      // イベントリマインダーの有効/無効
	reminderTime            time.Time // リマインダー送信時刻
	createdAt               time.Time // 作成日時
	updatedAt               time.Time // 更新日時
}

// NewNotificationSettings NotificationSettingsエンティティを作成する（デフォルト値設定付き）
func NewNotificationSettings(
	userID uuid.UUID,
) *NotificationSettings {
	now := time.Now()
	// デフォルトのリマインダー時刻（21:00）
	defaultReminderTime := time.Date(0, 1, 1, 21, 0, 0, 0, time.UTC)
	
	return &NotificationSettings{
		id:                   uuid.New(),
		userID:               userID,
		achievementEnabled:   true,  // デフォルト: 有効
		reminderEnabled:      true,  // デフォルト: 有効
		rivalUpdateEnabled:   true,  // デフォルト: 有効
		eventReminderEnabled: true,  // デフォルト: 有効
		reminderTime:         defaultReminderTime,
		createdAt:            now,
		updatedAt:            now,
	}
}

// Getters
func (s *NotificationSettings) ID() uuid.UUID                { return s.id }
func (s *NotificationSettings) UserID() uuid.UUID            { return s.userID }
func (s *NotificationSettings) AchievementEnabled() bool     { return s.achievementEnabled }
func (s *NotificationSettings) ReminderEnabled() bool        { return s.reminderEnabled }
func (s *NotificationSettings) RivalUpdateEnabled() bool     { return s.rivalUpdateEnabled }
func (s *NotificationSettings) EventReminderEnabled() bool   { return s.eventReminderEnabled }
func (s *NotificationSettings) ReminderTime() time.Time      { return s.reminderTime }
func (s *NotificationSettings) CreatedAt() time.Time         { return s.createdAt }
func (s *NotificationSettings) UpdatedAt() time.Time         { return s.updatedAt }