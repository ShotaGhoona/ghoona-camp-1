package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

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
		log.Println("🚀 マイグレーションを開始します...")
		
		// 基本マイグレーション実行
		if err := migrations.RunMigrations(db); err != nil {
			log.Fatal("🚨 基本マイグレーションに失敗しました:", err)
		}
		
		// SQLファイルからのマイグレーション実行
		if err := runSQLMigration(db, "001_initial_tables.sql"); err != nil {
			log.Fatal("🚨 SQLマイグレーションに失敗しました:", err)
		}
		
		log.Println("✅ マイグレーションが正常に完了しました")
		
	case "status":
		log.Println("📄 マイグレーション状態確認")
		if err := checkMigrationStatus(db); err != nil {
			log.Fatal("🚨 マイグレーション状態確認に失敗しました:", err)
		}
		
	default:
		fmt.Printf("Unknown action: %s\n", *action)
		fmt.Println("Available actions: up, status")
	}
}

// runSQLMigration はSQLファイルからマイグレーションを実行
func runSQLMigration(db interface{}, filename string) error {
	// SQLファイルのパスを取得
	sqlPath := filepath.Join("internal", "infrastructure", "gorm", "migration", filename)
	
	// SQLファイルを読み込み
	sqlBytes, err := ioutil.ReadFile(sqlPath)
	if err != nil {
		return fmt.Errorf("SQLファイルの読み込みに失敗しました: %w", err)
	}
	
	// TODO: BE-03-*でGORMモデル実装時に実際のマイグレーション実行
	log.Printf("📄 SQLファイル読み込み完了: %s (%d bytes)", filename, len(sqlBytes))
	log.Println("💡 TODO: BE-03-*でGORMモデル実装時に実際のSQL実行予定")
	
	return nil
}

// checkMigrationStatus はマイグレーション状態を確認
func checkMigrationStatus(db interface{}) error {
	// TODO: BE-03-*でGORMモデル実装時に実際の状態確認
	log.Println("📋 schema_migrations テーブルの状態確認")
	log.Println("💡 TODO: BE-03-*でGORMモデル実装時に詳細な状態確認実装予定")
	
	return nil
}