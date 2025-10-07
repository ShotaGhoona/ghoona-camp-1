package router

import (
	"ghoona-camp-backend/internal/interface/middleware"

	"github.com/gin-gonic/gin"
)

// setupEventRoutes イベント管理関連のルートを設定する
func (r *Router) setupEventRoutes(v1 *gin.RouterGroup) {
	// Clerk認証ミドルウェア
	auth := middleware.AuthMiddleware(r.container.ClerkService)
	optionalAuth := middleware.OptionalAuthMiddleware(r.container.ClerkService)

	// イベント管理エンドポイント
	eventsGroup := v1.Group("/events")
	{
		// 認証任意（公開情報）
		eventsGroup.GET("", optionalAuth, r.container.EventController.GetEvents)
		eventsGroup.GET("/:eventId", optionalAuth, r.container.EventController.GetEventByID)
		eventsGroup.GET("/:eventId/participants", optionalAuth, r.container.EventController.GetEventParticipants)

		// 認証必須（作成・編集・削除・参加）
		eventsGroup.Use(auth)
		eventsGroup.POST("", r.container.EventController.PostEvents)
		eventsGroup.PUT("/:eventId", r.container.EventController.PutEventByID)
		eventsGroup.DELETE("/:eventId", r.container.EventController.DeleteEventByID)
		eventsGroup.POST("/:eventId/participants", r.container.EventController.PostEventParticipants)
		eventsGroup.PUT("/:eventId/participants/:userId", r.container.EventController.PutEventParticipant)
	}
}
