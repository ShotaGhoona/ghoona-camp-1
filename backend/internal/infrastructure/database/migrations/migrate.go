package migrations

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// RunMigrations は自動マイグレーションを実行
func RunMigrations(db *gorm.DB) error {
	log.Println("Starting database migrations...")
	
	// TODO: BE-03-*で各ドメインのモデルを追加時に実装
	// 現在は基盤のみ実装
	
	// マイグレーション状態確認用のテーブルを作成
	if err := createMigrationTable(db); err != nil {
		return fmt.Errorf("failed to create migration table: %w", err)
	}
	
	log.Println("Database migrations completed successfully")
	return nil
}

// createMigrationTable はマイグレーション管理用テーブルを作成
func createMigrationTable(db *gorm.DB) error {
	// マイグレーション履歴テーブル
	migrationSQL := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version VARCHAR(255) PRIMARY KEY,
		applied_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`
	
	if err := db.Exec(migrationSQL).Error; err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}
	
	// 初期バージョンを記録
	var count int64
	db.Raw("SELECT COUNT(*) FROM schema_migrations WHERE version = ?", "001_initial").Scan(&count)
	
	if count == 0 {
		if err := db.Exec("INSERT INTO schema_migrations (version) VALUES (?)", "001_initial").Error; err != nil {
			return fmt.Errorf("failed to insert initial migration record: %w", err)
		}
		log.Println("Initial migration record created")
	}
	
	return nil
}

// GetMigrationStatus はマイグレーション状態を取得
func GetMigrationStatus(db *gorm.DB) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	
	// schema_migrationsテーブルの存在確認
	var tableExists bool
	err := db.Raw(`
		SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_schema = 'public' 
			AND table_name = 'schema_migrations'
		)
	`).Scan(&tableExists).Error
	
	if err != nil {
		return nil, fmt.Errorf("failed to check migration table existence: %w", err)
	}
	
	if !tableExists {
		return results, nil
	}
	
	// マイグレーション履歴を取得
	rows, err := db.Raw("SELECT version, applied_at FROM schema_migrations ORDER BY applied_at").Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to get migration status: %w", err)
	}
	defer rows.Close()
	
	for rows.Next() {
		var version string
		var appliedAt string
		if err := rows.Scan(&version, &appliedAt); err != nil {
			return nil, fmt.Errorf("failed to scan migration row: %w", err)
		}
		
		results = append(results, map[string]interface{}{
			"version":    version,
			"applied_at": appliedAt,
		})
	}
	
	return results, nil
}