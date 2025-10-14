package goal

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/goal/entity"
	"ghoona-camp-backend/internal/domain/goal/repository"
	goalModel "ghoona-camp-backend/internal/infrastructure/gorm/model/goal"
)

type goalRepository struct {
	db *gorm.DB
}

// NewGoalRepository コンストラクタ
func NewGoalRepository(db *gorm.DB) repository.GoalRepository {
	return &goalRepository{db: db}
}

// Create 目標を作成する
func (r *goalRepository) Create(ctx context.Context, goal *entity.Goal) error {
	gormGoal := r.toGORMGoal(goal)
	return r.db.WithContext(ctx).Create(gormGoal).Error
}

// GetByID IDで目標を取得する
func (r *goalRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Goal, error) {
	var gormGoal goalModel.Goal
	err := r.db.WithContext(ctx).First(&gormGoal, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.fromGORMGoal(&gormGoal)
}

// GetByUserID ユーザーIDで目標一覧を取得する
func (r *goalRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.Goal, error) {
	var gormGoals []goalModel.Goal
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&gormGoals).Error
	if err != nil {
		return nil, err
	}

	goals := make([]*entity.Goal, len(gormGoals))
	for i, gormGoal := range gormGoals {
		goal, err := r.fromGORMGoal(&gormGoal)
		if err != nil {
			return nil, err
		}
		goals[i] = goal
	}
	return goals, nil
}

// GetByUserIDAndStatus ユーザーIDとステータスで目標を取得する
func (r *goalRepository) GetByUserIDAndStatus(ctx context.Context, userID uuid.UUID, status string, limit, offset int) ([]*entity.Goal, error) {
	var gormGoals []goalModel.Goal
	err := r.db.WithContext(ctx).Where("user_id = ? AND status = ?", userID, status).Limit(limit).Offset(offset).Find(&gormGoals).Error
	if err != nil {
		return nil, err
	}

	goals := make([]*entity.Goal, len(gormGoals))
	for i, gormGoal := range gormGoals {
		goal, err := r.fromGORMGoal(&gormGoal)
		if err != nil {
			return nil, err
		}
		goals[i] = goal
	}
	return goals, nil
}

// Update 目標を更新する
func (r *goalRepository) Update(ctx context.Context, goal *entity.Goal) error {
	gormGoal := r.toGORMGoal(goal)
	return r.db.WithContext(ctx).Save(gormGoal).Error
}

// Delete 目標を削除する
func (r *goalRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&goalModel.Goal{}, "id = ?", id).Error
}

// List 目標一覧を取得する
func (r *goalRepository) List(ctx context.Context, limit, offset int) ([]*entity.Goal, error) {
	var gormGoals []goalModel.Goal
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&gormGoals).Error
	if err != nil {
		return nil, err
	}

	goals := make([]*entity.Goal, len(gormGoals))
	for i, gormGoal := range gormGoals {
		goal, err := r.fromGORMGoal(&gormGoal)
		if err != nil {
			return nil, err
		}
		goals[i] = goal
	}
	return goals, nil
}

// Search 目標を検索する
func (r *goalRepository) Search(ctx context.Context, query string, limit, offset int) ([]*entity.Goal, error) {
	// TODO: 検索仕様確定後実装（検索対象フィールド：title? description?）
	return nil, errors.New("TODO: 要件確定後実装")
}

// 型変換: Domain Entity → GORM Model
func (r *goalRepository) toGORMGoal(goal *entity.Goal) *goalModel.Goal {
	return &goalModel.Goal{
		ID:          goal.ID(),
		UserID:      goal.UserID(),
		Title:       goal.Title(),
		Description: goal.Description(),
		Status:      goal.Status(),
		Priority:    goal.Priority(),
		Category:    goal.Category(),
		TargetDate:  goal.TargetDate(),
		CreatedAt:   goal.CreatedAt(),
		UpdatedAt:   goal.UpdatedAt(),
	}
}

// 型変換: GORM Model → Domain Entity
func (r *goalRepository) fromGORMGoal(gormGoal *goalModel.Goal) (*entity.Goal, error) {
	// TODO: entity.NewGoal の引数確認後実装
	return nil, errors.New("TODO: entity constructor確認後実装")
}