# Feature Layer Strategy - User Domain

## 概要

User関連のFeature Layer設計戦略書。FSD（Feature-Sliced Design）アーキテクチャに基づき、エンドポイント1:1ではなく、ページ要件とユーザー操作を基準とした機能単位でFeature層を設計する。

## 基本方針

### 設計原則

1. **ページ駆動設計**: 13ページの要件から逆算して必要な機能を特定
2. **ユーザー操作重視**: APIエンドポイントではなく、ユーザーが実現したい操作を基準
3. **機能凝集**: 関連するCRUD操作を1つのFeatureに集約
4. **再利用性**: 複数ページで使い回せる設計

### 現状の課題

- 13個のAPIエンドポイントに対して13個のFeatureが存在（1:1対応）
- 機能が細分化されすぎて、ページ開発時に複数Featureを組み合わせる必要
- 関連する操作（GET+PUT等）が分散している

## 新Feature構造

### User Domain Feature Map

```
src/features/user/
├── profile-management/          # プロフィール管理
├── user-discovery/              # ユーザー探索・発見
├── rival-management/            # ライバル機能
└── user-statistics/             # ユーザー統計・ランキング
```

### 1. profile-management/

**責務**: 自分のプロフィール情報の表示・編集・管理

**対象ページ**:
- `/settings/profile` - プロフィール設定
- `/settings/vision` - ビジョン設定
- `/dashboard` - 自分の情報表示

**使用API**:
- `GET /auth/me` - 現在ユーザー情報取得
- `PUT /users/{userId}` - 基本情報更新
- `GET/PUT /users/{userId}/metadata` - メタデータ管理
- `GET/POST/PUT/DELETE /users/{userId}/social-links` - SNSリンク管理

**ディレクトリ構造**:
```
profile-management/
├── index.ts
├── lib/
│   ├── use-profile-data.ts          # プロフィール情報取得
│   ├── use-profile-update.ts        # 基本情報更新
│   ├── use-metadata-management.ts   # メタデータ管理
│   ├── use-vision-management.ts     # ビジョン専用
│   └── use-social-links.ts          # SNSリンク管理
├── model/
│   ├── profile-form-types.ts        # フォーム型定義
│   └── validation-schemas.ts        # バリデーション
└── ui/
    ├── ProfileEditForm.tsx
    ├── VisionEditForm.tsx
    ├── SocialLinksManager.tsx
    └── ProfilePreview.tsx
```

### 2. user-discovery/

**責務**: 他ユーザーの検索・発見・詳細表示

**対象ページ**:
- `/members` - メンバー一覧・検索
- `/ranking` - ランキング表示（ユーザー情報部分）

**使用API**:
- `GET /users` - ユーザー一覧・検索
- `GET /users/{userId}` - ユーザー詳細情報
- `GET /users/{userId}/metadata` - ユーザーメタデータ

**ディレクトリ構造**:
```
user-discovery/
├── index.ts
├── lib/
│   ├── use-users-search.ts          # 検索・フィルタリング
│   ├── use-user-detail.ts           # ユーザー詳細取得
│   └── use-members-list.ts          # メンバー一覧管理
├── model/
│   ├── search-types.ts              # 検索条件型
│   └── filter-types.ts              # フィルタ型
└── ui/
    ├── MemberSearchForm.tsx
    ├── UserCard.tsx
    ├── UserDetailModal.tsx
    └── MembersList.tsx
```

### 3. rival-management/

**責務**: ライバル機能の全操作（追加・削除・表示）

**対象ページ**:
- `/dashboard` - ライバル状況表示
- `/members` - ライバル追加・削除
- `/ranking` - ライバルのランキング表示

**使用API**:
- `GET /users/{userId}/rivals` - ライバル一覧取得
- `POST /users/{userId}/rivals` - ライバル追加
- `DELETE /users/{userId}/rivals/{rivalId}` - ライバル削除

**ディレクトリ構造**:
```
rival-management/
├── index.ts
├── lib/
│   ├── use-rivals-list.ts           # ライバル一覧管理
│   ├── use-rival-actions.ts         # 追加・削除操作
│   └── use-rival-status.ts          # ライバル状況取得
├── model/
│   └── rival-types.ts               # ライバル関連型
└── ui/
    ├── RivalsList.tsx
    ├── AddRivalButton.tsx
    ├── RemoveRivalButton.tsx
    └── RivalStatusCard.tsx
```

### 4. user-statistics/

**責務**: ユーザーの統計情報・ランキング・実績表示

**対象ページ**:
- `/dashboard` - 自分の統計表示
- `/ranking` - 全体ランキング
- `/activity` - 個人実績表示

**使用API**:
- `GET /auth/me` - 自分の統計情報
- `GET /users` - ランキング用ユーザー一覧
- `GET /users/{userId}` - 特定ユーザー統計

**ディレクトリ構造**:
```
user-statistics/
├── index.ts
├── lib/
│   ├── use-user-stats.ts            # 個人統計情報
│   ├── use-ranking-data.ts          # ランキング情報
│   └── use-stats-comparison.ts      # 比較統計
├── model/
│   └── stats-types.ts               # 統計関連型
└── ui/
    ├── StatsCard.tsx
    ├── RankingTable.tsx
    ├── StatsChart.tsx
    └── ComparisonView.tsx
```

## 移行戦略

### Phase 1: profile-management 統合

**現在のFeature**:
- `metadata-feature/metadata-get`
- `metadata-feature/metadata-update`
- `profile-feature/profile-get`
- `profile-feature/profile-update`
- `profile-feature/session-get`
- `social-links-feature/*`

**移行手順**:
1. `profile-management/` ディレクトリ作成
2. 既存のカスタムフックを統合・リファクタ
3. UIコンポーネントを新構造に移動
4. 既存のindex.tsを更新
5. import文の更新

### Phase 2: user-discovery 統合

**現在のFeature**:
- `user-feature/users-list`
- `user-feature/user-detail-get`

**移行手順**:
1. 検索・フィルタ機能を強化
2. メンバー一覧用UIコンポーネント作成
3. 既存コードを新構造に統合

### Phase 3: rival-management 統合

**現在のFeature**:
- `rivals-feature/rivals-get`
- `rivals-feature/rivals-create`
- `rivals-feature/rivals-delete`

**移行手順**:
1. CRUD操作を1つのカスタムフックに統合
2. ライバル関連UIを統合
3. ダッシュボード用コンポーネント作成

### Phase 4: user-statistics 作成

**新規Feature**:
- 既存の統計情報取得ロジックを集約
- ランキング用データ取得ロジック統合
- 統計表示用UIコンポーネント作成

## 実装詳細

### カスタムフック設計

```typescript
// profile-management/lib/use-profile-management.ts
export const useProfileManagement = () => {
  const { data: profile, refetch } = useProfileData();
  const updateProfileMutation = useProfileUpdate();
  const updateMetadataMutation = useMetadataUpdate();
  const socialLinksManager = useSocialLinks();

  const updateProfile = async (data: ProfileUpdateData) => {
    // 基本情報とメタデータを組み合わせて更新
    await Promise.all([
      updateProfileMutation.mutateAsync(data.basic),
      updateMetadataMutation.mutateAsync(data.metadata)
    ]);
    refetch();
  };

  return {
    profile,
    updateProfile,
    socialLinks: socialLinksManager,
    isLoading: updateProfileMutation.isLoading || updateMetadataMutation.isLoading
  };
};
```

### ページでの使用例

```tsx
// app/settings/profile/page.tsx
import { useProfileManagement } from '@/features/user/profile-management';

const ProfileSettingsPage = () => {
  const { profile, updateProfile, socialLinks } = useProfileManagement();

  return (
    <div>
      <ProfileEditForm 
        profile={profile} 
        onSubmit={updateProfile} 
      />
      <SocialLinksManager {...socialLinks} />
    </div>
  );
};
```

## 期待効果

### 開発効率向上
- 関連機能が1箇所に集約され、開発・保守が容易
- ページ開発時にFeatureの組み合わせがシンプル
- テストも機能単位で書きやすい

### コード品質向上
- 機能の凝集性が高まる
- 重複コードの削減
- 型安全性の向上

### スケーラビリティ
- 新機能追加時の影響範囲が明確
- Feature間の依存関係が整理される
- チーム開発時の作業分担が明確

## 今後の拡張

### 追加予定Feature
- `user-preferences/` - 通知設定、アカウント設定
- `user-achievements/` - 称号・実績管理
- `user-connections/` - フォロー・ネットワーク機能

### 長期的な展望
- 他ドメイン（attendance, events, goals等）でも同様の設計パターンを適用
- Domain間の依存関係を最小化
- マイクロフロントエンド化の準備