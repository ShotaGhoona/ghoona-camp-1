# Ghoona Camp フロントエンド開発タスクチェックリスト（FSD準拠）

## プロジェクト概要
朝活コミュニティアプリ「Ghoona Camp」のフロントエンド開発におけるタスク一覧です。  
FSD（Feature-Sliced Design）アーキテクチャとNext.js 15のApp Routerを使用して実装します。

---

## 📋 開発環境セットアップ

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ✅ | FE-01-setup-01 | 開発環境の構築 | Node.js、Next.js 15、TypeScript、必要な開発ツールのセットアップ | 3時間 | 🔴 |
| ◻️ | FE-01-setup-02 | プロジェクト初期化とFSD構造作成 | Next.js初期化、FSDディレクトリ構造の作成、基本設定ファイル | 4時間 | 🔴 |
| ◻️ | FE-01-setup-03 | 開発ツール・Linter設定 | ESLint、Prettier、FSD boundaries、pre-commit hooks設定 | 2時間 | 🔴 |

## 🔧 Shared層（共有基盤）

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-02-shared-01 | shadcn/ui セットアップ | shadcn/uiライブラリ導入、基本コンポーネントの設定 | 3時間 | 🔴 |
| ◻️ | FE-02-shared-02 | デザインシステム構築 | カラーパレット、タイポグラフィ、スペーシング、テーマ設定 | 6時間 | 🔴 |
| ◻️ | FE-02-shared-03 | 共通UIコンポーネント | Button、Card、Modal、Form、Loading等の基本コンポーネント | 8時間 | 🔴 |
| ◻️ | FE-02-shared-04 | APIクライアント設定 | Supabaseクライアント、Clerkクライアント、HTTP設定 | 4時間 | 🔴 |
| ◻️ | FE-02-shared-05 | 共通フック・ユーティリティ | 共通カスタムフック、日付・フォーマット・バリデーション関数 | 5時間 | 🔴 |
| ◻️ | FE-02-shared-06 | エラーハンドリング基盤 | Error Boundary、Toast通知、エラーページ実装 | 4時間 | 🔴 |

## 🏗️ Entities層（ドメインエンティティ）

### User Entity

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-03-user-01 | User型定義・モデル | ユーザー関連の型定義、インターフェース、バリデーションスキーマ | 3時間 | 🔴 |
| ◻️ | FE-03-user-02 | User API クライアント | ユーザー情報取得・更新、プロフィール管理のAPI呼び出し | 4時間 | 🔴 |
| ◻️ | FE-03-user-03 | User クエリキー管理 | React Queryのクエリキー一元管理、無効化戦略 | 2時間 | 🔴 |

### Attendance Entity

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-03-attend-01 | Attendance型定義・モデル | 出席ログ、統計、ランキング関連の型定義とバリデーション | 4時間 | 🔴 |
| ◻️ | FE-03-attend-02 | Attendance API クライアント | 出席ログ取得、統計取得、ランキング取得のAPI呼び出し | 5時間 | 🔴 |
| ◻️ | FE-03-attend-03 | Attendance クエリキー管理 | 出席関連データのクエリキー管理、リアルタイム更新 | 2時間 | 🔴 |

### Goal Entity

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-03-goal-01 | Goal型定義・モデル | 目標、進捗関連の型定義とバリデーションスキーマ | 3時間 | 🟡 |
| ◻️ | FE-03-goal-02 | Goal API クライアント | 目標CRUD、進捗管理のAPI呼び出し | 4時間 | 🟡 |
| ◻️ | FE-03-goal-03 | Goal クエリキー管理 | 目標関連データのクエリキー管理 | 2時間 | 🟡 |

### Event Entity

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-03-event-01 | Event型定義・モデル | イベント、参加者関連の型定義とバリデーション | 3時間 | 🟡 |
| ◻️ | FE-03-event-02 | Event API クライアント | イベントCRUD、参加管理のAPI呼び出し | 4時間 | 🟡 |
| ◻️ | FE-03-event-03 | Event クエリキー管理 | イベント関連データのクエリキー管理 | 2時間 | 🟡 |

### Title Entity

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-03-title-01 | Title型定義・モデル | 称号、実績関連の型定義とバリデーション | 3時間 | 🟡 |
| ◻️ | FE-03-title-02 | Title API クライアント | 称号取得、実績管理のAPI呼び出し | 3時間 | 🟡 |
| ◻️ | FE-03-title-03 | Title クエリキー管理 | 称号関連データのクエリキー管理 | 2時間 | 🟡 |

### Notification Entity

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-03-notify-01 | Notification型定義・モデル | 通知、設定関連の型定義とバリデーション | 3時間 | 🟡 |
| ◻️ | FE-03-notify-02 | Notification API クライアント | 通知CRUD、設定管理のAPI呼び出し | 3時間 | 🟡 |
| ◻️ | FE-03-notify-03 | Notification クエリキー管理 | 通知関連データのクエリキー管理 | 2時間 | 🟡 |

## ⚙️ Features層（機能実装）

### Attendance Features

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-04-attend-01 | attendance-log-get feature | 出席ログ取得機能、カスタムフック、エラーハンドリング | 4時間 | 🔴 |
| ◻️ | FE-04-attend-02 | attendance-stats-get feature | 出席統計取得機能、連続日数計算、進捗表示 | 5時間 | 🔴 |
| ◻️ | FE-04-attend-03 | attendance-ranking-get feature | ランキング取得機能、フィルタリング、ソート | 4時間 | 🔴 |
| ◻️ | FE-04-attend-04 | attendance-streak-get feature | 連続記録取得機能、アチーブメント判定 | 3時間 | 🟡 |

### User Features

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-04-user-01 | user-profile-get feature | ユーザープロフィール取得機能、メタデータ管理 | 3時間 | 🔴 |
| ◻️ | FE-04-user-02 | user-profile-update feature | プロフィール更新機能、画像アップロード、バリデーション | 5時間 | 🔴 |
| ◻️ | FE-04-user-03 | user-rivals-manage feature | ライバル管理機能、追加・削除、最大3人制限 | 4時間 | 🟡 |
| ◻️ | FE-04-user-04 | user-social-links feature | ソーシャルリンク管理機能、CRUD操作 | 3時間 | 🟡 |

### Goal Features

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-04-goal-01 | goal-create feature | 目標作成機能、フォーム、バリデーション | 4時間 | 🟡 |
| ◻️ | FE-04-goal-02 | goal-update feature | 目標更新機能、進捗管理、公開設定 | 4時間 | 🟡 |
| ◻️ | FE-04-goal-03 | goal-delete feature | 目標削除機能、確認ダイアログ、データ整合性 | 2時間 | 🟡 |
| ◻️ | FE-04-goal-04 | goal-list-get feature | 目標一覧取得機能、フィルタリング、ソート | 3時間 | 🟡 |

### Event Features

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-04-event-01 | event-create feature | イベント作成機能、日時選択、参加者設定 | 5時間 | 🟡 |
| ◻️ | FE-04-event-02 | event-participate feature | イベント参加機能、申込・キャンセル、定員管理 | 4時間 | 🟡 |
| ◻️ | FE-04-event-03 | event-list-get feature | イベント一覧取得機能、検索・フィルター | 3時間 | 🟡 |
| ◻️ | FE-04-event-04 | event-update feature | イベント更新機能、作成者権限チェック | 3時間 | 🟡 |

### Title Features

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-04-title-01 | title-get feature | 全称号取得機能、獲得条件表示 | 3時間 | 🟡 |
| ◻️ | FE-04-title-02 | achievement-get feature | 獲得履歴取得機能、現在設定中の称号 | 3時間 | 🟡 |
| ◻️ | FE-04-title-03 | title-change feature | 称号変更機能、獲得済み称号から選択 | 2時間 | 🟡 |

### Notification Features

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-04-notify-01 | notification-get feature | 通知取得機能、未読管理、ページネーション | 4時間 | 🟡 |
| ◻️ | FE-04-notify-02 | notification-read feature | 既読処理機能、一括既読、自動既読 | 3時間 | 🟡 |
| ◻️ | FE-04-notify-03 | notification-settings feature | 通知設定機能、カテゴリ別ON/OFF、時刻設定 | 4時間 | 🟡 |

### Auth Features

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-04-auth-01 | clerk-auth-integration | Clerk認証統合、ログイン・サインアップフロー | 5時間 | 🔴 |
| ◻️ | FE-04-auth-02 | auth-guard feature | 認証ガード機能、ルート保護、リダイレクト | 3時間 | 🔴 |
| ◻️ | FE-04-auth-03 | auth-state-manage | 認証状態管理、React Query統合、Zustand連携 | 4時間 | 🔴 |

## 🧩 Widgets層（複合UIブロック）

### Attendance Widgets

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-05-attend-01 | AttendanceCard widget | 出席カード表示、統計サマリー、連続日数 | 4時間 | 🔴 |
| ◻️ | FE-05-attend-02 | StreakBadge widget | 連続日数バッジ、レベル表示、アニメーション | 3時間 | 🔴 |
| ◻️ | FE-05-attend-03 | AttendanceChart widget | 出席グラフ、月次表示、データ可視化 | 6時間 | 🟡 |
| ◻️ | FE-05-attend-04 | CalendarView widget | カレンダー表示、出席状況、日付選択 | 8時間 | 🟡 |
| ◻️ | FE-05-attend-05 | RankingList widget | ランキング一覧、ライバル強調、フィルタリング | 5時間 | 🟡 |

### Goal Widgets

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-05-goal-01 | GoalCard widget | 目標カード表示、進捗バー、期限表示 | 4時間 | 🟡 |
| ◻️ | FE-05-goal-02 | ProgressBar widget | 進捗バー、パーセンテージ、アニメーション | 3時間 | 🟡 |
| ◻️ | FE-05-goal-03 | GoalForm widget | 目標フォーム、バリデーション、公開設定 | 5時間 | 🟡 |
| ◻️ | FE-05-goal-04 | GoalList widget | 目標一覧表示、フィルタリング、ソート | 4時間 | 🟡 |

### Event Widgets

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-05-event-01 | EventCard widget | イベントカード、参加状況、詳細リンク | 4時間 | 🟡 |
| ◻️ | FE-05-event-02 | EventList widget | イベント一覧、検索・フィルター、ページネーション | 5時間 | 🟡 |
| ◻️ | FE-05-event-03 | ParticipantList widget | 参加者一覧、アバター表示、参加状況 | 3時間 | 🟡 |
| ◻️ | FE-05-event-04 | EventForm widget | イベントフォーム、日時選択、画像アップロード | 6時間 | 🟡 |

### Title Widgets

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-05-title-01 | TitleCard widget | 称号カード、詳細情報、ストーリー | 4時間 | 🟡 |
| ◻️ | FE-05-title-02 | TitleBadge widget | 称号バッジ、レベル表示、テーマカラー | 3時間 | 🟡 |
| ◻️ | FE-05-title-03 | AchievementModal widget | 称号獲得モーダル、アニメーション、お祝い演出 | 5時間 | 🟢 |
| ◻️ | FE-05-title-04 | TitleGrid widget | 称号一覧グリッド、カテゴリ別表示、進捗表示 | 4時間 | 🟡 |

### Notification Widgets

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-05-notify-01 | NotificationList widget | 通知一覧、未読表示、アクション | 4時間 | 🟡 |
| ◻️ | FE-05-notify-02 | NotificationBadge widget | 通知バッジ、未読件数、リアルタイム更新 | 3時間 | 🟡 |
| ◻️ | FE-05-notify-03 | NotificationSettings widget | 通知設定フォーム、カテゴリ別設定 | 4時間 | 🟡 |

### Layout Widgets

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-05-layout-01 | Header widget | グローバルヘッダー、ユーザーメニュー、ナビゲーション | 5時間 | 🔴 |
| ◻️ | FE-05-layout-02 | Sidebar widget | サイドナビゲーション、メニュー状態管理、モバイル対応 | 5時間 | 🔴 |
| ◻️ | FE-05-layout-03 | Footer widget | フッター、リンク集、会社情報 | 2時間 | 🟡 |

## 📄 Page-components層（ページ統合）

### Dashboard Pages

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-06-dash-01 | Dashboard page component | ダッシュボードページ統合、widgetsオーケストレーション | 4時間 | 🔴 |
| ◻️ | FE-06-dash-02 | Dashboard hooks | ダッシュボード専用フック、データ統合、状態管理 | 3時間 | 🔴 |

### Attendance Pages

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-06-attend-01 | Attendance page component | 出席記録ページ統合、カレンダー・統計表示 | 4時間 | 🔴 |
| ◻️ | FE-06-attend-02 | Attendance hooks | 出席ページ専用フック、データフィルタリング | 3時間 | 🔴 |

### Goal Pages

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-06-goal-01 | Goals page component | 目標一覧ページ統合、CRUD操作統合 | 4時間 | 🟡 |
| ◻️ | FE-06-goal-02 | Goal detail page component | 目標詳細ページ、編集・削除機能統合 | 3時間 | 🟡 |
| ◻️ | FE-06-goal-03 | Goals page hooks | 目標ページ専用フック、フィルタリング・ソート | 3時間 | 🟡 |

### Event Pages

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-06-event-01 | Events page component | イベント一覧ページ統合、検索・フィルター統合 | 4時間 | 🟡 |
| ◻️ | FE-06-event-02 | Event detail page component | イベント詳細ページ、参加管理統合 | 3時間 | 🟡 |
| ◻️ | FE-06-event-03 | Events page hooks | イベントページ専用フック、参加状況管理 | 3時間 | 🟡 |

### Title Pages

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-06-title-01 | Titles page component | 称号一覧ページ統合、獲得状況表示 | 3時間 | 🟡 |
| ◻️ | FE-06-title-02 | Titles page hooks | 称号ページ専用フック、進捗計算 | 2時間 | 🟡 |

### Settings Pages

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-06-settings-01 | Profile settings page component | プロフィール設定ページ統合、フォーム管理 | 4時間 | 🔴 |
| ◻️ | FE-06-settings-02 | Notification settings page component | 通知設定ページ統合、設定管理 | 3時間 | 🟡 |
| ◻️ | FE-06-settings-03 | Account settings page component | アカウント設定ページ統合、Clerk連携 | 3時間 | 🟡 |

### Landing Pages

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-06-landing-01 | Landing page component | ランディングページ、未認証ユーザー向け | 5時間 | 🟡 |
| ◻️ | FE-06-landing-02 | Landing page hooks | ランディング専用フック、アニメーション管理 | 3時間 | 🟡 |

## 🔗 App層（Next.js統合）

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-07-app-01 | Root Layout設定 | Root Layout、グローバルプロバイダー設定 | 3時間 | 🔴 |
| ◻️ | FE-07-app-02 | ルート定義 | 全ページのルート定義、page.tsx、layout.tsx作成 | 4時間 | 🔴 |
| ◻️ | FE-07-app-03 | Loading・Error UI | Loading UI、Error UI、Not Found実装 | 3時間 | 🔴 |
| ◻️ | FE-07-app-04 | Providers設定 | Clerk、React Query、Zustand等のプロバイダー統合 | 4時間 | 🔴 |
| ◻️ | FE-07-app-05 | Metadata設定 | SEO・メタタグ最適化、OGP設定 | 3時間 | 🟡 |

## 🎨 高度なUI機能

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-08-ui-01 | データ可視化コンポーネント | チャート・グラフ、統計データ可視化 | 8時間 | 🟡 |
| ◻️ | FE-08-ui-02 | アニメーション・トランジション | ページ遷移、カード表示、モーダルアニメーション | 6時間 | 🟢 |
| ◻️ | FE-08-ui-03 | リアルタイム機能 | WebSocket接続、リアルタイム通知、即座更新 | 8時間 | 🟢 |

## 📱 レスポンシブ・PWA

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-09-mobile-01 | レスポンシブデザイン最適化 | 全画面のモバイル最適化、タッチ操作改善 | 10時間 | 🟡 |
| ◻️ | FE-09-mobile-02 | PWA設定 | Service Worker、マニフェスト、オフライン対応 | 6時間 | 🟢 |
| ◻️ | FE-09-mobile-03 | プッシュ通知対応 | ブラウザプッシュ通知、権限管理、設定UI | 8時間 | 🟢 |

## ♿ アクセシビリティ

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-10-access-01 | WCAG 2.1 AA準拠 | アクセシビリティガイドライン準拠、色覚対応 | 8時間 | 🟢 |
| ◻️ | FE-10-access-02 | スクリーンリーダー対応 | ARIA属性、セマンティックHTML、音声読み上げ対応 | 6時間 | 🟢 |
| ◻️ | FE-10-access-03 | キーボードナビゲーション | Tab順序最適化、ショートカットキー対応 | 4時間 | 🟢 |

## 🧪 テスト実装

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-11-test-01 | 単体テスト（Features） | features層のテスト、カスタムフックテスト | 12時間 | 🟡 |
| ◻️ | FE-11-test-02 | 単体テスト（Widgets） | widgets層のテスト、UIコンポーネントテスト | 10時間 | 🟡 |
| ◻️ | FE-11-test-03 | 統合テスト（Pages） | page-components層のテスト、ページ統合テスト | 8時間 | 🟡 |
| ◻️ | FE-11-test-04 | E2Eテスト | Playwright、主要フローのテスト | 10時間 | 🟢 |

## 🚀 パフォーマンス最適化

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-12-perf-01 | コード分割・遅延読み込み | 動的インポート、ページ分割、バンドル最適化 | 6時間 | 🟢 |
| ◻️ | FE-12-perf-02 | 画像最適化 | Next.js Image、WebP対応、レスポンシブ画像 | 4時間 | 🟢 |
| ◻️ | FE-12-perf-03 | キャッシュ戦略最適化 | React Query設定調整、キャッシュ最適化 | 4時間 | 🟢 |
| ◻️ | FE-12-perf-04 | Core Web Vitals最適化 | LCP、FID、CLS改善、パフォーマンス監視 | 6時間 | 🟢 |

## 🔒 セキュリティ対応

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-13-security-01 | XSS対策 | サニタイゼーション、Content Security Policy設定 | 4時間 | 🟡 |
| ◻️ | FE-13-security-02 | 認証セキュリティ強化 | トークン管理、セッション管理、CSRF対策 | 4時間 | 🟡 |
| ◻️ | FE-13-security-03 | 入力検証強化 | フロントエンド検証、サニタイゼーション、型安全性 | 4時間 | 🟡 |

## 📝 ドキュメント作成

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-14-docs-01 | コンポーネントドキュメント | Storybook設定、コンポーネントカタログ作成 | 8時間 | 🟢 |
| ◻️ | FE-14-docs-02 | FSD開発ガイド作成 | FSD構造説明、開発フロー、ベストプラクティス | 6時間 | 🟢 |

---

## 📊 総工数見積
- **総タスク数**: 114タスク
- **総見積工数**: 485時間
- **想定開発期間**: 約14-16週間（1人）

## 🔄 ステータス説明
- **◻️**: 未着手
- **🌀**: 進行中
- **✅**: 完了

## 🎯 優先度説明
- **🔴**: 高（最優先）
- **🟡**: 中（普通）
- **🟢**: 低（後回しOK）

## 🎯 FSDマイルストーン
1. **基盤構築（Shared層）**: FE-01-setup-01〜FE-02-shared-06（約3週間）
2. **エンティティ層**: FE-03-user-01〜FE-03-notify-03（約4週間）
3. **機能層（Features）**: FE-04-attend-01〜FE-04-auth-03（約5週間）
4. **ウィジェット層**: FE-05-attend-01〜FE-05-layout-03（約4週間）
5. **ページ統合**: FE-06-dash-01〜FE-07-app-05（約3週間）
6. **品質向上・最適化**: FE-08-ui-01〜FE-14-docs-02（約3週間）

## 🏗️ FSD依存関係
```
page-components ← widgets ← features ← entities ← shared
```

各レイヤーは下位レイヤーのみに依存し、同一レイヤー間の依存は禁止

---

_最終更新: 2025-01-21_