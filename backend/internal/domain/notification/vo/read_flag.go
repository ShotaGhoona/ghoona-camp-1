package vo

// ReadFlag 通知の既読/未読状態を表すValue Object
type ReadFlag bool

const (
	ReadFlagRead   ReadFlag = true  // 既読
	ReadFlagUnread ReadFlag = false // 未読
)

// NewReadFlag ReadFlagを作成する
func NewReadFlag(value bool) ReadFlag {
	return ReadFlag(value)
}

// IsRead 既読かどうかを判定する
func (f ReadFlag) IsRead() bool {
	return bool(f)
}

// String 文字列表現を返す
func (f ReadFlag) String() string {
	if f.IsRead() {
		return "read"
	}
	return "unread"
}

// Value 内部の真偽値を返す
func (f ReadFlag) Value() bool {
	return bool(f)
}