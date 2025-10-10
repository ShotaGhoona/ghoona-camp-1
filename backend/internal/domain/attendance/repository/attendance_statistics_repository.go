package repository

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/attendance/entity"
)

// AttendanceStatisticsRepository 参加統計リポジトリインターフェース
type AttendanceStatisticsRepository interface {
	// Create 参加統計を作成する
	Create(ctx context.Context, statistics *entity.AttendanceStatistics) error

	// GetByUserID ユーザーIDで統計を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.AttendanceStatistics, error)

	// Update 参加統計を更新する
	Update(ctx context.Context, statistics *entity.AttendanceStatistics) error

	// Delete 参加統計を削除する
	Delete(ctx context.Context, userID uuid.UUID) error

	// Upsert 統計を作成または更新する
	Upsert(ctx context.Context, statistics *entity.AttendanceStatistics) error

	// GetRankingByTotalDays 総参加日数ランキングを取得する
	GetRankingByTotalDays(ctx context.Context, limit, offset int) ([]*entity.AttendanceStatistics, error)

	// GetRankingByCurrentStreak 現在の連続参加日数ランキングを取得する
	GetRankingByCurrentStreak(ctx context.Context, limit, offset int) ([]*entity.AttendanceStatistics, error)

	// GetRankingByMaxStreak 最大連続参加日数ランキングを取得する
	GetRankingByMaxStreak(ctx context.Context, limit, offset int) ([]*entity.AttendanceStatistics, error)
}