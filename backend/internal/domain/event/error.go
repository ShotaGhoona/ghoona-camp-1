package event

import "errors"

// Event domain specific errors

// Event participation errors
var (
	ErrEventParticipantAlreadyRegistered = errors.New("既にこのイベントに参加登録済みです")
	ErrEventCapacityFull                 = errors.New("イベントの定員に達しています")
)

// Event validation errors
var (
	ErrEventTitleRequired     = errors.New("イベントタイトルは必須です")
	ErrEventEndTimeBeforeStart = errors.New("終了時間は開始時間より後に設定してください")
	ErrEventCreatorIDRequired = errors.New("作成者IDは必須です")
)

// Event participant validation errors
var (
	ErrEventParticipantEventIDRequired = errors.New("イベントIDは必須です")
	ErrEventParticipantUserIDRequired  = errors.New("ユーザーIDは必須です")
)