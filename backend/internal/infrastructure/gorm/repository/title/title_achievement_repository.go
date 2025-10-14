package title

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/title/entity"
	"ghoona-camp-backend/internal/domain/title/repository"
	titleModel "ghoona-camp-backend/internal/infrastructure/gorm/model/title"
)

type titleAchievementRepository struct {
	db *gorm.DB
}

// NewTitleAchievementRepository コンストラクタ
func NewTitleAchievementRepository(db *gorm.DB) repository.TitleAchievementRepository {
	return &titleAchievementRepository{db: db}
}

// Create タイトル達成を作成する
func (r *titleAchievementRepository) Create(ctx context.Context, achievement *entity.TitleAchievement) error {
	gormAchievement := r.toGORMTitleAchievement(achievement)
	return r.db.WithContext(ctx).Create(gormAchievement).Error
}

// GetByID IDで達成を取得する
func (r *titleAchievementRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.TitleAchievement, error) {
	var gormAchievement titleModel.TitleAchievement
	err := r.db.WithContext(ctx).First(&gormAchievement, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMTitleAchievement(&gormAchievement)
}

// GetByUserID ユーザーIDで達成一覧を取得する
func (r *titleAchievementRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.TitleAchievement, error) {
	var gormAchievements []titleModel.TitleAchievement
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&gormAchievements).Error
	if err != nil {
		return nil, err
	}

	achievements := make([]*entity.TitleAchievement, len(gormAchievements))
	for i, gormAchievement := range gormAchievements {
		achievement, err := r.fromGORMTitleAchievement(&gormAchievement)
		if err != nil {
			return nil, err
		}
		achievements[i] = achievement
	}
	return achievements, nil
}

// GetByTitleID タイトルIDで達成一覧を取得する
func (r *titleAchievementRepository) GetByTitleID(ctx context.Context, titleID uuid.UUID) ([]*entity.TitleAchievement, error) {
	var gormAchievements []titleModel.TitleAchievement
	err := r.db.WithContext(ctx).Where("title_id = ?", titleID).Find(&gormAchievements).Error
	if err != nil {
		return nil, err
	}

	achievements := make([]*entity.TitleAchievement, len(gormAchievements))
	for i, gormAchievement := range gormAchievements {
		achievement, err := r.fromGORMTitleAchievement(&gormAchievement)
		if err != nil {
			return nil, err
		}
		achievements[i] = achievement
	}
	return achievements, nil
}

// GetByUserAndTitle ユーザーIDとタイトルIDで達成を取得する
func (r *titleAchievementRepository) GetByUserAndTitle(ctx context.Context, userID, titleID uuid.UUID) (*entity.TitleAchievement, error) {
	var gormAchievement titleModel.TitleAchievement
	err := r.db.WithContext(ctx).Where("user_id = ? AND title_id = ?", userID, titleID).First(&gormAchievement).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMTitleAchievement(&gormAchievement)
}

// GetActiveByUserID ユーザーのアクティブタイトルを取得する
func (r *titleAchievementRepository) GetActiveByUserID(ctx context.Context, userID uuid.UUID) (*entity.TitleAchievement, error) {
	var gormAchievement titleModel.TitleAchievement
	err := r.db.WithContext(ctx).Where("user_id = ? AND is_active = ?", userID, true).First(&gormAchievement).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMTitleAchievement(&gormAchievement)
}

// Update タイトル達成を更新する
func (r *titleAchievementRepository) Update(ctx context.Context, achievement *entity.TitleAchievement) error {
	gormAchievement := r.toGORMTitleAchievement(achievement)
	return r.db.WithContext(ctx).Save(gormAchievement).Error
}

// Delete タイトル達成を削除する
func (r *titleAchievementRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&titleModel.TitleAchievement{}, "id = ?", id).Error
}

// SetActiveTitle ユーザーのアクティブタイトルを設定する
func (r *titleAchievementRepository) SetActiveTitle(ctx context.Context, userID, titleID uuid.UUID) error {
	// TODO: トランザクション処理で現在のActiveを無効化して新しいタイトルをアクティブに設定
	return errors.New("TODO: トランザクション処理実装必要")
}

// 型変換: Domain Entity → GORM Model
func (r *titleAchievementRepository) toGORMTitleAchievement(achievement *entity.TitleAchievement) *titleModel.TitleAchievement {
	return &titleModel.TitleAchievement{
		ID:          achievement.ID(),
		UserID:      achievement.UserID(),
		TitleID:     achievement.TitleID(),
		IsActive:    achievement.IsActive(),
		AchievedAt:  achievement.AchievedAt(),
		ActivatedAt: achievement.ActivatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *titleAchievementRepository) fromGORMTitleAchievement(gormAchievement *titleModel.TitleAchievement) (*entity.TitleAchievement, error) {
	// TODO: entity.NewTitleAchievement の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}