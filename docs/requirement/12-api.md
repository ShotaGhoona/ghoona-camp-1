# Ghoona Camp - API Design

## Overview
Ghoona CampアプリケーションのRESTful API設計書です。
認証にはClerkを使用し、データベースはSupabaseを使用します。

## Base URL
```
https://api.ghoona-camp.com/v1
```

## Authentication
Clerk認証トークンを使用します。
```
Authorization: Bearer <clerk_token>
```

## API Endpoints

### User Management
ユーザー認証・プロフィール管理関連のエンドポイント

| Method | Endpoint | Description | Access |
|--------|----------|-------------|---------|
| GET | `/auth/me` | 現在ログイン中のユーザー情報を取得。プロフィール設定画面の初期表示などで使用 | 🔐 |
| GET | `/users` | 全ユーザーの一覧を取得。メンバー検索・ランキング表示で使用 | 🔐 |
| GET | `/users/{userId}` | 指定したユーザーの詳細情報を取得。他のユーザーのプロフィール表示で使用 | 🔐 |
| PUT | `/users/{userId}` | ユーザーの基本情報（表示名、アバター等）を更新 | 👤 |
| GET | `/users/{userId}/metadata` | ユーザーの詳細メタデータ（スキル、興味、自己紹介等）を取得 | 🔐 |
| PUT | `/users/{userId}/metadata` | ユーザーの詳細メタデータを更新。プロフィール設定で使用 | 👤 |
| GET | `/users/{userId}/social-links` | ユーザーのSNSリンク・外部リンク一覧を取得 | 🔐 |
| POST | `/users/{userId}/social-links` | 新しいSNSリンク・外部リンクを追加 | 👤 |
| PUT | `/users/{userId}/social-links/{linkId}` | 既存のSNSリンク・外部リンクを更新 | 👤 |
| DELETE | `/users/{userId}/social-links/{linkId}` | SNSリンク・外部リンクを削除 | 👤 |
| GET | `/users/{userId}/rivals` | ユーザーが設定したライバル一覧を取得（最大3人） | 👤 |
| POST | `/users/{userId}/rivals` | 新しいライバルを追加。ダッシュボードでの比較表示で使用 | 👤 |
| DELETE | `/users/{userId}/rivals/{rivalId}` | ライバル関係を解除 | 👤 |

#### GET /users クエリパラメータ
- `page` (number): ページ番号（デフォルト: 1）
- `limit` (number): 1ページあたりの件数（デフォルト: 20）
- `search` (string): ユーザー名での検索
- `skills` (string): スキルでフィルタリング（カンマ区切り）
- `interests` (string): 興味・関心でフィルタリング（カンマ区切り）
- `sortBy` (string): ソート基準（name | attendanceDays | streakDays | createdAt）
- `order` (string): ソート順（asc | desc）

**例:** `GET /users?page=1&limit=20&search=john&skills=JavaScript,React&sortBy=attendanceDays&order=desc`

### Goal Management
目標設定・管理関連のエンドポイント

| Method | Endpoint | Description | Access |
|--------|----------|-------------|---------|
| GET | `/users/{userId}/goals` | 指定ユーザーの目標一覧を取得。公開目標のみ他ユーザーからも閲覧可能 | 👤📖 |
| POST | `/users/{userId}/goals` | 新しい目標を作成。タイトル、説明、期間、公開設定を指定 | 👤 |
| GET | `/goals/{goalId}` | 指定した目標の詳細情報を取得。進捗確認や編集画面で使用 | 👤📖 |
| PUT | `/goals/{goalId}` | 目標の内容を更新（タイトル、説明、期間、公開設定等） | 👤 |
| DELETE | `/goals/{goalId}` | 目標を削除。達成済みまたは不要になった目標の削除で使用 | 👤 |
| GET | `/goals/public` | 全ユーザーの公開目標一覧を取得。他ユーザーの目標閲覧で使用 | 🔐 |

### Event Management
朝活イベント・参加者管理関連のエンドポイント

| Method | Endpoint | Description | Access |
|--------|----------|-------------|---------|
| GET | `/events` | 開催予定・開催中のイベント一覧を取得。日時順でソート | 🔐 |
| POST | `/events` | 新しい朝活イベントを作成。タイトル、説明、日時、最大参加者数等を設定 | 🔐 |
| GET | `/events/{eventId}` | 指定イベントの詳細情報を取得。参加者一覧、説明、Discord情報等 | 🔐 |
| PUT | `/events/{eventId}` | イベント情報を更新。作成者のみが実行可能 | 👑 |
| DELETE | `/events/{eventId}` | イベントを削除。作成者のみが実行可能 | 👑 |
| GET | `/events/{eventId}/participants` | イベント参加者一覧を取得。参加状況確認で使用 | 🔐 |
| POST | `/events/{eventId}/participants` | イベントに参加申込を行う。定員チェックも実行 | 🔐 |
| PUT | `/events/{eventId}/participants/{userId}` | 参加ステータスを更新（参加→キャンセル等） | 👤 |
| DELETE | `/events/{eventId}/participants/{userId}` | イベント参加をキャンセル | 👤 |

### Title Management
称号・バッジ管理関連のエンドポイント

| Method | Endpoint | Description | Access |
|--------|----------|-------------|---------|
| GET | `/titles` | 全称号（8段階）の一覧を取得。レベル、必要日数、説明、画像等を含む | 🔐 |
| GET | `/titles/{titleId}` | 指定称号の詳細情報を取得。ストーリー、獲得条件、保持者数等 | 🔐 |
| GET | `/users/{userId}/achievements` | ユーザーの称号取得履歴・現在設定中の称号を取得 | 🔐 |
| PUT | `/users/{userId}/achievements/{titleId}` | 現在表示する称号を変更。獲得済み称号から選択 | 👤 |

### Attendance Management
Discord参加記録・統計管理関連のエンドポイント

| Method | Endpoint | Description | Access |
|--------|----------|-------------|---------|
| GET | `/users/{userId}/attendance/logs` | ユーザーのDiscord参加ログ詳細を取得。日時、参加時間、チャンネル等 | 👤 |
| POST | `/attendance/logs` | Discord Botからの参加ログ記録。入退室時刻を自動記録 | 🤖 |
| GET | `/users/{userId}/attendance/summaries` | ユーザーの日次参加サマリーを取得。カレンダー表示で使用 | 👤 |
| GET | `/users/{userId}/attendance/statistics` | ユーザーの参加統計（総日数、連続日数、称号進捗等）を取得 | 🔐 |
| GET | `/ranking/monthly` | 当月の参加日数ランキングを取得。ライバル表示も含む | 🔐 |
| GET | `/ranking/total` | 総合参加日数ランキングを取得。歴代順位表示 | 🔐 |
| GET | `/ranking/streak` | 連続参加日数ランキングを取得。現在の連続記録で順位付け | 🔐 |

### Notification Management
通知・リマインダー管理関連のエンドポイント

| Method | Endpoint | Description | Access |
|--------|----------|-------------|---------|
| GET | `/users/{userId}/notifications` | ユーザーの通知一覧を取得。称号獲得、ライバル更新、イベント等 | 👤 |
| PUT | `/notifications/{notificationId}` | 通知の既読ステータスを更新 | 👤 |
| DELETE | `/notifications/{notificationId}` | 不要な通知を削除 | 👤 |
| GET | `/users/{userId}/notification-settings` | ユーザーの通知設定を取得。各種通知のON/OFF、時刻設定等 | 👤 |
| PUT | `/users/{userId}/notification-settings` | 通知設定を更新。リマインダー時刻、通知種別の有効/無効等 | 👤 |

### System API
システム管理・監視関連のエンドポイント

| Method | Endpoint | Description | Access |
|--------|----------|-------------|---------|
| GET | `/health` | アプリケーションのヘルスチェック。サーバー状態、DB接続確認 | 🌍 |
| GET | `/version` | 現在のAPIバージョン情報を取得 | 🌍 |

## Discord Integration
Discord Bot連携用の専用エンドポイント

| Method | Endpoint | Description | Access |
|--------|----------|-------------|---------|
| POST | `/discord/attendance` | Discord参加ログを記録。出席チャンネル（6:00-6:30参加開始）のみ有効 | 🤖 |
| POST | `/discord/webhook` | Discordイベント（入退室等）のウェブフック受信 | 🤖 |
| GET | `/discord/users/{discordId}` | Discord IDから内部User IDへの変換。連携確認で使用 | 🤖 |

### Discord Bot認証
Discord BotからのAPIアクセスは**APIキー方式**で認証します。

```http
X-API-Key: your-bot-api-key
```

### 参加判定ロジック
- **対象チャンネル**: 専用の「出席チャンネル」のみ
- **有効参加時間**: 6:00-6:30の間に参加開始したユーザー
- **最低参加時間**: 制限なし（短時間でも有効）
- **退室時刻**: 6:30以前の退室でも有効参加とみなす

## Response Format

### Success Response
```json
{
  "data": {...},
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### Error Response
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid request parameters",
    "details": {...}
  },
  "timestamp": "2025-01-21T10:00:00Z"
}
```

## HTTP Status Codes

| Code | Description |
|------|-------------|
| 200 | OK - リクエスト成功 |
| 201 | Created - リソース作成成功 |
| 204 | No Content - 削除成功 |
| 400 | Bad Request - リクエストエラー |
| 401 | Unauthorized - 認証エラー |
| 403 | Forbidden | アクセス権限エラー |
| 404 | Not Found - リソースが見つからない |
| 409 | Conflict - データ競合 |
| 422 | Unprocessable Entity - バリデーションエラー |
| 500 | Internal Server Error - サーバーエラー |

## Access Control

### アクセス権限の説明
- **🌍 公開**: 認証不要
- **🔐 認証済み**: Clerk認証が必要
- **👤 本人のみ**: 自分のリソースのみアクセス可能
- **👑 作成者のみ**: リソースの作成者のみアクセス可能
- **🤖 Bot認証**: Discord Bot専用の認証
- **📖 公開設定**: is_publicフラグで制御
- **👤📖 本人または公開**: 本人のリソースか、is_publicがtrueの場合

### プライバシー制御
- ユーザーメタデータは本人のみ編集可能
- 目標はis_publicフラグで公開制御
- ビジョンはvision_publicフラグで公開制御
- 参加統計は本人のみ詳細閲覧可能