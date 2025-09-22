package notification

import "errors"

// TODO: BE-03-notification-01で実装予定
// 通知管理ドメイン固有のエラー定義

var (
	// 通知関連エラー（実装予定）
	ErrNotificationNotFound   = errors.New("通知が見つかりません")
	ErrInvalidChannel         = errors.New("無効な通知チャンネルです")
	ErrNotificationDisabled   = errors.New("通知が無効になっています")
	ErrQuietHours            = errors.New("サイレント時間中です")
	ErrInvalidPriority       = errors.New("無効な優先度です")
	ErrScheduleConflict      = errors.New("スケジュールが競合しています")
)

// TODO: 詳細なエラー定義とエラーハンドリング実装は後続タスクで行う