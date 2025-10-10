package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/user/entity"
	"ghoona-camp-backend/internal/domain/user/repository"
)

// RivalManagementService ライバル管理のビジネスロジックを提供するドメインサービス
type RivalManagementService struct {
	userRivalRepo repository.UserRivalRepository
}

// NewRivalManagementService コンストラクタ
func NewRivalManagementService(userRivalRepo repository.UserRivalRepository) *RivalManagementService {
	return &RivalManagementService{
		userRivalRepo: userRivalRepo,
	}
}

const MaxRivalsPerUser = 3

// CanAddRival ライバル追加が可能かチェックする
func (s *RivalManagementService) CanAddRival(ctx context.Context, userID, rivalUserID uuid.UUID) error {
	// 1. 自分自身をライバルに設定することの禁止
	if userID == rivalUserID {
		return errors.New("自分自身をライバルに設定することはできません")
	}

	// 2. 現在のライバル数を取得
	currentCount, err := s.userRivalRepo.CountByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// 3. 3人制限チェック
	if currentCount >= MaxRivalsPerUser {
		return errors.New("ライバルは最大3人までしか設定できません")
	}

	// 4. 既存ライバル一覧を取得して重複チェック
	existingRivals, err := s.userRivalRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// 5. 重複ライバル登録防止
	for _, rival := range existingRivals {
		if rival.RivalUserID() == rivalUserID {
			return errors.New("このユーザーは既にライバルに設定されています")
		}
	}

	return nil
}

// AddRival ライバルを追加する（ビジネスルールチェック付き）
func (s *RivalManagementService) AddRival(ctx context.Context, userID, rivalUserID uuid.UUID) (*entity.UserRival, error) {
	// ビジネスルールチェック
	if err := s.CanAddRival(ctx, userID, rivalUserID); err != nil {
		return nil, err
	}

	// ライバルエンティティ作成
	rival, err := entity.NewUserRival(userID, rivalUserID)
	if err != nil {
		return nil, err
	}

	// リポジトリに保存
	if err := s.userRivalRepo.Create(ctx, rival); err != nil {
		return nil, err
	}

	return rival, nil
}