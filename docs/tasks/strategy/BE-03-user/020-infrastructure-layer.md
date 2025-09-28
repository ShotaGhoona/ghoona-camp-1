# BE-03-user-02 Infrastructure層実装戦略書

## 概要
BE-03-user-02ではユーザー管理機能のInfrastructure層を実装します。具体的にはGORMモデルとリポジトリ実装を行い、Domain層で定義したインターフェースを具象化します。

## 実装範囲

### 対象ファイル
```
internal/infrastructure/gorm/
├── model/
│   └── user.go                           # ✏️ 4つのGORMモデル定義
└── repository/
    └── user_repository.go                # ✏️ 4つのリポジトリ実装
```

### 実装対象
1. **GORMモデル** - DBテーブル構造に対応するGoStruct
2. **リポジトリ実装** - Domain層のインターフェースの具象化

## GORM モデル設計

### 1. User モデル
```go
type User struct {
    ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    ClerkID   string     `gorm:"uniqueIndex;not null;size:255"`
    Email     string     `gorm:"uniqueIndex;not null;size:255"`
    Username  *string    `gorm:"size:100"`
    AvatarURL *string    `gorm:"type:text"`
    DiscordID *string    `gorm:"uniqueIndex;size:255"`
    IsActive  bool       `gorm:"default:true"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### 2. UserMetadata モデル
```go
type UserMetadata struct {
    ID               uuid.UUID     `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    UserID           uuid.UUID     `gorm:"type:uuid;uniqueIndex;not null"`
    DisplayName      *string       `gorm:"size:100"`
    ProfileImageURL  *string       `gorm:"type:text"`
    Tagline          *string       `gorm:"size:150"`
    Bio              *string       `gorm:"type:text"`
    Vision           *string       `gorm:"type:text"`
    VisionPublic     bool          `gorm:"default:false"`
    Timezone         string        `gorm:"size:50;default:'Asia/Tokyo'"`
    Skills           pq.StringArray `gorm:"type:text[]"`
    Interests        pq.StringArray `gorm:"type:text[]"`
    User             User          `gorm:"foreignKey:UserID"`
    CreatedAt        time.Time
    UpdatedAt        time.Time
}
```

### 3. UserSocialLink モデル
```go
type UserSocialLink struct {
    ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
    Platform  string    `gorm:"not null;size:50"`
    URL       string    `gorm:"not null;type:text"`
    Title     *string   `gorm:"size:100"`
    IsPublic  bool      `gorm:"default:true"`
    User      User      `gorm:"foreignKey:UserID"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### 4. UserRival モデル
```go
type UserRival struct {
    ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    UserID      uuid.UUID `gorm:"type:uuid;not null;index"`
    RivalUserID uuid.UUID `gorm:"type:uuid;not null;index"`
    User        User      `gorm:"foreignKey:UserID"`
    RivalUser   User      `gorm:"foreignKey:RivalUserID"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

## リポジトリ実装設計

### 共通パターン
- `gorm.DB`への依存性注入
- ドメインエンティティ ↔ GORMモデルの変換メソッド
- エラーハンドリング（`gorm.ErrRecordNotFound` → domain errors）
- トランザクション対応

### 1. UserRepository実装
```go
type userRepository struct {
    db *gorm.DB
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
func (r *userRepository) GetByClerkID(ctx context.Context, clerkID string) (*entity.User, error)
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error)
func (r *userRepository) Create(ctx context.Context, user *entity.User) error
func (r *userRepository) Update(ctx context.Context, user *entity.User) error
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error
```

### 2. UserMetadataRepository実装
```go
func (r *userMetadataRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserMetadata, error)
func (r *userMetadataRepository) Create(ctx context.Context, metadata *entity.UserMetadata) error
func (r *userMetadataRepository) Update(ctx context.Context, metadata *entity.UserMetadata) error
func (r *userMetadataRepository) Delete(ctx context.Context, userID uuid.UUID) error
```

### 3. UserSocialLinkRepository実装
```go
func (r *userSocialLinkRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserSocialLink, error)
func (r *userSocialLinkRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.UserSocialLink, error)
func (r *userSocialLinkRepository) Create(ctx context.Context, link *entity.UserSocialLink) error
func (r *userSocialLinkRepository) Update(ctx context.Context, link *entity.UserSocialLink) error
func (r *userSocialLinkRepository) Delete(ctx context.Context, id uuid.UUID) error
```

### 4. UserRivalRepository実装
```go
func (r *userRivalRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserRival, error)
func (r *userRivalRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int, error)
func (r *userRivalRepository) Create(ctx context.Context, rival *entity.UserRival) error
func (r *userRivalRepository) Delete(ctx context.Context, id uuid.UUID) error
```

## 実装方針

### 1. 責任の分離
- **GORMモデル**: DB構造のみに専念
- **リポジトリ**: CRUD操作とドメインエンティティ変換
- **ドメインロジック**: ドメイン層に委譲

### 2. エラーハンドリング
```go
// GORM NotFound → Domain Error変換例
if errors.Is(err, gorm.ErrRecordNotFound) {
    return nil, user.ErrUserNotFound
}
```

### 3. パフォーマンス考慮
- 必要な場合のみPreload使用
- インデックス設定でクエリ最適化
- N+1問題の回避

### 4. テスタビリティ
- インターフェース実装の確実性
- モック化可能な設計
- トランザクション境界の明確化

## 変換ロジック

### ドメインエンティティ → GORMモデル
```go
func toGormUser(domainUser *entity.User) *model.User {
    return &model.User{
        ID:        domainUser.ID,
        ClerkID:   domainUser.ClerkID,
        Email:     domainUser.Email,
        Username:  domainUser.Username,
        AvatarURL: domainUser.AvatarURL,
        DiscordID: domainUser.DiscordID,
        IsActive:  bool(domainUser.Status.IsActiveState()),
        CreatedAt: domainUser.CreatedAt,
        UpdatedAt: domainUser.UpdatedAt,
    }
}
```

### GORMモデル → ドメインエンティティ
```go
func toDomainUser(gormUser *model.User) *entity.User {
    status := value.UserStatusActive
    if !gormUser.IsActive {
        status = value.UserStatusInactive
    }
    
    return &entity.User{
        ID:        gormUser.ID,
        ClerkID:   gormUser.ClerkID,
        Email:     gormUser.Email,
        Username:  gormUser.Username,
        AvatarURL: gormUser.AvatarURL,
        DiscordID: gormUser.DiscordID,
        Status:    status,
        CreatedAt: gormUser.CreatedAt,
        UpdatedAt: gormUser.UpdatedAt,
    }
}
```

## 注意点

### 1. シンプル実装
- 複雑なクエリ最適化は避ける
- 基本的なCRUD操作に集中
- プリミティブな実装で要件充足

### 2. 型安全性
- `pq.StringArray`でPostgreSQLの配列型対応
- UUIDの適切なハンドリング
- null許可フィールドのポインタ使用

### 3. データベース制約
- 一意制約はDB制約で担保
- 外部キー制約で参照整合性確保
- インデックスで検索パフォーマンス最適化

## 成果物
1. **user.go** - 4つのGORMモデル定義
2. **user_repository.go** - 4つのリポジトリ実装クラス

オーバーエンジニアリングを避け、要件を満たす最小限の実装で進めます。