package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/user/entity"
	"ghoona-camp-backend/internal/domain/user/repository"
	userModel "ghoona-camp-backend/internal/infrastructure/gorm/model/user"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository コンストラクタ
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

// Create ユーザーを作成する
func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	gormUser := r.toGORMUser(user)
	return r.db.WithContext(ctx).Create(gormUser).Error
}

// GetByID IDでユーザーを取得する
func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var gormUser userModel.User
	err := r.db.WithContext(ctx).First(&gormUser, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMUser(&gormUser)
}

// GetByClerkID Clerk IDでユーザーを取得する
func (r *userRepository) GetByClerkID(ctx context.Context, clerkID string) (*entity.User, error) {
	var gormUser userModel.User
	err := r.db.WithContext(ctx).Where("clerk_id = ?", clerkID).First(&gormUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMUser(&gormUser)
}

// GetByEmail メールアドレスでユーザーを取得する
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var gormUser userModel.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&gormUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMUser(&gormUser)
}

// Update ユーザーを更新する
func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	gormUser := r.toGORMUser(user)
	return r.db.WithContext(ctx).Save(gormUser).Error
}

// Delete ユーザーを削除する
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&userModel.User{}, "id = ?", id).Error
}

// List ユーザー一覧を取得する
func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	var gormUsers []userModel.User
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&gormUsers).Error
	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, len(gormUsers))
	for i, gormUser := range gormUsers {
		user, err := r.fromGORMUser(&gormUser)
		if err != nil {
			return nil, err
		}
		users[i] = user
	}
	return users, nil
}

// Search ユーザーを検索する
func (r *userRepository) Search(ctx context.Context, query string, limit, offset int) ([]*entity.User, error) {
	// TODO: 検索仕様確定後実装（検索対象フィールド：username? email? display_name?）
	return nil, errors.New("TODO: 要件確定後実装")
}

// 型変換: Domain Entity → GORM Model
func (r *userRepository) toGORMUser(user *entity.User) *userModel.User {
	return &userModel.User{
		ID:        user.ID(),
		ClerkID:   user.ClerkID(),
		Email:     user.Email(),
		Username:  user.Username(),
		AvatarURL: user.AvatarURL(),
		DiscordID: user.DiscordID(),
		IsActive:  user.IsActive(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *userRepository) fromGORMUser(gormUser *userModel.User) (*entity.User, error) {
	// TODO: entity.NewUser の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}