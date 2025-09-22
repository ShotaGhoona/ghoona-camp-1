package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware はCORS設定ミドルウェア
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// 許可するオリジンのリスト
		allowedOrigins := []string{
			"http://localhost:3000",  // 開発環境のNext.js
			"http://localhost:3001",  // 開発環境の代替ポート
			"https://ghoona-camp.vercel.app", // 本番環境（例）
		}

		// オリジンチェック
		isAllowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				isAllowed = true
				break
			}
		}

		// 開発環境では全てのオリジンを許可
		// TODO: 本番環境では適切に制限する
		if origin == "" || isAllowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// CORS ヘッダーを設定
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		// プリフライトリクエストの処理
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// SecurityHeadersMiddleware はセキュリティヘッダーを設定
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// セキュリティヘッダーを設定
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		
		// HTTPS強制（本番環境のみ）
		// TODO: 本番環境判定を追加
		// c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		c.Next()
	}
}