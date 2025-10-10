package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/attendance/vo"
	uservo "ghoona-camp-backend/internal/domain/user/vo"
)

// AttendanceLog 参加ログエンティティ
type AttendanceLog struct {
	id               uuid.UUID               // ログID
	userID           uuid.UUID               // ユーザーID
	eventID          *uuid.UUID              // 関連イベントID (NULL可)
	discordChannelID string                  // 参加したDiscordチャンネルID
	joinedAt         time.Time               // Discord参加開始時刻
	leftAt           *time.Time              // Discord退出時刻
	duration         vo.DurationMinutes      // 参加時間(分)
	isValid          uservo.ActiveFlag       // 有効な参加かどうか
	createdAt        time.Time               // 記録作成日時
	updatedAt        time.Time               // 更新日時
}

// NewAttendanceLog AttendanceLogエンティティを作成する
func NewAttendanceLog(
	userID uuid.UUID,
	eventID *uuid.UUID,
	discordChannelID string,
	joinedAt time.Time,
	leftAt *time.Time,
	duration vo.DurationMinutes,
	isValid uservo.ActiveFlag,
) (*AttendanceLog, error) {
	// バリデーション
	if discordChannelID == "" {
		return nil, errors.New("DiscordチャンネルIDは必須です")
	}
	
	// 参加時間の妥当性チェック
	if !duration.IsValid() {
		return nil, errors.New("参加時間が不正です")
	}
	
	// 退出時刻が参加時刻より前でないことをチェック
	if leftAt != nil && leftAt.Before(joinedAt) {
		return nil, errors.New("退出時刻は参加時刻より後に設定してください")
	}

	now := time.Now()
	return &AttendanceLog{
		id:               uuid.New(),
		userID:           userID,
		eventID:          eventID,
		discordChannelID: discordChannelID,
		joinedAt:         joinedAt,
		leftAt:           leftAt,
		duration:         duration,
		isValid:          isValid,
		createdAt:        now,
		updatedAt:        now,
	}, nil
}

// Getters
func (l *AttendanceLog) ID() uuid.UUID                    { return l.id }
func (l *AttendanceLog) UserID() uuid.UUID                { return l.userID }
func (l *AttendanceLog) EventID() *uuid.UUID              { return l.eventID }
func (l *AttendanceLog) DiscordChannelID() string         { return l.discordChannelID }
func (l *AttendanceLog) JoinedAt() time.Time              { return l.joinedAt }
func (l *AttendanceLog) LeftAt() *time.Time               { return l.leftAt }
func (l *AttendanceLog) Duration() vo.DurationMinutes     { return l.duration }
func (l *AttendanceLog) IsValid() uservo.ActiveFlag       { return l.isValid }
func (l *AttendanceLog) CreatedAt() time.Time             { return l.createdAt }
func (l *AttendanceLog) UpdatedAt() time.Time             { return l.updatedAt }