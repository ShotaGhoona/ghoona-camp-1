package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"ghoona-camp-backend/internal/application/common"
)

// ErrorHandlerMiddleware はエラーハンドリングミドルウェア
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// パニックをキャッチしてエラーレスポンスを返す
				log.Printf("Panic recovered: %v\n%s", err, debug.Stack())
				
				errorResponse := common.NewAPIErrorResponse(
					"INTERNAL_ERROR",
					"内部エラーが発生しました",
					nil,
				)
				
				c.JSON(http.StatusInternalServerError, errorResponse)
				c.Abort()
			}
		}()

		c.Next()

		// エラーがある場合の処理
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			handleError(c, err.Err)
		}
	}
}

// handleError はエラーをHTTPレスポンスに変換
func handleError(c *gin.Context, err error) {
	// アプリケーションエラーの場合
	if appErr, ok := err.(*common.ApplicationError); ok {
		errorResponse := common.NewAPIErrorResponse(
			appErr.Code,
			appErr.Message,
			nil,
		)
		c.JSON(appErr.HTTPStatus, errorResponse)
		return
	}

	// ドメインエラーの場合はアプリケーションエラーに変換
	appErr := common.ConvertDomainError(err)
	if appErr != nil {
		errorResponse := common.NewAPIErrorResponse(
			appErr.Code,
			appErr.Message,
			nil,
		)
		c.JSON(appErr.HTTPStatus, errorResponse)
		return
	}

	// その他のエラーは内部エラーとして処理
	log.Printf("Unexpected error: %v", err)
	errorResponse := common.NewAPIErrorResponse(
		"INTERNAL_ERROR",
		"予期しないエラーが発生しました",
		nil,
	)
	c.JSON(http.StatusInternalServerError, errorResponse)
}

// ValidationErrorHandler はバリデーションエラー専用ハンドラー
func ValidationErrorHandler(c *gin.Context, errors []common.ValidationErrorDetail) {
	errorResponse := common.NewAPIErrorResponse(
		"VALIDATION_ERROR",
		"バリデーションエラーが発生しました",
		map[string]interface{}{
			"fields": errors,
		},
	)
	c.JSON(http.StatusBadRequest, errorResponse)
}

// NotFoundHandler は404エラーハンドラー
func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		errorResponse := common.NewAPIErrorResponse(
			"NOT_FOUND",
			"リクエストされたリソースが見つかりません",
			map[string]interface{}{
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
			},
		)
		c.JSON(http.StatusNotFound, errorResponse)
	}
}

// MethodNotAllowedHandler は405エラーハンドラー
func MethodNotAllowedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		errorResponse := common.NewAPIErrorResponse(
			"METHOD_NOT_ALLOWED",
			"このHTTPメソッドは許可されていません",
			map[string]interface{}{
				"method": c.Request.Method,
				"path":   c.Request.URL.Path,
			},
		)
		c.JSON(http.StatusMethodNotAllowed, errorResponse)
	}
}