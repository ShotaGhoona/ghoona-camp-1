package user

import (
	"context"

	"github.com/google/uuid"

	"ghoona-camp-backend/internal/application/dto/user"
	"ghoona-camp-backend/internal/application/transaction"
	domainUser "ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/user/entity"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/domain/user/service"
	"ghoona-camp-backend/internal/domain/user/value"
)

// UserSocialUseCase ユーザーソーシャルリンク操作のユースケース
type UserSocialUseCase interface {
	GetUserSocialLinks(ctx context.Context, userID uuid.UUID) (*user.SocialLinkListResponse, error)
	CreateSocialLink(ctx context.Context, userID uuid.UUID, req *user.CreateSocialLinkRequest) (*user.SocialLinkResponse, error)
	UpdateSocialLink(ctx context.Context, linkID uuid.UUID, req *user.UpdateSocialLinkRequest) (*user.SocialLinkResponse, error)
	DeleteSocialLink(ctx context.Context, linkID uuid.UUID) error
}

type userSocialUseCase struct {
	userRepo          repository.UserRepository
	socialLinkRepo    repository.UserSocialLinkRepository
	userService       *service.UserService
	validationService *service.UserValidationService
	txManager         transaction.Manager
}

// NewUserSocialUseCase 新しいUserSocialUseCaseを作成
func NewUserSocialUseCase(
	userRepo repository.UserRepository,
	socialLinkRepo repository.UserSocialLinkRepository,
	userService *service.UserService,
	validationService *service.UserValidationService,
	txManager transaction.Manager,
) UserSocialUseCase {
	return &userSocialUseCase{
		userRepo:          userRepo,
		socialLinkRepo:    socialLinkRepo,
		userService:       userService,
		validationService: validationService,
		txManager:         txManager,
	}
}

// GetUserSocialLinks ユーザーのソーシャルリンクを取得
func (u *userSocialUseCase) GetUserSocialLinks(ctx context.Context, userID uuid.UUID) (*user.SocialLinkListResponse, error) {
	// ユーザーの存在確認
	userEntity, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, domainUser.ErrUserNotFound
	}

	links, err := u.socialLinkRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &user.SocialLinkListResponse{
		SocialLinks: user.SocialLinkListFromEntities(links),
		Total:       len(links),
	}, nil
}

// CreateSocialLink ソーシャルリンクを作成
func (u *userSocialUseCase) CreateSocialLink(ctx context.Context, userID uuid.UUID, req *user.CreateSocialLinkRequest) (*user.SocialLinkResponse, error) {
	var response *user.SocialLinkResponse
	
	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ユーザーの存在確認
		userEntity, err := u.userRepo.GetByID(txCtx, userID)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return domainUser.ErrUserNotFound
		}

		// 同じプラットフォームの重複チェック
		if err := u.validationService.ValidateUserSocialLinkUniqueness(txCtx, userID, req.Platform); err != nil {
			return err
		}

		// プラットフォームの検証
		platform := value.Platform(req.Platform)
		if !platform.IsValid() {
			return domainUser.ErrInvalidPlatform
		}

		// URLを検証
		if err := u.userService.ValidateSocialLinkURL(platform, req.URL); err != nil {
			return err
		}

		// デフォルト値の設定
		isPublic := value.PublicTrue
		if req.IsPublic != nil && !*req.IsPublic {
			isPublic = value.PublicFalse
		}

		// ソーシャルリンクエンティティの作成
		link := entity.NewUserSocialLink(userID, platform, req.URL)
		
		// オプショナルフィールドの設定
		if req.Title != nil && *req.Title != "" {
			link.Title = req.Title
		}
		link.IsPublic = isPublic

		// ユーザーのバリデーション
		if err := u.validationService.ValidateUserSocialLink(link); err != nil {
			return err
		}

		// データベースに保存
		if err := u.socialLinkRepo.Create(txCtx, link); err != nil {
			return err
		}

		// DTOに変換
		response = user.SocialLinkResponseFromEntity(link)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateSocialLink ソーシャルリンクを更新
func (u *userSocialUseCase) UpdateSocialLink(ctx context.Context, linkID uuid.UUID, req *user.UpdateSocialLinkRequest) (*user.SocialLinkResponse, error) {
	var response *user.SocialLinkResponse
	
	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ソーシャルリンクを取得
		link, err := u.socialLinkRepo.GetByID(txCtx, linkID)
		if err != nil {
			return err
		}
		if link == nil {
			return domainUser.ErrUserSocialLinkNotFound
		}

		// フィールドの更新
		if req.URL != nil {
			// URLを検証
			if err := u.userService.ValidateSocialLinkURL(link.Platform, *req.URL); err != nil {
				return err
			}
			link.URL = *req.URL
		}
		if req.Title != nil {
			link.Title = req.Title
		}
		if req.IsPublic != nil {
			if *req.IsPublic {
				link.IsPublic = value.PublicTrue
			} else {
				link.IsPublic = value.PublicFalse
			}
		}

		// ユーザーのバリデーション
		if err := u.validationService.ValidateUserSocialLink(link); err != nil {
			return err
		}

		// データベースに保存
		if err := u.socialLinkRepo.Update(txCtx, link); err != nil {
			return err
		}

		// DTOに変換
		response = user.SocialLinkResponseFromEntity(link)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteSocialLink ソーシャルリンクを削除
func (u *userSocialUseCase) DeleteSocialLink(ctx context.Context, linkID uuid.UUID) error {
	return u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ソーシャルリンクの存在確認
		link, err := u.socialLinkRepo.GetByID(txCtx, linkID)
		if err != nil {
			return err
		}
		if link == nil {
			return domainUser.ErrUserSocialLinkNotFound
		}

		// ソーシャルリンクを削除
		return u.socialLinkRepo.Delete(txCtx, linkID)
	})
}