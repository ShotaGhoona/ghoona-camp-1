package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/attendance/entity"
)

// AttendanceSummaryRepository 参加サマリーリポジトリインターフェース
type AttendanceSummaryRepository interface {
	// Create 参加サマリーを作成する
	Create(ctx context.Context, summary *entity.AttendanceSummary) error

	// GetByID IDで参加サマリーを取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.AttendanceSummary, error)

	// GetByUserAndDate ユーザーIDと日付でサマリーを取得する
	GetByUserAndDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.AttendanceSummary, error)

	// GetByUserID ユーザーIDでサマリー一覧を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.AttendanceSummary, error)

	// GetByDateRange 日付範囲でサマリーを取得する
	GetByDateRange(ctx context.Context, userID uuid.UUID, dateFrom, dateTo time.Time) ([]*entity.AttendanceSummary, error)

	// Update 参加サマリーを更新する
	Update(ctx context.Context, summary *entity.AttendanceSummary) error

	// Delete 参加サマリーを削除する
	Delete(ctx context.Context, id uuid.UUID) error

	// Upsert サマリーを作成または更新する
	Upsert(ctx context.Context, summary *entity.AttendanceSummary) error
}