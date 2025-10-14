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

type userRivalRepository struct {
	db *gorm.DB
}

// NewUserRivalRepository コンストラクタ
func NewUserRivalRepository(db *gorm.DB) repository.UserRivalRepository {
	return &userRivalRepository{db: db}
}

// Create ユーザーライバルを作成する
func (r *userRivalRepository) Create(ctx context.Context, rival *entity.UserRival) error {
	gormRival := r.toGORMUserRival(rival)
	return r.db.WithContext(ctx).Create(gormRival).Error
}

// GetByID IDでライバルを取得する
func (r *userRivalRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.UserRival, error) {
	var gormRival userModel.UserRival
	err := r.db.WithContext(ctx).First(&gormRival, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMUserRival(&gormRival)
}

// GetByUserID ユーザーIDでライバル一覧を取得する
func (r *userRivalRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserRival, error) {
	var gormRivals []userModel.UserRival
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&gormRivals).Error
	if err != nil {
		return nil, err
	}

	rivals := make([]*entity.UserRival, len(gormRivals))
	for i, gormRival := range gormRivals {
		rival, err := r.fromGORMUserRival(&gormRival)
		if err != nil {
			return nil, err
		}
		rivals[i] = rival
	}
	return rivals, nil
}

// GetByUserAndRival ユーザーIDとライバルIDで取得する
func (r *userRivalRepository) GetByUserAndRival(ctx context.Context, userID, rivalID uuid.UUID) (*entity.UserRival, error) {
	var gormRival userModel.UserRival
	err := r.db.WithContext(ctx).Where("user_id = ? AND rival_id = ?", userID, rivalID).First(&gormRival).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMUserRival(&gormRival)
}

// CountByUserID ユーザーのライバル数を取得する
func (r *userRivalRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&userModel.UserRival{}).Where("user_id = ?", userID).Count(&count).Error
	return int(count), err
}

// Delete ユーザーライバルを削除する
func (r *userRivalRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&userModel.UserRival{}, "id = ?", id).Error
}

// DeleteByUserAndRival ユーザーIDとライバルIDで削除する
func (r *userRivalRepository) DeleteByUserAndRival(ctx context.Context, userID, rivalID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&userModel.UserRival{}, "user_id = ? AND rival_id = ?", userID, rivalID).Error
}

// 型変換: Domain Entity → GORM Model
func (r *userRivalRepository) toGORMUserRival(rival *entity.UserRival) *userModel.UserRival {
	return &userModel.UserRival{
		ID:        rival.ID(),
		UserID:    rival.UserID(),
		RivalID:   rival.RivalID(),
		CreatedAt: rival.CreatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *userRivalRepository) fromGORMUserRival(gormRival *userModel.UserRival) (*entity.UserRival, error) {
	// TODO: entity.NewUserRival の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}