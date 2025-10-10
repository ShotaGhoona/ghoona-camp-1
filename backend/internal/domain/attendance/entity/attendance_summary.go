package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/attendance/vo"
)

// AttendanceSummary 参加サマリーエンティティ
type AttendanceSummary struct {
	id                    uuid.UUID             // サマリーID
	userID                uuid.UUID             // ユーザーID
	date                  time.Time             // 参加日
	totalDuration         vo.DurationMinutes    // 1日の総参加時間(分)
	sessionCount          vo.SessionCount       // 参加セッション数
	firstJoinTime         *time.Time            // 最初の参加時刻
	lastLeaveTime         *time.Time            // 最後の退出時刻
	isMorningActive       bool                  // 朝活時間帯(6-7時)の参加有無
	createdAt             time.Time             // 作成日時
	updatedAt             time.Time             // 更新日時
}

// NewAttendanceSummary AttendanceSummaryエンティティを作成する
func NewAttendanceSummary(
	userID uuid.UUID,
	date time.Time,
	totalDuration vo.DurationMinutes,
	sessionCount vo.SessionCount,
	firstJoinTime *time.Time,
	lastLeaveTime *time.Time,
	isMorningActive bool,
) (*AttendanceSummary, error) {
	// バリデーション
	if !totalDuration.IsValid() {
		return nil, errors.New("総参加時間が不正です")
	}
	
	if !sessionCount.IsValid() {
		return nil, errors.New("セッション数が不正です")
	}
	
	// 最初の参加時刻と最後の退出時刻の整合性チェック
	if firstJoinTime != nil && lastLeaveTime != nil && lastLeaveTime.Before(*firstJoinTime) {
		return nil, errors.New("最後の退出時刻は最初の参加時刻より後に設定してください")
	}

	now := time.Now()
	return &AttendanceSummary{
		id:              uuid.New(),
		userID:          userID,
		date:            date,
		totalDuration:   totalDuration,
		sessionCount:    sessionCount,
		firstJoinTime:   firstJoinTime,
		lastLeaveTime:   lastLeaveTime,
		isMorningActive: isMorningActive,
		createdAt:       now,
		updatedAt:       now,
	}, nil
}

// Getters
func (s *AttendanceSummary) ID() uuid.UUID                 { return s.id }
func (s *AttendanceSummary) UserID() uuid.UUID             { return s.userID }
func (s *AttendanceSummary) Date() time.Time               { return s.date }
func (s *AttendanceSummary) TotalDuration() vo.DurationMinutes { return s.totalDuration }
func (s *AttendanceSummary) SessionCount() vo.SessionCount { return s.sessionCount }
func (s *AttendanceSummary) FirstJoinTime() *time.Time     { return s.firstJoinTime }
func (s *AttendanceSummary) LastLeaveTime() *time.Time     { return s.lastLeaveTime }
func (s *AttendanceSummary) IsMorningActive() bool         { return s.isMorningActive }
func (s *AttendanceSummary) CreatedAt() time.Time          { return s.createdAt }
func (s *AttendanceSummary) UpdatedAt() time.Time          { return s.updatedAt }