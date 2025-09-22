# BE-02-arch-02 戦略書 - データベース接続設定

**タスクID**: BE-02-arch-02  
**タスク名**: データベース接続設定  
**担当**: Backend Team  
**工数見積**: 4時間  
**優先度**: 🔴 高  

## 📋 タスク概要

BE-02-arch-01で構築したオニオンアーキテクチャ基盤に、実際のデータベース接続機能を実装します。Supabase PostgreSQLへの接続、GORM設定、マイグレーション基盤を構築し、後続のドメイン実装（BE-03-*）の準備を完了させます。

## 🎯 完了条件

### ✅ 必須条件
1. **Supabase接続**: 実際のSupabase PostgreSQLへの接続確立
2. **GORM設定**: 接続プール、ログ設定、パフォーマンス最適化
3. **マイグレーション基盤**: データベーススキーマ管理システム
4. **ヘルスチェック**: データベース接続状態の監視機能
5. **エラーハンドリング**: 接続エラーの適切な処理

### 🔧 技術要件
- **データベース**: Supabase PostgreSQL
- **ORM**: GORM v1.25+
- **接続方式**: 環境変数ベースの設定
- **マイグレーション**: 自動実行可能な仕組み
- **ログ**: 構造化ログでクエリ監視

## 🏗️ 実装戦略

### Phase 1: 現状分析（15分）
1. **既存コード確認**:
   - `internal/infrastructure/database/connection.go` の現状
   - 既存のConfig構造体とSupabase設定
   - .envファイルの設定状況

2. **要件確認**:
   - docs/requirement/11-db.md の詳細要件
   - データベーススキーマ設計
   - パフォーマンス要求

### Phase 2: データベース接続実装（90分）
1. **Config拡張**:
   ```go
   type Config struct {
       Port               string
       ClerkSecretKey     string
       ClerkWebhookSecret string
       Env                string
       
       // 新規追加
       Database struct {
           URL             string
           MaxOpenConns    int
           MaxIdleConns    int
           ConnMaxLifetime time.Duration
           LogLevel        string
       }
   }
   ```

2. **GORM接続実装**:
   - Supabase PostgreSQL接続文字列解析
   - 接続プール最適化
   - クエリログ設定
   - 接続テスト機能

3. **エラーハンドリング強化**:
   - 接続失敗時の適切なエラー
   - リトライ機能
   - タイムアウト設定

### Phase 3: マイグレーション基盤（60分）
1. **マイグレーション構造**:
   ```
   internal/infrastructure/database/
   ├── connection.go       # 接続管理
   ├── migrations/         # マイグレーションファイル
   │   ├── migrate.go     # マイグレーション実行
   │   └── 001_initial.sql # 初期スキーマ
   └── health.go          # ヘルスチェック
   ```

2. **機能実装**:
   - 自動マイグレーション実行
   - マイグレーション状態管理
   - ロールバック機能準備

### Phase 4: テスト・検証（30分）
1. **接続テスト**:
   - ローカル環境での接続確認
   - Supabase本番環境への接続テスト
   - 負荷テスト準備

2. **ヘルスチェック統合**:
   - `/health` エンドポイントにDB状態追加
   - 接続プール状態監視

### Phase 5: ドキュメント作成（15分）
1. **README更新**: データベース設定手順
2. **環境変数ガイド**: .env設定例

## 📁 影響ファイル

### 🔧 修正予定ファイル
- `internal/infrastructure/config/app_config.go` - DB設定追加
- `internal/infrastructure/database/connection.go` - 実装完成
- `internal/interface/router/router.go` - ヘルスチェック拡張
- `.env.example` - DB環境変数追加

### 🆕 新規作成ファイル
- `internal/infrastructure/database/migrations/migrate.go`
- `internal/infrastructure/database/migrations/001_initial.sql`
- `internal/infrastructure/database/health.go`

## ⚠️ 注意事項・リスク

### 🔴 高リスク
1. **Supabase接続情報**: 機密情報の適切な管理
2. **接続プール設定**: メモリ使用量との兼ね合い
3. **マイグレーション**: 本番環境での安全な実行

### 🟡 中リスク
1. **GORM設定**: パフォーマンス最適化
2. **ログ出力**: 機密情報の漏洩防止

### 📝 対策
- 環境変数での設定管理徹底
- 段階的実装とテスト
- ログ出力内容の審査

## 🔗 依存関係

### ⬅️ 前提条件
- ✅ BE-02-arch-01: オニオンアーキテクチャ基盤完了
- ✅ Supabaseプロジェクト設定完了
- ✅ 環境変数設定済み

### ➡️ 後続への影響
- **BE-03-user-01**: ユーザーモデル・リポジトリ実装
- **BE-03-***: 全ドメインでのデータベース利用
- **BE-04-***: トランザクション処理実装

## 🎯 成功指標

### ✅ 完了判定基準
1. **接続成功**: Supabaseへの安定接続
2. **GORM動作**: 基本的なCRUD操作実行可能
3. **マイグレーション**: 自動実行成功
4. **ヘルスチェック**: データベース状態監視機能
5. **ビルド成功**: エラーなしでコンパイル完了

### 📊 品質基準
- **接続速度**: 1秒以内での接続確立
- **エラーハンドリング**: 適切なエラーメッセージ
- **ログ品質**: 問題診断に必要な情報出力

## 🚀 次タスクへの準備

BE-02-arch-02完了により以下が可能になります：

1. **BE-03-user-01**: ユーザーエンティティ・リポジトリ実装
2. **BE-03-attendance-01**: 出席管理データモデル実装
3. **全ドメイン**: データベースを使用した実際のビジネスロジック実装

データベース基盤の完成により、**実際のアプリケーション機能開発**に本格移行できます！