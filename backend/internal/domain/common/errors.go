package common

import "errors"

// 共通ドメインエラー定義
var (
	// レコード関連エラー
	ErrNotFound          = errors.New("レコードが見つかりません")
	ErrAlreadyExists     = errors.New("レコードが既に存在します")
	ErrDuplicateEntry    = errors.New("重複するエントリです")
	
	// バリデーションエラー
	ErrInvalidInput      = errors.New("入力値が無効です")
	ErrInvalidFormat     = errors.New("フォーマットが無効です")
	ErrRequired          = errors.New("必須項目が未入力です")
	ErrTooLong           = errors.New("入力値が長すぎます")
	ErrTooShort          = errors.New("入力値が短すぎます")
	
	// 権限・認証エラー
	ErrUnauthorized      = errors.New("認証が必要です")
	ErrForbidden         = errors.New("アクセス権限がありません")
	ErrInvalidToken      = errors.New("無効なトークンです")
	ErrTokenExpired      = errors.New("トークンの有効期限が切れています")
	
	// ビジネスルールエラー
	ErrInvalidState      = errors.New("無効な状態です")
	ErrInvalidTransition = errors.New("無効な状態遷移です")
	ErrBusinessRule      = errors.New("ビジネスルール違反です")
	
	// 外部サービスエラー
	ErrExternalService   = errors.New("外部サービスエラーです")
	ErrNetworkError      = errors.New("ネットワークエラーです")
	ErrTimeout           = errors.New("タイムアウトしました")
)

// DomainError はドメイン固有のエラー型
type DomainError struct {
	Code    string
	Message string
	Cause   error
}

func (e *DomainError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Cause
}

// NewDomainError はドメインエラーを作成
func NewDomainError(code, message string, cause error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

// ValidationError はバリデーションエラー型
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// NewValidationError はバリデーションエラーを作成
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}