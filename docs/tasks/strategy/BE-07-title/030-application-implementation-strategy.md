# Title Application層 実装戦略書

## 概要
称号ドメインのApplication層実装の包括的戦略を記述します。既存のuserドメインApplication層のパターンに従い、8段階称号システムと称号獲得追跡、現在称号管理のためのDTO、UseCase、およびサポートインフラストラクチャを提供します。

## 実装日時
- 戦略策定: 2025-01-21
- 実装開始予定: 2025-01-21
- 完了予定: 2025-01-28（1週間）

## User Domain Application層パターン分析

### 1. DTOデザインパターン

#### Userドメインからのパターン分析
**一貫した構造の観察:**
- **Request DTOs**: `Create*Request`, `Update*Request`, `Add*Request`
- **Response DTOs**: `*Response`, `*ListResponse`  
- **Validation Tags**: `binding`タグの広範囲使用
- **ポインタ使用**: オプションフィールドはnill可能な値にポインタ使用
- **エンティティ変換**: エンティティからDTOへの`*FromEntity()`関数
- **リスト変換**: スライス変換用の`*ListFromEntities()`

**命名規則:**
```go
// リクエストパターン
type CreateUserRequest struct {}
type UpdateUserRequest struct {}
type AddRivalRequest struct {}

// レスポンスパターン  
type UserResponse struct {}
type UserListResponse struct {}
type RivalResponse struct {}
type RivalListResponse struct {}
```

#### Title Domain DTO実装戦略

**API要件分析に基づく:**

1. **Title DTOs**:
   - `TitleResponse` - 個別称号情報
   - `TitleListResponse` - 複数称号とメタデータ
   - 作成/更新リクエストは不要（管理者管理）

2. **Achievement DTOs**:
   - `AchievementResponse` - 個別獲得記録
   - `AchievementListResponse` - ユーザー獲得履歴
   - `SetCurrentTitleRequest` - 現在称号変更リクエスト
   - `UserTitleProgressResponse` - 次の称号への進捗

### 2. UseCase実装戦略

#### Userドメインからのパターン分析
**一貫したインターフェース構造:**
```go
type UserUseCase interface {
    GetUserByID(ctx context.Context, userID common.UUID) (*user.UserResponse, error)
    CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error)
    UpdateUser(ctx context.Context, userID common.UUID, req *user.UpdateUserRequest) (*user.UserResponse, error)
    DeleteUser(ctx context.Context, userID common.UUID) error
}
```

**実装パターン:**
- コンストラクタ関数による依存注入
- `transaction.Manager`を使用したトランザクション管理
- ビジネスロジック用ドメインサービス統合
- 入力検証用バリデーションサービス使用
- ドメイン固有エラーによるエラーハンドリング
- リターン文でのエンティティ→DTO変換

#### Title Domain UseCase戦略

**APIエンドポイントに基づく必要なUseCase:**

1. **TitleUseCase**:
   - `GetAllTitles(ctx context.Context, includeInactive bool) (*title.TitleListResponse, error)`
   - `GetTitleByID(ctx context.Context, titleID common.UUID) (*title.TitleResponse, error)`

2. **TitleAchievementUseCase**:
   - `GetUserAchievements(ctx context.Context, userID common.UUID, includeProgress bool) (*title.AchievementListResponse, error)`
   - `SetCurrentTitle(ctx context.Context, userID common.UUID, titleID common.UUID) (*title.AchievementResponse, error)`
   - `GetUserProgress(ctx context.Context, userID common.UUID) (*title.UserTitleProgressResponse, error)`

### 3. エラーハンドリングパターン

#### Userドメインエラーパターン分析
```go
// UseCaseでのドメインエラー使用
if userEntity == nil {
    return nil, domainUser.ErrUserNotFound
}

// トランザクションエラーハンドリング
err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
    // ビジネスロジック
    return nil
})
if err != nil {
    return nil, err
}
```

#### Title Domain エラー戦略
**一貫したエラーハンドリング:**
- Titleドメインエラー使用: `titleDomain.ErrTitleNotFound`, `titleDomain.ErrTitleNotAchieved`
- トランザクション境界を通してエラーコンテキストを維持
- デバッグ用エラーチェーン保持
- インフラストラクチャ漏れなしの適切なドメインエラー返却

### 4. トランザクション管理アプローチ

#### Userドメイントランザクションパターン
```go
var response *user.UserResponse

err := u.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
    // すべてのデータベース操作でtxCtxを使用
    // ビジネスロジック検証
    // エンティティ作成/変更
    // リポジトリ操作
    // DTO変換
    response = user.UserResponseFromEntity(entity)
    return nil
})

if err != nil {
    return nil, err
}
return response, nil
```

#### Title Domain トランザクション戦略
**アトミック操作が必要:**
- 現在称号変更（前回の解除と新規設定をアトミックに）
- 進捗更新付き獲得記録作成
- クロスドメイン操作（称号 + ユーザー検証）

**トランザクションスコープ:**
- 単一集約変更: シンプルトランザクション
- 複数集約操作: 慎重なトランザクション境界
- 読み取り専用操作: トランザクション不要

### 5. バリデーション戦略

#### Userドメインバリデーションパターン分析
```go
// bindingタグによる入力検証
type CreateUserRequest struct {
    Email    string  `json:"email" binding:"required,email"`
    Username *string `json:"username" binding:"omitempty,min=3,max=50"`
}

// バリデーションサービスによるドメイン検証
if err := u.validationService.ValidateUser(userEntity); err != nil {
    return err
}

// ドメインサービスによるビジネスルール検証
if err := u.rivalService.CanSetRival(txCtx, userID, rivalUserID); err != nil {
    return err
}
```

#### Title Domain バリデーション戦略
**多層バリデーション:**

1. **入力検証** (DTOレベル):
   ```go
   type SetCurrentTitleRequest struct {
       IsCurrent bool `json:"isCurrent" binding:"required"`
   }
   ```

2. **ドメイン検証** (エンティティレベル):
   - バリデーションサービス経由でTitleエンティティ検証
   - 獲得記録エンティティ検証

3. **ビジネスルール検証** (サービスレベル):
   - 獲得記録の資格チェック
   - 現在称号変更権限検証
   - 称号アクティベーション状態検証

### 6. 依存注入パターン

#### UserドメインDI分析
```go
func NewUserUseCase(
    userRepo repository.UserRepository,
    metadataRepo repository.UserMetadataRepository,
    socialLinkRepo repository.UserSocialLinkRepository,
    rivalRepo repository.UserRivalRepository,
    userService *service.UserService,
    validationService *service.UserValidationService,
    txManager transaction.Manager,
) UserUseCase {
    return &userUseCase{
        userRepo:          userRepo,
        metadataRepo:      metadataRepo,
        socialLinkRepo:    socialLinkRepo,
        rivalRepo:         rivalRepo,
        userService:       userService,
        validationService: validationService,
        txManager:         txManager,
    }
}
```

#### Title Domain DI戦略
**必要な依存関係:**

1. **TitleUseCase依存関係**:
   - `titleRepo repository.TitleRepository`
   - `validationService *service.TitleValidationService`

2. **TitleAchievementUseCase依存関係**:
   - `userRepo repository.UserRepository` （ユーザー存在確認用）
   - `titleRepo repository.TitleRepository`
   - `achievementRepo repository.TitleAchievementRepository`
   - `titleService *service.TitleService`
   - `validationService *service.TitleValidationService`
   - `txManager transaction.Manager`

**クロスドメイン依存関係:**
- ユーザー検証用Userリポジトリ
- 進捗計算用出席サービス統合

### 7. ファイル構造と命名規則

#### Userドメイン構造分析
```
backend/internal/application/
├── dto/
│   └── user/
│       ├── rival_dto.go
│       ├── social_link_dto.go
│       ├── user_dto.go
│       └── user_metadata_dto.go
└── usecase/
    └── user/
        ├── user_metadata_usecase.go
        ├── user_rival_usecase.go
        ├── user_social_usecase.go
        └── user_usecase.go
```

#### Title Domain構造戦略
**一貫したファイル構成:**
```
backend/internal/application/
├── dto/
│   └── title/
│       ├── title_dto.go                  # Title DTOs
│       ├── achievement_dto.go            # Achievement DTOs
│       └── progress_dto.go               # Progress DTOs
└── usecase/
    └── title/
        ├── title_usecase.go              # 称号管理UseCase
        ├── title_achievement_usecase.go  # 獲得記録管理UseCase
        └── title_progress_usecase.go     # 進捗追跡UseCase
```

**命名規則:**
- パッケージ: `title` （`user`と一貫）
- ファイル: `{entity}_{type}.go` パターン
- インターフェース: `{Entity}UseCase`
- 実装: `{entity}UseCase`
- コンストラクタ: `New{Entity}UseCase`

## 詳細実装プラン

### フェーズ1: DTO実装 (1-2日目)

#### 1.1 Title DTOs (`dto/title/title_dto.go`)
```go
// TitleResponse - 個別称号情報
type TitleResponse struct {
    ID           common.UUID `json:"id"`
    Level        int         `json:"level"`
    NameJP       string      `json:"nameJp"`
    NameEN       string      `json:"nameEn"`
    Description  string      `json:"description"`
    RequiredDays int         `json:"requiredDays"`
    ImageURL     *string     `json:"imageUrl"`
    ColorTheme   *string     `json:"colorTheme"`
    IsActive     bool        `json:"isActive"`
    CreatedAt    time.Time   `json:"createdAt"`
    UpdatedAt    time.Time   `json:"updatedAt"`
}

// TitleListResponse - 複数称号とメタデータ
type TitleListResponse struct {
    Titles []TitleResponse `json:"titles"`
    Total  int             `json:"total"`
}

// 変換関数
func TitleResponseFromEntity(title *entity.Title) *TitleResponse
func TitleListFromEntities(titles []*entity.Title) []TitleResponse
```

#### 1.2 Achievement DTOs (`dto/title/achievement_dto.go`)
```go
// AchievementResponse - 個別獲得記録
type AchievementResponse struct {
    ID          common.UUID   `json:"id"`
    UserID      common.UUID   `json:"userId"`
    Title       TitleResponse `json:"title"`
    AchievedAt  time.Time     `json:"achievedAt"`
    IsCurrent   bool          `json:"isCurrent"`
    CreatedAt   time.Time     `json:"createdAt"`
    UpdatedAt   time.Time     `json:"updatedAt"`
}

// AchievementListResponse - ユーザー獲得履歴
type AchievementListResponse struct {
    UserID       common.UUID           `json:"userId"`
    Achievements []AchievementResponse `json:"achievements"`
    CurrentTitle *AchievementResponse  `json:"currentTitle"`
    Total        int                   `json:"total"`
}

// SetCurrentTitleRequest - 現在称号変更リクエスト
type SetCurrentTitleRequest struct {
    IsCurrent bool `json:"isCurrent" binding:"required"`
}

// 変換関数
func AchievementResponseFromEntity(achievement *entity.TitleAchievement, title *entity.Title) *AchievementResponse
func AchievementListFromEntities(achievements []*entity.TitleAchievement, titles []*entity.Title) []AchievementResponse
```

#### 1.3 Progress DTOs (`dto/title/progress_dto.go`)
```go
// UserTitleProgressResponse - 次の称号への進捗
type UserTitleProgressResponse struct {
    UserID           common.UUID          `json:"userId"`
    CurrentTitle     *AchievementResponse `json:"currentTitle"`
    NextTitle        *TitleResponse       `json:"nextTitle"`
    CurrentDays      int                  `json:"currentDays"`
    RequiredDays     int                  `json:"requiredDays"`
    RemainingDays    int                  `json:"remainingDays"`
    ProgressPercent  float64              `json:"progressPercent"`
    HighestLevel     int                  `json:"highestLevel"`
    EligibleTitles   []TitleResponse      `json:"eligibleTitles"`
}
```

### フェーズ2: UseCase実装 (3-5日目)

#### 2.1 Title UseCase (`usecase/title/title_usecase.go`)
```go
type TitleUseCase interface {
    GetAllTitles(ctx context.Context, includeInactive bool) (*title.TitleListResponse, error)
    GetTitleByID(ctx context.Context, titleID common.UUID) (*title.TitleResponse, error)
}

type titleUseCase struct {
    titleRepo         repository.TitleRepository
    validationService *service.TitleValidationService
}

func NewTitleUseCase(
    titleRepo repository.TitleRepository,
    validationService *service.TitleValidationService,
) TitleUseCase {
    return &titleUseCase{
        titleRepo:         titleRepo,
        validationService: validationService,
    }
}
```

#### 2.2 Achievement UseCase (`usecase/title/title_achievement_usecase.go`)
```go
type TitleAchievementUseCase interface {
    GetUserAchievements(ctx context.Context, userID common.UUID, includeProgress bool) (*title.AchievementListResponse, error)
    SetCurrentTitle(ctx context.Context, userID common.UUID, titleID common.UUID) (*title.AchievementResponse, error)
}

type titleAchievementUseCase struct {
    userRepo          repository.UserRepository
    titleRepo         repository.TitleRepository
    achievementRepo   repository.TitleAchievementRepository
    titleService      *service.TitleService
    validationService *service.TitleValidationService
    txManager         transaction.Manager
}
```

#### 2.3 Progress UseCase (`usecase/title/title_progress_usecase.go`)
```go
type TitleProgressUseCase interface {
    GetUserProgress(ctx context.Context, userID common.UUID) (*title.UserTitleProgressResponse, error)
    CheckEligibleTitles(ctx context.Context, userID common.UUID, attendanceStats *value.AttendanceStatistics) ([]common.UUID, error)
}
```

### フェーズ3: 統合＆テスト (6-7日目)

#### 3.1 DI Container統合
```go
// container.goに追加
type Container struct {
    // 既存の依存関係...
    
    // Title UseCases
    TitleUseCase            titleUseCase.TitleUseCase
    TitleAchievementUseCase titleUseCase.TitleAchievementUseCase
    TitleProgressUseCase    titleUseCase.TitleProgressUseCase
}

func (c *Container) initTitleUseCases() {
    c.TitleUseCase = titleUseCase.NewTitleUseCase(
        c.TitleRepo,
        c.TitleValidationService,
    )
    
    c.TitleAchievementUseCase = titleUseCase.NewTitleAchievementUseCase(
        c.UserRepo,
        c.TitleRepo,
        c.TitleAchievementRepo,
        c.TitleService,
        c.TitleValidationService,
        c.TxManager,
    )
    
    c.TitleProgressUseCase = titleUseCase.NewTitleProgressUseCase(
        c.UserRepo,
        c.TitleRepo,
        c.TitleAchievementRepo,
        c.TitleService,
        c.AttendanceService, // クロスドメイン依存
    )
}
```

## クロスドメイン統合戦略

### 1. 出席サービス統合
**課題**: 称号進捗計算には出席統計が必要
**解決策**: インターフェース抽象化による依存注入

```go
type AttendanceStatsProvider interface {
    GetUserAttendanceStats(ctx context.Context, userID common.UUID) (*value.AttendanceStatistics, error)
}

// TitleProgressUseCaseでこのインターフェースを使用
type titleProgressUseCase struct {
    attendanceProvider AttendanceStatsProvider
    // その他の依存関係...
}
```

### 2. ユーザーサービス統合
**課題**: 称号操作でのユーザー検証
**解決策**: 慎重な境界管理を持つリポジトリレベル依存

```go
// 称号操作前にユーザー存在確認
userEntity, err := u.userRepo.GetByID(ctx, userID)
if err != nil {
    return nil, err
}
if userEntity == nil {
    return nil, domainUser.ErrUserNotFound
}
```

### 3. 権限ベースアクセス制御
**課題**: 自分vs他人の獲得記録での異なるデータアクセス
**解決策**: UseCaseレベル権限チェック

```go
func (u *titleAchievementUseCase) GetUserAchievements(
    ctx context.Context, 
    requestingUserID common.UUID,
    targetUserID common.UUID,
    includeProgress bool,
) (*title.AchievementListResponse, error) {
    isOwnData := requestingUserID == targetUserID
    // 権限に基づく異なるロジック実装
}
```

## エラーハンドリング戦略

### 1. ドメインエラー統合
```go
// Titleドメインエラー
import titleDomain "ghoona-camp-backend/internal/domain/title"

// UseCaseでの使用
if title == nil {
    return nil, titleDomain.ErrTitleNotFound
}

if !achievement.IsCurrentTitle() {
    return nil, titleDomain.ErrTitleNotCurrent
}
```

### 2. クロスドメインエラーハンドリング
```go
// 他ドメインからのエラーを適切にハンドリング
import userDomain "ghoona-camp-backend/internal/domain/user"

if user == nil {
    return nil, userDomain.ErrUserNotFound
}
```

### 3. バリデーションエラー集約
```go
// 複数のバリデーションエラーを収集して返却
type ValidationErrors struct {
    Errors []error
}

func (v ValidationErrors) Error() string {
    // エラーメッセージ結合
}
```

## セキュリティ考慮事項

### 1. 入力検証
- **DTO検証**: 包括的なbindingタグ
- **ドメイン検証**: エンティティレベルビジネスルール検証
- **権限検証**: ユーザー所有権確認

### 2. データアクセス制御
```go
// 自分の獲得記録変更のみ許可
func (u *titleAchievementUseCase) SetCurrentTitle(
    ctx context.Context,
    requestingUserID common.UUID, // JWTトークンから
    targetUserID common.UUID,     // URLパラメータから
    titleID common.UUID,
) (*title.AchievementResponse, error) {
    if requestingUserID != targetUserID {
        return nil, common.ErrUnauthorized
    }
    // ビジネスロジック継続
}
```

### 3. トランザクション整合性
- **アトミック操作**: 重要な状態変更でのトランザクション
- **分離レベル**: 適切なトランザクション分離
- **ロールバック安全性**: ロールバック付き適切なエラーハンドリング

## パフォーマンス最適化戦略

### 1. N+1クエリ防止
```go
// 関連エンティティの効率的ロード
func (u *titleAchievementUseCase) GetUserAchievements(
    ctx context.Context, 
    userID common.UUID,
) (*title.AchievementListResponse, error) {
    // 獲得記録取得
    achievements, err := u.achievementRepo.GetByUserID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    // ユニークな称号ID収集
    titleIDs := make([]common.UUID, len(achievements))
    for i, achievement := range achievements {
        titleIDs[i] = achievement.TitleID
    }
    
    // 称号バッチロード
    titles, err := u.titleRepo.GetByIDs(ctx, titleIDs)
    if err != nil {
        return nil, err
    }
    
    // 効率的ルックアップ用称号マップ作成
    titleMap := make(map[common.UUID]*entity.Title)
    for _, title := range titles {
        titleMap[title.ID] = title
    }
    
    // 効率的ルックアップでDTOに変換
    responses := make([]title.AchievementResponse, len(achievements))
    for i, achievement := range achievements {
        if title, exists := titleMap[achievement.TitleID]; exists {
            responses[i] = *title.AchievementResponseFromEntity(achievement, title)
        }
    }
    
    return &title.AchievementListResponse{
        Achievements: responses,
        Total:        len(responses),
    }, nil
}
```

### 2. キャッシュ戦略準備
```go
// 将来のキャッシュ実装用インターフェース
type TitleCacheProvider interface {
    GetTitle(ctx context.Context, titleID common.UUID) (*entity.Title, error)
    GetAllTitles(ctx context.Context) ([]*entity.Title, error)
    InvalidateTitle(ctx context.Context, titleID common.UUID) error
}

// キャッシュ準備完了のUseCase実装
type titleUseCase struct {
    titleRepo repository.TitleRepository
    cache     TitleCacheProvider // 将来使用のためオプション
}
```

## テスト戦略

### 1. ユニットテスト
**モック依存関係**:
```go
type mockTitleRepository struct {
    titles map[common.UUID]*entity.Title
}

func (m *mockTitleRepository) GetByID(ctx context.Context, id common.UUID) (*entity.Title, error) {
    if title, exists := m.titles[id]; exists {
        return title, nil
    }
    return nil, nil
}
```

**テスト構造**:
```go
func TestTitleUseCase_GetTitleByID(t *testing.T) {
    tests := []struct {
        name          string
        titleID       common.UUID
        mockTitle     *entity.Title
        expectedError error
        expectedTitle *title.TitleResponse
    }{
        {
            name:          "成功 - 称号が見つかった",
            titleID:       testTitleID,
            mockTitle:     testTitleEntity,
            expectedError: nil,
            expectedTitle: expectedTitleResponse,
        },
        {
            name:          "見つからない - 称号が存在しない",
            titleID:       nonExistentTitleID,
            mockTitle:     nil,
            expectedError: titleDomain.ErrTitleNotFound,
            expectedTitle: nil,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // テスト実装
        })
    }
}
```

### 2. 統合テスト
**トランザクションテスト**:
```go
func TestTitleAchievementUseCase_SetCurrentTitle_Transaction(t *testing.T) {
    // 失敗時のトランザクションロールバックテスト
    // 同時変更ハンドリングテスト
    // クロステーブル整合性テスト
}
```

**クロスドメイン統合**:
```go
func TestTitleProgressUseCase_GetUserProgress_WithAttendanceData(t *testing.T) {
    // 出席サービスとの統合テスト
    // 進捗計算精度テスト
    // 外部依存からのエラーハンドリングテスト
}
```

## 実装タイムライン

### 第1週: 基盤 (1-3日目)
- **1日目**: DTO実装とエンティティ変換関数
- **2日目**: 基本UseCaseインターフェースとシンプル読み取り操作
- **3日目**: 複雑なUseCase操作とトランザクション管理

### 第2週: 統合 (4-6日目)
- **4日目**: クロスドメイン統合と依存注入
- **5日目**: エラーハンドリングとバリデーション統合
- **6日目**: パフォーマンス最適化とキャッシュ準備

### 第3週: テスト＆改善 (7-9日目)
- **7日目**: ユニットテスト実装
- **8日目**: 統合テストとトランザクションテスト
- **9日目**: パフォーマンステストと最適化

## 品質保証チェックリスト

### コード品質
- [ ] userドメインパターンとの100%一貫性
- [ ] 包括的入力検証
- [ ] 適切なエラーハンドリングと伝播
- [ ] トランザクション境界管理
- [ ] クロスドメイン依存管理

### ビジネスロジック
- [ ] 8段階称号システム実装
- [ ] タイムスタンプ付き獲得記録追跡
- [ ] 現在称号アトミック更新
- [ ] 権限ベースアクセス制御
- [ ] 進捗計算精度

### パフォーマンス
- [ ] N+1クエリ防止
- [ ] 効率的バッチ操作
- [ ] 適切なインデックス活用
- [ ] メモリ効率的操作

### テスト
- [ ] >95%ユニットテストカバレッジ
- [ ] 統合テストカバレッジ
- [ ] トランザクションテストカバレッジ
- [ ] エラーシナリオカバレッジ
- [ ] パフォーマンステストベースライン

## リスク評価と軽減策

### 技術的リスク
1. **クロスドメイン複雑性**: インターフェース抽象化で軽減
2. **トランザクションデッドロック**: 順序付き操作で軽減
3. **パフォーマンス問題**: バッチ操作とキャッシュ準備で軽減

### ビジネスリスク
1. **データ不整合**: アトミックトランザクションで軽減
2. **権限バイパス**: UseCaseレベルセキュリティで軽減
3. **スケーラビリティ懸念**: 効率的クエリパターンで軽減

## 将来拡張準備

### 1. 通知統合
```go
type TitleNotificationProvider interface {
    NotifyTitleAchieved(ctx context.Context, userID common.UUID, titleID common.UUID) error
    NotifyCurrentTitleChanged(ctx context.Context, userID common.UUID, titleID common.UUID) error
}
```

### 2. バッチ処理統合
```go
type TitleBatchProcessor interface {
    ProcessTitleEligibility(ctx context.Context, userIDs []common.UUID) error
    CalculateUserProgress(ctx context.Context, userID common.UUID) error
}
```

### 3. 分析統合
```go
type TitleAnalyticsProvider interface {
    TrackTitleViewed(ctx context.Context, userID common.UUID, titleID common.UUID) error
    TrackAchievementShared(ctx context.Context, userID common.UUID, titleID common.UUID) error
}
```

## 結論

この実装戦略は、既存のuserドメインApplication層パターンとの100%一貫性を確保しながら、8段階称号システムの包括的機能を提供します。アプローチは以下を優先します：

1. **一貫性**: 同一パターンと命名規則
2. **保守性**: 明確な関心事分離と依存管理
3. **パフォーマンス**: 効率的データアクセスとトランザクション管理
4. **セキュリティ**: 適切な検証と権限チェック
5. **テスト容易性**: 包括的テスト戦略
6. **拡張性**: 将来拡張準備

実装は、userドメインで確立された高いコード品質基準を維持しながら、称号管理APIの堅牢な基盤を提供します。

---

**次ステップ**: この戦略文書に従ってフェーズ1 DTO実装を開始します。