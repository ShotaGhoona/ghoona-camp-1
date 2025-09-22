package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"ghoona-camp-backend/internal/application/common"
)

// ErrorHandlerMiddleware はエラーハンドリングミドルウェア
// 使用予定: router.goで全ルートに適用
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// パニックをキャッチしてエラーレスポンスを返す
				log.Printf("🚨 パニックが回復されました: %v\n%s", err, debug.Stack())
				
				errorResponse := common.NewErrorResponse(
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
// 使用予定: BE-03-*で各コントローラーエラー処理時
func handleError(c *gin.Context, err error) {
	// HTTPステータスとエラーコードを取得
	status := common.ConvertToHTTPStatus(err)
	code := common.ConvertToErrorCode(err)
	
	errorResponse := common.NewErrorResponse(code, err.Error(), nil)
	c.JSON(status, errorResponse)
}

// ValidationErrorHandler はバリデーションエラー専用ハンドラー
// 使用予定: BE-03-*でリクエストバリデーション時
func ValidationErrorHandler(c *gin.Context, errors []common.ValidationErrorDetail) {
	errorResponse := common.NewErrorResponse(
		"VALIDATION_ERROR",
		"バリデーションエラーが発生しました",
		map[string]interface{}{
			"fields": errors,
		},
	)
	c.JSON(http.StatusBadRequest, errorResponse)
}

// NotFoundHandler は404エラーハンドラー
// 使用予定: router.goで未定義ルート処理
func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		errorResponse := common.NewErrorResponse(
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
// 使用予定: router.goで未対応メソッド処理
func MethodNotAllowedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		errorResponse := common.NewErrorResponse(
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