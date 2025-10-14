package router

import (
	"ghoona-camp-backend/internal/presentation/controller"

	"github.com/gin-gonic/gin"
)

// setupAttendanceRoutes 参加記録管理関連の全ルートを設定する
func setupAttendanceRoutes(rg *gin.RouterGroup, attendanceController *controller.AttendanceController) {
	// ランキングルート（認証済みユーザーにパブリックアクセス）
	ranking := rg.Group("/ranking")
	{
		// GET /ranking/monthly - 月次参加ランキングを取得
		// アクセス: 🔐 認証済みユーザー
		// クエリパラメータ: month (1-12, デフォルト: 現在月), limit, offset
		ranking.GET("/monthly", attendanceController.GetRankingMonthly)

		// GET /ranking/total - 総合参加ランキングを取得
		// アクセス: 🔐 認証済みユーザー
		// クエリパラメータ: limit, offset
		ranking.GET("/total", attendanceController.GetRankingTotal)

		// GET /ranking/streak - 連続参加ランキングを取得
		// アクセス: 🔐 認証済みユーザー
		// クエリパラメータ: limit, offset
		ranking.GET("/streak", attendanceController.GetRankingStreak)
	}

	// 注意: ユーザー固有の参加記録ルートはuser_routes.goで定義されています
	// - GET /users/{userId}/attendance/summaries (所有者のみ)
	// - GET /users/{userId}/attendance/statistics (パブリックまたは所有者)
}
