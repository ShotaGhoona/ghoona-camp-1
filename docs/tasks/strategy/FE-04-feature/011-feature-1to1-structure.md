# Feature Layer 1:1 Structure - User Domain

## 概要

APIエンドポイントと1:1対応のFeature Layer設計。13個のAPIエンドポイントに対して13個のFeatureを作成し、関連機能ごとにグループ化した2階層構造とする。各Featureは単一の責務を持つ明確な構造とする。

## 設計原則

### 基本方針

1. **1エンドポイント = 1Feature**: 明確な責務分担
2. **機能グループ化**: 関連するAPIを同じfeature配下にまとめる
3. **予測可能性**: Featureの名前から機能が一目瞭然
4. **保守性**: バグ修正・機能変更の影響範囲が限定的
5. **学習コスト**: 新メンバーが理解しやすい構造

### 命名規則

- **グループ**: `{resource}-feature`
- **個別Feature**: `{resource}-{action}`
- **例**: `social-links-feature/social-links-create/`
- **一貫性**: 同じリソースは同じfeatureグループ内

## Feature Structure

```
src/features/user/
├── index.ts                     # 全Featureのre-export
│
├── auth-feature/
│   └── auth-me-get/             # GET /auth/me
│       ├── index.ts
│       ├── lib/
│       │   └── use-auth-me-get.ts
│       └── ui/
│
├── user-feature/
│   ├── index.ts                 # user-feature内のre-export
│   │
│   ├── users-list-get/          # GET /users
│   │   ├── index.ts
│   │   ├── lib/
│   │   │   └── use-users-list-get.ts
│   │   ├── model/
│   │   │   └── search-types.ts  # 検索・フィルタ型
│   │   └── ui/
│   │       ├── UsersList.tsx
│   │       └── UsersSearchForm.tsx
│   │
│   ├── user-detail-get/         # GET /users/{userId}
│   │   ├── index.ts
│   │   ├── lib/
│   │   │   └── use-user-detail-get.ts
│   │   └── ui/
│   │       └── UserDetailCard.tsx
│   │
│   └── user-basic-update/       # PUT /users/{userId}
│       ├── index.ts
│       ├── lib/
│       │   └── use-user-basic-update.ts
│       ├── model/
│       │   └── user-form-types.ts
│       └── ui/
│           └── UserBasicForm.tsx
│
├── metadata-feature/
│   ├── index.ts                 # metadata-feature内のre-export
│   │
│   ├── metadata-get/            # GET /users/{userId}/metadata
│   │   ├── index.ts
│   │   ├── lib/
│   │   │   └── use-metadata-get.ts
│   │   └── ui/
│   │
│   └── metadata-update/         # PUT /users/{userId}/metadata
│       ├── index.ts
│       ├── lib/
│       │   └── use-metadata-update.ts
│       ├── model/
│       │   ├── metadata-form-types.ts
│       │   └── validation-schemas.ts
│       └── ui/
│           ├── MetadataForm.tsx
│           └── SkillsSelector.tsx
│
├── social-links-feature/
│   ├── index.ts                 # social-links-feature内のre-export
│   │
│   ├── social-links-get/        # GET /users/{userId}/social-links
│   │   ├── index.ts
│   │   ├── lib/
│   │   │   └── use-social-links-get.ts
│   │   └── ui/
│   │
│   ├── social-links-create/     # POST /users/{userId}/social-links
│   │   ├── index.ts
│   │   ├── lib/
│   │   │   └── use-social-links-create.ts
│   │   ├── model/
│   │   │   └── social-link-form-types.ts
│   │   └── ui/
│   │       └── AddSocialLinkForm.tsx
│   │
│   ├── social-links-update/     # PUT /users/{userId}/social-links/{linkId}
│   │   ├── index.ts
│   │   ├── lib/
│   │   │   └── use-social-links-update.ts
│   │   ├── model/
│   │   │   └── social-link-edit-types.ts
│   │   └── ui/
│   │       └── EditSocialLinkForm.tsx
│   │
│   └── social-links-delete/     # DELETE /users/{userId}/social-links/{linkId}
│       ├── index.ts
│       ├── lib/
│       │   └── use-social-links-delete.ts
│       └── ui/
│           └── DeleteSocialLinkButton.tsx
│
└── rivals-feature/
    ├── index.ts                 # rivals-feature内のre-export
    │
    ├── rivals-get/              # GET /users/{userId}/rivals
    │   ├── index.ts
    │   ├── lib/
    │   │   └── use-rivals-get.ts
    │   └── ui/
    │       ├── RivalsList.tsx
    │       └── RivalCard.tsx
    │
    ├── rivals-create/           # POST /users/{userId}/rivals
    │   ├── index.ts
    │   ├── lib/
    │   │   └── use-rivals-create.ts
    │   ├── model/
    │   │   └── rival-form-types.ts
    │   └── ui/
    │       ├── AddRivalModal.tsx
    │       └── RivalSearchForm.tsx
    │
    └── rivals-delete/           # DELETE /users/{userId}/rivals/{rivalId}
        ├── index.ts
        ├── lib/
        │   └── use-rivals-delete.ts
        └── ui/
            └── RemoveRivalButton.tsx
```

## Feature詳細設計

### 基本構造パターン

各Featureは以下の標準構造を持つ：

```
feature-name/
├── index.ts                 # エクスポート定義
├── lib/
│   └── use-feature-name.ts  # カスタムフック
├── model/                   # 型定義（必要に応じて）
│   ├── form-types.ts
│   └── validation.ts
└── ui/                      # UIコンポーネント（必要に応じて）
    └── FeatureComponent.tsx
```

### カスタムフック命名規則

```typescript
// パターン: use + Feature名
export const useAuthMeGet = () => { /* */ };
export const useUsersListGet = () => { /* */ };
export const useUserDetailGet = () => { /* */ };
export const useUserBasicUpdate = () => { /* */ };
export const useMetadataGet = () => { /* */ };
export const useMetadataUpdate = () => { /* */ };
export const useSocialLinksGet = () => { /* */ };
export const useSocialLinksCreate = () => { /* */ };
export const useSocialLinksUpdate = () => { /* */ };
export const useSocialLinksDelete = () => { /* */ };
export const useRivalsGet = () => { /* */ };
export const useRivalsCreate = () => { /* */ };
export const useRivalsDelete = () => { /* */ };
```

### Feature index.ts パターン

```typescript
// features/user/auth-me-get/index.ts
export * from './lib/use-auth-me-get';
export * from './ui';

// features/user/metadata-update/index.ts  
export * from './lib/use-metadata-update';
export type * from './model/metadata-form-types';
export * from './ui';
```

### Feature Group index.ts パターン

```typescript
// features/user/user-feature/index.ts
export * from './users-list-get';
export * from './user-detail-get';
export * from './user-basic-update';

// features/user/metadata-feature/index.ts
export * from './metadata-get';
export * from './metadata-update';

// features/user/social-links-feature/index.ts
export * from './social-links-get';
export * from './social-links-create';
export * from './social-links-update';
export * from './social-links-delete';

// features/user/rivals-feature/index.ts
export * from './rivals-get';
export * from './rivals-create';
export * from './rivals-delete';
```

### メインindex.ts

```typescript
// features/user/index.ts

// === Authentication ===
export * from './auth-feature';

// === User Management ===
export * from './user-feature';

// === Metadata Management ===
export * from './metadata-feature';

// === Social Links Management ===
export * from './social-links-feature';

// === Rivals Management ===
export * from './rivals-feature';
```

## 実装例

### 1. 基本的なGETフック

```typescript
// features/user/auth-feature/auth-me-get/lib/use-auth-me-get.ts
import { useQuery } from '@tanstack/react-query';
import { getCurrentUser, userQueryKeys } from '@/entities/user';

export const useAuthMeGet = () => {
  return useQuery({
    queryKey: userQueryKeys.current(),
    queryFn: getCurrentUser,
    staleTime: 1000 * 60 * 5, // 5分
  });
};
```

### 2. 検索機能付きGETフック

```typescript
// features/user/user-feature/users-list-get/lib/use-users-list-get.ts
import { useQuery } from '@tanstack/react-query';
import { getUsers, userQueryKeys, type UsersQueryParams } from '@/entities/user';

export const useUsersListGet = (params?: UsersQueryParams) => {
  return useQuery({
    queryKey: userQueryKeys.list(params),
    queryFn: () => getUsers(params),
    enabled: true,
  });
};
```

### 3. Mutationフック

```typescript
// features/user/metadata-feature/metadata-update/lib/use-metadata-update.ts
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateUserMetadata, userQueryKeys, type UpdateUserMetadataDto } from '@/entities/user';

export const useMetadataUpdate = (userId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdateUserMetadataDto) => 
      updateUserMetadata({ userId, data }),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: userQueryKeys.metadata(userId)
      });
    },
  });
};
```

## ページでの使用例

### プロフィール設定ページ

```tsx
// app/settings/profile/page.tsx
import { 
  useAuthMeGet,
  useUserBasicUpdate,
  useMetadataGet,
  useMetadataUpdate,
  useSocialLinksGet,
  useSocialLinksCreate,
  useSocialLinksUpdate,
  useSocialLinksDelete
} from '@/features/user';

const ProfileSettingsPage = () => {
  // 必要なフックを個別にインポート
  const { data: currentUser } = useAuthMeGet();
  const { data: metadata } = useMetadataGet(currentUser?.id);
  const { data: socialLinks } = useSocialLinksGet(currentUser?.id);
  
  const updateBasicMutation = useUserBasicUpdate(currentUser?.id);
  const updateMetadataMutation = useMetadataUpdate(currentUser?.id);
  const createSocialLinkMutation = useSocialLinksCreate(currentUser?.id);

  return (
    <div>
      <UserBasicForm 
        user={currentUser} 
        onSubmit={updateBasicMutation.mutate} 
      />
      <MetadataForm 
        metadata={metadata} 
        onSubmit={updateMetadataMutation.mutate} 
      />
      <SocialLinksManager 
        links={socialLinks}
        onCreate={createSocialLinkMutation.mutate}
      />
    </div>
  );
};
```

### メンバー一覧ページ

```tsx
// app/members/page.tsx
import { 
  useUsersListGet,
  useUserDetailGet,
  useRivalsCreate 
} from '@/features/user';

const MembersPage = () => {
  const [selectedUserId, setSelectedUserId] = useState<string>();
  const [searchParams, setSearchParams] = useState({});

  const { data: users } = useUsersListGet(searchParams);
  const { data: userDetail } = useUserDetailGet(selectedUserId);
  const createRivalMutation = useRivalsCreate();

  return (
    <div>
      <UsersSearchForm onChange={setSearchParams} />
      <UsersList 
        users={users} 
        onUserSelect={setSelectedUserId} 
      />
      {userDetail && (
        <UserDetailModal 
          user={userDetail}
          onAddRival={(rivalUserId) => 
            createRivalMutation.mutate({ rivalUserId })
          }
        />
      )}
    </div>
  );
};
```

### Feature Group別インポート例

```tsx
// Feature Group単位でのインポート
import { 
  useMetadataGet,
  useMetadataUpdate 
} from '@/features/user/metadata-feature';

import { 
  useSocialLinksGet,
  useSocialLinksCreate,
  useSocialLinksUpdate,
  useSocialLinksDelete 
} from '@/features/user/social-links-feature';

import { 
  useRivalsGet,
  useRivalsCreate,
  useRivalsDelete 
} from '@/features/user/rivals-feature';
```

## メリット

### 1. 明確性・予測可能性
- Feature名から機能が一目瞭然
- どのAPIを使っているか明確
- デバッグ時の調査範囲が限定的

### 2. 論理的なグループ化
- 関連機能が視覚的に分かりやすい
- ファイル探索が容易（13個→5グループ）
- 機能ドメインごとの整理

### 3. 保守性
- 1つのAPIの変更が1つのFeatureにのみ影響
- テストも1API分の責務のみ
- 新メンバーの学習コストが低い

### 4. 段階的開発
- 必要なFeatureから順次実装可能
- APIが完成していないFeatureはスキップ可能
- 独立性が高いため並行開発しやすい

### 5. インポートの柔軟性
- 全体インポート: `from '@/features/user'`
- グループ単位: `from '@/features/user/social-links-feature'`
- 個別インポート: `from '@/features/user/social-links-feature/social-links-create'`

### 6. 将来の統合可能性
- 1:1構造からの統合は比較的容易
- グループ内で関連Featureを後からまとめることも可能

## 注意点

### 1. ディレクトリ階層の増加
- 3階層構造（user/ → feature/ → action/）
- パスが長くなる傾向

### 2. 似たようなロジックの重複
- 同じリソースの異なる操作で似たロジックが発生する可能性
- バリデーション等の共通処理は entities 層で対応

### 3. Feature間の連携
- 複数Featureの結果を組み合わせる際は page-components 層で実装

### 4. Feature Group間の依存
- 同じグループ内での依存は比較的自然
- 異なるグループ間での依存は慎重に検討

## 今後の展開

### 他ドメインでの適用
- `features/attendance/` - 出席管理関連
- `features/events/` - イベント管理関連  
- `features/goals/` - 目標管理関連

### 段階的な統合検討
- 開発が進んで類似パターンが見えたら統合を検討
- ただし、基本は1:1を維持