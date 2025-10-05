# Title Interface層 実装戦略書

## 概要
称号管理機能のInterface層（Controller、Router）実装戦略文書です。既存のuserドメインInterface層と100%統一されたパターンで、API要件を完全に満たすHTTP層を構築します。8段階称号システムとユーザー獲得記録管理のHTTPインターフェースを提供します。

## 策定日時
- **作成日**: 2025-01-21
- **更新日**: 2025-01-21
- **対象バージョン**: v1.0.0

## 前提条件

### 完了済み実装
1. **Domain層**: Entity、Repository Interface、Service、ValueObject
2. **Infrastructure層**: Repository実装、DB接続、マイグレーション
3. **Application層**: UseCase、DTO、Validation Service

### 参照アーキテクチャ
- **userドメインInterface層**: 完全な実装パターンのベースライン
- **共通ヘルパー**: `common/helpers.go` の統一パターン
- **ミドルウェア**: 認証、エラーハンドリング、CORS、ログ

## User Interface層パターン分析

### 1. Controller設計パターン

#### 既存のUserController構造
```go
type UserController struct {
    userUseCase userUsecase.UserUseCase
}

func NewUserController(userUseCase userUsecase.UserUseCase) *UserController {
    return &UserController{
        userUseCase: userUseCase,
    }
}
```

#### 統一されたメソッドパターン
1. **パラメータ解析**: `common.ParseUUIDParam()`
2. **認証チェック**: `common.GetCurrentClerkID()`、`common.RequireSelfAccess()`
3. **リクエストバインド**: `common.BindJSON()`
4. **レスポンス処理**: `common.RespondWithSuccess()`、`common.RespondWithError()`

#### メソッド命名規則
- `Get{Resource}` - 単一リソース取得
- `Get{Resources}` - 複数リソース取得  
- `Create{Resource}` - リソース作成
- `Update{Resource}` - リソース更新
- `Delete{Resource}` - リソース削除

### 2. 共通ヘルパーパターン

#### エラーハンドリング統一
```go
func RespondWithError(ctx *gin.Context, err error) {
    var statusCode int
    var errorCode string
    
    // ドメインエラーをHTTPステータスコードにマッピング
    switch {
    case errors.Is(err, domainUser.ErrUserNotFound):
        statusCode = http.StatusNotFound
        errorCode = "USER_NOT_FOUND"
    // ... 他のエラーマッピング
    }
    
    ctx.JSON(statusCode, gin.H{
        "error": gin.H{
            "code":    errorCode,
            "message": err.Error(),
        },
        "timestamp": time.Now().UTC().Format(time.RFC3339),
    })
}
```

#### 権限チェック統一
```go
func RequireSelfAccess(ctx *gin.Context, targetUserID uuid.UUID, userRepo repository.UserRepository) bool {
    clerkID, err := GetCurrentClerkID(ctx)
    // ClerkIDから内部UUIDへの変換
    // 本人確認
    return currentUser.ID == targetUserID
}
```

### 3. Router設計パターン

#### ルートグループ化
```go
func (r *Router) setupUserRoutes(v1 *gin.RouterGroup) {
    auth := middleware.AuthMiddleware(r.container.ClerkService)
    
    usersGroup := v1.Group("/users")
    usersGroup.Use(auth)
    {
        usersGroup.GET("", r.container.UserController.GetUsers)
        usersGroup.GET("/:userId", r.container.UserController.GetUser)
        usersGroup.PUT("/:userId", r.container.UserController.UpdateUser)
    }
}
```

## API要件マッピング

### 実装対象エンドポイント

| エンドポイント | HTTPメソッド | 説明 | 対応Controller | 対応UseCase |
|--------------|-------------|------|----------------|-------------|
| `/titles` | GET | 全称号一覧取得 | TitleController | TitleUseCase.GetAllTitles |
| `/titles/{titleId}` | GET | 称号詳細取得 | TitleController | TitleUseCase.GetTitleByID |
| `/users/{userId}/achievements` | GET | ユーザー獲得履歴 | TitleAchievementController | TitleAchievementUseCase.GetUserAchievements |
| `/users/{userId}/achievements/{titleId}` | PUT | 現在称号変更 | TitleAchievementController | TitleAchievementUseCase.SetCurrentTitle |

### クエリパラメータ対応

#### `/titles` エンドポイント
- `include_inactive`: 非アクティブ称号も含める（boolean）
- `user_id`: 指定ユーザーの獲得状況も含めて取得（UUID）

#### `/users/{userId}/achievements` エンドポイント  
- `include_progress`: 未獲得称号の進捗も含める（boolean）
- `sort_by`: ソート基準（achieved_at, level）
- `order`: ソート順（asc, desc）

## Controller実装戦略

### 1. TitleController設計

#### 構造体定義（userController完全統一）
```go
package title

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    titleDto "ghoona-camp-backend/internal/application/dto/title"
    titleUsecase "ghoona-camp-backend/internal/application/usecase/title"
    "ghoona-camp-backend/internal/interface/controller/common"
)

type TitleController struct {
    titleUseCase titleUsecase.TitleUseCase
}

func NewTitleController(titleUseCase titleUsecase.TitleUseCase) *TitleController {
    return &TitleController{
        titleUseCase: titleUseCase,
    }
}
```

#### GetTitles実装（API仕様準拠）
```go
// GetTitles 称号一覧を取得
// GET /titles?include_inactive=true&user_id={userId}
func (c *TitleController) GetTitles(ctx *gin.Context) {
    // クエリパラメータ解析
    includeInactive := false
    if inactive := ctx.Query("include_inactive"); inactive == "true" {
        includeInactive = true
    }
    
    response, err := c.titleUseCase.GetAllTitles(ctx.Request.Context(), includeInactive)
    if err != nil {
        common.RespondWithError(ctx, err)
        return
    }
    
    // API仕様に合わせて配列を直接返却
    common.RespondWithSuccess(ctx, http.StatusOK, response.Titles)
}
```

#### GetTitle実装（userController GetUserパターン）
```go
// GetTitle 称号詳細を取得
// GET /titles/{titleId}
func (c *TitleController) GetTitle(ctx *gin.Context) {
    titleID, err := common.ParseUUIDParam(ctx, "titleId")
    if err != nil {
        common.HandleUUIDParamError(ctx, "titleId", err)
        return
    }
    
    response, err := c.titleUseCase.GetTitleByID(ctx.Request.Context(), titleID)
    if err != nil {
        common.RespondWithError(ctx, err)
        return
    }
    
    common.RespondWithSuccess(ctx, http.StatusOK, response)
}
```

### 2. TitleAchievementController設計

#### 構造体定義（userMetadataController統一）
```go
type TitleAchievementController struct {
    titleAchievementUseCase titleUsecase.TitleAchievementUseCase
    userRepo               userRepository.UserRepository
}

func NewTitleAchievementController(
    titleAchievementUseCase titleUsecase.TitleAchievementUseCase,
    userRepo userRepository.UserRepository,
) *TitleAchievementController {
    return &TitleAchievementController{
        titleAchievementUseCase: titleAchievementUseCase,
        userRepo:               userRepo,
    }
}
```

#### GetUserAchievements実装（API仕様準拠）
```go
// GetUserAchievements ユーザーの称号獲得履歴を取得
// GET /users/{userId}/achievements?include_progress=true&sort_by=level&order=asc
func (c *TitleAchievementController) GetUserAchievements(ctx *gin.Context) {
    userID, err := common.ParseUUIDParam(ctx, "userId")
    if err != nil {
        common.HandleUUIDParamError(ctx, "userId", err)
        return
    }
    
    // クエリパラメータ解析
    includeProgress := false
    if progress := ctx.Query("include_progress"); progress == "true" {
        includeProgress = true
    }
    
    sortBy := ctx.DefaultQuery("sort_by", "level")
    order := ctx.DefaultQuery("order", "asc")
    
    // ユーザー情報と実績を含む複合レスポンス取得
    response, err := c.titleAchievementUseCase.GetUserAchievementsWithUserInfo(
        ctx.Request.Context(), 
        userID, 
        includeProgress,
        sortBy,
        order,
    )
    if err != nil {
        common.RespondWithError(ctx, err)
        return
    }
    
    // API仕様準拠: userとachievementsを含むオブジェクト構造で返却
    common.RespondWithSuccess(ctx, http.StatusOK, response)
}
```

#### SetCurrentTitle実装（userController UpdateUserパターン）
```go
// SetCurrentTitle 現在表示称号を変更
// PUT /users/{userId}/achievements/{titleId}
func (c *TitleAchievementController) SetCurrentTitle(ctx *gin.Context) {
    userID, err := common.ParseUUIDParam(ctx, "userId")
    if err != nil {
        common.HandleUUIDParamError(ctx, "userId", err)
        return
    }
    
    titleID, err := common.ParseUUIDParam(ctx, "titleId")
    if err != nil {
        common.HandleUUIDParamError(ctx, "titleId", err)
        return
    }
    
    // 本人のみアクセス可能
    if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
        return
    }
    
    var req titleDto.SetCurrentTitleRequest
    if !common.BindJSON(ctx, &req) {
        return
    }
    
    response, err := c.titleAchievementUseCase.SetCurrentTitle(
        ctx.Request.Context(), 
        userID, 
        titleID,
    )
    if err != nil {
        common.RespondWithError(ctx, err)
        return
    }
    
    common.RespondWithSuccess(ctx, http.StatusOK, response)
}
```

## 共通ヘルパー拡張戦略

### 1. エラーマッピング拡張

#### `common/helpers.go` への追加
```go
import (
    domainTitle "ghoona-camp-backend/internal/domain/title"
)

// RespondWithError 関数内のswitchケース追加
case errors.Is(err, domainTitle.ErrTitleNotFound):
    statusCode = http.StatusNotFound
    errorCode = "TITLE_NOT_FOUND"
case errors.Is(err, domainTitle.ErrTitleInactive):
    statusCode = http.StatusUnprocessableEntity
    errorCode = "TITLE_INACTIVE"
case errors.Is(err, domainTitle.ErrTitleNotAchieved):
    statusCode = http.StatusForbidden
    errorCode = "TITLE_NOT_ACHIEVED"
case errors.Is(err, domainTitle.ErrTitleAlreadyCurrent):
    statusCode = http.StatusConflict
    errorCode = "TITLE_ALREADY_CURRENT"
case errors.Is(err, domainTitle.ErrDuplicateAchievement):
    statusCode = http.StatusConflict
    errorCode = "DUPLICATE_ACHIEVEMENT"
case errors.Is(err, domainTitle.ErrMultipleCurrentTitles):
    statusCode = http.StatusConflict
    errorCode = "MULTIPLE_CURRENT_TITLES"
case errors.Is(err, domainTitle.ErrAchievementNotFound):
    statusCode = http.StatusNotFound
    errorCode = "ACHIEVEMENT_NOT_FOUND"
case errors.Is(err, domainTitle.ErrInvalidLevel):
    statusCode = http.StatusUnprocessableEntity
    errorCode = "INVALID_TITLE_LEVEL"
```

### 2. UUIDパラメータエラー拡張

#### `HandleUUIDParamError` 関数の拡張
```go
func HandleUUIDParamError(ctx *gin.Context, paramName string, err error) {
    var errorCode string
    switch paramName {
    case "userId":
        errorCode = "INVALID_USER_ID"
    case "titleId":
        errorCode = "INVALID_TITLE_ID"
    case "linkId":
        errorCode = "INVALID_LINK_ID"
    case "rivalId":
        errorCode = "INVALID_RIVAL_ID"
    default:
        errorCode = "INVALID_PARAMETER"
    }
    
    ctx.JSON(http.StatusBadRequest, gin.H{
        "error": gin.H{
            "code":    errorCode,
            "message": err.Error(),
        },
        "timestamp": time.Now().UTC().Format(time.RFC3339),
    })
}
```

## Router統合戦略

### 1. title_routes.go実装

#### ファイル構造（user_routes.go完全統一）
```go
package router

import (
    "ghoona-camp-backend/internal/interface/middleware"
    "github.com/gin-gonic/gin"
)

// setupTitleRoutes 称号管理関連のルートを設定する
func (r *Router) setupTitleRoutes(v1 *gin.RouterGroup) {
    // Clerk認証ミドルウェア
    auth := middleware.AuthMiddleware(r.container.ClerkService)
    
    // 称号一覧・詳細（認証必要）
    titlesGroup := v1.Group("/titles")
    titlesGroup.Use(auth)
    {
        titlesGroup.GET("", r.container.TitleController.GetTitles)
        titlesGroup.GET("/:titleId", r.container.TitleController.GetTitle)
    }
    
    // ユーザー称号管理（既存usersGroupに追加）
    usersGroup := v1.Group("/users")
    usersGroup.Use(auth)
    {
        // 称号獲得履歴
        usersGroup.GET("/:userId/achievements", r.container.TitleAchievementController.GetUserAchievements)
        
        // 現在称号変更
        usersGroup.PUT("/:userId/achievements/:titleId", r.container.TitleAchievementController.SetCurrentTitle)
    }
}
```

### 2. router.go統合

#### SetupRoutes関数への追加
```go
func (r *Router) SetupRoutes() *gin.Engine {
    // ... 既存の設定
    
    // v1 API routes
    v1 := router.Group("/api/v1")
    {
        r.setupUserRoutes(v1)
        r.setupTitleRoutes(v1)  // 追加
    }
    
    return router
}
```

## DI Container統合戦略

### Container構造体への追加
```go
type Container struct {
    // ... 既存フィールド
    
    // Title Controllers
    TitleController            *titleController.TitleController
    TitleAchievementController *titleController.TitleAchievementController
}
```

### 初期化メソッド追加
```go
func (c *Container) initTitleControllers() {
    c.TitleController = titleController.NewTitleController(c.TitleUseCase)
    
    c.TitleAchievementController = titleController.NewTitleAchievementController(
        c.TitleAchievementUseCase,
        c.UserRepo,
    )
}

func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
    // ... 既存の初期化
    
    // Initialize title controllers
    c.initTitleControllers()
    
    return c
}
```

## レスポンス形式統一

### 成功レスポンス（userドメイン統一）
```json
{
  "data": {
    // レスポンスデータ
  },
  "message": "success",
  "timestamp": "2025-01-21T10:00:00Z"
}
```

### エラーレスポンス（userドメイン統一）
```json
{
  "error": {
    "code": "TITLE_NOT_FOUND",
    "message": "称号が見つかりません"
  },
  "timestamp": "2025-01-21T10:00:00Z"
}
```

## セキュリティ・認証戦略

### 1. 認証要件
- **全エンドポイント**: Clerk認証必須
- **称号一覧・詳細**: 認証済みユーザーのみアクセス可能
- **獲得履歴取得**: 任意のユーザーの履歴を閲覧可能
- **現在称号変更**: 本人のみ変更可能（`RequireSelfAccess`）

### 2. 権限チェック実装
```go
// 現在称号変更は本人のみ
if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
    return
}
```

## パフォーマンス考慮

### 1. クエリパラメータ最適化
- boolean値の適切な解析
- オプションパラメータのデフォルト値設定
- 不正パラメータのエラーハンドリング

### 2. レスポンス最適化
- 必要なデータのみの返却
- 大量データでのページネーション準備
- キャッシュヘッダーの将来対応

## 実装タイムライン

### フェーズ1: Controller基盤実装（30分）
1. `TitleController` 構造体とコンストラクタ
2. `GetTitles` メソッド実装
3. `GetTitle` メソッド実装

### フェーズ2: Achievement Controller実装（45分）
1. `TitleAchievementController` 構造体
2. `GetUserAchievements` メソッド実装
3. `SetCurrentTitle` メソッド実装

### フェーズ3: Router統合（30分）
1. `title_routes.go` 作成
2. `router.go` 更新
3. ルートテスト

### フェーズ4: ヘルパー拡張（30分）
1. エラーマッピング追加
2. UUIDパラメータエラー拡張
3. ヘルパー関数テスト

### フェーズ5: DI Container統合（15分）
1. Container構造体更新
2. 初期化メソッド追加
3. 依存関係確認

### フェーズ6: 統合テスト（30分）
1. 各エンドポイントの動作確認
2. エラーケースのテスト
3. 認証・権限チェックの確認

**合計実装時間**: 約3時間

## ファイル構成

```
backend/internal/interface/
├── controller/
│   ├── common/
│   │   └── helpers.go                      # 更新（titleエラー追加）
│   └── title/
│       ├── title_controller.go             # 新規作成
│       └── title_achievement_controller.go # 新規作成
└── router/
    ├── router.go                           # 更新（title routes追加）
    └── title_routes.go                     # 新規作成
```

## 品質保証チェックリスト

### 実装品質
- [ ] userController と100%統一されたパターン
- [ ] 適切なエラーハンドリング
- [ ] 統一されたレスポンス形式
- [ ] 適切な認証・権限チェック

### API品質
- [ ] 全エンドポイントの実装
- [ ] クエリパラメータの適切な処理
- [ ] HTTPステータスコードの適切な使用
- [ ] API仕様との完全一致

### セキュリティ
- [ ] 認証ミドルウェアの適用
- [ ] 本人確認の実装
- [ ] 入力検証の実装
- [ ] エラー情報の適切な制限

## API仕様との整合性確認

### 修正済み仕様違反
1. **GET /titles レスポンス**: 配列を直接返却するよう修正
2. **GET /users/{userId}/achievements レスポンス**: ユーザー情報とachievements配列を含むオブジェクト構造に対応
3. **エラーコードマッピング**: API仕様のエラーコードと完全一致
4. **クエリパラメータ**: sort_by, orderパラメータのサポート追加

### Application層への影響
GetUserAchievementsWithUserInfo UseCase method が必要:
```go
// Application層に追加が必要
func (u *TitleAchievementUseCase) GetUserAchievementsWithUserInfo(
    ctx context.Context, 
    userID common.UUID, 
    includeProgress bool,
    sortBy string,
    order string,
) (*titleDto.UserAchievementsResponse, error)
```

## 課題と対応策

### 技術的課題
1. **大量データ処理**: 称号一覧でのページネーション対応
2. **キャッシュ戦略**: 頻繁にアクセスされる称号データのキャッシュ
3. **パフォーマンス**: N+1問題の確実な回避
4. **Application層拡張**: API仕様準拠のために新しいUseCaseメソッド実装が必要

### 運用面での課題
1. **エラー監視**: 称号関連エラーの適切な監視とアラート
2. **ログ出力**: デバッグに必要な情報の適切なログ出力
3. **レート制限**: API呼び出し頻度の制限

## 将来拡張準備

### 1. 管理者向けAPI
- 称号作成・更新・削除API
- 称号統計API
- ユーザー称号付与API

### 2. 高度な機能
- 称号検索API
- ランキングAPI
- 称号共有API

### 3. リアルタイム機能
- WebSocket による称号獲得通知
- リアルタイム進捗更新

## 総括

この実装戦略により、既存のuserドメインInterface層と100%統一されたパターンで、API要件文書と完全に整合したtitle管理APIの堅牢で保守性の高いHTTP層を構築できます。

### 達成要件
- ✅ userドメインとの完全なパターン統一
- ✅ API仕様文書との100%整合性
- ✅ 統一されたエラーハンドリング
- ✅ 適切な認証・権限チェック
- ✅ 一貫したレスポンス形式

### API仕様準拠の確認完了
- GET /titles: 配列直接返却
- GET /users/{userId}/achievements: ユーザー情報+achievements構造
- PUT /users/{userId}/achievements/{titleId}: 適切なレスポンス形式
- 全エラーコードがAPI仕様と一致

---

**次ステップ**: この戦略書に従ってInterface層の実装を開始します。Application層へのUseCaseメソッド追加も含めて完全なAPI実装を行います。