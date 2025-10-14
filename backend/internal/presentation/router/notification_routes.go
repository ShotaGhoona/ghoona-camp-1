package router

import (
	"ghoona-camp-backend/internal/presentation/controller"
	"ghoona-camp-backend/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

// setupNotificationRoutes 通知管理関連の全ルートを設定する
func setupNotificationRoutes(rg *gin.RouterGroup, notificationController *controller.NotificationController) {
	// 通知固有ルート（所有者のみアクセス可）
	notificationOnly := rg.Group("/notifications/:notificationId")
	notificationOnly.Use(middleware.NotificationOwnerOnly())
	{
		// PUT /notifications/{notificationId} - 通知の既読ステータスを更新
		// アクセス: 👤 通知所有者のみ
		// ボディ: 既読ステータスまたはその他の通知プロパティ
		notificationOnly.PUT("", notificationController.PutNotification)

		// DELETE /notifications/{notificationId} - 不要な通知を削除
		// アクセス: 👤 通知所有者のみ
		notificationOnly.DELETE("", notificationController.DeleteNotification)
	}

	// 注意: ユーザー固有の通知ルートはuser_routes.goで定義されています
	// - GET /users/{userId}/notifications (所有者のみ)
	// - GET /users/{userId}/notification-settings (所有者のみ)
	// - PUT /users/{userId}/notification-settings (所有者のみ)
}