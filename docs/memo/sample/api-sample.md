# REST API 設計書

## 概要

このドキュメントは、プロジェクト管理システムの REST API 設計を定義します。
各エンドポイントは RESTful な設計原則に従い、適切な HTTP メソッドとステータスコードを使用します。

## 権限管理

### 1. 認証コンテクスト (Auth Context)

- `POST   /login` ログイン **[全ユーザー]**
- `GET    /auth/status` 認証状態確認 **[全ユーザー]**
- `GET    /users` ユーザー一覧取得 **[pm, leader, superuser]**
- `GET    /users/{userId}` ユーザー詳細取得 **[pm, leader, superuser]**
- `POST   /users` ユーザー作成 **[leader, superuser]**
- `PUT    /users/{userId}` ユーザー更新 **[leader, superuser]**
- `DELETE /users/{userId}` ユーザー削除 **[leader, superuser]**
- `PATCH  /users/{userId}/password` パスワード変更 **[本人のみ]**
- `GET    /users/{userId}/assignments` ユーザーのアサインメント一覧取得 **[leader, superuser]**
- `POST   /users/{userId}/roles` ユーザーロール追加 **[leader, superuser]**
- `POST   /users/{userId}/skills` ユーザースキル追加 **[leader, superuser]**
- `PUT    /users/{userId}/skills/{skillId}` ユーザースキル更新 **[leader, superuser]**
- `DELETE /users/{userId}/skills/{skillId}` ユーザースキル削除 **[leader, superuser]**
- `GET    /authorities` 権限一覧取得 **[leader, superuser]**
- `GET    /roles` ロール一覧取得 **[全ユーザー]**
- `POST   /roles` ロール作成 **[leader, superuser]**

### 2. プロジェクトコンテクスト (Project Context)

- `GET    /projects` プロジェクト一覧取得 **[sales, leader, superuser]**
- `GET    /projects/{projectId}` プロジェクト詳細取得 **[sales, pm, leader, superuser]**
- `POST   /projects` プロジェクト作成 **[sales, leader, superuser]**
- `PUT    /projects/{projectId}` プロジェクト更新 **[sales, pm, leader, superuser]**
- `PATCH  /projects/{projectId}` プロジェクトステータス変更 **[sales, pm, leader, superuser]**
- `DELETE /projects/{projectId}` プロジェクト削除 **[leader, superuser]**
- `GET    /projects/categories` プロジェクトカテゴリー一覧取得 **[全ユーザー]**
- `POST   /projects/{projectId}/assignments` アサインメント追加 **[leader, superuser]**
- `PUT    /projects/{projectId}/assignments/{assignmentId}` アサインメント更新 **[leader, superuser]**
- `DELETE /projects/{projectId}/assignments/{assignmentId}` アサインメント削除 **[leader, superuser]**
- `POST   /projects/{projectId}/milestones` マイルストーン追加 **[sales, pm, leader, superuser]**
- `PUT    /projects/{projectId}/milestones/{milestoneId}` マイルストーン更新 **[pm, leader, superuser]**
- `DELETE /projects/{projectId}/milestones/{milestoneId}` マイルストーン削除 **[sales, pm, leader, superuser]**
- `GET    /projects/{projectId}/comments` コメント一覧取得 **[全ユーザー]**
- `POST   /projects/{projectId}/comments` コメント投稿 **[sales, pm, leader, superuser]**
- `PUT    /projects/{projectId}/comments/{commentId}` コメント更新 **[sales, pm, leader, superuser]**
- `DELETE /projects/{projectId}/comments/{commentId}` コメント削除 **[sales, pm, leader, superuser]**
- `GET    /projects/{projectId}/evaluations` 週次評価取得(最新) **[全ユーザー]**
- `POST   /projects/{projectId}/evaluations` 週次評価登録 **[sales, pm, leader, superuser]**
- `GET    /evaluations` 複数プロジェクトの最新評価取得 **[sales, leader, superuser]**
- `POST   /projects/{projectId}/requests` リクエスト作成 **[sales, pm, leader, superuser]**
- `GET    /milestones` マイルストーン一覧取得 **[全ユーザー]**
- `GET    /milestones/{milestoneId}` マイルストーン詳細取得 **[全ユーザー]**
- `POST   /milestones` マイルストーン作成 **[pm, leader, superuser]**
- `PUT    /milestones/{milestoneId}` マイルストーン更新 **[pm, leader, superuser]**
- `DELETE /milestones/{milestoneId}` マイルストーン削除 **[pm, leader, superuser]**
- `GET    /skills` スキル一覧取得 **[全ユーザー]**
- `POST   /skills` スキル作成 **[leader, superuser]**
- `GET    /requests` リクエスト一覧取得 **[leader, superuser]**
- `PATCH  /requests/{requestId}` リクエストステータス更新(確認中・承認・却下) **[leader, superuser]**
- `GET    /document-types` ドキュメントタイプ一覧取得 **[sales, pm, leader, superuser]**
- `POST   /document-types` ドキュメントタイプ作成 **[leader, superuser]**
- `GET    /availabilities` 週次稼働可能時間取得(最新) **[全ユーザー]**
- `GET    /demanded` 週次要求稼働時間取得(最新) **[全ユーザー]**
- `POST   /demanded` 週次要求稼働時間登録 **[pm, leader, superuser]**

## 共通仕様

### ベース URL

```
https://api.eagleai.com/v1
```

### 認証

すべての API はクッキーによる認証が必要です。

```
Cookie: session_id=<session_token>
```

### レスポンス形式

- Content-Type: application/json
- 日時フォーマット: ISO 8601 (YYYY-MM-DDTHH:mm:ssZ)

---

## 0. システム・認証 API

### System API

#### ヘルスチェック

```
GET /health
```

**レスポンス:**

```json
{
  "status": "ok"
}
```

### Authentication API

#### ログイン

```
POST /login
```

**リクエスト:**

```json
{
  "email": "tanaka@example.com",
  "password": "SecurePassword123!"
}
```

**レスポンス:**

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "田中太郎",
    "email": "tanaka@example.com",
    "github_username": "tanaka-taro",
    "authority": {
      "id": 2,
      "name": "pm"
    }
  },
  "message": "ログインに成功しました"
}
```

**Cookie 設定:**

```
Set-Cookie: session_id=<session_token>; HttpOnly; Secure; SameSite=Strict; Max-Age=86400
```

#### 認証状態確認

```
GET /auth/status
```

**レスポンス (認証済み):**

```json
{
  "authenticated": true,
  "user": {
    "id": 1,
    "name": "田中太郎",
    "email": "tanaka@example.com",
    "github_username": "tanaka-taro",
    "authority": {
      "id": 2,
      "name": "pm"
    }
  }
}
```

**レスポンス (未認証):**

```json
{
  "authenticated": false,
  "user": null
}
```

---

## 1. 認証・認可コンテキスト

### Users API

#### ユーザー一覧取得

```
GET /users
```

**クエリパラメータ:**

| パラメータ  | 型      | 必須 | 説明                           |
| ----------- | ------- | ---- | ------------------------------ |
| leader_only | boolean | No   | true の場合、leader のみを返す |

**レスポンス:**

```json
{
  "users": [
    {
      "id": 1,
      "name": "田中太郎",
      "email": "tanaka@example.com",
      "github_username": "tanaka-taro",
      "authority": {
        "id": 2,
        "name": "pm"
      },
      "roles": [
        {
          "id": 1,
          "name": "フロントエンジニア"
        }
      ],
      "skills": [
        {
          "id": 1,
          "name": "React",
          "level": 4
        }
      ],
      "created_at": "2024-01-01T09:00:00Z",
      "updated_at": "2024-01-01T09:00:00Z"
    }
  ],
  "total": 100
}
```

#### ユーザー詳細取得

```
GET /users/{userId}
```

**レスポンス:**

```json
{
  "id": 1,
  "name": "田中太郎",
  "email": "tanaka@example.com",
  "github_username": "tanaka-taro",
  "authority": {
    "id": 2,
    "name": "pm"
  },
  "roles": [
    {
      "id": 1,
      "name": "フロントエンジニア"
    }
  ],
  "skills": [
    {
      "id": 1,
      "name": "React",
      "level": 4
    }
  ],
  "created_at": "2024-01-01T09:00:00Z",
  "updated_at": "2024-01-01T09:00:00Z"
}
```

#### ユーザー作成

```
POST /users
```

**リクエスト:**

```json
{
  "name": "佐藤花子",
  "email": "sato@example.com",
  "password": "SecurePassword123!",
  "authority_id": 1,
  "github_username": "sato-hanako"
}
```

**レスポンス:**

```json
{
  "id": 2,
  "name": "佐藤花子",
  "email": "sato@example.com",
  "github_username": "sato-hanako",
  "authority": {
    "id": 1,
    "name": "member"
  },
  "created_at": "2024-01-02T10:00:00Z",
  "updated_at": "2024-01-02T10:00:00Z"
}
```

**備考:**
- `github_username` フィールドはオプション（null許可）
- GitHubユーザー名はシステム全体でユニーク

#### ユーザー更新

```
PUT /users/{userId}
```

**リクエスト:**

```json
{
  "name": "佐藤花子",
  "email": "sato.hanako@example.com",
  "authority_id": 2,
  "github_username": "sato-hanako"
}
```

**レスポンス:**

```json
{
  "id": 2,
  "name": "佐藤花子",
  "email": "sato.hanako@example.com",
  "github_username": "sato-hanako",
  "authority": {
    "id": 2,
    "name": "pm"
  },
  "created_at": "2024-01-02T10:00:00Z",
  "updated_at": "2024-01-03T11:00:00Z"
}
```

#### ユーザー削除

```
DELETE /users/{userId}
```

**レスポンス:**

```
204 No Content
```

#### パスワード変更

```
PATCH /users/{userId}/password
```

**リクエスト:**

```json
{
  "current_password": "OldPassword123!",
  "new_password": "NewPassword456!"
}
```

**レスポンス:**

```json
{
  "message": "パスワードを変更しました"
}
```

**備考:**

- ユーザーは自分のパスワードのみ変更可能
- current_password の検証が必須
- パスワードは 8 文字以上、大文字・小文字・数字・特殊文字を含む

#### ユーザーのアサインメント一覧取得

```
GET /users/{userId}/assignments
```

**レスポンス:**

```json
{
  "assignments": [
    {
      "id": 1,
      "project": {
        "id": 1,
        "name": "ECサイトリニューアル",
        "status": "in_progress"
      },
      "status": "assigned",
      "created_at": "2024-01-01T09:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
    },
    {
      "id": 2,
      "project": {
        "id": 2,
        "name": "新規CRMシステム開発",
        "status": "potential"
      },
      "status": "tempolary_assigned",
      "created_at": "2024-02-01T09:00:00Z",
      "updated_at": "2024-02-01T09:00:00Z"
    }
  ],
  "total": 2
}
```

### User Roles API

#### ユーザーロール追加

```
POST /users/{userId}/roles
```

**リクエスト:**

```json
{
  "role_id": 2
}
```

**レスポンス:**

```json
{
  "id": 3,
  "user_id": 1,
  "role": {
    "id": 2,
    "name": "バックエンドエンジニア"
  }
}
```

### User Skills API

#### ユーザースキル追加

```
POST /users/{userId}/skills
```

**リクエスト:**

```json
{
  "skill_id": 3,
  "level": 4
}
```

**レスポンス:**

```json
{
  "id": 5,
  "user_id": 1,
  "skill": {
    "id": 3,
    "name": "Docker"
  },
  "level": 4
}
```

#### ユーザースキル更新

```
PUT /users/{userId}/skills/{skillId}
```

**リクエスト:**

```json
{
  "level": 5
}
```

**レスポンス:**

```json
{
  "id": 5,
  "user_id": 1,
  "skill": {
    "id": 3,
    "name": "Docker"
  },
  "level": 5
}
```

#### ユーザースキル削除

```
DELETE /users/{userId}/skills/{skillId}
```

**レスポンス:**

```
204 No Content
```

**備考:**

- ユーザーが保有するスキルを削除
- user_id と skill_id の組み合わせが存在しない場合は 404 エラー

### Authorities API

#### 権限一覧取得

```
GET /authorities
```

**レスポンス:**

```json
{
  "authorities": [
    {
      "id": 1,
      "name": "member"
    },
    {
      "id": 2,
      "name": "pm"
    },
    {
      "id": 3,
      "name": "leader"
    },
    {
      "id": 4,
      "name": "superuser"
    }
  ]
}
```

### Roles API

#### ロール一覧取得

```
GET /roles
```

**レスポンス:**

```json
{
  "roles": [
    {
      "id": 1,
      "name": "フロントエンジニア"
    },
    {
      "id": 2,
      "name": "バックエンドエンジニア"
    },
    {
      "id": 3,
      "name": "インフラエンジニア"
    }
  ]
}
```

#### ロール作成

```
POST /roles
```

**リクエスト:**

```json
{
  "name": "DevOpsエンジニア"
}
```

**レスポンス:**

```json
{
  "id": 4,
  "name": "DevOpsエンジニア"
}
```

**備考:**

- ロール名は重複不可
- 作成後は各ユーザーのロールとして設定可能

### Skills API

#### スキル一覧取得

```
GET /skills
```

**レスポンス:**

```json
{
  "skills": [
    {
      "id": 1,
      "name": "React"
    },
    {
      "id": 2,
      "name": "Node.js"
    },
    {
      "id": 3,
      "name": "Docker"
    }
  ]
}
```

#### スキル作成

```
POST /skills
```

**リクエスト:**

```json
{
  "name": "TypeScript"
}
```

**レスポンス:**

```json
{
  "id": 4,
  "name": "TypeScript"
}
```

### Milestones API

#### マイルストーン一覧取得

```
GET /milestones
```

**レスポンス:**

```json
{
  "milestones": [
    {
      "id": 1,
      "name": "要件定義"
    },
    {
      "id": 2,
      "name": "設計"
    },
    {
      "id": 3,
      "name": "開発"
    },
    {
      "id": 4,
      "name": "テスト"
    },
    {
      "id": 5,
      "name": "リリース"
    }
  ]
}
```

#### マイルストーン詳細取得

```
GET /milestones/{milestoneId}
```

**レスポンス:**

```json
{
  "id": 1,
  "name": "要件定義"
}
```

#### マイルストーン作成

```
POST /milestones
```

**リクエスト:**

```json
{
  "name": "UAT"
}
```

**レスポンス:**

```json
{
  "id": 6,
  "name": "UAT"
}
```

**備考:**

- マイルストーン名は重複不可
- 作成後はプロジェクトのマイルストーンとして設定可能

#### マイルストーン更新

```
PUT /milestones/{milestoneId}
```

**リクエスト:**

```json
{
  "name": "ユーザー受け入れテスト"
}
```

**レスポンス:**

```json
{
  "id": 6,
  "name": "ユーザー受け入れテスト"
}
```

#### マイルストーン削除

```
DELETE /milestones/{milestoneId}
```

**レスポンス:**

```
204 No Content
```

**備考:**

- プロジェクトで使用中のマイルストーンは削除不可
- 削除前に関連するプロジェクトマイルストーンの確認が必要

---

## 2. プロジェクト管理コンテキスト

### Projects API

#### プロジェクト一覧取得

```
GET /projects?status=in_progress&category_id=1
```

**レスポンス:**

```json
{
  "projects": [
    {
      "id": 1,
      "name": "ECサイトリニューアル",
      "description": "既存ECサイトのフルリニューアル",
      "status": "in_progress",
      "category": {
        "id": 1,
        "name": "Web開発"
      },
      "man_month": 12,
      "created_at": "2024-01-01T09:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
    }
  ],
  "total": 25
}
```

#### プロジェクト詳細取得

```
GET /projects/{projectId}
```

**レスポンス:**

```json
{
  "id": 1,
  "name": "ECサイトリニューアル",
  "description": "既存ECサイトのフルリニューアル",
  "status": "in_progress",
  "category": {
    "id": 1,
    "name": "Web開発",
    "description": "Webアプリケーション開発プロジェクト"
  },
  "man_month": 12,
  "notion_link": "https://notion.so/project/123",
  "assignments": [
    {
      "id": 1,
      "user": {
        "id": 1,
        "name": "田中太郎"
      },
      "status": "assigned"
    }
  ],
  "milestones": [
    {
      "id": 1,
      "milestone": {
        "id": 1,
        "name": "要件定義"
      },
      "start_date": "2024-01-01",
      "end_date": "2024-01-31",
      "due_date": "2024-01-31"
    }
  ],
  "created_at": "2024-01-01T09:00:00Z",
  "updated_at": "2024-01-15T10:00:00Z"
}
```

#### プロジェクト作成

```
POST /projects
```

**リクエスト:**

```json
{
  "name": "新規CRMシステム開発",
  "description": "社内向けCRMシステムの新規開発",
  "project_category_id": 1,
  "man_month": 8,
  "notion_link": "https://notion.so/project/456"
}
```

**レスポンス:**

```json
{
  "id": 2,
  "name": "新規CRMシステム開発",
  "description": "社内向けCRMシステムの新規開発",
  "status": "potential",
  "category": {
    "id": 1,
    "name": "Web開発"
  },
  "man_month": 8,
  "notion_link": "https://notion.so/project/456",
  "created_at": "2024-02-01T09:00:00Z",
  "updated_at": "2024-02-01T09:00:00Z"
}
```

#### プロジェクト更新

```
PUT /projects/{projectId}
```

**リクエスト:**

```json
{
  "name": "新規CRMシステム開発",
  "description": "社内向けCRMシステムの新規開発(フェーズ1)",
  "man_month": 10
}
```

#### プロジェクトステータス変更

```
PATCH /projects/{projectId}
```

**リクエスト:**

```json
{
  "status": "in_progress"
}
```

**レスポンス:**

```json
{
  "id": 2,
  "name": "新規CRMシステム開発",
  "description": "社内向けCRMシステムの新規開発(フェーズ1)",
  "status": "in_progress",
  "category": {
    "id": 1,
    "name": "Web開発"
  },
  "man_month": 10,
  "notion_link": "https://notion.so/project/456",
  "created_at": "2024-02-01T09:00:00Z",
  "updated_at": "2024-02-02T10:00:00Z"
}
```

#### プロジェクト削除

```
DELETE /projects/{projectId}
```

**レスポンス:**

```
204 No Content
```

### Project Assignments API

#### アサインメント追加

```
POST /projects/{projectId}/assignments
```

**リクエスト:**

```json
[
  {
    "user_id": 2,
    "status": "tempolary_assigned"
  },
  {
    "user_id": 3,
    "status": "assigned"
  }
]
```

**レスポンス:**

```json
[
  {
    "id": 3,
    "project_id": 1,
    "user": {
      "id": 2,
      "name": "佐藤花子"
    },
    "status": "tempolary_assigned"
  },
  {
    "id": 4,
    "project_id": 1,
    "user": {
      "id": 3,
      "name": "鈴木一郎"
    },
    "status": "assigned"
  }
]
```

**備考:**

- 複数のアサインメントを一括で追加可能
- 配列でリクエストを送信し、配列でレスポンスを返す

#### アサインメント更新

```
PUT /projects/{projectId}/assignments/{assignmentId}
```

**リクエスト:**

```json
{
  "status": "assigned"
}
```

**レスポンス:**

```json
{
  "id": 3,
  "project_id": 1,
  "user": {
    "id": 2,
    "name": "佐藤花子"
  },
  "status": "assigned"
}
```

#### アサインメント削除

```
DELETE /projects/{projectId}/assignments/{assignmentId}
```

**レスポンス:**

```
204 No Content
```

### Project Milestones API

#### マイルストーン追加

```
POST /projects/{projectId}/milestones
```

**リクエスト:**

```json
{
  "milestone_id": 2,
  "start_date": "2024-02-01",
  "end_date": "2024-03-31",
  "due_date": "2024-03-31"
}
```

**レスポンス:**

```json
{
  "id": 2,
  "project_id": 1,
  "milestone": {
    "id": 2,
    "name": "設計"
  },
  "start_date": "2024-02-01",
  "end_date": "2024-03-31",
  "due_date": "2024-03-31"
}
```

#### マイルストーン更新

```
PUT /projects/{projectId}/milestones/{milestoneId}
```

**リクエスト:**

```json
{
  "start_date": "2024-02-01",
  "end_date": "2024-04-15",
  "due_date": "2024-04-15"
}
```

**レスポンス:**

```json
{
  "id": 2,
  "project_id": 1,
  "milestone": {
    "id": 2,
    "name": "設計"
  },
  "start_date": "2024-02-01",
  "end_date": "2024-04-15",
  "due_date": "2024-04-15"
}
```

#### マイルストーン削除

```
DELETE /projects/{projectId}/milestones/{milestoneId}
```

**レスポンス:**

```
204 No Content
```

### Project Comments API

#### コメント一覧取得

```
GET /projects/{projectId}/comments
```

**レスポンス:**

```json
{
  "comments": [
    {
      "id": 1,
      "project_id": 1,
      "user": {
        "id": 1,
        "name": "田中太郎"
      },
      "comment": "進捗順調です。",
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
    }
  ],
  "total": 5
}
```

#### コメント投稿

```
POST /projects/{projectId}/comments
```

**リクエスト:**

```json
{
  "comment": "仕様について確認事項があります。"
}
```

**レスポンス:**

```json
{
  "id": 2,
  "project_id": 1,
  "user": {
    "id": 2,
    "name": "佐藤花子"
  },
  "comment": "仕様について確認事項があります。",
  "created_at": "2024-01-16T11:00:00Z",
  "updated_at": "2024-01-16T11:00:00Z"
}
```

#### コメント更新

```
PUT /projects/{projectId}/comments/{commentId}
```

**リクエスト:**

```json
{
  "comment": "仕様について確認事項があります。詳細は別途メールで送付しました。"
}
```

**レスポンス:**

```json
{
  "id": 2,
  "project_id": 1,
  "user": {
    "id": 2,
    "name": "佐藤花子"
  },
  "comment": "仕様について確認事項があります。詳細は別途メールで送付しました。",
  "created_at": "2024-01-16T11:00:00Z",
  "updated_at": "2024-01-16T12:00:00Z"
}
```

#### コメント削除

```
DELETE /projects/{projectId}/comments/{commentId}
```

**レスポンス:**

```
204 No Content
```

### Weekly Schedule Evaluations API

#### 週次評価取得(最新)

```
GET /projects/{projectId}/evaluations
```

**レスポンス:**

```json
{
  "evaluations": [
    {
      "id": 1,
      "project_id": 1,
      "evaluation": 4,
      "week_start_date": "2024-01-01",
      "created_at": "2024-01-07T17:00:00Z",
      "updated_at": "2024-01-07T17:00:00Z"
    },
    {
      "id": 2,
      "project_id": 1,
      "evaluation": 3,
      "week_start_date": "2024-01-08",
      "created_at": "2024-01-14T17:00:00Z",
      "updated_at": "2024-01-14T17:00:00Z"
    }
  ]
}
```

#### 週次評価登録

```
POST /projects/{projectId}/evaluations
```

**リクエスト:**

```json
{
  "evaluation": 4
}
```

**レスポンス:**

```json
{
  "id": 3,
  "project_id": 1,
  "evaluation": 4,
  "week_start_date": "2024-01-15",
  "created_at": "2024-01-21T17:00:00Z",
  "updated_at": "2024-01-21T17:00:00Z"
}
```

## Evaluations API

### 複数プロジェクトの最新評価取得

```
GET /evaluations?project_id=1,2,3,4
```

**クエリパラメータ:**

- `project_id` (required): カンマ区切りのプロジェクト ID

**レスポンス:**

```json
[
  {
    "project_id": 1,
    "evaluation": 5
  },
  {
    "project_id": 2,
    "evaluation": 3
  },
  {
    "project_id": 3,
    "evaluation": 4
  },
  {
    "project_id": 4,
    "evaluation": 2
  }
]
```

**エラーレスポンス:**

- 400 Bad Request: `project_id` パラメータが未指定またはフォーマットが不正

```json
{
  "error": "project_id parameter is required"
}
```

```json
{
  "error": "invalid number: abc"
}
```

### Project Categories API

#### プロジェクトカテゴリー一覧取得

```
GET /projects/categories
```

**レスポンス:**

```json
{
  "categories": [
    {
      "id": 1,
      "name": "Web開発",
      "description": "Webアプリケーション開発プロジェクト",
      "in_progress_projects_count": 5,
      "potential_projects_count": 3
    },
    {
      "id": 2,
      "name": "モバイル開発",
      "description": "iOS/Androidアプリ開発プロジェクト",
      "in_progress_projects_count": 2,
      "potential_projects_count": 1
    }
  ]
}
```

---

## 3. 稼働管理コンテキスト

### Weekly Availabilities API

#### 週次稼働可能時間取得(最新)

```
GET /availabilities?user_id=1
GET /availabilities?user_id=1,2,3
```

**レスポンス:**

```json
{
  "availabilities": [
    {
      "id": 1,
      "user": {
        "id": 1,
        "name": "田中太郎"
      },
      "week_start_date": "2024-01-01",
      "available_hours": 40.0,
      "working_hours": 38.5,
      "created_at": "2024-01-01T09:00:00Z",
      "updated_at": "2024-01-07T18:00:00Z"
    }
  ]
}
```

#### 週次稼働可能時間登録

```
POST /availabilities
```

**リクエスト:**

```json
{
  "user_email": "tanaka@example.com",
  "week_start_date": "2024-01-01",
  "available_hours": 40.0
}
```

**レスポンス:**

```json
{
  "id": 2,
  "user": {
    "id": 1,
    "name": "田中太郎",
    "email": "tanaka@example.com"
  },
  "week_start_date": "2024-01-01",
  "available_hours": 40.0,
  "working_hours": 38.5,
  "created_at": "2024-01-01T09:00:00Z",
  "updated_at": "2024-01-01T09:00:00Z"
}
```

**備考:**
- このエンドポイントは主にGoogle FormからGAS経由で呼び出される
- `user_email`からユーザーIDを取得し、WeeklyAvailabilityレコードを作成
- `working_hours`は既存の値がある場合は引き継ぎ、ない場合は0で初期化
- 同一ユーザー・同一週の重複登録は既存レコードを更新
```

### Weekly Demanded API

#### 週次要求稼働時間取得(最新)

```
GET /demanded?assignment_id=1
```

**レスポンス:**

```json
{
  "demanded": [
    {
      "id": 1,
      "assignment": {
        "id": 1,
        "user": {
          "id": 1,
          "name": "田中太郎"
        },
        "project": {
          "id": 1,
          "name": "ECサイトリニューアル"
        }
      },
      "week_start_date": "2024-01-01",
      "demanded_hours": 30.0,
      "created_at": "2024-01-01T09:00:00Z",
      "updated_at": "2024-01-01T09:00:00Z"
    }
  ]
}
```

#### 週次要求稼働時間登録

```
POST /demanded
```

**リクエスト:**

```json
{
  "assignment_id": 1,
  "demanded_hours": 35.0
}
```

**レスポンス:**

```json
{
  "id": 2,
  "assignment": {
    "id": 1,
    "user": {
      "id": 1,
      "name": "田中太郎"
    },
    "project": {
      "id": 1,
      "name": "ECサイトリニューアル"
    }
  },
  "week_start_date": "2024-01-08",
  "demanded_hours": 35.0,
  "created_at": "2024-01-08T09:00:00Z",
  "updated_at": "2024-01-08T09:00:00Z"
}
```

---

## 4. 承認ワークフローコンテキスト

### Requests API

#### リクエスト一覧取得

```
GET /requests
```

**クエリパラメータ:**

| パラメータ | 型      | 必須 | 説明                                                          | デフォルト   |
| ---------- | ------- | ---- | ------------------------------------------------------------- | ------------ |
| status     | string  | No   | フィルタするステータス (pending, reviewing, approved, denied) <br/>**複数指定可能:** `status=pending&status=reviewing` または `status=pending,reviewing` | なし(全件) |
| page       | integer | No   | ページ番号(1 から開始)                                      | 1            |
| limit      | integer | No   | 1 ページあたりの件数(最大 100)                              | 10           |

**リクエスト例:**

```
# 単一ステータス
GET /requests?status=pending&page=2&limit=10

# 下記、どちらを使っても問題ないです。
# 複数ステータス(URL パラメータ形式)
GET /requests?status=pending&status=reviewing&page=1&limit=20

# 複数ステータス(カンマ区切り形式)
GET /requests?status=pending,reviewing&page=1&limit=20
```

**レスポンス:**

```json
{
  "requests": [
    {
      "id": 1,
      "title": "見積書承認依頼",
      "created_by": {
        "id": 1,
        "name": "田中太郎"
      },
      "approved_by": null,
      "description": "新規プロジェクトの見積書承認をお願いします。",
      "status": "pending",
      "document_type": {
        "id": 1,
        "name": "見積書"
      },
      "deadline": "2024-01-31T23:59:59Z",
      "notion_link": "https://notion.so/doc/789",
      "created_at": "2024-01-15T09:00:00Z",
      "updated_at": "2024-01-15T09:00:00Z"
    }
  ],
  "total": 45,
  "page": 2,
  "limit": 10,
  "total_pages": 5
}
```

**レスポンスフィールド:**

| フィールド  | 型      | 説明                 |
| ----------- | ------- | -------------------- |
| requests    | array   | リクエストの配列 (期限が近いもの順、期限なしは作成日降順)     |
| total       | integer | 条件に合致する全件数 |
| page        | integer | 現在のページ番号     |
| limit       | integer | 1 ページあたりの件数 |
| total_pages | integer | 総ページ数           |

**ソート順:**
- 期限(deadline)がある場合: 期限が近いもの順 (昇順)
- 期限がない場合: 作成日時の降順 (新しいもの順)

#### リクエスト作成

```
POST /projects/{projectId}/requests
```

**リクエスト:**

```json
{
  "title": "契約書承認依頼",
  "description": "A社との業務委託契約書の承認をお願いします。",
  "document_type_id": 3,
  "deadline": "2024-02-15T23:59:59Z",
  "notion_link": "https://notion.so/doc/999",
  "approved_by": 1 // オプション: 承認者ID(leaderまたはsuperuserのみ指定可能)
}
```

**レスポンス:**

```json
{
  "id": 2,
  "project_id": 1,
  "title": "契約書承認依頼",
  "created_by": {
    "id": 1,
    "name": "田中太郎"
  },
  "approved_by": 1,
  "description": "A社との業務委託契約書の承認をお願いします。",
  "status": "pending",
  "document_type": {
    "id": 3,
    "name": "契約書"
  },
  "deadline": "2024-02-15T23:59:59Z",
  "notion_link": "https://notion.so/doc/999",
  "created_at": "2024-01-20T10:00:00Z",
  "updated_at": "2024-01-20T10:00:00Z"
```

#### リクエストステータス更新(承認・却下)

```
PATCH /requests/{requestId}
```

**リクエスト(確認中の場合):**

```json
{
  "status": "reviewing"
}
```

**リクエスト(承認の場合):**

```json
{
  "status": "approved"
}
```

**リクエスト(却下の場合):**

```json
{
  "status": "denied"
}
```

**レスポンス:**

```json
{
  "id": 2,
  "project_id": 1,
  "title": "契約書承認依頼",
  "created_by": {
    "id": 1,
    "name": "田中太郎"
  },
  "approved_by": {
    "id": 3,
    "name": "山田部長"
  },
  "description": "A社との業務委託契約書の承認をお願いします。",
  "status": "approved",
  "document_type": {
    "id": 3,
    "name": "契約書"
  },
  "deadline": "2024-02-15T23:59:59Z",
  "notion_link": "https://notion.so/doc/999",
  "created_at": "2024-01-20T10:00:00Z",
  "updated_at": "2024-01-21T11:00:00Z"
}
```

### Document Types API

#### ドキュメントタイプ一覧取得

```
GET /document-types
```

**レスポンス:**

```json
{
  "document_types": [
    {
      "id": 1,
      "name": "見積書"
    },
    {
      "id": 2,
      "name": "要件定義書"
    },
    {
      "id": 3,
      "name": "契約書"
    }
  ]
}
```

#### ドキュメントタイプ作成

```
POST /document-types
```

**リクエスト:**

```json
{
  "name": "仕様書"
}
```

**レスポンス:**

```json
{
  "id": 4,
  "name": "仕様書"
}
```

**備考:**

- ドキュメントタイプ名は重複不可
- 作成後は承認リクエストのドキュメントタイプとして設定可能
