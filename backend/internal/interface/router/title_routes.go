package router

import (
	"ghoona-camp-backend/internal/interface/middleware"

	"github.com/gin-gonic/gin"
)

// setupTitleRoutes 称号管理関連のルートを設定する
func (r *Router) setupTitleRoutes(v1 *gin.RouterGroup) {
	// Clerk認証ミドルウェア
	auth := middleware.AuthMiddleware(r.container.ClerkService)
	
	// 称号一覧・詳細（認証必要）
	titlesGroup := v1.Group("/titles")
	titlesGroup.Use(auth)
	{
		titlesGroup.GET("", r.container.TitleController.GetTitles)
		titlesGroup.GET("/:titleId", r.container.TitleController.GetTitle)
	}
	
	// ユーザー称号管理（既存usersGroupに追加）
	usersGroup := v1.Group("/users")
	usersGroup.Use(auth)
	{
		// 称号獲得履歴
		usersGroup.GET("/:userId/achievements", r.container.TitleAchievementController.GetUserAchievements)
		
		// 現在称号変更
		usersGroup.PUT("/:userId/achievements/:titleId", r.container.TitleAchievementController.SetCurrentTitle)
	}
}