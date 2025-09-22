package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ghoona-camp-backend/internal/infrastructure/config"
)

// NewDatabase は新しいデータベース接続を作成
func NewDatabase(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	// TODO: BE-02-arch-02で実際のSupabase接続を実装
	// 現在は基盤のみ実装

	dsn := cfg.GetDSN()
	
	// GORM設定
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(getLogLevel(cfg)),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	// PostgreSQL接続
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 接続プール設定
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// 接続プール設定
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 接続テスト
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established successfully")
	return db, nil
}

// getLogLevel は環境に応じたログレベルを取得
func getLogLevel(cfg *config.DatabaseConfig) logger.LogLevel {
	// TODO: 設定から取得するように変更
	return logger.Info
}

// CloseDatabase はデータベース接続を閉じる
func CloseDatabase(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}
	
	return sqlDB.Close()
}