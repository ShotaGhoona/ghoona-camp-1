package vo

// PublicFlag 公開/非公開設定を表すValue Object
type PublicFlag bool

const (
	PublicFlagPublic  PublicFlag = true  // 公開
	PublicFlagPrivate PublicFlag = false // 非公開
)

// NewPublicFlag PublicFlagを作成する
func NewPublicFlag(value bool) PublicFlag {
	return PublicFlag(value)
}

// IsPublic 公開設定かどうかを判定する
func (f PublicFlag) IsPublic() bool {
	return bool(f)
}

// String 文字列表現を返す
func (f PublicFlag) String() string {
	if f.IsPublic() {
		return "public"
	}
	return "private"
}

// Value 内部の真偽値を返す
func (f PublicFlag) Value() bool {
	return bool(f)
}