package router

import (
	"ghoona-camp-backend/internal/presentation/controller"
	"ghoona-camp-backend/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

// setupUserRoutes ユーザー管理関連の全ルートを設定する
func setupUserRoutes(rg *gin.RouterGroup, userController *controller.UserController) {
	// GET /auth/me - 現在認証中のユーザー情報を取得
	// アクセス: 🔐 認証済みユーザーのみ
	rg.GET("/auth/me", userController.GetAuthMe)

	// GET /users - 検索機能付きの全ユーザー一覧を取得
	// アクセス: 🔐 認証済みユーザーのみ
	// クエリパラメータ: search, limit, offset
	rg.GET("/users", userController.GetUsers)

	// GET /users/{userId} - 指定ユーザーの詳細情報を取得
	// アクセス: 🔐 認証済みユーザーのみ
	rg.GET("/users/:userId", userController.GetUserDetail)

	// ユーザー固有ルート（所有者のみアクセス可）
	userOnly := rg.Group("/users/:userId")
	userOnly.Use(middleware.UserOwnerOnly())
	{
		// PUT /users/{userId} - ユーザーの基本情報とメタデータを更新
		// アクセス: 👤 所有者のみ
		userOnly.PUT("", userController.PutUser)

		// ソーシャルリンク管理
		// POST /users/{userId}/social-links - 新しいソーシャルリンクを追加
		// アクセス: 👤 所有者のみ
		userOnly.POST("/social-links", userController.PostSocialLink)

		// PUT /users/{userId}/social-links/{linkId} - 既存のソーシャルリンクを更新
		// アクセス: 👤 所有者のみ
		userOnly.PUT("/social-links/:linkId", userController.PutSocialLink)

		// DELETE /users/{userId}/social-links/{linkId} - ソーシャルリンクを削除
		// アクセス: 👤 所有者のみ
		userOnly.DELETE("/social-links/:linkId", userController.DeleteSocialLink)

		// ライバル管理
		// GET /users/{userId}/rivals - ユーザーのライバル一覧を取得（最大3人）
		// アクセス: 👤 所有者のみ
		userOnly.GET("/rivals", userController.GetUserRivals)

		// POST /users/{userId}/rivals - 新しいライバルを追加
		// アクセス: 👤 所有者のみ
		userOnly.POST("/rivals", userController.PostRival)

		// DELETE /users/{userId}/rivals/{rivalId} - ライバル関係を解除
		// アクセス: 👤 所有者のみ
		userOnly.DELETE("/rivals/:rivalId", userController.DeleteRival)

		// 参加記録管理
		// GET /users/{userId}/attendance/summaries - 日次参加サマリーを取得
		// アクセス: 👤 所有者のみ
		// クエリパラメータ: date_from, date_to
		userOnly.GET("/attendance/summaries", userController.GetAttendanceSummaries)

		// 通知管理
		// GET /users/{userId}/notifications - ユーザーの通知一覧を取得
		// アクセス: 👤 所有者のみ
		// クエリパラメータ: type, read, limit, offset
		userOnly.GET("/notifications", userController.GetUserNotifications)

		// GET /users/{userId}/notification-settings - 通知設定を取得
		// アクセス: 👤 所有者のみ
		userOnly.GET("/notification-settings", userController.GetNotificationSettings)

		// PUT /users/{userId}/notification-settings - 通知設定を更新
		// アクセス: 👤 所有者のみ
		userOnly.PUT("/notification-settings", userController.PutNotificationSettings)
	}

	// パブリックまたは所有者アクセス可能ルート
	publicOrOwner := rg.Group("/users/:userId")
	publicOrOwner.Use(middleware.PublicOrOwnerAccess())
	{
		// GET /users/{userId}/attendance/statistics - 参加統計を取得
		// アクセス: 🔐👤📖 パブリック閲覧可または所有者
		publicOrOwner.GET("/attendance/statistics", userController.GetAttendanceStatistics)

		// 称号実績管理
		// GET /users/{userId}/achievements - ユーザーの称号実績を取得
		// アクセス: 🔐 認証済みユーザー（パブリック情報）
		publicOrOwner.GET("/achievements", userController.GetUserAchievements)
	}

	// 称号実績更新（所有者のみ）
	achievementUpdate := rg.Group("/users/:userId/achievements/:titleId")
	achievementUpdate.Use(middleware.UserOwnerOnly())
	{
		// PUT /users/{userId}/achievements/{titleId} - 現在表示する称号を設定
		// アクセス: 👤 所有者のみ
		achievementUpdate.PUT("", userController.PutUserAchievement)
	}
}