package title

import "errors"

// TODO: BE-03-title-01で実装予定
// 称号管理ドメイン固有のエラー定義

var (
	// 称号関連エラー（実装予定）
	ErrTitleNotFound          = errors.New("称号が見つかりません")
	ErrTitleAlreadyAchieved   = errors.New("既に獲得済みの称号です")
	ErrInvalidLevel           = errors.New("無効なレベルです")
	ErrRequirementNotMet      = errors.New("獲得条件を満たしていません")
	ErrTitleNotUnlocked       = errors.New("称号が解放されていません")
)

// TODO: 詳細なエラー定義とエラーハンドリング実装は後続タスクで行う