# BE-03-user-04 Interface層実装 戦略書

## 概要
BE-03-user-04では、ユーザー管理のInterface層（コントローラー・ルーティング）を実装します。eagle-aiプロジェクトのパターンに準拠し、Ginフレームワークを使用してRESTful APIエンドポイントを構築します。

## アーキテクチャ方針

### 1. フレームワーク・技術選択
- **HTTPフレームワーク**: Gin (eagle-ai準拠)
- **認証方式**: Clerk認証トークン
- **パラメータ形式**: UUID (eagle-aiはintだがGhoonaCampはUUID)
- **レスポンス形式**: JSON

### 2. eagle-aiパターンとの差異対応

#### 共通点を踏襲
- Ginフレームワーク使用
- コントローラー構造とメソッド命名
- エラーハンドリングパターン (`respondWithError`)
- ルーティング設定の分離

#### Ghoona Camp固有の調整
- **UUID**: `parseIntParam` → `parseUUIDParam`に変更
- **Clerk認証**: eagle-aiの権限システム → Clerk認証ミドルウェア
- **権限チェック**: Authority-based → Clerk + 本人確認ベース

## 実装方針

### 1. コントローラー実装 (`user_controller.go`)

#### 実装するエンドポイント
1. **GET /auth/me** - 現在ユーザー情報取得
2. **GET /users/{userId}** - ユーザー詳細取得
3. **PUT /users/{userId}** - ユーザー基本情報更新
4. **GET /users/{userId}/metadata** - メタデータ取得
5. **PUT /users/{userId}/metadata** - メタデータ更新
6. **GET /users/{userId}/social-links** - ソーシャルリンク一覧
7. **POST /users/{userId}/social-links** - ソーシャルリンク追加
8. **PUT /users/{userId}/social-links/{linkId}** - ソーシャルリンク更新
9. **DELETE /users/{userId}/social-links/{linkId}** - ソーシャルリンク削除
10. **GET /users/{userId}/rivals** - ライバル一覧
11. **POST /users/{userId}/rivals** - ライバル追加
12. **DELETE /users/{userId}/rivals/{rivalId}** - ライバル削除

#### コントローラー構造
```go
type UserController struct {
    userUseCase usecase.UserUseCase
}
```

### 2. ヘルパー関数実装

#### UUID パラメータ解析
```go
func parseUUIDParam(ctx *gin.Context, paramName string) (uuid.UUID, error)
```

#### 認証ユーザー取得
```go
func getCurrentUser(ctx *gin.Context) (uuid.UUID, error)
```

#### 権限チェック
```go
func requireSelfOrAdmin(ctx *gin.Context, targetUserID uuid.UUID) bool
```

### 3. エラーハンドリング統一

#### ドメインエラー → HTTPステータスマッピング
- `ErrUserNotFound` → 404 Not Found
- `ErrDuplicateEmail` → 409 Conflict
- `ErrRivalLimitExceeded` → 422 Unprocessable Entity
- `ErrCannotRivalSelf` → 422 Unprocessable Entity
- バリデーションエラー → 422 Unprocessable Entity

#### レスポンス形式統一
```go
// 成功レスポンス
{
  "data": {...},
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}

// エラーレスポンス
{
  "error": {
    "code": "ERROR_CODE",
    "message": "エラーメッセージ"
  },
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### 4. ルーティング設定 (`user_routes.go`)

#### Clerk認証ミドルウェア統合
```go
func (r *Router) setupUserRoutes() {
    clerkAuth := r.container.ClerkMiddleware.RequireAuth()
    
    // 認証が必要なエンドポイント
    authGroup := r.engine.Group("/", clerkAuth)
    
    // 本人のみアクセス可能
    authGroup.GET("/auth/me", r.container.UserController.GetCurrentUser)
    authGroup.PUT("/users/:userId", requireSelf(), r.container.UserController.UpdateUser)
    
    // 公開アクセス（認証済み）
    authGroup.GET("/users/:userId", r.container.UserController.GetUser)
}
```

#### 権限制御パターン
1. **公開情報**: 認証済みユーザー全員がアクセス可能
2. **本人限定**: 自分の情報のみ編集可能
3. **プライバシー考慮**: is_public設定に基づく情報制御

## 技術仕様

### 1. 依存関係
```go
import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "ghoona-camp-backend/internal/application/usecase"
    userDto "ghoona-camp-backend/internal/application/dto/user"
)
```

### 2. バリデーション
- **DTOバリデーション**: Ginのbindingタグ活用
- **ビジネスロジック**: UseCase層でドメインバリデーション
- **UUID形式**: パラメータ解析時に検証

### 3. レスポンス最適化
- **メタデータ結合**: ユーザー情報とメタデータの効率的な取得
- **N+1問題回避**: 関連データの適切な取得戦略
- **不要フィールド除外**: プライバシー設定に基づく情報フィルタリング

## セキュリティ考慮

### 1. Clerk認証統合
- JWTトークン検証
- ユーザー識別子マッピング (ClerkID ↔ 内部UUID)
- 認証エラーの適切なハンドリング

### 2. 認可制御
- 本人確認チェック
- プライバシー設定の尊重
- 機密情報の適切なマスキング

### 3. 入力検証
- UUID形式検証
- SQLインジェクション防止
- XSS対策

## パフォーマンス考慮

### 1. キャッシュ戦略
- ユーザー情報: 短期キャッシュ（1分）
- メタデータ: 中期キャッシュ（5分）
- 関連データ: 適切なキャッシュ設計

### 2. データ取得最適化
- 必要最小限のデータ取得
- 関連テーブルの効率的なJOIN
- ページネーション対応準備

## 実装順序

1. **ヘルパー関数**: UUID解析、認証、エラーハンドリング
2. **基本コントローラー**: GetCurrentUser, GetUser, UpdateUser
3. **メタデータコントローラー**: メタデータCRUD
4. **ソーシャルリンクコントローラー**: リンク管理
5. **ライバルコントローラー**: ライバル管理
6. **ルーティング設定**: 全エンドポイントの統合
7. **テスト**: 各エンドポイントの動作確認

## 期待する成果物

- `internal/interface/controller/user_controller.go` (完全実装)
- `internal/interface/router/user_routes.go` (ルーティング設定)
- `internal/interface/router/router.go` (メインルーター更新)
- eagle-aiパターンに準拠した保守性の高いInterface層

## Ghoona Camp固有の要件

### API設計準拠
- `docs/requirement/api/user-management.md`の仕様に完全準拠
- RESTful APIの適切な実装
- エラーハンドリングの統一

### Clerk認証連携
- Clerk JWTトークンの検証
- ClerkIDと内部UUIDのマッピング
- 認証状態の適切な管理

### プライバシー制御
- vision_publicフラグの尊重
- is_publicフラグに基づくソーシャルリンク表示制御
- 本人以外からのアクセス時の情報制限

この戦略に基づいてBE-03-user-04の実装を進めます。