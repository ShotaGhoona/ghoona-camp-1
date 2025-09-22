package common

import (
	"time"
	domainCommon "ghoona-camp-backend/internal/domain/common"
)

// APIResponse は共通APIレスポンス形式
type APIResponse[T any] struct {
	Data      T          `json:"data,omitempty"`
	Message   string     `json:"message"`
	Timestamp time.Time  `json:"timestamp"`
	RequestID string     `json:"request_id,omitempty"`
}

// APIErrorResponse はエラーレスポンス形式
type APIErrorResponse struct {
	Error     *ErrorDetail `json:"error"`
	Message   string       `json:"message,omitempty"`
	Timestamp time.Time    `json:"timestamp"`
	RequestID string       `json:"request_id,omitempty"`
}

// ErrorDetail はエラー詳細
type ErrorDetail struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// PaginatedResponse はページネーション付きレスポンス
type PaginatedResponse[T any] struct {
	Data       []T                        `json:"data"`
	Pagination *domainCommon.Pagination   `json:"pagination"`
	Message    string                     `json:"message"`
	Timestamp  time.Time                  `json:"timestamp"`
	RequestID  string                     `json:"request_id,omitempty"`
}

// NewAPIResponse は成功レスポンスを作成
func NewAPIResponse[T any](data T, message string) *APIResponse[T] {
	return &APIResponse[T]{
		Data:      data,
		Message:   message,
		Timestamp: time.Now(),
	}
}

// NewAPIErrorResponse はエラーレスポンスを作成
func NewAPIErrorResponse(code, message string, details interface{}) *APIErrorResponse {
	return &APIErrorResponse{
		Error: &ErrorDetail{
			Code:    code,
			Message: message,
			Details: details,
		},
		Timestamp: time.Now(),
	}
}

// NewPaginatedResponse はページネーション付きレスポンスを作成
func NewPaginatedResponse[T any](data []T, pagination *domainCommon.Pagination, message string) *PaginatedResponse[T] {
	return &PaginatedResponse[T]{
		Data:       data,
		Pagination: pagination,
		Message:    message,
		Timestamp:  time.Now(),
	}
}

// WithRequestID はリクエストIDを設定
func (r *APIResponse[T]) WithRequestID(requestID string) *APIResponse[T] {
	r.RequestID = requestID
	return r
}

// WithRequestID はリクエストIDを設定
func (r *APIErrorResponse) WithRequestID(requestID string) *APIErrorResponse {
	r.RequestID = requestID
	return r
}

// WithRequestID はリクエストIDを設定
func (r *PaginatedResponse[T]) WithRequestID(requestID string) *PaginatedResponse[T] {
	r.RequestID = requestID
	return r
}