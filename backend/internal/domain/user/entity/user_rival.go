package entity

import (
	"time"

	"github.com/google/uuid"
)

// UserRival represents a rival relationship between users
type UserRival struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	RivalUserID uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewUserRival creates a new UserRival entity
func NewUserRival(userID, rivalUserID uuid.UUID) *UserRival {
	return &UserRival{
		ID:          uuid.New(),
		UserID:      userID,
		RivalUserID: rivalUserID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}