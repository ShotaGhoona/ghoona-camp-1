# Notification Management API - 詳細設計

## Overview
Ghoona Campアプリケーションの基本的な通知管理API設計書です。
朝活コミュニティの基本通知（称号獲得・リマインダー・イベント）を提供します。


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
| type | string | No | 通知タイプフィルタ (achievement, reminder, event) |
| is_read | boolean | No | 既読状態でのフィルタ |
| from_date | string | No | 開始日 (YYYY-MM-DD) |
| to_date | string | No | 終了日 (YYYY-MM-DD) |
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
        "data": {
          "title_id": "550e8400-e29b-41d4-a716-446655440004",
          "title_name": "朝陽の使者",
          "title_level": 4
        },
        "is_read": false,
        "sent_at": "2025-01-21T06:30:00Z",
        "created_at": "2025-01-21T06:30:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440101",
        "type": "reminder",
        "title": "🌅 朝活の時間です",
        "message": "おはようございます！今日も朝活を始めましょう。",
        "priority": "medium",
        "data": {
          "current_streak": 3
        },
        "is_read": true,
        "sent_at": "2025-01-21T06:00:00Z",
        "created_at": "2025-01-20T21:00:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440103",
        "type": "event",
        "title": "📚 参加予定のイベント開始まで30分",
        "message": "「朝の読書会」が30分後に開始されます。準備はいかがですか？今日のテーマは「TypeScriptの実践的活用法」です。",
        "data": {
          "event_id": "550e8400-e29b-41d4-a716-446655440200",
          "event_title": "朝の読書会",
          "start_time": "2025-01-21T06:30:00Z"
        },
        "is_read": false,
        "sent_at": "2025-01-21T06:00:00Z",
        "created_at": "2025-01-20T18:00:00Z"
      },
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
        "event": 6
      }
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


**レスポンス:**

```json
{
  "data": {
    "marked_count": 12,
    "remaining_unread": 0
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
      "event_reminder_enabled": true
    },
    "schedule_settings": {
      "reminder_time": "21:00:00"
    },
    "channel_preferences": {
      "in_app": {
        "enabled": true
      }
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
    "event_reminder_enabled": true
  },
  "schedule_settings": {
    "reminder_time": "21:30:00"
  },
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| reminder_time | HH:MM:SS形式 |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440500",
    "notification_types": {
      "achievement_enabled": true,
      "reminder_enabled": true,
      "event_reminder_enabled": true
    },
    "changes_applied": [
      "reminder_time: 21:00:00 → 21:30:00"
    ],
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Notification settings updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 3. システム通知送信 API

### 通知送信（バッチ処理用）

```
POST /notifications/send
```

**権限**: 🔐 システム認証

**説明**: バッチ処理からの通知送信。称号獲得通知等で使用

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
  "data": {
    "title_id": "550e8400-e29b-41d4-a716-446655440004",
    "title_name": "朝陽の使者",
    "title_level": 4
  }
}
```

**レスポンス:**

```json
{
  "data": {
    "notification_id": "550e8400-e29b-41d4-a716-446655440105"
  },
  "message": "Notification sent successfully",
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

## 基本ルール

### 通知タイプ
- **achievement**: 称号獲得通知
- **reminder**: 朝活リマインダー
- **event**: イベント関連通知

### 通知設定ルール
1. **基本設定**: 通知タイプ別のON/OFF設定
2. **リマインダー時刻**: 夜のリマインダー送信時刻設定
3. **アプリ内通知**: 基本的なアプリ内通知のみ

### アクセス制御
- **通知一覧**: 本人のみ閲覧可能
- **通知設定**: 本人のみ更新可能
- **システム送信**: バッチ処理のみ