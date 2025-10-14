package title

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/title/entity"
	"ghoona-camp-backend/internal/domain/title/repository"
	"ghoona-camp-backend/internal/domain/title/vo"
	titleModel "ghoona-camp-backend/internal/infrastructure/gorm/model/title"
)

type titleRepository struct {
	db *gorm.DB
}

// NewTitleRepository コンストラクタ
func NewTitleRepository(db *gorm.DB) repository.TitleRepository {
	return &titleRepository{db: db}
}

// Create タイトルを作成する
func (r *titleRepository) Create(ctx context.Context, title *entity.Title) error {
	gormTitle := r.toGORMTitle(title)
	return r.db.WithContext(ctx).Create(gormTitle).Error
}

// GetByID IDでタイトルを取得する
func (r *titleRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Title, error) {
	var gormTitle titleModel.Title
	err := r.db.WithContext(ctx).First(&gormTitle, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMTitle(&gormTitle)
}

// GetByLevel レベルで称号を取得する
func (r *titleRepository) GetByLevel(ctx context.Context, level vo.TitleLevel) (*entity.Title, error) {
	var gormTitle titleModel.Title
	err := r.db.WithContext(ctx).Where("level = ?", string(level)).First(&gormTitle).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMTitle(&gormTitle)
}

// GetAll すべての称号を取得する
func (r *titleRepository) GetAll(ctx context.Context) ([]*entity.Title, error) {
	var gormTitles []titleModel.Title
	err := r.db.WithContext(ctx).Find(&gormTitles).Error
	if err != nil {
		return nil, err
	}

	titles := make([]*entity.Title, len(gormTitles))
	for i, gormTitle := range gormTitles {
		title, err := r.fromGORMTitle(&gormTitle)
		if err != nil {
			return nil, err
		}
		titles[i] = title
	}
	return titles, nil
}

// GetActive 有効な称号一覧を取得する
func (r *titleRepository) GetActive(ctx context.Context) ([]*entity.Title, error) {
	var gormTitles []titleModel.Title
	err := r.db.WithContext(ctx).Where("is_active = ?", true).Find(&gormTitles).Error
	if err != nil {
		return nil, err
	}

	titles := make([]*entity.Title, len(gormTitles))
	for i, gormTitle := range gormTitles {
		title, err := r.fromGORMTitle(&gormTitle)
		if err != nil {
			return nil, err
		}
		titles[i] = title
	}
	return titles, nil
}

// Update タイトルを更新する
func (r *titleRepository) Update(ctx context.Context, title *entity.Title) error {
	gormTitle := r.toGORMTitle(title)
	return r.db.WithContext(ctx).Save(gormTitle).Error
}

// Delete タイトルを削除する
func (r *titleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&titleModel.Title{}, "id = ?", id).Error
}


// 型変換: Domain Entity → GORM Model
func (r *titleRepository) toGORMTitle(title *entity.Title) *titleModel.Title {
	return &titleModel.Title{
		ID:           title.ID(),
		Name:         title.Name(),
		Slug:         title.Slug(),
		Description:  title.Description(),
		Rarity:       title.Rarity(),
		Category:     title.Category(),
		ImageURL:     title.ImageURL(),
		Requirements: title.Requirements(),
		Conditions:   title.Conditions(),
		IsHidden:     title.IsHidden(),
		CreatedAt:    title.CreatedAt(),
		UpdatedAt:    title.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *titleRepository) fromGORMTitle(gormTitle *titleModel.Title) (*entity.Title, error) {
	// TODO: entity.NewTitle の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}