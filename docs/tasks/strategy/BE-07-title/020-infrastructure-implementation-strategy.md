# Title Infrastructure層 実装戦略書

## 概要
称号管理機能のInfrastructure層（データアクセス層）の実装戦略を定義します。Domain層の実装を踏まえ、PostgreSQL + GORMによるデータ永続化とRepository実装を行います。

## 前提条件

### Domain層実装状況
- ✅ Value Objects: `ActiveFlag`, `CurrentFlag`, `AttendanceStatistics`
- ✅ Entities: `Title`, `TitleAchievement` 
- ✅ Repositories Interface: `TitleRepository`, `TitleAchievementRepository`
- ✅ Services: `TitleService`
- ✅ Errors: 構造化エラー定義

### 既存Infrastructure基盤
- PostgreSQL + Supabase
- GORM v2
- 既存BaseRepository実装
- User domain Infrastructure実装済み

## 実装コンポーネント

### 1. GORMモデル定義

#### `model/title.go`
```go
// Title 称号基本情報のGORMモデル
type Title struct {
    ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    Level        int        `gorm:"unique;not null;check:level >= 1 AND level <= 8"`
    NameJP       string     `gorm:"not null;size:100"`
    NameEN       string     `gorm:"not null;size:100"`
    Description  string     `gorm:"not null;type:text"`
    RequiredDays int        `gorm:"not null;check:required_days >= 1"`
    ImageURL     *string    `gorm:"type:text"`
    ColorTheme   *string    `gorm:"size:50"`
    IsActive     bool       `gorm:"default:true"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

func (Title) TableName() string { return "titles" }
```

#### `model/title_achievement.go`
```go
// TitleAchievement ユーザー称号獲得実績のGORMモデル
type TitleAchievement struct {
    ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    UserID     uuid.UUID `gorm:"type:uuid;not null;index"`
    TitleID    uuid.UUID `gorm:"type:uuid;not null;index"`
    AchievedAt time.Time `gorm:"not null"`
    IsCurrent  bool      `gorm:"default:false;index"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
    
    // 外部キー制約
    Title User `gorm:"foreignKey:TitleID;references:ID"`
    User  User `gorm:"foreignKey:UserID;references:ID"`
}

func (TitleAchievement) TableName() string { return "title_achievements" }

// インデックス定義
func (TitleAchievement) Indexes() []gorm.Index {
    return []gorm.Index{
        {Fields: []string{"user_id", "title_id"}, Unique: true}, // 重複獲得防止
        {Fields: []string{"user_id", "is_current"}},             // 現在表示検索最適化
    }
}
```

### 2. エンティティ変換メソッド

#### ToEntity変換 (GORM → Domain)
```go
// Title.ToEntity
func (t *Title) ToEntity() (*entity.Title, error) {
    activeFlag := value.ActiveTrue
    if !t.IsActive {
        activeFlag = value.ActiveFalse
    }

    return &entity.Title{
        BaseEntity: common.BaseEntity{
            ID:        common.UUID(t.ID),
            CreatedAt: t.CreatedAt,
            UpdatedAt: t.UpdatedAt,
        },
        Level:        t.Level,
        NameJP:       t.NameJP,
        NameEN:       t.NameEN,
        Description:  t.Description,
        RequiredDays: t.RequiredDays,
        ImageURL:     t.ImageURL,
        ColorTheme:   t.ColorTheme,
        IsActive:     activeFlag,
    }, nil
}

// TitleAchievement.ToEntity
func (ta *TitleAchievement) ToEntity() (*entity.TitleAchievement, error) {
    currentFlag := value.CurrentFalse
    if ta.IsCurrent {
        currentFlag = value.CurrentTrue
    }

    return &entity.TitleAchievement{
        BaseEntity: common.BaseEntity{
            ID:        common.UUID(ta.ID),
            CreatedAt: ta.CreatedAt,
            UpdatedAt: ta.UpdatedAt,
        },
        UserID:     common.UUID(ta.UserID),
        TitleID:    common.UUID(ta.TitleID),
        AchievedAt: ta.AchievedAt,
        IsCurrent:  currentFlag,
    }, nil
}
```

#### FromEntity変換 (Domain → GORM)
```go
// FromEntityTitle
func FromEntityTitle(domainTitle *entity.Title) *Title {
    return &Title{
        ID:           uuid.UUID(domainTitle.ID),
        Level:        domainTitle.Level,
        NameJP:       domainTitle.NameJP,
        NameEN:       domainTitle.NameEN,
        Description:  domainTitle.Description,
        RequiredDays: domainTitle.RequiredDays,
        ImageURL:     domainTitle.ImageURL,
        ColorTheme:   domainTitle.ColorTheme,
        IsActive:     domainTitle.IsActive.Bool(),
        CreatedAt:    domainTitle.CreatedAt,
        UpdatedAt:    domainTitle.UpdatedAt,
    }
}
```

### 3. Repository実装

#### `repository/title_repository.go`
```go
type titleRepository struct {
    *baseGorm.BaseRepository
}

func NewTitleRepository(db *gorm.DB) repository.TitleRepository {
    return &titleRepository{
        BaseRepository: baseGorm.NewBaseRepository(db),
    }
}

// GetAll 全称号取得（レベル順）
func (r *titleRepository) GetAll(ctx context.Context) ([]*entity.Title, error) {
    var gormTitles []model.Title
    db := r.GetDB(ctx)
    
    err := db.Order("level ASC").Find(&gormTitles).Error
    if err != nil {
        return nil, err
    }
    
    return r.convertToEntities(gormTitles)
}

// GetActiveTitles アクティブ称号のみ取得
func (r *titleRepository) GetActiveTitles(ctx context.Context) ([]*entity.Title, error) {
    var gormTitles []model.Title
    db := r.GetDB(ctx)
    
    err := db.Where("is_active = ?", true).Order("level ASC").Find(&gormTitles).Error
    if err != nil {
        return nil, err
    }
    
    return r.convertToEntities(gormTitles)
}
```

#### `repository/title_achievement_repository.go`
```go
type titleAchievementRepository struct {
    *baseGorm.BaseRepository
}

// GetByUserID ユーザーの全獲得実績取得
func (r *titleAchievementRepository) GetByUserID(ctx context.Context, userID common.UUID) ([]*entity.TitleAchievement, error) {
    var gormAchievements []model.TitleAchievement
    db := r.GetDB(ctx)
    
    err := db.Where("user_id = ?", uuid.UUID(userID)).
             Order("achieved_at DESC").
             Find(&gormAchievements).Error
    if err != nil {
        return nil, err
    }
    
    return r.convertToEntities(gormAchievements)
}

// SetCurrent 現在表示称号設定（排他制御）
func (r *titleAchievementRepository) SetCurrent(ctx context.Context, userID common.UUID, titleID common.UUID) error {
    db := r.GetDB(ctx)
    
    return db.Transaction(func(tx *gorm.DB) error {
        // 既存の現在表示を解除
        err := tx.Model(&model.TitleAchievement{}).
                 Where("user_id = ? AND is_current = ?", uuid.UUID(userID), true).
                 Update("is_current", false).Error
        if err != nil {
            return err
        }
        
        // 指定称号を現在表示に設定
        return tx.Model(&model.TitleAchievement{}).
                 Where("user_id = ? AND title_id = ?", uuid.UUID(userID), uuid.UUID(titleID)).
                 Update("is_current", true).Error
    })
}
```

### 4. データベーススキーマ

#### マイグレーション定義
```sql
-- titles テーブル
CREATE TABLE titles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    level INTEGER UNIQUE NOT NULL CHECK (level >= 1 AND level <= 8),
    name_jp VARCHAR(100) NOT NULL,
    name_en VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    required_days INTEGER NOT NULL CHECK (required_days >= 1),
    image_url TEXT,
    color_theme VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- title_achievements テーブル
CREATE TABLE title_achievements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title_id UUID NOT NULL REFERENCES titles(id) ON DELETE CASCADE,
    achieved_at TIMESTAMP WITH TIME ZONE NOT NULL,
    is_current BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(user_id, title_id), -- 重複獲得防止
    INDEX idx_user_current (user_id, is_current),
    INDEX idx_user_id (user_id),
    INDEX idx_title_id (title_id)
);

-- 初期データ投入
INSERT INTO titles (level, name_jp, name_en, description, required_days) VALUES
(1, '新人キャンパー', 'Newcomer', '朝活の第一歩を踏み出した証', 1),
(2, '早起き習慣', 'Early Bird', '継続こそ力なり、7日間の挑戦達成', 7),
(3, '朝活戦士', 'Morning Warrior', '2週間の継続で朝活戦士の仲間入り', 14),
(4, '継続マスター', 'Consistency Master', '1ヶ月の継続で真の朝活人へ', 30),
(5, '朝活エキスパート', 'Morning Expert', '3ヶ月継続、朝活のエキスパート', 90),
(6, '朝活レジェンド', 'Morning Legend', '半年継続、伝説の朝活人', 180),
(7, '朝活マエストロ', 'Morning Maestro', '9ヶ月継続、朝活の指導者', 270),
(8, '朝活グランドマスター', 'Morning Grandmaster', '1年継続、朝活界の頂点', 365);
```

### 5. パフォーマンス最適化

#### インデックス戦略
1. **titles**: `level` (UNIQUE), `is_active` 
2. **title_achievements**: 
   - `(user_id, title_id)` UNIQUE 複合インデックス
   - `(user_id, is_current)` 現在表示検索用
   - `user_id`, `title_id` 単体インデックス

#### クエリ最適化
```go
// N+1問題対策 - JOIN使用
func (r *titleAchievementRepository) GetWithTitles(ctx context.Context, userID common.UUID) ([]*entity.TitleAchievement, error) {
    var gormAchievements []model.TitleAchievement
    db := r.GetDB(ctx)
    
    err := db.Preload("Title").
             Where("user_id = ?", uuid.UUID(userID)).
             Order("achieved_at DESC").
             Find(&gormAchievements).Error
    
    return r.convertToEntities(gormAchievements), err
}
```

### 6. エラーハンドリング

#### GORM エラー変換
```go
func (r *titleRepository) handleGORMError(err error) error {
    if err == nil {
        return nil
    }
    
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return title.ErrTitleNotFound
    }
    
    if errors.Is(err, gorm.ErrDuplicatedKey) {
        return title.ErrAchievementAlreadyExists
    }
    
    return err // その他のエラーはそのまま返す
}
```

### 7. トランザクション管理

#### 複合操作でのトランザクション
```go
func (r *titleAchievementRepository) AchieveTitle(ctx context.Context, userID common.UUID, titleID common.UUID) error {
    db := r.GetDB(ctx)
    
    return db.Transaction(func(tx *gorm.DB) error {
        // 既存獲得チェック
        var count int64
        err := tx.Model(&model.TitleAchievement{}).
                 Where("user_id = ? AND title_id = ?", uuid.UUID(userID), uuid.UUID(titleID)).
                 Count(&count).Error
        if err != nil {
            return err
        }
        
        if count > 0 {
            return title.ErrAchievementAlreadyExists
        }
        
        // 新規獲得登録
        achievement := &model.TitleAchievement{
            UserID:     uuid.UUID(userID),
            TitleID:    uuid.UUID(titleID),
            AchievedAt: time.Now(),
            IsCurrent:  false,
        }
        
        return tx.Create(achievement).Error
    })
}
```

## テスト戦略

### 1. Repository単体テスト
- SQLiteを使用したインメモリテスト
- 各Repository メソッドの動作確認
- エラーケースの検証

### 2. 統合テスト
- PostgreSQL Test Container使用
- マイグレーション実行確認
- 複合操作のトランザクション確認

### 3. パフォーマンステスト
- 大量データでのクエリ性能測定
- インデックス効果の確認
- N+1問題の検証

## 実装順序

### Phase 1: 基本CRUD実装
1. ✅ GORMモデル定義
2. ✅ 基本Repository実装
3. ✅ エンティティ変換実装
4. ✅ 単体テスト作成

### Phase 2: 高度な機能実装
1. ✅ 排他制御実装
2. ✅ トランザクション管理
3. ✅ パフォーマンス最適化
4. ✅ 統合テスト作成

### Phase 3: 本番対応
1. ✅ マイグレーション作成
2. ✅ 初期データ投入
3. ✅ エラーハンドリング強化
4. ✅ ドキュメント作成

## リスク・課題

### 技術的リスク
1. **マイグレーション失敗**: 既存データとの整合性
2. **パフォーマンス劣化**: 不適切なインデックス設計
3. **デッドロック**: 複合トランザクションでの競合

### 対策
1. **段階的マイグレーション**: バックアップとロールバック準備
2. **負荷テスト**: 本番想定でのパフォーマンス検証
3. **排他制御**: 適切なロック戦略とタイムアウト設定

## 品質保証

### コード品質
- GORMベストプラクティス遵守
- SQL インジェクション対策
- 型安全性の確保

### データ整合性
- 外部キー制約による参照整合性
- トランザクションによる ACID 特性
- バリデーション層での事前チェック

## 次ステップへの準備

### Application層への接続点
- Repository実装の抽象化確認
- エラーハンドリングの整備
- DTO変換の準備

### 運用面での準備
- ログ出力の標準化
- メトリクス収集ポイント設定
- 監視・アラート設定準備

---

**作成日**: 2025-01-21  
**対象実装**: BE-07-title-02  
**前段階**: BE-07-title-01 (Domain層実装完了)  
**次段階**: BE-07-title-03 (Application層実装)