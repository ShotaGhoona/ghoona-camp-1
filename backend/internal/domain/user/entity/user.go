package entity

import (
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/value"
)

// User represents a user account in the system
type User struct {
	ID        uuid.UUID
	ClerkID   string
	Email     string
	Username  *string
	AvatarURL *string
	DiscordID *string
	Status    value.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new User entity
func NewUser(clerkID, email string) *User {
	return &User{
		ID:        uuid.New(),
		ClerkID:   clerkID,
		Email:     email,
		Status:    value.UserStatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// IsActive checks if the user is in active state
func (u *User) IsActive() bool {
	return u.Status.IsActiveState()
}