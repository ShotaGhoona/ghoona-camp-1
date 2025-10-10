package attendance

import "errors"

// Attendance domain specific errors

// Attendance validation errors
var (
	ErrAttendanceUserIDRequired  = errors.New("ユーザーIDは必須です")
	ErrAttendanceEventIDRequired = errors.New("イベントIDは必須です")
)

// Attendance statistics validation errors
var (
	ErrAttendanceStatisticsUserIDRequired = errors.New("ユーザーIDは必須です")
)