package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ghoona-camp-backend/internal/infrastructure/config"
)

// NewDatabase は新しいデータベース接続を作成
func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	if cfg.Database.URL == "" {
		return nil, fmt.Errorf("データベースURLが設定されていません")
	}

	// GORMログレベル設定
	logLevel := getGORMLogLevel(cfg.Database.LogLevel)

	// GORM設定
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	// PostgreSQLドライバーで接続
	db, err := gorm.Open(postgres.Open(cfg.Database.URL), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("データベース接続に失敗しました: %w", err)
	}

	// 接続プール設定
	if err := setupConnectionPool(db, cfg.Database); err != nil {
		return nil, fmt.Errorf("接続プールの設定に失敗しました: %w", err)
	}

	// 接続テスト
	if err := testConnection(db); err != nil {
		return nil, fmt.Errorf("データベース接続テストに失敗しました: %w", err)
	}

	log.Printf("データベース接続が成功しました")
	return db, nil
}

// getGORMLogLevel はstringからGORMログレベルへ変換
func getGORMLogLevel(level string) logger.LogLevel {
	switch level {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Warn
	}
}

// setupConnectionPool は接続プールを設定
func setupConnectionPool(db *gorm.DB, dbConfig config.DatabaseConfig) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("sql.DBの取得に失敗しました: %w", err)
	}

	// 接続プール設定
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)

	log.Printf("接続プールが設定されました: MaxOpen=%d, MaxIdle=%d, MaxLifetime=%v",
		dbConfig.MaxOpenConns, dbConfig.MaxIdleConns, dbConfig.ConnMaxLifetime)

	return nil
}

// testConnection はデータベース接続をテスト
func testConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("sql.DBの取得に失敗しました: %w", err)
	}

	// シンプルなPingテスト
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("pingに失敗しました: %w", err)
	}

	// 簡単なクエリテスト
	var version string
	if err := db.WithContext(ctx).Raw("SELECT version()").Scan(&version).Error; err != nil {
		return fmt.Errorf("versionクエリに失敗しました: %w", err)
	}

	log.Printf("データベース接続テストが成功しました. PostgreSQL version: %s", version)
	return nil
}
