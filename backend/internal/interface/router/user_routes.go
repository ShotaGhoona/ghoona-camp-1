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
		// ユーザー一覧・作成
		usersGroup.GET("", r.container.UserController.GetUsers)
		usersGroup.POST("", r.container.UserController.CreateUser)
		
		// ユーザー基本操作
		usersGroup.GET("/:userId", r.container.UserController.GetUser)
		usersGroup.PUT("/:userId", r.container.UserController.UpdateUser)
		
		// ユーザーメタデータ
		usersGroup.GET("/:userId/metadata", r.container.UserMetadataController.GetUserMetadata)
		usersGroup.POST("/:userId/metadata", r.container.UserMetadataController.CreateUserMetadata)
		usersGroup.PUT("/:userId/metadata", r.container.UserMetadataController.UpdateUserMetadata)
		
		// ソーシャルリンク
		usersGroup.GET("/:userId/social-links", r.container.UserSocialController.GetUserSocialLinks)
		usersGroup.POST("/:userId/social-links", r.container.UserSocialController.CreateSocialLink)
		usersGroup.PUT("/:userId/social-links/:linkId", r.container.UserSocialController.UpdateSocialLink)
		usersGroup.DELETE("/:userId/social-links/:linkId", r.container.UserSocialController.DeleteSocialLink)
		
		// ライバル管理
		usersGroup.GET("/:userId/rivals", r.container.UserRivalController.GetUserRivals)
		usersGroup.POST("/:userId/rivals", r.container.UserRivalController.AddRival)
		usersGroup.DELETE("/:userId/rivals/:rivalId", r.container.UserRivalController.RemoveRival)
	}
}