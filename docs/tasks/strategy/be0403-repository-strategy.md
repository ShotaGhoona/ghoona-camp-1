# GORM Repository Implementation Strategy

## 概要

DDD (Domain-Driven Design) × Onion Architectureに基づくインフラ層のRepository実装戦略書です。
ドメイン層で定義されたRepository Interfaceを、GORM（Go ORM）を使用してPostgreSQL/Supabaseに対して実装します。

## 実装方針

### アーキテクチャ原則
- **依存性逆転**: Infrastructure層がDomain層のインターフェースに依存
- **シンプル設計**: YAGNI原則に基づく必要最小限の実装
- **基本CRUD優先**: 複雑な機能は後回し

### 実装範囲
- **6ドメイン**: User, Goal, Event, Title, Attendance, Notification
- **15Repository**: 各ドメインのRepository Interface完全実装（1:1対応）
- **基本CRUD**: 機械的に実装可能な標準操作のみ
- **要件不足箇所**: `//TODO` コメントで明示的に未実装とする

## ディレクトリ構造

```
backend/internal/infrastructure/gorm/
├── model/                    # GORMモデル（実装済み）
│   ├── user/
│   ├── goal/
│   ├── event/
│   ├── title/
│   ├── attendance/
│   └── notification/
└── repository/               # Repository実装（これから実装）
    ├── user/
    │   ├── user_repository.go
    │   ├── user_metadata_repository.go
    │   ├── user_vision_repository.go
    │   ├── user_social_link_repository.go
    │   └── user_rival_repository.go
    ├── goal/
    │   └── goal_repository.go
    ├── event/
    │   ├── event_repository.go
    │   └── event_participant_repository.go
    ├── title/
    │   ├── title_repository.go
    │   └── title_achievement_repository.go
    ├── attendance/
    │   ├── attendance_log_repository.go
    │   ├── attendance_summary_repository.go
    │   └── attendance_statistics_repository.go
    └── notification/
        ├── notification_repository.go
        └── notification_settings_repository.go
```

## 実装パターン

### 基本Repository構造

```go
// backend/internal/infrastructure/gorm/repository/user/user_repository.go
package user

import (
    "context"
    "errors"
    
    "github.com/google/uuid"
    "gorm.io/gorm"
    
    "ghoona-camp-backend/internal/domain/user/entity"
    "ghoona-camp-backend/internal/domain/user/repository"
    userModel "ghoona-camp-backend/internal/infrastructure/gorm/model/user"
)

type userRepository struct {
    db *gorm.DB
}

// NewUserRepository コンストラクタ
func NewUserRepository(db *gorm.DB) repository.UserRepository {
    return &userRepository{db: db}
}

// Create 基本CRUD
func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
    gormUser := r.toGORMUser(user)
    return r.db.WithContext(ctx).Create(gormUser).Error
}

// GetByID 基本CRUD
func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
    var gormUser userModel.User
    err := r.db.WithContext(ctx).First(&gormUser, "id = ?", id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return r.fromGORMUser(&gormUser)
}

// 型変換（Repository内に直接実装）
func (r *userRepository) toGORMUser(user *entity.User) *userModel.User {
    return &userModel.User{
        ID:        user.ID(),
        ClerkID:   user.ClerkID(),
        Email:     user.Email(),
        // ... 他フィールド
    }
}

func (r *userRepository) fromGORMUser(gormUser *userModel.User) (*entity.User, error) {
    // TODO: entity.NewUser の引数確認後実装
    return nil, errors.New("TODO: entity constructor")
}
```

## 実装戦略

### Phase 1: 基本CRUD実装

**機械的に実装可能なメソッド**
```go
// 基本CRUD操作
Create(ctx context.Context, entity *Entity) error
GetByID(ctx context.Context, id uuid.UUID) (*Entity, error)
Update(ctx context.Context, entity *Entity) error
Delete(ctx context.Context, id uuid.UUID) error

// 単純な検索
GetByClerkID(ctx context.Context, clerkID string) (*entity.User, error)
GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Goal, error)
CountByUserID(ctx context.Context, userID uuid.UUID) (int, error)
```

**実装順序** - 既存Domain Repository Interfaceと1:1対応
```
User Domain (5 repositories):
├── user_repository.go
├── user_metadata_repository.go
├── user_rival_repository.go
├── user_social_link_repository.go
└── user_vision_repository.go

Goal Domain (1 repository):
└── goal_repository.go

Event Domain (2 repositories):
├── event_repository.go
└── event_participant_repository.go

Title Domain (2 repositories):
├── title_repository.go
└── title_achievement_repository.go

Attendance Domain (3 repositories):
├── attendance_log_repository.go
├── attendance_statistics_repository.go
└── attendance_summary_repository.go

Notification Domain (2 repositories):
├── notification_repository.go
└── notification_settings_repository.go
```

### Phase 2: 複雑検索実装（後回し）

**要件不足メソッド** - TODO実装方針

**TODOとする判定基準:**
- 検索条件の詳細が不明（`query string`の対象フィールド）
- ソート・ランキング条件が不明（ASC/DESC、期間指定）
- 集計・統計の仕様が不明（期間、単位）
- フィルタリング条件が不明

**TODO実装パターン:**
```go
// TODO: 検索仕様確定後実装
func SearchByUserID(ctx context.Context, userID uuid.UUID, query string, limit, offset int) ([]*entity.Goal, error) {
    // TODO: query引数の検索対象フィールドを確認（タイトルのみ？説明も含む？）
    return nil, errors.New("TODO: 要件確定後実装")
}

// TODO: ランキング仕様確定後実装  
func GetRankingByTotalDays(ctx context.Context, limit, offset int) ([]*entity.AttendanceStatistics, error) {
    // TODO: ソート条件（ASC/DESC）・期間指定・同順位の扱いを確認
    return nil, errors.New("TODO: 要件確定後実装")
}

// TODO: 統計仕様確定後実装
func GetStatisticsByPeriod(ctx context.Context, userID uuid.UUID, from, to time.Time) ([]*entity.AttendanceSummary, error) {
    // TODO: 集計期間・集計単位（日別/月別）を確認
    return nil, errors.New("TODO: 要件確定後実装")
}
```

**機械的に実装可能な判定基準:**
- 単純なCRUD操作
- フィールド指定での単純検索（`GetByClerkID`など）
- 単純なカウント操作
- 外部キーでの取得操作

## エラーハンドリング

**シンプルなエラー処理** - 複雑化しない
```go
func (r *repository) handleError(err error) error {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil // ドメイン仕様：存在しない場合はnil
    }
    return err // その他はそのまま返す
}
```

## 実装チェックリスト

### Phase 1 完了基準
- [ ] 全15 Repository の1:1対応実装
  - [ ] User Domain (5 repositories)
  - [ ] Goal Domain (1 repository)  
  - [ ] Event Domain (2 repositories)
  - [ ] Title Domain (2 repositories)
  - [ ] Attendance Domain (3 repositories)
  - [ ] Notification Domain (2 repositories)
- [ ] 各Repositoryの型変換関数実装（Repository内）
- [ ] 基本CRUD・単純検索の完全実装
- [ ] 要件不足メソッドは`//TODO`で明示的にエラー返却
- [ ] コンパイルエラーなし

### Phase 2 完了基準（要件確定後）
- [ ] TODO解消（検索・ランキング・統計系）
- [ ] API仕様確定後の詳細実装

## 次のステップ

1. **Domain Repository Interface確認**: 既存15ファイルの詳細メソッド把握
2. **User Repository実装**: 5つのRepository実装（基本CRUD + 型変換 + TODO明記）
3. **他Domain Repository実装**: 10つのRepository実装（同パターンで横展開）
4. **要件確認**: API仕様確定後、TODO解消

**実装原則:**
- Domain Repository Interfaceとの1:1完全対応
- 要件が不明な部分は妄想で書かず`//TODO`で明記
- YAGNI原則に基づくシンプル実装