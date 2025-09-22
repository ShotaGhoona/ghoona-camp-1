package common

import (
	"errors"
	"net/http"
	domainCommon "ghoona-camp-backend/internal/domain/common"
)

// ConvertToHTTPStatus はドメインエラーからHTTPステータスコードを取得
// 使用予定: エラーハンドリングミドルウェア (BE-03-*で各コントローラー実装時)
func ConvertToHTTPStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	// ドメイン共通エラーの変換
	switch {
	case errors.Is(err, domainCommon.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, domainCommon.ErrAlreadyExists):
		return http.StatusConflict
	case errors.Is(err, domainCommon.ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.Is(err, domainCommon.ErrForbidden):
		return http.StatusForbidden
	case errors.Is(err, domainCommon.ErrInvalidInput):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// ConvertToErrorCode はドメインエラーからAPIエラーコードを取得
// 使用予定: エラーレスポンス生成 (BE-03-*で各コントローラー実装時)
func ConvertToErrorCode(err error) string {
	if err == nil {
		return ""
	}

	// ドメイン共通エラーの変換
	switch {
	case errors.Is(err, domainCommon.ErrNotFound):
		return "NOT_FOUND"
	case errors.Is(err, domainCommon.ErrAlreadyExists):
		return "ALREADY_EXISTS"
	case errors.Is(err, domainCommon.ErrUnauthorized):
		return "UNAUTHORIZED"
	case errors.Is(err, domainCommon.ErrForbidden):
		return "FORBIDDEN"
	case errors.Is(err, domainCommon.ErrInvalidInput):
		return "VALIDATION_ERROR"
	default:
		return "INTERNAL_ERROR"
	}
}

// ValidationErrorDetail はバリデーションエラーの詳細
// 使用予定: バリデーションエラーレスポンス (BE-03-*でリクエスト検証時)
type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}