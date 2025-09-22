package common

import (
	"time"
	domainCommon "ghoona-camp-backend/internal/domain/common"
)

// SuccessResponse は成功レスポンス形式
// 使用予定: 全API成功レスポンス (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
type SuccessResponse struct {
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Timestamp time.Time   `json:"timestamp"`
}

// ErrorResponse はエラーレスポンス形式  
// 使用予定: 全APIエラーレスポンス (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
type ErrorResponse struct {
	Error     *ErrorDetail `json:"error"`
	Timestamp time.Time    `json:"timestamp"`
}

// ErrorDetail はエラー詳細
// 使用予定: API仕様書通りのエラー構造 (code, message, details)
type ErrorDetail struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// PaginatedResponse はページネーション付きレスポンス
// 使用予定: リスト取得API (BE-03-user-04, BE-04-attend-04, BE-06-goal-04, BE-07-event-04, BE-08-title-04, BE-09-notify-04)
type PaginatedResponse struct {
	Data       interface{}                `json:"data"`
	Pagination *domainCommon.Pagination   `json:"pagination"`
	Message    string                     `json:"message"`
	Timestamp  time.Time                  `json:"timestamp"`
}

// NewSuccessResponse は成功レスポンスを作成
// 使用予定: 全成功レスポンス生成 (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
func NewSuccessResponse(data interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		Data:      data,
		Message:   message,
		Timestamp: time.Now(),
	}
}

// NewErrorResponse はエラーレスポンスを作成
// 使用予定: 全エラーレスポンス生成 (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
func NewErrorResponse(code, message string, details interface{}) *ErrorResponse {
	return &ErrorResponse{
		Error: &ErrorDetail{
			Code:    code,
			Message: message,
			Details: details,
		},
		Timestamp: time.Now(),
	}
}

// NewPaginatedResponse はページネーション付きレスポンスを作成
// 使用予定: リスト取得レスポンス生成 (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
func NewPaginatedResponse(data interface{}, pagination *domainCommon.Pagination, message string) *PaginatedResponse {
	return &PaginatedResponse{
		Data:       data,
		Pagination: pagination,
		Message:    message,
		Timestamp:  time.Now(),
	}
}