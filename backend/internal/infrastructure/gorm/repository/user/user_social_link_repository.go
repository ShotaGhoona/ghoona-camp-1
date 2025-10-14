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

type userSocialLinkRepository struct {
	db *gorm.DB
}

// NewUserSocialLinkRepository コンストラクタ
func NewUserSocialLinkRepository(db *gorm.DB) repository.UserSocialLinkRepository {
	return &userSocialLinkRepository{db: db}
}

// Create ユーザーソーシャルリンクを作成する
func (r *userSocialLinkRepository) Create(ctx context.Context, link *entity.UserSocialLink) error {
	gormLink := r.toGORMUserSocialLink(link)
	return r.db.WithContext(ctx).Create(gormLink).Error
}

// GetByID IDでソーシャルリンクを取得する
func (r *userSocialLinkRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.UserSocialLink, error) {
	var gormLink userModel.UserSocialLink
	err := r.db.WithContext(ctx).First(&gormLink, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMUserSocialLink(&gormLink)
}

// GetByUserID ユーザーIDでソーシャルリンク一覧を取得する
func (r *userSocialLinkRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserSocialLink, error) {
	var gormLinks []userModel.UserSocialLink
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&gormLinks).Error
	if err != nil {
		return nil, err
	}

	links := make([]*entity.UserSocialLink, len(gormLinks))
	for i, gormLink := range gormLinks {
		link, err := r.fromGORMUserSocialLink(&gormLink)
		if err != nil {
			return nil, err
		}
		links[i] = link
	}
	return links, nil
}

// Update ユーザーソーシャルリンクを更新する
func (r *userSocialLinkRepository) Update(ctx context.Context, link *entity.UserSocialLink) error {
	gormLink := r.toGORMUserSocialLink(link)
	return r.db.WithContext(ctx).Save(gormLink).Error
}

// Delete ユーザーソーシャルリンクを削除する
func (r *userSocialLinkRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&userModel.UserSocialLink{}, "id = ?", id).Error
}

// 型変換: Domain Entity → GORM Model
func (r *userSocialLinkRepository) toGORMUserSocialLink(link *entity.UserSocialLink) *userModel.UserSocialLink {
	return &userModel.UserSocialLink{
		ID:        link.ID(),
		UserID:    link.UserID(),
		Platform:  link.Platform(),
		URL:       link.URL(),
		CreatedAt: link.CreatedAt(),
		UpdatedAt: link.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *userSocialLinkRepository) fromGORMUserSocialLink(gormLink *userModel.UserSocialLink) (*entity.UserSocialLink, error) {
	// TODO: entity.NewUserSocialLink の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}