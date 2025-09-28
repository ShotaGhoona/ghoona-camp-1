# BE-03-user-03 Application層実装戦略書

## 概要
BE-03-user-03ではユーザー管理機能のApplication層を実装します。DTOとUseCaseを作成し、Domain層とInfrastructure層を結ぶアプリケーションロジックを構築します。

## 実装範囲

### 対象ファイル
```
internal/application/
├── dto/
│   └── user/
│       ├── user_dto.go                       # 🆕 ユーザー関連DTO
│       ├── user_metadata_dto.go              # 🆕 メタデータ関連DTO
│       ├── social_link_dto.go                # 🆕 ソーシャルリンク関連DTO
│       └── rival_dto.go                      # 🆕 ライバル関連DTO
└── usecase/
    └── user_service.go                       # ✏️ 既存ファイル拡張
```

## DTO設計（eagle-aiパターン準拠）

### 1. Request/Response分離方式
eagle-aiと同様に、リクエスト用とレスポンス用DTOを明確に分離

### 2. バリデーション戦略
```go
// Ginのbindingタグを使用（eagle-aiと同様）
type CreateUserRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Username string `json:"username" binding:"required,min=3,max=50"`
}
```

### 3. 変換ロジック
```go
// エンティティからレスポンスDTOへの変換関数
func UserResponseFromEntity(user *entity.User) *UserResponse {
    // 変換ロジック
}
```

## 各DTOファイルの設計

### 1. user_dto.go
```go
// Request DTOs
type CreateUserRequest struct {
    ClerkID   string  `json:"clerk_id" binding:"required"`
    Email     string  `json:"email" binding:"required,email"`
    Username  *string `json:"username" binding:"omitempty,min=3,max=50"`
    AvatarURL *string `json:"avatar_url" binding:"omitempty"`
    DiscordID *string `json:"discord_id" binding:"omitempty"`
}

type UpdateUserRequest struct {
    Username  *string `json:"username" binding:"omitempty,min=3,max=50"`
    AvatarURL *string `json:"avatar_url" binding:"omitempty"`
    DiscordID *string `json:"discord_id" binding:"omitempty"`
}

// Response DTOs
type UserResponse struct {
    ID        uuid.UUID             `json:"id"`
    ClerkID   string                `json:"clerk_id"`
    Email     string                `json:"email"`
    Username  *string               `json:"username"`
    AvatarURL *string               `json:"avatar_url"`
    DiscordID *string               `json:"discord_id"`
    Status    string                `json:"status"`
    Metadata  *UserMetadataResponse `json:"metadata,omitempty"`
    CreatedAt time.Time             `json:"created_at"`
    UpdatedAt time.Time             `json:"updated_at"`
}

type UserListResponse struct {
    Users []UserResponse `json:"users"`
    Total int            `json:"total"`
}

// 変換関数
func UserResponseFromEntity(user *entity.User) *UserResponse
func UserListFromEntities(users []*entity.User) []UserResponse
```

### 2. user_metadata_dto.go
```go
// Request DTOs
type UpdateUserMetadataRequest struct {
    DisplayName      *string  `json:"display_name" binding:"omitempty,max=100"`
    ProfileImageURL  *string  `json:"profile_image_url" binding:"omitempty"`
    Tagline          *string  `json:"tagline" binding:"omitempty,max=150"`
    Bio              *string  `json:"bio" binding:"omitempty,max=1000"`
    Vision           *string  `json:"vision" binding:"omitempty,max=2000"`
    VisionPublic     *bool    `json:"vision_public"`
    Timezone         *string  `json:"timezone" binding:"omitempty"`
    Skills           []string `json:"skills" binding:"max=20,dive,max=50"`
    Interests        []string `json:"interests" binding:"max=20,dive,max=50"`
}

// Response DTOs
type UserMetadataResponse struct {
    ID               uuid.UUID `json:"id"`
    UserID           uuid.UUID `json:"user_id"`
    DisplayName      *string   `json:"display_name"`
    ProfileImageURL  *string   `json:"profile_image_url"`
    Tagline          *string   `json:"tagline"`
    Bio              *string   `json:"bio"`
    Vision           *string   `json:"vision"`
    VisionPublic     bool      `json:"vision_public"`
    Timezone         string    `json:"timezone"`
    Skills           []string  `json:"skills"`
    Interests        []string  `json:"interests"`
    CreatedAt        time.Time `json:"created_at"`
    UpdatedAt        time.Time `json:"updated_at"`
}

// 変換関数
func UserMetadataResponseFromEntity(metadata *entity.UserMetadata) *UserMetadataResponse
```

### 3. social_link_dto.go
```go
// Request DTOs
type CreateSocialLinkRequest struct {
    Platform string  `json:"platform" binding:"required"`
    URL      string  `json:"url" binding:"required,url"`
    Title    *string `json:"title" binding:"omitempty,max=100"`
    IsPublic *bool   `json:"is_public"`
}

type UpdateSocialLinkRequest struct {
    URL      *string `json:"url" binding:"omitempty,url"`
    Title    *string `json:"title" binding:"omitempty,max=100"`
    IsPublic *bool   `json:"is_public"`
}

// Response DTOs
type SocialLinkResponse struct {
    ID        uuid.UUID `json:"id"`
    UserID    uuid.UUID `json:"user_id"`
    Platform  string    `json:"platform"`
    URL       string    `json:"url"`
    Title     *string   `json:"title"`
    IsPublic  bool      `json:"is_public"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type SocialLinkListResponse struct {
    SocialLinks []SocialLinkResponse `json:"social_links"`
    Total       int                  `json:"total"`
}

// 変換関数
func SocialLinkResponseFromEntity(link *entity.UserSocialLink) *SocialLinkResponse
func SocialLinkListFromEntities(links []*entity.UserSocialLink) []SocialLinkResponse
```

### 4. rival_dto.go
```go
// Request DTOs
type AddRivalRequest struct {
    RivalUserID uuid.UUID `json:"rival_user_id" binding:"required"`
}

// Response DTOs
type RivalResponse struct {
    ID          uuid.UUID    `json:"id"`
    UserID      uuid.UUID    `json:"user_id"`
    RivalUser   UserResponse `json:"rival_user"`
    CreatedAt   time.Time    `json:"created_at"`
}

type RivalListResponse struct {
    Rivals []RivalResponse `json:"rivals"`
    Total  int             `json:"total"`
    MaxRivals int          `json:"max_rivals"`
}

// 変換関数
func RivalResponseFromEntity(rival *entity.UserRival, rivalUser *entity.User) *RivalResponse
func RivalListFromEntities(rivals []*entity.UserRival, rivalUsers []*entity.User) []RivalResponse
```

## UseCase設計（eagle-aiパターン準拠）

### UserUseCase インターフェース
```go
type UserUseCase interface {
    // User基本操作
    GetUserByID(ctx context.Context, userID uuid.UUID) (*UserResponse, error)
    GetUserByClerkID(ctx context.Context, clerkID string) (*UserResponse, error)
    GetUsers(ctx context.Context) (*UserListResponse, error)
    CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error)
    UpdateUser(ctx context.Context, userID uuid.UUID, req *UpdateUserRequest) (*UserResponse, error)
    DeleteUser(ctx context.Context, userID uuid.UUID) error

    // Metadata操作
    GetUserMetadata(ctx context.Context, userID uuid.UUID) (*UserMetadataResponse, error)
    UpdateUserMetadata(ctx context.Context, userID uuid.UUID, req *UpdateUserMetadataRequest) (*UserMetadataResponse, error)

    // SocialLink操作
    GetUserSocialLinks(ctx context.Context, userID uuid.UUID) (*SocialLinkListResponse, error)
    CreateSocialLink(ctx context.Context, userID uuid.UUID, req *CreateSocialLinkRequest) (*SocialLinkResponse, error)
    UpdateSocialLink(ctx context.Context, linkID uuid.UUID, req *UpdateSocialLinkRequest) (*SocialLinkResponse, error)
    DeleteSocialLink(ctx context.Context, linkID uuid.UUID) error

    // Rival操作
    GetUserRivals(ctx context.Context, userID uuid.UUID) (*RivalListResponse, error)
    AddRival(ctx context.Context, userID uuid.UUID, req *AddRivalRequest) (*RivalResponse, error)
    RemoveRival(ctx context.Context, rivalID uuid.UUID) error
}
```

### 実装構造
```go
type userUseCase struct {
    userRepo         repository.UserRepository
    metadataRepo     repository.UserMetadataRepository
    socialLinkRepo   repository.UserSocialLinkRepository
    rivalRepo        repository.UserRivalRepository
    userService      *service.UserService
    rivalService     *service.RivalService
    validationService *service.UserValidationService
    txManager        transaction.Manager
}
```

## 実装方針

### 1. トランザクション管理（eagle-aiパターン）
```go
func (u *userUseCase) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
    result, err := u.txManager.ExecuteInTxWithResult(ctx, func(txCtx context.Context) (interface{}, error) {
        // ビジネスロジック実行
        // ...
        return *response, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    responseDTO := result.(UserResponse)
    return &responseDTO, nil
}
```

### 2. バリデーション戦略
- DTOレベル：Ginのbindingタグ
- ビジネスロジック：Domain Serviceを活用
- 重複チェック：リポジトリでの存在確認

### 3. エラーハンドリング
- Domain エラーをそのまま返す
- Application層での追加的なエラー処理
- 適切なエラーメッセージの維持

### 4. 依存性注入
```go
func NewUserUseCase(
    userRepo repository.UserRepository,
    metadataRepo repository.UserMetadataRepository,
    socialLinkRepo repository.UserSocialLinkRepository,
    rivalRepo repository.UserRivalRepository,
    userService *service.UserService,
    rivalService *service.RivalService,
    validationService *service.UserValidationService,
    txManager transaction.Manager,
) UserUseCase
```

## 特別な考慮事項

### 1. Clerk連携
- ClerkIDでのユーザー検索対応
- 外部認証との整合性確保

### 2. ライバル制限
- 最大3人制限のビジネスルール実装
- RivalServiceでの制限チェック

### 3. プライバシー制御
- VisionPublic, IsPublicフラグの適切な処理
- 他ユーザーからの閲覧制限

### 4. スキル・興味管理
- 配列フィールドの適切な処理
- 最大件数制限（20件）の実装

## テスト戦略

### 単体テスト
- 各UseCaseメソッドのテスト
- DTO変換ロジックのテスト
- バリデーションのテスト

### 結合テスト
- トランザクション動作の確認
- エラーハンドリングの確認
- Domain Serviceとの連携確認

## パフォーマンス考慮

### 1. データ取得最適化
- 必要最小限のデータ取得
- 関連データの効率的な取得

### 2. トランザクション最適化
- 適切なトランザクション境界
- デッドロック回避

## 成果物

1. **4つのDTOファイル** - リクエスト/レスポンス、変換ロジック
2. **UserUseCase実装** - 完全なビジネスロジック実装
3. **既存user_service.go拡張** - 新機能との統合

eagle-aiパターンに従い、保守可能で拡張性の高いApplication層を構築します。