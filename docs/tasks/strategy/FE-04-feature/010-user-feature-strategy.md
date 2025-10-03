# User Feature 実装戦略

## 概要
FSDアーキテクチャに基づき、既存のuser-entity群を活用してfeature層を実装する戦略書。

## 現状分析
- **entities層**: 完全実装済み（API、型定義、バリデーション、クエリキー）
- **features層**: ディレクトリ構造のみ、実装なし

## 実装方針

### 1. 実装対象ディレクトリ構造（lib層のみ）
```
frontend/src/features/user/
│
├── profile-feature/                      # ユーザー基本情報管理
│   ├── profile-get/                     # プロフィール取得
│   │   ├── index.ts
│   │   └── lib/
│   │       └── use-profile-get.ts       # useGetProfile()
│   ├── profile-update/                  # プロフィール更新
│   │   ├── index.ts
│   │   └── lib/
│   │       └── use-profile-update.ts    # useUpdateProfile()
│   └── session-get/                     # 認証セッション取得
│       ├── index.ts
│       └── lib/
│           └── use-session-get.ts       # useGetSession()
│
├── metadata-feature/                    # 詳細プロフィール管理
│   ├── index.ts
│   ├── metadata-get/                    # メタデータ取得
│   │   ├── index.ts
│   │   └── lib/
│   │       └── use-metadata-get.ts      # useGetMetadata()
│   └── metadata-update/                 # メタデータ更新
│       ├── index.ts
│       ├── lib/
│       │   └── use-metadata-update.ts   # useUpdateMetadata()
│       └── model/
│           └── form-types.ts            # フォーム専用型
│
├── user-feature/                        # ユーザー検索・一覧・詳細
│   ├── users-list/                      # ユーザー一覧・検索
│   │   ├── index.ts
│   │   └── lib/
│   │       └── use-users-list.ts        # useGetUsersList() + 検索・フィルター
│   └── user-detail-get/                 # 他ユーザー詳細表示
│       ├── index.ts
│       └── lib/
│           └── use-user-detail.ts       # useGetUserDetail()
│
├── rivals-feature/                      # ライバル関係管理
│   ├── rivals-get/                      # ライバル一覧取得
│   │   ├── index.ts
│   │   └── lib/
│   │       └── use-rivals-get.ts        # useGetRivals()
│   ├── rivals-create/                   # ライバル追加
│   │   ├── index.ts
│   │   └── lib/
│   │       └── use-rivals-create.ts     # useCreateRival()
│   └── rivals-delete/                   # ライバル削除
│       ├── index.ts
│       └── lib/
│           └── use-rivals-delete.ts     # useDeleteRival()
│
└── social-links-feature/                # SNSリンク管理
    ├── links-get/                       # リンク一覧取得
    │   ├── index.ts
    │   └── lib/
    │       └── use-links-get.ts         # useGetSocialLinks()
    ├── links-create/                    # リンク追加
    │   ├── index.ts
    │   └── lib/
    │       └── use-links-create.ts      # useCreateSocialLink()
    ├── links-update/                    # リンク更新
    │   ├── index.ts
    │   └── lib/
    │       └── use-links-update.ts      # useUpdateSocialLink()
    └── links-delete/                    # リンク削除
        ├── index.ts
        └── lib/
            └── use-links-delete.ts      # useDeleteSocialLink()
```

### 2. 各featureの責務
- **profile-feature**: 基本情報取得・更新、認証状態管理
- **metadata-feature**: プロフィール詳細の取得・編集
- **user-feature**: ユーザー検索・一覧・他ユーザー詳細表示
- **rivals-feature**: ライバル追加・削除・一覧表示（最大3人制限）
- **social-links-feature**: SNSリンクのCRUD操作（プラットフォーム制限）

### 3. 実装パターン（lib層のみ）
各featureスライスで以下を提供：
- **lib/**: React Queryベースのカスタムフック
- **model/**: フォーム型定義（必要に応じて）
- **ui/**: 後回し（当面実装対象外）

### 4. 依存関係
```
features/user/* → entities/user/* → shared/*
```

## 実装優先順位
1. **profile-feature** (認証・基本情報) - 最優先
   - session-get: GET /auth/me
   - profile-get: GET /users/{userId} 
   - profile-update: PUT /users/{userId}

2. **metadata-feature** (プロフィール編集) - 高優先
   - metadata-get: GET /users/{userId}/metadata
   - metadata-update: PUT /users/{userId}/metadata

3. **user-feature** (ユーザー一覧・検索・詳細) - 中優先
   - users-list: GET /users (検索・フィルター含む)
   - user-detail-get: GET /users/{userId} (詳細表示用)

4. **rivals-feature** (ライバル管理) - 中優先
   - rivals-get: GET /users/{userId}/rivals
   - rivals-create: POST /users/{userId}/rivals (最大3人制限)
   - rivals-delete: DELETE /users/{userId}/rivals/{rivalId}

5. **social-links-feature** (SNSリンク) - 低優先
   - links-get: GET /users/{userId}/social-links
   - links-create: POST /users/{userId}/social-links
   - links-update: PUT /users/{userId}/social-links/{linkId}
   - links-delete: DELETE /users/{userId}/social-links/{linkId}

## 技術仕様
- **データ取得**: React Query + use hook (React 19)
- **フォーム**: react-hook-form + Zod
- **状態管理**: React Query cache
- **エラーハンドリング**: Error Boundary + toast

---
_作成日: 2025-01-28_