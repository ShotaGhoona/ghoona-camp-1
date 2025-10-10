package goal

import "errors"

// Goal domain specific errors

// Goal validation errors
var (
	ErrGoalUserIDRequired     = errors.New("ユーザーIDは必須です")
	ErrGoalTitleRequired      = errors.New("目標タイトルは必須です")
	ErrGoalEndDateBeforeStart = errors.New("終了日は開始日より後に設定してください")
)

// Goal progress validation errors
var (
	ErrGoalProgressGoalIDRequired = errors.New("目標IDは必須です")
	ErrGoalProgressUserIDRequired = errors.New("ユーザーIDは必須です")
)