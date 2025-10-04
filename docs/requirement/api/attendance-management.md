# Attendance Management API - 詳細設計

## Overview
Ghoona Campアプリケーションの朝活参加記録管理API詳細設計書です。
Discord参加ログの自動記録・基本統計・ランキング機能を提供します。

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
        "discord_channel_id": "1234567890123456789",
        "joined_at": "2025-01-21T06:30:00Z",
        "left_at": "2025-01-21T07:15:00Z",
        "duration_minutes": 45,
        "is_valid": true,
        "is_morning_active": true,
        "created_at": "2025-01-21T06:30:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440101",
        "event": null,
        "discord_channel_id": "1234567890123456790",
        "joined_at": "2025-01-20T06:15:00Z",
        "left_at": "2025-01-20T06:50:00Z",
        "duration_minutes": 35,
        "is_valid": true,
        "is_morning_active": true,
        "created_at": "2025-01-20T06:15:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440102",
        "event": null,
        "discord_channel_id": "1234567890123456791",
        "joined_at": "2025-01-19T05:45:00Z",
        "left_at": "2025-01-19T06:05:00Z",
        "duration_minutes": 20,
        "is_valid": false,
        "is_morning_active": false,
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
        "events_attended": [
          {
            "event_id": "550e8400-e29b-41d4-a716-446655440200",
            "title": "朝の読書会",
            "duration_minutes": 45
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
        "events_attended": []
      },
      {
        "date": "2025-01-19",
        "total_duration_minutes": 0,
        "session_count": 0,
        "first_join_time": null,
        "last_leave_time": null,
        "is_morning_active": false,
        "events_attended": []
      }
    ],
    "calendar_data": {
      "attendance_days": 18,
      "total_days": 21,
      "attendance_rate": 85.7,
      "current_streak": 3,
      "max_streak": 7
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
      }
    },
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
        "is_current_user": true,
        "is_rival": false,
        "change_from_last_month": "+3"
      }
    ],
    "current_user_stats": {
      "rank": 15,
      "attendance_days": 18
    },
    "summary": {
      "total_participants": 147
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
        "is_current_user": true,
        "member_since_days": 98
      }
    ],
    "current_user_stats": {
      "rank": 47,
      "total_attendance_days": 45
    },
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
| limit | integer | No | 取得件数 (デフォルト: 20, 最大: 100) |

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
        "current_streak_days": 45,
        "streak_start_date": "2024-12-07",
        "max_streak_days": 89,
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
        "is_current_user": true
      }
    ],
    "current_user_stats": {
      "current_streak_rank": 23,
      "current_streak_days": 3,
      "max_streak_days": 15
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

## 基本ルール

### 参加ログ管理ルール
1. **Discord参加記録**: Botによる自動記録（6:00-7:00の参加開始が有効）
2. **最低参加時間**: 制限なし（短時間でも有効）
3. **有効性判定**: 朝活時間内（6:00-7:00）の参加開始
4. **統計更新**: 日次バッチ処理で自動計算

### ランキングシステム
- **月間ランキング**: 当月の参加日数
- **総合ランキング**: 累計参加日数
- **連続記録**: 現在の連続参加日数
- **更新頻度**: 毎日深夜0時に更新

### アクセス制御
- **参加ログ**: 本人のみ詳細閲覧可能
- **統計データ**: 本人のみ詳細、他者は基本統計のみ
- **ランキング**: 全ユーザー閲覧可能
- **Bot API**: Discord Bot専用認証
