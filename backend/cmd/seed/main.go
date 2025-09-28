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
)

func main() {
	// フラグ定義
	var action = flag.String("action", "run", "Seed action: run, clean")
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

	// シード実行
	switch *action {
	case "run":
		log.Println("🌱 シードデータの投入を開始します...")
		
		if err := runSeedData(db); err != nil {
			log.Fatal("🚨 シードデータの投入に失敗しました:", err)
		}
		
		log.Println("✅ シードデータの投入が正常に完了しました")
		
	case "clean":
		log.Println("🧹 シードデータのクリーンアップ")
		// TODO: BE-03-*でGORMモデル実装時にクリーンアップ機能実装
		log.Println("💡 TODO: BE-03-*でGORMモデル実装時にクリーンアップ機能実装予定")
		
	default:
		fmt.Printf("Unknown action: %s\n", *action)
		fmt.Println("Available actions: run, clean")
	}
}

// runSeedData はシードデータを投入
func runSeedData(db interface{}) error {
	// シードSQLファイルのパスを取得
	sqlPath := filepath.Join("internal", "infrastructure", "gorm", "migration", "seed_data.sql")
	
	// SQLファイルを読み込み
	sqlBytes, err := ioutil.ReadFile(sqlPath)
	if err != nil {
		return fmt.Errorf("シードSQLファイルの読み込みに失敗しました: %w", err)
	}
	
	// TODO: BE-03-*でGORMモデル実装時に実際のシードデータ投入
	log.Printf("📄 シードSQLファイル読み込み完了: seed_data.sql (%d bytes)", len(sqlBytes))
	log.Println("💡 TODO: BE-03-*でGORMモデル実装時に実際のSQL実行予定")
	log.Println("📊 投入予定データ: 称号8件、テストユーザー2件、イベント2件、統計2件、称号実績2件")
	
	return nil
}