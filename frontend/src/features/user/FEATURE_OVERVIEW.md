# User Features Overview

User関連の5つのFeature Groupと13個のFeatureの機能と役割をまとめたドキュメントです。

## 目次

- [全体構成](#全体構成)
- [auth-feature](#auth-feature)
- [user-feature](#user-feature)
- [metadata-feature](#metadata-feature)
- [social-links-feature](#social-links-feature)
- [rivals-feature](#rivals-feature)
- [設計パターン](#設計パターン)
- [使用例](#使用例)

---

## 全体構成

```
src/features/user/
├── index.ts                     # 全Featureのre-export
│
├── auth-feature/                # 認証関連機能
│   └── auth-me-get/             # GET /auth/me
│
├── user-feature/                # ユーザー基本管理
│   ├── users-list-get/          # GET /users
│   ├── user-detail-get/         # GET /users/{userId}
│   └── user-basic-update/       # PUT /users/{userId}
│
├── metadata-feature/            # メタデータ管理
│   ├── metadata-get/            # GET /users/{userId}/metadata
│   └── metadata-update/         # PUT /users/{userId}/metadata
│
├── social-links-feature/        # ソーシャルリンク管理
│   ├── social-links-get/        # GET /users/{userId}/social-links
│   ├── social-links-create/     # POST /users/{userId}/social-links
│   ├── social-links-update/     # PUT /users/{userId}/social-links/{linkId}
│   └── social-links-delete/     # DELETE /users/{userId}/social-links/{linkId}
│
└── rivals-feature/              # ライバル機能管理
    ├── rivals-get/              # GET /users/{userId}/rivals
    ├── rivals-create/           # POST /users/{userId}/rivals
    └── rivals-delete/           # DELETE /users/{userId}/rivals/{rivalId}
```

各Featureは以下の標準構造を持ちます：

```
feature-name/
├── index.ts                 # エクスポート定義
├── lib/
│   └── use-feature-name.ts  # React Queryベースのカスタムフック
├── model/                   # 型定義（必要に応じて）
│   ├── form-types.ts        # フォーム関連型定義
│   └── validation.ts        # Feature特有のバリデーション
```

---

## auth-feature

**責務**: ユーザー認証状態の管理とセッション情報取得

### 提供機能

#### カスタムフック
- `useAuthMeGet()` - 現在ログイン中のユーザー情報取得

#### 対応API
- `GET /auth/me` - 現在ユーザー情報取得

#### 特徴
- **キャッシュ戦略**: 5分間のstaleTime設定
- **リトライ**: 2回まで自動リトライ
- **用途**: プロフィール設定画面の初期表示、認証状態確認

### 使用例
```typescript
import { useAuthMeGet } from '@/features/user/auth-feature';

const { data: currentUser, isLoading, error } = useAuthMeGet();
```

---

## user-feature

**責務**: ユーザーの基本情報管理（一覧・詳細・更新）

### 提供機能

#### カスタムフック
- `useUsersListGet(params?)` - ユーザー一覧取得（検索・フィルタ対応）
- `useUserDetailGet(userId)` - 特定ユーザー詳細情報取得
- `useUserBasicUpdate(userId)` - ユーザー基本情報更新

#### 型定義
- `UserBasicUpdateFormData` - ユーザー更新フォーム型
- `UserBasicUpdateFormProps` - フォームコンポーネント用Props型

#### バリデーション
- `userBasicUpdateFormSchema` - Entity層を基盤としたフォーム検証

#### 対応API
- `GET /users` - ユーザー一覧・検索
- `GET /users/{userId}` - ユーザー詳細
- `PUT /users/{userId}` - ユーザー基本情報更新

#### 特徴
- **検索・フィルタ**: スキル、興味、並び順での絞り込み対応
- **キャッシュ戦略**: 一覧2分、詳細3分のstaleTime
- **自動無効化**: 更新時に関連クエリを自動無効化

### 使用例
```typescript
import { 
  useUsersListGet,
  useUserDetailGet,
  useUserBasicUpdate 
} from '@/features/user/user-feature';

// 検索パラメータ付きユーザー一覧取得
const { data: users } = useUsersListGet({
  search: 'john',
  skills: ['JavaScript', 'React'],
  sortBy: 'attendanceDays',
  order: 'desc'
});

// ユーザー詳細取得
const { data: userDetail } = useUserDetailGet(userId);

// ユーザー更新
const updateMutation = useUserBasicUpdate(userId);
updateMutation.mutate({
  username: 'newUsername',
  avatarUrl: 'https://new-avatar.jpg'
});
```

---

## metadata-feature

**責務**: ユーザーメタデータ管理（プロフィール詳細・スキル・興味）

### 提供機能

#### カスタムフック
- `useMetadataGet(userId)` - ユーザーメタデータ取得
- `useMetadataUpdate(userId)` - ユーザーメタデータ更新

#### 型定義
- `MetadataUpdateFormData` - メタデータ更新フォーム型
- `MetadataUpdateFormProps` - フォームコンポーネント用Props型
- `SkillsSelectorProps` - スキル選択コンポーネント用Props型

#### バリデーション
- `metadataUpdateFormSchema` - Entity層を基盤としたメタデータ検証

#### 対応API
- `GET /users/{userId}/metadata` - ユーザーメタデータ取得
- `PUT /users/{userId}/metadata` - ユーザーメタデータ更新

#### 特徴
- **豊富なフィールド**: 表示名、自己紹介、ビジョン、スキル、興味等
- **公開設定**: ビジョンの公開/非公開制御
- **配列管理**: スキル・興味の動的追加/削除対応

### 使用例
```typescript
import { 
  useMetadataGet,
  useMetadataUpdate,
  type MetadataUpdateFormData 
} from '@/features/user/metadata-feature';

const { data: metadata } = useMetadataGet(userId);
const updateMutation = useMetadataUpdate(userId);

const handleSubmit = (data: MetadataUpdateFormData) => {
  updateMutation.mutate({
    displayName: 'John Doe Jr.',
    bio: '新しい自己紹介文',
    skills: ['JavaScript', 'React', 'TypeScript'],
    interests: ['Web Development', 'AI']
  });
};
```

---

## social-links-feature

**責務**: SNSリンク・外部リンク管理（CRUD操作）

### 提供機能

#### カスタムフック
- `useSocialLinksGet(userId)` - ユーザーのSNSリンク一覧取得
- `useSocialLinksCreate(userId)` - 新しいSNSリンク追加
- `useSocialLinksUpdate(userId, linkId)` - 既存SNSリンク更新
- `useSocialLinksDelete(userId)` - SNSリンク削除

#### 型定義
- `SocialLinkCreateFormData` - SNSリンク作成フォーム型
- `SocialLinkUpdateFormData` - SNSリンク更新フォーム型
- `PlatformOption` - プラットフォーム選択用型
- `SocialLinkEditState` - 編集状態管理用型

#### バリデーション
- `socialLinkCreateFormSchema` - SNSリンク作成検証
- `socialLinkUpdateFormSchema` - SNSリンク更新検証

#### 対応API
- `GET /users/{userId}/social-links` - SNSリンク一覧取得
- `POST /users/{userId}/social-links` - SNSリンク追加
- `PUT /users/{userId}/social-links/{linkId}` - SNSリンク更新
- `DELETE /users/{userId}/social-links/{linkId}` - SNSリンク削除

#### サポートプラットフォーム
- `twitter`, `instagram`, `github`, `linkedin`
- `website`, `blog`, `youtube`, `facebook`, `discord`, `twitch`

#### 特徴
- **完全CRUD**: 作成・読取・更新・削除の全操作対応
- **プラットフォーム制御**: 許可されたプラットフォームのみ選択可能
- **公開設定**: リンクごとの公開/非公開制御

### 使用例
```typescript
import { 
  useSocialLinksGet,
  useSocialLinksCreate,
  useSocialLinksUpdate,
  useSocialLinksDelete 
} from '@/features/user/social-links-feature';

const { data: socialLinks } = useSocialLinksGet(userId);
const createMutation = useSocialLinksCreate(userId);
const updateMutation = useSocialLinksUpdate(userId, linkId);
const deleteMutation = useSocialLinksDelete(userId);

// 新規作成
createMutation.mutate({
  platform: 'github',
  url: 'https://github.com/username',
  title: 'GitHub Profile',
  isPublic: true
});

// 更新
updateMutation.mutate({
  url: 'https://github.com/new-username',
  title: 'Updated GitHub',
  isPublic: false
});

// 削除
deleteMutation.mutate(linkId);
```

---

## rivals-feature

**責務**: ライバル機能管理（ライバル追加・削除・一覧表示）

### 提供機能

#### カスタムフック
- `useRivalsGet(userId)` - ユーザーのライバル一覧取得
- `useRivalsCreate(userId)` - 新しいライバル追加
- `useRivalsDelete(userId)` - ライバル関係解除

#### 型定義
- `RivalCreateFormData` - ライバル作成フォーム型
- `RivalValidationContext` - ビジネスルール検証用型
- `RivalCreateModalState` - モーダル状態管理用型

#### バリデーション・ビジネスルール
- `rivalCreateFormSchema` - ライバル作成検証
- `validateRivalBusiness()` - 包括的ビジネスルール検証
- `MAX_RIVALS_COUNT` - 最大ライバル数定数（3人）

#### 対応API
- `GET /users/{userId}/rivals` - ライバル一覧取得
- `POST /users/{userId}/rivals` - ライバル追加
- `DELETE /users/{userId}/rivals/{rivalId}` - ライバル削除

#### ビジネスルール
- **上限制御**: 最大3人までライバル設定可能
- **自己防止**: 自分自身をライバルに設定不可
- **重複防止**: 同一ユーザーの重複ライバル設定不可

#### 特徴
- **厳密な制約**: Entity層とFeature層でのダブルチェック
- **包括的検証**: 全ビジネスルールを一括検証する関数提供
- **UI支援**: モーダル状態管理やフォーム制御用の型定義

### 使用例
```typescript
import { 
  useRivalsGet,
  useRivalsCreate,
  useRivalsDelete,
  validateRivalBusiness,
  MAX_RIVALS_COUNT 
} from '@/features/user/rivals-feature';

const { data: rivals } = useRivalsGet(userId);
const createMutation = useRivalsCreate(userId);
const deleteMutation = useRivalsDelete(userId);

// ビジネスルール検証
try {
  validateRivalBusiness({
    currentUserId: userId,
    existingRivals: rivals?.rivals || [],
    targetUserId: selectedUserId
  });
  
  // 検証通過後にライバル追加
  createMutation.mutate({
    rivalUserId: selectedUserId
  });
} catch (error) {
  console.error('Validation failed:', error);
}

// ライバル削除
deleteMutation.mutate(rivalId);
```

---

## 設計パターン

### 1:1 API Mapping

各APIエンドポイントに対応する専用のカスタムフックを提供：

```typescript
// 13個のAPIエンドポイント → 13個のカスタムフック
GET /auth/me                              → useAuthMeGet()
GET /users                                → useUsersListGet()
GET /users/{userId}                       → useUserDetailGet()
PUT /users/{userId}                       → useUserBasicUpdate()
GET /users/{userId}/metadata              → useMetadataGet()
PUT /users/{userId}/metadata              → useMetadataUpdate()
GET /users/{userId}/social-links          → useSocialLinksGet()
POST /users/{userId}/social-links         → useSocialLinksCreate()
PUT /users/{userId}/social-links/{linkId} → useSocialLinksUpdate()
DELETE /users/{userId}/social-links/{linkId} → useSocialLinksDelete()
GET /users/{userId}/rivals                → useRivalsGet()
POST /users/{userId}/rivals               → useRivalsCreate()
DELETE /users/{userId}/rivals/{rivalId}   → useRivalsDelete()
```

### カスタムフック設計パターン

#### Query Hooks（GET操作）
```typescript
export const useFeatureGet = (param?: string) => {
  return useQuery({
    queryKey: entityKeys.detail(param!),
    queryFn: () => apiFunction(param!),
    enabled: !!param,
    staleTime: 1000 * 60 * N, // 適切なキャッシュ時間
  });
};
```

#### Mutation Hooks（POST/PUT/DELETE操作）
```typescript
export const useFeatureUpdate = (id: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdateDto) => apiFunction({ id, data }),
    onSuccess: () => {
      // 関連クエリの無効化
      queryClient.invalidateQueries({ queryKey: entityKeys.detail(id) });
    },
    onError: (error) => {
      console.error('Operation failed:', error);
    },
  });
};
```

### Feature Layer Validation

Entity層のvalidationを基盤とし、Feature特有のロジックを追加：

```typescript
// Entity層を基盤として使用
export const featureFormSchema = entityFormSchema.extend({
  // Feature特有のバリデーション追加
});

// ビジネスルール検証関数
export const validateFeatureBusiness = (context: ValidationContext) => {
  // 複合的なビジネスルール検証
};
```

### エクスポート管理

階層的なエクスポート構造で利用しやすさを向上：

```typescript
// メインindex.ts
export * from './auth-feature';
export * from './user-feature';
export * from './metadata-feature';
export * from './social-links-feature';
export * from './rivals-feature';

// Feature Group index.ts
export * from './feature-action-1';
export * from './feature-action-2';

// Individual Feature index.ts
export * from './lib/use-feature-action';
export type * from './model/form-types';
```

---

## 使用例

### 基本的な使用方法

```typescript
// 全体インポート
import { 
  useAuthMeGet,
  useUserDetailGet,
  useMetadataUpdate,
  useSocialLinksCreate 
} from '@/features/user';

// Feature Group別インポート
import { useMetadataGet, useMetadataUpdate } from '@/features/user/metadata-feature';
import { useSocialLinksGet, useSocialLinksCreate } from '@/features/user/social-links-feature';
```

### プロフィール設定ページでの総合利用

```typescript
import { 
  useAuthMeGet,
  useUserBasicUpdate,
  useMetadataGet,
  useMetadataUpdate,
  useSocialLinksGet,
  useSocialLinksCreate,
  useSocialLinksUpdate,
  useSocialLinksDelete,
  useRivalsGet,
  useRivalsCreate,
  useRivalsDelete
} from '@/features/user';

const ProfileSettingsPage = () => {
  // 現在ユーザー取得
  const { data: currentUser } = useAuthMeGet();
  const userId = currentUser?.id;

  // 各データ取得
  const { data: metadata } = useMetadataGet(userId);
  const { data: socialLinks } = useSocialLinksGet(userId);
  const { data: rivals } = useRivalsGet(userId);

  // 更新・作成・削除 Mutations
  const updateBasicMutation = useUserBasicUpdate(userId!);
  const updateMetadataMutation = useMetadataUpdate(userId!);
  const createSocialLinkMutation = useSocialLinksCreate(userId!);
  const createRivalMutation = useRivalsCreate(userId!);

  return (
    <div>
      <UserBasicForm 
        user={currentUser} 
        onSubmit={updateBasicMutation.mutate} 
        isLoading={updateBasicMutation.isPending}
      />
      
      <MetadataForm 
        metadata={metadata} 
        onSubmit={updateMetadataMutation.mutate}
        isLoading={updateMetadataMutation.isPending}
      />
      
      <SocialLinksManager 
        links={socialLinks}
        onCreate={createSocialLinkMutation.mutate}
        onUpdate={(linkId, data) => updateSocialLinkMutation.mutate(data)}
        onDelete={deleteSocialLinkMutation.mutate}
      />
      
      <RivalsManager
        rivals={rivals}
        onCreate={createRivalMutation.mutate}
        onDelete={deleteRivalMutation.mutate}
        maxCount={MAX_RIVALS_COUNT}
      />
    </div>
  );
};
```

### メンバー一覧・詳細ページでの利用

```typescript
import { 
  useUsersListGet,
  useUserDetailGet,
  useRivalsCreate 
} from '@/features/user';

const MembersPage = () => {
  const [selectedUserId, setSelectedUserId] = useState<string>();
  const [searchParams, setSearchParams] = useState({});

  // ユーザー一覧取得（検索・フィルタ対応）
  const { data: usersResponse } = useUsersListGet(searchParams);
  
  // 選択ユーザーの詳細取得
  const { data: userDetail } = useUserDetailGet(selectedUserId);
  
  // ライバル追加機能
  const createRivalMutation = useRivalsCreate(currentUserId);

  return (
    <div>
      <UsersSearchForm 
        onChange={setSearchParams}
        filters={['skills', 'interests']}
        sorting={['name', 'attendanceDays', 'streakDays']}
      />
      
      <UsersList 
        users={usersResponse?.users} 
        pagination={usersResponse?.pagination}
        onUserSelect={setSelectedUserId} 
      />
      
      {userDetail && (
        <UserDetailModal 
          user={userDetail}
          onAddRival={(rivalUserId) => 
            createRivalMutation.mutate({ rivalUserId })
          }
          onClose={() => setSelectedUserId(undefined)}
        />
      )}
    </div>
  );
};
```

### エラーハンドリングとローディング状態

```typescript
const ProfilePage = () => {
  const { data: user, isLoading, error } = useAuthMeGet();
  const updateMutation = useUserBasicUpdate(user?.id!);

  if (isLoading) return <LoadingSpinner />;
  if (error) return <ErrorMessage error={error} />;

  const handleSubmit = async (data: UserBasicUpdateFormData) => {
    try {
      await updateMutation.mutateAsync(data);
      toast.success('プロフィールを更新しました');
    } catch (error) {
      toast.error('更新に失敗しました');
    }
  };

  return (
    <UserBasicForm 
      user={user}
      onSubmit={handleSubmit}
      isLoading={updateMutation.isPending}
      error={updateMutation.error}
    />
  );
};
```

---

## 利点とメリット

### 1. 明確な責務分担
- 1APIエンドポイント = 1カスタムフック
- 機能の境界が明確で理解しやすい
- デバッグ・テストの範囲が限定的

### 2. 段階的開発
- 必要なFeatureから順次実装可能
- 他Featureの実装状況に依存しない
- 並行開発が容易

### 3. 保守性
- 影響範囲が明確で変更リスクが低い
- Feature単位でのテスト・デバッグが容易
- 新メンバーの学習コストが低い

### 4. 柔軟なインポート
- 必要なFeatureのみインポート可能
- バンドルサイズの最適化
- 依存関係の明確化

### 5. Entity層との適切な分離
- Entity層: API仕様準拠のvalidation
- Feature層: ビジネスロジック特有のvalidation
- 責務が明確で重複を避けた設計

---

## 注意事項

### 直接インポート禁止

各Featureの内部実装（`lib/`, `model/` 内のファイル）は直接インポートせず、必ず各Featureの `index.ts` を通してインポートしてください。

```typescript
// ❌ 直接インポート（禁止）
import { useAuthMeGet } from '@/features/user/auth-feature/auth-me-get/lib/use-auth-me-get';

// ✅ 正しいインポート
import { useAuthMeGet } from '@/features/user/auth-feature';
// または
import { useAuthMeGet } from '@/features/user';
```

### Feature間の依存関係

- 同じFeature Group内での依存は自然
- 異なるFeature Group間での依存は慎重に検討
- 複数Featureの結果を組み合わせる際はPage Components層で実装

### パフォーマンス考慮

- 必要なFeatureのみインポートしてバンドルサイズを最適化
- React Queryのキャッシュ戦略を適切に設定
- 無効化処理は必要最小限に留める