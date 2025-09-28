package router

import (
	"ghoona-camp-backend/internal/interface/middleware"

	"github.com/gin-gonic/gin"
)

// setupUserRoutes ユーザー管理関連のルートを設定する
func (r *Router) setupUserRoutes(v1 *gin.RouterGroup) {
	// Clerk認証ミドルウェア
	auth := middleware.AuthMiddleware(r.container.ClerkService)
	
	// 認証エンドポイント
	authGroup := v1.Group("/auth")
	authGroup.Use(auth)
	{
		authGroup.GET("/me", r.container.UserController.GetCurrentUser)
	}
	
	// ユーザー管理エンドポイント
	usersGroup := v1.Group("/users")
	usersGroup.Use(auth)
	{
		// ユーザー基本操作
		usersGroup.GET("/:userId", r.container.UserController.GetUser)
		usersGroup.PUT("/:userId", r.container.UserController.UpdateUser)
		
		// ユーザーメタデータ
		usersGroup.GET("/:userId/metadata", r.container.UserController.GetUserMetadata)
		usersGroup.POST("/:userId/metadata", r.container.UserController.CreateUserMetadata)
		usersGroup.PUT("/:userId/metadata", r.container.UserController.UpdateUserMetadata)
		
		// ソーシャルリンク
		usersGroup.GET("/:userId/social-links", r.container.UserController.GetUserSocialLinks)
		usersGroup.POST("/:userId/social-links", r.container.UserController.CreateSocialLink)
		usersGroup.PUT("/:userId/social-links/:linkId", r.container.UserController.UpdateSocialLink)
		usersGroup.DELETE("/:userId/social-links/:linkId", r.container.UserController.DeleteSocialLink)
		
		// ライバル管理
		usersGroup.GET("/:userId/rivals", r.container.UserController.GetUserRivals)
		usersGroup.POST("/:userId/rivals", r.container.UserController.AddRival)
		usersGroup.DELETE("/:userId/rivals/:rivalId", r.container.UserController.RemoveRival)
	}
}