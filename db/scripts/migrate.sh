#!/bin/bash
# migrate.sh
# データベースマイグレーション実行スクリプト
# 作成日: 2025-09-28

set -e

# カラー定義
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# ヘルプ表示
show_help() {
    echo "データベースマイグレーション管理スクリプト"
    echo ""
    echo "使用方法:"
    echo "  $0 [オプション]"
    echo ""
    echo "オプション:"
    echo "  up                    全てのマイグレーションを実行"
    echo "  down                  最新のマイグレーションをロールバック"
    echo "  status                マイグレーション状況を確認"
    echo "  seed [environment]    シードデータを投入"
    echo "  reset                 データベースを初期状態にリセット"
    echo "  help                  このヘルプを表示"
    echo ""
    echo "環境変数:"
    echo "  DATABASE_URL          データベース接続URL（必須）"
    echo ""
    echo "例:"
    echo "  $0 up                 # 全マイグレーション実行"
    echo "  $0 seed development   # 開発環境用シードデータ投入"
    echo "  $0 seed testing       # テスト環境用シードデータ投入"
    echo "  $0 seed production    # 本番環境用シードデータ投入"
}

# 環境変数チェック
check_env() {
    if [ -z "$DATABASE_URL" ]; then
        echo -e "${RED}❌ エラー: DATABASE_URL 環境変数が設定されていません${NC}"
        echo "例: export DATABASE_URL='postgresql://user:password@localhost:5432/dbname'"
        exit 1
    fi
}

# データベース接続確認
check_connection() {
    echo -e "${BLUE}🔗 データベース接続を確認中...${NC}"
    if ! psql "$DATABASE_URL" -c "SELECT 1;" > /dev/null 2>&1; then
        echo -e "${RED}❌ データベースに接続できません${NC}"
        exit 1
    fi
    echo -e "${GREEN}✅ データベース接続成功${NC}"
}

# マイグレーションテーブル作成
create_migration_table() {
    echo -e "${BLUE}📋 マイグレーション管理テーブルを確認中...${NC}"
    psql "$DATABASE_URL" -c "
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version VARCHAR(255) PRIMARY KEY,
            applied_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
        );
    " > /dev/null
}

# マイグレーション実行
run_migrations() {
    echo -e "${BLUE}🚀 マイグレーションを実行中...${NC}"
    
    # スクリプトの場所から相対パスを取得
    SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    MIGRATIONS_DIR="$SCRIPT_DIR/../migrations"
    
    # db/migrations ディレクトリのSQLファイルを順番に実行
    for migration_file in "$MIGRATIONS_DIR"/*.sql; do
        if [ -f "$migration_file" ]; then
            filename=$(basename "$migration_file")
            version="${filename%.*}"
            
            # 既に適用されているかチェック
            applied=$(psql "$DATABASE_URL" -t -c "SELECT COUNT(*) FROM schema_migrations WHERE version = '$version';" | tr -d ' ')
            
            if [ "$applied" = "0" ]; then
                echo -e "${YELLOW}📄 実行中: $filename${NC}"
                psql "$DATABASE_URL" -f "$migration_file"
                psql "$DATABASE_URL" -c "INSERT INTO schema_migrations (version) VALUES ('$version');" > /dev/null
                echo -e "${GREEN}✅ 完了: $filename${NC}"
            else
                echo -e "${BLUE}⏭️  スキップ: $filename (既に適用済み)${NC}"
            fi
        fi
    done
    
    echo -e "${GREEN}🎉 全てのマイグレーションが完了しました${NC}"
}

# シードデータ投入
run_seeds() {
    local environment=$1
    
    if [ -z "$environment" ]; then
        echo -e "${RED}❌ 環境を指定してください (development, testing, production)${NC}"
        exit 1
    fi
    
    # スクリプトの場所から相対パスを取得
    SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    local seed_file="$SCRIPT_DIR/../seeds/${environment}.sql"
    
    if [ ! -f "$seed_file" ]; then
        echo -e "${RED}❌ シードファイルが見つかりません: $seed_file${NC}"
        exit 1
    fi
    
    echo -e "${BLUE}🌱 ${environment} 環境のシードデータを投入中...${NC}"
    psql "$DATABASE_URL" -f "$seed_file"
    echo -e "${GREEN}✅ シードデータの投入が完了しました${NC}"
}

# マイグレーション状況確認
show_status() {
    echo -e "${BLUE}📊 マイグレーション状況:${NC}"
    echo ""
    
    # 適用済みマイグレーション一覧
    echo -e "${GREEN}適用済みマイグレーション:${NC}"
    psql "$DATABASE_URL" -c "
        SELECT version, applied_at 
        FROM schema_migrations 
        ORDER BY applied_at;
    " 2>/dev/null || echo "マイグレーション管理テーブルが存在しません"
    
    echo ""
    
    # 未適用マイグレーション確認
    echo -e "${YELLOW}利用可能なマイグレーションファイル:${NC}"
    # スクリプトの場所から相対パスを取得
    SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    MIGRATIONS_DIR="$SCRIPT_DIR/../migrations"
    for migration_file in "$MIGRATIONS_DIR"/*.sql; do
        if [ -f "$migration_file" ]; then
            filename=$(basename "$migration_file")
            version="${filename%.*}"
            applied=$(psql "$DATABASE_URL" -t -c "SELECT COUNT(*) FROM schema_migrations WHERE version = '$version';" 2>/dev/null | tr -d ' ' || echo "0")
            
            if [ "$applied" = "0" ]; then
                echo -e "${RED}❌ $filename (未適用)${NC}"
            else
                echo -e "${GREEN}✅ $filename (適用済み)${NC}"
            fi
        fi
    done
}

# データベースリセット
reset_database() {
    echo -e "${RED}⚠️  警告: データベースを完全にリセットします${NC}"
    echo -e "${RED}この操作により全てのデータが削除されます${NC}"
    echo ""
    read -p "続行しますか？ (yes/no): " confirm
    
    if [ "$confirm" != "yes" ]; then
        echo "キャンセルしました"
        exit 0
    fi
    
    echo -e "${BLUE}🗑️  データベースをリセット中...${NC}"
    
    # 全テーブル削除
    psql "$DATABASE_URL" -c "
        DROP SCHEMA public CASCADE;
        CREATE SCHEMA public;
        GRANT ALL ON SCHEMA public TO postgres;
        GRANT ALL ON SCHEMA public TO public;
        CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";
    "
    
    echo -e "${GREEN}✅ データベースのリセットが完了しました${NC}"
    echo -e "${YELLOW}💡 マイグレーションを再実行してください: $0 up${NC}"
}

# メイン処理
main() {
    case "${1:-help}" in
        "up")
            check_env
            check_connection
            create_migration_table
            run_migrations
            ;;
        "down")
            echo -e "${RED}❌ ロールバック機能は未実装です${NC}"
            echo -e "${YELLOW}💡 必要に応じて手動でSQLを実行してください${NC}"
            exit 1
            ;;
        "status")
            check_env
            check_connection
            create_migration_table
            show_status
            ;;
        "seed")
            check_env
            check_connection
            run_seeds "$2"
            ;;
        "reset")
            check_env
            check_connection
            reset_database
            ;;
        "help")
            show_help
            ;;
        *)
            echo -e "${RED}❌ 不明なオプション: $1${NC}"
            echo ""
            show_help
            exit 1
            ;;
    esac
}

# スクリプト実行
main "$@"