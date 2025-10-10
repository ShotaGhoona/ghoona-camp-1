# BE-0401: Application DTO Strategy

## Overview
Ghoona Camp BackendのApplication層におけるDTO（Data Transfer Object）の設計戦略書です。
DDD × Onion Architectureに従い、**APIエンドポイントと1:1対応**するtype-safeで保守性の高いDTO構造を定義します。

## Design Philosophy: API Endpoint 1:1 Mapping

各DTOファイルが特定のAPIエンドポイントと完全に1:1対応することで、以下のメリットを実現：
- **可読性向上**: どのファイルがどのAPIに対応するか一目瞭然
- **保守性向上**: APIの変更時に影響範囲が明確
- **開発効率向上**: 新しいエンドポイント追加時の作業が定型化

## Directory Structure

```
internal/application/dto/
├── common/
│   ├── pagination.go        # ページネーション共通DTO
│   ├── response.go          # 標準レスポンス形式
│   └── error.go            # エラーレスポンス形式
├── users/
│   ├── request/
│   │   ├── put_user.go                    # PUT /users/{userId}
│   │   ├── post_social_link.go            # POST /users/{userId}/social-links
│   │   ├── put_social_link.go             # PUT /users/{userId}/social-links/{linkId}
│   │   └── post_rival.go                  # POST /users/{userId}/rivals
│   └── response/
│       ├── get_auth_me.go                 # GET /auth/me
│       ├── get_users.go                   # GET /users
│       ├── get_user_detail.go             # GET /users/{userId}
│       ├── put_user.go                    # PUT /users/{userId}
│       ├── post_social_link.go            # POST /users/{userId}/social-links
│       ├── put_social_link.go             # PUT /users/{userId}/social-links/{linkId}
│       ├── get_user_rivals.go             # GET /users/{userId}/rivals
│       └── post_rival.go                  # POST /users/{userId}/rivals
├── goals/
│   ├── request/
│   │   ├── post_goal.go                   # POST /goals
│   │   └── put_goal.go                    # PUT /goals/{goalId}
│   └── response/
│       ├── get_goals_me.go                # GET /goals/me
│       ├── get_goals_public.go            # GET /goals/public
│       ├── post_goal.go                   # POST /goals
│       └── put_goal.go                    # PUT /goals/{goalId}
├── events/
│   ├── request/
│   │   ├── post_event.go                  # POST /events
│   │   ├── put_event.go                   # PUT /events/{eventId}
│   │   └── put_event_participant.go       # PUT /events/{eventId}/participants/{userId}
│   └── response/
│       ├── get_events.go                  # GET /events
│       ├── post_event.go                  # POST /events
│       ├── get_event_detail.go            # GET /events/{eventId}
│       ├── put_event.go                   # PUT /events/{eventId}
│       ├── post_event_participant.go      # POST /events/{eventId}/participants
│       └── put_event_participant.go       # PUT /events/{eventId}/participants/{userId}
├── titles/
│   └── response/
│       ├── get_titles.go                  # GET /titles
│       ├── get_title_detail.go            # GET /titles/{titleId}
│       ├── get_user_achievements.go       # GET /users/{userId}/achievements
│       └── put_user_achievement.go        # PUT /users/{userId}/achievements/{titleId}
└── attendance/
    └── response/
        ├── get_attendance_summaries.go    # GET /users/{userId}/attendance/summaries
        ├── get_attendance_statistics.go   # GET /users/{userId}/attendance/statistics
        ├── get_ranking_monthly.go         # GET /ranking/monthly
        ├── get_ranking_total.go           # GET /ranking/total
        └── get_ranking_streak.go          # GET /ranking/streak
```

## API Endpoint Analysis

### 実装対象エンドポイント数（Notification除く）

| ドメイン | Request DTO | Response DTO | 詳細 |
|---------|------------|-------------|------|
| **Users** | 4個 | 8個 | PUT /users/{userId}, POST /users/{userId}/social-links, PUT /users/{userId}/social-links/{linkId}, POST /users/{userId}/rivals |
| **Goals** | 2個 | 4個 | POST /goals, PUT /goals/{goalId} |
| **Events** | 3個 | 6個 | POST /events, PUT /events/{eventId}, PUT /events/{eventId}/participants/{userId} |
| **Titles** | 0個 | 4個 | Read-only + URL parameter update |
| **Attendance** | 0個 | 5個 | Read-only APIs |
| **System** | 0個 | 1個 | GET /health |

**合計: 32個のエンドポイント**
- **Request DTO**: 9個（Body付きリクエスト）
- **Response DTO**: 28個（DELETE除く全レスポンス）
- **DELETE**: 4個（204 No Content レスポンス）

### 1:1マッピングの利点

1. **明確な責務分離**: 各ファイルが単一のAPIエンドポイントに特化
2. **変更影響の局所化**: API仕様変更時の影響範囲が明確
3. **開発者体験向上**: ファイル名から対応APIが即座に判別可能
4. **テスト容易性**: エンドポイント毎の単体テストが書きやすい

## Common Package Design

可読性を優先し、本当に共通で使用される最低限の要素のみを`common/`に配置：

### 1. `pagination.go`
```go
type Pagination struct {
    Total   int  `json:"total"`
    Limit   int  `json:"limit"`
    Offset  int  `json:"offset"`
    HasMore bool `json:"has_more"`
}
```

### 2. `response.go`
```go
type Response[T any] struct {
    Data      T      `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type ListResponse[T any] struct {
    Data struct {
        Items      []T         `json:"items"`
        Pagination *Pagination `json:"pagination,omitempty"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}
```

### 3. `error.go`
```go
type ErrorResponse struct {
    Error struct {
        Code    string `json:"code"`
        Message string `json:"message"`
    } `json:"error"`
    Timestamp string `json:"timestamp"`
}
```

## Design Principles

1. **ドメイン境界の明確化**: 各APIドメインでパッケージを分離
2. **Request/Response分離**: 入力と出力の責務を明確に分離
3. **最小限の共通化**: 本当に必要な共通要素のみをcommonに配置
4. **型安全性**: Goの型システムを最大限活用
5. **JSON互換性**: API仕様書のJSONレスポンスと完全一致
6. **拡張性**: 将来的な機能追加に対応可能な構造

## Implementation Guidelines

### Naming Convention

#### File Naming
- Request DTO: `{http_method}_{resource}.go` (例: `post_goal.go`, `put_user.go`)
- Response DTO: `{http_method}_{resource}.go` (例: `get_goals_me.go`, `get_user_detail.go`)

#### Struct Naming  
- Request DTO: `{HttpMethod}{Resource}Request` (例: `PostGoalRequest`, `PutUserRequest`)
- Response DTO: `{HttpMethod}{Resource}Response` (例: `GetGoalsMeResponse`, `GetUserDetailResponse`)

#### Examples
```go
// ファイル: post_goal.go
type PostGoalRequest struct { ... }

// ファイル: get_goals_me.go  
type GetGoalsMeResponse struct { ... }

// ファイル: put_user.go
type PutUserRequest struct { ... }
type PutUserResponse struct { ... }
```

### Validation
- Request DTOにはvalidationタグを付与
- Required/Optional フィールドを明確に区別
- ビジネスルールに基づいた制約を設定

### Mapping Strategy
- Domain Entity ↔ DTO の変換は明示的なマッピング関数で実装
- Auto-mapping ライブラリは使用せず、手動変換で型安全性を確保

## Next Steps

### Phase 1: Foundation
1. ✅ `common/` パッケージの実装完了
2. 🔄 既存DTOファイルの1:1対応への移行

### Phase 2: Domain Implementation (Priority Order)
1. **Users** (8 request + 8 response = 16 files) - 最重要・依存関係が多い
2. **Goals** (2 request + 4 response = 6 files) - ユーザー機能の基本
3. **Events** (3 request + 6 response = 9 files) - コア機能
4. **Titles** (0 request + 4 response = 4 files) - ゲーミフィケーション
5. **Attendance** (0 request + 5 response = 5 files) - 統計・ランキング

### Implementation Rules
- 1ファイル = 1エンドポイント
- APIドキュメントとの完全一致を保証
- 新規エンドポイント追加時は必ず対応DTOファイルを作成