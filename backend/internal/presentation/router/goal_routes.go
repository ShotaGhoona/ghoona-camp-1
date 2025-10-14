package router

import (
	"ghoona-camp-backend/internal/presentation/controller"
	"ghoona-camp-backend/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

// setupGoalRoutes 目標管理関連の全ルートを設定する
func setupGoalRoutes(rg *gin.RouterGroup, goalController *controller.GoalController) {
	// GET /goals/me - 現在のユーザーの目標を取得（プライベート + パブリック）
	// アクセス: 👤 所有者のみ
	// クエリパラメータ: search, is_public, started_from, started_to, ended_from, ended_to, limit, offset
	meGoals := rg.Group("/goals/me")
	meGoals.Use(middleware.UserContextRequired())
	{
		meGoals.GET("", goalController.GetGoalsMe)
	}

	// GET /goals/public - 全ユーザーの公開目標を取得
	// アクセス: 🔐 認証済みユーザー
	// クエリパラメータ: search, user, started_from, started_to, ended_from, ended_to, limit, offset
	rg.GET("/goals/public", goalController.GetGoalsPublic)

	// POST /goals - 新しい目標を作成
	// アクセス: 👤 認証済みユーザー（自分用の目標を作成）
	// ボディ: user_id, title, description, started_at, ended_at, is_public
	rg.POST("/goals", goalController.PostGoal)

	// 目標固有ルート（作成者のみアクセス可）
	goalOnly := rg.Group("/goals/:goalId")
	goalOnly.Use(middleware.GoalCreatorOnly())
	{
		// PUT /goals/{goalId} - 目標の内容を更新
		// アクセス: 👤 目標作成者のみ
		// ボディ: title, description, ended_at, is_public, is_active
		goalOnly.PUT("", goalController.PutGoal)

		// DELETE /goals/{goalId} - 目標を削除
		// アクセス: 👤 目標作成者のみ
		goalOnly.DELETE("", goalController.DeleteGoal)
	}
}