package user

import (
	"context"
	"log"

	"github.com/google/uuid"

	"ghoona-camp-backend/internal/application/dto/user"
	"ghoona-camp-backend/internal/application/transaction"
	domainUser "ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/user/entity"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/domain/user/service"
	"ghoona-camp-backend/internal/domain/user/value"
)

// UserMetadataUseCase ユーザーメタデータ操作のユースケース
type UserMetadataUseCase interface {
	GetUserMetadata(ctx context.Context, userID uuid.UUID) (*user.UserMetadataResponse, error)
	CreateUserMetadata(ctx context.Context, userID uuid.UUID, req *user.CreateUserMetadataRequest) (*user.UserMetadataResponse, error)
	UpdateUserMetadata(ctx context.Context, userID uuid.UUID, req *user.UpdateUserMetadataRequest) (*user.UserMetadataResponse, error)
}

type userMetadataUseCase struct {
	userRepo          repository.UserRepository
	metadataRepo      repository.UserMetadataRepository
	validationService *service.UserValidationService
	txManager         transaction.Manager
}

// NewUserMetadataUseCase 新しいUserMetadataUseCaseを作成
func NewUserMetadataUseCase(
	userRepo repository.UserRepository,
	metadataRepo repository.UserMetadataRepository,
	validationService *service.UserValidationService,
	txManager transaction.Manager,
) UserMetadataUseCase {
	return &userMetadataUseCase{
		userRepo:          userRepo,
		metadataRepo:      metadataRepo,
		validationService: validationService,
		txManager:         txManager,
	}
}

// GetUserMetadata ユーザーメタデータを取得
func (u *userMetadataUseCase) GetUserMetadata(ctx context.Context, userID uuid.UUID) (*user.UserMetadataResponse, error) {
	// ユーザーの存在確認
	userEntity, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, domainUser.ErrUserNotFound
	}

	metadata, err := u.metadataRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if metadata == nil {
		return nil, domainUser.ErrUserMetadataNotFound
	}

	return user.UserMetadataResponseFromEntity(metadata), nil
}

// CreateUserMetadata ユーザーメタデータを作成
func (u *userMetadataUseCase) CreateUserMetadata(ctx context.Context, userID uuid.UUID, req *user.CreateUserMetadataRequest) (*user.UserMetadataResponse, error) {
	var response *user.UserMetadataResponse
	
	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ユーザーの存在確認
		userEntity, err := u.userRepo.GetByID(txCtx, userID)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return domainUser.ErrUserNotFound
		}

		// 既存メタデータの確認
		existingMetadata, err := u.metadataRepo.GetByUserID(txCtx, userID)
		if err != nil {
			log.Printf("Failed to check existing metadata for userID %s: %v", userID, err)
			return err
		}
		if existingMetadata != nil {
			return domainUser.ErrUserMetadataAlreadyExists
		}

		// デフォルト値の設定
		timezone := "Asia/Tokyo"
		if req.Timezone != "" {
			timezone = req.Timezone
		}

		visionPublic := value.PublicFalse
		if req.VisionPublic != nil && *req.VisionPublic {
			visionPublic = value.PublicTrue
		}

		skills := req.Skills
		if skills == nil {
			skills = []string{}
		}

		interests := req.Interests
		if interests == nil {
			interests = []string{}
		}

		// メタデータエンティティの作成
		metadata := entity.NewUserMetadata(userID)
		
		// フィールドの設定
		if req.DisplayName != nil && *req.DisplayName != "" {
			metadata.DisplayName = req.DisplayName
		}
		if req.ProfileImageURL != nil && *req.ProfileImageURL != "" {
			metadata.ProfileImageURL = req.ProfileImageURL
		}
		if req.Tagline != nil && *req.Tagline != "" {
			metadata.Tagline = req.Tagline
		}
		if req.Bio != nil && *req.Bio != "" {
			metadata.Bio = req.Bio
		}
		if req.Vision != nil && *req.Vision != "" {
			metadata.Vision = req.Vision
		}
		metadata.VisionPublic = visionPublic
		metadata.Timezone = timezone
		metadata.Skills = skills
		metadata.Interests = interests

		// ユーザーのバリデーション
		if err := u.validationService.ValidateUserMetadata(metadata); err != nil {
			return err
		}

		// データベースに保存
		if err := u.metadataRepo.Create(txCtx, metadata); err != nil {
			return err
		}

		// DTOに変換
		response = user.UserMetadataResponseFromEntity(metadata)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateUserMetadata ユーザーメタデータを更新
func (u *userMetadataUseCase) UpdateUserMetadata(ctx context.Context, userID uuid.UUID, req *user.UpdateUserMetadataRequest) (*user.UserMetadataResponse, error) {
	var response *user.UserMetadataResponse
	
	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ユーザーの存在確認
		userEntity, err := u.userRepo.GetByID(txCtx, userID)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return domainUser.ErrUserNotFound
		}

		// ユーザーを取得
		metadata, err := u.metadataRepo.GetByUserID(txCtx, userID)
		if err != nil {
			return err
		}
		if metadata == nil {
			return domainUser.ErrUserMetadataNotFound
		}

		// フィールドの更新
		if req.DisplayName != nil {
			metadata.DisplayName = req.DisplayName
		}
		if req.ProfileImageURL != nil {
			metadata.ProfileImageURL = req.ProfileImageURL
		}
		if req.Tagline != nil {
			metadata.Tagline = req.Tagline
		}
		if req.Bio != nil {
			metadata.Bio = req.Bio
		}
		if req.Vision != nil {
			metadata.Vision = req.Vision
		}
		if req.VisionPublic != nil {
			if *req.VisionPublic {
				metadata.VisionPublic = value.PublicTrue
			} else {
				metadata.VisionPublic = value.PublicFalse
			}
		}
		if req.Timezone != nil {
			metadata.Timezone = *req.Timezone
		}
		if req.Skills != nil {
			metadata.Skills = req.Skills
		}
		if req.Interests != nil {
			metadata.Interests = req.Interests
		}

		// ユーザーのバリデーション
		if err := u.validationService.ValidateUserMetadata(metadata); err != nil {
			return err
		}

		// データベースに保存
		if err := u.metadataRepo.Update(txCtx, metadata); err != nil {
			return err
		}

		// DTOに変換
		response = user.UserMetadataResponseFromEntity(metadata)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}