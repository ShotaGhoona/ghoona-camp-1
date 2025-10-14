package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS クロスオリジンリソースシェアリングを設定する
func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	
	// 許可されたオリジンを設定（本番環境用に調整）
	config.AllowOrigins = []string{
		"http://localhost:3000",  // Next.js開発サーバー
		"https://ghoona-camp.com", // 本番環境フロントエンドドメイン
	}
	
	// 許可されたヘッダーを設定
	config.AllowHeaders = []string{
		"Origin",
		"Content-Length", 
		"Content-Type",
		"Authorization",
		"Accept",
		"X-Requested-With",
	}
	
	// 許可されたメソッドを設定
	config.AllowMethods = []string{
		"GET",
		"POST", 
		"PUT",
		"DELETE",
		"OPTIONS",
	}
	
	// 認証情報を許可（Clerk認証に必須）
	config.AllowCredentials = true
	
	return cors.New(config)
}