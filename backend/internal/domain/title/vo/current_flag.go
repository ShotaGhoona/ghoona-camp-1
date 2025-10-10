package vo

// CurrentFlag 現在設定中の称号かどうかを表すValue Object
type CurrentFlag bool

const (
	CurrentFlagCurrent    CurrentFlag = true  // 現在設定中
	CurrentFlagNotCurrent CurrentFlag = false // 設定していない
)

// NewCurrentFlag CurrentFlagを作成する
func NewCurrentFlag(value bool) CurrentFlag {
	return CurrentFlag(value)
}

// IsCurrent 現在設定中かどうかを判定する
func (f CurrentFlag) IsCurrent() bool {
	return bool(f)
}

// String 文字列表現を返す
func (f CurrentFlag) String() string {
	if f.IsCurrent() {
		return "current"
	}
	return "not_current"
}

// Value 内部の真偽値を返す
func (f CurrentFlag) Value() bool {
	return bool(f)
}