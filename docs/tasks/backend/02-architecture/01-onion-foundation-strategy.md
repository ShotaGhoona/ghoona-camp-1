# BE-02-arch-01: オニオンアーキテクチャ基盤構築戦略書

## 📋 タスク概要
**タスクID**: BE-02-arch-01  
**タスク名**: オニオンアーキテクチャ基盤の構築  
**内容**: domain/application/infrastructure/interface層の基本構造実装  
**工数見積**: 8時間  
**優先度**: 🔴 高（最優先）  

## 🎯 目標
既存の6つの境界付きコンテキスト（user, attendance, goal, event, title, notification）について、オニオンアーキテクチャの4層構造を実装し、基盤となる共通コンポーネントを構築する。

**重要**: このタスクはアーキテクチャの**基盤構築**のみに集中。具体的なビジネスロジックは後続タスク（BE-03-user-01以降）で実装。今回は骨組み・インターフェース定義・共通基盤の作成に専念。

## 📖 前提条件・参考資料
- BE-01-setup-01,02,03 完了済み
- eagle-ai プロジェクト構造を参考
- Ghoona Camp 要件（product vision, tech spec, DB design, API design）
- DDD × オニオンアーキテクチャ原則

## 🏗️ 実装範囲

### 1. ドメイン層（Domain Layer）基盤
**場所**: `internal/domain/`

#### 1.1 共通基盤
- `internal/domain/common/errors.go` - 共通ドメインエラー定義
- `internal/domain/common/types.go` - 共通型定義
- `internal/domain/common/specification.go` - 仕様パターン基盤

#### 1.2 コンテキスト別構造（6ドメイン）
各コンテキスト（user, attendance, goal, event, title, notification）で以下の**空の構造**を実装：

```
internal/domain/{context}/
├── entity/           # エンティティ（//TODO実装）
├── repository/       # リポジトリインターフェース（//TODO実装）
├── service/          # ドメインサービス（//TODO実装）
├── value/           # 値オブジェクト（//TODO実装）
└── errors.go        # コンテキスト固有エラー（//TODO実装）
```

#### 1.3 今回実装範囲（基盤のみ）
- **ディレクトリ構造作成**
- **空のファイル作成**（//TODOコメント付き）
- **基本的なインターフェース定義**のみ
- 具体的なエンティティ実装は**BE-03-user-01以降**で実施

### 2. アプリケーション層（Application Layer）基盤
**場所**: `internal/application/`

#### 2.1 共通基盤
- `internal/application/transaction/manager.go` - トランザクション管理
- `internal/application/common/response.go` - 共通レスポンス形式
- `internal/application/common/errors.go` - アプリケーションエラー変換

#### 2.2 コンテキスト別構造
```
internal/application/
├── dto/
│   └── {context}/    # コンテキスト別DTO
├── usecase/
│   └── {context}/    # ユースケース実装
└── transaction/      # トランザクション管理
```

### 3. インフラストラクチャ層（Infrastructure Layer）基盤
**場所**: `internal/infrastructure/`

#### 3.1 データベース基盤
- `internal/infrastructure/database/connection.go` - Supabase接続
- `internal/infrastructure/gorm/base_repository.go` - 基底リポジトリ
- `internal/infrastructure/gorm/transaction.go` - トランザクション実装

#### 3.2 外部サービス基盤
- `internal/infrastructure/clerk/auth_service.go` - Clerk認証連携
- `internal/infrastructure/discord/webhook.go` - Discord連携
- `internal/infrastructure/config/app_config.go` - 設定管理

#### 3.3 GORM構造
```
internal/infrastructure/gorm/
├── model/           # データベースモデル
├── repository/      # リポジトリ実装
└── migration/       # マイグレーション
```

### 4. インターフェース層（Interface Layer）基盤
**場所**: `internal/interface/`

#### 4.1 HTTP基盤
- `internal/interface/middleware/auth.go` - Clerk認証ミドルウェア
- `internal/interface/middleware/cors.go` - CORS設定
- `internal/interface/middleware/error.go` - エラーハンドリング
- `internal/interface/middleware/logger.go` - ログミドルウェア

#### 4.2 ルーティング構造
```
internal/interface/
├── controller/      # HTTPコントローラー
├── middleware/      # ミドルウェア
├── router/          # ルーティング設定
└── dto/            # HTTPリクエスト/レスポンスDTO
```

## 🎮 実装計画

### Phase 1: 共通基盤・ディレクトリ構造（3時間）
1. 4層×6コンテキストのディレクトリ構造作成
2. 空ファイル作成（//TODOコメント付き）
3. 共通エラー・型定義・トランザクション管理基盤
4. 設定管理システム

### Phase 2: インターフェース定義（2時間）
1. 基本的なリポジトリインターフェースの雛形
2. 基本的なユースケースインターフェースの雛形  
3. 基本的なエンティティインターフェースの雛形
4. DTO構造の雛形

### Phase 3: インフラ基盤（2時間）
1. Supabase接続設定（実際の接続は次タスク）
2. GORM基底リポジトリ
3. 設定ファイル読み込み
4. ミドルウェア基盤

### Phase 4: HTTPインターフェース基盤（1時間）
1. ルーティング基盤
2. エラーハンドリングミドルウェア
3. ヘルスチェック強化

**重要**: 実際のビジネスロジック・DB接続・認証実装は次タスク以降！

## 🔧 技術実装詳細

### 1. Clerk認証統合
```go
// internal/infrastructure/clerk/auth_service.go
type ClerkAuthService struct {
    secretKey string
}

func (s *ClerkAuthService) VerifyToken(token string) (*ClerkUser, error) {
    // Clerk JWT検証ロジック
}
```

### 2. トランザクション管理
```go
// internal/application/transaction/manager.go
type Manager interface {
    ExecuteInTx(ctx context.Context, fn func(context.Context) error) error
}

// GORM実装
func (m *gormManager) ExecuteInTx(ctx context.Context, fn func(context.Context) error) error {
    return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        txCtx := context.WithValue(ctx, "tx", tx)
        return fn(txCtx)
    })
}
```

### 3. エラーハンドリング戦略
```go
// 4段階エラー変換
Domain Error → Application Error → Infrastructure Error → HTTP Error

// 例: ユーザー不存在
domain.ErrUserNotFound → app.ErrNotFound → http.StatusNotFound
```

### 4. 設定管理
```go
// internal/infrastructure/config/app_config.go
type Config struct {
    Server   ServerConfig
    Database DatabaseConfig  
    Clerk    ClerkConfig
    Discord  DiscordConfig
}

func LoadConfig() (*Config, error) {
    // .env + 環境変数読み込み
}
```

## 📊 データベース連携

### Supabase接続
```go
// internal/infrastructure/database/connection.go
func NewSupabaseDB(config *config.DatabaseConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require",
        config.User, config.Password, config.Host, config.Port, config.Database)
    
    return gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
}
```

### 基底リポジトリ
```go
// internal/infrastructure/gorm/base_repository.go
type BaseRepository struct {
    db *gorm.DB
}

func (r *BaseRepository) GetDB(ctx context.Context) *gorm.DB {
    if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
        return tx // トランザクション中
    }
    return r.db // 通常のDB接続
}
```

## 🧪 テスト戦略
1. **ドメイン層**: 単体テスト（エンティティ・値オブジェクト）
2. **アプリケーション層**: モックを使用した単体テスト
3. **インフラ層**: 統合テスト（実DB使用）
4. **インターフェース層**: APIテスト

## 📝 成果物
1. 6つのドメインコンテキストの基本構造
2. 4層アーキテクチャの実装基盤
3. Clerk認証統合
4. Supabaseデータベース連携
5. 共通エラーハンドリングシステム
6. トランザクション管理システム

## ✅ 完了条件
- [ ] 4層×6コンテキストの完全なディレクトリ構造作成
- [ ] 各層の基本インターフェース定義（//TODO付き空実装）
- [ ] 共通基盤コンポーネント実装
- [ ] 設定管理システム実装
- [ ] エラーハンドリング基盤実装
- [ ] ルーティング基盤実装
- [ ] ヘルスチェックAPI拡張
- [ ] Docker環境での基本動作確認

**注意**: 実際のDB接続・認証・ビジネスロジックは含まない（次タスクで実装）

## 🎯 次のステップ
**BE-02-arch-02**: データベース接続設定（Supabase接続、GORM設定、マイグレーション基盤の実装）

## 🚨 リスク・考慮事項
1. **Clerk統合**: トークン検証ロジックの複雑性
2. **Supabase接続**: 認証情報・接続プール設定
3. **トランザクション**: 複数テーブルにまたがる処理
4. **エラーハンドリング**: 多層アーキテクチャでのエラー伝播

この戦略に基づいて、Ghoona Campのオニオンアーキテクチャ基盤を構築します。