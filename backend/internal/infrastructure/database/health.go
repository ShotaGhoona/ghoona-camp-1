package database

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// HealthStatus はデータベースの健康状態を表す
type HealthStatus struct {
	Status  string `json:"status"`
	Latency string `json:"latency"`
	Error   string `json:"error,omitempty"`
}

// CheckHealth はデータベースの健康状態をチェック
func CheckHealth(db *gorm.DB) HealthStatus {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlDB, err := db.DB()
	if err != nil {
		return HealthStatus{
			Status: "error",
			Error:  err.Error(),
		}
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return HealthStatus{
			Status: "error",
			Error:  err.Error(),
		}
	}

	return HealthStatus{
		Status:  "ok",
		Latency: time.Since(start).String(),
	}
}