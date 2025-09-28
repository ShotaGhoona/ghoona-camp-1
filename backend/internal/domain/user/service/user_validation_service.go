package service

import (
	"context"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/user/entity"
	"ghoona-camp-backend/internal/domain/user/repository"
)

// UserValidationService handles all validation logic for user domain
type UserValidationService struct {
	socialLinkRepo repository.UserSocialLinkRepository
}

// NewUserValidationService creates a new UserValidationService
func NewUserValidationService(socialLinkRepo repository.UserSocialLinkRepository) *UserValidationService {
	return &UserValidationService{
		socialLinkRepo: socialLinkRepo,
	}
}

// ValidateUser validates the user entity
func (s *UserValidationService) ValidateUser(u *entity.User) error {
	if u.ClerkID == "" {
		return user.ErrInvalidClerkID
	}
	if u.Email == "" {
		return user.ErrInvalidEmail
	}
	if !u.Status.IsValid() {
		return user.ErrInvalidUsername // TODO: 適切なエラーに修正予定
	}
	return nil
}

// ValidateUserMetadata validates the user metadata entity
func (s *UserValidationService) ValidateUserMetadata(um *entity.UserMetadata) error {
	if um.DisplayName != nil && len(*um.DisplayName) > 100 {
		return user.ErrInvalidDisplayName
	}
	if um.Tagline != nil && len(*um.Tagline) > 150 {
		return user.ErrInvalidTagline
	}
	if um.Bio != nil && len(*um.Bio) > 1000 {
		return user.ErrInvalidBio
	}
	if um.Vision != nil && len(*um.Vision) > 2000 {
		return user.ErrInvalidVision
	}
	if len(um.Skills) > 20 {
		return user.ErrTooManySkills
	}
	if len(um.Interests) > 20 {
		return user.ErrTooManyInterests
	}
	return nil
}

// ValidateUserSocialLink validates the user social link entity
func (s *UserValidationService) ValidateUserSocialLink(usl *entity.UserSocialLink) error {
	if !usl.Platform.IsValid() {
		return user.ErrInvalidPlatform
	}
	if usl.URL == "" {
		return user.ErrInvalidURL
	}
	if usl.Title != nil && len(*usl.Title) > 100 {
		return user.ErrInvalidTitle
	}
	return nil
}

// ValidateUserRival validates the user rival entity
func (s *UserValidationService) ValidateUserRival(ur *entity.UserRival) error {
	if ur.UserID == ur.RivalUserID {
		return user.ErrCannotRivalSelf
	}
	return nil
}

// ValidateUserSocialLinkUniqueness checks if user can add a social link for the platform
func (s *UserValidationService) ValidateUserSocialLinkUniqueness(ctx context.Context, userID uuid.UUID, platform string) error {
	existingLinks, err := s.socialLinkRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	
	for _, link := range existingLinks {
		if link.Platform.String() == platform {
			return user.ErrDuplicatePlatform
		}
	}
	
	return nil
}