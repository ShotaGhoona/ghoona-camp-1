# FE-03-entity 実装戦略書（改訂版）

## 概要
Ghoona Camp フロントエンドにおけるEntities層の実装戦略を定義する。FSD（Feature-Sliced Design）アーキテクチャに基づき、eagle-aiプロジェクトの実装パターンを参考にしつつ、Ghoona Campの要件に最適化したビジネスエンティティの型定義、API クライアント、クエリキー管理を統一的に実装する。

## 戦略目標
1. **型安全性の確保**: 厳密なTypeScript型定義による実行時エラーの防止
2. **APIクライアントの統一**: Clerk認証とGhoona Camp APIレスポンス形式に対応した一貫性のあるAPI呼び出しパターン
3. **効率的なキャッシュ管理**: React Query + Zustandによる最適化されたデータフェッチング・状態管理
4. **保守性の向上**: FSDルールに準拠し、eagle-aiパターンを参考にした構造化されたコードベース

## 実装アプローチ

### 1. 段階的実装戦略

#### Phase 1: Core User Entity（最優先 - Week 1）
- **User Entity** (user-entity/)
  - 基本ユーザー情報の型定義
  - 認証関連APIクライアント
  - ユーザー管理のクエリキー

#### Phase 2: User関連サブエンティティ（高優先 - Week 1-2）
- **User Metadata Entity** (metadata-entity/)
  - プロフィール詳細情報
  - スキル・興味関心管理
- **User Social Links Entity** (social-links-entity/)
  - SNS・外部リンク管理
- **User Rivals Entity** (rivals-entity/)
  - ライバル関係管理

#### Phase 3: 朝活コアエンティティ（中優先 - Week 2-3）
- **Attendance Entity**
  - 出席ログ・統計
- **Goals Entity**
  - 目標管理
- **Events Entity**
  - イベント管理

#### Phase 4: 付加価値エンティティ（低優先 - Week 3-4）
- **Titles Entity**
  - 称号・実績管理
- **Notifications Entity**
  - 通知管理

### 2. アーキテクチャ設計原則（eagle-ai参考）

#### 2.1 ディレクトリ構造の統一（eagle-aiパターン準拠）
```
entities/
├── user/
│   ├── user-entity/
│   │   ├── api/
│   │   │   ├── user-api.ts      # ユーザー関連API
│   │   │   └── index.ts
│   │   ├── model/
│   │   │   ├── types.ts         # User型定義
│   │   │   ├── store.ts         # Zustand状態管理（sliceの代替）
│   │   │   └── index.ts
│   │   ├── lib/
│   │   │   ├── validators.ts    # Zodバリデーション関数
│   │   │   └── index.ts
│   │   ├── utils/
│   │   │   ├── query-keys.ts    # React Queryキー
│   │   │   └── index.ts
│   │   └── index.ts
│   ├── metadata-entity/         # 同様の構造
│   ├── social-links-entity/     # 同様の構造
│   ├── rivals-entity/          # 同様の構造
│   └── utils/
│       └── query-keys.ts        # 全ユーザー関連クエリキー統合
```

#### 2.2 型定義戦略（Ghoona Camp仕様準拠）
- **API Response Types**: Ghoona Camp統一レスポンス形式（data, message, timestamp）
- **UUID型の徹底**: すべてのIDをstring型（UUID）で統一
- **Client Types**: フロントエンド専用の型（UI状態、フォーム等）
- **Validation Schemas**: Zodによる実行時型検証（eagle-aiの独自バリデーションから移行）
- **DTO Types**: 作成・更新用のData Transfer Object型定義

#### 2.3 APIクライアント設計（Clerk認証対応）
- **shared/api/client使用**: 統一されたAPIクライアントを活用
- **Clerk認証統合**: `useAuth().getToken()`でJWTトークンを自動取得
- **Ghoona Campレスポンス対応**: `{ data, message, timestamp }`形式への対応
- **エラーハンドリング**: shared/apiクライアントのエラー処理を活用
- **Type Safety**: レスポンス型の厳密な定義

#### 2.4 状態管理戦略（Redux → Zustand移行）
- **Zustand Store**: Redux sliceの代替として軽量な状態管理
- **React Query**: APIデータのキャッシュ・同期はReact Queryに委譲
- **直接アクセス**: Zustandの特性を活かし、selectorsファイルは作らずに直接状態アクセス
- **状態分離**: 認証状態（Zustand）とAPIキャッシュ（React Query）を明確に分離
- **Custom Hooks**: 複雑な状態計算が必要な場合はlib/内でカスタムフックとして実装

#### 2.5 クエリキー管理戦略（eagle-aiパターン採用）
- **階層的管理**: `userKeys.detail(id)`のような階層構造
- **無効化戦略**: 関連データの適切な無効化ルール
- **依存関係**: エンティティ間のデータ依存関係の明確化

### 3. 実装詳細

#### 3.1 User Entity 実装仕様（範囲を限定）

##### 🎯 user-entityの責務範囲
- **基本ユーザー情報のみ**: id, clerk_id, email, username, avatar_url, discord_id等
- **認証状態管理**: GET /auth/meの基本情報部分
- **基本情報更新**: PUT /users/{userId}
- **他エンティティへの委譲**: metadata → metadata-entity, attendance_stats → attendance-entity, social_links → social-links-entity

##### 型定義 (model/types.ts) - user-entity範囲のみ
```typescript
// === API Response Types（user-entityが担当する範囲のみ） ===

// 基本ユーザー情報（user-entity の核心）
export interface UserResponse {
  id: string;
  clerk_id: string;
  email: string;
  username: string;
  avatar_url: string;
  discord_id?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

// === GET /auth/me レスポンス（user-entityが扱う部分のみ） ===
export interface AuthMeResponse {
  data: UserResponse & {
    metadata: unknown; // metadata-entityで定義予定
    attendance_stats: unknown; // attendance-entityで定義予定
  };
  message: string;
  timestamp: string;
}

// === PUT /users/{userId} レスポンス（統一形式） ===
export interface UserUpdateResponse {
  data: {
    id: string;
    username: string;
    avatar_url: string;
    updated_at: string;
  };
  message: string;
  timestamp: string;
}

// === Client Types（camelCase - フロントエンド用） ===
export interface User {
  id: string;
  clerkId: string;
  email: string;
  username: string;
  avatarUrl: string;
  discordId?: string;
  isActive: boolean;
  createdAt: Date;
  updatedAt: Date;
}

// === DTO Types（API送信用） ===
export interface UpdateUserDto {
  username?: string;
  avatar_url?: string;
}
```

##### 🚫 他のエンティティに委譲する型
- `UserMetadataResponse` → `metadata-entity/model/types.ts`
- `SocialLinkResponse` → `social-links-entity/model/types.ts`
- `AttendanceStatsResponse` → `attendance-entity/model/types.ts`
- `UsersListResponse` → 必要に応じて後で追加（複数entityを組み合わせる場合）
- `UserDetailResponse` → 複数entityの統合が必要な場合

##### APIクライアント (api/user-api.ts) - user-entity範囲のみ
```typescript
import { apiClient } from '@/shared/api';
import type { 
  AuthMeResponse,
  UserUpdateResponse,
  UpdateUserDto 
} from '../model/types';

/** 現在のユーザー情報を取得（認証状態確認） */
export const getCurrentUser = async (): Promise<AuthMeResponse['data']> => {
  const { data } = await apiClient.get<AuthMeResponse>('/auth/me');
  return data.data; // 基本情報 + metadata(unknown) + attendance_stats(unknown)
};

/** ユーザーの基本情報を更新 */
export const updateUser = async ({
  id,
  data: updateData,
}: {
  id: string;
  data: UpdateUserDto;
}): Promise<UserUpdateResponse['data']> => {
  const { data } = await apiClient.put<UserUpdateResponse>(`/users/${id}`, updateData);
  return data.data; // 更新されたユーザー基本情報
};
```

##### 🚫 他のエンティティで実装するAPI
- `getUsers()` → 複数entityの組み合わせが必要な場合に実装
- `getUserDetail()` → 複数entityの組み合わせが必要な場合に実装  
- `getUserMetadata()` → `metadata-entity/api/`で実装
- `updateUserMetadata()` → `metadata-entity/api/`で実装
- `getSocialLinks()` → `social-links-entity/api/`で実装

##### クエリキー管理 (utils/query-keys.ts) - user-entity範囲のみ
```typescript
export const userKeys = {
  all: ['users'] as const,
  // 基本情報更新関連のみ
  update: (id: string) => [...userKeys.all, 'update', id] as const,
} as const;

// 認証関連（Clerkとの区別）
export const authKeys = {
  all: ['auth'] as const,
  me: () => [...authKeys.all, 'me'] as const, // GET /auth/me用
} as const;
```

##### 🚫 他のエンティティで管理するクエリキー
- `userKeys.lists()` → ユーザー一覧系は複数entityの組み合わせ
- `userKeys.detail()` → ユーザー詳細系は複数entityの組み合わせ
- `userKeys.metadata()` → `metadata-entity/utils/query-keys.ts`で管理

##### Zustand Store (model/store.ts) - user-entity基本情報のみ
```typescript
import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import type { User } from './types';

// user-entityが管理する状態（基本情報のみ）
interface UserState {
  // 基本ユーザー情報のみ保持
  currentUser: User | null;
  isAuthenticated: boolean;
  isLoading: boolean;
}

interface UserActions {
  setCurrentUser: (user: User | null) => void;
  updateCurrentUser: (userData: Partial<User>) => void;
  setLoading: (loading: boolean) => void;
  logout: () => void;
}

export const useUserStore = create<UserState & UserActions>()(
  devtools(
    (set) => ({
      // State
      currentUser: null,
      isAuthenticated: false,
      isLoading: false,

      // Actions
      setCurrentUser: (user) =>
        set({
          currentUser: user,
          isAuthenticated: !!user,
          isLoading: false,
        }),

      updateCurrentUser: (userData) =>
        set((state) => ({
          currentUser: state.currentUser
            ? { ...state.currentUser, ...userData }
            : null,
        })),

      setLoading: (loading) => set({ isLoading: loading }),

      logout: () =>
        set({
          currentUser: null,
          isAuthenticated: false,
          isLoading: false,
        }),
    }),
    { name: 'user-store' }
  )
);

// 便利なselectors
export const useCurrentUser = () => useUserStore((state) => state.currentUser);
export const useIsAuthenticated = () => useUserStore((state) => state.isAuthenticated);
```

##### 📝 設計判断の説明
- **基本情報のみ保持**: metadata, attendance_statsは他のentityのStoreで管理
- **シンプルな構造**: user-entityの責務を明確に限定
- **他entityとの協調**: 必要に応じて複数のStoreを組み合わせて使用

##### バリデーション (lib/validators.ts) - user-entity範囲のみ
```typescript
import { z } from 'zod';

/** ユーザー名のバリデーション */
export const usernameSchema = z
  .string()
  .min(3, 'ユーザー名は3文字以上で入力してください')
  .max(50, 'ユーザー名は50文字以内で入力してください')
  .regex(/^[a-zA-Z0-9_-]+$/, '英数字、アンダースコア、ハイフンのみ使用可能です');

/** ユーザー基本情報更新フォームのバリデーション */
export const updateUserFormSchema = z.object({
  username: usernameSchema.optional(),
  avatar_url: z.string().url('有効なURLを入力してください').optional(),
});

/** ユーザーエンティティのスキーマ */
export const userSchema = z.object({
  id: z.string().uuid(),
  clerkId: z.string(),
  email: z.string().email(),
  username: z.string(),
  avatarUrl: z.string().url(),
  discordId: z.string().optional(),
  isActive: z.boolean(),
  createdAt: z.date(),
  updatedAt: z.date(),
});

// Type Guards
export const isUser = (value: unknown): value is User => {
  return userSchema.safeParse(value).success;
};

export const isValidUpdateUserDto = (value: unknown): value is UpdateUserDto => {
  return updateUserFormSchema.safeParse(value).success;
};

// Type exports
export type UpdateUserFormData = z.infer<typeof updateUserFormSchema>;
export type User = z.infer<typeof userSchema>;
```

##### 🚫 他のエンティティで実装するバリデーション
- メールアドレス検証 → `metadata-entity/lib/validators.ts`  
- プロフィール項目検証 → `metadata-entity/lib/validators.ts`
- SNSリンク検証 → `social-links-entity/lib/validators.ts`

#### 3.2 データフロー設計（Clerk + React Query + Zustand）

##### 1. 認証フロー
```
Clerk認証 → getCurrentUser() API → Zustand Store更新 → React Query キャッシュ → UI表示
```

##### 2. プロフィール更新フロー
```
フォーム入力 → Zodバリデーション → updateUser() API → React Query無効化 → Zustand更新 → UI反映
```

##### 3. ユーザー一覧表示フロー
```
ページアクセス → getUsers() API → React Query キャッシュ → ページネーション → UI表示
```

##### 4. 実装時のポイント
- **認証状態**: ClerkとZustandの両方で管理（Clerkは認証、Zustandはユーザー情報）
- **APIキャッシュ**: React Queryで自動管理
- **エラーハンドリング**: shared/apiクライアントで一元処理
- **型安全性**: APIレスポンス → 内部型へのマッピング関数を実装

### 4. 品質保証戦略

#### 4.1 型安全性
- **never型の活用**: 網羅性チェック
- **branded型**: 値の意味を型で表現
- **discriminated union**: 状態の型安全な表現

#### 4.2 バリデーション
- **Zod Schema**: API レスポンスの実行時検証
- **Form Validation**: ユーザー入力の事前検証
- **Type Guards**: 型の実行時チェック

#### 4.3 エラーハンドリング
- **Error Boundary**: 予期せぬエラーのキャッチ
- **Toast Notification**: ユーザーフレンドリーなエラー表示
- **Retry Logic**: 一時的な障害への対応

### 5. パフォーマンス最適化

#### 5.1 React Query 最適化
- **Stale Time**: データの鮮度管理（5分）
- **Cache Time**: キャッシュ保持時間（10分）
- **Background Refetch**: バックグラウンド更新の有効化

#### 5.2 バンドルサイズ最適化
- **Tree Shaking**: 未使用コードの除去
- **Code Splitting**: エンティティごとの分割
- **Dynamic Import**: 必要時のみ読み込み

### 6. 開発フロー

#### 6.1 実装順序
1. **型定義** → **バリデーションスキーマ** → **APIクライアント** → **クエリキー**
2. **単体テスト** → **統合テスト** → **型テスト**
3. **ドキュメント** → **コードレビュー** → **デプロイ**

#### 6.2 品質チェック
- **ESLint**: コード品質チェック
- **TypeScript**: 型チェック
- **Vitest**: 単体テスト
- **MSW**: APIモック

### 7. リスク管理

#### 7.1 技術リスク
- **API仕様変更**: バージョニングによる対応
- **型不整合**: 実行時バリデーションによる検証
- **パフォーマンス劣化**: 監視とアラートの設定

#### 7.2 開発リスク
- **FSD違反**: ESLint ルールによる自動検出
- **型安全性低下**: 厳格なTSConfig設定
- **テストカバレッジ不足**: 最低80%のカバレッジ要求

### 8. 成功指標

#### 8.1 技術指標
- **型安全性**: TypeScript エラー 0件
- **ESLint**: 警告・エラー 0件
- **テストカバレッジ**: 80%以上
- **バンドルサイズ**: 前バージョン比較で10%以内

#### 8.2 品質指標
- **API応答時間**: 平均200ms以下
- **エラー率**: 1%以下
- **ユーザビリティ**: フォーム送信成功率95%以上

## 実装タイムライン（eagle-aiパターン適用版）

### Week 1: User Core Entity（FE-03-user-01~03）
- **Day 1**: User Entity型定義（types.ts, Zodスキーマ）
- **Day 2**: Zustand Store実装（store.ts） - selectorsファイルなしのシンプル構成
- **Day 3**: User API クライアント実装（user-api.ts）
- **Day 4**: クエリキー管理（query-keys.ts）
- **Day 5**: バリデーション・テスト（validators.ts）

### Week 2: User Sub-Entities
- **Day 1-2**: Metadata Entity実装（同様パターン）
- **Day 3**: Social Links Entity実装
- **Day 4**: Rivals Entity実装
- **Day 5**: 統合テスト・バグ修正・ドキュメント

### Week 3-4: 他エンティティ展開
- **Attendance Entity**: eagle-aiパターンを朝活データに適用
- **Goals, Events, Titles, Notifications**: 同様の構造で展開

## 結論（改訂版）- Entity境界の明確化
eagle-aiプロジェクトの実装パターンを参考にしつつ、**FSDの単一責任原則**に基づき、各entityの責務範囲を明確に分離したEntities層を構築する。

### 主な採用パターン
1. **eagle-aiのディレクトリ構造**: api/, model/, lib/, utils/の4層構造
2. **eagle-aiのクエリキー管理**: 階層的なキー構造
3. **eagle-aiのAPI関数パターン**: 名前付き関数エクスポート
4. **Ghoona Camp仕様**: UUID、Clerk認証、統一レスポンス形式
5. **技術スタック適用**: Redux → Zustand、独自validation → Zod
6. **Zustand最適化**: selectorsファイルを廃止し、直接アクセスによるシンプルな状態管理

### 🎯 重要な改善点：Entity責務の明確化
- **user-entity**: 基本ユーザー情報（id, email, username等）のみ
- **metadata-entity**: プロフィール詳細情報（display_name, bio, skills等）
- **attendance-entity**: 出席統計情報
- **social-links-entity**: SNSリンク情報
- **複合的な機能**: features層で複数entityを組み合わせて実装

### 📝 今回の反省点
1. **スコープの明確化不足**: 最初の実装で他entityの責務まで含めてしまった
2. **FSD原則の軽視**: 単一責任原則を守らず、過度に複雑な型定義を作成
3. **境界の曖昧さ**: entityの責務範囲を事前に明確にしていなかった

これらの反省を活かし、**明確に分離された責務**を持つ、保守性の高いEntities層を確立する。

---

*作成日: 2025-01-21*  
*作成者: Claude*  
*バージョン: 1.0*