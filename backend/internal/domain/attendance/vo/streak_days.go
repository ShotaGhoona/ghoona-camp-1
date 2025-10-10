package vo

import "errors"

// StreakDays 連続参加日数を表すValue Object
type StreakDays int

const (
	MinStreakDays = 0     // 最小連続参加日数
	MaxStreakDays = 10000 // 最大連続参加日数（現実的な上限）
)

// NewStreakDays StreakDaysを作成する
func NewStreakDays(value int) (StreakDays, error) {
	if value < MinStreakDays || value > MaxStreakDays {
		return 0, errors.New("連続参加日数は0-10000の範囲で指定してください")
	}
	return StreakDays(value), nil
}

// Value 内部の整数値を返す
func (s StreakDays) Value() int {
	return int(s)
}

// IsValid 有効な連続参加日数かどうかを検証する
func (s StreakDays) IsValid() bool {
	return int(s) >= MinStreakDays && int(s) <= MaxStreakDays
}