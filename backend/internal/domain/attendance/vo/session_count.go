package vo

import "errors"

// SessionCount セッション数を表すValue Object
type SessionCount int

const (
	MinSessionCount = 0   // 最小セッション数
	MaxSessionCount = 100 // 最大セッション数（現実的な上限）
)

// NewSessionCount SessionCountを作成する
func NewSessionCount(value int) (SessionCount, error) {
	if value < MinSessionCount || value > MaxSessionCount {
		return 0, errors.New("セッション数は0-100の範囲で指定してください")
	}
	return SessionCount(value), nil
}

// Value 内部の整数値を返す
func (s SessionCount) Value() int {
	return int(s)
}

// IsValid 有効なセッション数かどうかを検証する
func (s SessionCount) IsValid() bool {
	return int(s) >= MinSessionCount && int(s) <= MaxSessionCount
}