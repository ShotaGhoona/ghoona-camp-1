package attendance

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/attendance/entity"
	"ghoona-camp-backend/internal/domain/attendance/repository"
	attendanceModel "ghoona-camp-backend/internal/infrastructure/gorm/model/attendance"
)

type attendanceSummaryRepository struct {
	db *gorm.DB
}

// NewAttendanceSummaryRepository コンストラクタ
func NewAttendanceSummaryRepository(db *gorm.DB) repository.AttendanceSummaryRepository {
	return &attendanceSummaryRepository{db: db}
}

// Create 参加サマリーを作成する
func (r *attendanceSummaryRepository) Create(ctx context.Context, summary *entity.AttendanceSummary) error {
	gormSummary := r.toGORMAttendanceSummary(summary)
	return r.db.WithContext(ctx).Create(gormSummary).Error
}

// GetByID IDで参加サマリーを取得する
func (r *attendanceSummaryRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.AttendanceSummary, error) {
	var gormSummary attendanceModel.AttendanceSummary
	err := r.db.WithContext(ctx).First(&gormSummary, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMAttendanceSummary(&gormSummary)
}

// GetByUserAndDate ユーザーIDと日付でサマリーを取得する
func (r *attendanceSummaryRepository) GetByUserAndDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.AttendanceSummary, error) {
	var gormSummary attendanceModel.AttendanceSummary
	err := r.db.WithContext(ctx).Where("user_id = ? AND DATE(summary_date) = DATE(?)", userID, date).First(&gormSummary).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMAttendanceSummary(&gormSummary)
}

// GetByUserID ユーザーIDでサマリー一覧を取得する
func (r *attendanceSummaryRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.AttendanceSummary, error) {
	var gormSummaries []attendanceModel.AttendanceSummary
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&gormSummaries).Error
	if err != nil {
		return nil, err
	}

	summaries := make([]*entity.AttendanceSummary, len(gormSummaries))
	for i, gormSummary := range gormSummaries {
		summary, err := r.fromGORMAttendanceSummary(&gormSummary)
		if err != nil {
			return nil, err
		}
		summaries[i] = summary
	}
	return summaries, nil
}

// GetByDateRange 日付範囲でサマリーを取得する
func (r *attendanceSummaryRepository) GetByDateRange(ctx context.Context, userID uuid.UUID, dateFrom, dateTo time.Time) ([]*entity.AttendanceSummary, error) {
	var gormSummaries []attendanceModel.AttendanceSummary
	err := r.db.WithContext(ctx).Where("user_id = ? AND summary_date >= ? AND summary_date <= ?", userID, dateFrom, dateTo).Find(&gormSummaries).Error
	if err != nil {
		return nil, err
	}

	summaries := make([]*entity.AttendanceSummary, len(gormSummaries))
	for i, gormSummary := range gormSummaries {
		summary, err := r.fromGORMAttendanceSummary(&gormSummary)
		if err != nil {
			return nil, err
		}
		summaries[i] = summary
	}
	return summaries, nil
}

// Update 参加サマリーを更新する
func (r *attendanceSummaryRepository) Update(ctx context.Context, summary *entity.AttendanceSummary) error {
	gormSummary := r.toGORMAttendanceSummary(summary)
	return r.db.WithContext(ctx).Save(gormSummary).Error
}

// Delete 参加サマリーを削除する
func (r *attendanceSummaryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&attendanceModel.AttendanceSummary{}, "id = ?", id).Error
}

// Upsert サマリーを作成または更新する
func (r *attendanceSummaryRepository) Upsert(ctx context.Context, summary *entity.AttendanceSummary) error {
	gormSummary := r.toGORMAttendanceSummary(summary)
	return r.db.WithContext(ctx).Save(gormSummary).Error
}

// 型変換: Domain Entity → GORM Model
func (r *attendanceSummaryRepository) toGORMAttendanceSummary(summary *entity.AttendanceSummary) *attendanceModel.AttendanceSummary {
	return &attendanceModel.AttendanceSummary{
		ID:                 summary.ID(),
		UserID:             summary.UserID(),
		SummaryDate:        summary.SummaryDate(),
		TotalDuration:      summary.TotalDuration(),
		EventCount:         summary.EventCount(),
		GoalCompletionRate: summary.GoalCompletionRate(),
		Notes:              summary.Notes(),
		CreatedAt:          summary.CreatedAt(),
		UpdatedAt:          summary.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *attendanceSummaryRepository) fromGORMAttendanceSummary(gormSummary *attendanceModel.AttendanceSummary) (*entity.AttendanceSummary, error) {
	// TODO: entity.NewAttendanceSummary の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}