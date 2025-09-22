package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	
	"ghoona-camp-backend/internal/infrastructure/config"
	"ghoona-camp-backend/internal/interface/middleware"
)

func main() {
	// 設定読み込み
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("設定の読み込みに失敗しました:", err)
	}

	// データベース接続（TODO: BE-02-arch-02で実装）
	// db, err := database.NewDatabase(&cfg.Database)
	// if err != nil {
	//     log.Fatal("データベース接続に失敗しました:", err)
	// }
	// defer database.CloseDatabase(db)

	// Ginエンジン初期化
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}
	
	r := gin.New()

	// ミドルウェア設定
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorHandlerMiddleware())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.SecurityHeadersMiddleware())

	// ヘルスチェックエンドポイント
	r.GET("/health", func(c *gin.Context) {
		// TODO: データベース接続チェックを追加
		healthStatus := gin.H{
			"status":      "ok",
			"service":     "ghoona-camp-backend",
			"version":     "1.0.0",
			"environment": cfg.Server.Environment,
			"timestamp":   time.Now().UTC().Format(time.RFC3339),
		}

		// TODO: データベース接続チェック
		// if err := db.DB().Ping(); err != nil {
		//     healthStatus["status"] = "error"
		//     healthStatus["database"] = "disconnected"
		//     c.JSON(503, healthStatus)
		//     return
		// }
		// healthStatus["database"] = "connected"

		c.JSON(200, healthStatus)
	})

	// メトリクスエンドポイント（TODO: 後で詳細実装）
	r.GET("/metrics", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"metrics": "TODO: implement metrics",
		})
	})

	// API v1 ルートグループ
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
				"timestamp": time.Now().UTC().Format(time.RFC3339),
			})
		})
		
		// TODO: BE-03以降で各ドメインのルートを追加
		// v1.Use(middleware.AuthMiddleware(authService))
		// 
		// users := v1.Group("/users")
		// users.GET("", userController.GetUsers)
		// 
		// goals := v1.Group("/goals")
		// goals.GET("", goalController.GetGoals)
		//
		// events := v1.Group("/events")
		// events.GET("", eventController.GetEvents)
	}

	// 404ハンドラー
	r.NoRoute(middleware.NotFoundHandler())
	r.NoMethod(middleware.MethodNotAllowedHandler())

	// サーバー設定
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// グレースフルシャットダウンのためのgoroutine
	go func() {
		log.Printf("Ghoona Camp Backendサーバーをポート%sで起動します", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("サーバーの起動に失敗しました:", err)
		}
	}()

	// シャットダウンシグナルを待機
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("サーバーをシャットダウンしています...")

	// グレースフルシャットダウン
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("サーバーのシャットダウンに失敗しました:", err)
	}
	
	log.Println("サーバーが正常にシャットダウンされました")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}