# Ghoona Camp フロントエンド開発タスクチェックリスト（FSD準拠）

## プロジェクト概要
朝活コミュニティアプリ「Ghoona Camp」のフロントエンド開発におけるタスク一覧です。  
FSD（Feature-Sliced Design）アーキテクチャとNext.js 15のApp Routerを使用して実装します。

---

## 01. 📋 開発環境セットアップ

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ✅ | FE-01-setup-01 | 開発環境の構築 | Node.js、Next.js 15、TypeScript、必要な開発ツールのセットアップ | 3時間 | 🔴 |
| ✅ | FE-01-setup-02 | プロジェクト初期化とFSD構造作成 | Next.js初期化、FSDディレクトリ構造の作成、基本設定ファイル | 4時間 | 🔴 |
| ✅ | FE-01-setup-03 | 開発ツール・Linter設定 | ESLint、Prettier、FSD boundaries、pre-commit hooks設定 | 2時間 | 🔴 |

## 02. 🔧 Shared層（共有基盤）

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ✅ | FE-02-shared-01 | shadcn/ui セットアップ | shadcn/uiライブラリ導入、基本コンポーネントの設定 | 3時間 | 🔴 |
| ◻️ | FE-02-shared-02 | デザインシステム構築 | カラーパレット、タイポグラフィ、スペーシング、テーマ設定 | 6時間 | 🔴 |
| ◻️ | FE-02-shared-03 | 共通UIコンポーネント | Button、Card、Modal、Form、Loading等の基本コンポーネント | 8時間 | 🔴 |
| ◻️ | FE-02-shared-04 | APIクライアント設定 | Supabaseクライアント、Clerkクライアント、HTTP設定 | 4時間 | 🔴 |
| ◻️ | FE-02-shared-05 | 共通フック・ユーティリティ | 共通カスタムフック、日付・フォーマット・バリデーション関数 | 5時間 | 🔴 |
| ◻️ | FE-02-shared-06 | エラーハンドリング基盤 | Error Boundary、Toast通知、エラーページ実装 | 4時間 | 🔴 |

## 03. 🏗️ Entities層（ドメインエンティティ）

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ✅ | FE-03-entity-01 | User Entity 完全実装 | User型定義、API、クエリキー、Metadata、Social Links、Rivals | 18時間 | 🔴 |
| ✅ | FE-03-entity-02 | Title Entity 完全実装 | Title型定義、API、クエリキー、Achievements | 12時間 | 🔴 |
| ◻️ | FE-03-entity-03 | Attendance Entity 完全実装 | Attendance型定義、API、クエリキー、統計・ランキング | 11時間 | 🔴 |
| ◻️ | FE-03-entity-04 | Goal Entity 完全実装 | Goal型定義、API、クエリキー、進捗管理 | 9時間 | 🟡 |
| ◻️ | FE-03-entity-05 | Event Entity 完全実装 | Event型定義、API、クエリキー、参加者管理 | 9時間 | 🟡 |
| ◻️ | FE-03-entity-06 | Notification Entity 完全実装 | Notification型定義、API、クエリキー、設定管理 | 8時間 | 🟡 |

## 04. ⚙️ Features層（機能実装）

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-04-user | User Features 完全実装 | プロフィール取得・更新、ライバル管理、ソーシャルリンク | 15時間 | 🔴 |
| ◻️ | FE-04-title | Title Features 完全実装 | 称号取得、実績管理、称号変更機能 | 8時間 | 🟡 |
| ◻️ | FE-04-attendance | Attendance Features 完全実装 | 出席ログ・統計・ランキング取得機能 | 13時間 | 🔴 |
| ◻️ | FE-04-goal | Goal Features 完全実装 | 目標CRUD、進捗管理、一覧表示機能 | 13時間 | 🟡 |
| ◻️ | FE-04-event | Event Features 完全実装 | イベントCRUD、参加管理、検索機能 | 15時間 | 🟡 |
| ◻️ | FE-04-notification | Notification Features 完全実装 | 通知取得・既読処理・設定管理機能 | 11時間 | 🟡 |
| ◻️ | FE-04-auth | Auth Features 完全実装 | Clerk認証統合、ガード機能、状態管理 | 12時間 | 🔴 |

## 05. 🧩 Widgets層（複合UIブロック）

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-05-layout | Layout Widgets 完全実装 | Header、Sidebar、Footer widget | 12時間 | 🔴 |
| ◻️ | FE-05-user | User Widgets 完全実装 | UserCard、ProfileForm、RivalsList widget | 12時間 | 🔴 |
| ◻️ | FE-05-attendance | Attendance Widgets 完全実装 | AttendanceCard、StreakBadge、Chart、Calendar、Ranking widget | 26時間 | 🔴 |
| ◻️ | FE-05-title | Title Widgets 完全実装 | TitleCard、TitleBadge、AchievementModal、TitleGrid widget | 16時間 | 🟡 |
| ◻️ | FE-05-goal | Goal Widgets 完全実装 | GoalCard、ProgressBar、GoalForm、GoalList widget | 16時間 | 🟡 |
| ◻️ | FE-05-event | Event Widgets 完全実装 | EventCard、EventList、ParticipantList、EventForm widget | 18時間 | 🟡 |
| ◻️ | FE-05-notification | Notification Widgets 完全実装 | NotificationList、NotificationBadge、NotificationSettings widget | 11時間 | 🟡 |

## 06. 📄 Page-components層（ページ統合）

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-06-auth | Authentication Pages 完全実装 | Landing、Sign In、Sign Up page component | 11時間 | 🔴 |
| ◻️ | FE-06-main | Main Application Pages 完全実装 | Dashboard、Ranking、Activity、Events、Goals、Titles、Members page | 27時間 | 🔴 |
| ◻️ | FE-06-settings | Settings Pages 完全実装 | Profile、Vision、Notification、Account settings page | 13時間 | 🔴 |
| ◻️ | FE-06-error | Error Pages 完全実装 | 404、500 error page component | 4時間 | 🟡 |

## 07. 🔗 App層（Next.js統合）

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-07-app-01 | Root Layout設定 | Root Layout、グローバルプロバイダー設定 | 3時間 | 🔴 |
| ◻️ | FE-07-app-02 | ルート定義 | 全ページのルート定義、page.tsx、layout.tsx作成 | 4時間 | 🔴 |
| ◻️ | FE-07-app-03 | Loading・Error UI | Loading UI、Error UI、Not Found実装 | 3時間 | 🔴 |
| ◻️ | FE-07-app-04 | Providers設定 | Clerk、React Query、Zustand等のプロバイダー統合 | 4時間 | 🔴 |
| ◻️ | FE-07-app-05 | Metadata設定 | SEO・メタタグ最適化、OGP設定 | 3時間 | 🟡 |

## 08. 🎨 高度なUI機能

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-08-ui-01 | データ可視化コンポーネント | チャート・グラフ、統計データ可視化 | 8時間 | 🟡 |
| ◻️ | FE-08-ui-02 | アニメーション・トランジション | ページ遷移、カード表示、モーダルアニメーション | 6時間 | 🟢 |
| ◻️ | FE-08-ui-03 | リアルタイム機能 | WebSocket接続、リアルタイム通知、即座更新 | 8時間 | 🟢 |

## 09. 📱 レスポンシブ・PWA

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-09-mobile-01 | レスポンシブデザイン最適化 | 全画面のモバイル最適化、タッチ操作改善 | 10時間 | 🟡 |
| ◻️ | FE-09-mobile-02 | PWA設定 | Service Worker、マニフェスト、オフライン対応 | 6時間 | 🟢 |
| ◻️ | FE-09-mobile-03 | プッシュ通知対応 | ブラウザプッシュ通知、権限管理、設定UI | 8時間 | 🟢 |

## 10. ♿ アクセシビリティ

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-10-access-01 | WCAG 2.1 AA準拠 | アクセシビリティガイドライン準拠、色覚対応 | 8時間 | 🟢 |
| ◻️ | FE-10-access-02 | スクリーンリーダー対応 | ARIA属性、セマンティックHTML、音声読み上げ対応 | 6時間 | 🟢 |
| ◻️ | FE-10-access-03 | キーボードナビゲーション | Tab順序最適化、ショートカットキー対応 | 4時間 | 🟢 |

## 11. 🧪 テスト実装

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-11-test-01 | 単体テスト（Features） | features層のテスト、カスタムフックテスト | 12時間 | 🟡 |
| ◻️ | FE-11-test-02 | 単体テスト（Widgets） | widgets層のテスト、UIコンポーネントテスト | 10時間 | 🟡 |
| ◻️ | FE-11-test-03 | 統合テスト（Pages） | page-components層のテスト、ページ統合テスト | 8時間 | 🟡 |
| ◻️ | FE-11-test-04 | E2Eテスト | Playwright、主要フローのテスト | 10時間 | 🟢 |

## 12. 🚀 パフォーマンス最適化

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-12-perf-01 | コード分割・遅延読み込み | 動的インポート、ページ分割、バンドル最適化 | 6時間 | 🟢 |
| ◻️ | FE-12-perf-02 | 画像最適化 | Next.js Image、WebP対応、レスポンシブ画像 | 4時間 | 🟢 |
| ◻️ | FE-12-perf-03 | キャッシュ戦略最適化 | React Query設定調整、キャッシュ最適化 | 4時間 | 🟢 |
| ◻️ | FE-12-perf-04 | Core Web Vitals最適化 | LCP、FID、CLS改善、パフォーマンス監視 | 6時間 | 🟢 |

## 13. 🔒 セキュリティ対応

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-13-security-01 | XSS対策 | サニタイゼーション、Content Security Policy設定 | 4時間 | 🟡 |
| ◻️ | FE-13-security-02 | 認証セキュリティ強化 | トークン管理、セッション管理、CSRF対策 | 4時間 | 🟡 |
| ◻️ | FE-13-security-03 | 入力検証強化 | フロントエンド検証、サニタイゼーション、型安全性 | 4時間 | 🟡 |

## 14. 📝 ドキュメント作成

| ステータス | ID | タスク名 | 内容 | 工数見積 | 優先度 |
|---|---|---|---|---|---|
| ◻️ | FE-14-docs-01 | コンポーネントドキュメント | Storybook設定、コンポーネントカタログ作成 | 8時間 | 🟢 |
| ◻️ | FE-14-docs-02 | FSD開発ガイド作成 | FSD構造説明、開発フロー、ベストプラクティス | 6時間 | 🟢 |

---

## 📊 総工数見積
- **総タスク数**: 56タスク
- **総見積工数**: 540時間
- **想定開発期間**: 約16-18週間（1人）

## 🔄 ステータス説明
- **◻️**: 未着手
- **🌀**: 進行中
- **✅**: 完了

## 🎯 優先度説明
- **🔴**: 高（最優先）
- **🟡**: 中（普通）
- **🟢**: 低（後回しOK）

## 🎯 FSDマイルストーン
1. **基盤構築（01-02）**: セットアップ〜Shared層（約4週間）
2. **エンティティ層（03）**: 全エンティティ実装（約4週間）
3. **機能層（04）**: Features層実装（約5週間）
4. **ウィジェット層（05）**: Widgets層実装（約4週間）
5. **ページ統合（06-07）**: Page-components〜App層（約3週間）
6. **品質向上・最適化（08-14）**: 高度機能〜ドキュメント（約4週間）

## 🏗️ FSD依存関係
```
page-components ← widgets ← features ← entities ← shared
```

各レイヤーは下位レイヤーのみに依存し、同一レイヤー間の依存は禁止

---

_最終更新: 2025-01-28_