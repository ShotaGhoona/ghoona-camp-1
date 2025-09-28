package value

// UserStatus はユーザーアカウントの状態を表す
type UserStatus string

const (
	UserStatusActive    UserStatus = "active"    // アクティブ
	UserStatusInactive  UserStatus = "inactive"  // 非アクティブ
	UserStatusSuspended UserStatus = "suspended" // 停止中
	UserStatusDeleted   UserStatus = "deleted"   // 削除済み
)

// IsValid はユーザーステータスが有効かどうかを確認する
func (s UserStatus) IsValid() bool {
	switch s {
	case UserStatusActive, UserStatusInactive, UserStatusSuspended, UserStatusDeleted:
		return true
	default:
		return false
	}
}

// String はユーザーステータスの文字列表現を返す
func (s UserStatus) String() string {
	return string(s)
}

// IsActiveState はユーザーがアクション（ログイン、機能利用）を実行できるかどうかを確認する
func (s UserStatus) IsActiveState() bool {
	return s == UserStatusActive
}