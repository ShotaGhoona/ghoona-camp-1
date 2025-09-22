package attendance

import "errors"

// TODO: BE-03-attendance-01で実装予定
// 出席管理ドメイン固有のエラー定義

var (
	// 出席記録関連エラー（実装予定）
	ErrAttendanceNotFound     = errors.New("出席記録が見つかりません")
	ErrInvalidDuration        = errors.New("無効な滞在時間です")
	ErrDuplicateAttendance    = errors.New("既に出席記録が存在します")
	ErrInvalidTimeRange       = errors.New("無効な時間範囲です")
	ErrChannelNotActive       = errors.New("Discordチャンネルがアクティブではありません")
)

// TODO: 詳細なエラー定義とエラーハンドリング実装は後続タスクで行う