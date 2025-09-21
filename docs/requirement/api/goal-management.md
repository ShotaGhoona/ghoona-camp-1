# Goal Management API - 詳細設計

## Overview
Ghoona Campアプリケーションの目標管理関連API詳細設計書です。
ユーザーの朝活目標設定・進捗管理・公開設定を中心とした機能を提供します。

## アーキテクチャ適合性
このAPI設計は以下のアーキテクチャ原則に従います：
- **オニオンアーキテクチャ**: Domain → Application → Infrastructure → Interface層の依存関係
- **ドメイン駆動設計**: `goal`コンテキストを中心としたエンティティ設計
- **クリーンな境界**: DTO、ドメインエンティティ、GORMモデルの適切な分離

## ドメイン設計対応

### ドメインエンティティ
```go
// internal/domain/goal/entity/goal.go
type Goal struct {
    ID          UUID
    UserID      UUID  
    Title       string
    Description string
    StartedAt   time.Time
    EndedAt     *time.Time
    IsActive    bool
    IsPublic    bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// internal/domain/goal/entity/goal_progress.go  
type GoalProgress struct {
    ID              UUID
    GoalID          UUID
    Date            time.Time
    Status          ProgressStatus
    Note            string
    DurationMinutes int
    QualityRating   *int
}

// internal/domain/goal/value/progress_status.go
type ProgressStatus string
const (
    ProgressCompleted ProgressStatus = "completed"
    ProgressPartial   ProgressStatus = "partial" 
    ProgressSkipped   ProgressStatus = "skipped"
)
```

### リポジトリインターフェース
```go
// internal/domain/goal/repository/goal_repository.go
type GoalRepository interface {
    FindByUserID(ctx context.Context, userID UUID) ([]*entity.Goal, error)
    FindByID(ctx context.Context, id UUID) (*entity.Goal, error)
    Create(ctx context.Context, goal *entity.Goal) error
    Update(ctx context.Context, goal *entity.Goal) error
    Delete(ctx context.Context, id UUID) error
    FindPublicGoals(ctx context.Context, filters GoalFilters) ([]*entity.Goal, error)
}

type GoalProgressRepository interface {
    FindByGoalID(ctx context.Context, goalID UUID) ([]*entity.GoalProgress, error)
    CreateProgress(ctx context.Context, progress *entity.GoalProgress) error
    UpdateProgress(ctx context.Context, progress *entity.GoalProgress) error
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

## 1. 個人目標管理 API

### ユーザーの目標一覧取得

```
GET /users/{userId}/goals
```

**権限**: 👤📖 本人または公開目標

**説明**: 指定ユーザーの目標一覧を取得。公開目標のみ他ユーザーからも閲覧可能

**パラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| userId | UUID | Yes | ユーザーID |

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| status | string | No | フィルタ (active, completed, archived) |
| page | integer | No | ページ番号 (デフォルト: 1) |
| limit | integer | No | 1ページあたりの件数 (デフォルト: 10, 最大: 50) |
| sort_by | string | No | ソート基準 (created_at, started_at, ended_at) |
| order | string | No | ソート順 (asc, desc) デフォルト: desc |

**レスポンス:**

```json
{
  "data": {
    "goals": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440100",
        "user_id": "550e8400-e29b-41d4-a716-446655440000",
        "title": "毎日読書30分",
        "description": "朝活時間を使って技術書や自己啓発書を読む習慣をつける",
        "started_at": "2025-01-01",
        "ended_at": "2025-03-31",
        "is_active": true,
        "is_public": true,
        "progress": {
          "total_days": 90,
          "completed_days": 15,
          "completion_rate": 16.7,
          "current_streak": 3,
          "remaining_days": 75
        },
        "created_at": "2024-12-15T10:30:00Z",
        "updated_at": "2025-01-15T08:45:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440101",
        "user_id": "550e8400-e29b-41d4-a716-446655440000",
        "title": "筋トレ週4回",
        "description": "健康的な体作りのために筋トレを継続する",
        "started_at": "2025-01-01",
        "ended_at": "2025-06-30",
        "is_active": true,
        "is_public": false,
        "progress": {
          "total_days": 181,
          "completed_days": 8,
          "completion_rate": 4.4,
          "current_streak": 2,
          "remaining_days": 173
        },
        "created_at": "2024-12-20T15:20:00Z",
        "updated_at": "2025-01-14T19:30:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 1,
      "total_count": 2,
      "limit": 10,
      "has_next": false,
      "has_prev": false
    },
    "summary": {
      "total_goals": 2,
      "active_goals": 2,
      "completed_goals": 0,
      "public_goals": 1
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 新規目標作成

```
POST /users/{userId}/goals
```

**権限**: 👤 本人のみ

**説明**: 新しい目標を作成。タイトル、説明、期間、公開設定を指定

**リクエスト:**

```json
{
  "title": "英語学習1時間",
  "description": "TOEIC900点を目指して毎朝1時間英語学習を続ける。単語暗記30分、リスニング30分の構成で進める。",
  "started_at": "2025-02-01",
  "ended_at": "2025-08-31",
  "is_public": true
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| title | 必須、最大200文字 |
| description | 最大1000文字 |
| started_at | 必須、YYYY-MM-DD形式、今日以降 |
| ended_at | YYYY-MM-DD形式、started_at以降 |
| is_public | boolean、デフォルト: false |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440102",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "英語学習1時間",
    "description": "TOEIC900点を目指して毎朝1時間英語学習を続ける。単語暗記30分、リスニング30分の構成で進める。",
    "started_at": "2025-02-01",
    "ended_at": "2025-08-31",
    "is_active": true,
    "is_public": true,
    "progress": {
      "total_days": 212,
      "completed_days": 0,
      "completion_rate": 0.0,
      "current_streak": 0,
      "remaining_days": 212
    },
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

**権限**: 👤📖 本人または公開目標

**説明**: 指定した目標の詳細情報を取得。進捗確認や編集画面で使用

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
      "avatar_url": "https://img.clerk.com/preview.png",
      "current_title": {
        "level": 3,
        "name_jp": "陽光探求者",
        "color_theme": "#FFB800"
      }
    },
    "title": "毎日読書30分",
    "description": "朝活時間を使って技術書や自己啓発書を読む習慣をつける",
    "started_at": "2025-01-01",
    "ended_at": "2025-03-31",
    "is_active": true,
    "is_public": true,
    "progress": {
      "total_days": 90,
      "completed_days": 15,
      "completion_rate": 16.7,
      "current_streak": 3,
      "max_streak": 5,
      "remaining_days": 75,
      "days_since_start": 15,
      "expected_completion_rate": 16.7,
      "on_track": true
    },
    "milestones": [
      {
        "date": "2025-01-15",
        "days_completed": 15,
        "milestone_type": "streak_achieved",
        "description": "3日連続達成"
      }
    ],
    "recent_activity": [
      {
        "date": "2025-01-15",
        "status": "completed",
        "note": "Vue.js公式ドキュメント読了"
      },
      {
        "date": "2025-01-14", 
        "status": "completed",
        "note": "TypeScript実践入門 第3章"
      },
      {
        "date": "2025-01-13",
        "status": "skipped",
        "note": null
      }
    ],
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

**説明**: 目標の内容を更新（タイトル、説明、期間、公開設定等）

**リクエスト:**

```json
{
  "title": "毎日読書45分",
  "description": "朝活時間を使って技術書や自己啓発書を読む習慣をつける。目標時間を30分から45分に延長。",
  "ended_at": "2025-04-30",
  "is_public": true,
  "is_active": true
}
```

**バリデーション:**
- `started_at`は変更不可（目標開始後は変更できない）
- `ended_at`は現在日以降である必要がある
- アクティブな目標のみ更新可能

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440100",
    "title": "毎日読書45分",
    "description": "朝活時間を使って技術書や自己啓発書を読む習慣をつける。目標時間を30分から45分に延長。",
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

**説明**: 目標を削除。達成済みまたは不要になった目標の削除で使用

**注意事項:**
- 論理削除（`is_active = false`）
- 進捗データは保持される
- 削除後は復元不可（アーカイブ扱い）

**レスポンス:**

```json
{
  "message": "Goal deleted successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 2. 公開目標閲覧 API

### 公開目標一覧取得

```
GET /goals/public
```

**権限**: 🔐 認証済み

**説明**: 全ユーザーの公開目標一覧を取得。他ユーザーの目標閲覧で使用

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| category | string | No | カテゴリフィルタ (study, exercise, creative, lifestyle, etc.) |
| status | string | No | ステータスフィルタ (active, completed) |
| search | string | No | タイトル・説明での部分検索 |
| user_id | UUID | No | 特定ユーザーの公開目標のみ |
| page | integer | No | ページ番号 (デフォルト: 1) |
| limit | integer | No | 1ページあたりの件数 (デフォルト: 20, 最大: 50) |
| sort_by | string | No | ソート基準 (created_at, started_at, popularity) |
| order | string | No | ソート順 (asc, desc) デフォルト: desc |

**レスポンス:**

```json
{
  "data": {
    "goals": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440100",
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
        "title": "毎日読書30分",
        "description": "朝活時間を使って技術書や自己啓発書を読む習慣をつける",
        "started_at": "2025-01-01",
        "ended_at": "2025-03-31",
        "progress": {
          "completion_rate": 16.7,
          "current_streak": 3,
          "on_track": true
        },
        "category": "study",
        "likes_count": 12,
        "comments_count": 3,
        "is_liked": false,
        "created_at": "2024-12-15T10:30:00Z"
      },
      {
        "id": "550e8400-e29b-41d4-a716-446655440105",
        "user": {
          "id": "550e8400-e29b-41d4-a716-446655440002",
          "display_name": "佐藤花子",
          "username": "sato_hanako",
          "avatar_url": "https://img.clerk.com/sato.png",
          "current_title": {
            "level": 4,
            "name_jp": "朝陽の使者",
            "color_theme": "#FF6B35"
          }
        },
        "title": "ヨガ週5回",
        "description": "心身の健康のために朝ヨガを継続する",
        "started_at": "2024-12-01",
        "ended_at": "2025-06-01",
        "progress": {
          "completion_rate": 85.2,
          "current_streak": 12,
          "on_track": true
        },
        "category": "exercise",
        "likes_count": 25,
        "comments_count": 8,
        "is_liked": true,
        "created_at": "2024-11-25T14:20:00Z"
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
    "filters": {
      "available_categories": ["study", "exercise", "creative", "lifestyle", "business"],
      "total_active": 35,
      "total_completed": 12
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 3. 目標進捗管理 API

### 目標進捗記録

```
POST /goals/{goalId}/progress
```

**権限**: 👤 本人のみ

**説明**: 目標の日次進捗を記録。朝活参加と連携して自動/手動で記録

**リクエスト:**

```json
{
  "date": "2025-01-21",
  "status": "completed",
  "note": "TypeScript実践入門 第4章完了。async/awaitの理解が深まった",
  "duration_minutes": 45,
  "quality_rating": 4
}
```

**バリデーション:**

| フィールド | 制約 |
|-----------|------|
| date | 必須、YYYY-MM-DD形式、過去30日以内 |
| status | 必須、enum(completed, partial, skipped) |
| note | 最大500文字 |
| duration_minutes | 0以上の整数 |
| quality_rating | 1-5の整数（満足度） |

**レスポンス:**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440200",
    "goal_id": "550e8400-e29b-41d4-a716-446655440100",
    "date": "2025-01-21",
    "status": "completed",
    "note": "TypeScript実践入門 第4章完了。async/awaitの理解が深まった",
    "duration_minutes": 45,
    "quality_rating": 4,
    "streak_updated": {
      "previous_streak": 3,
      "current_streak": 4,
      "is_new_record": false
    },
    "created_at": "2025-01-21T10:00:00Z"
  },
  "message": "Progress recorded successfully",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 目標進捗履歴取得

```
GET /goals/{goalId}/progress
```

**権限**: 👤📖 本人または公開目標

**説明**: 目標の進捗履歴を取得。カレンダー表示やグラフ表示で使用

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| from_date | string | No | 開始日 (YYYY-MM-DD) |
| to_date | string | No | 終了日 (YYYY-MM-DD) |
| status | string | No | ステータスフィルタ (completed, partial, skipped) |
| limit | integer | No | 取得件数 (デフォルト: 30, 最大: 100) |

**レスポンス:**

```json
{
  "data": {
    "goal": {
      "id": "550e8400-e29b-41d4-a716-446655440100",
      "title": "毎日読書30分",
      "started_at": "2025-01-01",
      "ended_at": "2025-03-31"
    },
    "progress_entries": [
      {
        "date": "2025-01-21",
        "status": "completed",
        "note": "TypeScript実践入門 第4章完了",
        "duration_minutes": 45,
        "quality_rating": 4
      },
      {
        "date": "2025-01-20",
        "status": "completed",
        "note": "Vue.js公式ドキュメント読了",
        "duration_minutes": 30,
        "quality_rating": 5
      },
      {
        "date": "2025-01-19",
        "status": "skipped",
        "note": "体調不良のため休み",
        "duration_minutes": 0,
        "quality_rating": null
      }
    ],
    "statistics": {
      "total_days": 21,
      "completed_days": 15,
      "partial_days": 2,
      "skipped_days": 4,
      "completion_rate": 71.4,
      "average_duration": 38.2,
      "average_quality": 4.1,
      "current_streak": 4,
      "max_streak": 7
    }
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

---

## 4. 目標統計・分析 API

### 個人目標統計

```
GET /users/{userId}/goals/statistics
```

**権限**: 👤 本人のみ

**説明**: ユーザーの目標達成統計を取得。ダッシュボード表示で使用

**クエリパラメータ:**

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| period | string | No | 期間 (week, month, quarter, year, all) デフォルト: month |

**レスポンス:**

```json
{
  "data": {
    "period": "month",
    "period_range": {
      "from": "2025-01-01",
      "to": "2025-01-31"
    },
    "goals_summary": {
      "total_goals": 3,
      "active_goals": 2,
      "completed_goals": 0,
      "paused_goals": 1,
      "average_completion_rate": 45.6
    },
    "progress_summary": {
      "total_progress_days": 42,
      "completed_days": 28,
      "partial_days": 6,
      "skipped_days": 8,
      "overall_completion_rate": 66.7,
      "longest_streak": 12,
      "current_active_streaks": 2
    },
    "categories": [
      {
        "category": "study",
        "goals_count": 2,
        "completion_rate": 72.5,
        "total_duration_minutes": 1260
      },
      {
        "category": "exercise", 
        "goals_count": 1,
        "completion_rate": 45.2,
        "total_duration_minutes": 840
      }
    ],
    "achievements": [
      {
        "type": "streak_milestone",
        "description": "10日連続達成",
        "achieved_at": "2025-01-15T06:30:00Z",
        "goal_title": "毎日読書30分"
      },
      {
        "type": "duration_milestone",
        "description": "累計20時間達成",
        "achieved_at": "2025-01-18T07:15:00Z",
        "goal_title": "筋トレ週4回"
      }
    ],
    "trends": {
      "completion_rate_trend": "increasing",
      "average_quality_trend": "stable",
      "consistency_score": 78.5
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
| GOAL_NOT_FOUND | 目標が見つからない | 404 |
| GOAL_NOT_ACCESSIBLE | 目標へのアクセス権限なし | 403 |
| GOAL_VALIDATION_ERROR | 目標データが無効 | 422 |
| PROGRESS_ALREADY_EXISTS | 指定日の進捗が既に存在 | 409 |
| PROGRESS_DATE_INVALID | 進捗記録日が無効 | 422 |
| GOAL_LIMIT_EXCEEDED | 目標数の上限超過 | 422 |
| GOAL_PERIOD_INVALID | 目標期間が無効 | 422 |

## ビジネスルール（ドメイン層実装）

### ドメインルール
```go
// internal/domain/goal/errors.go
var (
    ErrGoalLimitExceeded    = errors.New("目標数の上限（10個）を超えています")
    ErrGoalPeriodTooLong    = errors.New("目標期間は最大1年間です")
    ErrProgressDateInvalid  = errors.New("進捗記録は過去30日以内のみ可能です")
    ErrGoalNotPublic        = errors.New("非公開の目標です")
    ErrProgressAlreadyExists = errors.New("指定日の進捗が既に存在します")
)

// internal/domain/goal/entity/goal.go
func (g *Goal) Validate() error {
    if g.Title == "" {
        return errors.New("目標タイトルは必須です")
    }
    if g.EndedAt != nil && g.EndedAt.Sub(g.StartedAt) > 365*24*time.Hour {
        return ErrGoalPeriodTooLong
    }
    return nil
}

func (g *Goal) CanBeAccessedBy(userID UUID) bool {
    return g.UserID == userID || g.IsPublic
}

// internal/domain/goal/entity/goal_progress.go
func (gp *GoalProgress) Validate() error {
    if gp.Date.After(time.Now()) {
        return errors.New("未来日の進捗記録はできません")
    }
    if time.Since(gp.Date) > 30*24*time.Hour {
        return ErrProgressDateInvalid
    }
    if gp.QualityRating != nil && (*gp.QualityRating < 1 || *gp.QualityRating > 5) {
        return errors.New("品質評価は1-5の範囲で入力してください")
    }
    return nil
}
```

### 目標作成・管理のルール
1. **目標数制限**: ユーザー1人につき同時進行可能な目標は最大10個（ドメインサービスで制御）
2. **期間制限**: 目標期間は最大1年間（Goal.Validate()で制御）
3. **進捗記録**: 過去30日以内の日付のみ記録可能（GoalProgress.Validate()で制御）
4. **公開設定**: 一度公開した目標は非公開に変更不可（ユースケース層で制御）
5. **削除制限**: 進捗が記録されている目標は論理削除のみ（ドメインサービスで制御）

### 進捗記録のルール
1. **重複防止**: 同一日の進捗は1回のみ記録可能（リポジトリ層で制御）
2. **遡及記録**: 過去分の記録は可能だが、未来日は不可（GoalProgress.Validate()で制御）
3. **ストリーク計算**: 連続日数は日次バッチで自動計算
4. **品質評価**: 1-5段階での満足度評価（GoalProgress.Validate()で制御）

## 注意事項

### プライバシー考慮
- 公開設定の目標のみ他ユーザーから閲覧可能
- 進捗の詳細（note、quality_rating）は本人のみ閲覧可能
- 統計情報は本人のみアクセス可能

### パフォーマンス考慮
- 進捗履歴は期間指定を推奨
- 公開目標一覧はページネーション必須
- 統計計算は日次バッチで事前計算

### キャッシュ戦略
- 公開目標一覧: 10分キャッシュ
- 目標統計: 1時間キャッシュ
- 進捗履歴: 5分キャッシュ