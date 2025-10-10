package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/attendance/vo"
)

// AttendanceStatistics 参加統計エンティティ
type AttendanceStatistics struct {
	id                   uuid.UUID             // 統計ID
	userID               uuid.UUID             // ユーザーID
	totalAttendanceDays  int                   // 総参加日数
	currentStreakDays    vo.StreakDays         // 現在の連続参加日数
	maxStreakDays        vo.StreakDays         // 最大連続参加日数
	lastAttendanceDate   *time.Time            // 最後の参加日
	firstAttendanceDate  *time.Time            // 初回参加日
	totalDuration        vo.DurationMinutes    // 総参加時間(分)
	createdAt            time.Time             // 作成日時
	updatedAt            time.Time             // 更新日時
}

// NewAttendanceStatistics AttendanceStatisticsエンティティを作成する
func NewAttendanceStatistics(
	userID uuid.UUID,
	totalAttendanceDays int,
	currentStreakDays vo.StreakDays,
	maxStreakDays vo.StreakDays,
	lastAttendanceDate *time.Time,
	firstAttendanceDate *time.Time,
	totalDuration vo.DurationMinutes,
) (*AttendanceStatistics, error) {
	// バリデーション
	if totalAttendanceDays < 0 {
		return nil, errors.New("総参加日数は0以上で指定してください")
	}
	
	if !currentStreakDays.IsValid() {
		return nil, errors.New("現在の連続参加日数が不正です")
	}
	
	if !maxStreakDays.IsValid() {
		return nil, errors.New("最大連続参加日数が不正です")
	}
	
	if !totalDuration.IsValid() {
		return nil, errors.New("総参加時間が不正です")
	}
	
	// 連続参加日数の整合性チェック
	if currentStreakDays.Value() > maxStreakDays.Value() {
		return nil, errors.New("現在の連続参加日数は最大連続参加日数を超えることはできません")
	}
	
	// 日付の整合性チェック
	if firstAttendanceDate != nil && lastAttendanceDate != nil && lastAttendanceDate.Before(*firstAttendanceDate) {
		return nil, errors.New("最後の参加日は初回参加日より後に設定してください")
	}

	now := time.Now()
	return &AttendanceStatistics{
		id:                  uuid.New(),
		userID:              userID,
		totalAttendanceDays: totalAttendanceDays,
		currentStreakDays:   currentStreakDays,
		maxStreakDays:       maxStreakDays,
		lastAttendanceDate:  lastAttendanceDate,
		firstAttendanceDate: firstAttendanceDate,
		totalDuration:       totalDuration,
		createdAt:           now,
		updatedAt:           now,
	}, nil
}

// Getters
func (s *AttendanceStatistics) ID() uuid.UUID                     { return s.id }
func (s *AttendanceStatistics) UserID() uuid.UUID                 { return s.userID }
func (s *AttendanceStatistics) TotalAttendanceDays() int           { return s.totalAttendanceDays }
func (s *AttendanceStatistics) CurrentStreakDays() vo.StreakDays   { return s.currentStreakDays }
func (s *AttendanceStatistics) MaxStreakDays() vo.StreakDays       { return s.maxStreakDays }
func (s *AttendanceStatistics) LastAttendanceDate() *time.Time     { return s.lastAttendanceDate }
func (s *AttendanceStatistics) FirstAttendanceDate() *time.Time    { return s.firstAttendanceDate }
func (s *AttendanceStatistics) TotalDuration() vo.DurationMinutes  { return s.totalDuration }
func (s *AttendanceStatistics) CreatedAt() time.Time              { return s.createdAt }
func (s *AttendanceStatistics) UpdatedAt() time.Time              { return s.updatedAt }