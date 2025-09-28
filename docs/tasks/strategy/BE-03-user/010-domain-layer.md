# BE-03-user-01: ユーザー管理Domain層実装戦略書

## 概要
ユーザー管理機能のDomain層を実装します。オニオンアーキテクチャの最内層として、ビジネスロジックとドメインルールを外部依存なしで実装します。

## スコープ

### 実装対象
- **エンティティ**: User、UserMetadata、UserSocialLink、UserRival
- **値オブジェクト**: Platform、UserStatus、PrivacyLevel
- **リポジトリインターフェース**: UserRepository、UserMetadataRepository、UserSocialLinkRepository、UserRivalRepository
- **ドメインサービス**: UserDomainService
- **ドメインエラー**: user_errors.go

### 実装しないもの（スコープ外）
- GORMモデル（Infrastructure層）
- データベース接続（Infrastructure層）
- HTTP ハンドラー（Interface層）
- DTO（Application層）
- ユースケース（Application層）

## アーキテクチャ設計

### ディレクトリ構造
```
backend/internal/domain/user/
├── entity/
│   ├── user.go                 # User エンティティ
│   ├── user_metadata.go        # UserMetadata エンティティ
│   ├── user_social_link.go     # UserSocialLink エンティティ
│   └── user_rival.go           # UserRival エンティティ
├── repository/
│   ├── user_repository.go      # User リポジトリI/F
│   ├── user_metadata_repository.go  # UserMetadata リポジトリI/F
│   ├── user_social_link_repository.go # UserSocialLink リポジトリI/F
│   └── user_rival_repository.go    # UserRival リポジトリI/F
├── service/
│   └── user_domain_service.go  # ドメインサービス
├── value/
│   ├── platform.go             # Platform値オブジェクト
│   ├── user_status.go          # UserStatus値オブジェクト
│   └── privacy_level.go        # PrivacyLevel値オブジェクト
└── errors.go                   # ドメインエラー定義
```

### エンティティ設計方針

#### 1. User エンティティ
```go
type User struct {
    ID        uuid.UUID  // プライマリキー
    ClerkID   string     // Clerk認証ID（外部キー）
    Email     string     // メールアドレス（一意制約）
    Username  *string    // ユーザー名（可変、一意制約）
    AvatarURL *string    // アバター画像URL
    DiscordID *string    // Discord ID（Discord連携用）
    Status    UserStatus // ユーザーステータス
    CreatedAt time.Time  // 作成日時
    UpdatedAt time.Time  // 更新日時
}
```

**ビジネスルール**:
- ClerkIDは不変（変更不可）
- Emailは必須、変更可能（Clerk側との同期要）
- Usernameは任意、変更可能、重複不可
- DiscordIDは任意、Discord連携時のみ設定

#### 2. UserMetadata エンティティ
```go
type UserMetadata struct {
    ID               uuid.UUID     // プライマリキー
    UserID           uuid.UUID     // User外部キー
    DisplayName      *string       // 表示名
    ProfileImageURL  *string       // プロフィール画像URL
    Tagline          *string       // キャッチフレーズ（150文字以内）
    Bio              *string       // 自己紹介（1000文字以内）
    Vision           *string       // ビジョン（500文字以内）
    VisionPublic     bool          // ビジョン公開設定
    Timezone         string        // タイムゾーン（デフォルト: Asia/Tokyo）
    Skills           []string      // スキル（最大20個）
    Interests        []string      // 興味・関心（最大20個）
    PrivacyLevel     PrivacyLevel  // プライバシーレベル
    CreatedAt        time.Time     // 作成日時
    UpdatedAt        time.Time     // 更新日時
}
```

**ビジネスルール**:
- UserIDは必須、User.IDとの整合性必要
- 各フィールドの文字数制限
- Skills/Interestsの個数制限
- VisionPublicはVisionが設定されている場合のみ有効

#### 3. UserSocialLink エンティティ
```go
type UserSocialLink struct {
    ID        uuid.UUID // プライマリキー
    UserID    uuid.UUID // User外部キー
    Platform  Platform  // プラットフォーム
    URL       string    // リンクURL
    Title     *string   // リンクタイトル
    IsPublic  bool      // 公開設定
    CreatedAt time.Time // 作成日時
    UpdatedAt time.Time // 更新日時
}
```

**ビジネスルール**:
- ユーザー1人につき、同一Platformは1つまで
- URLの形式バリデーション
- Platformごとの固有バリデーション

#### 4. UserRival エンティティ
```go
type UserRival struct {
    ID          uuid.UUID // プライマリキー
    UserID      uuid.UUID // User外部キー
    RivalUserID uuid.UUID // ライバルUser外部キー
    CreatedAt   time.Time // 作成日時
    UpdatedAt   time.Time // 更新日時
}
```

**ビジネスルール**:
- 自分自身をライバルに設定不可
- 同一ユーザーとの重複関係不可
- ライバル関係は一方向（相互ではない）

### 値オブジェクト設計

#### Platform（プラットフォーム）
```go
type Platform string
const (
    PlatformTwitter   Platform = "twitter"
    PlatformGitHub    Platform = "github"
    PlatformLinkedIn  Platform = "linkedin"
    PlatformWebsite   Platform = "website"
    PlatformBlog      Platform = "blog"
    PlatformYouTube   Platform = "youtube"
    PlatformInstagram Platform = "instagram"
)
```

#### UserStatus（ユーザーステータス）
```go
type UserStatus string
const (
    UserStatusActive    UserStatus = "active"
    UserStatusInactive  UserStatus = "inactive"
    UserStatusSuspended UserStatus = "suspended"
    UserStatusDeleted   UserStatus = "deleted"
)
```

#### PrivacyLevel（プライバシーレベル）
```go
type PrivacyLevel string
const (
    PrivacyPublic   PrivacyLevel = "public"   // 全体公開
    PrivacyFriends  PrivacyLevel = "friends"  // フレンドのみ
    PrivacyPrivate  PrivacyLevel = "private"  // 非公開
)
```

### ドメインサービス設計

#### UserDomainService
複数エンティティにまたがるビジネスロジックを処理：

```go
type UserDomainService struct {}

// ValidateUsername - ユーザー名の妥当性チェック
func (s *UserDomainService) ValidateUsername(username string) error

// CanSetRival - ライバル設定可能性チェック
func (s *UserDomainService) CanSetRival(userID, rivalUserID uuid.UUID) error

// GenerateDefaultDisplayName - デフォルト表示名生成
func (s *UserDomainService) GenerateDefaultDisplayName(user *User) string

// ValidateSocialLinkURL - ソーシャルリンクURL検証
func (s *UserDomainService) ValidateSocialLinkURL(platform Platform, url string) error
```

### リポジトリインターフェース設計

#### UserRepository
```go
type UserRepository interface {
    FindByID(ctx context.Context, id uuid.UUID) (*User, error)
    FindByClerkID(ctx context.Context, clerkID string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    FindByUsername(ctx context.Context, username string) (*User, error)
    FindAll(ctx context.Context, filters UserFilters) ([]*User, error)
    Create(ctx context.Context, user *User) error
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id uuid.UUID) error
    ExistsByUsername(ctx context.Context, username string) (bool, error)
    ExistsByEmail(ctx context.Context, email string) (bool, error)
}
```

#### UserMetadataRepository
```go
type UserMetadataRepository interface {
    FindByUserID(ctx context.Context, userID uuid.UUID) (*UserMetadata, error)
    Create(ctx context.Context, metadata *UserMetadata) error
    Update(ctx context.Context, metadata *UserMetadata) error
    Delete(ctx context.Context, userID uuid.UUID) error
}
```

## 実装詳細

### 1. バリデーション戦略
- **エンティティレベル**: 基本的な制約（必須項目、文字数等）
- **ドメインサービスレベル**: 複雑なビジネスルール
- **値オブジェクトレベル**: 値の妥当性チェック

### 2. エラーハンドリング
```go
// errors.go
var (
    ErrUserNotFound         = errors.New("ユーザーが見つかりません")
    ErrDuplicateUsername    = errors.New("ユーザー名が既に使用されています")
    ErrDuplicateEmail       = errors.New("メールアドレスが既に使用されています")
    ErrInvalidClerkID       = errors.New("無効なClerk IDです")
    ErrInvalidSocialLinkURL = errors.New("無効なソーシャルリンクURLです")
    ErrSelfRival            = errors.New("自分自身をライバルに設定できません")
    ErrDuplicateRival       = errors.New("既にライバル関係が存在します")
)
```

### 3. テスト戦略
- **エンティティテスト**: バリデーション、ビジネスルールの検証
- **値オブジェクトテスト**: 値の妥当性、変換処理の検証
- **ドメインサービステスト**: 複雑なロジックの検証

## 依存関係
- **外部依存なし**: Domain層は完全に独立
- **標準ライブラリのみ**: time、context、errors等
- **UUID生成**: google/uuid ライブラリ

## 工数見積詳細

| 実装項目 | 工数 | 備考 |
|---------|------|------|
| エンティティ実装（4つ） | 2時間 | バリデーション含む |
| 値オブジェクト実装（3つ） | 0.5時間 | 定数定義、バリデーション |
| リポジトリI/F実装（4つ） | 1時間 | インターフェース定義 |
| ドメインサービス実装 | 1時間 | ビジネスロジック |
| エラー定義・テスト | 0.5時間 | エラー定数、簡易テスト |
| **合計** | **4時間** | |

## 完了条件
1. ✅ 全エンティティが実装され、バリデーションが動作する
2. ✅ 全値オブジェクトが実装され、定数が定義されている
3. ✅ 全リポジトリインターフェースが定義されている
4. ✅ ドメインサービスが実装され、ビジネスルールが適用される
5. ✅ ドメインエラーが適切に定義されている
6. ✅ Go modulesでコンパイルが通る
7. ✅ 基本的なユニットテストが実装されている

## 注意事項
- Infrastructure層の実装は行わない（次のタスクで実装）
- 外部ライブラリへの依存は最小限に抑える
- ビジネスロジックはDomain層に集約する
- TODO コメントで次フェーズでの実装予定を記載する