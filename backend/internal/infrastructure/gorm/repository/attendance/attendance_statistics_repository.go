package attendance

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/attendance/entity"
	"ghoona-camp-backend/internal/domain/attendance/repository"
	attendanceModel "ghoona-camp-backend/internal/infrastructure/gorm/model/attendance"
)

type attendanceStatisticsRepository struct {
	db *gorm.DB
}

// NewAttendanceStatisticsRepository コンストラクタ
func NewAttendanceStatisticsRepository(db *gorm.DB) repository.AttendanceStatisticsRepository {
	return &attendanceStatisticsRepository{db: db}
}

// Create 参加統計を作成する
func (r *attendanceStatisticsRepository) Create(ctx context.Context, statistics *entity.AttendanceStatistics) error {
	gormStatistics := r.toGORMAttendanceStatistics(statistics)
	return r.db.WithContext(ctx).Create(gormStatistics).Error
}

// GetByUserID ユーザーIDで統計を取得する
func (r *attendanceStatisticsRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.AttendanceStatistics, error) {
	var gormStatistics attendanceModel.AttendanceStatistics
	err := r.db.WithContext(ctx).First(&gormStatistics, "user_id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMAttendanceStatistics(&gormStatistics)
}

// Update 参加統計を更新する
func (r *attendanceStatisticsRepository) Update(ctx context.Context, statistics *entity.AttendanceStatistics) error {
	gormStatistics := r.toGORMAttendanceStatistics(statistics)
	return r.db.WithContext(ctx).Save(gormStatistics).Error
}

// Delete 参加統計を削除する
func (r *attendanceStatisticsRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&attendanceModel.AttendanceStatistics{}, "user_id = ?", userID).Error
}

// Upsert 統計を作成または更新する
func (r *attendanceStatisticsRepository) Upsert(ctx context.Context, statistics *entity.AttendanceStatistics) error {
	gormStatistics := r.toGORMAttendanceStatistics(statistics)
	return r.db.WithContext(ctx).Save(gormStatistics).Error
}

// GetRankingByTotalDays 総参加日数ランキングを取得する
func (r *attendanceStatisticsRepository) GetRankingByTotalDays(ctx context.Context, limit, offset int) ([]*entity.AttendanceStatistics, error) {
	var gormStatistics []attendanceModel.AttendanceStatistics
	err := r.db.WithContext(ctx).Order("total_days DESC").Limit(limit).Offset(offset).Find(&gormStatistics).Error
	if err != nil {
		return nil, err
	}

	statistics := make([]*entity.AttendanceStatistics, len(gormStatistics))
	for i, gormStat := range gormStatistics {
		stat, err := r.fromGORMAttendanceStatistics(&gormStat)
		if err != nil {
			return nil, err
		}
		statistics[i] = stat
	}
	return statistics, nil
}

// GetRankingByCurrentStreak 現在の連続参加日数ランキングを取得する
func (r *attendanceStatisticsRepository) GetRankingByCurrentStreak(ctx context.Context, limit, offset int) ([]*entity.AttendanceStatistics, error) {
	var gormStatistics []attendanceModel.AttendanceStatistics
	err := r.db.WithContext(ctx).Order("current_streak DESC").Limit(limit).Offset(offset).Find(&gormStatistics).Error
	if err != nil {
		return nil, err
	}

	statistics := make([]*entity.AttendanceStatistics, len(gormStatistics))
	for i, gormStat := range gormStatistics {
		stat, err := r.fromGORMAttendanceStatistics(&gormStat)
		if err != nil {
			return nil, err
		}
		statistics[i] = stat
	}
	return statistics, nil
}

// GetRankingByMaxStreak 最大連続参加日数ランキングを取得する
func (r *attendanceStatisticsRepository) GetRankingByMaxStreak(ctx context.Context, limit, offset int) ([]*entity.AttendanceStatistics, error) {
	var gormStatistics []attendanceModel.AttendanceStatistics
	err := r.db.WithContext(ctx).Order("max_streak DESC").Limit(limit).Offset(offset).Find(&gormStatistics).Error
	if err != nil {
		return nil, err
	}

	statistics := make([]*entity.AttendanceStatistics, len(gormStatistics))
	for i, gormStat := range gormStatistics {
		stat, err := r.fromGORMAttendanceStatistics(&gormStat)
		if err != nil {
			return nil, err
		}
		statistics[i] = stat
	}
	return statistics, nil
}

// 型変換: Domain Entity → GORM Model
func (r *attendanceStatisticsRepository) toGORMAttendanceStatistics(statistics *entity.AttendanceStatistics) *attendanceModel.AttendanceStatistics {
	return &attendanceModel.AttendanceStatistics{
		UserID:        statistics.UserID(),
		TotalDays:     statistics.TotalDays(),
		CurrentStreak: statistics.CurrentStreak(),
		MaxStreak:     statistics.MaxStreak(),
		LastAttendance: statistics.LastAttendance(),
		CreatedAt:     statistics.CreatedAt(),
		UpdatedAt:     statistics.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *attendanceStatisticsRepository) fromGORMAttendanceStatistics(gormStatistics *attendanceModel.AttendanceStatistics) (*entity.AttendanceStatistics, error) {
	// TODO: entity.NewAttendanceStatistics の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}