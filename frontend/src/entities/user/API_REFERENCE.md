# User API Reference

User関連の全てのAPIエンドポイントのリクエスト・レスポンス例をまとめたドキュメントです。

## 目次

- [認証・基本情報](#認証基本情報)
- [ユーザーメタデータ](#ユーザーメタデータ)
- [ライバル機能](#ライバル機能)
- [ソーシャルリンク](#ソーシャルリンク)

---

## 認証・基本情報

### GET /auth/me
現在ログイン中のユーザー情報を取得

**Request:**
```http
GET /auth/me
Authorization: Bearer <token>
```

**Response:**
```json
{
  "data": {
    "id": "user_123",
    "clerkId": "clerk_abc123",
    "email": "john.doe@example.com",
    "username": "johndoe",
    "avatarUrl": "https://example.com/avatars/john.jpg",
    "discordId": "discord_456",
    "isActive": true,
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-03-10T15:45:00Z",
    "metadata": {
      "displayName": "John Doe",
      "tagline": "Full-stack developer"
    },
    "attendanceStats": {
      "totalAttendanceDays": 45,
      "currentStreakDays": 7
    }
  },
  "message": "User information retrieved successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### GET /users
ユーザー一覧を取得（検索・フィルタリング対応）

**Request:**
```http
GET /users?page=1&limit=20&search=john&skills=JavaScript,React&sort_by=attendance_days&order=desc
```

**Query Parameters:**
- `page`: ページ番号（デフォルト: 1）
- `limit`: 1ページあたりの件数（デフォルト: 20）
- `search`: ユーザー名での検索
- `skills`: スキルでフィルタリング（カンマ区切り）
- `interests`: 興味・関心でフィルタリング（カンマ区切り）
- `sort_by`: ソート基準（name | attendance_days | streak_days | created_at）
- `order`: ソート順（asc | desc）

**Response:**
```json
{
  "data": {
    "users": [
      {
        "id": "user_123",
        "displayName": "John Doe",
        "username": "johndoe",
        "avatarUrl": "https://example.com/avatars/john.jpg",
        "tagline": "Full-stack developer passionate about React",
        "skills": ["JavaScript", "React", "Node.js", "TypeScript"],
        "interests": ["Web Development", "AI", "Open Source"],
        "currentTitle": {
          "level": 3,
          "nameJp": "シニアデベロッパー",
          "nameEn": "Senior Developer",
          "colorTheme": "#3B82F6"
        },
        "attendanceStats": {
          "totalAttendanceDays": 45,
          "currentStreakDays": 7
        },
        "createdAt": "2024-01-15T10:30:00Z"
      }
    ],
    "pagination": {
      "currentPage": 1,
      "totalPages": 5,
      "totalCount": 89,
      "limit": 20,
      "hasNext": true,
      "hasPrev": false
    }
  },
  "message": "Users retrieved successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### GET /users/{userId}
特定ユーザーの詳細情報を取得

**Request:**
```http
GET /users/user_123
```

**Response:**
```json
{
  "data": {
    "id": "user_123",
    "displayName": "John Doe",
    "username": "johndoe",
    "avatarUrl": "https://example.com/avatars/john.jpg",
    "profileImageUrl": "https://example.com/profiles/john_large.jpg",
    "tagline": "Full-stack developer passionate about React",
    "bio": "5年間のWeb開発経験を持つフルスタックエンジニア。特にReactとNode.jsが得意分野です。オープンソースプロジェクトにも積極的に貢献しています。",
    "vision": "テクノロジーの力で世界をより良い場所にしたい",
    "timezone": "Asia/Tokyo",
    "skills": ["JavaScript", "React", "Node.js", "TypeScript", "Python", "Docker"],
    "interests": ["Web Development", "AI/ML", "Open Source", "Tech Meetups"],
    "socialLinks": [
      {
        "id": "link_1",
        "platform": "github",
        "url": "https://github.com/johndoe",
        "title": "GitHub",
        "isPublic": true
      },
      {
        "id": "link_2",
        "platform": "twitter",
        "url": "https://twitter.com/johndoe_dev",
        "title": "Twitter",
        "isPublic": true
      }
    ],
    "currentTitle": {
      "id": "title_senior_dev",
      "level": 3,
      "nameJp": "シニアデベロッパー",
      "nameEn": "Senior Developer",
      "description": "3年以上の開発経験を持つエンジニア",
      "colorTheme": "#3B82F6",
      "achievedAt": "2024-02-01T09:00:00Z"
    },
    "attendanceStats": {
      "totalAttendanceDays": 45,
      "currentStreakDays": 7,
      "maxStreakDays": 21,
      "firstAttendanceDate": "2024-01-15T10:30:00Z",
      "lastAttendanceDate": "2024-03-15T11:45:00Z",
      "totalDurationMinutes": 2700
    },
    "achievements": [
      {
        "titleId": "first_attendance",
        "level": 1,
        "nameJp": "初回参加",
        "achievedAt": "2024-01-15T10:30:00Z"
      },
      {
        "titleId": "week_streak",
        "level": 2,
        "nameJp": "1週間連続参加",
        "achievedAt": "2024-01-22T10:30:00Z"
      }
    ],
    "isRival": false,
    "createdAt": "2024-01-15T10:30:00Z"
  },
  "message": "User details retrieved successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### PUT /users/{userId}
ユーザーの基本情報を更新

**Request:**
```http
PUT /users/user_123
Content-Type: application/json

{
  "username": "john_doe_updated",
  "avatarUrl": "https://example.com/avatars/john_new.jpg"
}
```

**Response:**
```json
{
  "data": {
    "id": "user_123",
    "username": "john_doe_updated",
    "avatarUrl": "https://example.com/avatars/john_new.jpg",
    "updatedAt": "2024-03-15T12:00:00Z"
  },
  "message": "User updated successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

---

## ユーザーメタデータ

### GET /users/{userId}/metadata
ユーザーのメタデータ（プロフィール詳細）を取得

**Request:**
```http
GET /users/user_123/metadata
```

**Response:**
```json
{
  "data": {
    "id": "metadata_456",
    "userId": "user_123",
    "displayName": "John Doe",
    "profileImageUrl": "https://example.com/profiles/john_large.jpg",
    "tagline": "Full-stack developer passionate about React",
    "bio": "5年間のWeb開発経験を持つフルスタックエンジニア。特にReactとNode.jsが得意分野です。オープンソースプロジェクトにも積極的に貢献しています。",
    "vision": "テクノロジーの力で世界をより良い場所にしたい",
    "visionPublic": true,
    "timezone": "Asia/Tokyo",
    "skills": ["JavaScript", "React", "Node.js", "TypeScript", "Python"],
    "interests": ["Web Development", "AI/ML", "Open Source"],
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-03-10T15:45:00Z"
  },
  "message": "User metadata retrieved successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### PUT /users/{userId}/metadata
ユーザーのメタデータを更新

**Request:**
```http
PUT /users/user_123/metadata
Content-Type: application/json

{
  "displayName": "John Doe Jr.",
  "tagline": "Senior Full-stack Developer",
  "bio": "6年間のWeb開発経験を持つシニアフルスタックエンジニア。React、Node.js、AWS等の技術に精通。",
  "vision": "最新技術で社会課題を解決するプロダクトを作りたい",
  "visionPublic": true,
  "timezone": "Asia/Tokyo",
  "skills": ["JavaScript", "React", "Node.js", "TypeScript", "Python", "AWS"],
  "interests": ["Web Development", "AI/ML", "Open Source", "SaaS Development"]
}
```

**Response:**
```json
{
  "data": {
    "id": "metadata_456",
    "displayName": "John Doe Jr.",
    "tagline": "Senior Full-stack Developer",
    "visionPublic": true,
    "skills": ["JavaScript", "React", "Node.js", "TypeScript", "Python", "AWS"],
    "interests": ["Web Development", "AI/ML", "Open Source", "SaaS Development"],
    "updatedAt": "2024-03-15T12:00:00Z"
  },
  "message": "User metadata updated successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

---

## ライバル機能

### GET /users/{userId}/rivals
ユーザーのライバル一覧を取得

**Request:**
```http
GET /users/user_123/rivals
```

**Response:**
```json
{
  "data": {
    "rivals": [
      {
        "id": "rival_relation_1",
        "rivalUser": {
          "id": "user_456",
          "displayName": "Alice Smith",
          "username": "alicesmith",
          "avatarUrl": "https://example.com/avatars/alice.jpg",
          "currentTitle": {
            "level": 3,
            "nameJp": "シニアデベロッパー",
            "colorTheme": "#10B981"
          },
          "attendanceStats": {
            "totalAttendanceDays": 52,
            "currentStreakDays": 12
          }
        },
        "createdAt": "2024-02-01T10:00:00Z"
      },
      {
        "id": "rival_relation_2",
        "rivalUser": {
          "id": "user_789",
          "displayName": "Bob Johnson",
          "username": "bobjohnson",
          "avatarUrl": "https://example.com/avatars/bob.jpg",
          "currentTitle": {
            "level": 2,
            "nameJp": "ミドルデベロッパー",
            "colorTheme": "#F59E0B"
          },
          "attendanceStats": {
            "totalAttendanceDays": 38,
            "currentStreakDays": 5
          }
        },
        "createdAt": "2024-02-15T14:30:00Z"
      }
    ],
    "count": 2,
    "maxRivals": 3
  },
  "message": "User rivals retrieved successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### POST /users/{userId}/rivals
新しいライバルを追加

**Request:**
```http
POST /users/user_123/rivals
Content-Type: application/json

{
  "rivalUserId": "user_999"
}
```

**Response:**
```json
{
  "data": {
    "id": "rival_relation_3",
    "rivalUser": {
      "id": "user_999",
      "displayName": "Charlie Brown",
      "username": "charliebrown",
      "avatarUrl": "https://example.com/avatars/charlie.jpg",
      "currentTitle": {
        "level": 4,
        "nameJp": "リードデベロッパー",
        "colorTheme": "#8B5CF6"
      },
      "attendanceStats": {
        "totalAttendanceDays": 67,
        "currentStreakDays": 15
      }
    },
    "createdAt": "2024-03-15T12:00:00Z"
  },
  "message": "Rival added successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### DELETE /users/{userId}/rivals/{rivalId}
ライバル関係を解除

**Request:**
```http
DELETE /users/user_123/rivals/rival_relation_1
```

**Response:**
```http
204 No Content
```

---

## ソーシャルリンク

### GET /users/{userId}/social-links
ユーザーのソーシャルリンク一覧を取得

**Request:**
```http
GET /users/user_123/social-links
```

**Response:**
```json
{
  "data": {
    "socialLinks": [
      {
        "id": "link_1",
        "platform": "github",
        "url": "https://github.com/johndoe",
        "title": "GitHub",
        "isPublic": true,
        "createdAt": "2024-01-15T10:30:00Z",
        "updatedAt": "2024-01-15T10:30:00Z"
      },
      {
        "id": "link_2",
        "platform": "twitter",
        "url": "https://twitter.com/johndoe_dev",
        "title": "Twitter",
        "isPublic": true,
        "createdAt": "2024-01-20T15:45:00Z",
        "updatedAt": "2024-02-10T11:20:00Z"
      },
      {
        "id": "link_3",
        "platform": "linkedin",
        "url": "https://linkedin.com/in/johndoe",
        "title": "LinkedIn",
        "isPublic": false,
        "createdAt": "2024-02-01T09:15:00Z",
        "updatedAt": "2024-02-01T09:15:00Z"
      }
    ]
  },
  "message": "Social links retrieved successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### POST /users/{userId}/social-links
新しいソーシャルリンクを追加

**Request:**
```http
POST /users/user_123/social-links
Content-Type: application/json

{
  "platform": "youtube",
  "url": "https://youtube.com/@johndoe",
  "title": "YouTube Channel",
  "isPublic": true
}
```

**Response:**
```json
{
  "data": {
    "id": "link_4",
    "platform": "youtube",
    "url": "https://youtube.com/@johndoe",
    "title": "YouTube Channel",
    "isPublic": true,
    "createdAt": "2024-03-15T12:00:00Z",
    "updatedAt": "2024-03-15T12:00:00Z"
  },
  "message": "Social link created successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### PUT /users/{userId}/social-links/{linkId}
既存のソーシャルリンクを更新

**Request:**
```http
PUT /users/user_123/social-links/link_2
Content-Type: application/json

{
  "url": "https://twitter.com/johndoe_developer",
  "title": "Twitter (Developer)",
  "isPublic": true
}
```

**Response:**
```json
{
  "data": {
    "id": "link_2",
    "platform": "twitter",
    "url": "https://twitter.com/johndoe_developer",
    "title": "Twitter (Developer)",
    "isPublic": true,
    "createdAt": "2024-01-20T15:45:00Z",
    "updatedAt": "2024-03-15T12:00:00Z"
  },
  "message": "Social link updated successfully",
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### DELETE /users/{userId}/social-links/{linkId}
ソーシャルリンクを削除

**Request:**
```http
DELETE /users/user_123/social-links/link_3
```

**Response:**
```http
204 No Content
```

---

## エラーレスポンス

全てのエンドポイントで共通のエラーレスポンス形式を使用します。

### 400 Bad Request
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": [
      {
        "field": "email",
        "message": "Invalid email format"
      }
    ]
  },
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### 401 Unauthorized
```json
{
  "error": {
    "code": "UNAUTHORIZED",
    "message": "Authentication required"
  },
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### 403 Forbidden
```json
{
  "error": {
    "code": "FORBIDDEN",
    "message": "Access denied"
  },
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### 404 Not Found
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "User not found"
  },
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### 409 Conflict
```json
{
  "error": {
    "code": "CONFLICT",
    "message": "Rival relationship already exists"
  },
  "timestamp": "2024-03-15T12:00:00Z"
}
```

### 500 Internal Server Error
```json
{
  "error": {
    "code": "INTERNAL_SERVER_ERROR",
    "message": "An unexpected error occurred"
  },
  "timestamp": "2024-03-15T12:00:00Z"
}
```

---

## 利用可能なプラットフォーム（ソーシャルリンク）

- `twitter`
- `instagram`
- `github`
- `linkedin`
- `website`
- `blog`
- `youtube`
- `facebook`
- `discord`
- `twitch`