# Title Management API - 詳細設計

## Overview
Ghoona Campアプリケーションの称号・バッジ管理関連API詳細設計書です。
8段階の称号システムとユーザーの称号獲得・表示管理を中心とした機能を提供します。

## アーキテクチャ適合性
このAPI設計は以下のアーキテクチャ原則に従います：
- **オニオンアーキテクチャ**: Domain → Application → Infrastructure → Interface層の依存関係
- **ドメイン駆動設計**: `title`コンテキストを中心としたエンティティ設計
- **クリーンな境界**: DTO、ドメインエンティティ、GORMモデルの適切な分離

## ドメイン設計対応

### ドメインエンティティ
```go
// internal/domain/title/entity/title.go
type Title struct {
    ID           UUID
    Level        int
    NameJP       string
    NameEN       string
    Description  string
    RequiredDays int
    ImageURL     string
    ColorTheme   string
    IsActive     bool
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

// internal/domain/title/entity/title_achievement.go
type TitleAchievement struct {
    ID          UUID
    UserID      UUID
    TitleID     UUID
    AchievedAt  time.Time
    IsCurrent   bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// internal/domain/title/value/title_level.go
type TitleLevel int
const (
    LevelSleeper       TitleLevel = 1  // まどろみ見習い
    LevelEarlyRiser    TitleLevel = 2  // 早起き戦士
    LevelDawnSeeker    TitleLevel = 3  // 陽光探求者
    LevelSunMessenger  TitleLevel = 4  // 朝陽の使者
    LevelDawnGuardian  TitleLevel = 5  // 暁の守護者
    LevelMorningMaster TitleLevel = 6  // 朝活の覇者
    LevelDawnEmperor   TitleLevel = 7  // 夜明けの皇帝
    LevelEternalDawn   TitleLevel = 8  // 永遠の夜明け
)
```

### リポジトリインターフェース
```go
// internal/domain/title/repository/title_repository.go
type TitleRepository interface {
    FindAll(ctx context.Context) ([]*entity.Title, error)
    FindByID(ctx context.Context, id UUID) (*entity.Title, error)
    FindByLevel(ctx context.Context, level int) (*entity.Title, error)
    Create(ctx context.Context, title *entity.Title) error
    Update(ctx context.Context, title *entity.Title) error
    Delete(ctx context.Context, id UUID) error
}

type TitleAchievementRepository interface {
    FindByUserID(ctx context.Context, userID UUID) ([]*entity.TitleAchievement, error)
    FindByTitleID(ctx context.Context, titleID UUID) ([]*entity.TitleAchievement, error)
    FindCurrentByUserID(ctx context.Context, userID UUID) (*entity.TitleAchievement, error)
    Create(ctx context.Context, achievement *entity.TitleAchievement) error
    Update(ctx context.Context, achievement *entity.TitleAchievement) error
    CountByTitleID(ctx context.Context, titleID UUID) (int, error)
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

## 1. 称号管理 API

### 全称号一覧取得

```
GET /titles
```

**権限**: 🔐 認証済み

**説明**: 全称号（8段階）の一覧を取得。レベル、必要日数、説明、画像等を含む

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| include_inactive | boolean | No | 非アクティブな称号も含める (デフォルト: false) |
| user_id | UUID | No | 指定ユーザーの獲得状況も含めて取得 |

**レスポンス:**

```json
{
  "data": {
    "titles": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440001",
        "level": 1,
        "name_jp": "まどろみ見習い",
        "name_en": "Sleeper",
        "description": "朝活の世界への第一歩。まだ眠気と戦いながらも、新しい習慣への道を歩み始めた勇気ある者。1日の参加で獲得できる称号。",
        "required_days": 1,
        "image_url": "https://cdn.ghoona-camp.com/titles/level-1.png",
        "color_theme": "#B0BEC5",
        "holders_count": 324,
        "achievement_rate": 98.5,
        "user_status": {
          "is_achieved": true,
          "achieved_at": "2024-10-15T06:30:00Z",
          "is_current": false
        },
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440002",
        "level": 2,
        "name_jp": "早起き戦士",
        "name_en": "Early Riser",
        "description": "5日間の継続により、朝の時間を味方につけ始めた戦士。規則正しい生活リズムの基盤を築き上げた証。",
        "required_days": 5,
        "image_url": "https://cdn.ghoona-camp.com/titles/level-2.png",
        "color_theme": "#4ECDC4",
        "holders_count": 287,
        "achievement_rate": 87.2,
        "user_status": {
          "is_achieved": true,
          "achieved_at": "2024-10-20T06:30:00Z",
          "is_current": false
        },
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440003",
        "level": 3,
        "name_jp": "陽光探求者",
        "name_en": "Dawn Seeker",
        "description": "15日間の朝活により、朝の光を求め続ける者。自己成長への意欲と継続力を証明した探求者。",
        "required_days": 15,
        "image_url": "https://cdn.ghoona-camp.com/titles/level-3.png",
        "color_theme": "#FFB800",
        "holders_count": 198,
        "achievement_rate": 60.1,
        "user_status": {
          "is_achieved": true,
          "achieved_at": "2024-11-05T06:30:00Z",
          "is_current": true
        },
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440004",
        "level": 4,
        "name_jp": "朝陽の使者",
        "name_en": "Sun Messenger",
        "description": "30日間の継続により、朝の力を他者にも伝える使者となった。コミュニティに良い影響を与える存在。",
        "required_days": 30,
        "image_url": "https://cdn.ghoona-camp.com/titles/level-4.png",
        "color_theme": "#FF6B35",
        "holders_count": 145,
        "achievement_rate": 44.0,
        "user_status": {
          "is_achieved": false,
          "achieved_at": null,
          "is_current": false,
          "progress": {
            "current_days": 28,
            "remaining_days": 2,
            "progress_rate": 93.3
          }
        },
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440005",
        "level": 5,
        "name_jp": "暁の守護者",
        "name_en": "Dawn Guardian",
        "description": "60日間の朝活により、朝の時間を守り抜く守護者となった。揺るぎない意志と習慣を身につけた証。",
        "required_days": 60,
        "image_url": "https://cdn.ghoona-camp.com/titles/level-5.png",
        "color_theme": "#8E44AD",
        "holders_count": 89,
        "achievement_rate": 27.0,
        "user_status": {
          "is_achieved": false,
          "achieved_at": null,
          "is_current": false,
          "progress": {
            "current_days": 28,
            "remaining_days": 32,
            "progress_rate": 46.7
          }
        },
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440006",
        "level": 6,
        "name_jp": "朝活の覇者",
        "name_en": "Morning Master",
        "description": "100日間の朝活により、朝の時間を完全に制覇した覇者。朝活コミュニティのリーダー的存在。",
        "required_days": 100,
        "image_url": "https://cdn.ghoona-camp.com/titles/level-6.png",
        "color_theme": "#E74C3C",
        "holders_count": 42,
        "achievement_rate": 12.7,
        "user_status": {
          "is_achieved": false,
          "achieved_at": null,
          "is_current": false,
          "progress": {
            "current_days": 28,
            "remaining_days": 72,
            "progress_rate": 28.0
          }
        },
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440007",
        "level": 7,
        "name_jp": "夜明けの皇帝",
        "name_en": "Dawn Emperor",
        "description": "200日間の朝活により、夜明けを支配する皇帝となった。朝活の境地に達した稀有な存在。",
        "required_days": 200,
        "image_url": "https://cdn.ghoona-camp.com/titles/level-7.png",
        "color_theme": "#D4AF37",
        "holders_count": 18,
        "achievement_rate": 5.5,
        "user_status": {
          "is_achieved": false,
          "achieved_at": null,
          "is_current": false,
          "progress": {
            "current_days": 28,
            "remaining_days": 172,
            "progress_rate": 14.0
          }
        },
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440008",
        "level": 8,
        "name_jp": "永遠の夜明け",
        "name_en": "Eternal Dawn",
        "description": "365日間の朝活により、永遠の夜明けを体現する伝説的存在。朝活の神髄を極めた真の覇者。",
        "required_days": 365,
        "image_url": "https://cdn.ghoona-camp.com/titles/level-8.png",
        "color_theme": "#9B59B6",
        "holders_count": 5,
        "achievement_rate": 1.5,
        "user_status": {
          "is_achieved": false,
          "achieved_at": null,
          "is_current": false,
          "progress": {
            "current_days": 28,
            "remaining_days": 337,
            "progress_rate": 7.7
          }
        },
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      }
    ],
    "user_summary": {
      "achieved_count": 3,
      "total_count": 8,
      "current_title": {
        "id": "550e8400-e29b-41d4-a716-446655440003",
        "level": 3,
        "name_jp": "陽光探求者",
        "color_theme": "#FFB800"
      },
      "next_title": {
        "id": "550e8400-e29b-41d4-a716-446655440004",
        "level": 4,
        "name_jp": "朝陽の使者",
        "progress_rate": 93.3,
        "remaining_days": 2
      }
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 称号詳細取得

```
GET /titles/{titleId}
```

**権限**: 🔐 認証済み

**説明**: 指定称号の詳細情報を取得。ストーリー、獲得条件、保持者数等

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| titleId | UUID | Yes | 称号ID |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440004",
    "level": 4,
    "name_jp": "朝陽の使者",
    "name_en": "Sun Messenger",
    "description": "30日間の継続により、朝の力を他者にも伝える使者となった。コミュニティに良い影響を与える存在。",
    "story": "まだ多くの人が夢の中にいる時間、あなたは既に新しい一日の扉を開いている。30日間という決して短くない期間を朝活に捧げたあなたは、今や朝の力を知り尽くした使者として、同じ道を歩む仲間たちに希望の光を与える存在となった。あなたの継続する姿は、多くの人にとって朝活への憧れと動機を与える貴重な存在である。",
    "required_days": 30,
    "image_url": "https://cdn.ghoona-camp.com/titles/level-4.png",
    "color_theme": "#FF6B35",
    "badge_design": {
      "primary_color": "#FF6B35",
      "secondary_color": "#FFE5DB",
      "icon": "🌅",
      "pattern": "radial_gradient"
    },
    "achievement_conditions": [
      {
        "type": "attendance_days",
        "value": 30,
        "description": "30日間の朝活参加"
      },
      {
        "type": "consistency",
        "value": 80,
        "description": "80%以上の出席率維持"
      }
    ],
    "statistics": {
      "total_holders": 145,
      "achievement_rate": 44.0,
      "average_achievement_days": 42,
      "recent_achievers_count": 12
    },
    "recent_achievers": [
      {
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440010",
          "display_name": "山田花子",
          "username": "yamada_hanako",
          "avatar_url": "https://img.clerk.com/yamada.png"
        },
        "achieved_at": "2025-01-20T06:30:00Z"
      },
      {
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440011",
          "display_name": "鈴木太郎",
          "username": "suzuki_taro",
          "avatar_url": "https://img.clerk.com/suzuki.png"
        },
        "achieved_at": "2025-01-18T06:30:00Z"
      }
    ],
    "user_status": {
      "is_achieved": false,
      "progress": {
        "current_days": 28,
        "remaining_days": 2,
        "progress_rate": 93.3,
        "estimated_achievement_date": "2025-01-23"
      }
    },
    "perks": [
      {
        "type": "profile_decoration",
        "description": "プロフィールに特別なオレンジ色のフレームが表示"
      },
      {
        "type": "comment_badge",
        "description": "コメント投稿時に称号バッジが表示"
      },
      {
        "type": "priority_support",
        "description": "イベント参加時の優先枠利用可能"
      }
    ],
    "next_title": {
      "id": "550e8400-e29b-41d4-a716-446655440005",
      "level": 5,
      "name_jp": "暁の守護者",
      "required_days": 60,
      "additional_days_needed": 32
    },
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 2. ユーザー称号管理 API

### ユーザー称号取得履歴

```
GET /users/{userId}/achievements
```

**権限**: 🔐 認証済み（本人の場合は全て、他人の場合は公開情報のみ）

**説明**: ユーザーの称号取得履歴・現在設定中の称号を取得

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| userId | UUID | Yes | ユーザーID |

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| include_progress | boolean | No | 未獲得称号の進捗も含める (デフォルト: false) |
| sort_by | string | No | ソート基準 (achieved_at, level) デフォルト: level |
| order | string | No | ソート順 (asc, desc) デフォルト: asc |

**レスポンス:**

```json
{
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "display_name": "田中太郎",
      "username": "tanaka_taro",
      "avatar_url": "https://img.clerk.com/preview.png"
    },
    "current_title": {
      "id": "550e8400-e29b-41d4-a716-446655440003",
      "level": 3,
      "name_jp": "陽光探求者",
      "name_en": "Dawn Seeker",
      "color_theme": "#FFB800",
      "image_url": "https://cdn.ghoona-camp.com/titles/level-3.png",
      "achieved_at": "2024-11-05T06:30:00Z",
      "days_held": 77
    },
    "achievements": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440100",
        "title": {
          "id": "550e8400-e29b-41d4-a716-446655440001",
          "level": 1,
          "name_jp": "まどろみ見習い",
          "name_en": "Sleeper",
          "color_theme": "#B0BEC5",
          "image_url": "https://cdn.ghoona-camp.com/titles/level-1.png"
        },
        "achieved_at": "2024-10-15T06:30:00Z",
        "is_current": false,
        "achievement_rank": 187,
        "total_achievers_at_time": 186
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440101",
        "title": {
          "id": "550e8400-e29b-41d4-a716-446655440002",
          "level": 2,
          "name_jp": "早起き戦士",
          "name_en": "Early Riser",
          "color_theme": "#4ECDC4",
          "image_url": "https://cdn.ghoona-camp.com/titles/level-2.png"
        },
        "achieved_at": "2024-10-20T06:30:00Z",
        "is_current": false,
        "achievement_rank": 142,
        "total_achievers_at_time": 141
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440102",
        "title": {
          "id": "550e8400-e29b-41d4-a716-446655440003",
          "level": 3,
          "name_jp": "陽光探求者",
          "name_en": "Dawn Seeker",
          "color_theme": "#FFB800",
          "image_url": "https://cdn.ghoona-camp.com/titles/level-3.png"
        },
        "achieved_at": "2024-11-05T06:30:00Z",
        "is_current": true,
        "achievement_rank": 89,
        "total_achievers_at_time": 88
      }
    ],
    "next_title_progress": {
      "title": {
        "id": "550e8400-e29b-41d4-a716-446655440004",
        "level": 4,
        "name_jp": "朝陽の使者",
        "required_days": 30
      },
      "progress": {
        "current_days": 28,
        "remaining_days": 2,
        "progress_rate": 93.3,
        "estimated_achievement_date": "2025-01-23"
      }
    },
    "statistics": {
      "total_achieved": 3,
      "max_level_achieved": 3,
      "achievement_speed_rank": "fast",
      "consistency_score": 85.2,
      "total_attendance_days": 45
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 現在の表示称号変更

```
PUT /users/{userId}/achievements/{titleId}
```

**権限**: 👤 本人のみ

**説明**: 現在表示する称号を変更。獲得済み称号から選択

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| userId | UUID | Yes | ユーザーID |
| titleId | UUID | Yes | 称号ID |

**リクエスト:**

```json
{
  "is_current": true
}
```

**バリデーション:**
- 指定された称号を既に獲得している必要がある
- 同時に表示できる称号は1つのみ
- 非アクティブな称号は設定不可

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440102",
    "title": {
      "id": "550e8400-e29b-41d4-a716-446655440003",
      "level": 3,
      "name_jp": "陽光探求者",
      "name_en": "Dawn Seeker",
      "color_theme": "#FFB800",
      "image_url": "https://cdn.ghoona-camp.com/titles/level-3.png"
    },
    "is_current": true,
    "updated_at": "2025-01-21T10:00:00Z",
    "previous_title": {
      "id": "550e8400-e29b-41d4-a716-446655440002",
      "level": 2,
      "name_jp": "早起き戦士"
    }
  },
  "message": "Current title updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 3. 称号ランキング・統計 API

### 称号別ランキング

```
GET /titles/{titleId}/leaderboard
```

**権限**: 🔐 認証済み

**説明**: 指定称号の獲得者ランキングを取得。獲得日時順で表示

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| period | string | No | 期間 (week, month, all) デフォルト: all |
| page | integer | No | ページ番号 (デフォルト: 1) |
| limit | integer | No | 1ページあたりの件数 (デフォルト: 20, 最大: 50) |

**レスポンス:**

```json
{
  "data": {
    "title": {
      "id": "550e8400-e29b-41d4-a716-446655440004",
      "level": 4,
      "name_jp": "朝陽の使者",
      "name_en": "Sun Messenger",
      "color_theme": "#FF6B35"
    },
    "leaderboard": [
      {
        "rank": 1,
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440020",
          "display_name": "佐藤太郎",
          "username": "sato_taro",
          "avatar_url": "https://img.clerk.com/sato.png",
          "current_title": {
            "level": 6,
            "name_jp": "朝活の覇者",
            "color_theme": "#E74C3C"
          }
        },
        "achieved_at": "2024-08-15T06:30:00Z",
        "achievement_day": 30,
        "consistency_rate": 100.0,
        "is_current_user": false
      },
      {
        "rank": 2,
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440021",
          "display_name": "田中花子",
          "username": "tanaka_hanako",
          "avatar_url": "https://img.clerk.com/tanaka.png",
          "current_title": {
            "level": 5,
            "name_jp": "暁の守護者",
            "color_theme": "#8E44AD"
          }
        },
        "achieved_at": "2024-09-02T06:30:00Z",
        "achievement_day": 32,
        "consistency_rate": 93.8,
        "is_current_user": false
      },
      {
        "rank": 3,
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440000",
          "display_name": "田中太郎",
          "username": "tanaka_taro",
          "avatar_url": "https://img.clerk.com/preview.png",
          "current_title": {
            "level": 3,
            "name_jp": "陽光探求者",
            "color_theme": "#FFB800"
          }
        },
        "achieved_at": null,
        "achievement_day": null,
        "consistency_rate": null,
        "progress": {
          "current_days": 28,
          "remaining_days": 2,
          "progress_rate": 93.3
        },
        "is_current_user": true
      }
    ],
    "statistics": {
      "total_achievers": 145,
      "average_achievement_days": 42,
      "fastest_achievement_days": 30,
      "current_user_rank": null,
      "current_user_progress_rank": 3
    },
    "pagination": {
      "current_page": 1,
      "total_pages": 8,
      "total_count": 148,
      "limit": 20,
      "has_next": true,
      "has_prev": false
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 称号統計情報

```
GET /titles/statistics
```

**権限**: 🔐 認証済み

**説明**: 全称号の統計情報を取得。獲得率、トレンド等

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| period | string | No | 集計期間 (week, month, quarter, year) デフォルト: month |

**レスポンス:**

```json
{
  "data": {
    "period": "month",
    "period_range": {
      "from": "2025-01-01",
      "to": "2025-01-31"
    },
    "overall_statistics": {
      "total_users": 329,
      "total_achievements_this_period": 47,
      "most_achieved_title": {
        "id": "550e8400-e29b-41d4-a716-446655440001",
        "level": 1,
        "name_jp": "まどろみ見習い",
        "achievements_count": 15
      },
      "rarest_title": {
        "id": "550e8400-e29b-41d4-a716-446655440008",
        "level": 8,
        "name_jp": "永遠の夜明け",
        "holders_count": 5,
        "achievement_rate": 1.5
      }
    },
    "title_statistics": [
      {
        "title": {
          "id": "550e8400-e29b-41d4-a716-446655440001",
          "level": 1,
          "name_jp": "まどろみ見習い",
          "color_theme": "#B0BEC5"
        },
        "holders_count": 324,
        "achievement_rate": 98.5,
        "new_achievers_this_period": 15,
        "average_days_to_achieve": 1.0,
        "trend": "stable"
      },
      {
        "title": {
          "id": "550e8400-e29b-41d4-a716-446655440002",
          "level": 2,
          "name_jp": "早起き戦士",
          "color_theme": "#4ECDC4"
        },
        "holders_count": 287,
        "achievement_rate": 87.2,
        "new_achievers_this_period": 12,
        "average_days_to_achieve": 5.2,
        "trend": "increasing"
      }
    ],
    "user_distribution": {
      "level_1": 324,
      "level_2": 287,
      "level_3": 198,
      "level_4": 145,
      "level_5": 89,
      "level_6": 42,
      "level_7": 18,
      "level_8": 5
    },
    "trends": {
      "monthly_growth_rate": 14.6,
      "retention_after_level_1": 88.6,
      "average_progression_speed": "normal"
    }
  },
  "message": "success",
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
| TITLE_NOT_FOUND | 称号が見つからない | 404 |
| TITLE_NOT_ACHIEVED | 未獲得の称号です | 403 |
| TITLE_ALREADY_CURRENT | 既に現在の称号に設定済み | 409 |
| ACHIEVEMENT_NOT_FOUND | 称号獲得記録が見つからない | 404 |
| INVALID_TITLE_LEVEL | 無効な称号レベル | 422 |
| TITLE_INACTIVE | 非アクティブな称号です | 422 |
| ACHIEVEMENT_CONDITIONS_NOT_MET | 獲得条件を満たしていません | 422 |

## ビジネスルール（ドメイン層実装）

### ドメインルール
```go
// internal/domain/title/errors.go
var (
    ErrTitleNotFound          = errors.New("称号が見つかりません")
    ErrTitleNotAchieved      = errors.New("未獲得の称号です")
    ErrInvalidTitleLevel     = errors.New("無効な称号レベルです")
    ErrTitleAlreadyCurrent   = errors.New("既に現在の称号に設定済みです")
    ErrConditionsNotMet      = errors.New("獲得条件を満たしていません")
    ErrTitleInactive         = errors.New("非アクティブな称号です")
)

// internal/domain/title/entity/title.go
func (t *Title) Validate() error {
    if t.Level < 1 || t.Level > 8 {
        return ErrInvalidTitleLevel
    }
    if t.NameJP == "" || t.NameEN == "" {
        return errors.New("称号名は必須です")
    }
    if t.RequiredDays <= 0 {
        return errors.New("必要日数は1以上である必要があります")
    }
    return nil
}

func (t *Title) IsAchievableBy(attendanceDays int) bool {
    return attendanceDays >= t.RequiredDays && t.IsActive
}

// internal/domain/title/entity/title_achievement.go
func (ta *TitleAchievement) Validate() error {
    if ta.UserID == uuid.Nil || ta.TitleID == uuid.Nil {
        return errors.New("ユーザーIDと称号IDは必須です")
    }
    if ta.AchievedAt.After(time.Now()) {
        return errors.New("未来日の獲得日時は設定できません")
    }
    return nil
}

func (ta *TitleAchievement) CanBeSetAsCurrent() bool {
    return ta.UserID != uuid.Nil && ta.TitleID != uuid.Nil
}

// internal/domain/title/value/title_level.go
func (tl TitleLevel) IsValid() bool {
    return tl >= 1 && tl <= 8
}

func (tl TitleLevel) GetRequiredDays() int {
    requirements := map[TitleLevel]int{
        LevelSleeper:       1,
        LevelEarlyRiser:    5,
        LevelDawnSeeker:    15,
        LevelSunMessenger:  30,
        LevelDawnGuardian:  60,
        LevelMorningMaster: 100,
        LevelDawnEmperor:   200,
        LevelEternalDawn:   365,
    }
    return requirements[tl]
}
```

### 称号システムのルール
1. **段階的獲得**: 称号は順次獲得（レベル3を取得するには1,2の獲得が必要）
2. **参加日数ベース**: 朝活参加日数に基づいて自動獲得（Title.IsAchievableBy()で制御）
3. **表示称号**: 獲得済み称号から1つのみ選択可能（TitleAchievement.CanBeSetAsCurrent()で制御）
4. **獲得条件**: 各称号には最低参加日数が設定（TitleLevel.GetRequiredDays()で制御）
5. **非可逆性**: 一度獲得した称号は削除不可（論理削除のみ）

### 称号獲得判定のルール
1. **自動判定**: 日次バッチで参加日数を確認し自動付与
2. **重複防止**: 同一称号の重複獲得不可（リポジトリ層で制御）
3. **遡及適用**: 過去の参加実績も加算対象
4. **一貫性保証**: 下位称号未獲得の場合は上位称号も獲得不可

## 注意事項

### 称号システム設計
- 称号は8段階の固定システム（1-8レベル）
- 各称号には独自のストーリーとビジュアルデザインを設定
- 獲得率の低い高レベル称号ほど希少価値を演出

### パフォーマンス考慮
- 称号一覧は静的データのため積極的キャッシュ
- ユーザー進捗計算は日次バッチで事前計算
- ランキング表示は定期的に更新

### キャッシュ戦略
- 称号マスターデータ: 1時間キャッシュ
- ユーザー称号情報: 10分キャッシュ
- 称号統計情報: 1時間キャッシュ
- ランキングデータ: 30分キャッシュ