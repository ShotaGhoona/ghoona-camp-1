package vo

import "errors"

// RecurrencePattern 定期開催イベントの繰り返しパターンを表すValue Object
type RecurrencePattern string

const (
	RecurrencePatternDaily  RecurrencePattern = "daily"  // 毎日
	RecurrencePatternWeekly RecurrencePattern = "weekly" // 毎週
)

// NewRecurrencePattern RecurrencePatternを作成する
func NewRecurrencePattern(value string) (RecurrencePattern, error) {
	pattern := RecurrencePattern(value)
	if !pattern.isValid() {
		return "", errors.New("不正な繰り返しパターンです")
	}
	return pattern, nil
}

// isValid 有効な繰り返しパターン値かどうかを検証する
func (p RecurrencePattern) isValid() bool {
	switch p {
	case RecurrencePatternDaily, RecurrencePatternWeekly:
		return true
	default:
		return false
	}
}

// String 文字列表現を返す
func (p RecurrencePattern) String() string {
	return string(p)
}

// Value 内部の文字列値を返す
func (p RecurrencePattern) Value() string {
	return string(p)
}