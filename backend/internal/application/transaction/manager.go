package transaction

import (
	"context"
	"gorm.io/gorm"
)

// Manager はトランザクション管理のインターフェース
type Manager interface {
	ExecuteInTx(ctx context.Context, fn func(context.Context) error) error
}

// manager はGORMを使用したトランザクション管理実装
type manager struct {
	db *gorm.DB
}

// NewManager は新しいトランザクションマネージャーを作成
func NewManager(db *gorm.DB) Manager {
	return &manager{db: db}
}

// ExecuteInTx はトランザクション内で関数を実行
func (m *manager) ExecuteInTx(ctx context.Context, fn func(context.Context) error) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// コンテキストにトランザクションを設定
		txCtx := context.WithValue(ctx, "tx", tx)
		return fn(txCtx)
	})
}

// GetDB はコンテキストからDBまたはトランザクションを取得
func GetDB(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx // トランザクション中
	}
	return defaultDB // 通常のDB接続
}