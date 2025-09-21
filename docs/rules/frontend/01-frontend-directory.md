# Ghoona Camp FSD (Feature-Sliced Design) ディレクトリ構成ガイド

## 概要

朝活コミュニティアプリGhoona CampにおけるFSD（Feature-Sliced Design）とNext.js 15のApp Routerを統合する際の実装パターンとベストプラクティスを示します。

## ディレクトリ構造の統合

### Next.js App Router構造とFSDの共存

```
frontend/
└── src/                         # すべてのソースコード
    ├── app/                     # Next.js App Router + FSDアプリケーション層
    │   ├── dashboard/           # ダッシュボードルート
    │   ├── goals/               # 目標管理ルート
    │   ├── events/              # イベント管理ルート
    │   ├── titles/              # 称号管理ルート
    │   ├── attendance/          # 出席記録ルート
    │   ├── notifications/       # 通知管理ルート
    │   ├── settings/            # 設定ルート
    │   ├── providers/           # グローバルプロバイダー（Clerk、React Queryなど）
    │   ├── store/               # 状態管理設定
    │   ├── layout.tsx           # Root Layout
    │   ├── page.tsx             # Home Page
    │   └── globals.css
    │
    ├── page-components/         # ページ固有のオーケストレーション層
    ├── widgets/                 # 複合UIコンポーネント層
    ├── features/                # 機能固有のビジネスロジック層
    ├── entities/                # ビジネスエンティティ層
    └── shared/                  # 共有リソース層
```

## FSDレイヤーの責務分担

### `src/app/` (Next.js App Router + FSD Application Layer)

- **責務**: ルーティング、メタデータ、レイアウト定義、グローバル設定
- **含むもの**:
  - ルート定義: `page.tsx`, `layout.tsx`, `loading.tsx`, `error.tsx`
  - グローバル設定: `providers/` (Clerk認証、React Query等), `store/`
- **含まないもの**: ビジネスロジック、複雑な状態管理、再利用可能なUIコンポーネント

### `src/page-components/`

- **責務**: ページ固有の状態管理とオーケストレーション
- **含むもの**:
  - ページコンテナコンポーネント
  - ページ固有のカスタムフック
  - ページレベルの状態管理
- **含まないもの**: 再利用可能なUIコンポーネント

### `src/widgets/`

- **責務**: 複数のfeatureを組み合わせた大きなUIブロック
- **含むもの**: ヘッダー、ドメイン固有のウィジェット（出席管理、目標管理、イベント管理、称号管理など）
- **含まないもの**: ページ固有のロジック、他のwidgets（同一レイヤー依存禁止）
- **設計原則**:
  - **Props駆動型**: 外部から動作を制御可能な設計
  - **ロジック分離**: UIロジックはカスタムフックに分離（`lib/`ディレクトリ）
  - **ドメイン別編成**: 機能ドメインごとにウィジェットを整理

### `src/features/`

- **責務**: 独立した機能単位のコンポーネント
- **含むもの**: 機能に必要なUI、ロジック、API呼び出し
- **含まないもの**: ページ全体の状態管理
- **組織化**: エンティティごと、操作ごとに細分化（例: `attendance/log/log-get/`）
- **命名規則**: `use[Entity][Action].ts`（例: `useGetAttendance.ts`, `useCreateGoal.ts`）

### `src/entities/`

- **責務**: ビジネスエンティティの定義とAPI
- **含むもの**:
  - APIクライアント
  - 型定義とバリデーター
  - クエリキーの一元管理（`utils/query-keys.ts`）
  - エンティティ固有のフック
- **含まないもの**: 複雑なUIコンポーネント

### `src/shared/`

- **責務**: アプリケーション全体で共有されるリソース
- **含むもの**:
  - UIコンポーネント（shadcn/ui）
  - APIクライアント設定（Supabase、Clerk連携）
  - 共通フック
  - ユーティリティ関数
  - アプリ設定
- **含まないもの**: ビジネスロジック

## 詳細なディレクトリ構造例

### FSDレイヤーを使用した構成

```
src/
├── app/                          # Next.js App Router + グローバル設定
│   ├── dashboard/                # ダッシュボードルート
│   │   └── page.tsx
│   ├── goals/                    # 目標管理ルート
│   │   ├── [id]/
│   │   │   └── page.tsx
│   │   └── page.tsx
│   ├── events/                   # イベント管理ルート
│   │   ├── [id]/
│   │   │   └── page.tsx
│   │   └── page.tsx
│   ├── titles/                   # 称号管理ルート
│   │   └── page.tsx
│   ├── attendance/               # 出席記録ルート
│   │   └── page.tsx
│   ├── notifications/            # 通知管理ルート
│   │   └── page.tsx
│   ├── settings/                 # 設定ルート
│   │   └── page.tsx
│   ├── providers/                # グローバルプロバイダー
│   │   ├── clerk-provider.tsx
│   │   ├── query-provider.tsx
│   │   └── index.tsx
│   ├── store/                    # 状態管理設定
│   │   └── index.ts
│   ├── layout.tsx
│   ├── page.tsx
│   └── globals.css
│
├── page-components/              # ページ固有のコンポーネント
│   ├── dashboard/                # ダッシュボードページ
│   │   ├── lib/
│   │   │   └── use-dashboard-page.ts
│   │   └── index.ts
│   ├── goals/                    # 目標一覧ページ
│   │   ├── lib/
│   │   │   └── use-goals-page.ts
│   │   └── index.ts
│   ├── goal-detail/              # 目標詳細ページ
│   │   ├── lib/
│   │   └── index.ts
│   ├── events/                   # イベント一覧ページ
│   │   ├── lib/
│   │   └── index.ts
│   ├── event-detail/             # イベント詳細ページ
│   │   ├── lib/
│   │   └── index.ts
│   ├── titles/                   # 称号管理ページ
│   │   ├── lib/
│   │   └── index.ts
│   ├── attendance/               # 出席記録ページ
│   │   ├── lib/
│   │   └── index.ts
│   └── settings/                 # 設定ページ
│       ├── lib/
│       │   └── use-settings-page.ts
│       └── index.ts
│
├── widgets/                      # 複合UIブロック
│   ├── header/                   # グローバルヘッダー
│   │   ├── ui/
│   │   ├── lib/
│   │   └── index.ts
│   ├── attendance/               # 出席管理関連ウィジェット
│   │   ├── attendanceCard/       # 出席カード
│   │   ├── streakBadge/          # 連続日数バッジ
│   │   ├── attendanceChart/      # 出席グラフ
│   │   └── calendarView/         # カレンダー表示
│   ├── goals/                    # 目標管理関連ウィジェット
│   │   ├── goalCard/             # 目標カード
│   │   ├── progressBar/          # 進捗バー
│   │   └── goalForm/             # 目標フォーム
│   ├── events/                   # イベント関連ウィジェット
│   │   ├── eventCard/            # イベントカード
│   │   ├── eventList/            # イベント一覧
│   │   └── participantList/      # 参加者一覧
│   ├── titles/                   # 称号関連ウィジェット
│   │   ├── titleCard/            # 称号カード
│   │   ├── titleBadge/           # 称号バッジ
│   │   └── achievementModal/     # 獲得モーダル
│   └── notifications/            # 通知関連ウィジェット
│       ├── notificationList/     # 通知一覧
│       └── notificationBadge/    # 通知バッジ
│
├── features/                     # 機能単位の実装
│   ├── attendance/               # 出席管理機能
│   │   ├── attendance-log/       # 出席ログ
│   │   │   ├── log-get/
│   │   │   │   ├── lib/
│   │   │   │   │   └── useAttendanceLogs.ts
│   │   │   │   └── index.ts
│   │   │   └── log-create/
│   │   ├── attendance-stats/     # 出席統計
│   │   │   ├── stats-get/
│   │   │   └── ranking-get/
│   │   └── streak/               # 連続記録
│   │       └── streak-get/
│   ├── goals/                    # 目標管理機能
│   │   ├── goal-create/          # 目標作成
│   │   ├── goal-update/          # 目標更新
│   │   ├── goal-delete/          # 目標削除
│   │   ├── goal-get/             # 目標取得
│   │   └── progress/             # 進捗管理
│   │       ├── progress-create/
│   │       └── progress-get/
│   ├── events/                   # イベント管理機能
│   │   ├── event-create/         # イベント作成
│   │   ├── event-update/         # イベント更新
│   │   ├── event-delete/         # イベント削除
│   │   ├── event-get/            # イベント取得
│   │   └── participation/        # 参加管理
│   │       ├── participate/
│   │       └── cancel/
│   ├── titles/                   # 称号管理機能
│   │   ├── title-get/            # 称号取得
│   │   ├── achievement-get/      # 獲得履歴取得
│   │   └── title-change/         # 称号変更
│   ├── notifications/            # 通知管理機能
│   │   ├── notification-get/     # 通知取得
│   │   ├── notification-read/    # 既読処理
│   │   └── settings/             # 通知設定
│   └── user/                     # ユーザー管理機能
│       ├── user-get/
│       ├── user-update/
│       └── profile/
│
├── entities/                     # ビジネスエンティティ
│   ├── attendance/               # 出席エンティティ
│   │   ├── api/                  # APIクライアント
│   │   ├── model/                # 型定義
│   │   ├── utils/
│   │   │   └── query-keys.ts    # クエリキー管理
│   │   └── index.ts
│   ├── goals/                    # 目標エンティティ
│   │   ├── api/
│   │   ├── model/
│   │   ├── utils/
│   │   └── index.ts
│   ├── events/                   # イベントエンティティ
│   │   ├── api/
│   │   ├── model/
│   │   ├── utils/
│   │   └── index.ts
│   ├── titles/                   # 称号エンティティ
│   │   ├── api/
│   │   ├── model/
│   │   ├── utils/
│   │   └── index.ts
│   ├── notifications/            # 通知エンティティ
│   │   ├── api/
│   │   ├── model/
│   │   ├── utils/
│   │   └── index.ts
│   ├── user/                     # ユーザーエンティティ
│   │   ├── api/
│   │   ├── model/
│   │   └── index.ts
│   └── ranking/                  # ランキング計算
│       ├── lib/
│       └── index.ts
│
└── shared/                       # 共有リソース
    ├── ui/                       # UIコンポーネント（shadcn/ui）
    │   ├── button.tsx
    │   ├── dialog.tsx
    │   ├── select.tsx
    │   └── ...
    ├── api/                      # APIクライアント設定
    │   ├── supabase-client.ts    # Supabase設定
    │   ├── clerk-client.ts       # Clerk設定
    │   └── index.ts
    ├── hooks/                    # 共通カスタムフック
    │   ├── use-toast.ts
    │   └── use-mobile.tsx
    ├── model/                    # 共通型定義
    │   └── index.ts
    └── utils/                    # ユーティリティ関数
        ├── date.ts
        ├── format.ts
        └── index.ts
```

## 各レイヤーのインポートルール

### 依存関係の方向（上位から下位へのみ）

```
page-components
    ↓
  widgets
    ↓
  features
    ↓
  entities
    ↓
   shared
```

### 重要な制約事項

1. **同一レイヤー間の依存禁止**: widgets間、features間、entities間での相互参照は禁止
2. **上位レイヤーへの依存禁止**: 下位レイヤーから上位レイヤーへの参照は禁止
3. **ESLintによる強制**: FSD boundaries ルールにより自動的にチェック

## まとめ

朝活コミュニティアプリGhoona CampにおけるFSDとNext.js App Routerの統合ポイント：

1. **src/app/はルーティングとグローバル設定を統合**: Next.js App Router、Clerk認証、React Queryを同一ディレクトリで管理
2. **page-components/でページオーケストレーション**: widgetsとfeaturesを組み合わせてページを構成
3. **widgets/はドメイン別に編成**: 出席管理、目標管理、イベント管理、称号管理など朝活ドメインごとに整理
4. **features/は細粒度で組織化**: エンティティと操作ごとに分離（例: `attendance/log/log-get/`）
5. **entities/でクエリキーを一元管理**: `utils/query-keys.ts`でクエリキーを集中管理
6. **shared/でSupabase・Clerk統合**: APIクライアント設定とshadcn/uiの統合
7. **すべてのモジュールでindex.tsを使用**: クリーンなPublic APIの提供
8. **一貫した命名規則**: `use[Entity][Action].ts`パターンの徹底

これらのパターンを活用することで、朝活コミュニティ機能に特化したスケーラブルで保守性の高いNext.jsアプリケーションを構築できます。

---

_最終更新: 2025-01-21_
