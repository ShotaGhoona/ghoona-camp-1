# Attendance Management API - 詳細設計

## Overview
Ghoona Campアプリケーションの参加記録・統計管理関連API詳細設計書です。
Discord参加ログ・出席統計・ランキング機能を中心とした朝活参加管理を提供します。

## アーキテクチャ適合性
このAPI設計は以下のアーキテクチャ原則に従います：
- **オニオンアーキテクチャ**: Domain → Application → Infrastructure → Interface層の依存関係
- **ドメイン駆動設計**: `attendance`コンテキストを中心としたエンティティ設計
- **クリーンな境界**: DTO、ドメインエンティティ、GORMモデルの適切な分離

## ドメイン設計対応

### ドメインエンティティ
```go
// internal/domain/attendance/entity/attendance_log.go
type AttendanceLog struct {
    ID               UUID
    UserID           UUID
    EventID          *UUID
    DiscordChannelID string
    JoinedAt         time.Time
    LeftAt           *time.Time
    DurationMinutes  int
    IsValid          bool
    CreatedAt        time.Time
    UpdatedAt        time.Time
}

// internal/domain/attendance/entity/attendance_summary.go
type AttendanceSummary struct {
    ID                     UUID
    UserID                 UUID
    Date                   time.Time
    TotalDurationMinutes   int
    SessionCount           int
    FirstJoinTime          *time.Time
    LastLeaveTime          *time.Time
    IsMorningActive        bool
    CreatedAt              time.Time
    UpdatedAt              time.Time
}

// internal/domain/attendance/entity/attendance_statistics.go
type AttendanceStatistics struct {
    ID                   UUID
    UserID               UUID
    TotalAttendanceDays  int
    CurrentStreakDays    int
    MaxStreakDays        int
    LastAttendanceDate   *time.Time
    FirstAttendanceDate  *time.Time
    TotalDurationMinutes int
    CreatedAt            time.Time
    UpdatedAt            time.Time
}

// internal/domain/attendance/value/morning_hours.go
type MorningHours struct {
    StartHour int // 6:00
    EndHour   int // 7:00
}

// internal/domain/attendance/value/attendance_validity.go
type AttendanceValidity struct {
    MinDurationMinutes int
    ValidChannels      []string
    MorningTimeRange   MorningHours
}
```

### リポジトリインターフェース
```go
// internal/domain/attendance/repository/attendance_log_repository.go
type AttendanceLogRepository interface {
    GetByUserID(ctx context.Context, userID UUID, filters LogFilters) ([]*entity.AttendanceLog, error)
    GetByDate(ctx context.Context, date time.Time) ([]*entity.AttendanceLog, error)
    GetByEventID(ctx context.Context, eventID UUID) ([]*entity.AttendanceLog, error)
    Create(ctx context.Context, log *entity.AttendanceLog) error
    Update(ctx context.Context, log *entity.AttendanceLog) error
    Delete(ctx context.Context, id UUID) error
}

type AttendanceSummaryRepository interface {
    GetByUserID(ctx context.Context, userID UUID, filters SummaryFilters) ([]*entity.AttendanceSummary, error)
    GetByDate(ctx context.Context, userID UUID, date time.Time) (*entity.AttendanceSummary, error)
    Create(ctx context.Context, summary *entity.AttendanceSummary) error
    Update(ctx context.Context, summary *entity.AttendanceSummary) error
    GetMonthlyData(ctx context.Context, userID UUID, year, month int) ([]*entity.AttendanceSummary, error)
}

type AttendanceStatisticsRepository interface {
    GetByUserID(ctx context.Context, userID UUID) (*entity.AttendanceStatistics, error)
    Create(ctx context.Context, stats *entity.AttendanceStatistics) error
    Update(ctx context.Context, stats *entity.AttendanceStatistics) error
    GetRankings(ctx context.Context, rankingType string, limit int) ([]*AttendanceRanking, error)
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

## 1. 参加ログ管理 API

### ユーザー参加ログ取得

```
GET /users/{userId}/attendance/logs
```

**権限**: 👤 本人のみ

**説明**: ユーザーのDiscord参加ログ詳細を取得。日時、参加時間、チャンネル等

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| userId | UUID | Yes | ユーザーID |

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| from_date | string | No | 開始日 (YYYY-MM-DD) |
| to_date | string | No | 終了日 (YYYY-MM-DD) |
| channel_id | string | No | Discordチャンネルでのフィルタ |
| min_duration | integer | No | 最低参加時間（分）でのフィルタ |
| is_valid | boolean | No | 有効な参加のみ取得 |
| page | integer | No | ページ番号 (デフォルト: 1) |
| limit | integer | No | 1ページあたりの件数 (デフォルト: 30, 最大: 100) |
| sort_by | string | No | ソート基準 (joined_at, duration_minutes) デフォルト: joined_at |
| order | string | No | ソート順 (asc, desc) デフォルト: desc |

**レスポンス:**

```json
{
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "display_name": "田中太郎",
      "username": "tanaka_taro"
    },
    "logs": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440100",
        "event": {
          "id": "550e8400-e29b-41d4-a716-446655440200",
          "title": "朝の読書会",
          "creator": "佐藤花子"
        },
        "discord_channel": {
          "id": "1234567890123456789",
          "name": "朝の読書会",
          "type": "voice"
        },
        "joined_at": "2025-01-21T06:30:00Z",
        "left_at": "2025-01-21T07:15:00Z",
        "duration_minutes": 45,
        "is_valid": true,
        "is_morning_active": true,
        "session_quality": {
          "consistency_score": 85,
          "engagement_level": "high"
        },
        "related_goal": {
          "id": "550e8400-e29b-41d4-a716-446655440300",
          "title": "毎日読書30分",
          "progress_contributed": true
        },
        "created_at": "2025-01-21T06:30:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440101",
        "event": null,
        "discord_channel": {
          "id": "1234567890123456790",
          "name": "一般朝活",
          "type": "voice"
        },
        "joined_at": "2025-01-20T06:15:00Z",
        "left_at": "2025-01-20T06:50:00Z",
        "duration_minutes": 35,
        "is_valid": true,
        "is_morning_active": true,
        "session_quality": {
          "consistency_score": 92,
          "engagement_level": "high"
        },
        "related_goal": null,
        "created_at": "2025-01-20T06:15:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440102",
        "event": null,
        "discord_channel": {
          "id": "1234567890123456791",
          "name": "雑談ルーム",
          "type": "voice"
        },
        "joined_at": "2025-01-19T05:45:00Z",
        "left_at": "2025-01-19T06:05:00Z",
        "duration_minutes": 20,
        "is_valid": false,
        "is_morning_active": false,
        "session_quality": null,
        "related_goal": null,
        "validation_reason": "朝活時間外での参加",
        "created_at": "2025-01-19T05:45:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 3,
      "total_count": 87,
      "limit": 30,
      "has_next": true,
      "has_prev": false
    },
    "summary": {
      "total_sessions": 87,
      "valid_sessions": 78,
      "total_duration_minutes": 3420,
      "average_duration_minutes": 43.8,
      "morning_active_sessions": 72,
      "longest_session_minutes": 120,
      "shortest_valid_session_minutes": 15
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 参加ログ記録（Discord Bot用）

```
POST /attendance/logs
```

**権限**: 🤖 Bot認証

**説明**: Discord Botからの参加ログ記録。入退室時刻を自動記録

**認証ヘッダー:**
```
X-API-Key: your-bot-api-key
```

**リクエスト:**

```json
{
  "user_discord_id": "123456789012345678",
  "discord_channel_id": "1234567890123456789",
  "event_type": "join",
  "timestamp": "2025-01-21T06:30:00Z",
  "event_id": "550e8400-e29b-41d4-a716-446655440200"
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| user_discord_id | 必須、有効なDiscord User ID |
| discord_channel_id | 必須、有効なDiscord Channel ID |
| event_type | 必須、enum(join, leave) |
| timestamp | 必須、ISO8601形式 |
| event_id | 任意、関連イベントID |

**レスポンス:**

```json
{
  "data": {
    "log_id": "550e8400-e29b-41d4-a716-446655440103",
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "display_name": "田中太郎"
    },
    "action": "join",
    "timestamp": "2025-01-21T06:30:00Z",
    "is_morning_active": true,
    "session_info": {
      "current_session_id": "550e8400-e29b-41d4-a716-446655440104",
      "expected_validity": true,
      "related_event": {
        "id": "550e8400-e29b-41d4-a716-446655440200",
        "title": "朝の読書会"
      }
    }
  },
  "message": "Attendance logged successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

**退室ログのレスポンス例:**

```json
{
  "data": {
    "log_id": "550e8400-e29b-41d4-a716-446655440103",
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "display_name": "田中太郎"
    },
    "action": "leave",
    "timestamp": "2025-01-21T07:15:00Z",
    "session_completed": {
      "duration_minutes": 45,
      "is_valid": true,
      "is_morning_active": true,
      "goals_progressed": [
        {
          "goal_id": "550e8400-e29b-41d4-a716-446655440300",
          "title": "毎日読書30分",
          "progress_updated": true
        }
      ],
      "achievements_earned": [
        {
          "type": "daily_goal",
          "description": "本日の目標達成"
        }
      ]
    }
  },
  "message": "Session completed successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 2. 参加サマリー管理 API

### ユーザー参加サマリー取得

```
GET /users/{userId}/attendance/summaries
```

**権限**: 👤 本人のみ

**説明**: ユーザーの日次参加サマリーを取得。カレンダー表示で使用

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| year | integer | No | 年 (デフォルト: 現在年) |
| month | integer | No | 月 (デフォルト: 現在月) |
| from_date | string | No | 開始日 (YYYY-MM-DD) |
| to_date | string | No | 終了日 (YYYY-MM-DD) |
| view_type | string | No | 表示タイプ (calendar, list) デフォルト: calendar |

**レスポンス:**

```json
{
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "display_name": "田中太郎",
      "username": "tanaka_taro"
    },
    "period": {
      "year": 2025,
      "month": 1,
      "from_date": "2025-01-01",
      "to_date": "2025-01-31"
    },
    "summaries": [
      {
        "date": "2025-01-21",
        "total_duration_minutes": 45,
        "session_count": 1,
        "first_join_time": "06:30:00",
        "last_leave_time": "07:15:00",
        "is_morning_active": true,
        "quality_metrics": {
          "consistency_score": 85,
          "goal_alignment": 100,
          "community_engagement": 78
        },
        "goals_progressed": [
          {
            "goal_id": "550e8400-e29b-41d4-a716-446655440300",
            "title": "毎日読書30分",
            "target_minutes": 30,
            "actual_minutes": 45,
            "achievement_rate": 150
          }
        ],
        "events_attended": [
          {
            "event_id": "550e8400-e29b-41d4-a716-446655440200",
            "title": "朝の読書会",
            "duration_minutes": 45
          }
        ],
        "milestones": [
          {
            "type": "streak_continuation",
            "description": "3日連続参加",
            "achievement_level": "good"
          }
        ]
      },
      {
        "date": "2025-01-20",
        "total_duration_minutes": 35,
        "session_count": 1,
        "first_join_time": "06:15:00",
        "last_leave_time": "06:50:00",
        "is_morning_active": true,
        "quality_metrics": {
          "consistency_score": 92,
          "goal_alignment": 100,
          "community_engagement": 65
        },
        "goals_progressed": [
          {
            "goal_id": "550e8400-e29b-41d4-a716-446655440300",
            "title": "毎日読書30分",
            "target_minutes": 30,
            "actual_minutes": 35,
            "achievement_rate": 117
          }
        ],
        "events_attended": [],
        "milestones": []
      },
      {
        "date": "2025-01-19",
        "total_duration_minutes": 0,
        "session_count": 0,
        "first_join_time": null,
        "last_leave_time": null,
        "is_morning_active": false,
        "quality_metrics": null,
        "goals_progressed": [],
        "events_attended": [],
        "milestones": [],
        "absence_reason": "weekend"
      }
    ],
    "calendar_data": {
      "attendance_days": 18,
      "total_days": 21,
      "attendance_rate": 85.7,
      "streak_days": 3,
      "best_day": {
        "date": "2025-01-15",
        "duration_minutes": 120,
        "quality_score": 95
      },
      "patterns": {
        "most_active_hour": "06:30",
        "average_session_duration": 42,
        "preferred_channels": ["朝の読書会", "一般朝活"]
      }
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 月間カレンダーデータ取得

```
GET /users/{userId}/attendance/calendar
```

**権限**: 👤 本人のみ

**説明**: カレンダー表示用の月間参加データを最適化して取得

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| year | integer | Yes | 年 |
| month | integer | Yes | 月 |

**レスポンス:**

```json
{
  "data": {
    "calendar": {
      "year": 2025,
      "month": 1,
      "days": [
        {
          "date": 1,
          "status": "attended",
          "duration_minutes": 45,
          "quality_level": "high",
          "session_count": 1,
          "goals_achieved": 1,
          "events_count": 1
        },
        {
          "date": 2,
          "status": "attended",
          "duration_minutes": 30,
          "quality_level": "medium",
          "session_count": 1,
          "goals_achieved": 1,
          "events_count": 0
        },
        {
          "date": 3,
          "status": "missed",
          "duration_minutes": 0,
          "quality_level": null,
          "session_count": 0,
          "goals_achieved": 0,
          "events_count": 0
        }
      ]
    },
    "summary": {
      "total_attendance_days": 18,
      "total_duration_minutes": 810,
      "average_duration": 45,
      "current_streak": 3,
      "best_streak": 7,
      "goals_achievement_rate": 88.9,
      "quality_distribution": {
        "high": 8,
        "medium": 7,
        "low": 3
      }
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 3. 参加統計 API

### ユーザー参加統計取得

```
GET /users/{userId}/attendance/statistics
```

**権限**: 🔐 認証済み（公開統計のみ、本人の場合は詳細も含む）

**説明**: ユーザーの参加統計（総日数、連続日数、称号進捗等）を取得

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| period | string | No | 期間 (week, month, quarter, year, all) デフォルト: all |
| include_trends | boolean | No | トレンド分析を含める (デフォルト: false) |
| include_goals | boolean | No | 目標との関連統計を含める (デフォルト: false) |

**レスポンス:**

```json
{
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "display_name": "田中太郎",
      "username": "tanaka_taro",
      "current_title": {
        "level": 3,
        "name_jp": "陽光探求者",
        "color_theme": "#FFB800"
      }
    },
    "overall_statistics": {
      "total_attendance_days": 45,
      "current_streak_days": 3,
      "max_streak_days": 15,
      "first_attendance_date": "2024-10-15",
      "last_attendance_date": "2025-01-21",
      "total_duration_minutes": 2025,
      "average_session_duration": 45,
      "attendance_rate": 78.3,
      "consistency_score": 82.1
    },
    "period_statistics": {
      "period": "month",
      "period_range": {
        "from": "2025-01-01",
        "to": "2025-01-31"
      },
      "attendance_days": 18,
      "total_duration_minutes": 810,
      "session_count": 20,
      "average_daily_duration": 45,
      "quality_metrics": {
        "morning_active_rate": 94.4,
        "goal_achievement_rate": 88.9,
        "community_engagement_score": 76.5
      }
    },
    "streaks": {
      "current_streak": {
        "days": 3,
        "start_date": "2025-01-19",
        "status": "active"
      },
      "longest_streak": {
        "days": 15,
        "start_date": "2024-11-01",
        "end_date": "2024-11-15",
        "achievement_unlocked": "強い意志の証"
      },
      "recent_streaks": [
        {
          "days": 7,
          "start_date": "2024-12-20",
          "end_date": "2024-12-26"
        },
        {
          "days": 12,
          "start_date": "2024-12-01",
          "end_date": "2024-12-12"
        }
      ]
    },
    "title_progress": {
      "current_title": {
        "id": "550e8400-e29b-41d4-a716-446655440003",
        "level": 3,
        "name_jp": "陽光探求者",
        "achieved_at": "2024-11-05T06:30:00Z",
        "days_held": 77
      },
      "next_title": {
        "id": "550e8400-e29b-41d4-a716-446655440004",
        "level": 4,
        "name_jp": "朝陽の使者",
        "required_days": 30,
        "current_progress": 28,
        "remaining_days": 2,
        "progress_rate": 93.3,
        "estimated_achievement": "2025-01-23"
      }
    },
    "rankings": {
      "total_days_rank": 47,
      "current_streak_rank": 23,
      "monthly_duration_rank": 15,
      "consistency_rank": 31
    },
    "achievements": [
      {
        "type": "milestone",
        "title": "初心者卒業",
        "description": "15日連続参加達成",
        "achieved_at": "2024-11-15T06:30:00Z",
        "icon": "🏆"
      },
      {
        "type": "quality",
        "title": "朝活エキスパート",
        "description": "月間平均45分以上参加",
        "achieved_at": "2024-12-31T23:59:59Z",
        "icon": "⭐"
      }
    ],
    "trends": {
      "weekly_pattern": {
        "monday": 85.7,
        "tuesday": 92.3,
        "wednesday": 76.9,
        "thursday": 88.5,
        "friday": 80.8,
        "saturday": 69.2,
        "sunday": 73.1
      },
      "monthly_trend": "improving",
      "duration_trend": "stable",
      "consistency_trend": "improving"
    },
    "goals_relationship": {
      "goals_count": 2,
      "goals_supported_by_attendance": 2,
      "average_goal_achievement_rate": 91.5,
      "most_progressed_goal": {
        "id": "550e8400-e29b-41d4-a716-446655440300",
        "title": "毎日読書30分",
        "contribution_rate": 95.2
      }
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 4. ランキング API

### 月間参加日数ランキング

```
GET /ranking/monthly
```

**権限**: 🔐 認証済み

**説明**: 当月の参加日数ランキングを取得。ライバル表示も含む

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| year | integer | No | 年 (デフォルト: 現在年) |
| month | integer | No | 月 (デフォルト: 現在月) |
| limit | integer | No | 取得件数 (デフォルト: 20, 最大: 100) |
| include_rivals | boolean | No | ライバルのランキングも含める (デフォルト: true) |

**レスポンス:**

```json
{
  "data": {
    "period": {
      "year": 2025,
      "month": 1,
      "from_date": "2025-01-01",
      "to_date": "2025-01-31"
    },
    "rankings": [
      {
        "rank": 1,
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440010",
          "display_name": "佐藤花子",
          "username": "sato_hanako",
          "avatar_url": "https://img.clerk.com/sato.png",
          "current_title": {
            "level": 6,
            "name_jp": "朝活の覇者",
            "color_theme": "#E74C3C"
          }
        },
        "attendance_days": 21,
        "total_duration_minutes": 1260,
        "average_duration": 60,
        "current_streak": 21,
        "quality_score": 94.5,
        "is_current_user": false,
        "is_rival": true,
        "change_from_last_month": "+2"
      },
      {
        "rank": 2,
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440011",
          "display_name": "山田太郎",
          "username": "yamada_taro",
          "avatar_url": "https://img.clerk.com/yamada.png",
          "current_title": {
            "level": 5,
            "name_jp": "暁の守護者",
            "color_theme": "#8E44AD"
          }
        },
        "attendance_days": 20,
        "total_duration_minutes": 1100,
        "average_duration": 55,
        "current_streak": 8,
        "quality_score": 88.2,
        "is_current_user": false,
        "is_rival": false,
        "change_from_last_month": "0"
      },
      {
        "rank": 15,
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
        "attendance_days": 18,
        "total_duration_minutes": 810,
        "average_duration": 45,
        "current_streak": 3,
        "quality_score": 82.1,
        "is_current_user": true,
        "is_rival": false,
        "change_from_last_month": "+3"
      }
    ],
    "current_user_stats": {
      "rank": 15,
      "attendance_days": 18,
      "percentile": 67.3,
      "gap_to_next_rank": {
        "rank": 14,
        "attendance_days_needed": 1,
        "achievable_by": "2025-01-22"
      },
      "rivals_comparison": [
        {
          "rival": {
            "id": "550e8400-e29b-41d4-a716-446655440010",
            "display_name": "佐藤花子",
            "rank": 1
          },
          "gap": 3,
          "status": "behind"
        }
      ]
    },
    "summary": {
      "total_participants": 147,
      "average_attendance_days": 12.3,
      "top_10_average": 19.8,
      "median_attendance_days": 11
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 総合参加日数ランキング

```
GET /ranking/total
```

**権限**: 🔐 認証済み

**説明**: 総合参加日数ランキングを取得。歴代順位表示

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| limit | integer | No | 取得件数 (デフォルト: 20, 最大: 100) |
| include_inactive | boolean | No | 非アクティブユーザーも含める (デフォルト: false) |

**レスポンス:**

```json
{
  "data": {
    "rankings": [
      {
        "rank": 1,
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440020",
          "display_name": "鈴木一郎",
          "username": "suzuki_ichiro",
          "avatar_url": "https://img.clerk.com/suzuki.png",
          "current_title": {
            "level": 8,
            "name_jp": "永遠の夜明け",
            "color_theme": "#9B59B6"
          }
        },
        "total_attendance_days": 412,
        "first_attendance_date": "2024-01-15",
        "total_duration_minutes": 18540,
        "max_streak_days": 89,
        "current_streak": 45,
        "achievement_rate": 98.3,
        "is_current_user": false,
        "member_since_days": 371
      },
      {
        "rank": 47,
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
        "total_attendance_days": 45,
        "first_attendance_date": "2024-10-15",
        "total_duration_minutes": 2025,
        "max_streak_days": 15,
        "current_streak": 3,
        "achievement_rate": 78.3,
        "is_current_user": true,
        "member_since_days": 98
      }
    ],
    "current_user_stats": {
      "rank": 47,
      "total_attendance_days": 45,
      "percentile": 68.1,
      "hall_of_fame_status": "rising_star",
      "next_milestone": {
        "attendance_days": 60,
        "title_unlock": "暁の守護者",
        "estimated_date": "2025-02-15"
      }
    },
    "hall_of_fame": [
      {
        "category": "longest_streak",
        "record_holder": "鈴木一郎",
        "value": 89,
        "achievement_date": "2024-11-30"
      },
      {
        "category": "fastest_to_100_days",
        "record_holder": "佐藤花子",
        "value": 105,
        "achievement_date": "2024-06-15"
      }
    ]
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 連続参加日数ランキング

```
GET /ranking/streak
```

**権限**: 🔐 認証済み

**説明**: 連続参加日数ランキングを取得。現在の連続記録で順位付け

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| type | string | No | ランキングタイプ (current, max) デフォルト: current |
| limit | integer | No | 取得件数 (デフォルト: 20, 最大: 100) |
| min_streak | integer | No | 最低連続日数でのフィルタ (デフォルト: 1) |

**レスポンス:**

```json
{
  "data": {
    "ranking_type": "current",
    "rankings": [
      {
        "rank": 1,
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440020",
          "display_name": "鈴木一郎",
          "username": "suzuki_ichiro",
          "avatar_url": "https://img.clerk.com/suzuki.png",
          "current_title": {
            "level": 8,
            "name_jp": "永遠の夜明け",
            "color_theme": "#9B59B6"
          }
        },
        "current_streak_days": 45,
        "streak_start_date": "2024-12-07",
        "max_streak_days": 89,
        "streak_quality": {
          "consistency_rate": 100.0,
          "morning_active_rate": 97.8,
          "average_duration": 58
        },
        "streak_milestones": [
          {
            "days": 30,
            "achieved_at": "2025-01-05",
            "reward": "朝陽の使者獲得"
          }
        ],
        "is_current_user": false
      },
      {
        "rank": 23,
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
        "current_streak_days": 3,
        "streak_start_date": "2025-01-19",
        "max_streak_days": 15,
        "streak_quality": {
          "consistency_rate": 100.0,
          "morning_active_rate": 100.0,
          "average_duration": 42
        },
        "streak_milestones": [],
        "is_current_user": true
      }
    ],
    "current_user_stats": {
      "current_streak_rank": 23,
      "max_streak_rank": 15,
      "current_streak_days": 3,
      "max_streak_days": 15,
      "streak_potential": "high",
      "next_streak_milestone": {
        "days": 5,
        "reward": "継続の力",
        "estimated_date": "2025-01-23"
      }
    },
    "streak_insights": {
      "danger_zone_threshold": 1,
      "users_in_danger": 23,
      "average_current_streak": 8.4,
      "longest_active_streak": 45,
      "streak_distribution": {
        "1-7_days": 89,
        "8-14_days": 34,
        "15-30_days": 18,
        "31+_days": 6
      }
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
| ATTENDANCE_LOG_NOT_FOUND | 参加ログが見つからない | 404 |
| INVALID_DISCORD_USER | 無効なDiscordユーザーID | 422 |
| INVALID_CHANNEL | 無効なDiscordチャンネル | 422 |
| DUPLICATE_LOG_ENTRY | 重複するログエントリ | 409 |
| SESSION_NOT_ACTIVE | アクティブなセッションがない | 422 |
| INVALID_TIMESTAMP | 無効なタイムスタンプ | 422 |
| BOT_AUTH_REQUIRED | Bot認証が必要 | 401 |
| STATISTICS_NOT_AVAILABLE | 統計データが利用できない | 404 |

## ビジネスルール（ドメイン層実装）

### ドメインルール
```go
// internal/domain/attendance/errors.go
var (
    ErrInvalidDiscordUser    = errors.New("無効なDiscordユーザーIDです")
    ErrInvalidChannel        = errors.New("無効なDiscordチャンネルです")
    ErrSessionNotActive      = errors.New("アクティブなセッションがありません")
    ErrDuplicateLogEntry     = errors.New("重複するログエントリです")
    ErrInvalidTimeRange      = errors.New("無効な時間範囲です")
    ErrMinimumDurationNotMet = errors.New("最低参加時間を満たしていません")
)

// internal/domain/attendance/entity/attendance_log.go
func (al *AttendanceLog) Validate() error {
    if al.UserID == uuid.Nil {
        return errors.New("ユーザーIDは必須です")
    }
    if al.DiscordChannelID == "" {
        return errors.New("DiscordチャンネルIDは必須です")
    }
    if al.JoinedAt.After(time.Now()) {
        return errors.New("未来の参加時刻は設定できません")
    }
    if al.LeftAt != nil && al.LeftAt.Before(al.JoinedAt) {
        return ErrInvalidTimeRange
    }
    return nil
}

func (al *AttendanceLog) CalculateDuration() int {
    if al.LeftAt == nil {
        return 0
    }
    duration := al.LeftAt.Sub(al.JoinedAt)
    return int(duration.Minutes())
}

func (al *AttendanceLog) IsMorningActive() bool {
    hour := al.JoinedAt.Hour()
    return hour >= 6 && hour < 7
}

func (al *AttendanceLog) IsValidSession(minDuration int) bool {
    return al.DurationMinutes >= minDuration && al.IsMorningActive()
}

// internal/domain/attendance/entity/attendance_summary.go
func (as *AttendanceSummary) Validate() error {
    if as.UserID == uuid.Nil {
        return errors.New("ユーザーIDは必須です")
    }
    if as.Date.After(time.Now().Truncate(24 * time.Hour)) {
        return errors.New("未来日のサマリーは作成できません")
    }
    if as.SessionCount < 0 || as.TotalDurationMinutes < 0 {
        return errors.New("セッション数と参加時間は0以上である必要があります")
    }
    return nil
}

func (as *AttendanceSummary) CalculateQualityScore() float64 {
    if as.SessionCount == 0 {
        return 0
    }
    
    avgDuration := float64(as.TotalDurationMinutes) / float64(as.SessionCount)
    qualityScore := avgDuration / 60.0 * 100 // 1時間を基準とした品質スコア
    
    if as.IsMorningActive {
        qualityScore *= 1.2 // 朝活ボーナス
    }
    
    if qualityScore > 100 {
        qualityScore = 100
    }
    
    return qualityScore
}

// internal/domain/attendance/entity/attendance_statistics.go
func (stats *AttendanceStatistics) UpdateStreak(attendanceDate time.Time) {
    if stats.LastAttendanceDate == nil {
        stats.CurrentStreakDays = 1
        stats.MaxStreakDays = 1
        stats.LastAttendanceDate = &attendanceDate
        return
    }
    
    lastDate := *stats.LastAttendanceDate
    daysDiff := int(attendanceDate.Sub(lastDate).Hours() / 24)
    
    if daysDiff == 1 {
        // 連続参加
        stats.CurrentStreakDays++
        if stats.CurrentStreakDays > stats.MaxStreakDays {
            stats.MaxStreakDays = stats.CurrentStreakDays
        }
    } else if daysDiff > 1 {
        // 連続記録リセット
        stats.CurrentStreakDays = 1
    }
    
    stats.LastAttendanceDate = &attendanceDate
}

func (stats *AttendanceStatistics) CalculateAttendanceRate(totalPossibleDays int) float64 {
    if totalPossibleDays == 0 {
        return 0
    }
    return float64(stats.TotalAttendanceDays) / float64(totalPossibleDays) * 100
}
```

### 参加ログ管理のルール
1. **朝活時間判定**: 6:00-7:00の参加開始を朝活として認定（AttendanceLog.IsMorningActive()で制御）
2. **最低参加時間**: 設定可能な最低時間（デフォルト0分）（AttendanceLog.IsValidSession()で制御）
3. **重複セッション防止**: 同一時間帯の重複参加ログ防止（リポジトリ層で制御）
4. **自動統計更新**: 参加ログ作成時に統計を自動更新（ドメインサービスで制御）

### 統計計算のルール
1. **連続日数計算**: 日単位での連続参加判定（AttendanceStatistics.UpdateStreak()で制御）
2. **品質スコア**: 参加時間と朝活率に基づく品質評価（AttendanceSummary.CalculateQualityScore()で制御）
3. **ランキング更新**: 日次バッチで順位を再計算
4. **データ整合性**: 統計データと実ログの整合性を定期チェック

## 注意事項

### Discord連携
- Discord Botからの認証はAPIキー方式
- 入退室ログは自動記録、手動補正も可能
- チャンネル制限による有効性判定

### プライバシー考慮
- 詳細ログは本人のみアクセス可能
- 公開統計は基本情報のみ
- ランキングは匿名化オプション対応

### パフォーマンス考慮
- 統計データは事前計算でキャッシュ
- ログデータは適切なインデックス設計
- ランキングは定期更新で負荷分散

### キャッシュ戦略
- 個人統計: 10分キャッシュ
- ランキングデータ: 1時間キャッシュ
- カレンダーデータ: 5分キャッシュ
- 月間サマリー: 30分キャッシュ