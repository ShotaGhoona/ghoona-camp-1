package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"ghoona-camp-backend/internal/infrastructure/clerk"
)

// TODO: BE-02-arch-02で実際のClerk認証を実装
// 現在は基盤のみ実装

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

		// TODO: 実際のClerk JWT検証を実装
		// 現在は基盤のみ
		/*
		user, err := authService.VerifyToken(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
				"code":  "INVALID_TOKEN",
			})
			c.Abort()
			return
		}

		// ユーザー情報をコンテキストに設定
		c.Set("user", user)
		c.Set("user_id", user.ID)
		*/

		// 仮実装: トークンが "test-token" の場合は認証成功とする
		if token == "test-token" {
			c.Set("user_id", "test-user-id")
			c.Set("user", map[string]interface{}{
				"id":       "test-user-id",
				"email":    "test@example.com",
				"username": "testuser",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
				"code":  "INVALID_TOKEN",
			})
			c.Abort()
			return
		}

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

		// TODO: 実際のClerk JWT検証を実装
		// 認証に成功した場合のみユーザー情報を設定
		if token == "test-token" {
			c.Set("user_id", "test-user-id")
			c.Set("user", map[string]interface{}{
				"id":       "test-user-id",
				"email":    "test@example.com",
				"username": "testuser",
			})
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
func GetUser(c *gin.Context) (map[string]interface{}, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	
	userMap, ok := user.(map[string]interface{})
	return userMap, ok
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