package value

// PublicFlag はシンプルな公開/非公開の可視性フラグを表す
// vision_publicやis_publicなどのフィールドがbooleanであるAPI要件に対応
type PublicFlag bool

const (
	PublicTrue  PublicFlag = true  // 公開
	PublicFalse PublicFlag = false // 非公開
)

// String は公開フラグの文字列表現を返す
func (p PublicFlag) String() string {
	if bool(p) {
		return "public"
	}
	return "private"
}

// Bool はboolean表現を返す
func (p PublicFlag) Bool() bool {
	return bool(p)
}