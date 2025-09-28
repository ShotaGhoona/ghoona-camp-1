package value

// UserStatus represents the status of a user account
type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusInactive  UserStatus = "inactive"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusDeleted   UserStatus = "deleted"
)

// IsValid checks if the user status is valid
func (s UserStatus) IsValid() bool {
	switch s {
	case UserStatusActive, UserStatusInactive, UserStatusSuspended, UserStatusDeleted:
		return true
	default:
		return false
	}
}

// String returns the string representation of the user status
func (s UserStatus) String() string {
	return string(s)
}

// IsActiveState checks if the user can perform actions (login, use features)
func (s UserStatus) IsActiveState() bool {
	return s == UserStatusActive
}