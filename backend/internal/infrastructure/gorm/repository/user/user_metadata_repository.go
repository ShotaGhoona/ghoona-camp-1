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

type userMetadataRepository struct {
	db *gorm.DB
}

// NewUserMetadataRepository コンストラクタ
func NewUserMetadataRepository(db *gorm.DB) repository.UserMetadataRepository {
	return &userMetadataRepository{db: db}
}

// Create ユーザーメタデータを作成する
func (r *userMetadataRepository) Create(ctx context.Context, metadata *entity.UserMetadata) error {
	gormMetadata := r.toGORMUserMetadata(metadata)
	return r.db.WithContext(ctx).Create(gormMetadata).Error
}

// GetByUserID ユーザーIDでメタデータを取得する
func (r *userMetadataRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserMetadata, error) {
	var gormMetadata userModel.UserMetadata
	err := r.db.WithContext(ctx).First(&gormMetadata, "user_id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMUserMetadata(&gormMetadata)
}

// Update ユーザーメタデータを更新する
func (r *userMetadataRepository) Update(ctx context.Context, metadata *entity.UserMetadata) error {
	gormMetadata := r.toGORMUserMetadata(metadata)
	return r.db.WithContext(ctx).Save(gormMetadata).Error
}

// Delete ユーザーメタデータを削除する
func (r *userMetadataRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&userModel.UserMetadata{}, "user_id = ?", userID).Error
}

// Upsert ユーザーメタデータを作成または更新する
func (r *userMetadataRepository) Upsert(ctx context.Context, metadata *entity.UserMetadata) error {
	gormMetadata := r.toGORMUserMetadata(metadata)
	return r.db.WithContext(ctx).Save(gormMetadata).Error
}

// 型変換: Domain Entity → GORM Model
func (r *userMetadataRepository) toGORMUserMetadata(metadata *entity.UserMetadata) *userModel.UserMetadata {
	return &userModel.UserMetadata{
		UserID:              metadata.UserID(),
		DisplayName:         metadata.DisplayName(),
		Bio:                 metadata.Bio(),
		Location:            metadata.Location(),
		Website:             metadata.Website(),
		DateOfBirth:         metadata.DateOfBirth(),
		Gender:              metadata.Gender(),
		Occupation:          metadata.Occupation(),
		EducationLevel:      metadata.EducationLevel(),
		Interests:           metadata.Interests(),
		Timezone:            metadata.Timezone(),
		LanguagePreference:  metadata.LanguagePreference(),
		PrivacyLevel:        metadata.PrivacyLevel(),
		NotificationPrefs:   metadata.NotificationPrefs(),
		CustomFields:        metadata.CustomFields(),
		CreatedAt:           metadata.CreatedAt(),
		UpdatedAt:           metadata.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *userMetadataRepository) fromGORMUserMetadata(gormMetadata *userModel.UserMetadata) (*entity.UserMetadata, error) {
	// TODO: entity.NewUserMetadata の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}