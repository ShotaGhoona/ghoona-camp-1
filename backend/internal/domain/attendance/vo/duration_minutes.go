package vo

import "errors"

// DurationMinutes 参加時間（分）を表すValue Object
type DurationMinutes int

const (
	MinDurationMinutes = 0    // 最小参加時間
	MaxDurationMinutes = 1440 // 最大参加時間（24時間）
)

// NewDurationMinutes DurationMinutesを作成する
func NewDurationMinutes(value int) (DurationMinutes, error) {
	if value < MinDurationMinutes || value > MaxDurationMinutes {
		return 0, errors.New("参加時間は0-1440分の範囲で指定してください")
	}
	return DurationMinutes(value), nil
}

// Value 内部の整数値を返す
func (d DurationMinutes) Value() int {
	return int(d)
}

// IsValid 有効な参加時間かどうかを検証する
func (d DurationMinutes) IsValid() bool {
	return int(d) >= MinDurationMinutes && int(d) <= MaxDurationMinutes
}