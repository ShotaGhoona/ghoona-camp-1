# Notification Management API - 詳細設計

## Overview
Ghoona Campアプリケーションの通知・リマインダー管理関連API詳細設計書です。
称号獲得・リマインダー・ライバル更新・イベント通知を中心とした通知システムを提供します。

## アーキテクチャ適合性
このAPI設計は以下のアーキテクチャ原則に従います：
- **オニオンアーキテクチャ**: Domain → Application → Infrastructure → Interface層の依存関係
- **ドメイン駆動設計**: `notification`コンテキストを中心としたエンティティ設計
- **クリーンな境界**: DTO、ドメインエンティティ、GORMモデルの適切な分離

## ドメイン設計対応

### ドメインエンティティ
```go
// internal/domain/notification/entity/notification.go
type Notification struct {
    ID          UUID
    UserID      UUID
    Type        NotificationType
    Title       string
    Message     string
    Data        map[string]interface{}
    IsRead      bool
    ScheduledAt *time.Time
    SentAt      *time.Time
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// internal/domain/notification/entity/notification_settings.go
type NotificationSettings struct {
    ID                     UUID
    UserID                 UUID
    AchievementEnabled     bool
    ReminderEnabled        bool
    RivalUpdateEnabled     bool
    EventReminderEnabled   bool
    GoalProgressEnabled    bool
    CommunityEnabled       bool
    ReminderTime           time.Time
    QuietHoursStart        *time.Time
    QuietHoursEnd          *time.Time
    PreferredChannels      []NotificationChannel
    CreatedAt              time.Time
    UpdatedAt              time.Time
}

// internal/domain/notification/value/notification_type.go
type NotificationType string
const (
    NotificationAchievement  NotificationType = "achievement"
    NotificationReminder     NotificationType = "reminder"
    NotificationRivalUpdate  NotificationType = "rival_update"
    NotificationEvent        NotificationType = "event"
    NotificationGoalProgress NotificationType = "goal_progress"
    NotificationCommunity    NotificationType = "community"
    NotificationSystem       NotificationType = "system"
)

// internal/domain/notification/value/notification_channel.go
type NotificationChannel string
const (
    ChannelInApp    NotificationChannel = "in_app"
    ChannelEmail    NotificationChannel = "email"
    ChannelPush     NotificationChannel = "push"
    ChannelDiscord  NotificationChannel = "discord"
)

// internal/domain/notification/value/notification_priority.go
type NotificationPriority string
const (
    PriorityLow    NotificationPriority = "low"
    PriorityMedium NotificationPriority = "medium"
    PriorityHigh   NotificationPriority = "high"
    PriorityUrgent NotificationPriority = "urgent"
)
```

### リポジトリインターフェース
```go
// internal/domain/notification/repository/notification_repository.go
type NotificationRepository interface {
    FindByUserID(ctx context.Context, userID UUID, filters NotificationFilters) ([]*entity.Notification, error)
    FindByID(ctx context.Context, id UUID) (*entity.Notification, error)
    Create(ctx context.Context, notification *entity.Notification) error
    Update(ctx context.Context, notification *entity.Notification) error
    Delete(ctx context.Context, id UUID) error
    MarkAsRead(ctx context.Context, id UUID) error
    MarkAllAsRead(ctx context.Context, userID UUID) error
    GetUnreadCount(ctx context.Context, userID UUID) (int, error)
}

type NotificationSettingsRepository interface {
    FindByUserID(ctx context.Context, userID UUID) (*entity.NotificationSettings, error)
    Create(ctx context.Context, settings *entity.NotificationSettings) error
    Update(ctx context.Context, settings *entity.NotificationSettings) error
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

## 1. 通知管理 API

### ユーザー通知一覧取得

```
GET /users/{userId}/notifications
```

**権限**: 👤 本人のみ

**説明**: ユーザーの通知一覧を取得。称号獲得、ライバル更新、イベント等

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| userId | UUID | Yes | ユーザーID |

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| type | string | No | 通知タイプフィルタ (achievement, reminder, rival_update, event, goal_progress, community, system) |
| is_read | boolean | No | 既読状態でのフィルタ |
| from_date | string | No | 開始日 (YYYY-MM-DD) |
| to_date | string | No | 終了日 (YYYY-MM-DD) |
| priority | string | No | 優先度フィルタ (low, medium, high, urgent) |
| page | integer | No | ページ番号 (デフォルト: 1) |
| limit | integer | No | 1ページあたりの件数 (デフォルト: 20, 最大: 50) |
| sort_by | string | No | ソート基準 (created_at, sent_at) デフォルト: created_at |
| order | string | No | ソート順 (asc, desc) デフォルト: desc |

**レスポンス:**

```json
{
  "data": {
    "notifications": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440100",
        "type": "achievement",
        "title": "🏆 新しい称号を獲得しました！",
        "message": "おめでとうございます！「朝陽の使者」の称号を獲得しました。30日間の継続的な朝活参加により、コミュニティの模範となる存在に成長されました。",
        "priority": "high",
        "data": {
          "achievement_type": "title",
          "title_id": "550e8400-e29b-41d4-a716-446655440004",
          "title_name": "朝陽の使者",
          "title_level": 4,
          "title_color": "#FF6B35",
          "achievement_date": "2025-01-21T06:30:00Z",
          "streak_days": 30,
          "celebration_animation": "golden_burst"
        },
        "is_read": false,
        "action_buttons": [
          {
            "type": "view_title",
            "label": "称号を確認",
            "action": "/titles/550e8400-e29b-41d4-a716-446655440004"
          },
          {
            "type": "share",
            "label": "シェア",
            "action": "share_achievement"
          }
        ],
        "scheduled_at": null,
        "sent_at": "2025-01-21T06:30:00Z",
        "created_at": "2025-01-21T06:30:00Z",
        "expires_at": "2025-02-21T06:30:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440101",
        "type": "reminder",
        "title": "🌅 朝活の時間です",
        "message": "おはようございます！今日も素晴らしい朝活を始めましょう。あなたの目標「毎日読書30分」まで残り2分で達成です。",
        "priority": "medium",
        "data": {
          "reminder_type": "daily_morning",
          "current_streak": 3,
          "goals_today": [
            {
              "goal_id": "550e8400-e29b-41d4-a716-446655440300",
              "title": "毎日読書30分",
              "progress_today": 28,
              "target": 30,
              "remaining": 2
            }
          ],
          "weather": {
            "condition": "sunny",
            "temperature": 12,
            "message": "快晴で朝活に最適な天気です"
          },
          "motivation_quote": "継続は力なり。今日も一歩前進しましょう。"
        },
        "is_read": true,
        "action_buttons": [
          {
            "type": "join_discord",
            "label": "朝活に参加",
            "action": "discord://join/1234567890123456789"
          },
          {
            "type": "snooze",
            "label": "10分後に再通知",
            "action": "snooze_10min"
          }
        ],
        "scheduled_at": "2025-01-21T06:00:00Z",
        "sent_at": "2025-01-21T06:00:00Z",
        "created_at": "2025-01-20T21:00:00Z",
        "expires_at": "2025-01-21T12:00:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440102",
        "type": "rival_update",
        "title": "🔥 ライバルが新記録を達成",
        "message": "ライバルの佐藤花子さんが15日連続参加を達成しました！あなたの現在の記録は3日です。負けずに頑張りましょう！",
        "priority": "medium",
        "data": {
          "rival_user": {
            "id": "550e8400-e29b-41d4-a716-446655440010",
            "display_name": "佐藤花子",
            "username": "sato_hanako",
            "avatar_url": "https://img.clerk.com/sato.png",
            "current_title": {
              "level": 4,
              "name_jp": "朝陽の使者"
            }
          },
          "achievement_type": "streak_milestone",
          "milestone_value": 15,
          "your_current_value": 3,
          "gap": 12,
          "encouragement_level": "motivational"
        },
        "is_read": false,
        "action_buttons": [
          {
            "type": "view_rival",
            "label": "ライバルを見る",
            "action": "/users/550e8400-e29b-41d4-a716-446655440010"
          },
          {
            "type": "join_challenge",
            "label": "チャレンジ参加",
            "action": "join_morning_challenge"
          }
        ],
        "scheduled_at": null,
        "sent_at": "2025-01-20T07:30:00Z",
        "created_at": "2025-01-20T07:30:00Z",
        "expires_at": "2025-01-27T07:30:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440103",
        "type": "event",
        "title": "📚 参加予定のイベント開始まで30分",
        "message": "「朝の読書会」が30分後に開始されます。準備はいかがですか？今日のテーマは「TypeScriptの実践的活用法」です。",
        "priority": "high",
        "data": {
          "event": {
            "id": "550e8400-e29b-41d4-a716-446655440200",
            "title": "朝の読書会",
            "start_time": "2025-01-21T06:30:00Z",
            "duration_minutes": 60,
            "creator": "田中花子",
            "participants_count": 8,
            "max_participants": 10
          },
          "reminder_type": "30_minutes_before",
          "preparation_items": [
            "TypeScript実践入門の第4章を準備",
            "DiscordのVCチャンネルをチェック",
            "質問があれば事前にメモしておく"
          ],
          "discord_channel": {
            "id": "1234567890123456789",
            "name": "朝の読書会",
            "invite_url": "https://discord.gg/abc123"
          }
        },
        "is_read": false,
        "action_buttons": [
          {
            "type": "join_event",
            "label": "イベントに参加",
            "action": "discord://join/1234567890123456789"
          },
          {
            "type": "view_event",
            "label": "詳細を見る",
            "action": "/events/550e8400-e29b-41d4-a716-446655440200"
          },
          {
            "type": "cancel_participation",
            "label": "参加をキャンセル",
            "action": "cancel_event_participation"
          }
        ],
        "scheduled_at": "2025-01-21T06:00:00Z",
        "sent_at": "2025-01-21T06:00:00Z",
        "created_at": "2025-01-20T18:00:00Z",
        "expires_at": "2025-01-21T07:30:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440104",
        "type": "goal_progress",
        "title": "🎯 目標達成まであと少し！",
        "message": "「毎日読書30分」の今月の達成率が85%に到達しました。月末まで5日、あと3日参加すれば90%達成です！",
        "priority": "medium",
        "data": {
          "goal": {
            "id": "550e8400-e29b-41d4-a716-446655440300",
            "title": "毎日読書30分",
            "period": "monthly",
            "current_achievement_rate": 85.0,
            "target_achievement_rate": 90.0,
            "days_remaining": 5,
            "sessions_needed": 3
          },
          "progress_trend": "improving",
          "milestone_type": "achievement_rate",
          "encouragement_message": "順調な進捗です！この調子で継続しましょう。",
          "next_milestone": {
            "value": 90,
            "reward": "月間達成バッジ",
            "estimated_date": "2025-01-29"
          }
        },
        "is_read": true,
        "action_buttons": [
          {
            "type": "view_goal",
            "label": "目標を確認",
            "action": "/goals/550e8400-e29b-41d4-a716-446655440300"
          },
          {
            "type": "log_progress",
            "label": "今日の進捗を記録",
            "action": "log_goal_progress"
          }
        ],
        "scheduled_at": null,
        "sent_at": "2025-01-20T20:00:00Z",
        "created_at": "2025-01-20T20:00:00Z",
        "expires_at": "2025-02-01T00:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 3,
      "total_count": 47,
      "limit": 20,
      "has_next": true,
      "has_prev": false
    },
    "summary": {
      "total_notifications": 47,
      "unread_count": 12,
      "by_type": {
        "achievement": 3,
        "reminder": 25,
        "rival_update": 8,
        "event": 6,
        "goal_progress": 4,
        "community": 1
      },
      "urgent_count": 0,
      "today_count": 5
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 通知既読更新

```
PUT /notifications/{notificationId}
```

**権限**: 👤 本人のみ

**説明**: 通知の既読ステータスを更新

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| notificationId | UUID | Yes | 通知ID |

**リクエスト:**

```json
{
  "is_read": true,
  "action_taken": "view_title"
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| is_read | 必須、boolean |
| action_taken | 任意、実行されたアクション |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440100",
    "is_read": true,
    "read_at": "2025-01-21T10:00:00Z",
    "action_taken": "view_title",
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Notification status updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 通知削除

```
DELETE /notifications/{notificationId}
```

**権限**: 👤 本人のみ

**説明**: 不要な通知を削除

**制限事項:**
- 未読の重要な通知（urgent）は削除不可
- システム通知は削除不可
- 削除後は復元不可

**レスポンス:**

```json
{
  "message": "Notification deleted successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 全通知既読化

```
PUT /users/{userId}/notifications/mark-all-read
```

**権限**: 👤 本人のみ

**説明**: ユーザーの全未読通知を既読にする

**リクエスト:**

```json
{
  "types": ["achievement", "reminder", "rival_update"],
  "before_date": "2025-01-21"
}
```

**レスポンス:**

```json
{
  "data": {
    "marked_count": 12,
    "remaining_unread": 0,
    "excluded_urgent_count": 0
  },
  "message": "All notifications marked as read",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 2. 通知設定 API

### 通知設定取得

```
GET /users/{userId}/notification-settings
```

**権限**: 👤 本人のみ

**説明**: ユーザーの通知設定を取得。各種通知のON/OFF、時刻設定等

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440500",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "notification_types": {
      "achievement_enabled": true,
      "reminder_enabled": true,
      "rival_update_enabled": true,
      "event_reminder_enabled": true,
      "goal_progress_enabled": true,
      "community_enabled": false
    },
    "schedule_settings": {
      "reminder_time": "21:00:00",
      "morning_reminder_time": "06:00:00",
      "quiet_hours": {
        "enabled": true,
        "start": "22:00:00",
        "end": "05:30:00"
      },
      "weekend_reminders": false,
      "holiday_reminders": false
    },
    "channel_preferences": {
      "in_app": {
        "enabled": true,
        "types": ["achievement", "reminder", "rival_update", "event", "goal_progress", "community"]
      },
      "email": {
        "enabled": true,
        "types": ["achievement", "event"],
        "digest_frequency": "weekly",
        "email_address": "tanaka@example.com"
      },
      "push": {
        "enabled": true,
        "types": ["reminder", "event"],
        "sound_enabled": true,
        "vibration_enabled": true
      },
      "discord": {
        "enabled": false,
        "types": [],
        "dm_enabled": false,
        "server_notifications": false
      }
    },
    "advanced_settings": {
      "achievement_celebration": {
        "enabled": true,
        "animation_level": "full",
        "sound_effects": true,
        "share_prompt": true
      },
      "rival_updates": {
        "frequency": "immediate",
        "achievements_only": false,
        "streak_updates": true,
        "goal_progress": true
      },
      "event_reminders": {
        "advance_times": ["24h", "1h", "30m"],
        "custom_reminders": true,
        "creator_updates": true,
        "participant_updates": false
      },
      "goal_progress": {
        "milestone_notifications": true,
        "daily_summary": false,
        "weekly_summary": true,
        "trend_alerts": true
      }
    },
    "personalization": {
      "notification_tone": "gentle",
      "motivational_level": "high",
      "language": "ja",
      "timezone": "Asia/Tokyo",
      "cultural_context": "japanese"
    },
    "created_at": "2024-10-15T10:30:00Z",
    "updated_at": "2025-01-15T14:20:00Z"
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 通知設定更新

```
PUT /users/{userId}/notification-settings
```

**権限**: 👤 本人のみ

**説明**: 通知設定を更新。リマインダー時刻、通知種別の有効/無効等

**リクエスト:**

```json
{
  "notification_types": {
    "achievement_enabled": true,
    "reminder_enabled": true,
    "rival_update_enabled": false,
    "event_reminder_enabled": true,
    "goal_progress_enabled": true,
    "community_enabled": true
  },
  "schedule_settings": {
    "reminder_time": "21:30:00",
    "morning_reminder_time": "05:45:00",
    "quiet_hours": {
      "enabled": true,
      "start": "22:30:00",
      "end": "05:30:00"
    },
    "weekend_reminders": true,
    "holiday_reminders": false
  },
  "channel_preferences": {
    "email": {
      "enabled": true,
      "types": ["achievement", "event", "goal_progress"],
      "digest_frequency": "daily"
    },
    "push": {
      "enabled": true,
      "types": ["reminder", "event", "achievement"],
      "sound_enabled": false,
      "vibration_enabled": true
    }
  },
  "advanced_settings": {
    "rival_updates": {
      "frequency": "daily_digest",
      "achievements_only": true
    },
    "event_reminders": {
      "advance_times": ["1h", "30m"],
      "custom_reminders": false
    }
  },
  "personalization": {
    "notification_tone": "energetic",
    "motivational_level": "medium"
  }
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| reminder_time | HH:MM:SS形式 |
| quiet_hours.start/end | HH:MM:SS形式 |
| channel_preferences.email.types | 有効な通知タイプの配列 |
| personalization.notification_tone | enum(gentle, standard, energetic) |
| personalization.motivational_level | enum(low, medium, high) |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440500",
    "notification_types": {
      "achievement_enabled": true,
      "reminder_enabled": true,
      "rival_update_enabled": false,
      "event_reminder_enabled": true,
      "goal_progress_enabled": true,
      "community_enabled": true
    },
    "changes_applied": [
      "rival_update_enabled: true → false",
      "community_enabled: false → true",
      "reminder_time: 21:00:00 → 21:30:00",
      "notification_tone: gentle → energetic"
    ],
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Notification settings updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 3. 通知配信・スケジュール API

### 即座通知送信（システム用）

```
POST /notifications/send
```

**権限**: 🔐 システム認証

**説明**: システムからの即座通知送信。称号獲得、イベント更新等で使用

**認証ヘッダー:**
```
X-System-Key: your-system-api-key
```

**リクエスト:**

```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "type": "achievement",
  "title": "🏆 新しい称号を獲得しました！",
  "message": "おめでとうございます！「朝陽の使者」の称号を獲得しました。",
  "priority": "high",
  "data": {
    "achievement_type": "title",
    "title_id": "550e8400-e29b-41d4-a716-446655440004",
    "title_name": "朝陽の使者",
    "title_level": 4,
    "celebration_animation": "golden_burst"
  },
  "channels": ["in_app", "push", "email"],
  "action_buttons": [
    {
      "type": "view_title",
      "label": "称号を確認",
      "action": "/titles/550e8400-e29b-41d4-a716-446655440004"
    }
  ],
  "expires_at": "2025-02-21T06:30:00Z"
}
```

**レスポンス:**

```json
{
  "data": {
    "notification_id": "550e8400-e29b-41d4-a716-446655440105",
    "delivery_status": {
      "in_app": {
        "status": "delivered",
        "delivered_at": "2025-01-21T10:00:00Z"
      },
      "push": {
        "status": "delivered",
        "delivered_at": "2025-01-21T10:00:01Z",
        "device_tokens": 2
      },
      "email": {
        "status": "queued",
        "scheduled_for": "2025-01-21T10:05:00Z"
      }
    },
    "user_settings_applied": {
      "quiet_hours_checked": false,
      "user_preferences_respected": true,
      "filtered_channels": []
    }
  },
  "message": "Notification sent successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### スケジュール通知作成（システム用）

```
POST /notifications/schedule
```

**権限**: 🔐 システム認証

**説明**: スケジュール通知の作成。リマインダー、イベント通知等で使用

**リクエスト:**

```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "type": "reminder",
  "title": "🌅 朝活の時間です",
  "message": "おはようございます！今日も素晴らしい朝活を始めましょう。",
  "priority": "medium",
  "scheduled_at": "2025-01-22T06:00:00Z",
  "data": {
    "reminder_type": "daily_morning",
    "current_streak": 4,
    "weather": {
      "condition": "cloudy",
      "temperature": 8
    }
  },
  "channels": ["in_app", "push"],
  "recurring": {
    "enabled": true,
    "pattern": "daily",
    "end_date": "2025-12-31",
    "skip_weekends": false,
    "skip_holidays": true
  }
}
```

**レスポンス:**

```json
{
  "data": {
    "notification_id": "550e8400-e29b-41d4-a716-446655440106",
    "scheduled_at": "2025-01-22T06:00:00Z",
    "recurring_schedule": {
      "next_5_deliveries": [
        "2025-01-22T06:00:00Z",
        "2025-01-23T06:00:00Z",
        "2025-01-24T06:00:00Z",
        "2025-01-27T06:00:00Z",
        "2025-01-28T06:00:00Z"
      ],
      "skipped_dates": ["2025-01-25", "2025-01-26"],
      "total_scheduled": 340
    },
    "estimated_delivery_channels": ["in_app", "push"]
  },
  "message": "Notification scheduled successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 通知配信状況確認

```
GET /notifications/{notificationId}/delivery-status
```

**権限**: 🔐 システム認証

**説明**: 特定通知の配信状況を確認

**レスポンス:**

```json
{
  "data": {
    "notification_id": "550e8400-e29b-41d4-a716-446655440105",
    "overall_status": "partially_delivered",
    "created_at": "2025-01-21T10:00:00Z",
    "channels": {
      "in_app": {
        "status": "delivered",
        "delivered_at": "2025-01-21T10:00:00Z",
        "read_at": "2025-01-21T10:05:23Z",
        "action_taken": "view_title"
      },
      "push": {
        "status": "delivered",
        "delivered_at": "2025-01-21T10:00:01Z",
        "devices": [
          {
            "platform": "ios",
            "status": "delivered",
            "delivered_at": "2025-01-21T10:00:01Z"
          },
          {
            "platform": "android",
            "status": "failed",
            "error": "invalid_token",
            "attempted_at": "2025-01-21T10:00:01Z"
          }
        ]
      },
      "email": {
        "status": "delivered",
        "delivered_at": "2025-01-21T10:05:00Z",
        "opened_at": "2025-01-21T11:23:45Z",
        "clicked": true,
        "bounce_reason": null
      },
      "discord": {
        "status": "skipped",
        "reason": "user_preference_disabled"
      }
    },
    "user_interaction": {
      "first_seen_at": "2025-01-21T10:05:23Z",
      "actions_taken": ["view_title"],
      "engagement_score": 85
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 4. 通知テンプレート・個人化 API

### 通知テンプレート一覧

```
GET /notification-templates
```

**権限**: 🔐 システム認証

**説明**: 利用可能な通知テンプレート一覧を取得

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| type | string | No | 通知タイプでのフィルタ |
| language | string | No | 言語でのフィルタ (ja, en) |
| active_only | boolean | No | アクティブなテンプレートのみ |

**レスポンス:**

```json
{
  "data": {
    "templates": [
      {
        "id": "achievement_title_unlocked",
        "type": "achievement",
        "name": "称号獲得通知",
        "language": "ja",
        "title_template": "🏆 新しい称号「{title_name}」を獲得しました！",
        "message_template": "おめでとうございます！{streak_days}日間の継続により「{title_name}」の称号を獲得しました。{encouragement_message}",
        "variables": [
          {
            "name": "title_name",
            "type": "string",
            "required": true,
            "description": "獲得した称号名"
          },
          {
            "name": "streak_days",
            "type": "integer",
            "required": true,
            "description": "連続参加日数"
          },
          {
            "name": "encouragement_message",
            "type": "string",
            "required": false,
            "description": "励ましメッセージ"
          }
        ],
        "personalization_options": {
          "tone": ["gentle", "standard", "energetic"],
          "motivational_level": ["low", "medium", "high"],
          "celebration_style": ["minimal", "standard", "full"]
        },
        "action_buttons": [
          {
            "type": "view_title",
            "label_template": "「{title_name}」を確認",
            "action_template": "/titles/{title_id}"
          },
          {
            "type": "share",
            "label_template": "シェア",
            "action_template": "share_achievement"
          }
        ],
        "is_active": true,
        "created_at": "2024-01-01T00:00:00Z"
      }
    ]
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 個人化通知プレビュー

```
POST /notifications/preview
```

**権限**: 👤 本人のみ

**説明**: ユーザーの設定に基づいた通知のプレビューを生成

**リクエスト:**

```json
{
  "template_id": "achievement_title_unlocked",
  "variables": {
    "title_name": "朝陽の使者",
    "streak_days": 30,
    "title_level": 4
  },
  "personalization": {
    "tone": "energetic",
    "motivational_level": "high",
    "celebration_style": "full"
  },
  "channel": "in_app"
}
```

**レスポンス:**

```json
{
  "data": {
    "preview": {
      "title": "🏆 新しい称号「朝陽の使者」を獲得しました！",
      "message": "素晴らしいです！30日間の継続により「朝陽の使者」の称号を獲得しました。あなたの努力が実を結んだ瞬間です！この調子で更なる高みを目指しましょう！",
      "priority": "high",
      "visual_elements": {
        "celebration_animation": "golden_burst_energetic",
        "color_scheme": "#FF6B35",
        "icon": "🏆",
        "background_effect": "sparkles"
      },
      "action_buttons": [
        {
          "type": "view_title",
          "label": "「朝陽の使者」を確認",
          "style": "primary"
        },
        {
          "type": "share",
          "label": "みんなにシェア！",
          "style": "secondary"
        }
      ],
      "sound_effect": "achievement_fanfare",
      "estimated_engagement": 92
    },
    "personalization_applied": {
      "tone_adjustments": "より活発な表現に調整",
      "motivational_boost": "高レベルの励ましメッセージを追加",
      "celebration_enhancement": "フル演出でお祝い効果を最大化"
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
| NOTIFICATION_NOT_FOUND | 通知が見つからない | 404 |
| NOTIFICATION_ACCESS_DENIED | 通知へのアクセス権限なし | 403 |
| INVALID_NOTIFICATION_TYPE | 無効な通知タイプ | 422 |
| SETTINGS_NOT_FOUND | 通知設定が見つからない | 404 |
| QUIET_HOURS_ACTIVE | サイレント時間中 | 422 |
| DELIVERY_FAILED | 通知配信失敗 | 500 |
| TEMPLATE_NOT_FOUND | テンプレートが見つからない | 404 |
| INVALID_SCHEDULE | 無効なスケジュール設定 | 422 |
| RATE_LIMIT_EXCEEDED | 送信制限超過 | 429 |

## ビジネスルール（ドメイン層実装）

### ドメインルール
```go
// internal/domain/notification/errors.go
var (
    ErrNotificationNotFound     = errors.New("通知が見つかりません")
    ErrInvalidNotificationType  = errors.New("無効な通知タイプです")
    ErrQuietHoursActive         = errors.New("サイレント時間中です")
    ErrDeliveryFailed          = errors.New("通知配信に失敗しました")
    ErrInvalidSchedule         = errors.New("無効なスケジュール設定です")
    ErrRateLimitExceeded       = errors.New("送信制限を超過しています")
    ErrTemplateNotFound        = errors.New("テンプレートが見つかりません")
)

// internal/domain/notification/entity/notification.go
func (n *Notification) Validate() error {
    if n.UserID == uuid.Nil {
        return errors.New("ユーザーIDは必須です")
    }
    if n.Title == "" {
        return errors.New("通知タイトルは必須です")
    }
    if len(n.Title) > 100 {
        return errors.New("通知タイトルは100文字以内で入力してください")
    }
    if len(n.Message) > 500 {
        return errors.New("通知メッセージは500文字以内で入力してください")
    }
    if !n.Type.IsValid() {
        return ErrInvalidNotificationType
    }
    return nil
}

func (n *Notification) CanBeDelivered(settings *NotificationSettings) bool {
    if !n.IsTypeEnabled(settings) {
        return false
    }
    if n.IsInQuietHours(settings) {
        return false
    }
    return true
}

func (n *Notification) IsTypeEnabled(settings *NotificationSettings) bool {
    switch n.Type {
    case NotificationAchievement:
        return settings.AchievementEnabled
    case NotificationReminder:
        return settings.ReminderEnabled
    case NotificationRivalUpdate:
        return settings.RivalUpdateEnabled
    case NotificationEvent:
        return settings.EventReminderEnabled
    case NotificationGoalProgress:
        return settings.GoalProgressEnabled
    case NotificationCommunity:
        return settings.CommunityEnabled
    default:
        return true // システム通知は常に配信
    }
}

func (n *Notification) IsInQuietHours(settings *NotificationSettings) bool {
    if settings.QuietHoursStart == nil || settings.QuietHoursEnd == nil {
        return false
    }
    
    now := time.Now()
    start := *settings.QuietHoursStart
    end := *settings.QuietHoursEnd
    
    currentTime := time.Date(0, 1, 1, now.Hour(), now.Minute(), now.Second(), 0, time.UTC)
    
    if start.Before(end) {
        return currentTime.After(start) && currentTime.Before(end)
    } else {
        return currentTime.After(start) || currentTime.Before(end)
    }
}

func (n *Notification) MarkAsRead() {
    n.IsRead = true
    n.UpdatedAt = time.Now()
}

// internal/domain/notification/entity/notification_settings.go
func (ns *NotificationSettings) Validate() error {
    if ns.UserID == uuid.Nil {
        return errors.New("ユーザーIDは必須です")
    }
    if ns.QuietHoursStart != nil && ns.QuietHoursEnd != nil {
        // サイレント時間の妥当性チェック
        if ns.QuietHoursStart.Equal(*ns.QuietHoursEnd) {
            return errors.New("サイレント時間の開始と終了は異なる時刻である必要があります")
        }
    }
    return nil
}

func (ns *NotificationSettings) IsDeliveryAllowed(notificationType NotificationType, currentTime time.Time) bool {
    if !ns.isTypeEnabled(notificationType) {
        return false
    }
    if ns.isInQuietHours(currentTime) {
        return false
    }
    return true
}

// internal/domain/notification/value/notification_type.go
func (nt NotificationType) IsValid() bool {
    switch nt {
    case NotificationAchievement, NotificationReminder, NotificationRivalUpdate,
         NotificationEvent, NotificationGoalProgress, NotificationCommunity, NotificationSystem:
        return true
    default:
        return false
    }
}

func (nt NotificationType) GetDefaultPriority() NotificationPriority {
    switch nt {
    case NotificationAchievement:
        return PriorityHigh
    case NotificationReminder:
        return PriorityMedium
    case NotificationRivalUpdate:
        return PriorityMedium
    case NotificationEvent:
        return PriorityHigh
    case NotificationGoalProgress:
        return PriorityMedium
    case NotificationCommunity:
        return PriorityLow
    case NotificationSystem:
        return PriorityUrgent
    default:
        return PriorityMedium
    }
}

// internal/domain/notification/value/notification_channel.go
func (nc NotificationChannel) IsValid() bool {
    switch nc {
    case ChannelInApp, ChannelEmail, ChannelPush, ChannelDiscord:
        return true
    default:
        return false
    }
}

func (nc NotificationChannel) GetDeliveryDelay() time.Duration {
    switch nc {
    case ChannelInApp:
        return 0
    case ChannelPush:
        return 1 * time.Second
    case ChannelEmail:
        return 5 * time.Minute
    case ChannelDiscord:
        return 2 * time.Second
    default:
        return 0
    }
}
```

### 通知配信のルール
1. **優先度制御**: 通知タイプによる優先度の自動設定（NotificationType.GetDefaultPriority()で制御）
2. **サイレント時間**: ユーザー設定による配信時間制限（Notification.IsInQuietHours()で制御）
3. **チャンネル別遅延**: 配信チャンネルごとの適切な遅延時間（NotificationChannel.GetDeliveryDelay()で制御）
4. **レート制限**: ユーザー・システム別の送信制限（ドメインサービスで制御）

### 通知個人化のルール
1. **設定継承**: デフォルト設定からユーザー設定への階層的継承
2. **文脈適応**: ユーザーの活動パターンに基づく通知タイミング最適化
3. **A/Bテスト**: 通知効果測定のための配信パターン実験
4. **配信最適化**: 開封率・エンゲージメント率に基づく配信チャンネル選択

## 注意事項

### プライバシー考慮
- 通知内容は暗号化して保存
- 削除された通知は完全削除（30日後）
- ユーザー設定の輸出・削除機能

### 配信信頼性
- 配信失敗時の自動リトライ機能
- チャンネル障害時の代替配信
- 配信状況の詳細ログ記録

### パフォーマンス考慮
- 大量配信時のバッチ処理
- 配信キューの負荷分散
- 通知履歴の自動アーカイブ

### キャッシュ戦略
- 通知設定: 30分キャッシュ
- 通知一覧: リアルタイム更新
- 配信テンプレート: 1時間キャッシュ
- 配信状況: 5分キャッシュ