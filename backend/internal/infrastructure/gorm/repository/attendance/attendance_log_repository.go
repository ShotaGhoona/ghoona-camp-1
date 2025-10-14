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

type attendanceLogRepository struct {
	db *gorm.DB
}

// NewAttendanceLogRepository コンストラクタ
func NewAttendanceLogRepository(db *gorm.DB) repository.AttendanceLogRepository {
	return &attendanceLogRepository{db: db}
}

// Create 参加ログを作成する
func (r *attendanceLogRepository) Create(ctx context.Context, log *entity.AttendanceLog) error {
	gormLog := r.toGORMAttendanceLog(log)
	return r.db.WithContext(ctx).Create(gormLog).Error
}

// GetByID IDで参加ログを取得する
func (r *attendanceLogRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.AttendanceLog, error) {
	var gormLog attendanceModel.AttendanceLog
	err := r.db.WithContext(ctx).First(&gormLog, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMAttendanceLog(&gormLog)
}

// GetByUserID ユーザーIDで参加ログ一覧を取得する
func (r *attendanceLogRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.AttendanceLog, error) {
	var gormLogs []attendanceModel.AttendanceLog
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&gormLogs).Error
	if err != nil {
		return nil, err
	}

	logs := make([]*entity.AttendanceLog, len(gormLogs))
	for i, gormLog := range gormLogs {
		log, err := r.fromGORMAttendanceLog(&gormLog)
		if err != nil {
			return nil, err
		}
		logs[i] = log
	}
	return logs, nil
}

// GetByUserAndDate ユーザーIDと日付で参加ログを取得する
func (r *attendanceLogRepository) GetByUserAndDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.AttendanceLog, error) {
	var gormLog attendanceModel.AttendanceLog
	err := r.db.WithContext(ctx).Where("user_id = ? AND DATE(attendance_date) = DATE(?)", userID, date).First(&gormLog).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMAttendanceLog(&gormLog)
}

// GetByDateRange 日付範囲で参加ログを取得する
func (r *attendanceLogRepository) GetByDateRange(ctx context.Context, userID uuid.UUID, dateFrom, dateTo time.Time) ([]*entity.AttendanceLog, error) {
	var gormLogs []attendanceModel.AttendanceLog
	err := r.db.WithContext(ctx).Where("user_id = ? AND attendance_date >= ? AND attendance_date <= ?", userID, dateFrom, dateTo).Find(&gormLogs).Error
	if err != nil {
		return nil, err
	}

	logs := make([]*entity.AttendanceLog, len(gormLogs))
	for i, gormLog := range gormLogs {
		log, err := r.fromGORMAttendanceLog(&gormLog)
		if err != nil {
			return nil, err
		}
		logs[i] = log
	}
	return logs, nil
}

// GetByType タイプで参加ログを取得する
func (r *attendanceLogRepository) GetByType(ctx context.Context, userID uuid.UUID, logType string, limit, offset int) ([]*entity.AttendanceLog, error) {
	var gormLogs []attendanceModel.AttendanceLog
	err := r.db.WithContext(ctx).Where("user_id = ? AND log_type = ?", userID, logType).Limit(limit).Offset(offset).Find(&gormLogs).Error
	if err != nil {
		return nil, err
	}

	logs := make([]*entity.AttendanceLog, len(gormLogs))
	for i, gormLog := range gormLogs {
		log, err := r.fromGORMAttendanceLog(&gormLog)
		if err != nil {
			return nil, err
		}
		logs[i] = log
	}
	return logs, nil
}

// Update 参加ログを更新する
func (r *attendanceLogRepository) Update(ctx context.Context, log *entity.AttendanceLog) error {
	gormLog := r.toGORMAttendanceLog(log)
	return r.db.WithContext(ctx).Save(gormLog).Error
}

// Delete 参加ログを削除する
func (r *attendanceLogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&attendanceModel.AttendanceLog{}, "id = ?", id).Error
}

// 型変換: Domain Entity → GORM Model
func (r *attendanceLogRepository) toGORMAttendanceLog(log *entity.AttendanceLog) *attendanceModel.AttendanceLog {
	return &attendanceModel.AttendanceLog{
		ID:             log.ID(),
		UserID:         log.UserID(),
		AttendanceDate: log.AttendanceDate(),
		LogType:        log.LogType(),
		EventID:        log.EventID(),
		GoalID:         log.GoalID(),
		Duration:       log.Duration(),
		Notes:          log.Notes(),
		CreatedAt:      log.CreatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *attendanceLogRepository) fromGORMAttendanceLog(gormLog *attendanceModel.AttendanceLog) (*entity.AttendanceLog, error) {
	// TODO: entity.NewAttendanceLog の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}