# Event Management API - 詳細設計

## Overview
Ghoona Campアプリケーションのイベント管理関連API詳細設計書です。
朝活コミュニティでのイベント作成・参加管理・Discord連携を中心とした機能を提供します。

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

**説明**: 開催予定・開催中のイベント一覧を取得

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| date_from | string | No | 開始日フィルタ (YYYY-MM-DD) |
| date_to | string | No | 終了日フィルタ (YYYY-MM-DD) |
| event_type | string | No | イベントタイプフィルタ (general, study, exercise, meditation, creative, business) |
| search | string | No | タイトル・説明での部分検索 |
| page | integer | No | ページ番号 (デフォルト: 1) |
| limit | integer | No | 1ページあたりの件数 (デフォルト: 20, 最大: 50) |

**レスポンス:**

```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440200",
      "creator": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "display_name": "田中太郎",
        "username": "tanaka_taro",
        "avatar_url": "https://img.clerk.com/preview.png"
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
      "participant_count": 6,
      "is_user_registered": true,
      "is_active": true,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-20T16:45:00Z"
    }
  ],
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### イベント作成

```
POST /events
```

**権限**: 🔐 認証済み

**説明**: 新しい朝活イベントを作成

**リクエスト:**

```json
{
  "title": "朝のプログラミング勉強会",
  "description": "みんなでプログラミングの課題に取り組みましょう。今回はReactのフック機能について学習します。",
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
| title | 必須、1-200文字 |
| description | 最大2000文字 |
| event_type | 必須、enum(general, study, exercise, meditation, creative, business) |
| scheduled_date | 必須、YYYY-MM-DD形式 |
| start_time | 必須、HH:MM:SS形式 |
| end_time | 必須、HH:MM:SS形式、start_time以降 |
| max_participants | 1-100の整数 |
| discord_channel_id | Discord チャンネルID（文字列） |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440202",
    "title": "朝のプログラミング勉強会",
    "description": "みんなでプログラミングの課題に取り組みましょう。今回はReactのフック機能について学習します。",
    "event_type": "study",
    "scheduled_date": "2025-01-25",
    "start_time": "06:30:00",
    "end_time": "07:30:00",
    "max_participants": 12,
    "is_recurring": false,
    "recurrence_pattern": null,
    "discord_channel_id": "1234567890123456791",
    "participant_count": 0,
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

**説明**: 指定イベントの詳細情報を取得

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
      "avatar_url": "https://img.clerk.com/preview.png"
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
    "participant_count": 6,
    "is_user_registered": true,
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

**説明**: イベント情報を更新

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| eventId | UUID | Yes | イベントID |

**リクエスト:**

```json
{
  "title": "朝の読書会（上級者向け）",
  "description": "技術書を一緒に読んで議論しましょう。今週はVue.js公式ドキュメントの応用編を読みます。",
  "max_participants": 8,
  "discord_channel_id": "1234567890123456789"
}
```

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440200",
    "title": "朝の読書会（上級者向け）",
    "description": "技術書を一緒に読んで議論しましょう。今週はVue.js公式ドキュメントの応用編を読みます。",
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

**説明**: イベントを削除（論理削除）

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| eventId | UUID | Yes | イベントID |

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

**説明**: イベント参加者一覧を取得

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| eventId | UUID | Yes | イベントID |

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| status | string | No | 参加状態フィルタ (registered, cancelled) |

**レスポンス:**

```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440300",
      "user": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "display_name": "田中太郎",
        "username": "tanaka_taro",
        "avatar_url": "https://img.clerk.com/preview.png"
      },
      "status": "registered",
      "created_at": "2025-01-20T15:30:00Z",
      "updated_at": "2025-01-20T15:30:00Z"
    }
  ],
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### イベント参加申込

```
POST /events/{eventId}/participants
```

**権限**: 🔐 認証済み

**説明**: イベントに参加申込を行う

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| eventId | UUID | Yes | イベントID |

**バリデーション:**
- イベントが有効でアクティブである
- 定員に空きがある
- 同一ユーザーの重複参加不可

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440302",
    "event_id": "550e8400-e29b-41d4-a716-446655440200",
    "user_id": "550e8400-e29b-41d4-a716-446655440002",
    "status": "registered",
    "created_at": "2025-01-21T10:00:00Z"
  },
  "message": "Event registration successful",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### イベント参加ステータス更新

```
PUT /events/{eventId}/participants/{userId}
```

**権限**: 👤 本人のみ

**説明**: イベント参加ステータスを更新（参加 ↔ キャンセル）

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| eventId | UUID | Yes | イベントID |
| userId | UUID | Yes | ユーザーID |

**リクエスト:**

```json
{
  "status": "cancelled"
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| status | 必須、enum(registered, cancelled) |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440302",
    "event_id": "550e8400-e29b-41d4-a716-446655440200",
    "user_id": "550e8400-e29b-41d4-a716-446655440002",
    "status": "cancelled",
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Participation status updated successfully",
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
| PARTICIPANT_NOT_FOUND | 参加者が見つからない | 404 |

## 基本ルール

### イベント管理ルール
1. **イベント作成**: タイトル、説明、日時、定員、イベントタイプを設定
2. **参加管理**: 登録・キャンセル機能
3. **権限管理**: 作成者のみ編集・削除可能
4. **Discord連携**: チャンネルIDによる連携（実際の参加ログは自動記録）

### 参加システム
- **参加登録**: 事前の参加意思表示
- **ステータス管理**: `registered` ↔ `cancelled` での状態変更（履歴保持）
- **実際の参加**: Discord参加ログで自動記録（attendance_logsテーブル）
- **定員管理**: max_participantsによる事前登録制限

### アクセス制御
- **全ユーザー**: イベント一覧・詳細閲覧、参加申込可能
- **作成者**: イベント編集・削除権限
- **参加者**: 自分の参加ステータス変更権限