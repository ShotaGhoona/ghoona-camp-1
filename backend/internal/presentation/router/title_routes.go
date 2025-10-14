package router

import (
	"ghoona-camp-backend/internal/presentation/controller"

	"github.com/gin-gonic/gin"
)

// setupTitleRoutes 称号管理関連の全ルートを設定する
func setupTitleRoutes(rg *gin.RouterGroup, titleController *controller.TitleController) {
	// GET /titles - 全称号一覧を取得（8レベル）
	// アクセス: 🔐 認証済みユーザー
	// 返却値: level, name_jp, name_en, description, required_days, image_url, color_theme
	rg.GET("/titles", titleController.GetTitles)

	// GET /titles/{titleId} - 指定称号の詳細情報を取得
	// アクセス: 🔐 認証済みユーザー
	// 返却値: 称号情報 + holder_count, rarity_percentage
	rg.GET("/titles/:titleId", titleController.GetTitleDetail)
}