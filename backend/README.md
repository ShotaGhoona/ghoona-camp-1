# Ghoona Camp Backend

朝活コミュニティアプリ「Ghoona Camp」のバックエンドAPI

## アーキテクチャ

- **DDD (Domain-Driven Design)** × **オニオンアーキテクチャ**
- **境界付きコンテキスト**: user, attendance, goal, event, title, notification

## 技術スタック

- **言語**: Go 1.25+
- **フレームワーク**: Gin
- **ORM**: GORM
- **データベース**: Supabase (PostgreSQL)
- **認証**: Clerk JWT

## 開発環境

```bash
# 依存関係インストール
go mod download

# 環境変数設定
cp .env.example .env

# サーバー起動
go run cmd/api/main.go
```

## ディレクトリ構造

```
backend/
├── cmd/              # エントリーポイント
├── internal/         # 内部パッケージ
│   ├── domain/       # ドメイン層
│   ├── application/  # アプリケーション層
│   ├── infrastructure/ # インフラ層
│   └── interface/    # インターフェース層
└── ...
```