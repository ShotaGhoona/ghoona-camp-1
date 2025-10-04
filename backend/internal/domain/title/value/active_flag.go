package value

// ActiveFlag は称号のアクティブ状態を表す
type ActiveFlag bool

const (
	ActiveTrue  ActiveFlag = true  // アクティブ
	ActiveFalse ActiveFlag = false // 非アクティブ
)

// String はアクティブフラグの文字列表現を返す
func (a ActiveFlag) String() string {
	if bool(a) {
		return "active"
	}
	return "inactive"
}

// Bool はboolean表現を返す
func (a ActiveFlag) Bool() bool {
	return bool(a)
}

// IsValid はアクティブフラグが有効かどうかを確認する
func (a ActiveFlag) IsValid() bool {
	return a == ActiveTrue || a == ActiveFalse
}