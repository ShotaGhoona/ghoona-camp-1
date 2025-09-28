# Title Management API - 詳細設計

## Overview
Ghoona Campアプリケーションの称号・バッジ管理関連API詳細設計書です。
8段階の称号システムとユーザーの称号獲得・表示管理を中心とした機能を提供します。

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
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440001",
      "level": 1,
      "name_jp": "まどろみ見習い",
      "name_en": "Sleeper",
      "description": "朝活の世界への第一歩。1日の参加で獲得できる称号。",
      "required_days": 1,
      "image_url": "https://cdn.ghoona-camp.com/titles/level-1.png",
      "color_theme": "#B0BEC5",
      "is_active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    {
      "id": "550e8400-e29b-41d4-a716-446655440002",
      "level": 2,
      "name_jp": "早起き戦士",
      "name_en": "Early Riser",
      "description": "5日間の継続により、朝の時間を味方につけ始めた戦士。",
      "required_days": 5,
      "image_url": "https://cdn.ghoona-camp.com/titles/level-2.png",
      "color_theme": "#4ECDC4",
      "is_active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
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
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "level": 1,
    "name_jp": "まどろみ見習い",
    "name_en": "Sleeper",
    "description": "朝活の世界への第一歩。1日の参加で獲得できる称号。",
    "required_days": 1,
    "image_url": "https://cdn.ghoona-camp.com/titles/level-1.png",
    "color_theme": "#B0BEC5",
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
    "achievements": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440100",
        "title": {
          "id": "550e8400-e29b-41d4-a716-446655440001",
          "level": 1,
          "name_jp": "まどろみ見習い",
          "name_en": "Sleeper",
          "description": "朝活の世界への第一歩。1日の参加で獲得できる称号。",
          "required_days": 1,
          "image_url": "https://cdn.ghoona-camp.com/titles/level-1.png",
          "color_theme": "#B0BEC5",
          "is_active": true,
          "created_at": "2024-01-01T00:00:00Z",
          "updated_at": "2024-01-01T00:00:00Z"
        },
        "achieved_at": "2024-10-15T06:30:00Z",
        "is_current": true,
        "created_at": "2024-10-15T06:30:00Z",
        "updated_at": "2024-10-15T06:30:00Z"
      }
    ]
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
      "id": "550e8400-e29b-41d4-a716-446655440001",
      "level": 1,
      "name_jp": "まどろみ見習い",
      "name_en": "Sleeper",
      "description": "朝活の世界への第一歩。1日の参加で獲得できる称号。",
      "required_days": 1,
      "image_url": "https://cdn.ghoona-camp.com/titles/level-1.png",
      "color_theme": "#B0BEC5",
      "is_active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    "achieved_at": "2024-10-15T06:30:00Z",
    "is_current": true,
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Current title updated successfully",
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

## 称号システムのルール

### 称号レベル
- レベル1-8の8段階システム
- 各レベルには固定の必要参加日数を設定

### 称号名
- レベル1: まどろみ見習い (Sleeper) - 1日
- レベル2: 早起き戦士 (Early Riser) - 5日
- レベル3: 陽光探求者 (Dawn Seeker) - 15日  
- レベル4: 朝陽の使者 (Sun Messenger) - 30日
- レベル5: 暁の守護者 (Dawn Guardian) - 60日
- レベル6: 朝活の覇者 (Morning Master) - 100日
- レベル7: 夜明けの皇帝 (Dawn Emperor) - 200日
- レベル8: 永遠の夜明け (Eternal Dawn) - 365日

### 基本ルール
1. **参加日数ベース**: 朝活参加日数に基づいて自動獲得
2. **表示称号**: 獲得済み称号から1つのみ選択可能
3. **非可逆性**: 一度獲得した称号は削除不可