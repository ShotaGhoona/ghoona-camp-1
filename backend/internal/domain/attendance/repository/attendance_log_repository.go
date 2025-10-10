package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/attendance/entity"
)

// AttendanceLogRepository 参加ログリポジトリインターフェース
type AttendanceLogRepository interface {
	// Create 参加ログを作成する
	Create(ctx context.Context, log *entity.AttendanceLog) error

	// GetByID IDで参加ログを取得する
	GetByID(ctx context.Context, id uuid.UUID) (*entity.AttendanceLog, error)

	// GetByUserID ユーザーIDで参加ログ一覧を取得する
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.AttendanceLog, error)

	// GetByDateRange 日付範囲で参加ログを取得する
	GetByDateRange(ctx context.Context, userID uuid.UUID, dateFrom, dateTo time.Time) ([]*entity.AttendanceLog, error)

	// GetByEventID イベントIDで参加ログを取得する
	GetByEventID(ctx context.Context, eventID uuid.UUID) ([]*entity.AttendanceLog, error)

	// Update 参加ログを更新する
	Update(ctx context.Context, log *entity.AttendanceLog) error

	// Delete 参加ログを削除する
	Delete(ctx context.Context, id uuid.UUID) error
}