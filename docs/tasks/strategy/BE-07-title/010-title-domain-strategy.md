# Title Domain 実装戦略書

## 概要
称号管理機能のDomain層実装戦略書です。8段階の称号システムとユーザーの称号獲得・表示管理を実装します。

## 要件整理

### 称号システムの仕様
- **8段階の称号レベル**: レベル1（1日）〜レベル8（365日）
- **参加日数ベース**: `attendance_statistics.total_attendance_days`に基づく自動獲得
- **表示称号管理**: 獲得済み称号から1つのみ選択可能
- **進捗計算**: 未獲得称号への進捗情報提供

### データベース設計
- `titles`: 称号マスターデータ（8件、シードデータ投入済み）
- `title_achievements`: ユーザー称号獲得実績

## Domain層の設計方針

### アーキテクチャ原則
1. **Domain層の独立性**: 他ドメインへの直接依存を避ける
2. **既存コードとの統一性**: userドメインのパターンに準拠
3. **オーバーエンジニアリング回避**: 要件に必要な機能のみ実装

### 依存関係の解決
```
出席統計データの取得:
Application層 → AttendanceService (統計取得)
             → TitleService (称号判定, 統計データを引数で受け取り)
```

## 実装コンポーネント

### 1. Entity層

#### `entity/title.go`
```go
type Title struct {
    ID            uuid.UUID
    Level         int           // 1-8
    NameJP        string        // 日本語名
    NameEN        string        // 英語名  
    Description   string        // 説明
    RequiredDays  int          // 必要参加日数
    ImageURL      *string      // 画像URL（オプショナル）
    ColorTheme    *string      // テーマカラー（オプショナル）
    IsActive      value.ActiveFlag // 有効状態
    CreatedAt     time.Time
    UpdatedAt     time.Time
}
```

#### `entity/title_achievement.go`
```go
type TitleAchievement struct {
    ID          uuid.UUID
    UserID      uuid.UUID
    TitleID     uuid.UUID
    AchievedAt  time.Time
    IsCurrent   value.PublicFlag // 現在表示中フラグ
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### 2. Value Object層

#### `value/active_flag.go`
```go
type ActiveFlag bool
const (
    ActiveTrue  ActiveFlag = true
    ActiveFalse ActiveFlag = false
)
```

#### `value/attendance_statistics.go`
```go
// 他ドメインとの結合を避けるための値オブジェクト
type AttendanceStatistics struct {
    TotalAttendanceDays int
}
```

### 3. Repository層

#### `repository/title_repository.go`
```go
type TitleRepository interface {
    GetAll(ctx context.Context) ([]*entity.Title, error)
    GetByID(ctx context.Context, id uuid.UUID) (*entity.Title, error)
    GetByLevel(ctx context.Context, level int) (*entity.Title, error)
    GetActiveTitles(ctx context.Context) ([]*entity.Title, error)
}
```

#### `repository/title_achievement_repository.go`
```go
type TitleAchievementRepository interface {
    GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.TitleAchievement, error)
    GetCurrentByUserID(ctx context.Context, userID uuid.UUID) (*entity.TitleAchievement, error)
    Create(ctx context.Context, achievement *entity.TitleAchievement) error
    Update(ctx context.Context, achievement *entity.TitleAchievement) error
    SetCurrent(ctx context.Context, userID, titleID uuid.UUID) error
}
```

### 4. Service層

#### `service/title_service.go`
```go
type TitleService struct{}

// 獲得可能な称号をチェック
func (s *TitleService) CheckEligibleTitles(stats value.AttendanceStatistics, allTitles []*entity.Title) []*entity.Title

// 次の称号への進捗計算
func (s *TitleService) CalculateProgress(stats value.AttendanceStatistics, currentLevel int, nextTitle *entity.Title) int

// 称号表示変更の妥当性チェック
func (s *TitleService) ValidateCurrentTitleChange(userAchievements []*entity.TitleAchievement, targetTitleID uuid.UUID) error
```

### 5. Error定義

#### `errors.go`
```go
var (
    // Title errors
    ErrTitleNotFound        = errors.New("称号が見つかりません")
    ErrTitleInactive        = errors.New("非アクティブな称号です")
    ErrInvalidTitleLevel    = errors.New("無効な称号レベルです")
    
    // Achievement errors  
    ErrAchievementNotFound     = errors.New("称号獲得記録が見つかりません")
    ErrTitleNotAchieved       = errors.New("未獲得の称号です")
    ErrTitleAlreadyCurrent    = errors.New("既に現在の称号に設定済み")
    ErrAchievementAlreadyExists = errors.New("称号実績が既に存在します")
)
```

## 既存コードとの統一性

### パターンの踏襲
1. **Entityコンストラクタ**: `NewTitle()`, `NewTitleAchievement()`
2. **Value Object**: `IsValid()`, `String()`メソッド
3. **Repository**: 標準的なCRUD + ドメイン固有メソッド
4. **Service**: バリデーション + ビジネスロジック
5. **Error**: 日本語メッセージ、明確な分類

### コーディング規約
- パッケージ構造: `domain/title/{entity,value,repository,service}`
- 命名規則: userドメインと同様のパターン
- インターフェース設計: 最小限かつ明確な責務分離

## 実装優先順位

### Phase 1: 基本構造 (2時間)
1. `value/active_flag.go`
2. `value/attendance_statistics.go` 
3. `entity/title.go`
4. `entity/title_achievement.go`

### Phase 2: Repository (1.5時間)
1. `repository/title_repository.go`
2. `repository/title_achievement_repository.go`

### Phase 3: Business Logic (1.5時間)
1. `service/title_service.go`
2. `errors.go`

## 注意点

### オーバーエンジニアリング回避
- 複雑な称号条件（連続日数等）は将来拡張として除外
- シンプルな参加日数ベースのみ実装
- 不要な抽象化を避ける

### テスタビリティ
- 外部依存（attendance統計）を引数で受け取る設計
- 純粋関数を意識したService実装
- モックしやすいRepository設計

### 拡張性考慮
- 将来的な称号条件追加に対応可能な設計
- レベル上限（8段階）の変更に柔軟対応
- 新しい称号タイプ追加への備え

---

この戦略に基づいて、userドメインと統一性を保ちながら、要件を満たす最小限の実装を行います。