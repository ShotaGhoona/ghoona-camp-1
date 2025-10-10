# Domain Layer Implementation Summary

## リビジョン履歴

| バージョン | 日付 | 変更内容 | 担当者 |
|-----------|------|---------|--------|
| v1.2 | 2025-01-21 | ドメインエラーハンドリング実装 - カスタムエラー型とリファクタリング | Claude |
| v1.1 | 2025-01-21 | Domain Services追加 - 必要最小限のビジネスロジック実装 | Claude |
| v1.0 | 2025-01-21 | 初版作成 - ドメイン層完全実装完了 | Claude |

## 概要

### 実装目的
Ghoona Camp朝活コミュニティアプリのバックエンドにおいて、DDD (Domain-Driven Design) × Onion Architectureに基づくドメイン層を完全実装。ビジネスロジックとデータアクセスを明確に分離し、保守性と拡張性を確保する。

### 対象範囲
- 6ドメイン（User, Goal, Event, Title, Attendance, Notification）
- 22個のValue Objects（制約実装）
- 15個のEntities（ビジネスロジック）
- 15個のRepository Interfaces（データアクセス抽象化）
- 2個のDomain Services（複雑なビジネスルール管理）
- 6個のDomain Error Files（カスタムエラーハンドリング）

## 実装内容

### 📁 ディレクトリ構造
```
backend/internal/domain/
├── user/        (認証・プロフィール管理)
├── goal/        (目標設定・追跡)
├── event/       (朝活イベント管理)
├── title/       (称号・ゲーミフィケーション)
├── attendance/  (Discord参加記録)
└── notification/ (通知・リマインダー)

各ドメイン内:
├── entity/     (ビジネスロジック)
├── repository/ (データアクセス抽象化)
├── vo/         (制約・ルール)
├── service/    (複合ビジネスロジック) ※User・Eventのみ
└── error.go    (ドメイン固有エラー定義)
```

### 🔒 主要なValue Objects制約

#### 称号システム (TitleLevel)
```go
// 8段階の成長ストーリー
Level 1: 1日    → "まどろみ見習い"
Level 2: 7日    → "早起き候補生"  
Level 3: 30日   → "朝活探検家"
Level 4: 100日  → "朝の住人"
Level 5: 200日  → "夜明けの戦士"
Level 6: 365日  → "朝光の使者"
Level 7: 500日  → "暁の守護者"
Level 8: 1000日 → "朝活の伝説"
```

#### 参加時間制約 (DurationMinutes)
```go
// 現実的な時間制限
MinDuration: 0分
MaxDuration: 1440分 (24時間)
// 異常値を排除し、正確な統計を保証
```

#### プラットフォーム標準化 (Platform)
```go
// 10種類の主要SNSに限定
twitter, instagram, github, linkedin, 
facebook, youtube, website, blog, 
portfolio, other
// UI/UX統一と開発効率化
```

### 🏗️ Entity設計パターン

#### 時間整合性チェック
```go
// Event Entity例
if !endTime.After(startTime) {
    return nil, errors.New("終了時間は開始時間より後に設定してください")
}
```

#### 必須フィールド検証
```go
// User Entity例
if clerkID == "" {
    return nil, errors.New("Clerk IDは必須です")
}
if email == "" {
    return nil, errors.New("メールアドレスは必須です")
}
```

### 🔄 Repository Interface設計

#### 標準的なCRUD操作
```go
type UserRepository interface {
    Create(ctx context.Context, user *entity.User) error
    GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
    Update(ctx context.Context, user *entity.User) error
    Delete(ctx context.Context, id uuid.UUID) error
}
```

#### ドメイン固有メソッド
```go
// ライバル関係
CountByUserID(ctx context.Context, userID uuid.UUID) (int, error)

// 統計・ランキング
GetRankingByTotalDays(ctx context.Context, limit, offset int) ([]*entity.AttendanceStatistics, error)

// 検索・フィルタ
SearchByUserID(ctx context.Context, userID uuid.UUID, query string, limit, offset int) ([]*entity.Goal, error)
```

### 🔧 Domain Services設計

#### RivalManagementService (User Domain)
```go
// ライバル管理のビジネスルール
- 自分自身をライバルに設定することの禁止
- ライバル登録の3人制限チェック
- ユーザーの重複ライバル登録防止

// 主要メソッド
CanAddRival(ctx context.Context, userID, rivalUserID uuid.UUID) error
AddRival(ctx context.Context, userID, rivalUserID uuid.UUID) (*entity.UserRival, error)
```

#### EventParticipationService (Event Domain)
```go
// イベント参加のビジネスルール
- イベント参加者数の上限チェック（max_participantsに対して）
- 同一ユーザーの重複参加防止
- キャンセル済み参加者の再参加対応

// 主要メソッド
CanJoinEvent(ctx context.Context, eventID, userID uuid.UUID) error
JoinEvent(ctx context.Context, eventID, userID uuid.UUID) (*entity.EventParticipant, error)
```

### 🚨 ドメインエラーハンドリング

#### カスタムエラー型の設計
```go
// User Domain例
var (
    ErrCannotAddSelfAsRival    = errors.New("自分自身をライバルに設定することはできません")
    ErrMaxRivalsLimitExceeded  = errors.New("ライバルは最大3人までしか設定できません")
    ErrRivalAlreadyExists      = errors.New("このユーザーは既にライバルに設定されています")
)
```

#### エラーメッセージの国際化対応
- **日本語メッセージ**: ユーザー向けエラーメッセージは日本語で統一
- **定数化**: エラーメッセージの変更時の影響範囲を限定
- **ドメイン分離**: 各ドメインのerror.goで独立したエラー管理

#### Domain Servicesでの活用
```go
// Before: インラインエラー
return errors.New("自分自身をライバルに設定することはできません")

// After: ドメインエラー使用
return user.ErrCannotAddSelfAsRival
```

## 共通パターン・ガイドライン

### Value Object実装パターン
```go
// 1. 制約定義
const (
    MinValue = 0
    MaxValue = 1000
)

// 2. バリデーション付きコンストラクタ
func NewValueObject(value int) (ValueObject, error) {
    if value < MinValue || value > MaxValue {
        return 0, errors.New("範囲外の値です")
    }
    return ValueObject(value), nil
}

// 3. 値取得メソッド
func (v ValueObject) Value() int {
    return int(v)
}
```

### Entity設計パターン
```go
// 1. 不変フィールド（ID, 作成日時）
id        uuid.UUID
createdAt time.Time

// 2. バリデーション付きコンストラクタ
func NewEntity(...) (*Entity, error) {
    // ビジネスルール検証
    // UUID生成、時刻設定
    return &Entity{...}, nil
}

// 3. ゲッターのみ（エンカプセレーション）
func (e *Entity) ID() uuid.UUID { return e.id }
```

### Repository Interface設計原則
- **Context必須**: すべてのメソッドで`context.Context`を第一引数
- **エラーハンドリング**: 戻り値の最後は必ず`error`
- **ページネーション**: `limit, offset int`で統一
- **型安全性**: UUIDは`uuid.UUID`型、enumは専用VO型使用

### Domain Service設計パターン
```go
// 1. 単一責任の原則
type SomeService struct {
    specificRepo repository.SpecificRepository  // 必要最小限の依存
}

// 2. ビジネスルールチェックメソッド
func (s *SomeService) CanPerformAction(ctx context.Context, params...) error {
    // 複数の制約をチェック
    // エラーメッセージは日本語で具体的に
}

// 3. アクション実行メソッド（チェック付き）
func (s *SomeService) PerformAction(ctx context.Context, params...) (*entity.Result, error) {
    // 1. ビジネスルールチェック
    if err := s.CanPerformAction(ctx, params...); err != nil {
        return nil, err
    }
    // 2. エンティティ作成・操作
    // 3. リポジトリ保存
}
```

## 技術的成果・効果

### 🎯 ビジネス価値の実現
- **継続促進**: 段階的称号システムによるモチベーション向上
- **データ品質**: 型安全性による実行時エラー排除
- **開発効率**: 一貫したパターンによる学習コスト削減
- **国際化対応**: 日英バイリンガル称号システム

### 🛡️ 技術的優位性
- **保守性**: YAGNI原則による適切な複雑度
- **拡張性**: Interface分離による段階的機能追加
- **テスタビリティ**: Repository抽象化による単体テスト容易性
- **パフォーマンス**: Value Objectによる早期バリデーション
- **ビジネスルール集約**: Domain Serviceによる制約の一元管理
- **エラーハンドリング**: カスタムドメインエラーによる一貫性と保守性向上

## 備考・注意事項

### Infrastructure層での実装時の注意点
1. **Repository実装**: Infrastructure層でPostgreSQL/Supabaseとの接続実装が必要
2. **トランザクション管理**: 複数Repository操作時の整合性保証
3. **エラーハンドリング**: ドメインエラーとインフラエラーの適切な変換
4. **パフォーマンス**: N+1問題の回避とクエリ最適化

### 今後の拡張ポイント
- **追加ドメインサービス**: 称号判定、ランキング計算、ライバル推薦等
- **ドメインイベント**: 非同期処理（通知送信、統計更新等）
- **仕様パターン**: 複雑な検索・フィルタ条件
- **集約ルート**: トランザクション境界の明確化

### 実装済みDomain Servicesの活用
- **Application層**: UseCaseでDomain Serviceを呼び出し
- **単体テスト**: Repository Mockを使用したビジネスロジックテスト
- **エラーハンドリング**: 日本語メッセージをそのままAPI応答に利用可能

### ドメインエラーの運用指針
- **一貫性**: 全ドメインで統一されたエラーハンドリングパターン
- **保守性**: エラーメッセージ変更時の影響範囲の局所化
- **可読性**: ドメイン固有の意味を持つエラー名称
- **国際化**: 日本語メッセージによるユーザビリティ向上

このドメイン層実装により、Ghoona Campアプリケーションは「技術的に堅牢で、ビジネス要求に忠実で、エラーハンドリングが一貫した」基盤を獲得しました。