package user

import "errors"

// TODO: BE-03-user-01で実装予定
// ユーザードメイン固有のエラー定義

var (
	// ユーザー関連エラー（実装予定）
	ErrUserNotFound         = errors.New("ユーザーが見つかりません")
	ErrInvalidEmail         = errors.New("無効なメールアドレスです")
	ErrDuplicateEmail       = errors.New("このメールアドレスは既に使用されています")
	ErrRivalLimitExceeded   = errors.New("ライバルは最大3人まで設定できます")
	ErrCannotRivalSelf      = errors.New("自分自身をライバルに設定できません")
	ErrInvalidPlatform      = errors.New("無効なプラットフォームです")
	ErrInvalidURL           = errors.New("無効なURL形式です")
)

// TODO: 詳細なエラー定義とエラーハンドリング実装は後続タスクで行う