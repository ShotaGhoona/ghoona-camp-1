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
		log.Println("No .env file found")
	}

	// Initialize application configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Set Gin mode based on environment
	if appConfig.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup database connection
	db, err := setupDatabase(appConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize dependency injection container
	container := di.NewContainer(db, appConfig)

	// Setup routes
	r := router.NewRouter(container)
	engine := r.Setup()

	// Create and start server
	server := config.NewServer(engine, appConfig)
	if err := server.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func setupDatabase(cfg *config.Config) (*gorm.DB, error) {
	// TODO: BE-02-arch-02で実際のSupabase接続を実装
	// 現在は基盤のみ
	return database.NewDatabase(&cfg.Database)
}