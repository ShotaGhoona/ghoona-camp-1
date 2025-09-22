package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"ghoona-camp-backend/internal/infrastructure/config"
	"ghoona-camp-backend/internal/infrastructure/database"
	"ghoona-camp-backend/internal/infrastructure/database/migrations"
)

func main() {
	// フラグ定義
	var action = flag.String("action", "up", "Migration action: up, down, status")
	flag.Parse()

	// 環境変数読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("⚙️ .envファイルが見つかりません（環境変数から取得します）")
	}

	// 設定読み込み
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("🚨 設定の読み込みに失敗しました:", err)
	}

	// データベース接続
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("🚨 データベースへの接続に失敗しました:", err)
	}

	// マイグレーション実行
	switch *action {
	case "up":
		if err := migrations.RunMigrations(db); err != nil {
			log.Fatal("🚨 マイグレーションに失敗しました:", err)
		}
		log.Println("✅ マイグレーションが正常に完了しました")
	case "status":
		log.Println("📄 マイグレーション状態確認 - TODO: 実装予定")
	default:
		fmt.Printf("Unknown action: %s\n", *action)
		fmt.Println("Available actions: up, status")
	}
}