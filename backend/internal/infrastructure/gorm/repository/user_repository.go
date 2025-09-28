package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/user/entity"
	"ghoona-camp-backend/internal/domain/user/repository"
	baseGorm "ghoona-camp-backend/internal/infrastructure/gorm"
	"ghoona-camp-backend/internal/infrastructure/gorm/model"
)

// userRepository はUserRepositoryインターフェースの実装
type userRepository struct {
	*baseGorm.BaseRepository
}

// NewUserRepository は新しいUserRepositoryを作成する
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		BaseRepository: baseGorm.NewBaseRepository(db),
	}
}

// GetByID はIDでユーザーを取得する
func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var gormUser model.User
	db := r.GetDB(ctx)
	err := db.Where("id = ?", id).First(&gormUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return gormUser.ToEntity()
}

// GetByClerkID はClerk IDでユーザーを取得する
func (r *userRepository) GetByClerkID(ctx context.Context, clerkID string) (*entity.User, error) {
	var gormUser model.User
	db := r.GetDB(ctx)
	err := db.Where("clerk_id = ?", clerkID).First(&gormUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return gormUser.ToEntity()
}

// GetByEmail はメールアドレスでユーザーを取得する
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var gormUser model.User
	db := r.GetDB(ctx)
	err := db.Where("email = ?", email).First(&gormUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return gormUser.ToEntity()
}

// Create はユーザーを作成する
func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	gormUser := model.FromEntity(user)
	db := r.GetDB(ctx)
	err := db.Create(gormUser).Error
	if err != nil {
		return err
	}
	// 作成後のIDを反映
	user.ID = gormUser.ID
	return nil
}

// Update はユーザーを更新する
func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	gormUser := model.FromEntity(user)
	db := r.GetDB(ctx)
	return db.Save(gormUser).Error
}

// GetAll は全ユーザーを取得する
func (r *userRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	var gormUsers []model.User
	db := r.GetDB(ctx)
	
	// 基本クエリ（アクティブなユーザーのみ）
	err := db.Where("is_active = ?", true).Order("created_at DESC").Find(&gormUsers).Error
	if err != nil {
		return nil, err
	}
	
	// エンティティに変換
	users := make([]*entity.User, len(gormUsers))
	for i, gormUser := range gormUsers {
		entityUser, err := gormUser.ToEntity()
		if err != nil {
			return nil, err
		}
		users[i] = entityUser
	}
	
	return users, nil
}

// Delete はユーザーを削除する
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	db := r.GetDB(ctx)
	return db.Delete(&model.User{}, id).Error
}

// userMetadataRepository はUserMetadataRepositoryインターフェースの実装
type userMetadataRepository struct {
	*baseGorm.BaseRepository
}

// NewUserMetadataRepository は新しいUserMetadataRepositoryを作成する
func NewUserMetadataRepository(db *gorm.DB) repository.UserMetadataRepository {
	return &userMetadataRepository{
		BaseRepository: baseGorm.NewBaseRepository(db),
	}
}

// GetByUserID はユーザーIDでメタデータを取得する
func (r *userMetadataRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserMetadata, error) {
	var gormMetadata model.UserMetadata
	db := r.GetDB(ctx)
	err := db.Where("user_id = ?", userID).First(&gormMetadata).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return gormMetadata.ToEntity()
}

// Create はメタデータを作成する
func (r *userMetadataRepository) Create(ctx context.Context, metadata *entity.UserMetadata) error {
	gormMetadata := model.FromEntityUserMetadata(metadata)
	db := r.GetDB(ctx)
	err := db.Create(gormMetadata).Error
	if err != nil {
		return err
	}
	metadata.ID = gormMetadata.ID
	return nil
}

// Update はメタデータを更新する
func (r *userMetadataRepository) Update(ctx context.Context, metadata *entity.UserMetadata) error {
	gormMetadata := model.FromEntityUserMetadata(metadata)
	db := r.GetDB(ctx)
	return db.Save(gormMetadata).Error
}

// Delete はメタデータを削除する
func (r *userMetadataRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	db := r.GetDB(ctx)
	return db.Where("user_id = ?", userID).Delete(&model.UserMetadata{}).Error
}

// userSocialLinkRepository はUserSocialLinkRepositoryインターフェースの実装
type userSocialLinkRepository struct {
	*baseGorm.BaseRepository
}

// NewUserSocialLinkRepository は新しいUserSocialLinkRepositoryを作成する
func NewUserSocialLinkRepository(db *gorm.DB) repository.UserSocialLinkRepository {
	return &userSocialLinkRepository{
		BaseRepository: baseGorm.NewBaseRepository(db),
	}
}

// GetByUserID はユーザーIDでソーシャルリンク一覧を取得する
func (r *userSocialLinkRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserSocialLink, error) {
	var gormLinks []model.UserSocialLink
	db := r.GetDB(ctx)
	err := db.Where("user_id = ?", userID).Find(&gormLinks).Error
	if err != nil {
		return nil, err
	}

	links := make([]*entity.UserSocialLink, len(gormLinks))
	for i, gormLink := range gormLinks {
		entityLink, err := gormLink.ToEntity()
		if err != nil {
			return nil, err
		}
		links[i] = entityLink
	}
	return links, nil
}

// GetByID はIDでソーシャルリンクを取得する
func (r *userSocialLinkRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.UserSocialLink, error) {
	var gormLink model.UserSocialLink
	db := r.GetDB(ctx)
	err := db.Where("id = ?", id).First(&gormLink).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return gormLink.ToEntity()
}

// Create はソーシャルリンクを作成する
func (r *userSocialLinkRepository) Create(ctx context.Context, link *entity.UserSocialLink) error {
	gormLink := model.FromEntityUserSocialLink(link)
	db := r.GetDB(ctx)
	err := db.Create(gormLink).Error
	if err != nil {
		return err
	}
	link.ID = gormLink.ID
	return nil
}

// Update はソーシャルリンクを更新する
func (r *userSocialLinkRepository) Update(ctx context.Context, link *entity.UserSocialLink) error {
	gormLink := model.FromEntityUserSocialLink(link)
	db := r.GetDB(ctx)
	return db.Save(gormLink).Error
}

// Delete はソーシャルリンクを削除する
func (r *userSocialLinkRepository) Delete(ctx context.Context, id uuid.UUID) error {
	db := r.GetDB(ctx)
	return db.Delete(&model.UserSocialLink{}, id).Error
}

// userRivalRepository はUserRivalRepositoryインターフェースの実装
type userRivalRepository struct {
	*baseGorm.BaseRepository
}

// NewUserRivalRepository は新しいUserRivalRepositoryを作成する
func NewUserRivalRepository(db *gorm.DB) repository.UserRivalRepository {
	return &userRivalRepository{
		BaseRepository: baseGorm.NewBaseRepository(db),
	}
}

// GetByUserID はユーザーIDでライバル一覧を取得する
func (r *userRivalRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserRival, error) {
	var gormRivals []model.UserRival
	db := r.GetDB(ctx)
	err := db.Where("user_id = ?", userID).Find(&gormRivals).Error
	if err != nil {
		return nil, err
	}

	rivals := make([]*entity.UserRival, len(gormRivals))
	for i, gormRival := range gormRivals {
		entityRival, err := gormRival.ToEntity()
		if err != nil {
			return nil, err
		}
		rivals[i] = entityRival
	}
	return rivals, nil
}

// CountByUserID はユーザーIDでライバル数をカウントする
func (r *userRivalRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int64
	db := r.GetDB(ctx)
	err := db.Model(&model.UserRival{}).Where("user_id = ?", userID).Count(&count).Error
	return int(count), err
}

// Create はライバル関係を作成する
func (r *userRivalRepository) Create(ctx context.Context, rival *entity.UserRival) error {
	gormRival := model.FromEntityUserRival(rival)
	db := r.GetDB(ctx)
	err := db.Create(gormRival).Error
	if err != nil {
		return err
	}
	rival.ID = gormRival.ID
	return nil
}

// Delete はライバル関係を削除する
func (r *userRivalRepository) Delete(ctx context.Context, id uuid.UUID) error {
	db := r.GetDB(ctx)
	return db.Delete(&model.UserRival{}, id).Error
}
