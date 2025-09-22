package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/di"
	"ghoona-camp-backend/internal/infrastructure/config"
	"ghoona-camp-backend/internal/infrastructure/database"
	"ghoona-camp-backend/internal/interface/router"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚙️ .envファイルが見つかりません（環境変数から取得します）")
	}

	// Initialize application configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("🚨 設定の読み込みに失敗しました:", err)
	}

	// Set Gin mode based on environment
	if appConfig.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup database connection
	db, err := setupDatabase(appConfig)
	if err != nil {
		log.Fatal("🚨 データベースへの接続に失敗しました:", err)
	}

	// Initialize dependency injection container
	container := di.NewContainer(db, appConfig)

	// Setup routes
	r := router.NewRouter(container)
	engine := r.Setup()

	// Create and start server
	server := config.NewServer(engine, appConfig)
	if err := server.Start(); err != nil {
		log.Fatal("🚨 サーバーの起動に失敗しました:", err)
	}
}

func setupDatabase(cfg *config.Config) (*gorm.DB, error) {
	return database.NewDatabase(cfg)
}