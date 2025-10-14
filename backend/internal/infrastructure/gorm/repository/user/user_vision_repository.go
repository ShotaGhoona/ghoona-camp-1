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

type userVisionRepository struct {
	db *gorm.DB
}

// NewUserVisionRepository コンストラクタ
func NewUserVisionRepository(db *gorm.DB) repository.UserVisionRepository {
	return &userVisionRepository{db: db}
}

// Create ユーザービジョンを作成する
func (r *userVisionRepository) Create(ctx context.Context, vision *entity.UserVision) error {
	gormVision := r.toGORMUserVision(vision)
	return r.db.WithContext(ctx).Create(gormVision).Error
}

// GetByID IDでビジョンを取得する
func (r *userVisionRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.UserVision, error) {
	var gormVision userModel.UserVision
	err := r.db.WithContext(ctx).First(&gormVision, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMUserVision(&gormVision)
}

// GetByUserID ユーザーIDでビジョン一覧を取得する
func (r *userVisionRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserVision, error) {
	var gormVisions []userModel.UserVision
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&gormVisions).Error
	if err != nil {
		return nil, err
	}

	visions := make([]*entity.UserVision, len(gormVisions))
	for i, gormVision := range gormVisions {
		vision, err := r.fromGORMUserVision(&gormVision)
		if err != nil {
			return nil, err
		}
		visions[i] = vision
	}
	return visions, nil
}

// GetByUserIDAndStatus ユーザーIDとステータスでビジョンを取得する
func (r *userVisionRepository) GetByUserIDAndStatus(ctx context.Context, userID uuid.UUID, status string) ([]*entity.UserVision, error) {
	var gormVisions []userModel.UserVision
	err := r.db.WithContext(ctx).Where("user_id = ? AND status = ?", userID, status).Find(&gormVisions).Error
	if err != nil {
		return nil, err
	}

	visions := make([]*entity.UserVision, len(gormVisions))
	for i, gormVision := range gormVisions {
		vision, err := r.fromGORMUserVision(&gormVision)
		if err != nil {
			return nil, err
		}
		visions[i] = vision
	}
	return visions, nil
}

// Update ユーザービジョンを更新する
func (r *userVisionRepository) Update(ctx context.Context, vision *entity.UserVision) error {
	gormVision := r.toGORMUserVision(vision)
	return r.db.WithContext(ctx).Save(gormVision).Error
}

// Delete ユーザービジョンを削除する
func (r *userVisionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&userModel.UserVision{}, "id = ?", id).Error
}

// 型変換: Domain Entity → GORM Model
func (r *userVisionRepository) toGORMUserVision(vision *entity.UserVision) *userModel.UserVision {
	return &userModel.UserVision{
		ID:          vision.ID(),
		UserID:      vision.UserID(),
		Title:       vision.Title(),
		Description: vision.Description(),
		Status:      vision.Status(),
		TargetDate:  vision.TargetDate(),
		CreatedAt:   vision.CreatedAt(),
		UpdatedAt:   vision.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *userVisionRepository) fromGORMUserVision(gormVision *userModel.UserVision) (*entity.UserVision, error) {
	// TODO: entity.NewUserVision の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}