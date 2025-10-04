package user

import (
	"context"

	"ghoona-camp-backend/internal/application/dto/user"
	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/domain/common"
	domainUser "ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/user/entity"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/domain/user/service"
)

// UserRivalUseCase ユーザーライバル操作のユースケース
type UserRivalUseCase interface {
	GetUserRivals(ctx context.Context, userID common.UUID) (*user.RivalListResponse, error)
	AddRival(ctx context.Context, userID common.UUID, req *user.AddRivalRequest) (*user.RivalResponse, error)
	RemoveRival(ctx context.Context, rivalID common.UUID) error
}

type userRivalUseCase struct {
	userRepo          repository.UserRepository
	rivalRepo         repository.UserRivalRepository
	rivalService      *service.RivalService
	validationService *service.UserValidationService
	txManager         transaction.Manager
}

// NewUserRivalUseCase 新しいUserRivalUseCaseを作成
func NewUserRivalUseCase(
	userRepo repository.UserRepository,
	rivalRepo repository.UserRivalRepository,
	rivalService *service.RivalService,
	validationService *service.UserValidationService,
	txManager transaction.Manager,
) UserRivalUseCase {
	return &userRivalUseCase{
		userRepo:          userRepo,
		rivalRepo:         rivalRepo,
		rivalService:      rivalService,
		validationService: validationService,
		txManager:         txManager,
	}
}

// GetUserRivals ユーザーのライバルを取得
func (u *userRivalUseCase) GetUserRivals(ctx context.Context, userID common.UUID) (*user.RivalListResponse, error) {
	// ユーザーの存在確認
	userEntity, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, domainUser.ErrUserNotFound
	}

	rivals, err := u.rivalRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// ライバルユーザーの詳細を取得
	rivalUsers := make([]*entity.User, 0, len(rivals))
	for _, rival := range rivals {
		rivalUser, err := u.userRepo.GetByID(ctx, rival.RivalUserID)
		if err == nil && rivalUser != nil {
			rivalUsers = append(rivalUsers, rivalUser)
		}
	}

	return &user.RivalListResponse{
		Rivals:    user.RivalListFromEntities(rivals, rivalUsers),
		Total:     len(rivals),
		MaxRivals: 3,
	}, nil
}

// AddRival ライバルを追加
func (u *userRivalUseCase) AddRival(ctx context.Context, userID common.UUID, req *user.AddRivalRequest) (*user.RivalResponse, error) {
	var response *user.RivalResponse
	
	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ユーザーの存在確認
		userEntity, err := u.userRepo.GetByID(txCtx, userID)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return domainUser.ErrUserNotFound
		}

		// ライバルユーザーの存在確認
		rivalUser, err := u.userRepo.GetByID(txCtx, req.RivalUserID)
		if err != nil {
			return err
		}
		if rivalUser == nil {
			return domainUser.ErrUserNotFound
		}

		// ライバル設定可能性をチェック
		if err := u.rivalService.CanSetRival(txCtx, userID, req.RivalUserID); err != nil {
			return err
		}

		// ライバルエンティティを作成
		rival := entity.NewUserRival(userID, req.RivalUserID)

		// ユーザーのバリデーション
		if err := u.validationService.ValidateUserRival(rival); err != nil {
			return err
		}

		// データベースに保存
		if err := u.rivalRepo.Create(txCtx, rival); err != nil {
			return err
		}

		// DTOに変換
		response = user.RivalResponseFromEntity(rival, rivalUser)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

// RemoveRival ライバルを削除
func (u *userRivalUseCase) RemoveRival(ctx context.Context, rivalID common.UUID) error {
	return u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ライバルを削除
		return u.rivalRepo.Delete(txCtx, rivalID)
	})
}