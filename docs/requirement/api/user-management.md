# User Management API - 詳細設計

## Overview
Ghoona Campアプリケーションのユーザー管理関連API詳細設計書です。
Clerk認証とユーザープロフィール管理を中心とした機能を提供します。

## アーキテクチャ適合性
このAPI設計は以下のアーキテクチャ原則に従います：
- **オニオンアーキテクチャ**: Domain → Application → Infrastructure → Interface層の依存関係
- **ドメイン駆動設計**: `user`コンテキストを中心としたエンティティ設計
- **クリーンな境界**: DTO、ドメインエンティティ、GORMモデルの適切な分離

## ドメイン設計対応

### ドメインエンティティ
```go
// internal/domain/user/entity/user.go
type User struct {
    ID        UUID
    ClerkID   string
    Email     string
    Username  string
    AvatarURL string
    DiscordID *string
    IsActive  bool
    CreatedAt time.Time
    UpdatedAt time.Time
}

// internal/domain/user/entity/user_metadata.go
type UserMetadata struct {
    ID               UUID
    UserID           UUID
    DisplayName      string
    ProfileImageURL  string
    Tagline          string
    Bio              string
    Vision           string
    VisionPublic     bool
    Timezone         string
    Skills           []string
    Interests        []string
    CreatedAt        time.Time
    UpdatedAt        time.Time
}

// internal/domain/user/entity/user_social_link.go
type UserSocialLink struct {
    ID        UUID
    UserID    UUID
    Platform  Platform
    URL       string
    Title     string
    IsPublic  bool
    CreatedAt time.Time
    UpdatedAt time.Time
}

// internal/domain/user/entity/user_rival.go
type UserRival struct {
    ID          UUID
    UserID      UUID
    RivalUserID UUID
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// internal/domain/user/value/platform.go
type Platform string
const (
    PlatformTwitter   Platform = "twitter"
    PlatformGitHub    Platform = "github"
    PlatformLinkedIn  Platform = "linkedin"
    PlatformWebsite   Platform = "website"
    PlatformBlog      Platform = "blog"
)
```

### リポジトリインターフェース
```go
// internal/domain/user/repository/user_repository.go
type UserRepository interface {
    GetByID(ctx context.Context, id UUID) (*entity.User, error)
    GetByClerkID(ctx context.Context, clerkID string) (*entity.User, error)
    GetByEmail(ctx context.Context, email string) (*entity.User, error)
    GetAll(ctx context.Context, filters UserFilters) ([]*entity.User, error)
    Create(ctx context.Context, user *entity.User) error
    Update(ctx context.Context, user *entity.User) error
    Delete(ctx context.Context, id UUID) error
}

type UserMetadataRepository interface {
    GetByUserID(ctx context.Context, userID UUID) (*entity.UserMetadata, error)
    Create(ctx context.Context, metadata *entity.UserMetadata) error
    Update(ctx context.Context, metadata *entity.UserMetadata) error
}

type UserSocialLinkRepository interface {
    GetByUserID(ctx context.Context, userID UUID) ([]*entity.UserSocialLink, error)
    GetByID(ctx context.Context, id UUID) (*entity.UserSocialLink, error)
    Create(ctx context.Context, link *entity.UserSocialLink) error
    Update(ctx context.Context, link *entity.UserSocialLink) error
    Delete(ctx context.Context, id UUID) error
}

type UserRivalRepository interface {
    GetByUserID(ctx context.Context, userID UUID) ([]*entity.UserRival, error)
    Create(ctx context.Context, rival *entity.UserRival) error
    Delete(ctx context.Context, id UUID) error
    CountByUserID(ctx context.Context, userID UUID) (int, error)
}
```

## Base URL
```
https://api.ghoona-camp.com/v1
```

## Authentication
```
Authorization: Bearer <clerk_token>
```

---

## 1. 認証 API

### 現在のユーザー情報取得

```
GET /auth/me
```

**権限**: 🔐 認証済み

**説明**: 現在ログイン中のユーザーの基本情報とメタデータを取得

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "clerk_id": "user_2NiQOyOGrKhcgQjEkRMfPqLj9zF",
    "email": "tanaka@example.com",
    "username": "tanaka_taro",
    "avatar_url": "https://img.clerk.com/preview.png",
    "discord_id": "123456789012345678",
    "is_active": true,
    "metadata": {
      "display_name": "田中太郎",
      "profile_image_url": "https://example.com/profile/tanaka.jpg",
      "tagline": "朝活で人生変える！",
      "bio": "エンジニアとして働きながら、朝活を通じて自己成長を目指しています。",
      "vision": "2025年末までにフルスタックエンジニアになる",
      "vision_public": false,
      "timezone": "Asia/Tokyo",
      "skills": ["JavaScript", "React", "Node.js"],
      "interests": ["プログラミング", "読書", "筋トレ"]
    },
    "attendance_stats": {
      "total_attendance_days": 45,
      "current_streak_days": 7,
      "max_streak_days": 15,
      "current_title": {
        "id": "550e8400-e29b-41d4-a716-446655440003",
        "level": 3,
        "name_jp": "陽光探求者",
        "name_en": "Dawn Seeker",
        "color_theme": "#FFB800"
      }
    },
    "created_at": "2024-12-01T10:30:00Z",
    "updated_at": "2024-12-15T08:45:00Z"
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

**エラーレスポンス:**

```json
{
  "error": {
    "code": "UNAUTHORIZED",
    "message": "Invalid or expired token"
  },
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 2. ユーザー一覧・検索 API

### ユーザー一覧取得

```
GET /users
```

**権限**: 🔐 認証済み

**説明**: 全ユーザーの一覧を取得。メンバー検索・ランキング表示で使用

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| page | integer | No | ページ番号 (デフォルト: 1) |
| limit | integer | No | 1ページあたりの件数 (デフォルト: 20, 最大: 100) |
| search | string | No | 表示名・ユーザー名での部分検索 |
| skills | string | No | スキルでのフィルタ (カンマ区切り) |
| interests | string | No | 興味でのフィルタ (カンマ区切り) |
| sort_by | string | No | ソート基準 (name, attendance_days, streak_days, created_at) |
| order | string | No | ソート順 (asc, desc) デフォルト: desc |

**レスポンス:**

```json
{
  "data": {
    "users": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "display_name": "田中太郎",
        "username": "tanaka_taro",
        "avatar_url": "https://img.clerk.com/preview.png",
        "tagline": "朝活で人生変える！",
        "skills": ["JavaScript", "React", "Node.js"],
        "interests": ["プログラミング", "読書"],
        "current_title": {
          "level": 3,
          "name_jp": "陽光探求者",
          "name_en": "Dawn Seeker",
          "color_theme": "#FFB800"
        },
        "attendance_stats": {
          "total_attendance_days": 45,
          "current_streak_days": 7
        },
        "created_at": "2024-12-01T10:30:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 5,
      "total_count": 98,
      "limit": 20,
      "has_next": true,
      "has_prev": false
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### ユーザー詳細取得

```
GET /users/{userId}
```

**権限**: 🔐 認証済み

**説明**: 指定したユーザーの詳細情報を取得。プロフィールポップアップ表示で使用

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| userId | UUID | Yes | ユーザーID |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "display_name": "田中太郎",
    "username": "tanaka_taro",
    "avatar_url": "https://img.clerk.com/preview.png",
    "profile_image_url": "https://example.com/profile/tanaka.jpg",
    "tagline": "朝活で人生変える！",
    "bio": "エンジニアとして働きながら、朝活を通じて自己成長を目指しています。",
    "vision": "2025年末までにフルスタックエンジニアになる",
    "timezone": "Asia/Tokyo",
    "skills": ["JavaScript", "React", "Node.js"],
    "interests": ["プログラミング", "読書", "筋トレ"],
    "social_links": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440010",
        "platform": "github",
        "url": "https://github.com/tanaka-taro",
        "title": "GitHub Profile",
        "is_public": true
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440011",
        "platform": "twitter",
        "url": "https://twitter.com/tanaka_dev",
        "title": "Twitter",
        "is_public": true
      }
    ],
    "current_title": {
      "id": "550e8400-e29b-41d4-a716-446655440003",
      "level": 3,
      "name_jp": "陽光探求者",
      "name_en": "Dawn Seeker",
      "description": "朝の光を求め続ける者。30日間の朝活を達成し、成長への道筋を見つけた。",
      "color_theme": "#FFB800",
      "achieved_at": "2024-12-10T06:30:00Z"
    },
    "attendance_stats": {
      "total_attendance_days": 45,
      "current_streak_days": 7,
      "max_streak_days": 15,
      "first_attendance_date": "2024-10-15",
      "last_attendance_date": "2024-12-15",
      "total_duration_minutes": 2700
    },
    "achievements": [
      {
        "title_id": "550e8400-e29b-41d4-a716-446655440001",
        "level": 1,
        "name_jp": "まどろみ見習い",
        "achieved_at": "2024-10-20T06:30:00Z"
      },
      {
        "title_id": "550e8400-e29b-41d4-a716-446655440002",
        "level": 2,
        "name_jp": "早起き戦士",
        "achieved_at": "2024-11-05T06:30:00Z"
      }
    ],
    "is_rival": false,
    "created_at": "2024-12-01T10:30:00Z"
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 3. ユーザー情報更新 API

### ユーザー基本情報更新

```
PUT /users/{userId}
```

**権限**: 👤 本人のみ

**説明**: ユーザーの基本情報（表示名、アバター等）を更新

**リクエスト:**

```json
{
  "username": "tanaka_taro_new",
  "avatar_url": "https://img.clerk.com/new-avatar.png"
}
```

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "tanaka_taro_new",
    "avatar_url": "https://img.clerk.com/new-avatar.png",
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "User updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### ユーザーメタデータ取得

```
GET /users/{userId}/metadata
```

**権限**: 🔐 認証済み（プライバシー設定により一部制限）

**説明**: ユーザーの詳細メタデータを取得

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440100",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "display_name": "田中太郎",
    "profile_image_url": "https://example.com/profile/tanaka.jpg",
    "tagline": "朝活で人生変える！",
    "bio": "エンジニアとして働きながら、朝活を通じて自己成長を目指しています。",
    "vision": "2025年末までにフルスタックエンジニアになる",
    "vision_public": false,
    "timezone": "Asia/Tokyo",
    "skills": ["JavaScript", "React", "Node.js"],
    "interests": ["プログラミング", "読書", "筋トレ"],
    "created_at": "2024-12-01T10:30:00Z",
    "updated_at": "2024-12-15T08:45:00Z"
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### ユーザーメタデータ更新

```
PUT /users/{userId}/metadata
```

**権限**: 👤 本人のみ

**説明**: ユーザーの詳細メタデータを更新。プロフィール設定で使用

**リクエスト:**

```json
{
  "display_name": "田中太郎",
  "profile_image_url": "https://example.com/profile/tanaka_new.jpg",
  "tagline": "朝活で人生を変えた！",
  "bio": "朝活を始めて半年、毎日が充実しています。エンジニアとして働きながら、朝の時間を有効活用して成長を続けています。",
  "vision": "2025年末までにフルスタックエンジニアになって、自分のサービスを作る",
  "vision_public": true,
  "timezone": "Asia/Tokyo",
  "skills": ["JavaScript", "TypeScript", "React", "Node.js", "Go"],
  "interests": ["プログラミング", "読書", "筋トレ", "瞑想"]
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| display_name | 最大100文字 |
| tagline | 最大150文字 |
| bio | 最大1000文字 |
| vision | 最大2000文字 |
| skills | 配列、各要素最大50文字、最大20個 |
| interests | 配列、各要素最大50文字、最大20個 |
| timezone | 有効なタイムゾーン文字列 |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440100",
    "display_name": "田中太郎",
    "tagline": "朝活で人生を変えた！",
    "vision_public": true,
    "skills": ["JavaScript", "TypeScript", "React", "Node.js", "Go"],
    "interests": ["プログラミング", "読書", "筋トレ", "瞑想"],
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Metadata updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 4. SNS・外部リンク管理 API

### SNSリンク一覧取得

```
GET /users/{userId}/social-links
```

**権限**: 🔐 認証済み（is_public=trueのもののみ、本人の場合は全て）

**説明**: ユーザーのSNSリンク・外部リンク一覧を取得

**レスポンス:**

```json
{
  "data": {
    "social_links": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440010",
        "platform": "github",
        "url": "https://github.com/tanaka-taro",
        "title": "GitHub Profile",
        "is_public": true,
        "created_at": "2024-12-01T10:30:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440011",
        "platform": "twitter",
        "url": "https://twitter.com/tanaka_dev",
        "title": "Twitter",
        "is_public": true,
        "created_at": "2024-12-01T10:30:00Z"
      }
    ]
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### SNSリンク追加

```
POST /users/{userId}/social-links
```

**権限**: 👤 本人のみ

**説明**: 新しいSNSリンク・外部リンクを追加

**リクエスト:**

```json
{
  "platform": "website",
  "url": "https://tanaka-dev.com",
  "title": "個人ブログ",
  "is_public": true
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| platform | 必須、最大50文字 (twitter, instagram, github, linkedin, website, blog, youtube, etc.) |
| url | 必須、有効なURL形式 |
| title | 最大100文字 |
| is_public | boolean、デフォルト: true |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440012",
    "platform": "website",
    "url": "https://tanaka-dev.com",
    "title": "個人ブログ",
    "is_public": true,
    "created_at": "2025-01-21T10:00:00Z"
  },
  "message": "Social link added successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### SNSリンク更新

```
PUT /users/{userId}/social-links/{linkId}
```

**権限**: 👤 本人のみ

**説明**: 既存のSNSリンク・外部リンクを更新

**リクエスト:**

```json
{
  "url": "https://tanaka-dev.com/blog",
  "title": "技術ブログ",
  "is_public": false
}
```

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440012",
    "platform": "website",
    "url": "https://tanaka-dev.com/blog",
    "title": "技術ブログ",
    "is_public": false,
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Social link updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### SNSリンク削除

```
DELETE /users/{userId}/social-links/{linkId}
```

**権限**: 👤 本人のみ

**説明**: SNSリンク・外部リンクを削除

**レスポンス:**

```json
{
  "message": "Social link deleted successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 5. ライバル管理 API

### ライバル一覧取得

```
GET /users/{userId}/rivals
```

**権限**: 👤 本人のみ

**説明**: ユーザーが設定したライバル一覧を取得（最大3人）

**レスポンス:**

```json
{
  "data": {
    "rivals": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440020",
        "rival_user": {
          "id": "550e8400-e29b-41d4-a716-446655440001",
          "display_name": "佐藤花子",
          "username": "sato_hanako",
          "avatar_url": "https://img.clerk.com/sato.png",
          "current_title": {
            "level": 4,
            "name_jp": "朝陽の使者",
            "color_theme": "#FF6B35"
          },
          "attendance_stats": {
            "total_attendance_days": 62,
            "current_streak_days": 12
          }
        },
        "created_at": "2024-11-01T10:30:00Z"
      }
    ],
    "count": 1,
    "max_rivals": 3
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### ライバル追加

```
POST /users/{userId}/rivals
```

**権限**: 👤 本人のみ

**説明**: 新しいライバルを追加。ダッシュボードでの比較表示で使用

**リクエスト:**

```json
{
  "rival_user_id": "550e8400-e29b-41d4-a716-446655440002"
}
```

**バリデーション:**
- 最大3人まで設定可能
- 自分自身は設定不可
- 既に設定済みのユーザーは重複不可
- 対象ユーザーが存在し、アクティブである必要あり

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440021",
    "rival_user": {
      "id": "550e8400-e29b-41d4-a716-446655440002",
      "display_name": "鈴木一郎",
      "username": "suzuki_ichiro",
      "avatar_url": "https://img.clerk.com/suzuki.png",
      "current_title": {
        "level": 2,
        "name_jp": "早起き戦士",
        "color_theme": "#4ECDC4"
      }
    },
    "created_at": "2025-01-21T10:00:00Z"
  },
  "message": "Rival added successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

**エラーレスポンス例:**

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Maximum number of rivals reached",
    "details": {
      "current_count": 3,
      "max_allowed": 3
    }
  },
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### ライバル削除

```
DELETE /users/{userId}/rivals/{rivalId}
```

**権限**: 👤 本人のみ

**説明**: ライバル関係を解除

**レスポンス:**

```json
{
  "message": "Rival relationship removed successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## エラーハンドリング

### 共通エラーレスポンス

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "エラーメッセージ",
    "details": {}
  },
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### エラーコード一覧

| Code | Description | HTTP Status |
|------|-------------|-------------|
| UNAUTHORIZED | 認証が必要 | 401 |
| FORBIDDEN | アクセス権限なし | 403 |
| USER_NOT_FOUND | ユーザーが見つからない | 404 |
| VALIDATION_ERROR | リクエストデータが無効 | 422 |
| RESOURCE_NOT_FOUND | リソースが見つからない | 404 |
| DUPLICATE_ENTRY | 重複するデータ | 409 |
| RATE_LIMIT_EXCEEDED | レート制限超過 | 429 |

## ビジネスルール（ドメイン層実装）

### ドメインルール
```go
// internal/domain/user/errors.go
var (
    ErrUserNotFound         = errors.New("ユーザーが見つかりません")
    ErrInvalidEmail         = errors.New("無効なメールアドレスです")
    ErrDuplicateEmail       = errors.New("このメールアドレスは既に使用されています")
    ErrRivalLimitExceeded   = errors.New("ライバルは最大3人まで設定できます")
    ErrCannotRivalSelf      = errors.New("自分自身をライバルに設定できません")
    ErrInvalidPlatform      = errors.New("無効なプラットフォームです")
    ErrInvalidURL           = errors.New("無効なURL形式です")
)

// internal/domain/user/entity/user.go
func (u *User) Validate() error {
    if u.Email == "" {
        return errors.New("メールアドレスは必須です")
    }
    if !isValidEmail(u.Email) {
        return ErrInvalidEmail
    }
    return nil
}

func (u *User) IsActive() bool {
    return u.IsActive
}

// internal/domain/user/entity/user_metadata.go
func (um *UserMetadata) Validate() error {
    if len(um.DisplayName) > 100 {
        return errors.New("表示名は100文字以内で入力してください")
    }
    if len(um.Tagline) > 150 {
        return errors.New("一言プロフィールは150文字以内で入力してください")
    }
    if len(um.Bio) > 1000 {
        return errors.New("自己紹介は1000文字以内で入力してください")
    }
    if len(um.Vision) > 2000 {
        return errors.New("ビジョンは2000文字以内で入力してください")
    }
    if len(um.Skills) > 20 {
        return errors.New("スキルは最大20個まで設定できます")
    }
    if len(um.Interests) > 20 {
        return errors.New("興味・関心は最大20個まで設定できます")
    }
    return nil
}

// internal/domain/user/entity/user_social_link.go
func (usl *UserSocialLink) Validate() error {
    if !usl.Platform.IsValid() {
        return ErrInvalidPlatform
    }
    if !isValidURL(usl.URL) {
        return ErrInvalidURL
    }
    if len(usl.Title) > 100 {
        return errors.New("タイトルは100文字以内で入力してください")
    }
    return nil
}

// internal/domain/user/value/platform.go
func (p Platform) IsValid() bool {
    switch p {
    case PlatformTwitter, PlatformGitHub, PlatformLinkedIn, PlatformWebsite, PlatformBlog:
        return true
    default:
        return false
    }
}

// internal/domain/user/entity/user_rival.go
func (ur *UserRival) Validate() error {
    if ur.UserID == ur.RivalUserID {
        return ErrCannotRivalSelf
    }
    return nil
}
```

### ユーザー管理のルール
1. **メール重複防止**: 同一メールアドレスでの重複登録不可（User.Validate()で制御）
2. **ライバル制限**: 最大3人まで設定可能（ドメインサービスで制御）
3. **プロフィール文字数制限**: 各フィールドに適切な文字数制限（UserMetadata.Validate()で制御）
4. **SNSリンク検証**: URL形式とプラットフォームの有効性チェック（UserSocialLink.Validate()で制御）
5. **アカウント状態管理**: 論理削除による非アクティブ化（User.IsActive()で制御）

## 注意事項

### プライバシー考慮
- `vision`フィールドは`vision_public=true`または本人の場合のみ表示
- SNSリンクは`is_public=true`または本人の場合のみ表示
- 他ユーザーの詳細情報は公開設定に基づいて制限

### レート制限
- 認証API: 10req/min
- 一般API: 100req/min
- 更新API: 30req/min

### キャッシュ戦略
- ユーザー一覧: 5分キャッシュ
- ユーザー詳細: 1分キャッシュ
- 統計情報: 10分キャッシュ