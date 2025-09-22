package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Server represents HTTP server configuration
// 使用予定: main.goでサーバー起動・停止制御
type Server struct {
	engine *gin.Engine
	config *Config
}

// NewServer creates a new server instance
// 使用予定: main.goでサーバーインスタンス作成時
func NewServer(engine *gin.Engine, config *Config) *Server {
	return &Server{
		engine: engine,
		config: config,
	}
}

// Start starts the HTTP server with graceful shutdown
// 使用予定: main.goでサーバー起動時
func (s *Server) Start() error {
	srv := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.engine,
	}

	// Server run in goroutine
	go func() {
		log.Printf("🚀 Ghoona Camp バックエンドサーバーをポート%sで起動します", s.config.Port)
		log.Printf("⚙️ 環境: %s", s.config.Env)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("🚨 サーバーの起動に失敗しました: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("⏹️ サーバーを停止します...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("🚨 サーバーを強制的に停止できませんでした: %w", err)
	}

	log.Println("✅ サーバーが停止しました")
	return nil
}
