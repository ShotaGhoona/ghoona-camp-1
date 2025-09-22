package common

import (
	"errors"
	"fmt"
	"net/http"
	domainCommon "ghoona-camp-backend/internal/domain/common"
)

// ApplicationError はアプリケーション層のエラー型
type ApplicationError struct {
	Code       string
	Message    string
	HTTPStatus int
	Cause      error
}

func (e *ApplicationError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *ApplicationError) Unwrap() error {
	return e.Cause
}

// NewApplicationError はアプリケーションエラーを作成
func NewApplicationError(code, message string, httpStatus int, cause error) *ApplicationError {
	return &ApplicationError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
		Cause:      cause,
	}
}

// エラーコード定数
const (
	ErrorCodeValidation   = "VALIDATION_ERROR"
	ErrorCodeNotFound     = "NOT_FOUND"
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	ErrorCodeForbidden    = "FORBIDDEN"
	ErrorCodeConflict     = "CONFLICT"
	ErrorCodeInternal     = "INTERNAL_ERROR"
	ErrorCodeBadRequest   = "BAD_REQUEST"
)

// 事前定義されたアプリケーションエラー
var (
	ErrValidation   = NewApplicationError(ErrorCodeValidation, "バリデーションエラーが発生しました", http.StatusBadRequest, nil)
	ErrNotFound     = NewApplicationError(ErrorCodeNotFound, "リソースが見つかりません", http.StatusNotFound, nil)
	ErrUnauthorized = NewApplicationError(ErrorCodeUnauthorized, "認証が必要です", http.StatusUnauthorized, nil)
	ErrForbidden    = NewApplicationError(ErrorCodeForbidden, "アクセスが禁止されています", http.StatusForbidden, nil)
	ErrConflict     = NewApplicationError(ErrorCodeConflict, "競合が発生しました", http.StatusConflict, nil)
	ErrInternal     = NewApplicationError(ErrorCodeInternal, "内部エラーが発生しました", http.StatusInternalServerError, nil)
	ErrBadRequest   = NewApplicationError(ErrorCodeBadRequest, "不正なリクエストです", http.StatusBadRequest, nil)
)

// ConvertDomainError はドメインエラーをアプリケーションエラーに変換
func ConvertDomainError(err error) *ApplicationError {
	if err == nil {
		return nil
	}

	// 既にアプリケーションエラーの場合はそのまま返す
	var appErr *ApplicationError
	if errors.As(err, &appErr) {
		return appErr
	}

	// ドメインエラーの場合は変換
	var domainErr *domainCommon.DomainError
	if errors.As(err, &domainErr) {
		return convertDomainErrorToAppError(domainErr)
	}

	// バリデーションエラーの場合は変換
	var validationErr *domainCommon.ValidationError
	if errors.As(err, &validationErr) {
		return NewApplicationError(
			ErrorCodeValidation,
			validationErr.Error(),
			http.StatusBadRequest,
			err,
		)
	}

	// 共通ドメインエラーの変換
	switch {
	case errors.Is(err, domainCommon.ErrNotFound):
		return NewApplicationError(ErrorCodeNotFound, err.Error(), http.StatusNotFound, err)
	case errors.Is(err, domainCommon.ErrAlreadyExists):
		return NewApplicationError(ErrorCodeConflict, err.Error(), http.StatusConflict, err)
	case errors.Is(err, domainCommon.ErrUnauthorized):
		return NewApplicationError(ErrorCodeUnauthorized, err.Error(), http.StatusUnauthorized, err)
	case errors.Is(err, domainCommon.ErrForbidden):
		return NewApplicationError(ErrorCodeForbidden, err.Error(), http.StatusForbidden, err)
	case errors.Is(err, domainCommon.ErrInvalidInput):
		return NewApplicationError(ErrorCodeValidation, err.Error(), http.StatusBadRequest, err)
	default:
		return NewApplicationError(ErrorCodeInternal, "予期しないエラーが発生しました", http.StatusInternalServerError, err)
	}
}

// convertDomainErrorToAppError はドメインエラーをアプリケーションエラーに変換
func convertDomainErrorToAppError(domainErr *domainCommon.DomainError) *ApplicationError {
	switch domainErr.Code {
	case "NOT_FOUND":
		return NewApplicationError(ErrorCodeNotFound, domainErr.Message, http.StatusNotFound, domainErr)
	case "VALIDATION":
		return NewApplicationError(ErrorCodeValidation, domainErr.Message, http.StatusBadRequest, domainErr)
	case "UNAUTHORIZED":
		return NewApplicationError(ErrorCodeUnauthorized, domainErr.Message, http.StatusUnauthorized, domainErr)
	case "FORBIDDEN":
		return NewApplicationError(ErrorCodeForbidden, domainErr.Message, http.StatusForbidden, domainErr)
	case "CONFLICT":
		return NewApplicationError(ErrorCodeConflict, domainErr.Message, http.StatusConflict, domainErr)
	default:
		return NewApplicationError(ErrorCodeInternal, domainErr.Message, http.StatusInternalServerError, domainErr)
	}
}

// ValidationErrorDetail はバリデーションエラーの詳細
type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewValidationError は複数のバリデーションエラーを持つアプリケーションエラーを作成
func NewValidationError(details []ValidationErrorDetail) *ApplicationError {
	return NewApplicationError(
		ErrorCodeValidation,
		"バリデーションエラーが発生しました",
		http.StatusBadRequest,
		nil,
	)
}