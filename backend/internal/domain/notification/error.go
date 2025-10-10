package notification

import "errors"

// Notification domain specific errors

// Notification validation errors
var (
	ErrNotificationUserIDRequired = errors.New("ユーザーIDは必須です")
	ErrNotificationTitleRequired  = errors.New("通知タイトルは必須です")
	ErrNotificationMessageRequired = errors.New("通知メッセージは必須です")
)