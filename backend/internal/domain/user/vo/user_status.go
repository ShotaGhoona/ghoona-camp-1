package vo

import "errors"

// UserStatus ユーザーのアカウント状態を表すValue Object
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"   // アクティブ
	UserStatusInactive UserStatus = "inactive" // 非アクティブ
)

// NewUserStatus UserStatusを作成する
func NewUserStatus(value string) (UserStatus, error) {
	status := UserStatus(value)
	if !status.isValid() {
		return "", errors.New("不正なユーザーステータスです")
	}
	return status, nil
}

// isValid 有効なステータス値かどうかを検証する
func (s UserStatus) isValid() bool {
	switch s {
	case UserStatusActive, UserStatusInactive:
		return true
	default:
		return false
	}
}

// IsActive アクティブ状態かどうかを判定する
func (s UserStatus) IsActive() bool {
	return s == UserStatusActive
}

// String 文字列表現を返す
func (s UserStatus) String() string {
	return string(s)
}

// Value 内部の文字列値を返す
func (s UserStatus) Value() string {
	return string(s)
}