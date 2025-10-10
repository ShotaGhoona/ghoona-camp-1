# Domain & Repository Implementation Strategy

## 実装方針

### アーキテクチャ
- **DDD × Onion Architecture**
- **Repository Pattern**でデータアクセス抽象化
- **Value Objects**による制約実装

### 実装順序
1. **Value Objects** (22個) - ビジネス制約の実装
2. **Entities** (15個) - ドメインロジックの実装  
3. **Repository Interfaces** - データアクセス抽象化

## 6ドメインの責務

### User Domain
- 認証基盤、プロフィール管理、SNS連携、ライバル機能

### Goal Domain  
- 個人目標設定・追跡、公開/非公開制御

### Event Domain
- 朝活イベント管理、参加者管理、Discord連携

### Title Domain
- 称号システム、ゲーミフィケーション、実績管理

### Attendance Domain
- Discord参加記録、統計計算、連続日数管理

### Notification Domain
- 通知配信、ユーザー設定、リマインダー機能

## Repository設計原則

### インターフェース定義
```go
type UserRepository interface {
    Create(ctx context.Context, user *entity.User) error
    GetByID(ctx context.Context, id string) (*entity.User, error)
    Update(ctx context.Context, user *entity.User) error
    Delete(ctx context.Context, id string) error
}
```

### 実装分離
- **Domain層**: インターフェース定義のみ
- **Infrastructure層**: 具体的な実装(Supabase)

## 品質保証

### Value Objects制約
- 型安全性によるコンパイル時チェック
- ビジネスルール違反の防止

### Entity設計
- 最小限のコンストラクタ
- 不変性の保証
- ビジネスロジックの分離

## 実装完了後の効果

- **データ整合性**: 型レベルでの制約保証
- **保守性**: YAGNI原則による適切な複雑度
- **拡張性**: 段階的機能追加が可能
- **テスタビリティ**: Repository抽象化による単体テスト容易性