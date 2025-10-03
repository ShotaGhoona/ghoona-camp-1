# Profile Feature 実装戦略

## 概要
ユーザー基本情報管理を担うprofile-featureの詳細実装戦略書。
認証状態管理、プロフィール取得・更新機能を提供する。

## 対象スライス

### 1. session-get
**責務**: 認証セッション管理・現在ユーザー情報取得
- **API**: `GET /auth/me`
- **権限**: 🔐 認証済み
- **用途**: ログイン状態確認、プロフィール設定画面の初期表示

### 2. profile-get  
**責務**: ユーザー基本情報取得（自分・他ユーザー両対応）
- **API**: `GET /users/{userId}`
- **権限**: 🔐 認証済み
- **用途**: プロフィール表示、他ユーザー情報表示

### 3. profile-update
**責務**: ユーザー基本情報更新
- **API**: `PUT /users/{userId}`
- **権限**: 👤 本人のみ
- **用途**: ユーザー名・アバター画像の更新

## 実装詳細

### session-get/lib/use-session-get.ts
```typescript
import { use, useMemo } from 'react';
import { getCurrentUser } from '@/entities/user';

export const useGetSession = () => {
  const sessionPromise = useMemo(() => getCurrentUser(), []);
  
  const sessionData = use(sessionPromise);
  
  return {
    user: sessionData,
    isAuthenticated: !!sessionData,
  };
};
```

### profile-get/lib/use-profile-get.ts
```typescript
import { use, useMemo } from 'react';
import { getUserDetail } from '@/entities/user';

export const useGetProfile = (userId: string) => {
  const profilePromise = useMemo(
    () => getUserDetail(userId),
    [userId]
  );
  
  const profileData = use(profilePromise);
  
  return {
    profile: profileData,
  };
};
```

### profile-update/lib/use-profile-update.ts
```typescript
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateUser, userKeys, type UpdateUserDto } from '@/entities/user';

export const useUpdateProfile = (userId: string) => {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (data: UpdateUserDto) => 
      updateUser({ id: userId, data }),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: userKeys.detail(userId),
      });
      queryClient.invalidateQueries({
        queryKey: userKeys.session(),
      });
    },
  });
};
```

## データフロー

### 認証フロー
```
1. useGetSession() 実行
2. getCurrentUser() API呼び出し
3. 認証状態判定 (isAuthenticated)
4. ユーザー情報返却
```

### プロフィール取得フロー
```
1. useGetProfile(userId) 実行
2. getUserDetail(userId) API呼び出し
3. プロフィールデータ返却
4. UI表示
```

### プロフィール更新フロー
```
1. useUpdateProfile(userId) 実行
2. mutate(updateData) 実行
3. updateUser() API呼び出し
4. キャッシュ無効化
5. 更新完了
```

## API連携仕様

### GET /auth/me レスポンス処理
```typescript
// entities/user/user-api.ts からの変換
interface AuthMeResponse {
  data: {
    id: string;
    clerk_id: string;
    email: string;
    username: string;
    avatar_url: string;
    // ... その他フィールド
  };
}
```

### GET /users/{userId} レスポンス処理
```typescript
// 詳細プロフィール情報取得
interface UserDetailResponse {
  data: {
    id: string;
    display_name: string;
    username: string;
    avatar_url: string;
    // ... その他フィールド
  };
}
```

### PUT /users/{userId} リクエスト処理
```typescript
// 更新可能フィールド
interface UpdateUserDto {
  username?: string;
  avatar_url?: string;
}
```

## エラーハンドリング

### 認証エラー
- 401 Unauthorized → ログイン画面リダイレクト
- 403 Forbidden → 権限エラー表示

### ネットワークエラー
- タイムアウト → リトライ機能
- 接続エラー → オフライン状態表示

### バリデーションエラー
- 422 Unprocessable Entity → フィールド別エラー表示

## キャッシュ戦略

### session-get
- **キャッシュ時間**: 5分
- **無効化タイミング**: ログアウト時、プロフィール更新時

### profile-get
- **キャッシュ時間**: 1分
- **無効化タイミング**: プロフィール更新時

### profile-update
- **キャッシュ無効化**: session, profile-get両方

## 使用例

### 認証ガード
```typescript
const { user, isAuthenticated } = useGetSession();

if (!isAuthenticated) {
  return <LoginRequired />;
}
```

### プロフィール表示
```typescript
const { profile } = useGetProfile(userId);

return (
  <div>
    <img src={profile.avatar_url} />
    <h1>{profile.display_name}</h1>
  </div>
);
```

### プロフィール編集
```typescript
const updateProfile = useUpdateProfile(userId);

const handleUpdate = (data: UpdateUserDto) => {
  updateProfile.mutate(data);
};
```

## 実装順序
1. **session-get**: 認証基盤として最優先
2. **profile-get**: セッション依存機能
3. **profile-update**: CRUD完成

---
_作成日: 2025-01-28_