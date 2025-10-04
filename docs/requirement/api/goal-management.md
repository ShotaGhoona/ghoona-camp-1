# Goal Management API - 詳細設計

## Overview
Ghoona Campアプリケーションの目標管理関連API詳細設計書です。
朝活セッション期間中に取り組む個人目標の簡単な設定・共有機能を提供します。

## Base URL
```
https://api.ghoona-camp.com/v1
```

## Authentication
```
Authorization: Bearer <clerk_token>
```

---

## 1. 目標管理 API

### ユーザーの目標一覧取得

```
GET /users/{userId}/goals
```

**権限**: 🔐 認証済み（本人の場合は全て、他人の場合は公開目標のみ）

**説明**: 指定ユーザーの朝活目標一覧を取得

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| userId | UUID | Yes | ユーザーID |

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| is_active | boolean | No | アクティブな目標のみ (デフォルト: true) |

**レスポンス:**

```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440100",
      "title": "朝活で読書習慣",
      "description": "朝活時間を使って技術書を読む",
      "started_at": "2025-01-01",
      "ended_at": "2025-03-31",
      "is_active": true,
      "is_public": true,
      "created_at": "2024-12-15T10:30:00Z",
      "updated_at": "2025-01-15T08:45:00Z"
    }
  ],
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 新規目標作成

```
POST /users/{userId}/goals
```

**権限**: 👤 本人のみ

**説明**: 新しい朝活目標を作成

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| userId | UUID | Yes | ユーザーID |

**リクエスト:**

```json
{
  "title": "朝活で英語学習",
  "description": "朝活時間で英語のリスニング練習",
  "started_at": "2025-02-01",
  "ended_at": "2025-08-31",
  "is_public": true
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| title | 必須、1-200文字 |
| description | 最大1000文字 |
| started_at | 必須、YYYY-MM-DD形式 |
| ended_at | YYYY-MM-DD形式、started_at以降 |
| is_public | boolean、デフォルト: false |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440102",
    "title": "朝活で英語学習",
    "description": "朝活時間で英語のリスニング練習",
    "started_at": "2025-02-01",
    "ended_at": "2025-08-31",
    "is_active": true,
    "is_public": true,
    "created_at": "2025-01-21T10:00:00Z",
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Goal created successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 目標詳細取得

```
GET /goals/{goalId}
```

**権限**: 🔐 認証済み（本人の場合は全て、他人の場合は公開目標のみ）

**説明**: 指定した目標の詳細情報を取得

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| goalId | UUID | Yes | 目標ID |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440100",
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "display_name": "田中太郎",
      "username": "tanaka_taro",
      "avatar_url": "https://img.clerk.com/preview.png"
    },
    "title": "朝活で読書習慣",
    "description": "朝活時間を使って技術書を読む",
    "started_at": "2025-01-01",
    "ended_at": "2025-03-31",
    "is_active": true,
    "is_public": true,
    "created_at": "2024-12-15T10:30:00Z",
    "updated_at": "2025-01-15T08:45:00Z"
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 目標更新

```
PUT /goals/{goalId}
```

**権限**: 👤 本人のみ

**説明**: 目標の内容を更新

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| goalId | UUID | Yes | 目標ID |

**リクエスト:**

```json
{
  "title": "朝活で英語とプログラミング",
  "description": "朝活時間で英語学習とプログラミング練習",
  "ended_at": "2025-04-30",
  "is_public": true,
  "is_active": true
}
```

**バリデーション:**
- `started_at`は変更不可
- `ended_at`は変更可能（現在日以降）
- アクティブな目標のみ更新可能

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440100",
    "title": "朝活で英語とプログラミング",
    "description": "朝活時間で英語学習とプログラミング練習",
    "ended_at": "2025-04-30",
    "is_public": true,
    "is_active": true,
    "updated_at": "2025-01-21T10:00:00Z"
  },
  "message": "Goal updated successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 目標削除

```
DELETE /goals/{goalId}
```

**権限**: 👤 本人のみ

**説明**: 目標を削除（論理削除）

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| goalId | UUID | Yes | 目標ID |

**レスポンス:**

```json
{
  "message": "Goal deleted successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 2. 公開目標一覧 API

### 公開目標一覧取得

```
GET /goals/public
```

**権限**: 🔐 認証済み

**説明**: 全ユーザーの公開目標一覧を取得

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| search | string | No | タイトル・説明での部分検索 |
| page | integer | No | ページ番号 (デフォルト: 1) |
| limit | integer | No | 1ページあたりの件数 (デフォルト: 20, 最大: 50) |

**レスポンス:**

```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440100",
      "user": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "display_name": "田中太郎",
        "username": "tanaka_taro",
        "avatar_url": "https://img.clerk.com/preview.png"
      },
      "title": "朝活で読書習慣",
      "description": "朝活時間を使って技術書を読む",
      "started_at": "2025-01-01",
      "ended_at": "2025-03-31",
      "created_at": "2024-12-15T10:30:00Z"
    }
  ],
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

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
| GOAL_NOT_FOUND | 目標が見つからない | 404 |
| GOAL_NOT_ACCESSIBLE | 目標へのアクセス権限なし | 403 |
| GOAL_VALIDATION_ERROR | 目標データが無効 | 422 |
| GOAL_PERIOD_INVALID | 目標期間が無効 | 422 |

## 基本ルール

### 目標管理ルール
1. **目標作成**: シンプルなタイトル・説明・期間設定
2. **公開設定**: 個人目標をコミュニティで共有可能
3. **削除**: 論理削除（`is_active = false`）
4. **権限**: 本人のみ作成・更新・削除可能

### アクセス制御
- **本人**: 全ての目標（公開・非公開）を閲覧・編集可能
- **他ユーザー**: 公開目標のみ閲覧可能
- **未認証**: アクセス不可