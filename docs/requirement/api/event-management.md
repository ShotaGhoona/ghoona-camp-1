# Event Management API - 詳細設計

## Overview
Ghoona Campアプリケーションのイベント管理関連API詳細設計書です。
朝活イベントの作成・参加管理・Discord連携を中心とした機能を提供します。

## アーキテクチャ適合性
このAPI設計は以下のアーキテクチャ原則に従います：
- **オニオンアーキテクチャ**: Domain → Application → Infrastructure → Interface層の依存関係
- **ドメイン駆動設計**: `event`コンテキストを中心としたエンティティ設計
- **クリーンな境界**: DTO、ドメインエンティティ、GORMモデルの適切な分離

## ドメイン設計対応

### ドメインエンティティ
```go
// internal/domain/event/entity/event.go
type Event struct {
    ID                UUID
    CreatorID         UUID
    Title             string
    Description       string
    EventType         EventType
    ScheduledDate     time.Time
    StartTime         time.Time
    EndTime           time.Time
    MaxParticipants   *int
    IsRecurring       bool
    RecurrencePattern *RecurrencePattern
    DiscordChannelID  *string
    IsActive          bool
    CreatedAt         time.Time
    UpdatedAt         time.Time
}

// internal/domain/event/entity/event_participant.go
type EventParticipant struct {
    ID        UUID
    EventID   UUID
    UserID    UUID
    Status    ParticipantStatus
    CreatedAt time.Time
    UpdatedAt time.Time
}

// internal/domain/event/value/event_type.go
type EventType string
const (
    EventTypeGeneral     EventType = "general"
    EventTypeStudy       EventType = "study"
    EventTypeExercise    EventType = "exercise"
    EventTypeMeditation  EventType = "meditation"
    EventTypeCreative    EventType = "creative"
    EventTypeBusiness    EventType = "business"
)

// internal/domain/event/value/participant_status.go
type ParticipantStatus string
const (
    ParticipantStatusRegistered ParticipantStatus = "registered"
    ParticipantStatusCancelled  ParticipantStatus = "cancelled"
    ParticipantStatusAttended   ParticipantStatus = "attended"
    ParticipantStatusNoShow     ParticipantStatus = "no_show"
)

// internal/domain/event/value/recurrence_pattern.go
type RecurrencePattern string
const (
    RecurrenceDaily   RecurrencePattern = "daily"
    RecurrenceWeekly  RecurrencePattern = "weekly"
    RecurrenceMonthly RecurrencePattern = "monthly"
)
```

### リポジトリインターフェース
```go
// internal/domain/event/repository/event_repository.go
type EventRepository interface {
    GetByID(ctx context.Context, id UUID) (*entity.Event, error)
    FindAll(ctx context.Context, filters EventFilters) ([]*entity.Event, error)
    GetByCreatorID(ctx context.Context, creatorID UUID) ([]*entity.Event, error)
    FindUpcoming(ctx context.Context, limit int) ([]*entity.Event, error)
    Create(ctx context.Context, event *entity.Event) error
    Update(ctx context.Context, event *entity.Event) error
    Delete(ctx context.Context, id UUID) error
}

type EventParticipantRepository interface {
    GetByEventID(ctx context.Context, eventID UUID) ([]*entity.EventParticipant, error)
    GetByUserID(ctx context.Context, userID UUID) ([]*entity.EventParticipant, error)
    GetByEventAndUser(ctx context.Context, eventID, userID UUID) (*entity.EventParticipant, error)
    Create(ctx context.Context, participant *entity.EventParticipant) error
    Update(ctx context.Context, participant *entity.EventParticipant) error
    Delete(ctx context.Context, id UUID) error
    CountByEventID(ctx context.Context, eventID UUID) (int, error)
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

## 1. イベント管理 API

### イベント一覧取得

```
GET /events
```

**権限**: 🔐 認証済み

**説明**: 開催予定・開催中のイベント一覧を取得。日時順でソート

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| date_from | string | No | 開始日フィルタ (YYYY-MM-DD) |
| date_to | string | No | 終了日フィルタ (YYYY-MM-DD) |
| event_type | string | No | イベントタイプフィルタ (general, study, exercise, etc.) |
| creator_id | UUID | No | 作成者でのフィルタ |
| status | string | No | イベント状態 (upcoming, ongoing, completed) |
| search | string | No | タイトル・説明での部分検索 |
| page | integer | No | ページ番号 (デフォルト: 1) |
| limit | integer | No | 1ページあたりの件数 (デフォルト: 20, 最大: 50) |
| sort_by | string | No | ソート基準 (scheduled_date, created_at, participants_count) |
| order | string | No | ソート順 (asc, desc) デフォルト: asc |

**レスポンス:**

```json
{
  "data": {
    "events": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440200",
        "creator": {
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
        "title": "朝の読書会",
        "description": "技術書を一緒に読んで議論しましょう。今週はVue.js公式ドキュメントを読みます。",
        "event_type": "study",
        "scheduled_date": "2025-01-23",
        "start_time": "06:30:00",
        "end_time": "07:30:00",
        "max_participants": 10,
        "is_recurring": true,
        "recurrence_pattern": "weekly",
        "discord_channel_id": "1234567890123456789",
        "participants": {
          "registered_count": 6,
          "max_participants": 10,
          "available_slots": 4,
          "is_full": false
        },
        "user_participation": {
          "is_registered": true,
          "status": "registered",
          "registered_at": "2025-01-20T15:30:00Z"
        },
        "is_active": true,
        "created_at": "2025-01-15T10:30:00Z",
        "updated_at": "2025-01-20T16:45:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440201",
        "creator": {
          "id": "550e8400-e29b-41d4-a716-446655440001",
          "display_name": "佐藤花子",
          "username": "sato_hanako",
          "avatar_url": "https://img.clerk.com/sato.png",
          "current_title": {
            "level": 4,
            "name_jp": "朝陽の使者",
            "color_theme": "#FF6B35"
          }
        },
        "title": "朝ヨガセッション",
        "description": "心と体を整える朝ヨガで一日をスタートしましょう。初心者も大歓迎です。",
        "event_type": "exercise",
        "scheduled_date": "2025-01-24",
        "start_time": "06:00:00",
        "end_time": "06:45:00",
        "max_participants": 8,
        "is_recurring": false,
        "recurrence_pattern": null,
        "discord_channel_id": "1234567890123456790",
        "participants": {
          "registered_count": 8,
          "max_participants": 8,
          "available_slots": 0,
          "is_full": true
        },
        "user_participation": {
          "is_registered": false,
          "status": null,
          "registered_at": null
        },
        "is_active": true,
        "created_at": "2025-01-18T14:20:00Z",
        "updated_at": "2025-01-21T09:15:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 2,
      "total_count": 25,
      "limit": 20,
      "has_next": true,
      "has_prev": false
    },
    "summary": {
      "upcoming_events": 15,
      "ongoing_events": 2,
      "my_registered_events": 3,
      "available_events": 8
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### イベント作成

```
POST /events
```

**権限**: 🔐 認証済み

**説明**: 新しい朝活イベントを作成。タイトル、説明、日時、最大参加者数等を設定

**リクエスト:**

```json
{
  "title": "朝のプログラミング勉強会",
  "description": "みんなでプログラミングの課題に取り組みましょう。今回はReactのフック機能について学習します。質問や議論も大歓迎です。",
  "event_type": "study",
  "scheduled_date": "2025-01-25",
  "start_time": "06:30:00",
  "end_time": "07:30:00",
  "max_participants": 12,
  "is_recurring": false,
  "recurrence_pattern": null,
  "discord_channel_id": "1234567890123456791"
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| title | 必須、最大200文字 |
| description | 最大2000文字 |
| event_type | 必須、有効なEventType |
| scheduled_date | 必須、YYYY-MM-DD形式、今日以降 |
| start_time | 必須、HH:MM:SS形式 |
| end_time | 必須、HH:MM:SS形式、start_time以降 |
| max_participants | 1以上の整数、最大100 |
| recurrence_pattern | is_recurringがtrueの場合必須 |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440202",
    "creator": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "display_name": "田中太郎",
      "username": "tanaka_taro"
    },
    "title": "朝のプログラミング勉強会",
    "description": "みんなでプログラミングの課題に取り組みましょう。今回はReactのフック機能について学習します。質問や議論も大歓迎です。",
    "event_type": "study",
    "scheduled_date": "2025-01-25",
    "start_time": "06:30:00",
    "end_time": "07:30:00",
    "max_participants": 12,
    "is_recurring": false,
    "recurrence_pattern": null,
    "discord_channel_id": "1234567890123456791",
    "participants": {
      "registered_count": 0,
      "max_participants": 12,
      "available_slots": 12,
      "is_full": false
    },
    "is_active": true,
    "created_at": "2025-01-21T10:00:00Z",
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Event created successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### イベント詳細取得

```
GET /events/{eventId}
```

**権限**: 🔐 認証済み

**説明**: 指定イベントの詳細情報を取得。参加者一覧、説明、Discord情報等

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| eventId | UUID | Yes | イベントID |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440200",
    "creator": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "display_name": "田中太郎",
      "username": "tanaka_taro",
      "avatar_url": "https://img.clerk.com/preview.png",
      "bio": "エンジニアとして働きながら朝活を通じて成長中",
      "current_title": {
        "level": 3,
        "name_jp": "陽光探求者",
        "color_theme": "#FFB800"
      }
    },
    "title": "朝の読書会",
    "description": "技術書を一緒に読んで議論しましょう。今週はVue.js公式ドキュメントを読みます。参加者同士で疑問点を共有し、理解を深めていきましょう。",
    "event_type": "study",
    "scheduled_date": "2025-01-23",
    "start_time": "06:30:00",
    "end_time": "07:30:00",
    "max_participants": 10,
    "is_recurring": true,
    "recurrence_pattern": "weekly",
    "discord_channel_id": "1234567890123456789",
    "discord_info": {
      "channel_name": "朝の読書会",
      "channel_url": "https://discord.gg/abc123",
      "is_channel_active": true
    },
    "participants": {
      "registered_count": 6,
      "max_participants": 10,
      "available_slots": 4,
      "is_full": false,
      "waiting_list_count": 0
    },
    "user_participation": {
      "is_registered": true,
      "status": "registered",
      "registered_at": "2025-01-20T15:30:00Z",
      "can_cancel": true
    },
    "schedule_info": {
      "is_upcoming": true,
      "is_ongoing": false,
      "is_completed": false,
      "time_until_start": "2 days 8 hours",
      "duration_minutes": 60
    },
    "related_events": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440203",
        "title": "朝の読書会",
        "scheduled_date": "2025-01-30",
        "start_time": "06:30:00",
        "is_recurring_instance": true
      }
    ],
    "is_active": true,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-20T16:45:00Z"
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### イベント更新

```
PUT /events/{eventId}
```

**権限**: 👑 作成者のみ

**説明**: イベント情報を更新。作成者のみが実行可能

**リクエスト:**

```json
{
  "title": "朝の読書会（上級者向け）",
  "description": "技術書を一緒に読んで議論しましょう。今週はVue.js公式ドキュメントの応用編を読みます。上級者向けの内容になります。",
  "max_participants": 8,
  "discord_channel_id": "1234567890123456789"
}
```

**制限事項:**
- 開催日時の変更は開催24時間前まで
- 参加者がいる場合の定員減少は不可
- イベントタイプの変更は不可

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440200",
    "title": "朝の読書会（上級者向け）",
    "description": "技術書を一緒に読んで議論しましょう。今週はVue.js公式ドキュメントの応用編を読みます。上級者向けの内容になります。",
    "max_participants": 8,
    "discord_channel_id": "1234567890123456789",
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Event updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### イベント削除

```
DELETE /events/{eventId}
```

**権限**: 👑 作成者のみ

**説明**: イベントを削除。作成者のみが実行可能

**制限事項:**
- 参加者がいる場合は削除不可（キャンセル処理が必要）
- 開催済みのイベントは論理削除のみ

**レスポンス:**

```json
{
  "message": "Event deleted successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 2. イベント参加管理 API

### イベント参加者一覧取得

```
GET /events/{eventId}/participants
```

**権限**: 🔐 認証済み

**説明**: イベント参加者一覧を取得。参加状況確認で使用

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| status | string | No | 参加状態フィルタ (registered, cancelled, attended, no_show) |
| sort_by | string | No | ソート基準 (registered_at, name) |
| order | string | No | ソート順 (asc, desc) デフォルト: asc |

**レスポンス:**

```json
{
  "data": {
    "event": {
      "id": "550e8400-e29b-41d4-a716-446655440200",
      "title": "朝の読書会",
      "scheduled_date": "2025-01-23",
      "max_participants": 10
    },
    "participants": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440300",
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
        "status": "registered",
        "registered_at": "2025-01-20T15:30:00Z",
        "updated_at": "2025-01-20T15:30:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440301",
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440001",
          "display_name": "佐藤花子",
          "username": "sato_hanako",
          "avatar_url": "https://img.clerk.com/sato.png",
          "current_title": {
            "level": 4,
            "name_jp": "朝陽の使者",
            "color_theme": "#FF6B35"
          }
        },
        "status": "registered",
        "registered_at": "2025-01-19T09:15:00Z",
        "updated_at": "2025-01-19T09:15:00Z"
      }
    ],
    "summary": {
      "total_participants": 6,
      "registered": 5,
      "cancelled": 1,
      "attended": 0,
      "no_show": 0,
      "available_slots": 4
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### イベント参加申込

```
POST /events/{eventId}/participants
```

**権限**: 🔐 認証済み

**説明**: イベントに参加申込を行う。定員チェックも実行

**リクエスト:**

```json
{
  "note": "Vue.js初心者ですが、よろしくお願いします！"
}
```

**バリデーション:**
- イベントが有効でアクティブである
- 定員に空きがある
- 同一ユーザーの重複参加不可
- 開催日時が未来である

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440302",
    "event": {
      "id": "550e8400-e29b-41d4-a716-446655440200",
      "title": "朝の読書会",
      "scheduled_date": "2025-01-23",
      "start_time": "06:30:00"
    },
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440002",
      "display_name": "鈴木一郎",
      "username": "suzuki_ichiro"
    },
    "status": "registered",
    "registered_at": "2025-01-21T10:00:00Z",
    "reminder_info": {
      "will_send_reminder": true,
      "reminder_time": "2025-01-23T06:00:00Z"
    }
  },
  "message": "Event registration successful",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 参加ステータス更新

```
PUT /events/{eventId}/participants/{userId}
```

**権限**: 👤 本人のみ

**説明**: 参加ステータスを更新（参加→キャンセル等）

**リクエスト:**

```json
{
  "status": "cancelled",
  "reason": "急用のため参加できなくなりました"
}
```

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440302",
    "status": "cancelled",
    "updated_at": "2025-01-21T10:00:00Z",
    "event_slots": {
      "available_slots": 5,
      "waiting_list_notified": false
    }
  },
  "message": "Participation status updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### イベント参加キャンセル

```
DELETE /events/{eventId}/participants/{userId}
```

**権限**: 👤 本人のみ

**説明**: イベント参加をキャンセル

**制限事項:**
- 開催2時間前まで可能
- キャンセル済みの参加は削除不可

**レスポンス:**

```json
{
  "message": "Event participation cancelled successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 3. イベント検索・フィルタ API

### イベント種別一覧取得

```
GET /events/types
```

**権限**: 🔐 認証済み

**説明**: 利用可能なイベントタイプ一覧を取得

**レスポンス:**

```json
{
  "data": {
    "event_types": [
      {
        "type": "general",
        "name_jp": "一般",
        "name_en": "General",
        "description": "一般的な朝活イベント",
        "icon": "🌅",
        "color": "#6B7280"
      },
      {
        "type": "study",
        "name_jp": "学習",
        "name_en": "Study",
        "description": "読書・プログラミング等の学習系",
        "icon": "📚",
        "color": "#3B82F6"
      },
      {
        "type": "exercise",
        "name_jp": "運動",
        "name_en": "Exercise", 
        "description": "ヨガ・ストレッチ・筋トレ等",
        "icon": "💪",
        "color": "#EF4444"
      },
      {
        "type": "meditation",
        "name_jp": "瞑想",
        "name_en": "Meditation",
        "description": "瞑想・マインドフルネス",
        "icon": "🧘",
        "color": "#8B5CF6"
      },
      {
        "type": "creative",
        "name_jp": "創作",
        "name_en": "Creative",
        "description": "絵画・音楽・執筆等の創作活動",
        "icon": "🎨",
        "color": "#F59E0B"
      },
      {
        "type": "business",
        "name_jp": "ビジネス",
        "name_en": "Business",
        "description": "起業・副業・スキルアップ",
        "icon": "💼",
        "color": "#10B981"
      }
    ]
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 人気イベント取得

```
GET /events/popular
```

**権限**: 🔐 認証済み

**説明**: 参加者数の多い人気イベントを取得

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| period | string | No | 集計期間 (week, month, quarter) デフォルト: month |
| limit | integer | No | 取得件数 (デフォルト: 10, 最大: 20) |

**レスポンス:**

```json
{
  "data": {
    "period": "month",
    "popular_events": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440200",
        "title": "朝の読書会",
        "event_type": "study",
        "creator": {
          "display_name": "田中太郎",
          "username": "tanaka_taro"
        },
        "total_participants": 45,
        "average_rating": 4.8,
        "recurrence_count": 6,
        "next_scheduled_date": "2025-01-23"
      }
    ]
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
| EVENT_NOT_FOUND | イベントが見つからない | 404 |
| EVENT_ACCESS_DENIED | イベントへのアクセス権限なし | 403 |
| EVENT_VALIDATION_ERROR | イベントデータが無効 | 422 |
| EVENT_FULL | イベントが満員 | 409 |
| REGISTRATION_DUPLICATE | 既に参加登録済み | 409 |
| REGISTRATION_DEADLINE_PASSED | 参加登録期限切れ | 422 |
| CANCELLATION_DEADLINE_PASSED | キャンセル期限切れ | 422 |
| EVENT_NOT_MODIFIABLE | イベント変更不可 | 422 |
| PARTICIPANT_NOT_FOUND | 参加者が見つからない | 404 |

## ビジネスルール（ドメイン層実装）

### ドメインルール
```go
// internal/domain/event/errors.go
var (
    ErrEventNotFound           = errors.New("イベントが見つかりません")
    ErrEventFull              = errors.New("イベントが満員です")
    ErrDuplicateRegistration  = errors.New("既に参加登録済みです")
    ErrRegistrationDeadline   = errors.New("参加登録期限を過ぎています")
    ErrCancellationDeadline   = errors.New("キャンセル期限を過ぎています")
    ErrInvalidEventType       = errors.New("無効なイベントタイプです")
    ErrInvalidTimeRange       = errors.New("開始時間は終了時間より前である必要があります")
    ErrPastEventDate          = errors.New("過去の日付にはイベントを作成できません")
)

// internal/domain/event/entity/event.go
func (e *Event) Validate() error {
    if e.Title == "" {
        return errors.New("イベントタイトルは必須です")
    }
    if len(e.Title) > 200 {
        return errors.New("イベントタイトルは200文字以内で入力してください")
    }
    if !e.EventType.IsValid() {
        return ErrInvalidEventType
    }
    if e.StartTime.After(e.EndTime) || e.StartTime.Equal(e.EndTime) {
        return ErrInvalidTimeRange
    }
    if e.ScheduledDate.Before(time.Now().Truncate(24 * time.Hour)) {
        return ErrPastEventDate
    }
    if e.MaxParticipants != nil && *e.MaxParticipants < 1 {
        return errors.New("最大参加者数は1以上である必要があります")
    }
    if e.MaxParticipants != nil && *e.MaxParticipants > 100 {
        return errors.New("最大参加者数は100以下である必要があります")
    }
    return nil
}

func (e *Event) CanBeModifiedBy(userID UUID) bool {
    return e.CreatorID == userID
}

func (e *Event) IsRegistrationOpen() bool {
    now := time.Now()
    eventDateTime := time.Date(
        e.ScheduledDate.Year(), e.ScheduledDate.Month(), e.ScheduledDate.Day(),
        e.StartTime.Hour(), e.StartTime.Minute(), e.StartTime.Second(), 0,
        e.ScheduledDate.Location(),
    )
    return now.Before(eventDateTime) && e.IsActive
}

func (e *Event) IsCancellationAllowed() bool {
    now := time.Now()
    eventDateTime := time.Date(
        e.ScheduledDate.Year(), e.ScheduledDate.Month(), e.ScheduledDate.Day(),
        e.StartTime.Hour(), e.StartTime.Minute(), e.StartTime.Second(), 0,
        e.ScheduledDate.Location(),
    )
    // 開催2時間前まで
    return now.Before(eventDateTime.Add(-2 * time.Hour))
}

// internal/domain/event/value/event_type.go
func (et EventType) IsValid() bool {
    switch et {
    case EventTypeGeneral, EventTypeStudy, EventTypeExercise, 
         EventTypeMeditation, EventTypeCreative, EventTypeBusiness:
        return true
    default:
        return false
    }
}

// internal/domain/event/entity/event_participant.go
func (ep *EventParticipant) Validate() error {
    if ep.EventID == uuid.Nil || ep.UserID == uuid.Nil {
        return errors.New("イベントIDとユーザーIDは必須です")
    }
    if !ep.Status.IsValid() {
        return errors.New("無効な参加ステータスです")
    }
    return nil
}

func (ep *EventParticipant) CanBeCancelled() bool {
    return ep.Status == ParticipantStatusRegistered
}

// internal/domain/event/value/participant_status.go
func (ps ParticipantStatus) IsValid() bool {
    switch ps {
    case ParticipantStatusRegistered, ParticipantStatusCancelled,
         ParticipantStatusAttended, ParticipantStatusNoShow:
        return true
    default:
        return false
    }
}
```

### イベント管理のルール
1. **イベント作成制限**: 過去の日付にはイベント作成不可（Event.Validate()で制御）
2. **定員管理**: 最大参加者数は1-100人の範囲（Event.Validate()で制御）
3. **参加登録期限**: イベント開始時刻まで（Event.IsRegistrationOpen()で制御）
4. **キャンセル期限**: イベント開始2時間前まで（Event.IsCancellationAllowed()で制御）
5. **重複参加防止**: 同一イベントへの重複参加不可（ドメインサービスで制御）
6. **イベント変更権限**: 作成者のみ変更可能（Event.CanBeModifiedBy()で制御）

## 注意事項

### Discord連携
- Discord チャンネルIDは任意設定
- 参加者には Discord チャンネル情報を通知
- Discord 参加状況は attendance_logs テーブルで管理

### 定期イベント
- 定期イベントは個別インスタンスとして作成
- 元イベントの変更は将来のインスタンスにのみ適用
- 過去のインスタンスは変更不可

### キャッシュ戦略
- イベント一覧: 5分キャッシュ
- イベント詳細: 1分キャッシュ
- 参加者情報: リアルタイム更新