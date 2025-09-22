package goal

import "errors"

// TODO: BE-03-goal-01で実装予定
// 目標管理ドメイン固有のエラー定義

var (
	// 目標関連エラー（実装予定）
	ErrGoalNotFound          = errors.New("目標が見つかりません")
	ErrGoalLimitExceeded     = errors.New("目標数の上限（10個）を超えています")
	ErrGoalPeriodTooLong     = errors.New("目標期間は最大1年間です")
	ErrProgressDateInvalid   = errors.New("進捗記録は過去30日以内のみ可能です")
	ErrGoalNotPublic         = errors.New("非公開の目標です")
	ErrProgressAlreadyExists = errors.New("指定日の進捗が既に存在します")
)

// TODO: 詳細なエラー定義とエラーハンドリング実装は後続タスクで行う