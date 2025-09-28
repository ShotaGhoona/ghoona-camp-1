#!/bin/bash
# db-setup.sh
# データベース初期セットアップスクリプト
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
    echo "データベース初期セットアップスクリプト"
    echo ""
    echo "使用方法:"
    echo "  $0 [環境] [オプション]"
    echo ""
    echo "環境:"
    echo "  development           開発環境セットアップ"
    echo "  testing               テスト環境セットアップ"
    echo "  production            本番環境セットアップ"
    echo ""
    echo "オプション:"
    echo "  --reset               既存データを削除してセットアップ"
    echo "  --migrations-only     マイグレーションのみ実行"
    echo "  --seeds-only          シードデータのみ投入"
    echo "  --help                このヘルプを表示"
    echo ""
    echo "環境変数:"
    echo "  DATABASE_URL          データベース接続URL（必須）"
    echo ""
    echo "例:"
    echo "  $0 development        # 開発環境を完全セットアップ"
    echo "  $0 testing --reset    # テスト環境をリセットしてセットアップ"
    echo "  $0 production         # 本番環境セットアップ（称号のみ）"
}

# 環境変数チェック
check_env() {
    if [ -z "$DATABASE_URL" ]; then
        echo -e "${RED}❌ エラー: DATABASE_URL 環境変数が設定されていません${NC}"
        echo "例: export DATABASE_URL='postgresql://user:password@localhost:5432/dbname'"
        exit 1
    fi
}

# 引数解析
parse_args() {
    ENVIRONMENT=""
    RESET=false
    MIGRATIONS_ONLY=false
    SEEDS_ONLY=false
    
    for arg in "$@"; do
        case $arg in
            --reset)
                RESET=true
                ;;
            --migrations-only)
                MIGRATIONS_ONLY=true
                ;;
            --seeds-only)
                SEEDS_ONLY=true
                ;;
            --help)
                show_help
                exit 0
                ;;
            development|testing|production)
                ENVIRONMENT=$arg
                ;;
            *)
                echo -e "${RED}❌ 不明な引数: $arg${NC}"
                show_help
                exit 1
                ;;
        esac
    done
    
    if [ -z "$ENVIRONMENT" ]; then
        echo -e "${RED}❌ 環境を指定してください${NC}"
        show_help
        exit 1
    fi
}

# セットアップ実行
run_setup() {
    local env=$1
    local script_dir="$(dirname "$0")"
    local migrate_script="$script_dir/migrate.sh"
    
    echo -e "${BLUE}🚀 $env 環境のデータベースセットアップを開始します${NC}"
    echo ""
    
    # リセットオプション
    if [ "$RESET" = true ]; then
        echo -e "${YELLOW}🗑️  データベースをリセットします...${NC}"
        "$migrate_script" reset
        echo ""
    fi
    
    # マイグレーション実行
    if [ "$SEEDS_ONLY" = false ]; then
        echo -e "${BLUE}📋 マイグレーションを実行します...${NC}"
        "$migrate_script" up
        echo ""
    fi
    
    # シードデータ投入
    if [ "$MIGRATIONS_ONLY" = false ]; then
        echo -e "${BLUE}🌱 シードデータを投入します...${NC}"
        "$migrate_script" seed "$env"
        echo ""
    fi
    
    # セットアップ完了メッセージ
    echo -e "${GREEN}🎉 $env 環境のセットアップが完了しました！${NC}"
    echo ""
    
    # 環境別の完了メッセージ
    case $env in
        "development")
            echo -e "${GREEN}開発環境での利用が可能になりました:${NC}"
            echo "  - 充実したテストデータでアプリケーションをテスト"
            echo "  - 6名の開発・テスト用ユーザーアカウント"
            echo "  - サンプルイベントと出席統計データ"
            echo "  - 全ての機能のデモンストレーションが可能"
            ;;
        "testing")
            echo -e "${GREEN}テスト環境での利用が可能になりました:${NC}"
            echo "  - 最小限のテストデータで自動テストを実行"
            echo "  - 3名のテスト用ユーザーアカウント"
            echo "  - CI/CDパイプラインでの動作確認"
            echo "  - 高速なテスト実行が可能"
            ;;
        "production")
            echo -e "${GREEN}本番環境での利用が可能になりました:${NC}"
            echo "  - 称号マスターデータのみ投入済み"
            echo "  - ユーザーの実際の利用でデータが蓄積"
            echo "  - 本番サービスの提供準備完了"
            ;;
    esac
    
    echo ""
    echo -e "${BLUE}📊 セットアップ状況確認:${NC}"
    "$migrate_script" status
}

# メイン処理
main() {
    if [ $# -eq 0 ]; then
        show_help
        exit 1
    fi
    
    parse_args "$@"
    check_env
    run_setup "$ENVIRONMENT"
}

# スクリプト実行
main "$@"