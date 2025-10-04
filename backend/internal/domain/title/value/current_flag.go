package value

// CurrentFlag は称号の現在表示状態を表す
type CurrentFlag bool

const (
	CurrentTrue  CurrentFlag = true  // 現在表示中
	CurrentFalse CurrentFlag = false // 非表示
)

// String は現在表示フラグの文字列表現を返す
func (c CurrentFlag) String() string {
	if bool(c) {
		return "current"
	}
	return "not_current"
}

// Bool はboolean表現を返す
func (c CurrentFlag) Bool() bool {
	return bool(c)
}