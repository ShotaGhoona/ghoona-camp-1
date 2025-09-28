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
)

// UserUseCase ユーザー基本操作のユースケース
type UserUseCase interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*user.UserResponse, error)
	GetUserByClerkID(ctx context.Context, clerkID string) (*user.UserResponse, error)
	GetUsers(ctx context.Context) (*user.UserListResponse, error)
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, req *user.UpdateUserRequest) (*user.UserResponse, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUserRepo() repository.UserRepository
}

type userUseCase struct {
	userRepo          repository.UserRepository
	metadataRepo      repository.UserMetadataRepository
	socialLinkRepo    repository.UserSocialLinkRepository
	rivalRepo         repository.UserRivalRepository
	userService       *service.UserService
	validationService *service.UserValidationService
	txManager         transaction.Manager
}

// NewUserUseCase 新しいUserUseCaseを作成
func NewUserUseCase(
	userRepo repository.UserRepository,
	metadataRepo repository.UserMetadataRepository,
	socialLinkRepo repository.UserSocialLinkRepository,
	rivalRepo repository.UserRivalRepository,
	userService *service.UserService,
	validationService *service.UserValidationService,
	txManager transaction.Manager,
) UserUseCase {
	return &userUseCase{
		userRepo:          userRepo,
		metadataRepo:      metadataRepo,
		socialLinkRepo:    socialLinkRepo,
		rivalRepo:         rivalRepo,
		userService:       userService,
		validationService: validationService,
		txManager:         txManager,
	}
}

// GetUserByID IDでユーザーを取得
func (u *userUseCase) GetUserByID(ctx context.Context, userID uuid.UUID) (*user.UserResponse, error) {
	userEntity, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, domainUser.ErrUserNotFound
	}

	// メタデータを取得
	metadata, _ := u.metadataRepo.GetByUserID(ctx, userID)

	return user.UserResponseFromEntityWithMetadata(userEntity, metadata), nil
}

// GetUserByClerkID Clerk IDでユーザーを取得
func (u *userUseCase) GetUserByClerkID(ctx context.Context, clerkID string) (*user.UserResponse, error) {
	userEntity, err := u.userRepo.GetByClerkID(ctx, clerkID)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, domainUser.ErrUserNotFound
	}

	// メタデータを取得
	metadata, _ := u.metadataRepo.GetByUserID(ctx, userEntity.ID)

	return user.UserResponseFromEntityWithMetadata(userEntity, metadata), nil
}

// GetUsers 全ユーザーを取得
func (u *userUseCase) GetUsers(ctx context.Context) (*user.UserListResponse, error) {
	// 全ユーザーを取得
	users, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// 既存のヘルパー関数を使用してDTOに変換
	userResponses := make([]user.UserResponse, len(users))
	for i, usr := range users {
		response := user.UserResponseFromEntity(usr)
		if response != nil {
			userResponses[i] = *response
		}
	}

	return &user.UserListResponse{
		Users: userResponses,
		Total: len(users),
	}, nil
}

// CreateUser ユーザーを作成
func (u *userUseCase) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {
	var response *user.UserResponse
	
	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// メールアドレスの重複チェック
		existingUser, err := u.userRepo.GetByEmail(txCtx, req.Email)
		if err != nil {
			return err
		}
		if existingUser != nil {
			return domainUser.ErrDuplicateEmail
		}

		// Clerk IDの重複チェック
		existingUserByClerk, err := u.userRepo.GetByClerkID(txCtx, req.ClerkID)
		if err != nil {
			return err
		}
		if existingUserByClerk != nil {
			return domainUser.ErrDuplicateEmail // Clerk IDが重複している場合のエラー
		}

		// ユーザーエンティティの作成
		userEntity := entity.NewUser(req.ClerkID, req.Email)
		
		// オプショナルフィールドの設定
		if req.Username != nil {
			userEntity.Username = req.Username
		}
		if req.AvatarURL != nil {
			userEntity.AvatarURL = req.AvatarURL
		}
		if req.DiscordID != nil {
			userEntity.DiscordID = req.DiscordID
		}

		// ユーザーのバリデーション
		if err := u.validationService.ValidateUser(userEntity); err != nil {
			return err
		}

		// データベースに保存
		if err := u.userRepo.Create(txCtx, userEntity); err != nil {
			return err
		}

		// 作成されたユーザーを取得
		createdUser, err := u.userRepo.GetByClerkID(txCtx, userEntity.ClerkID)
		if err != nil {
			return err
		}

		// DTOに変換
		response = user.UserResponseFromEntity(createdUser)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateUser ユーザー情報を更新
func (u *userUseCase) UpdateUser(ctx context.Context, userID uuid.UUID, req *user.UpdateUserRequest) (*user.UserResponse, error) {
	var response *user.UserResponse
	
	err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ユーザーを取得
		userEntity, err := u.userRepo.GetByID(txCtx, userID)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return domainUser.ErrUserNotFound
		}

		// ユーザー情報を更新
		if req.Username != nil {
			userEntity.Username = req.Username
		}
		if req.AvatarURL != nil {
			userEntity.AvatarURL = req.AvatarURL
		}
		if req.DiscordID != nil {
			userEntity.DiscordID = req.DiscordID
		}

		// ユーザーのバリデーション
		if err := u.validationService.ValidateUser(userEntity); err != nil {
			return err
		}

		// データベースに保存
		if err := u.userRepo.Update(txCtx, userEntity); err != nil {
			return err
		}

		// DTOに変換
		response = user.UserResponseFromEntity(userEntity)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteUser ユーザーを削除
func (u *userUseCase) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ユーザーの存在確認
		userEntity, err := u.userRepo.GetByID(txCtx, userID)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return domainUser.ErrUserNotFound
		}

		// 関連するメタデータを削除
		_ = u.metadataRepo.Delete(txCtx, userID)

		// ソーシャルリンクを削除
		socialLinks, _ := u.socialLinkRepo.GetByUserID(txCtx, userID)
		for _, link := range socialLinks {
			_ = u.socialLinkRepo.Delete(txCtx, link.ID)
		}

		// ライバルを削除
		rivals, _ := u.rivalRepo.GetByUserID(txCtx, userID)
		for _, rival := range rivals {
			_ = u.rivalRepo.Delete(txCtx, rival.ID)
		}

		// ユーザーを削除
		return u.userRepo.Delete(txCtx, userID)
	})
}

// GetUserRepo UserRepositoryを取得（Controller用）
func (u *userUseCase) GetUserRepo() repository.UserRepository {
	return u.userRepo
}