package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 共通型定義

// PaginationQuery ページネーション用クエリパラメータ
type PaginationQuery struct {
	Limit  int `form:"limit,default=20" binding:"min=1,max=100"`
	Offset int `form:"offset,default=0" binding:"min=0"`
}

// DateRangeQuery 日付範囲指定用クエリパラメータ
type DateRangeQuery struct {
	DateFrom string `form:"date_from" binding:"omitempty,datetime=2006-01-02"`
	DateTo   string `form:"date_to" binding:"omitempty,datetime=2006-01-02"`
}

// SearchQuery 検索用クエリパラメータ
type SearchQuery struct {
	Search string `form:"search" binding:"omitempty,max=100"`
}

// APIResponse 統一APIレスポンス形式
type APIResponse struct {
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message"`
	Timestamp string      `json:"timestamp"`
}

// APIErrorResponse エラーレスポンス形式
type APIErrorResponse struct {
	Error     ErrorDetail `json:"error"`
	Timestamp string      `json:"timestamp"`
}

// ErrorDetail エラー詳細
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// PaginationMeta ページネーション情報
type PaginationMeta struct {
	Total   int  `json:"total"`
	Limit   int  `json:"limit"`
	Offset  int  `json:"offset"`
	HasMore bool `json:"has_more"`
}

// 共通レスポンスヘルパー関数

// SendSuccessResponse 成功レスポンスを送信
func SendSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Data:      data,
		Message:   "success",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

// SendErrorResponse エラーレスポンスを送信
func SendErrorResponse(c *gin.Context, statusCode int, code, message string) {
	c.JSON(statusCode, APIErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

// SendValidationError バリデーションエラーを送信
func SendValidationError(c *gin.Context, message string) {
	SendErrorResponse(c, 400, "VALIDATION_ERROR", message)
}

// SendNotFoundError NotFoundエラーを送信
func SendNotFoundError(c *gin.Context, resource string) {
	SendErrorResponse(c, 404, "NOT_FOUND", resource+" not found")
}

// SendUnauthorizedError 認証エラーを送信
func SendUnauthorizedError(c *gin.Context) {
	SendErrorResponse(c, 401, "UNAUTHORIZED", "Authentication required")
}

// SendForbiddenError 認可エラーを送信
func SendForbiddenError(c *gin.Context) {
	SendErrorResponse(c, 403, "FORBIDDEN", "Access denied")
}

// SendInternalError 内部エラーを送信
func SendInternalError(c *gin.Context) {
	SendErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", "Internal server error")
}