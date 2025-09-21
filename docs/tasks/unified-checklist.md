# Ghoona Camp 開発タスク統合表

## プロジェクト概要
朝活コミュニティアプリ「Ghoona Camp」の全開発タスクを一つの表にまとめました。
バックエンド・インフラ・フロントエンドの全タスクが含まれており、Notionへのインポート用として作成されています。

---

| ステータス | チーム | ID | カテゴリ | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|---|---|
| ◻️ | Backend | BE-01-setup-01 | 📋 開発環境セットアップ | 開発環境の構築 | Docker、Go、PostgreSQL（Supabase）、開発ツールのセットアップ | 4時間 | 🔴 |
| ◻️ | Backend | BE-01-setup-02 | 📋 開発環境セットアップ | プロジェクト初期化 | Go modulesの設定、ディレクトリ構造の作成、基本設定ファイルの作成 | 2時間 | 🔴 |
| ◻️ | Backend | BE-01-setup-03 | 📋 開発環境セットアップ | Docker化設定 | アプリケーションのDocker化、コンテナ最適化 | 3時間 | 🟡 |
| ◻️ | Backend | BE-02-arch-01 | 🏗️ アーキテクチャ基盤実装 | オニオンアーキテクチャ基盤の構築 | domain/application/infrastructure/interface層の基本構造実装 | 8時間 | 🔴 |
| ◻️ | Backend | BE-02-arch-02 | 🏗️ アーキテクチャ基盤実装 | データベース接続設定 | Supabase接続、GORM設定、マイグレーション基盤の実装 | 4時間 | 🔴 |
| ◻️ | Backend | BE-02-arch-03 | 🏗️ アーキテクチャ基盤実装 | Clerk認証システム統合 | Clerk JWT検証、認証ミドルウェア、ユーザー情報連携の実装 | 6時間 | 🔴 |
| ◻️ | Backend | BE-02-arch-04 | 🏗️ アーキテクチャ基盤実装 | エラーハンドリング基盤 | ドメインエラー定義、エラー変換層、HTTPレスポンス標準化 | 4時間 | 🟡 |
| ◻️ | Backend | BE-02-arch-05 | 🏗️ アーキテクチャ基盤実装 | ロギング・監視システム | 構造化ログ、メトリクス収集、アプリケーション監視の実装 | 4時間 | 🟡 |
| ◻️ | Backend | BE-03-user-01 | 👤 ユーザー管理機能 | Userエンティティ・ドメインロジック | ユーザードメインモデル、バリデーション、ビジネスルールの実装 | 4時間 | 🔴 |
| ◻️ | Backend | BE-03-user-02 | 👤 ユーザー管理機能 | UserリポジトリとGORMモデル | ユーザーデータ永続化、リポジトリインターフェース・実装の作成 | 3時間 | 🔴 |
| ◻️ | Backend | BE-03-user-03 | 👤 ユーザー管理機能 | ユーザー管理ユースケース | ユーザー登録・更新・取得・削除のビジネスロジック実装 | 4時間 | 🔴 |
| ◻️ | Backend | BE-03-user-04 | 👤 ユーザー管理機能 | ユーザー管理API | HTTPハンドラー、ルーティング、リクエスト・レスポンスDTO実装 | 3時間 | 🔴 |
| ◻️ | Backend | BE-03-user-05 | 👤 ユーザー管理機能 | ユーザープロフィール機能 | プロフィール画像、自己紹介、設定項目の管理機能 | 4時間 | 🟡 |
| ◻️ | Backend | BE-03-user-06 | 👤 ユーザー管理機能 | ユーザーソーシャルリンク機能 | SNSリンク・外部リンク管理、ライバル設定機能の実装 | 4時間 | 🟢 |
| ◻️ | Backend | BE-04-attend-01 | 📅 出席管理機能 | Attendanceエンティティ・ドメインロジック | 出席記録ドメインモデル、連続記録計算、バリデーション実装 | 5時間 | 🔴 |
| ◻️ | Backend | BE-04-attend-02 | 📅 出席管理機能 | AttendanceリポジトリとGORMモデル | 出席データ永続化、リポジトリインターフェース・実装の作成 | 4時間 | 🔴 |
| ◻️ | Backend | BE-04-attend-03 | 📅 出席管理機能 | 出席記録ユースケース | 出席登録・取得・統計計算のビジネスロジック実装 | 6時間 | 🔴 |
| ◻️ | Backend | BE-04-attend-04 | 📅 出席管理機能 | 出席管理API | 出席記録CRUD、統計取得APIのハンドラー・ルーティング実装 | 4時間 | 🔴 |
| ◻️ | Backend | BE-04-attend-05 | 📅 出席管理機能 | 連続出席・ランキング機能 | 連続記録計算、ユーザーランキング、アチーブメント判定 | 6時間 | 🟡 |
| ◻️ | Backend | BE-05-discord-01 | 🤖 Discord Bot統合機能 | Discord API統合基盤 | Discord Bot認証、API接続、基本構造の実装 | 6時間 | 🔴 |
| ◻️ | Backend | BE-05-discord-02 | 🤖 Discord Bot統合機能 | Discord参加ログ記録システム | 音声チャンネル入退室の自動記録、出席判定ロジック実装 | 8時間 | 🔴 |
| ◻️ | Backend | BE-05-discord-03 | 🤖 Discord Bot統合機能 | Discord Webhook処理 | Discordイベント受信、ユーザーID連携、リアルタイム処理 | 6時間 | 🔴 |
| ◻️ | Backend | BE-06-goal-01 | 🎯 目標管理機能 | Goalエンティティ・ドメインロジック | 目標ドメインモデル、進捗計算、達成判定の実装 | 5時間 | 🟡 |
| ◻️ | Backend | BE-06-goal-02 | 🎯 目標管理機能 | GoalリポジトリとGORMモデル | 目標データ永続化、リポジトリインターフェース・実装の作成 | 3時間 | 🟡 |
| ◻️ | Backend | BE-06-goal-03 | 🎯 目標管理機能 | 目標管理ユースケース | 目標作成・更新・削除・進捗更新のビジネスロジック実装 | 5時間 | 🟡 |
| ◻️ | Backend | BE-06-goal-04 | 🎯 目標管理機能 | 目標管理API | 目標CRUD、進捗取得APIのハンドラー・ルーティング実装 | 4時間 | 🟡 |
| ◻️ | Backend | BE-06-goal-05 | 🎯 目標管理機能 | 目標カテゴリ・テンプレート機能 | 目標カテゴリ管理、推奨目標テンプレート提供機能 | 4時間 | 🟢 |
| ◻️ | Backend | BE-07-event-01 | 🎪 イベント管理機能 | Eventエンティティ・ドメインロジック | イベントドメインモデル、参加管理、状態遷移の実装 | 5時間 | 🟡 |
| ◻️ | Backend | BE-07-event-02 | 🎪 イベント管理機能 | EventリポジトリとGORMモデル | イベントデータ永続化、リポジトリインターフェース・実装の作成 | 4時間 | 🟡 |
| ◻️ | Backend | BE-07-event-03 | 🎪 イベント管理機能 | イベント管理ユースケース | イベント作成・更新・削除・参加管理のビジネスロジック実装 | 6時間 | 🟡 |
| ◻️ | Backend | BE-07-event-04 | 🎪 イベント管理機能 | イベント管理API | イベントCRUD、参加管理APIのハンドラー・ルーティング実装 | 4時間 | 🟡 |
| ◻️ | Backend | BE-07-event-05 | 🎪 イベント管理機能 | イベント検索・フィルタリング機能 | 日時・カテゴリ・参加状況による検索・フィルタリング機能 | 4時間 | 🟢 |
| ◻️ | Backend | BE-08-title-01 | 🏆 称号管理機能 | Titleエンティティ・ドメインロジック | 称号ドメインモデル、獲得条件判定、称号レベル管理の実装 | 4時間 | 🟡 |
| ◻️ | Backend | BE-08-title-02 | 🏆 称号管理機能 | TitleリポジトリとGORMモデル | 称号データ永続化、リポジトリインターフェース・実装の作成 | 3時間 | 🟡 |
| ◻️ | Backend | BE-08-title-03 | 🏆 称号管理機能 | 称号管理ユースケース | 称号獲得判定・付与・変更のビジネスロジック実装 | 5時間 | 🟡 |
| ◻️ | Backend | BE-08-title-04 | 🏆 称号管理機能 | 称号管理API | 称号取得・変更APIのハンドラー・ルーティング実装 | 3時間 | 🟡 |
| ◻️ | Backend | BE-08-title-05 | 🏆 称号管理機能 | アチーブメントシステム | 称号獲得条件の自動判定・通知システムの実装 | 6時間 | 🟢 |
| ◻️ | Backend | BE-09-notify-01 | 📢 通知管理機能 | Notificationエンティティ・ドメインロジック | 通知ドメインモデル、通知タイプ管理、配信設定の実装 | 4時間 | 🟡 |
| ◻️ | Backend | BE-09-notify-02 | 📢 通知管理機能 | NotificationリポジトリとGORMモデル | 通知データ永続化、リポジトリインターフェース・実装の作成 | 3時間 | 🟡 |
| ◻️ | Backend | BE-09-notify-03 | 📢 通知管理機能 | 通知管理ユースケース | 通知作成・配信・既読管理のビジネスロジック実装 | 5時間 | 🟡 |
| ◻️ | Backend | BE-09-notify-04 | 📢 通知管理機能 | 通知管理API | 通知CRUD、既読管理APIのハンドラー・ルーティング実装 | 3時間 | 🟡 |
| ◻️ | Backend | BE-09-notify-05 | 📢 通知管理機能 | リアルタイム通知システム | WebSocket・プッシュ通知による即座通知システムの実装 | 8時間 | 🟢 |
| ◻️ | Backend | BE-10-batch-01 | ⏰ バッチ処理システム | バッチ処理基盤構築 | 定期実行基盤、エラーハンドリング、監視・ログ機能の実装 | 6時間 | 🟡 |
| ◻️ | Backend | BE-10-batch-02 | ⏰ バッチ処理システム | 日次バッチ処理実装 | Discord参加ログ集計、統計更新、称号付与の自動処理 | 8時間 | 🟡 |
| ◻️ | Backend | BE-10-batch-03 | ⏰ バッチ処理システム | 週次バッチ処理実装 | 週次リマインダー送信、ライバル更新通知の自動処理 | 6時間 | 🟢 |
| ◻️ | Backend | BE-10-batch-04 | ⏰ バッチ処理システム | 月次バッチ処理実装 | 月次ランキング確定、古いデータクリーンアップの自動処理 | 6時間 | 🟢 |
| ◻️ | Backend | BE-10-batch-05 | ⏰ バッチ処理システム | バッチ処理監視・復旧システム | 処理失敗検知、自動リトライ、手動実行機能の実装 | 4時間 | 🟢 |
| ◻️ | Backend | BE-11-system-01 | 🔧 システム機能 | ヘルスチェック・監視API | アプリケーション・DB状態監視、メトリクス公開API | 3時間 | 🟡 |
| ◻️ | Backend | BE-11-system-02 | 🔧 システム機能 | API レート制限 | API利用制限、適切なレスポンス制御の実装 | 4時間 | 🟡 |
| ◻️ | Backend | BE-11-system-03 | 🔧 システム機能 | セキュリティ強化 | CORS設定、セキュリティヘッダー、入力検証の強化 | 4時間 | 🔴 |
| ◻️ | Backend | BE-12-test-01 | 🧪 テスト実装 | ドメイン層ユニットテスト | エンティティ・値オブジェクト・ドメインサービスのテスト | 12時間 | 🟡 |
| ◻️ | Backend | BE-12-test-02 | 🧪 テスト実装 | アプリケーション層テスト | ユースケース・DTOのテスト、モック使用統合テスト | 10時間 | 🟡 |
| ◻️ | Backend | BE-12-test-03 | 🧪 テスト実装 | インフラストラクチャ層テスト | リポジトリ実装・外部サービス統合のテスト | 8時間 | 🟡 |
| ◻️ | Backend | BE-12-test-04 | 🧪 テスト実装 | API統合テスト | HTTPエンドポイント・認証・エラーハンドリングのテスト | 10時間 | 🟡 |
| ◻️ | Backend | BE-12-test-05 | 🧪 テスト実装 | パフォーマンステスト | 負荷テスト・メモリ使用量・レスポンス時間の測定 | 6時間 | 🟢 |
| ◻️ | Backend | BE-12-test-06 | 🧪 テスト実装 | Discord Bot統合テスト | Discord API連携、Webhook処理、バッチ処理の統合テスト | 8時間 | 🟢 |
| ◻️ | Backend | BE-13-docs-01 | 📝 ドキュメント作成 | API仕様書作成 | OpenAPI/Swagger仕様書、エンドポイント詳細ドキュメント | 8時間 | 🟡 |
| ◻️ | Backend | BE-13-docs-02 | 📝 ドキュメント作成 | 開発ガイド作成 | アーキテクチャ説明、開発手順、ベストプラクティス | 4時間 | 🟢 |
| ◻️ | Infrastructure | INF-01-infra-01 | 🏗️ インフラストラクチャ基盤 | AWS アカウント・IAM設定 | AWSアカウント作成、IAMユーザー・ロール・ポリシー設定 | 4時間 | 🔴 |
| ◻️ | Infrastructure | INF-01-infra-02 | 🏗️ インフラストラクチャ基盤 | VPC・ネットワーク設計 | VPC、サブネット、セキュリティグループ、ルートテーブル設定 | 6時間 | 🔴 |
| ◻️ | Infrastructure | INF-01-infra-03 | 🏗️ インフラストラクチャ基盤 | Supabase環境構築 | Supabaseプロジェクト作成、データベース設定、接続設定 | 3時間 | 🔴 |
| ◻️ | Infrastructure | INF-01-infra-04 | 🏗️ インフラストラクチャ基盤 | CDN・ドメイン設定 | CloudFront設定、Route53ドメイン管理、SSL証明書設定 | 4時間 | 🟡 |
| ◻️ | Infrastructure | INF-01-infra-05 | 🏗️ インフラストラクチャ基盤 | 環境分離設計 | 開発・ステージング・本番環境の分離、命名規則設定 | 3時間 | 🔴 |
| ◻️ | Infrastructure | INF-02-container-01 | 🐳 コンテナ・オーケストレーション | Docker基盤構築 | Dockerfile最適化、マルチステージビルド、セキュリティ設定 | 4時間 | 🔴 |
| ◻️ | Infrastructure | INF-02-container-02 | 🐳 コンテナ・オーケストレーション | ECR設定 | Amazon ECRリポジトリ作成、プッシュ・プル権限設定 | 2時間 | 🔴 |
| ◻️ | Infrastructure | INF-02-container-03 | 🐳 コンテナ・オーケストレーション | ECS/Fargate設定 | ECSクラスター、タスク定義、サービス設定 | 8時間 | 🔴 |
| ◻️ | Infrastructure | INF-02-container-04 | 🐳 コンテナ・オーケストレーション | ロードバランサー設定 | ALB設定、ヘルスチェック、ターゲットグループ管理 | 4時間 | 🔴 |
| ◻️ | Infrastructure | INF-02-container-05 | 🐳 コンテナ・オーケストレーション | Auto Scaling設定 | CPU・メモリベースのオートスケーリング設定 | 3時間 | 🟡 |
| ◻️ | Infrastructure | INF-03-cicd-01 | 🚀 CI/CD パイプライン | GitHub Actions基盤 | ワークフロー設定、シークレット管理、権限設定 | 4時間 | 🔴 |
| ◻️ | Infrastructure | INF-03-cicd-02 | 🚀 CI/CD パイプライン | バックエンドCI/CD | Go APIのビルド・テスト・デプロイパイプライン | 6時間 | 🔴 |
| ◻️ | Infrastructure | INF-03-cicd-03 | 🚀 CI/CD パイプライン | フロントエンドCI/CD | Next.jsのビルド・テスト・デプロイパイプライン | 6時間 | 🔴 |
| ◻️ | Infrastructure | INF-03-cicd-04 | 🚀 CI/CD パイプライン | Discord BotCI/CD | Bot専用のデプロイパイプライン、環境変数管理 | 4時間 | 🔴 |
| ◻️ | Infrastructure | INF-03-cicd-05 | 🚀 CI/CD パイプライン | 並列実行最適化 | パイプライン並列化、キャッシュ戦略、実行時間短縮 | 4時間 | 🟡 |
| ◻️ | Infrastructure | INF-04-monitor-01 | 📊 監視・ロギング | CloudWatch基盤 | ログ収集、メトリクス設定、ダッシュボード作成 | 6時間 | 🟡 |
| ◻️ | Infrastructure | INF-04-monitor-02 | 📊 監視・ロギング | アプリケーション監視 | APM設定、レスポンス時間・エラー率監視 | 4時間 | 🟡 |
| ◻️ | Infrastructure | INF-04-monitor-03 | 📊 監視・ロギング | インフラ監視 | CPU・メモリ・ディスク・ネットワーク監視 | 4時間 | 🟡 |
| ◻️ | Infrastructure | INF-04-monitor-04 | 📊 監視・ロギング | アラート設定 | 閾値設定、SNS通知、Slack連携、エスカレーション | 4時間 | 🟡 |
| ◻️ | Infrastructure | INF-04-monitor-05 | 📊 監視・ロギング | ログ分析基盤 | 構造化ログ、検索・分析環境、ログ保存期間設定 | 5時間 | 🟡 |
| ◻️ | Infrastructure | INF-05-security-01 | 🔒 セキュリティ・コンプライアンス | WAF・DDoS対策 | AWS WAF設定、DDoS Protection、レート制限 | 4時間 | 🔴 |
| ◻️ | Infrastructure | INF-05-security-02 | 🔒 セキュリティ・コンプライアンス | シークレット管理 | AWS Secrets Manager、環境変数暗号化、ローテーション | 3時間 | 🔴 |
| ◻️ | Infrastructure | INF-05-security-03 | 🔒 セキュリティ・コンプライアンス | ネットワークセキュリティ | NACLs、セキュリティグループ最適化、VPN設定 | 4時間 | 🔴 |
| ◻️ | Infrastructure | INF-05-security-04 | 🔒 セキュリティ・コンプライアンス | 脆弱性スキャン | 定期脆弱性スキャン、セキュリティパッチ適用自動化 | 3時間 | 🟡 |
| ◻️ | Infrastructure | INF-05-security-05 | 🔒 セキュリティ・コンプライアンス | アクセスログ・監査 | CloudTrail、Config設定、コンプライアンス監視 | 4時間 | 🟡 |
| ◻️ | Infrastructure | INF-06-backup-01 | 💾 バックアップ・災害復旧 | データベースバックアップ | Supabase自動バックアップ、ポイントインタイム復旧設定 | 3時間 | 🟡 |
| ◻️ | Infrastructure | INF-06-backup-02 | 💾 バックアップ・災害復旧 | アプリケーションバックアップ | 設定ファイル、シークレット、デプロイメント資産のバックアップ | 3時間 | 🟡 |
| ◻️ | Infrastructure | INF-06-backup-03 | 💾 バックアップ・災害復旧 | 災害復旧計画 | RTO・RPO設定、復旧手順書、定期復旧テスト | 6時間 | 🟡 |
| ◻️ | Infrastructure | INF-06-backup-04 | 💾 バックアップ・災害復旧 | データアーカイブ | 古いデータの自動アーカイブ、コスト最適化 | 3時間 | 🟢 |
| ◻️ | Infrastructure | INF-07-optimize-01 | 🚦 パフォーマンス・コスト最適化 | パフォーマンス監視 | レスポンス時間、スループット、リソース使用率の継続監視 | 4時間 | 🟡 |
| ◻️ | Infrastructure | INF-07-optimize-02 | 🚦 パフォーマンス・コスト最適化 | キャッシュ戦略 | CloudFront、ElastiCache設定、キャッシュ最適化 | 5時間 | 🟢 |
| ◻️ | Infrastructure | INF-07-optimize-03 | 🚦 パフォーマンス・コスト最適化 | コスト監視・最適化 | Cost Explorer、予算アラート、リソース使用量最適化 | 4時間 | 🟢 |
| ◻️ | Infrastructure | INF-07-optimize-04 | 🚦 パフォーマンス・コスト最適化 | リソース自動停止 | 開発環境の自動停止、スケジュール管理 | 3時間 | 🟢 |
| ◻️ | Infrastructure | INF-08-qa-01 | 📋 プロジェクト管理・QA | 要件定義・仕様策定 | 機能要件・非機能要件の詳細化、承認プロセス | 8時間 | 🟢 |
| ◻️ | Infrastructure | INF-08-qa-02 | 📋 プロジェクト管理・QA | テスト計画策定 | テスト戦略、テストケース設計、受け入れ基準定義 | 6時間 | 🟢 |
| ◻️ | Infrastructure | INF-08-qa-03 | 📋 プロジェクト管理・QA | QA環境構築 | テスト専用環境、テストデータ準備、自動テスト基盤 | 8時間 | 🟢 |
| ◻️ | Infrastructure | INF-08-qa-04 | 📋 プロジェクト管理・QA | ユーザビリティテスト | ユーザーテスト計画、実施、フィードバック収集・分析 | 10時間 | 🟢 |
| ◻️ | Infrastructure | INF-08-qa-05 | 📋 プロジェクト管理・QA | セキュリティテスト | ペネトレーションテスト、脆弱性診断、セキュリティ監査 | 8時間 | 🟢 |
| ◻️ | Infrastructure | INF-09-release-01 | 🚀 リリース管理・運用 | リリース戦略策定 | ブルーグリーンデプロイ、カナリアリリース、ロールバック戦略 | 4時間 | 🟡 |
| ◻️ | Infrastructure | INF-09-release-02 | 🚀 リリース管理・運用 | 本番リリース準備 | 本番環境最終確認、データ移行、リリースチェックリスト | 6時間 | 🟡 |
| ◻️ | Infrastructure | INF-09-release-03 | 🚀 リリース管理・運用 | 運用監視体制 | 24/7監視体制、インシデント対応手順、エスカレーション | 6時間 | 🟡 |
| ◻️ | Infrastructure | INF-09-release-04 | 🚀 リリース管理・運用 | パフォーマンステスト | 負荷テスト、ストレステスト、キャパシティプランニング | 8時間 | 🟢 |
| ◻️ | Infrastructure | INF-09-release-05 | 🚀 リリース管理・運用 | 運用自動化 | 定期メンテナンス自動化、ヘルスチェック、自動復旧 | 6時間 | 🟢 |
| ◻️ | Infrastructure | INF-10-docs-01 | 📝 ドキュメント・ナレッジ管理 | インフラ設計書作成 | アーキテクチャ図、ネットワーク構成図、セキュリティ要件 | 8時間 | 🟢 |
| ◻️ | Infrastructure | INF-10-docs-02 | 📝 ドキュメント・ナレッジ管理 | 運用手順書作成 | デプロイ手順、トラブルシューティング、緊急対応手順 | 6時間 | 🟢 |
| ◻️ | Infrastructure | INF-10-docs-03 | 📝 ドキュメント・ナレッジ管理 | 開発者向けガイド | 環境構築手順、デバッグ方法、FAQ作成 | 4時間 | 🟢 |
| ◻️ | Infrastructure | INF-10-docs-04 | 📝 ドキュメント・ナレッジ管理 | 監視・アラート設定書 | 監視項目一覧、アラート設定、対応手順書 | 4時間 | 🟢 |
| ◻️ | Infrastructure | INF-11-tools-01 | 🔧 ツール・自動化 | Infrastructure as Code | Terraform/CDK、設定管理、バージョン管理 | 10時間 | 🟢 |
| ◻️ | Infrastructure | INF-11-tools-02 | 🔧 ツール・自動化 | 開発者体験向上 | ローカル開発環境自動化、開発ツール整備 | 6時間 | 🟢 |
| ◻️ | Infrastructure | INF-11-tools-03 | 🔧 ツール・自動化 | 自動テスト基盤 | CI/CDでの自動テスト実行、テストレポート生成 | 6時間 | 🟢 |
| ◻️ | Frontend | FE-01-setup-01 | 📋 開発環境セットアップ | 開発環境の構築 | Node.js、Next.js 15、TypeScript、必要な開発ツールのセットアップ | 3時間 | 🔴 |
| ◻️ | Frontend | FE-01-setup-02 | 📋 開発環境セットアップ | プロジェクト初期化とFSD構造作成 | Next.js初期化、FSDディレクトリ構造の作成、基本設定ファイル | 4時間 | 🔴 |
| ◻️ | Frontend | FE-01-setup-03 | 📋 開発環境セットアップ | 開発ツール・Linter設定 | ESLint、Prettier、FSD boundaries、pre-commit hooks設定 | 2時間 | 🔴 |
| ◻️ | Frontend | FE-02-shared-01 | 🔧 Shared層（共有基盤） | shadcn/ui セットアップ | shadcn/uiライブラリ導入、基本コンポーネントの設定 | 3時間 | 🔴 |
| ◻️ | Frontend | FE-02-shared-02 | 🔧 Shared層（共有基盤） | デザインシステム構築 | カラーパレット、タイポグラフィ、スペーシング、テーマ設定 | 6時間 | 🔴 |
| ◻️ | Frontend | FE-02-shared-03 | 🔧 Shared層（共有基盤） | 共通UIコンポーネント | Button、Card、Modal、Form、Loading等の基本コンポーネント | 8時間 | 🔴 |
| ◻️ | Frontend | FE-02-shared-04 | 🔧 Shared層（共有基盤） | APIクライアント設定 | Supabaseクライアント、Clerkクライアント、HTTP設定 | 4時間 | 🔴 |
| ◻️ | Frontend | FE-02-shared-05 | 🔧 Shared層（共有基盤） | 共通フック・ユーティリティ | 共通カスタムフック、日付・フォーマット・バリデーション関数 | 5時間 | 🔴 |
| ◻️ | Frontend | FE-02-shared-06 | 🔧 Shared層（共有基盤） | エラーハンドリング基盤 | Error Boundary、Toast通知、エラーページ実装 | 4時間 | 🔴 |
| ◻️ | Frontend | FE-03-user-01 | # User Entity | User型定義・モデル | ユーザー関連の型定義、インターフェース、バリデーションスキーマ | 3時間 | 🔴 |
| ◻️ | Frontend | FE-03-user-02 | # User Entity | User API クライアント | ユーザー情報取得・更新、プロフィール管理のAPI呼び出し | 4時間 | 🔴 |
| ◻️ | Frontend | FE-03-user-03 | # User Entity | User クエリキー管理 | React Queryのクエリキー一元管理、無効化戦略 | 2時間 | 🔴 |
| ◻️ | Frontend | FE-03-attend-01 | # Attendance Entity | Attendance型定義・モデル | 出席ログ、統計、ランキング関連の型定義とバリデーション | 4時間 | 🔴 |
| ◻️ | Frontend | FE-03-attend-02 | # Attendance Entity | Attendance API クライアント | 出席ログ取得、統計取得、ランキング取得のAPI呼び出し | 5時間 | 🔴 |
| ◻️ | Frontend | FE-03-attend-03 | # Attendance Entity | Attendance クエリキー管理 | 出席関連データのクエリキー管理、リアルタイム更新 | 2時間 | 🔴 |
| ◻️ | Frontend | FE-03-goal-01 | # Goal Entity | Goal型定義・モデル | 目標、進捗関連の型定義とバリデーションスキーマ | 3時間 | 🟡 |
| ◻️ | Frontend | FE-03-goal-02 | # Goal Entity | Goal API クライアント | 目標CRUD、進捗管理のAPI呼び出し | 4時間 | 🟡 |
| ◻️ | Frontend | FE-03-goal-03 | # Goal Entity | Goal クエリキー管理 | 目標関連データのクエリキー管理 | 2時間 | 🟡 |
| ◻️ | Frontend | FE-03-event-01 | # Event Entity | Event型定義・モデル | イベント、参加者関連の型定義とバリデーション | 3時間 | 🟡 |
| ◻️ | Frontend | FE-03-event-02 | # Event Entity | Event API クライアント | イベントCRUD、参加管理のAPI呼び出し | 4時間 | 🟡 |
| ◻️ | Frontend | FE-03-event-03 | # Event Entity | Event クエリキー管理 | イベント関連データのクエリキー管理 | 2時間 | 🟡 |
| ◻️ | Frontend | FE-03-title-01 | # Title Entity | Title型定義・モデル | 称号、実績関連の型定義とバリデーション | 3時間 | 🟡 |
| ◻️ | Frontend | FE-03-title-02 | # Title Entity | Title API クライアント | 称号取得、実績管理のAPI呼び出し | 3時間 | 🟡 |
| ◻️ | Frontend | FE-03-title-03 | # Title Entity | Title クエリキー管理 | 称号関連データのクエリキー管理 | 2時間 | 🟡 |
| ◻️ | Frontend | FE-03-notify-01 | # Notification Entity | Notification型定義・モデル | 通知、設定関連の型定義とバリデーション | 3時間 | 🟡 |
| ◻️ | Frontend | FE-03-notify-02 | # Notification Entity | Notification API クライアント | 通知CRUD、設定管理のAPI呼び出し | 3時間 | 🟡 |
| ◻️ | Frontend | FE-03-notify-03 | # Notification Entity | Notification クエリキー管理 | 通知関連データのクエリキー管理 | 2時間 | 🟡 |
| ◻️ | Frontend | FE-04-attend-01 | # Attendance Features | attendance-log-get feature | 出席ログ取得機能、カスタムフック、エラーハンドリング | 4時間 | 🔴 |
| ◻️ | Frontend | FE-04-attend-02 | # Attendance Features | attendance-stats-get feature | 出席統計取得機能、連続日数計算、進捗表示 | 5時間 | 🔴 |
| ◻️ | Frontend | FE-04-attend-03 | # Attendance Features | attendance-ranking-get feature | ランキング取得機能、フィルタリング、ソート | 4時間 | 🔴 |
| ◻️ | Frontend | FE-04-attend-04 | # Attendance Features | attendance-streak-get feature | 連続記録取得機能、アチーブメント判定 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-04-user-01 | # User Features | user-profile-get feature | ユーザープロフィール取得機能、メタデータ管理 | 3時間 | 🔴 |
| ◻️ | Frontend | FE-04-user-02 | # User Features | user-profile-update feature | プロフィール更新機能、画像アップロード、バリデーション | 5時間 | 🔴 |
| ◻️ | Frontend | FE-04-user-03 | # User Features | user-rivals-manage feature | ライバル管理機能、追加・削除、最大3人制限 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-04-user-04 | # User Features | user-social-links feature | ソーシャルリンク管理機能、CRUD操作 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-04-goal-01 | # Goal Features | goal-create feature | 目標作成機能、フォーム、バリデーション | 4時間 | 🟡 |
| ◻️ | Frontend | FE-04-goal-02 | # Goal Features | goal-update feature | 目標更新機能、進捗管理、公開設定 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-04-goal-03 | # Goal Features | goal-delete feature | 目標削除機能、確認ダイアログ、データ整合性 | 2時間 | 🟡 |
| ◻️ | Frontend | FE-04-goal-04 | # Goal Features | goal-list-get feature | 目標一覧取得機能、フィルタリング、ソート | 3時間 | 🟡 |
| ◻️ | Frontend | FE-04-event-01 | # Event Features | event-create feature | イベント作成機能、日時選択、参加者設定 | 5時間 | 🟡 |
| ◻️ | Frontend | FE-04-event-02 | # Event Features | event-participate feature | イベント参加機能、申込・キャンセル、定員管理 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-04-event-03 | # Event Features | event-list-get feature | イベント一覧取得機能、検索・フィルター | 3時間 | 🟡 |
| ◻️ | Frontend | FE-04-event-04 | # Event Features | event-update feature | イベント更新機能、作成者権限チェック | 3時間 | 🟡 |
| ◻️ | Frontend | FE-04-title-01 | # Title Features | title-get feature | 全称号取得機能、獲得条件表示 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-04-title-02 | # Title Features | achievement-get feature | 獲得履歴取得機能、現在設定中の称号 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-04-title-03 | # Title Features | title-change feature | 称号変更機能、獲得済み称号から選択 | 2時間 | 🟡 |
| ◻️ | Frontend | FE-04-notify-01 | # Notification Features | notification-get feature | 通知取得機能、未読管理、ページネーション | 4時間 | 🟡 |
| ◻️ | Frontend | FE-04-notify-02 | # Notification Features | notification-read feature | 既読処理機能、一括既読、自動既読 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-04-notify-03 | # Notification Features | notification-settings feature | 通知設定機能、カテゴリ別ON/OFF、時刻設定 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-04-auth-01 | # Auth Features | clerk-auth-integration | Clerk認証統合、ログイン・サインアップフロー | 5時間 | 🔴 |
| ◻️ | Frontend | FE-04-auth-02 | # Auth Features | auth-guard feature | 認証ガード機能、ルート保護、リダイレクト | 3時間 | 🔴 |
| ◻️ | Frontend | FE-04-auth-03 | # Auth Features | auth-state-manage | 認証状態管理、React Query統合、Zustand連携 | 4時間 | 🔴 |
| ◻️ | Frontend | FE-05-attend-01 | # Attendance Widgets | AttendanceCard widget | 出席カード表示、統計サマリー、連続日数 | 4時間 | 🔴 |
| ◻️ | Frontend | FE-05-attend-02 | # Attendance Widgets | StreakBadge widget | 連続日数バッジ、レベル表示、アニメーション | 3時間 | 🔴 |
| ◻️ | Frontend | FE-05-attend-03 | # Attendance Widgets | AttendanceChart widget | 出席グラフ、月次表示、データ可視化 | 6時間 | 🟡 |
| ◻️ | Frontend | FE-05-attend-04 | # Attendance Widgets | CalendarView widget | カレンダー表示、出席状況、日付選択 | 8時間 | 🟡 |
| ◻️ | Frontend | FE-05-attend-05 | # Attendance Widgets | RankingList widget | ランキング一覧、ライバル強調、フィルタリング | 5時間 | 🟡 |
| ◻️ | Frontend | FE-05-goal-01 | # Goal Widgets | GoalCard widget | 目標カード表示、進捗バー、期限表示 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-05-goal-02 | # Goal Widgets | ProgressBar widget | 進捗バー、パーセンテージ、アニメーション | 3時間 | 🟡 |
| ◻️ | Frontend | FE-05-goal-03 | # Goal Widgets | GoalForm widget | 目標フォーム、バリデーション、公開設定 | 5時間 | 🟡 |
| ◻️ | Frontend | FE-05-goal-04 | # Goal Widgets | GoalList widget | 目標一覧表示、フィルタリング、ソート | 4時間 | 🟡 |
| ◻️ | Frontend | FE-05-event-01 | # Event Widgets | EventCard widget | イベントカード、参加状況、詳細リンク | 4時間 | 🟡 |
| ◻️ | Frontend | FE-05-event-02 | # Event Widgets | EventList widget | イベント一覧、検索・フィルター、ページネーション | 5時間 | 🟡 |
| ◻️ | Frontend | FE-05-event-03 | # Event Widgets | ParticipantList widget | 参加者一覧、アバター表示、参加状況 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-05-event-04 | # Event Widgets | EventForm widget | イベントフォーム、日時選択、画像アップロード | 6時間 | 🟡 |
| ◻️ | Frontend | FE-05-title-01 | # Title Widgets | TitleCard widget | 称号カード、詳細情報、ストーリー | 4時間 | 🟡 |
| ◻️ | Frontend | FE-05-title-02 | # Title Widgets | TitleBadge widget | 称号バッジ、レベル表示、テーマカラー | 3時間 | 🟡 |
| ◻️ | Frontend | FE-05-title-03 | # Title Widgets | AchievementModal widget | 称号獲得モーダル、アニメーション、お祝い演出 | 5時間 | 🟢 |
| ◻️ | Frontend | FE-05-title-04 | # Title Widgets | TitleGrid widget | 称号一覧グリッド、カテゴリ別表示、進捗表示 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-05-notify-01 | # Notification Widgets | NotificationList widget | 通知一覧、未読表示、アクション | 4時間 | 🟡 |
| ◻️ | Frontend | FE-05-notify-02 | # Notification Widgets | NotificationBadge widget | 通知バッジ、未読件数、リアルタイム更新 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-05-notify-03 | # Notification Widgets | NotificationSettings widget | 通知設定フォーム、カテゴリ別設定 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-05-layout-01 | # Layout Widgets | Header widget | グローバルヘッダー、ユーザーメニュー、ナビゲーション | 5時間 | 🔴 |
| ◻️ | Frontend | FE-05-layout-02 | # Layout Widgets | Sidebar widget | サイドナビゲーション、メニュー状態管理、モバイル対応 | 5時間 | 🔴 |
| ◻️ | Frontend | FE-05-layout-03 | # Layout Widgets | Footer widget | フッター、リンク集、会社情報 | 2時間 | 🟡 |
| ◻️ | Frontend | FE-06-dash-01 | # Dashboard Pages | Dashboard page component | ダッシュボードページ統合、widgetsオーケストレーション | 4時間 | 🔴 |
| ◻️ | Frontend | FE-06-dash-02 | # Dashboard Pages | Dashboard hooks | ダッシュボード専用フック、データ統合、状態管理 | 3時間 | 🔴 |
| ◻️ | Frontend | FE-06-attend-01 | # Attendance Pages | Attendance page component | 出席記録ページ統合、カレンダー・統計表示 | 4時間 | 🔴 |
| ◻️ | Frontend | FE-06-attend-02 | # Attendance Pages | Attendance hooks | 出席ページ専用フック、データフィルタリング | 3時間 | 🔴 |
| ◻️ | Frontend | FE-06-goal-01 | # Goal Pages | Goals page component | 目標一覧ページ統合、CRUD操作統合 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-06-goal-02 | # Goal Pages | Goal detail page component | 目標詳細ページ、編集・削除機能統合 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-06-goal-03 | # Goal Pages | Goals page hooks | 目標ページ専用フック、フィルタリング・ソート | 3時間 | 🟡 |
| ◻️ | Frontend | FE-06-event-01 | # Event Pages | Events page component | イベント一覧ページ統合、検索・フィルター統合 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-06-event-02 | # Event Pages | Event detail page component | イベント詳細ページ、参加管理統合 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-06-event-03 | # Event Pages | Events page hooks | イベントページ専用フック、参加状況管理 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-06-title-01 | # Title Pages | Titles page component | 称号一覧ページ統合、獲得状況表示 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-06-title-02 | # Title Pages | Titles page hooks | 称号ページ専用フック、進捗計算 | 2時間 | 🟡 |
| ◻️ | Frontend | FE-06-settings-01 | # Settings Pages | Profile settings page component | プロフィール設定ページ統合、フォーム管理 | 4時間 | 🔴 |
| ◻️ | Frontend | FE-06-settings-02 | # Settings Pages | Notification settings page component | 通知設定ページ統合、設定管理 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-06-settings-03 | # Settings Pages | Account settings page component | アカウント設定ページ統合、Clerk連携 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-06-landing-01 | # Landing Pages | Landing page component | ランディングページ、未認証ユーザー向け | 5時間 | 🟡 |
| ◻️ | Frontend | FE-06-landing-02 | # Landing Pages | Landing page hooks | ランディング専用フック、アニメーション管理 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-07-app-01 | 🔗 App層（Next.js統合） | Root Layout設定 | Root Layout、グローバルプロバイダー設定 | 3時間 | 🔴 |
| ◻️ | Frontend | FE-07-app-02 | 🔗 App層（Next.js統合） | ルート定義 | 全ページのルート定義、page.tsx、layout.tsx作成 | 4時間 | 🔴 |
| ◻️ | Frontend | FE-07-app-03 | 🔗 App層（Next.js統合） | Loading・Error UI | Loading UI、Error UI、Not Found実装 | 3時間 | 🔴 |
| ◻️ | Frontend | FE-07-app-04 | 🔗 App層（Next.js統合） | Providers設定 | Clerk、React Query、Zustand等のプロバイダー統合 | 4時間 | 🔴 |
| ◻️ | Frontend | FE-07-app-05 | 🔗 App層（Next.js統合） | Metadata設定 | SEO・メタタグ最適化、OGP設定 | 3時間 | 🟡 |
| ◻️ | Frontend | FE-08-ui-01 | 🎨 高度なUI機能 | データ可視化コンポーネント | チャート・グラフ、統計データ可視化 | 8時間 | 🟡 |
| ◻️ | Frontend | FE-08-ui-02 | 🎨 高度なUI機能 | アニメーション・トランジション | ページ遷移、カード表示、モーダルアニメーション | 6時間 | 🟢 |
| ◻️ | Frontend | FE-08-ui-03 | 🎨 高度なUI機能 | リアルタイム機能 | WebSocket接続、リアルタイム通知、即座更新 | 8時間 | 🟢 |
| ◻️ | Frontend | FE-09-mobile-01 | 📱 レスポンシブ・PWA | レスポンシブデザイン最適化 | 全画面のモバイル最適化、タッチ操作改善 | 10時間 | 🟡 |
| ◻️ | Frontend | FE-09-mobile-02 | 📱 レスポンシブ・PWA | PWA設定 | Service Worker、マニフェスト、オフライン対応 | 6時間 | 🟢 |
| ◻️ | Frontend | FE-09-mobile-03 | 📱 レスポンシブ・PWA | プッシュ通知対応 | ブラウザプッシュ通知、権限管理、設定UI | 8時間 | 🟢 |
| ◻️ | Frontend | FE-10-access-01 | ♿ アクセシビリティ | WCAG 2.1 AA準拠 | アクセシビリティガイドライン準拠、色覚対応 | 8時間 | 🟢 |
| ◻️ | Frontend | FE-10-access-02 | ♿ アクセシビリティ | スクリーンリーダー対応 | ARIA属性、セマンティックHTML、音声読み上げ対応 | 6時間 | 🟢 |
| ◻️ | Frontend | FE-10-access-03 | ♿ アクセシビリティ | キーボードナビゲーション | Tab順序最適化、ショートカットキー対応 | 4時間 | 🟢 |
| ◻️ | Frontend | FE-11-test-01 | 🧪 テスト実装 | 単体テスト（Features） | features層のテスト、カスタムフックテスト | 12時間 | 🟡 |
| ◻️ | Frontend | FE-11-test-02 | 🧪 テスト実装 | 単体テスト（Widgets） | widgets層のテスト、UIコンポーネントテスト | 10時間 | 🟡 |
| ◻️ | Frontend | FE-11-test-03 | 🧪 テスト実装 | 統合テスト（Pages） | page-components層のテスト、ページ統合テスト | 8時間 | 🟡 |
| ◻️ | Frontend | FE-11-test-04 | 🧪 テスト実装 | E2Eテスト | Playwright、主要フローのテスト | 10時間 | 🟢 |
| ◻️ | Frontend | FE-12-perf-01 | 🚀 パフォーマンス最適化 | コード分割・遅延読み込み | 動的インポート、ページ分割、バンドル最適化 | 6時間 | 🟢 |
| ◻️ | Frontend | FE-12-perf-02 | 🚀 パフォーマンス最適化 | 画像最適化 | Next.js Image、WebP対応、レスポンシブ画像 | 4時間 | 🟢 |
| ◻️ | Frontend | FE-12-perf-03 | 🚀 パフォーマンス最適化 | キャッシュ戦略最適化 | React Query設定調整、キャッシュ最適化 | 4時間 | 🟢 |
| ◻️ | Frontend | FE-12-perf-04 | 🚀 パフォーマンス最適化 | Core Web Vitals最適化 | LCP、FID、CLS改善、パフォーマンス監視 | 6時間 | 🟢 |
| ◻️ | Frontend | FE-13-security-01 | 🔒 セキュリティ対応 | XSS対策 | サニタイゼーション、Content Security Policy設定 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-13-security-02 | 🔒 セキュリティ対応 | 認証セキュリティ強化 | トークン管理、セッション管理、CSRF対策 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-13-security-03 | 🔒 セキュリティ対応 | 入力検証強化 | フロントエンド検証、サニタイゼーション、型安全性 | 4時間 | 🟡 |
| ◻️ | Frontend | FE-14-docs-01 | 📝 ドキュメント作成 | コンポーネントドキュメント | Storybook設定、コンポーネントカタログ作成 | 8時間 | 🟢 |
| ◻️ | Frontend | FE-14-docs-02 | 📝 ドキュメント作成 | FSD開発ガイド作成 | FSD構造説明、開発フロー、ベストプラクティス | 6時間 | 🟢 |

---

## 📊 統計情報
- **総タスク数**: 227タスク
  - **バックエンドタスク**: 58タスク
  - **インフラタスク**: 50タスク
  - **フロントエンドタスク**: 119タスク
- **総見積工数**: 1048時間
  - **バックエンド工数**: 295時間
  - **インフラ工数**: 268時間
  - **フロントエンド工数**: 485時間
- **想定開発期間**: 約16-20週間（3チーム並行）

## 🔄 ステータス説明
- **◻️**: 未着手
- **🌀**: 進行中
- **✅**: 完了

## 🎯 優先度説明
- **🔴**: 高（最優先）
- **🟡**: 中（普通）
- **🟢**: 低（後回しOK）

## 👥 チーム説明
- **Backend**: バックエンドAPI・Discord Bot開発
- **Infrastructure**: インフラ・DevOps・プロジェクト管理
- **Frontend**: フロントエンドUI・UX開発（FSD準拠）

---

_作成日: 2025-01-21 (Notion インポート用)_
