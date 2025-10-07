// Package value イベントドメインの値オブジェクトを定義する
package value

// RecurrencePattern イベントの繰り返しパターンを表す値オブジェクト
// 朝活の習慣化を支援するため、毎日・毎週・毎月の定期開催を可能にする
// noneは単発イベント、その他は自動的な次回イベント生成に使用される
type RecurrencePattern string

const (
	RecurrencePatternNone    RecurrencePattern = "none"
	RecurrencePatternDaily   RecurrencePattern = "daily"
	RecurrencePatternWeekly  RecurrencePattern = "weekly"
	RecurrencePatternMonthly RecurrencePattern = "monthly"
)

// IsValid 繰り返しパターンが有効かどうかを判定する
func (r RecurrencePattern) IsValid() bool {
	switch r {
	case RecurrencePatternNone, RecurrencePatternDaily, RecurrencePatternWeekly, RecurrencePatternMonthly:
		return true
	default:
		return false
	}
}

// String 繰り返しパターンを文字列として返す
func (r RecurrencePattern) String() string {
	return string(r)
}