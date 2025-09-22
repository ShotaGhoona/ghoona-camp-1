package gorm

import (
	"context"
	"gorm.io/gorm"
	"ghoona-camp-backend/internal/application/transaction"
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