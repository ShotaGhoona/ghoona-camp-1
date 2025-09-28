package service

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/user/repository"
)

// RivalService handles rival-related business logic
type RivalService struct {
	userRivalRepo repository.UserRivalRepository
}

// NewRivalService creates a new RivalService
func NewRivalService(userRivalRepo repository.UserRivalRepository) *RivalService {
	return &RivalService{
		userRivalRepo: userRivalRepo,
	}
}

// CanSetRival checks if a user can set another user as rival
func (s *RivalService) CanSetRival(ctx context.Context, userID, rivalUserID uuid.UUID) error {
	if userID == rivalUserID {
		return user.ErrCannotRivalSelf
	}
	
	count, err := s.userRivalRepo.CountByUserID(ctx, userID)
	if err != nil {
		return err
	}
	
	if count >= 3 {
		return user.ErrRivalLimitExceeded
	}
	
	return nil
}