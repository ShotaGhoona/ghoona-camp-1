package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title/entity"
	"ghoona-camp-backend/internal/domain/title/repository"
	baseGorm "ghoona-camp-backend/internal/infrastructure/gorm"
	"ghoona-camp-backend/internal/infrastructure/gorm/model"
)

// titleRepository はTitleRepositoryインターフェースの実装
type titleRepository struct {
	*baseGorm.BaseRepository
}

// NewTitleRepository は新しいTitleRepositoryを作成する
func NewTitleRepository(db *gorm.DB) repository.TitleRepository {
	return &titleRepository{
		BaseRepository: baseGorm.NewBaseRepository(db),
	}
}

// GetAll は全称号を取得する（レベル順）
func (r *titleRepository) GetAll(ctx context.Context) ([]*entity.Title, error) {
	var gormTitles []model.Title
	db := r.GetDB(ctx)
	
	err := db.Order("level ASC").Find(&gormTitles).Error
	if err != nil {
		return nil, err
	}
	
	return r.convertToEntities(gormTitles)
}

// GetByID はIDで称号を取得する
func (r *titleRepository) GetByID(ctx context.Context, id common.UUID) (*entity.Title, error) {
	var gormTitle model.Title
	db := r.GetDB(ctx)
	err := db.Where("id = ?", uuid.UUID(id)).First(&gormTitle).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return gormTitle.ToEntity()
}

// GetByLevel はレベルで称号を取得する
func (r *titleRepository) GetByLevel(ctx context.Context, level int) (*entity.Title, error) {
	var gormTitle model.Title
	db := r.GetDB(ctx)
	err := db.Where("level = ?", level).First(&gormTitle).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return gormTitle.ToEntity()
}

// GetActiveTitles はアクティブな称号のみを取得する（レベル順）
func (r *titleRepository) GetActiveTitles(ctx context.Context) ([]*entity.Title, error) {
	var gormTitles []model.Title
	db := r.GetDB(ctx)
	
	err := db.Where("is_active = ?", true).Order("level ASC").Find(&gormTitles).Error
	if err != nil {
		return nil, err
	}
	
	return r.convertToEntities(gormTitles)
}

// convertToEntities はGORMモデルリストからエンティティリストに変換する
func (r *titleRepository) convertToEntities(gormTitles []model.Title) ([]*entity.Title, error) {
	if len(gormTitles) == 0 {
		return []*entity.Title{}, nil
	}
	
	titles := make([]*entity.Title, len(gormTitles))
	for i, gormTitle := range gormTitles {
		entityTitle, err := gormTitle.ToEntity()
		if err != nil {
			return nil, err
		}
		titles[i] = entityTitle
	}
	
	return titles, nil
}

// titleAchievementRepository はTitleAchievementRepositoryインターフェースの実装
type titleAchievementRepository struct {
	*baseGorm.BaseRepository
}

// NewTitleAchievementRepository は新しいTitleAchievementRepositoryを作成する
func NewTitleAchievementRepository(db *gorm.DB) repository.TitleAchievementRepository {
	return &titleAchievementRepository{
		BaseRepository: baseGorm.NewBaseRepository(db),
	}
}

// GetByUserID はユーザーIDで称号獲得実績一覧を取得する
func (r *titleAchievementRepository) GetByUserID(ctx context.Context, userID common.UUID) ([]*entity.TitleAchievement, error) {
	var gormAchievements []model.TitleAchievement
	db := r.GetDB(ctx)
	
	err := db.Where("user_id = ?", uuid.UUID(userID)).
		Order("achieved_at DESC").
		Find(&gormAchievements).Error
	if err != nil {
		return nil, err
	}
	
	return r.convertAchievementsToEntities(gormAchievements)
}

// GetCurrentByUserID はユーザーの現在表示称号を取得する
func (r *titleAchievementRepository) GetCurrentByUserID(ctx context.Context, userID common.UUID) (*entity.TitleAchievement, error) {
	var gormAchievement model.TitleAchievement
	db := r.GetDB(ctx)
	err := db.Where("user_id = ? AND is_current = ?", uuid.UUID(userID), true).First(&gormAchievement).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return gormAchievement.ToEntity()
}

// GetByUserIDAndTitleID はユーザーIDと称号IDで特定の獲得実績を取得する
func (r *titleAchievementRepository) GetByUserIDAndTitleID(ctx context.Context, userID common.UUID, titleID common.UUID) (*entity.TitleAchievement, error) {
	var gormAchievement model.TitleAchievement
	db := r.GetDB(ctx)
	err := db.Where("user_id = ? AND title_id = ?", uuid.UUID(userID), uuid.UUID(titleID)).First(&gormAchievement).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return gormAchievement.ToEntity()
}

// Create は称号獲得実績を作成する
func (r *titleAchievementRepository) Create(ctx context.Context, achievement *entity.TitleAchievement) error {
	gormAchievement := model.FromEntityTitleAchievement(achievement)
	db := r.GetDB(ctx)
	err := db.Create(gormAchievement).Error
	if err != nil {
		return err
	}
	// 作成後のIDを反映
	achievement.ID = common.UUID(gormAchievement.ID)
	return nil
}

// Update は称号獲得実績を更新する
func (r *titleAchievementRepository) Update(ctx context.Context, achievement *entity.TitleAchievement) error {
	gormAchievement := model.FromEntityTitleAchievement(achievement)
	db := r.GetDB(ctx)
	return db.Save(gormAchievement).Error
}

// SetCurrent は現在表示称号を設定する（排他制御付き）
func (r *titleAchievementRepository) SetCurrent(ctx context.Context, userID common.UUID, titleID common.UUID) error {
	db := r.GetDB(ctx)
	
	return db.Transaction(func(tx *gorm.DB) error {
		// 既存の現在表示を解除
		err := tx.Model(&model.TitleAchievement{}).
			Where("user_id = ? AND is_current = ?", uuid.UUID(userID), true).
			Update("is_current", false).Error
		if err != nil {
			return err
		}
		
		// 指定称号を現在表示に設定
		return tx.Model(&model.TitleAchievement{}).
			Where("user_id = ? AND title_id = ?", uuid.UUID(userID), uuid.UUID(titleID)).
			Update("is_current", true).Error
	})
}

// convertAchievementsToEntities はGORMモデルリストからエンティティリストに変換する
func (r *titleAchievementRepository) convertAchievementsToEntities(gormAchievements []model.TitleAchievement) ([]*entity.TitleAchievement, error) {
	if len(gormAchievements) == 0 {
		return []*entity.TitleAchievement{}, nil
	}
	
	achievements := make([]*entity.TitleAchievement, len(gormAchievements))
	for i, gormAchievement := range gormAchievements {
		entityAchievement, err := gormAchievement.ToEntity()
		if err != nil {
			return nil, err
		}
		achievements[i] = entityAchievement
	}
	
	return achievements, nil
}