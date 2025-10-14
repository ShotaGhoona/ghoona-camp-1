package router

import (
	"ghoona-camp-backend/internal/presentation/controller"
	"ghoona-camp-backend/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

// setupEventRoutes イベント管理関連の全ルートを設定する
func setupEventRoutes(rg *gin.RouterGroup, eventController *controller.EventController) {
	// GET /events - 開催予定/開催中のイベント一覧を取得
	// アクセス: 🔐 認証済みユーザー
	// クエリパラメータ: status (upcoming/ongoing/past), creator, date_from, date_to, limit, offset
	rg.GET("/events", eventController.GetEvents)

	// POST /events - 新しい朝活イベントを作成
	// アクセス: 🔐 認証済みユーザー（イベント作成者になる）
	// ボディ: title, description, event_type, scheduled_date, start_time, end_time, max_participants, is_recurring, recurrence_pattern, discord_channel_id
	rg.POST("/events", eventController.PostEvent)

	// GET /events/{eventId} - 指定イベントの詳細情報を取得
	// アクセス: 🔐 認証済みユーザー
	// 返却値: イベント情報 + 作成者情報 + 参加者一覧
	rg.GET("/events/:eventId", eventController.GetEventDetail)

	// イベント作成者のみのルート
	eventCreator := rg.Group("/events/:eventId")
	eventCreator.Use(middleware.EventCreatorOnly())
	{
		// PUT /events/{eventId} - イベント情報を更新
		// アクセス: 👑 イベント作成者のみ
		// ボディ: title, description, max_participants, is_active
		eventCreator.PUT("", eventController.PutEvent)

		// DELETE /events/{eventId} - イベントを削除
		// アクセス: 👑 イベント作成者のみ
		eventCreator.DELETE("", eventController.DeleteEvent)
	}

	// イベント参加ルート
	participation := rg.Group("/events/:eventId/participants")
	{
		// POST /events/{eventId}/participants - イベントに参加登録
		// アクセス: 🔐 認証済みユーザー
		// 自動的に現在のユーザーIDを使用
		participation.POST("", eventController.PostEventParticipant)

		// ユーザー固有の参加ルート（参加者またはイベント作成者のみ）
		userParticipation := participation.Group("/:userId")
		userParticipation.Use(middleware.ParticipantOrCreatorOnly())
		{
			// PUT /events/{eventId}/participants/{userId} - 参加ステータスを更新
			// アクセス: 👤 参加者本人またはイベント作成者
			// ボディ: status (registered/cancelled)
			userParticipation.PUT("", eventController.PutEventParticipant)
		}
	}
}