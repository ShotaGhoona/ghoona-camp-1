package gorm

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/domain/common"
)

// BaseRepository は基底リポジトリ
type BaseRepository struct {
	db *gorm.DB
}

// NewBaseRepository は基底リポジトリを作成
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

// GetDB はコンテキストに応じたDB接続を取得
// トランザクション内であればtx、そうでなければ通常のdbを返す
func (r *BaseRepository) GetDB(ctx context.Context) *gorm.DB {
	return transaction.GetDB(ctx, r.db)
}

// DB は基本DB接続を取得
func (r *BaseRepository) DB() *gorm.DB {
	return r.db
}

// WithTx はトランザクション付きリポジトリを作成
func (r *BaseRepository) WithTx(tx *gorm.DB) *BaseRepository {
	return &BaseRepository{db: tx}
}

// HealthCheck はデータベース接続状態をチェック
func (r *BaseRepository) HealthCheck(ctx context.Context) error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// Ping with context
	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	return nil
}

// GetConnectionStats はデータベース接続統計を取得
func (r *BaseRepository) GetConnectionStats() (map[string]interface{}, error) {
	sqlDB, err := r.db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	stats := sqlDB.Stats()
	return map[string]interface{}{
		"open_connections":     stats.OpenConnections,
		"in_use":              stats.InUse,
		"idle":                stats.Idle,
		"wait_count":          stats.WaitCount,
		"wait_duration":       stats.WaitDuration.String(),
		"max_idle_closed":     stats.MaxIdleClosed,
		"max_idle_time_closed": stats.MaxIdleTimeClosed,
		"max_lifetime_closed": stats.MaxLifetimeClosed,
	}, nil
}

// ExecuteInTransaction はトランザクション内で関数を実行
func (r *BaseRepository) ExecuteInTransaction(ctx context.Context, fn func(*gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

// SetLogLevel はログレベルを動的に変更
func (r *BaseRepository) SetLogLevel(level logger.LogLevel) {
	r.db.Logger = r.db.Logger.LogMode(level)
}

// CreateBaseEntity は基底エンティティの作成時共通処理
func (r *BaseRepository) CreateBaseEntity(entity interface{}) error {
	now := time.Now()
	
	// BaseEntityが実装されている場合の処理
	if baseEntity, ok := entity.(interface {
		SetCreatedAt(time.Time)
		SetUpdatedAt(time.Time)
	}); ok {
		baseEntity.SetCreatedAt(now)
		baseEntity.SetUpdatedAt(now)
	}
	
	return nil
}

// UpdateBaseEntity は基底エンティティの更新時共通処理
func (r *BaseRepository) UpdateBaseEntity(entity interface{}) error {
	now := time.Now()
	
	// BaseEntityが実装されている場合の処理
	if baseEntity, ok := entity.(interface {
		SetUpdatedAt(time.Time)
	}); ok {
		baseEntity.SetUpdatedAt(now)
	}
	
	return nil
}

// Paginate はページネーション処理を実行
func (r *BaseRepository) Paginate(ctx context.Context, query *gorm.DB, page, limit int, result interface{}) (*common.Pagination, error) {
	var total int64
	
	// 総件数を取得
	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count records: %w", err)
	}
	
	// ページネーション計算
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	
	offset := (page - 1) * limit
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	
	// データ取得
	if err := query.Offset(offset).Limit(limit).Find(result).Error; err != nil {
		return nil, fmt.Errorf("failed to find records: %w", err)
	}
	
	pagination := &common.Pagination{
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalCount:  int(total),
		HasNext:     page < totalPages,
		HasPrev:     page > 1,
	}
	
	return pagination, nil
}

// IsRecordNotFound はGORMのErrRecordNotFoundをチェック
func (r *BaseRepository) IsRecordNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}

// HandleError は共通エラーハンドリング
func (r *BaseRepository) HandleError(err error, operation string) error {
	if err == nil {
		return nil
	}
	
	if r.IsRecordNotFound(err) {
		return common.ErrNotFound
	}
	
	return fmt.Errorf("%s failed: %w", operation, err)
}