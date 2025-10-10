package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/event/vo"
	eventvo "ghoona-camp-backend/internal/domain/event/vo"
)

// Event イベントエンティティ
type Event struct {
	id                uuid.UUID                  // イベントID
	creatorID         uuid.UUID                  // 作成者ID
	title             string                     // イベント名
	description       string                     // イベント詳細
	eventType         vo.EventType               // イベントタイプ
	scheduledDate     time.Time                  // 開催日
	startTime         time.Time                  // 開始時間
	endTime           time.Time                  // 終了時間
	maxParticipants   *int                       // 最大参加者数
	isRecurring       bool                       // 定期開催かどうか
	recurrencePattern *vo.RecurrencePattern      // 繰り返しパターン
	discordChannelID  string                     // 対応するDiscordチャンネルID
	isActive          eventvo.ActiveFlag          // イベント有効状態
	createdAt         time.Time                  // 作成日時
	updatedAt         time.Time                  // 更新日時
}

// NewEvent Eventエンティティを作成する
func NewEvent(
	creatorID uuid.UUID,
	title string,
	description string,
	eventType vo.EventType,
	scheduledDate time.Time,
	startTime time.Time,
	endTime time.Time,
	maxParticipants *int,
	isRecurring bool,
	recurrencePattern *vo.RecurrencePattern,
	discordChannelID string,
	isActive eventvo.ActiveFlag,
) (*Event, error) {
	// バリデーション
	if title == "" {
		return nil, errors.New("イベント名は必須です")
	}
	
	// 時間制約チェック (start_time < end_time)
	if !endTime.After(startTime) {
		return nil, errors.New("終了時間は開始時間より後に設定してください")
	}
	
	// 最大参加者数チェック
	if maxParticipants != nil && *maxParticipants <= 0 {
		return nil, errors.New("最大参加者数は1以上で指定してください")
	}

	now := time.Now()
	return &Event{
		id:                uuid.New(),
		creatorID:         creatorID,
		title:             title,
		description:       description,
		eventType:         eventType,
		scheduledDate:     scheduledDate,
		startTime:         startTime,
		endTime:           endTime,
		maxParticipants:   maxParticipants,
		isRecurring:       isRecurring,
		recurrencePattern: recurrencePattern,
		discordChannelID:  discordChannelID,
		isActive:          isActive,
		createdAt:         now,
		updatedAt:         now,
	}, nil
}

// Getters
func (e *Event) ID() uuid.UUID                         { return e.id }
func (e *Event) CreatorID() uuid.UUID                  { return e.creatorID }
func (e *Event) Title() string                         { return e.title }
func (e *Event) Description() string                   { return e.description }
func (e *Event) EventType() vo.EventType               { return e.eventType }
func (e *Event) ScheduledDate() time.Time              { return e.scheduledDate }
func (e *Event) StartTime() time.Time                  { return e.startTime }
func (e *Event) EndTime() time.Time                    { return e.endTime }
func (e *Event) MaxParticipants() *int                 { return e.maxParticipants }
func (e *Event) IsRecurring() bool                     { return e.isRecurring }
func (e *Event) RecurrencePattern() *vo.RecurrencePattern { return e.recurrencePattern }
func (e *Event) DiscordChannelID() string              { return e.discordChannelID }
func (e *Event) IsActive() eventvo.ActiveFlag           { return e.isActive }
func (e *Event) CreatedAt() time.Time                  { return e.createdAt }
func (e *Event) UpdatedAt() time.Time                  { return e.updatedAt }