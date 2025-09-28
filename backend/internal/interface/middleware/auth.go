package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"ghoona-camp-backend/internal/infrastructure/clerk"
)

// AuthMiddleware はClerk認証ミドルウェア
func AuthMiddleware(authService *clerk.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization ヘッダーを取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
				"code":  "MISSING_AUTH_HEADER",
			})
			c.Abort()
			return
		}

		// Bearer トークンの形式チェック
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format",
				"code":  "INVALID_AUTH_FORMAT",
			})
			c.Abort()
			return
		}

		// トークンを抽出
		token := strings.TrimPrefix(authHeader, bearerPrefix)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is empty",
				"code":  "EMPTY_TOKEN",
			})
			c.Abort()
			return
		}

		// Clerk JWT検証を実行
		user, err := authService.VerifyToken(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
				"code":  "INVALID_TOKEN",
				"details": err.Error(),
			})
			c.Abort()
			return
		}

		// ユーザー情報をコンテキストに設定
		c.Set("user", user)
		c.Set("user_id", user.ID)

		c.Next()
	}
}

// OptionalAuthMiddleware は任意認証ミドルウェア（認証がなくてもリクエストを続行）
func OptionalAuthMiddleware(authService *clerk.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// 認証情報がない場合は匿名ユーザーとして続行
			c.Next()
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			c.Next()
			return
		}

		token := strings.TrimPrefix(authHeader, bearerPrefix)
		if token == "" {
			c.Next()
			return
		}

		// Clerk JWT検証を試行
		// 認証に成功した場合のみユーザー情報を設定
		user, err := authService.VerifyToken(c.Request.Context(), token)
		if err == nil {
			c.Set("user", user)
			c.Set("user_id", user.ID)
		}

		c.Next()
	}
}

// GetUserID はコンテキストからユーザーIDを取得
func GetUserID(c *gin.Context) (string, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return "", false
	}
	
	id, ok := userID.(string)
	return id, ok
}

// GetUser はコンテキストからユーザー情報を取得
func GetUser(c *gin.Context) (*clerk.ClerkUser, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	
	clerkUser, ok := user.(*clerk.ClerkUser)
	return clerkUser, ok
}

// RequireUserID はユーザーIDが必要な場合のヘルパー
func RequireUserID(c *gin.Context) (string, bool) {
	userID, exists := GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID is required",
			"code":  "USER_ID_REQUIRED",
		})
		return "", false
	}
	return userID, true
}