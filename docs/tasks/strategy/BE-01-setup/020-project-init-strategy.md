# BE-01-setup-02 プロジェクト初期化戦略

## 🎯 30秒キャッチアップ

**目標**: Ghoona Campバックエンドプロジェクトの初期化とオニオンアーキテクチャ構造作成  
**作業時間**: 2時間  
**対象**: Go modules設定、ディレクトリ構造作成、基本設定ファイル作成  
**アーキテクチャ**: eagle-ai準拠のオニオンアーキテクチャ + DDD

## 📋 実装内容

### 1. プロジェクト基本構造作成（30分）
- `backend/` ディレクトリ作成
- `go mod init` でGo modules初期化
- 基本的な依存関係追加（Gin、GORM等）

### 2. オニオンアーキテクチャ構造作成（60分）
- `cmd/` 階層（api, migrate, seed）
- `internal/` 4層構造作成
- 境界付きコンテキスト別ディレクトリ作成

### 3. 基本設定ファイル作成（30分）
- 環境変数設定（.env.example）
- 基本的なconfig構造
- main.goのスケルトン実装

## 🏗️ 作成するディレクトリ構造

```
backend/
├── cmd/
│   ├── api/
│   │   └── main.go              # APIサーバーエントリーポイント
│   ├── migrate/
│   │   └── main.go              # マイグレーション実行
│   └── seed/
│       └── main.go              # シードデータ投入
├── internal/
│   ├── domain/                  # ドメイン層（境界付きコンテキスト別）
│   │   ├── user/                # ユーザー管理コンテキスト
│   │   │   ├── entity/
│   │   │   ├── repository/
│   │   │   ├── value/
│   │   │   └── errors.go
│   │   ├── attendance/          # 出席管理コンテキスト
│   │   │   ├── entity/
│   │   │   ├── repository/
│   │   │   ├── value/
│   │   │   └── errors.go
│   │   ├── goal/                # 目標管理コンテキスト
│   │   ├── event/               # イベント管理コンテキスト
│   │   ├── title/               # 称号管理コンテキスト
│   │   ├── notification/        # 通知管理コンテキスト
│   │   └── common/              # 共通ドメイン
│   │       └── errors.go
│   ├── application/             # アプリケーション層
│   │   ├── dto/                 # コンテキスト別DTO
│   │   │   ├── user/
│   │   │   ├── attendance/
│   │   │   ├── goal/
│   │   │   ├── event/
│   │   │   ├── title/
│   │   │   └── notification/
│   │   ├── usecase/             # ユースケース実装
│   │   └── transaction/
│   │       └── manager.go
│   ├── infrastructure/          # インフラ層
│   │   ├── config/
│   │   │   ├── app_config.go    # アプリ設定
│   │   │   └── server.go        # サーバー設定
│   │   ├── database/
│   │   │   └── database.go      # DB接続
│   │   ├── gorm/
│   │   │   ├── model/           # GORMモデル
│   │   │   └── repository/      # リポジトリ実装
│   │   ├── jwt/
│   │   │   └── token_service.go # JWT認証
│   │   └── external/            # 外部API（Discord等）
│   ├── interface/               # インターフェース層
│   │   ├── controller/          # HTTPコントローラー
│   │   ├── middleware/          # ミドルウェア
│   │   │   ├── cors.go
│   │   │   ├── logger.go
│   │   │   ├── error.go
│   │   │   └── auth.go
│   │   └── router/              # ルーティング
│   │       └── router.go
│   └── di/                      # DI管理
│       └── container.go
├── go.mod
├── go.sum
├── .env.example                 # 環境変数テンプレート
├── .gitignore                   # Git除外設定
└── README.md                    # プロジェクト説明
```

## 📦 追加する依存関係

### Web Framework
- `github.com/gin-gonic/gin` - HTTPフレームワーク
- `github.com/gin-contrib/cors` - CORS対応

### データベース
- `gorm.io/gorm` - ORM
- `gorm.io/driver/postgres` - PostgreSQLドライバー

### 設定管理
- `github.com/joho/godotenv` - 環境変数読み込み

### JWT認証
- `github.com/golang-jwt/jwt/v5` - JWT処理

### ログ・ユーティリティ
- `github.com/sirupsen/logrus` - 構造化ログ
- `github.com/google/uuid` - UUID生成

## 🔧 基本設定ファイル

### .env.example
```env
# Database
DATABASE_URL=postgresql://postgres:password@localhost:5432/ghoona_camp_dev
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-anon-key

# JWT
JWT_SECRET=your-jwt-secret-key
JWT_EXPIRES_IN=24h

# Server
PORT=8080
GIN_MODE=debug

# Clerk
CLERK_SECRET_KEY=your-clerk-secret-key

# Discord
DISCORD_BOT_TOKEN=your-discord-bot-token
DISCORD_API_URL=https://discord.com/api/v10
```

### main.go スケルトン
```go
package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // 環境変数読み込み
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Ginエンジン初期化
    r := gin.Default()

    // ヘルスチェックエンドポイント
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
            "service": "ghoona-camp-backend",
        })
    })

    // サーバー起動
    port := getEnv("PORT", "8080")
    log.Printf("Starting server on port %s", port)
    r.Run(":" + port)
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

## ✅ 完了基準

- [ ] Go modules初期化完了
- [ ] eagle-ai準拠ディレクトリ構造作成完了
- [ ] 基本依存関係追加完了
- [ ] 環境変数テンプレート作成
- [ ] main.goスケルトン実装
- [ ] ヘルスチェックAPI動作確認
- [ ] .gitignore設定完了

## 🎯 次のステップ
**BE-02-arch-01**: オニオンアーキテクチャ基盤の構築（domain/application/infrastructure/interface層の基本構造実装）