// Package entity イベントドメインのエンティティを定義する
// エンティティは一意の識別子を持ち、ビジネスロジックと状態を管理する
package entity

import (
	"time"

	"ghoona-camp-backend/internal/domain/event"
	"ghoona-camp-backend/internal/domain/event/value"

	"github.com/google/uuid"
)

// Event 朝活イベントを表すエンティティ
type Event struct {
	ID                uuid.UUID
	CreatorID         uuid.UUID
	Title             string
	Description       string
	EventType         value.EventType
	ScheduledDate     time.Time
	StartTime         time.Time
	EndTime           time.Time
	MaxParticipants   int
	IsRecurring       bool
	RecurrencePattern value.RecurrencePattern
	DiscordChannelID  string
	IsActive          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// NewEvent 新しいイベントを作成するファクトリメソッド
func NewEvent(
	creatorID uuid.UUID,
	title, description string,
	eventType value.EventType,
	scheduledDate, startTime, endTime time.Time,
	maxParticipants int,
	isRecurring bool,
	recurrencePattern value.RecurrencePattern,
	discordChannelID string,
) *Event {
	now := time.Now()
	return &Event{
		ID:                uuid.New(),
		CreatorID:         creatorID,
		Title:             title,
		Description:       description,
		EventType:         eventType,
		ScheduledDate:     scheduledDate,
		StartTime:         startTime,
		EndTime:           endTime,
		MaxParticipants:   maxParticipants,
		IsRecurring:       isRecurring,
		RecurrencePattern: recurrencePattern,
		DiscordChannelID:  discordChannelID,
		IsActive:          true,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}

// TimeSlot 開始時刻と終了時刻を取得する
func (e *Event) TimeSlot() (start, end time.Time) {
	return e.StartTime, e.EndTime
}

// Status 指定時刻での現在のステータスを計算する
func (e *Event) Status(now time.Time) value.EventStatus {
	if !e.IsActive {
		return value.EventStatusCancelled
	}
	switch {
	case now.Before(e.StartTime):
		return value.EventStatusUpcoming
	case now.Equal(e.StartTime) || (now.After(e.StartTime) && now.Before(e.EndTime)):
		return value.EventStatusOngoing
	default:
		return value.EventStatusCompleted
	}
}

// IsUpcoming 指定時刻でイベントが開催前かどうかを判定する
func (e *Event) IsUpcoming(now time.Time) bool {
	return e.Status(now) == value.EventStatusUpcoming
}

// IsOngoing 指定時刻でイベントが開催中かどうかを判定する
func (e *Event) IsOngoing(now time.Time) bool {
	return e.Status(now) == value.EventStatusOngoing
}

// IsCompleted 指定時刻でイベントが完了済みかどうかを判定する
func (e *Event) IsCompleted(now time.Time) bool {
	return e.Status(now) == value.EventStatusCompleted
}

// Cancel イベントをキャンセルする（論理削除）
func (e *Event) Cancel() error {
	if !e.IsActive {
		return nil
	}
	e.IsActive = false
	e.UpdatedAt = time.Now()
	return nil
}

// Reschedule イベントの日時を変更する
func (e *Event) Reschedule(date, start, end time.Time) error {
	now := time.Now()
	if e.IsCompleted(now) {
		return event.ErrEventCompleted
	}
	if end.Before(start) || end.Equal(start) {
		return event.ErrInvalidTimeSlot
	}
	
	e.ScheduledDate = date
	e.StartTime = start
	e.EndTime = end
	e.UpdatedAt = time.Now()
	return nil
}

// CanAcceptRegistration 新規参加登録を受け入れ可能かどうかを判定する
func (e *Event) CanAcceptRegistration(now time.Time, currentRegistered int) bool {
	if !e.IsActive || e.IsCompleted(now) {
		return false
	}
	return currentRegistered < e.MaxParticipants
}

// ChangeCapacity 定員を変更する
func (e *Event) ChangeCapacity(newMax int, currentRegistered int) error {
	if newMax < currentRegistered {
		return event.ErrCannotReduceCapacity
	}
	e.MaxParticipants = newMax
	e.UpdatedAt = time.Now()
	return nil
}

// UpdateBasicInfo 基本情報を更新する
func (e *Event) UpdateBasicInfo(title, description string, eventType value.EventType) error {
	if title == "" {
		return event.ErrInvalidTitle
	}
	if !eventType.IsValid() {
		return event.ErrInvalidEventType
	}
	
	e.Title = title
	e.Description = description
	e.EventType = eventType
	e.UpdatedAt = time.Now()
	return nil
}

// AssignDiscord DiscordチャンネルIDを設定する
func (e *Event) AssignDiscord(channelID string) {
	e.DiscordChannelID = channelID
	e.UpdatedAt = time.Now()
}

// ClearDiscord DiscordチャンネルIDをクリアする
func (e *Event) ClearDiscord() {
	e.DiscordChannelID = ""
	e.UpdatedAt = time.Now()
}
