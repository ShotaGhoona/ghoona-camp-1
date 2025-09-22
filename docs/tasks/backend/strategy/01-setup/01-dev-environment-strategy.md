# BE-01-setup-01 開発環境構築戦略

## 🎯 30秒キャッチアップ

**目標**: Ghoona Camp バックエンド開発に必要な環境・ツールの構築  
**作業時間**: 4時間  
**対象**: Docker、Go、PostgreSQL（Supabase）、開発ツールのセットアップ  
**注意**: プロジェクト作成は BE-01-setup-02 で実施

## 📋 実装内容

### 1. Go環境セットアップ（60分）
- Go 1.23+ のインストール・設定
- GOPATH、GOROOT設定確認
- go version、go env動作確認
- 基本的なGoコマンド実行テスト

### 2. Docker環境セットアップ（90分）
- Docker Engine インストール
- Docker Compose インストール・設定
- Docker動作確認（hello-world実行）
- PostgreSQLコンテナ起動テスト

### 3. データベース環境セットアップ（60分）
- Supabaseアカウント作成・設定
- 接続情報取得・環境変数設定
- psqlクライアントでの接続確認
- 基本的なSQL実行テスト

### 4. 開発ツールセットアップ（30分）
- golangci-lint インストール
- gofmt、goimports 設定
- IDE/エディタ設定（Go extension等）
- Git設定確認

## 🔧 環境構築確認事項

### Go環境確認
- `go version` で Go 1.23+ が表示される
- `go env GOPATH` でGOPATHが設定されている
- `go env GOPROXY` でプロキシ設定が適切

### Docker環境確認  
- `docker --version` でDockerバージョン表示
- `docker-compose --version` でComposeバージョン表示
- `docker run hello-world` が正常実行される
- PostgreSQLコンテナが起動できる

### データベース接続確認
- Supabaseプロジェクト作成完了
- 接続文字列の取得・保存
- psqlクライアントでの接続成功
- 基本的なテーブル作成・削除が可能

### 開発ツール確認
- `golangci-lint --version` でバージョン表示
- `gofmt` `goimports` コマンド実行可能
- IDEでGo拡張機能が動作

## 🔧 必要なツール・技術

### 基本環境
- **Go**: 1.23+ (最新LTS推奨)
- **Docker**: 最新安定版
- **Docker Compose**: v2.x
- **Git**: 2.30+

### データベース
- **Supabase**: PostgreSQL SaaS
- **psql**: PostgreSQLクライアント

### 開発ツール
- **golangci-lint**: Goコード品質チェック
- **gofmt**: Goコードフォーマッター
- **goimports**: import文自動整理

## ✅ 完了基準

- [ ] Go 1.23+ インストール・動作確認
- [ ] Docker & Docker Compose インストール・動作確認
- [ ] Supabaseアカウント作成・接続確認
- [ ] PostgreSQLクライアント接続成功
- [ ] 開発ツール（linter等）インストール完了
- [ ] IDEのGo開発環境設定完了

## 🎯 次のステップ
**BE-01-setup-02**: Go modules設定、ディレクトリ構造作成、基本設定ファイル作成