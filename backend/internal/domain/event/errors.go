package event

import "errors"

// TODO: BE-03-event-01で実装予定
// イベント管理ドメイン固有のエラー定義

var (
	// イベント関連エラー（実装予定）
	ErrEventNotFound           = errors.New("イベントが見つかりません")
	ErrEventFull              = errors.New("イベントが満員です")
	ErrDuplicateRegistration  = errors.New("既に参加登録済みです")
	ErrRegistrationDeadline   = errors.New("参加登録期限を過ぎています")
	ErrCancellationDeadline   = errors.New("キャンセル期限を過ぎています")
	ErrInvalidEventType       = errors.New("無効なイベントタイプです")
	ErrInvalidTimeRange       = errors.New("開始時間は終了時間より前である必要があります")
	ErrPastEventDate          = errors.New("過去の日付にはイベントを作成できません")
)

// TODO: 詳細なエラー定義とエラーハンドリング実装は後続タスクで行う