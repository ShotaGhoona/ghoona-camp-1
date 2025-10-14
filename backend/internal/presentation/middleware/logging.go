package middleware

import (
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger 構造化ログミドルウェアを設定する
func Logger() gin.HandlerFunc {
	return logger.SetLogger(
		logger.WithLogger(func(c *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(
				zerolog.ConsoleWriter{
					Out:     gin.DefaultWriter,
					NoColor: false,
				},
			).With().
				Str("component", "gin").
				Logger()
		}),
		logger.WithUTC(true),
		logger.WithSkipPath([]string{"/health"}), // ヘルスチェックのログをスキップ
	)
}

// Recovery パニックリカバリーミドルウェアを設定する
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		log.Error().
			Interface("error", recovered).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("user_agent", c.Request.UserAgent()).
			Str("client_ip", c.ClientIP()).
			Msg("Panic recovered")

		c.JSON(500, gin.H{
			"error": gin.H{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": "Internal server error",
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	})
}