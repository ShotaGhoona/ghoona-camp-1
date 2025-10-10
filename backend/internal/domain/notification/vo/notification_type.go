package vo

import "errors"

// NotificationType 通知の種別を表すValue Object
type NotificationType string

const (
	NotificationTypeAchievement NotificationType = "achievement"   // 称号獲得
	NotificationTypeReminder    NotificationType = "reminder"      // リマインダー
	NotificationTypeRivalUpdate NotificationType = "rival_update"  // ライバル更新
	NotificationTypeEvent       NotificationType = "event"         // イベント関連
)

// NewNotificationType NotificationTypeを作成する
func NewNotificationType(value string) (NotificationType, error) {
	notificationType := NotificationType(value)
	if !notificationType.isValid() {
		return "", errors.New("不正な通知タイプです")
	}
	return notificationType, nil
}

// isValid 有効な通知タイプ値かどうかを検証する
func (t NotificationType) isValid() bool {
	switch t {
	case NotificationTypeAchievement, NotificationTypeReminder,
		NotificationTypeRivalUpdate, NotificationTypeEvent:
		return true
	default:
		return false
	}
}

// String 文字列表現を返す
func (t NotificationType) String() string {
	return string(t)
}

// Value 内部の文字列値を返す
func (t NotificationType) Value() string {
	return string(t)
}