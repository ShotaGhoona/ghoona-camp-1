# Database Scripts

データベース関連のスクリプト群です。

## スクリプト一覧

### migrate.sh
マイグレーション管理スクリプト

**使用方法:**
```bash
# 全マイグレーション実行
./db/scripts/migrate.sh up

# マイグレーション状況確認
./db/scripts/migrate.sh status

# シードデータ投入
./db/scripts/migrate.sh seed development
./db/scripts/migrate.sh seed testing
./db/scripts/migrate.sh seed production

# データベースリセット
./db/scripts/migrate.sh reset
```

### db-setup.sh
データベース初期セットアップスクリプト

**使用方法:**
```bash
# 開発環境セットアップ
./db/scripts/db-setup.sh development

# テスト環境セットアップ
./db/scripts/db-setup.sh testing

# 本番環境セットアップ
./db/scripts/db-setup.sh production

# リセットしてセットアップ
./db/scripts/db-setup.sh development --reset

# マイグレーションのみ
./db/scripts/db-setup.sh development --migrations-only

# シードデータのみ
./db/scripts/db-setup.sh development --seeds-only
```

## 環境変数

すべてのスクリプトで以下の環境変数が必要です：

```bash
export DATABASE_URL="postgresql://user:password@host:port/database"
```

## ディレクトリ構造

```
db/
├── migrations/           # マイグレーションファイル
│   ├── 001_create_users.sql
│   ├── 002_create_user_metadata.sql
│   ├── 003_create_user_social_links.sql
│   ├── 004_create_user_rivals.sql
│   ├── 005_create_events.sql
│   ├── 006_create_attendance.sql
│   └── 007_create_titles.sql
├── seeds/               # シードデータ
│   ├── development.sql  # 開発環境用
│   ├── testing.sql      # テスト環境用
│   └── production.sql   # 本番環境用
└── README.md

└── scripts/            # データベーススクリプト
    ├── migrate.sh       # マイグレーション管理
    ├── db-setup.sh     # 初期セットアップ
    └── README.md       # このファイル
```

## セットアップ手順

### 開発環境
```bash
# 1. 環境変数設定
export DATABASE_URL="postgresql://postgres.xxx:password@host:5432/postgres"

# 2. 開発環境セットアップ
./db/scripts/db-setup.sh development
```

### テスト環境
```bash
# 1. 環境変数設定
export DATABASE_URL="postgresql://test_user:password@host:5432/test_db"

# 2. テスト環境セットアップ
./db/scripts/db-setup.sh testing
```

### 本番環境
```bash
# 1. 環境変数設定
export DATABASE_URL="postgresql://prod_user:password@host:5432/prod_db"

# 2. 本番環境セットアップ（称号マスターデータのみ）
./db/scripts/db-setup.sh production
```

## トラブルシューティング

### 権限エラー
スクリプトに実行権限がない場合：
```bash
chmod +x db/scripts/*.sh
```

### 接続エラー
DATABASE_URLの確認：
```bash
echo $DATABASE_URL
psql $DATABASE_URL -c "SELECT 1;"
```

### マイグレーション状況確認
```bash
./db/scripts/migrate.sh status
```