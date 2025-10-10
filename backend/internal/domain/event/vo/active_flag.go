package vo

// ActiveFlag エンティティの有効/無効状態を表すValue Object
type ActiveFlag bool

const (
	ActiveFlagActive   ActiveFlag = true  // 有効
	ActiveFlagInactive ActiveFlag = false // 無効
)

// NewActiveFlag ActiveFlagを作成する
func NewActiveFlag(value bool) ActiveFlag {
	return ActiveFlag(value)
}

// IsActive 有効状態かどうかを判定する
func (f ActiveFlag) IsActive() bool {
	return bool(f)
}

// String 文字列表現を返す
func (f ActiveFlag) String() string {
	if f.IsActive() {
		return "active"
	}
	return "inactive"
}

// Value 内部の真偽値を返す
func (f ActiveFlag) Value() bool {
	return bool(f)
}