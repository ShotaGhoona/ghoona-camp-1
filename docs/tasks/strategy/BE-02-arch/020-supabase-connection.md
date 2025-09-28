# BE-02-arch-02 Supabase接続設定戦略書

## 📁 影響を及ぼすファイル・ディレクトリ（修正・新規作成対象）

```
backend/
├── internal/
│   ├── infrastructure/
│   │   ├── config/
│   │   │   └── app_config.go                         # ✅ 既に実装済み（DB設定含む）
│   │   ├── database/
│   │   │   ├── connection.go                         # ✅ 既に実装済み（Supabase接続完成）
│   │   │   ├── health.go                             # ✅ 既に実装済み（ヘルスチェック完成）
│   │   │   └── migrations/
│   │   │       └── migrate.go                        # ✅ 既に実装済み（基盤のみ）
│   │   └── gorm/
│   │       ├── base_repository.go                    # ✏️ 軽微修正（トランザクション対応強化）
│   │       └── migration/                            # 📁 新規ディレクトリ
│   │           ├── 001_initial_tables.sql            # 🆕 初期テーブル作成SQL
│   │           └── seed_data.sql                     # 🆕 開発用初期データ
│   ├── interface/
│   │   ├── controller/
│   │   │   └── system_controller.go                  # 🆕 DB接続テスト用コントローラー
│   │   └── router/
│   │       └── router.go                             # ✏️ DB接続テスト用ルート追加
│   └── di/
│       └── container.go                              # ✏️ システムコントローラーのDI追加
├── cmd/
│   ├── api/main.go                                   # ✅ 既に実装済み（DB接続統合完了）
│   ├── migrate/main.go                               # ✏️ 実際のマイグレーション実行実装
│   └── seed/main.go                                  # ✏️ 実際のシード実行実装
└── .env.example                                      # ✏️ 実際のSupabase接続例に更新

frontend/
└── src/
    └── app/
        └── test/
            └── connection/
                └── page.tsx                          # 🆕 DB接続テストページ作成
```

## 🎯 タスク概要

**BE-02-arch-02: Supabase接続設定**
- **目的**: 既存の基盤を活用し、実際のSupabaseとの接続を確立
- **工数**: 2時間（大幅短縮、既に基盤が実装済みのため）
- **優先度**: 🔴 高（クリティカル）

## 📋 実装内容（チェックリスト準拠スコープ）

**BE-02-arch-02: データベース接続設定**
- **チェックリスト内容**: Supabase接続、GORM設定、マイグレーション基盤の実装
- **スコープ外**: システムAPI、フロントエンドテストページ（別タスク）

### ✅ 既に実装済みの項目

1. **設定管理**: `app_config.go` - DatabaseConfig構造体完全実装済み
2. **DB接続**: `connection.go` - Supabase PostgreSQL接続完全実装済み
3. **ヘルスチェック**: `health.go` - DB状態監視機能完全実装済み
4. **マイグレーション基盤**: `migrate.go` - 基本的なマイグレーション管理実装済み
5. **main.go統合**: DB接続のエントリーポイント統合完了
6. **環境変数設定**: `.env.example`にDB設定テンプレート実装済み

### 🔄 実装が必要な項目（チェックリスト準拠）

### Phase 1: マイグレーション実装 (1時間)

#### 1.1 SQLファイル作成（チェックリスト範囲内）
- `backend/internal/infrastructure/gorm/migration/001_initial_tables.sql`: 基本テーブル構造
- `backend/internal/infrastructure/gorm/migration/seed_data.sql`: 開発用初期データ

#### 1.2 マイグレーション実行（チェックリスト範囲内）
- `cmd/migrate/main.go`: 実際のマイグレーション実行コマンド
- `cmd/seed/main.go`: 実際のシード実行コマンド
- マイグレーション基盤の`migrate.go`とSQLファイルの統合

### Phase 2: GORM設定強化 (1時間)

#### 2.1 GORM基盤強化（チェックリスト範囲内）
- `base_repository.go`: トランザクション対応とエラーハンドリング強化
- GORM設定の最適化とパフォーマンス調整

#### 2.2 設定ファイル更新（チェックリスト範囲内）
- `.env.example`: 実際のSupabase接続例に更新

### ❌ スコープ外項目（今回実装しない）

#### システムAPI関連
- `system_controller.go`: DB接続テスト用API ← **BE-11-system-01で実装**
- `/api/v1/system/health`: DB状態確認API ← **BE-11-system-01で実装**
- システムAPIルーティング ← **BE-11-system-01で実装**

#### フロントエンド関連
- `frontend/src/app/test/connection/page.tsx` ← **フロントエンドタスクで実装**
- DB接続状況表示機能 ← **フロントエンドタスクで実装**

#### 理由
- チェックリストの**BE-02-arch-02**は「Supabase接続、GORM設定、マイグレーション基盤」に限定
- システムAPIは**BE-11-system-01**（ヘルスチェック・監視API）で実装予定
- フロントエンドは別のタスク体系で管理

## ✅ 完了条件（チェックリスト準拠）

1. **Supabase接続確立**: PostgreSQLへの安定した接続 ✅（既に実装済み）
2. **GORM設定完了**: ORMとSupabaseの統合完了 ✅（既に実装済み）
3. **マイグレーション基盤実装**: SQLファイル管理とマイグレーション実行
4. **マイグレーション動作確認**: `go run cmd/migrate/main.go`でテーブル作成成功
5. **シード動作確認**: `go run cmd/seed/main.go`で初期データ投入成功
6. **Docker環境対応**: コンテナ環境でのDB接続確認

## 🎯 技術仕様

### データベース仕様
- **DB**: Supabase PostgreSQL
- **ORM**: GORM v1.31.0
- **ドライバー**: postgres driver
- **接続プール**: 最大10接続
- **タイムアウト**: 30秒

### マイグレーション仕様
- SQLファイル管理: `internal/infrastructure/gorm/migration/`
- バージョン管理: `schema_migrations`テーブル
- 実行コマンド: `go run cmd/migrate/main.go`、`go run cmd/seed/main.go`

### 環境変数
```env
# Supabase Database
DATABASE_URL=postgresql://[user]:[password]@[host]:[port]/[database]
DB_HOST=db.your-project.supabase.co
DB_PORT=5432
DB_NAME=postgres
DB_USER=postgres
DB_PASSWORD=your-password
DB_SSL_MODE=require
```

## 🔧 動作確認方法

### 1. Docker環境起動
```bash
docker-compose up --build
```

### 2. マイグレーション実行
```bash
# コンテナ内で実行
docker exec -it ghoona_backend go run cmd/migrate/main.go
```

### 3. DB接続確認
```bash
# アプリケーション起動でDB接続確認
# 起動ログにDB接続成功メッセージが表示される
docker-compose logs backend
```

### 4. TODO: システムAPI確認（BE-11-system-01で実装）
```bash
# 将来的にヘルスチェックAPI実装予定
# curl http://localhost:8080/api/v1/system/health
```

## 🚨 注意事項

1. **Supabase設定**: 実際のSupabaseプロジェクト作成が必要
2. **環境変数**: 本物のDB接続情報設定必須
3. **セキュリティ**: 接続情報の適切な管理
4. **接続プール**: 適切なプール設定でメモリ効率化

## 📈 次のステップへの準備

このタスク完了により以下が可能になります：

- **BE-02-arch-03**: Clerk認証でのユーザー情報DB保存
- **BE-03-user-01**: ユーザーエンティティのDB永続化
- **BE-04-***: 各ドメインエンティティのDB操作
- **BE-10-batch-***: バッチ処理でのDB操作

**重要**: このタスクはすべての後続開発の基盤となる重要な実装です。