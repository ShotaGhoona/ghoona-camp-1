# Backend ディレクトリガイド（1分で理解）

## 📁 全体のディレクトリ構造

```
backend/
├── cmd/              # 【実行可能ファイル】
├── internal/         # 【メインのビジネスロジック】
├── config/           # 【設定ファイル】
├── docs/             # 【ドキュメント】
├── db_data/          # 【データベースファイル】
├── go.mod            # 【Goモジュール定義】
├── go.sum            # 【依存関係のチェックサム】
├── .env              # 【環境変数】
├── Dockerfile        # 【Dockerイメージ定義】
└── docker-compose.yml # 【コンテナ構成】
```

## 🚀 cmd/ - 実行可能ファイル

```
cmd/
├── api/
│   └── main.go      # APIサーバー起動
├── migrate/
│   └── main.go      # DBマイグレーション実行
└── seed/
    └── main.go      # 初期データ投入
```

**役割:**
- 各種エントリーポイント
- `go run cmd/api/main.go` で起動
- 依存性注入（DI）の設定

## 📦 Go モジュール管理

### go.mod
```go
module github.com/ghoona_camp/backend

go 1.23

require (
    github.com/gin-gonic/gin v1.9.1      // Webフレームワーク
    gorm.io/gorm v1.25.5                 // ORM
    gorm.io/driver/mysql v1.5.2          // MySQLドライバー
    github.com/golang-jwt/jwt/v5 v5.0.0  // JWT認証
    github.com/joho/godotenv v1.5.1      // 環境変数
)
```

### go.sum
- 依存パッケージの整合性チェック用
- 自動生成されるので触らない
- `go mod tidy` で更新

## 🏗️ internal/ - メインロジック（詳細）

```
backend/internal/
├── domain/           # 【ビジネスの核心】
│   ├── entity/       # User, Attendance, Goal, Event, Title等のモデル
│   ├── repository/   # インターフェース定義のみ
│   └── errors.go     # ビジネスエラー
│
├── application/      # 【ビジネスの流れ】
│   ├── usecase/      # ログイン、出席記録、目標管理、イベント参加等
│   ├── dto/          # request/response構造体
│   └── transaction/  # トランザクション管理
│
├── infrastructure/   # 【外部との接続】
│   ├── database/     # DB接続設定（Supabase）
│   ├── gorm/         # GORM実装
│   │   ├── model/    # テーブル定義
│   │   └── repository/ # リポジトリ実装
│   ├── jwt/          # トークン生成
│   ├── clerk/        # Clerk認証連携
│   └── external/     # 外部API（通知サービス等）
│
└── interface/        # 【HTTPの入り口】
    ├── handler/      # コントローラー
    ├── middleware/   # 認証、CORS、ログ
    └── router/       # ルーティング設定
```

## 🎯 Gin特有のファイル配置

### 1. ルーター設定
`internal/interface/router/router.go`
```go
func SetupRouter(handlers...) *gin.Engine {
    r := gin.Default()
    r.Use(middleware.CORS())
    r.Use(middleware.Logger())
    
    // ルート定義
    v1 := r.Group("/api/v1")
    {
        v1.POST("/login", userHandler.Login)
        v1.GET("/users", middleware.Auth(), userHandler.GetUsers)
    }
    return r
}
```

### 2. メイン起動ファイル
`cmd/api/main.go`
```go
func main() {
    // 環境変数読み込み
    godotenv.Load()
    
    // DB接続
    db := database.NewDB()
    
    // 依存性注入
    userRepo := repository.NewUserRepository(db)
    userUseCase := usecase.NewUserUseCase(userRepo)
    userHandler := handler.NewUserHandler(userUseCase)
    
    // サーバー起動
    r := router.SetupRouter(userHandler)
    r.Run(":8080")
}
```

## 💡 クイックリファレンス

| 作業内容 | ファイル/コマンド |
|---------|----------------|
| サーバー起動 | `go run cmd/api/main.go` |
| DB初期化 | `go run cmd/migrate/main.go` |
| パッケージ追加 | `go get github.com/xxx/yyy` |
| 依存関係整理 | `go mod tidy` |
| 新規API追加 | 1. handler作成 2. router追加 |
| 環境変数設定 | `.env` ファイル編集 |

## 📝 実装手順（新機能追加）

1. **entity定義** → `domain/entity/新機能.go`（例：attendance.go, goal.go, event.go, title.go）
2. **リポジトリI/F** → `domain/repository/新機能_repository.go`
3. **DBモデル** → `infrastructure/gorm/model/新機能.go`
4. **リポジトリ実装** → `infrastructure/gorm/repository/新機能_repository.go`
5. **ユースケース** → `application/usecase/新機能_usecase.go`
6. **ハンドラー** → `interface/handler/新機能_handler.go`
7. **ルート追加** → `interface/router/router.go`
8. **DI設定** → `cmd/api/main.go`

## ⚡ よく使うコマンド

```bash
# 開発サーバー起動
docker-compose up

# Goモジュール管理
go mod init                  # 初期化
go get パッケージ名           # パッケージ追加
go mod tidy                  # 不要な依存削除

# ビルド・実行
go build cmd/api/main.go     # ビルド
go run cmd/api/main.go       # 実行
go test ./...                # テスト実行
```