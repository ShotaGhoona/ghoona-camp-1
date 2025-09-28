package value

// PublicFlag represents a simple public/private visibility flag
// This aligns with the API requirements where fields like vision_public and is_public are boolean
type PublicFlag bool

const (
	PublicTrue  PublicFlag = true  // 公開
	PublicFalse PublicFlag = false // 非公開
)

// String returns the string representation of the public flag
func (p PublicFlag) String() string {
	if bool(p) {
		return "public"
	}
	return "private"
}

// Bool returns the boolean representation
func (p PublicFlag) Bool() bool {
	return bool(p)
}