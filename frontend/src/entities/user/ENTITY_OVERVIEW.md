# User Entities Overview

User関連の4つのEntityの機能と役割をまとめたドキュメントです。

## 目次

- [全体構成](#全体構成)
- [user-entity](#user-entity)
- [metadata-entity](#metadata-entity)
- [rivals-entity](#rivals-entity)
- [social-links-entity](#social-links-entity)
- [共通設計パターン](#共通設計パターン)
- [使用例](#使用例)

---

## 全体構成

```
src/entities/user/
├── user-entity/          # ユーザー基本情報管理
├── metadata-entity/      # ユーザープロフィール詳細管理
├── rivals-entity/        # ライバル機能管理
├── social-links-entity/  # SNSリンク管理
├── API_REFERENCE.md      # 全APIエンドポイント仕様
└── ENTITY_OVERVIEW.md    # 本ドキュメント
```

各Entityは以下の標準構造を持ちます：

```
{entity-name}/
├── index.ts              # エクスポート定義
├── api/                  # API呼び出し関数
├── model/                # 型定義
├── lib/
│   ├── validators.ts     # バリデーション機能
│   └── query-keys.ts     # React Query キー管理
```

---

## user-entity

**責務**: ユーザーの基本情報管理（認証・基本プロフィール・一覧取得）

### 提供機能

#### API関数
- `getCurrentUser()` - 現在のユーザー情報取得（認証確認）
- `getUsers(params?)` - ユーザー一覧取得（検索・フィルタ対応）
- `getUserDetail(id)` - 特定ユーザー詳細情報取得
- `updateUser({id, data})` - ユーザー基本情報更新

#### 型定義
- `User` - ユーザー基本情報型
- `AuthMeResponse` - 認証ユーザー情報レスポンス型
- `UsersListResponse` - ユーザー一覧レスポンス型
- `UserDetailResponse` - ユーザー詳細レスポンス型
- `UpdateUserDto` - ユーザー更新リクエスト型
- `UsersQueryParams` - ユーザー検索パラメータ型

#### バリデーション
- `usernameSchema` - ユーザー名バリデーション
- `updateUserFormSchema` - ユーザー更新フォームバリデーション
- `userSchema` - ユーザーエンティティバリデーション

#### Query Keys
- `userKeys.session()` - 認証ユーザー情報
- `userKeys.list(filters)` - ユーザー一覧
- `userKeys.detail(id)` - ユーザー詳細

### 対応API
- `GET /auth/me` - 現在ユーザー情報
- `GET /users` - ユーザー一覧・検索
- `GET /users/{userId}` - ユーザー詳細
- `PUT /users/{userId}` - ユーザー基本情報更新

---

## metadata-entity

**責務**: ユーザーのメタデータ管理（プロフィール詳細・スキル・興味）

### 提供機能

#### API関数
- `getUserMetadata(userId)` - ユーザーメタデータ取得
- `updateUserMetadata({userId, data})` - ユーザーメタデータ更新

#### 型定義
- `UserMetadata` - ユーザーメタデータ型
- `UserMetadataDetailResponse` - メタデータ詳細レスポンス型
- `UserMetadataUpdateResponse` - メタデータ更新レスポンス型
- `UpdateUserMetadataDto` - メタデータ更新リクエスト型

#### バリデーション
- `displayNameSchema` - 表示名バリデーション
- `taglineSchema` - 一言プロフィールバリデーション
- `bioSchema` - 自己紹介バリデーション
- `visionSchema` - ビジョンバリデーション
- `skillsSchema` - スキル配列バリデーション
- `interestsSchema` - 興味・関心配列バリデーション
- `timezoneSchema` - タイムゾーンバリデーション
- `updateUserMetadataFormSchema` - メタデータ更新フォームバリデーション

#### Query Keys
- `metadataKeys.detail(userId)` - ユーザーメタデータ
- `metadataKeys.update(userId)` - メタデータ更新

### 対応API
- `GET /users/{userId}/metadata` - ユーザーメタデータ取得
- `PUT /users/{userId}/metadata` - ユーザーメタデータ更新

### 管理項目
- **基本情報**: 表示名、プロフィール画像URL、一言プロフィール
- **詳細情報**: 自己紹介、ビジョン、タイムゾーン
- **スキル・興味**: スキル配列、興味・関心配列
- **公開設定**: ビジョンの公開/非公開設定

---

## rivals-entity

**責務**: ライバル機能管理（ライバル追加・削除・一覧表示）

### 提供機能

#### API関数
- `getUserRivals(userId)` - ユーザーのライバル一覧取得
- `createRival({userId, data})` - 新しいライバル追加
- `deleteRival({userId, rivalId})` - ライバル関係解除

#### 型定義
- `RivalUser` - ライバルユーザー情報型
- `Rival` - ライバル関係型
- `RivalsList` - ライバル一覧型
- `UserRivalsListResponse` - ライバル一覧レスポンス型
- `RivalCreateResponse` - ライバル作成レスポンス型
- `CreateRivalDto` - ライバル作成リクエスト型

#### バリデーション
- `rivalUserIdSchema` - ライバルユーザーIDバリデーション
- `createRivalFormSchema` - ライバル作成フォームバリデーション
- `validateRivalCount(count)` - ライバル数上限チェック
- `validateNotSelfRival(userId, rivalUserId)` - 自分自身をライバル設定防止
- `validateDuplicateRival(existingRivals, newRivalUserId)` - 重複ライバル防止

#### Query Keys
- `rivalsKeys.list(userId)` - ライバル一覧
- `rivalsKeys.create(userId)` - ライバル作成
- `rivalsKeys.delete(userId, rivalId)` - ライバル削除

### 対応API
- `GET /users/{userId}/rivals` - ライバル一覧取得
- `POST /users/{userId}/rivals` - ライバル追加
- `DELETE /users/{userId}/rivals/{rivalId}` - ライバル削除

### ビジネスルール
- **上限**: 最大3人までライバル設定可能
- **制限**: 自分自身をライバルに設定不可
- **重複**: 同一ユーザーの重複ライバル設定不可

---

## social-links-entity

**責務**: SNSリンク管理（ソーシャルメディア・外部リンク管理）

### 提供機能

#### API関数
- `getUserSocialLinks(userId)` - ユーザーのSNSリンク一覧取得
- `createSocialLink({userId, data})` - 新しいSNSリンク追加
- `updateSocialLink({userId, linkId, data})` - 既存SNSリンク更新
- `deleteSocialLink({userId, linkId})` - SNSリンク削除

#### 型定義
- `SocialLink` - SNSリンク型
- `UserSocialLinksListResponse` - SNSリンク一覧レスポンス型
- `SocialLinkCreateResponse` - SNSリンク作成レスポンス型
- `SocialLinkUpdateResponse` - SNSリンク更新レスポンス型
- `CreateSocialLinkDto` - SNSリンク作成リクエスト型
- `UpdateSocialLinkDto` - SNSリンク更新リクエスト型
- `SocialPlatform` - サポートプラットフォーム型

#### バリデーション
- `platformSchema` - プラットフォームバリデーション
- `urlSchema` - URLバリデーション
- `titleSchema` - タイトルバリデーション
- `createSocialLinkFormSchema` - SNSリンク作成フォームバリデーション
- `updateSocialLinkFormSchema` - SNSリンク更新フォームバリデーション

#### Query Keys
- `socialLinksKeys.list(userId)` - SNSリンク一覧
- `socialLinksKeys.create(userId)` - SNSリンク作成
- `socialLinksKeys.update(userId, linkId)` - SNSリンク更新
- `socialLinksKeys.delete(userId, linkId)` - SNSリンク削除

### 対応API
- `GET /users/{userId}/social-links` - SNSリンク一覧取得
- `POST /users/{userId}/social-links` - SNSリンク追加
- `PUT /users/{userId}/social-links/{linkId}` - SNSリンク更新
- `DELETE /users/{userId}/social-links/{linkId}` - SNSリンク削除

### サポートプラットフォーム
- `twitter`, `instagram`, `github`, `linkedin`
- `website`, `blog`, `youtube`, `facebook`
- `discord`, `twitch`

---

## 共通設計パターン

### Entity構造の統一性

各Entityは以下の統一された構造を持ちます：

```typescript
// index.ts - 統一されたエクスポート形式
export * from './api/{entity}-api';           // API関数
export type * from './model/{entity}-types';  // 型定義
export * from './lib/validators';              // バリデーション
export type * from './lib/validators';         // バリデーション型
export * from './lib/query-keys';              // Query Keys
```

### API関数命名規則

```typescript
// GET操作
export const get{EntityName} = async () => { /* */ };
export const get{EntityName}Detail = async (id) => { /* */ };

// POST操作
export const create{EntityName} = async ({data}) => { /* */ };

// PUT操作
export const update{EntityName} = async ({id, data}) => { /* */ };

// DELETE操作
export const delete{EntityName} = async ({id}) => { /* */ };
```

### 型定義命名規則

```typescript
// Core型
export interface {EntityName} { /* */ }

// API Response型
export interface {EntityName}DetailResponse { /* */ }
export interface {EntityName}CreateResponse { /* */ }
export interface {EntityName}UpdateResponse { /* */ }

// DTO型
export interface Create{EntityName}Dto { /* */ }
export interface Update{EntityName}Dto { /* */ }
```

### バリデーション統一パターン

```typescript
// フィールド単位のスキーマ
export const {fieldName}Schema = z.string().min(1);

// フォーム全体のスキーマ
export const {operation}{EntityName}FormSchema = z.object({ /* */ });

// Type Guards
export const is{EntityName} = (value: unknown): value is {EntityName} => { /* */ };
export const isValid{Operation}{EntityName}Dto = (value: unknown) => { /* */ };
```

### Query Keys統一パターン

```typescript
export const {entityName}Keys = {
  all: ['{entityName}'] as const,
  lists: () => [...{entityName}Keys.all, 'list'] as const,
  list: (filters) => [...{entityName}Keys.lists(), filters] as const,
  detail: (id) => [...{entityName}Keys.all, 'detail', id] as const,
  create: (id) => [...{entityName}Keys.all, 'create', id] as const,
  update: (id) => [...{entityName}Keys.all, 'update', id] as const,
  delete: (id) => [...{entityName}Keys.all, 'delete', id] as const,
} as const;
```

---

## 使用例

### 基本的な使用方法

```typescript
// Feature層での使用例
import { 
  getCurrentUser,
  getUserMetadata,
  getUserRivals,
  getUserSocialLinks,
  updateUser,
  updateUserMetadata,
  createRival,
  createSocialLink,
  userKeys,
  metadataKeys,
  rivalsKeys,
  socialLinksKeys,
  type User,
  type UserMetadata,
  type UpdateUserDto
} from '@/entities/user';

// React Queryでの使用例
const useProfileData = (userId: string) => {
  const userQuery = useQuery({
    queryKey: userKeys.detail(userId),
    queryFn: () => getUserDetail(userId)
  });

  const metadataQuery = useQuery({
    queryKey: metadataKeys.detail(userId),
    queryFn: () => getUserMetadata(userId)
  });

  const rivalsQuery = useQuery({
    queryKey: rivalsKeys.list(userId),
    queryFn: () => getUserRivals(userId)
  });

  const socialLinksQuery = useQuery({
    queryKey: socialLinksKeys.list(userId),
    queryFn: () => getUserSocialLinks(userId)
  });

  return {
    user: userQuery.data,
    metadata: metadataQuery.data,
    rivals: rivalsQuery.data,
    socialLinks: socialLinksQuery.data,
    isLoading: userQuery.isLoading || metadataQuery.isLoading
  };
};
```

### Mutation操作の例

```typescript
// プロフィール更新
const useProfileUpdate = (userId: string) => {
  const queryClient = useQueryClient();

  const updateBasicMutation = useMutation({
    mutationFn: (data: UpdateUserDto) => updateUser({ id: userId, data }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: userKeys.detail(userId) });
    }
  });

  const updateMetadataMutation = useMutation({
    mutationFn: (data: UpdateUserMetadataDto) => 
      updateUserMetadata({ userId, data }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: metadataKeys.detail(userId) });
    }
  });

  return { updateBasicMutation, updateMetadataMutation };
};
```

### バリデーション使用例

```typescript
import { 
  updateUserFormSchema,
  updateUserMetadataFormSchema,
  createRivalFormSchema,
  createSocialLinkFormSchema
} from '@/entities/user';

// フォームバリデーション
const profileFormData = updateUserFormSchema.parse(formValues);
const metadataFormData = updateUserMetadataFormSchema.parse(formValues);
const rivalFormData = createRivalFormSchema.parse(formValues);
const socialLinkFormData = createSocialLinkFormSchema.parse(formValues);
```

---

## 注意事項

### 直接インポート禁止

各Entityの内部実装（`api/`, `model/`, `lib/` 内のファイル）は直接インポートせず、必ず各Entityの `index.ts` を通してインポートしてください。

```typescript
// ❌ 直接インポート（禁止）
import { getCurrentUser } from '@/entities/user/user-entity/api/user-api';

// ✅ 正しいインポート
import { getCurrentUser } from '@/entities/user/user-entity';
// または
import { getCurrentUser } from '@/entities/user'; // 将来的にメインindex.tsが作成された場合
```

### 型の一貫性

すべてのAPI関数・型定義はcamelCase形式で統一されており、API仕様書（`API_REFERENCE.md`）と完全に一致しています。

### Entity間の依存関係

各Entityは独立しており、相互に直接依存しません。複数Entityのデータが必要な場合は、Feature層で組み合わせて使用してください。