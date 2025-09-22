package database

import (
	"fmt"

	"gorm.io/gorm"

	"ghoona-camp-backend/internal/infrastructure/config"
)

// NewDatabase は新しいデータベース接続を作成
// 使用予定: BE-02-arch-02で実際のSupabase接続実装
func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	// TODO: BE-02-arch-02で実装予定
	// - 実際のSupabase接続文字列作成
	// - 接続プール設定
	// - マイグレーション実行
	
	// 現在は基盤のみ - 実際の接続は後続タスクで実装
	return nil, fmt.Errorf("database connection not implemented yet - will be implemented in BE-02-arch-02")
}