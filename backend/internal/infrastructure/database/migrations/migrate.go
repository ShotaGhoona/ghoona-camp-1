package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

// RunMigrations は自動マイグレーションを実行
func RunMigrations(db *gorm.DB) error {
	// TODO: BE-03-*で各ドメインのモデルを追加時に実装
	// 現在は基盤のみ実装
	
	if err := createMigrationTable(db); err != nil {
		return fmt.Errorf("🚨 マイグレーションテーブルの作成に失敗しました: %w", err)
	}
	
	return nil
}

// createMigrationTable はマイグレーション管理用テーブルを作成
func createMigrationTable(db *gorm.DB) error {
	migrationSQL := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version VARCHAR(255) PRIMARY KEY,
		applied_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`
	
	if err := db.Exec(migrationSQL).Error; err != nil {
		return fmt.Errorf("🚨 schema_migrationsテーブルの作成に失敗しました: %w", err)
	}
	
	var count int64
	db.Raw("SELECT COUNT(*) FROM schema_migrations WHERE version = ?", "001_initial").Scan(&count)
	
	if count == 0 {
		if err := db.Exec("INSERT INTO schema_migrations (version) VALUES (?)", "001_initial").Error; err != nil {
			return fmt.Errorf("🚨 初期マイグレーションレコードの挿入に失敗しました: %w", err)
		}
	}
	
	return nil
}

