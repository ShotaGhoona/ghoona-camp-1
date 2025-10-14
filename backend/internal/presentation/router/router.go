package router

import (
	"ghoona-camp-backend/internal/presentation/controller"
	"ghoona-camp-backend/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 適切なミドルウェアを使用して全APIルートを設定する
func SetupRoutes(
	userController *controller.UserController,
	goalController *controller.GoalController,
	titleController *controller.TitleController,
	attendanceController *controller.AttendanceController,
	eventController *controller.EventController,
	notificationController *controller.NotificationController,
) *gin.Engine {
	r := gin.Default()

	// グローバルミドルウェア
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// APIバージョニング
	v1 := r.Group("/v1")

	// パブリックルート（認証不要）
	setupPublicRoutes(v1)

	// 認証必須ルート
	auth := v1.Group("/")
	auth.Use(middleware.ClerkAuth())
	setupAuthenticatedRoutes(auth, userController, goalController, titleController, attendanceController, eventController, notificationController)

	return r
}

// setupPublicRoutes 認証が不要なルートを設定する
func setupPublicRoutes(rg *gin.RouterGroup) {
	// システムAPI - ヘルスチェック
	rg.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"timestamp": "2025-01-10T00:00:00Z",
		})
	})
}

// setupAuthenticatedRoutes 認証が必要なルートを設定する
func setupAuthenticatedRoutes(
	rg *gin.RouterGroup,
	userController *controller.UserController,
	goalController *controller.GoalController,
	titleController *controller.TitleController,
	attendanceController *controller.AttendanceController,
	eventController *controller.EventController,
	notificationController *controller.NotificationController,
) {
	// ユーザー管理ルート
	setupUserRoutes(rg, userController)

	// 目標管理ルート
	setupGoalRoutes(rg, goalController)

	// 称号管理ルート
	setupTitleRoutes(rg, titleController)

	// 参加記録管理ルート
	setupAttendanceRoutes(rg, attendanceController)

	// イベント管理ルート
	setupEventRoutes(rg, eventController)

	// 通知管理ルート
	setupNotificationRoutes(rg, notificationController)
}