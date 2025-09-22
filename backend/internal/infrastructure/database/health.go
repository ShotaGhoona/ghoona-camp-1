package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// HealthStatus はデータベースの健康状態を表す
type HealthStatus struct {
	Status      string        `json:"status"`
	Latency     time.Duration `json:"latency"`
	Connections int           `json:"connections"`
	Error       string        `json:"error,omitempty"`
}

// CheckHealth はデータベースの健康状態をチェック
func CheckHealth(db *gorm.DB) HealthStatus {
	start := time.Now()
	
	// コンテキストでタイムアウト設定
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// 基本的な接続テスト
	sqlDB, err := db.DB()
	if err != nil {
		return HealthStatus{
			Status:  "error",
			Latency: time.Since(start),
			Error:   fmt.Sprintf("failed to get sql.DB: %v", err),
		}
	}
	
	// Pingテスト
	if err := sqlDB.PingContext(ctx); err != nil {
		return HealthStatus{
			Status:  "error", 
			Latency: time.Since(start),
			Error:   fmt.Sprintf("ping failed: %v", err),
		}
	}
	
	// 接続プール統計を取得
	stats := sqlDB.Stats()
	
	// 簡単なクエリテスト
	var result int
	if err := db.WithContext(ctx).Raw("SELECT 1").Scan(&result).Error; err != nil {
		return HealthStatus{
			Status:      "degraded",
			Latency:     time.Since(start),
			Connections: stats.OpenConnections,
			Error:       fmt.Sprintf("query test failed: %v", err),
		}
	}
	
	return HealthStatus{
		Status:      "healthy",
		Latency:     time.Since(start),
		Connections: stats.OpenConnections,
	}
}

// GetConnectionStats は接続プールの詳細統計を取得
func GetConnectionStats(db *gorm.DB) (map[string]interface{}, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}
	
	stats := sqlDB.Stats()
	
	return map[string]interface{}{
		"max_open_connections":     stats.MaxOpenConnections,
		"open_connections":         stats.OpenConnections,
		"in_use":                  stats.InUse,
		"idle":                    stats.Idle,
		"wait_count":              stats.WaitCount,
		"wait_duration":           stats.WaitDuration.String(),
		"max_idle_closed":         stats.MaxIdleClosed,
		"max_idle_time_closed":    stats.MaxIdleTimeClosed,
		"max_lifetime_closed":     stats.MaxLifetimeClosed,
	}, nil
}