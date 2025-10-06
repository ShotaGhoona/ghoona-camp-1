package service

import (
	"context"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/user/repository"
)

// RivalService はライバル関連のビジネスロジックを処理する
type RivalService struct {
	userRivalRepo repository.UserRivalRepository
}

// NewRivalService は新しいRivalServiceを作成する
func NewRivalService(userRivalRepo repository.UserRivalRepository) *RivalService {
	return &RivalService{
		userRivalRepo: userRivalRepo,
	}
}

// CanSetRival はユーザーが他のユーザーをライバルに設定できるかチェックする
func (s *RivalService) CanSetRival(ctx context.Context, userID, rivalUserID common.UUID) error {
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