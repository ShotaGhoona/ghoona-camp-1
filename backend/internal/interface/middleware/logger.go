package middleware

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerConfig はロガー設定
type LoggerConfig struct {
	Output    io.Writer
	SkipPaths []string
}

// LoggerMiddleware はログミドルウェア
func LoggerMiddleware(config ...LoggerConfig) gin.HandlerFunc {
	var conf LoggerConfig
	if len(config) > 0 {
		conf = config[0]
	}

	// デフォルト設定
	if conf.Output == nil {
		conf.Output = os.Stdout
	}

	// スキップパスのマップ化
	skipPaths := make(map[string]bool)
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return gin.LoggerWithConfig(gin.LoggerConfig{
		Output: conf.Output,
		Formatter: func(param gin.LogFormatterParams) string {
			// ヘルスチェック等のスキップ
			if skipPaths[param.Path] {
				return ""
			}

			// ログフォーマット
			return fmt.Sprintf("[%s] %s %s %d %s %s %s\n",
				param.TimeStamp.Format("2006/01/02 15:04:05"),
				param.Method,
				param.Path,
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.ErrorMessage,
			)
		},
		SkipPaths: conf.SkipPaths,
	})
}

// AccessLogMiddleware はアクセスログ専用ミドルウェア
func AccessLogMiddleware() gin.HandlerFunc {
	return LoggerMiddleware(LoggerConfig{
		SkipPaths: []string{
			"/health",
			"/metrics",
			"/favicon.ico",
		},
	})
}

// RequestLoggerMiddleware は詳細なリクエストログミドルウェア
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエスト開始時間
		start := time.Now()
		
		// ユーザーIDを取得（認証済みの場合）
		userID := "anonymous"
		if id, exists := c.Get("user_id"); exists {
			if uid, ok := id.(string); ok {
				userID = uid
			}
		}

		// リクエスト情報をログ
		fmt.Printf("📡 [REQUEST] %s %s %s - ユーザー: %s\n",
			start.Format("2006/01/02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			userID,
		)

		c.Next()

		// レスポンス情報をログ
		latency := time.Since(start)
		fmt.Printf("✅ [RESPONSE] %s %s %s %d %s - ユーザー: %s\n",
			time.Now().Format("2006/01/02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
			userID,
		)
	}
}

// ErrorLogMiddleware はエラー専用ログミドルウェア
func ErrorLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// エラーがある場合のみログ
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				fmt.Printf("❌ [ERROR] %s %s %s - %s\n",
					time.Now().Format("2006/01/02 15:04:05"),
					c.Request.Method,
					c.Request.URL.Path,
					err.Error(),
				)
			}
		}
	}
}