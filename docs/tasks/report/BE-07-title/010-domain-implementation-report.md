# Title Domain 実装レポート

## 概要
称号管理機能のDomain層実装が完了しました。8段階の称号システムとユーザーの称号獲得・表示管理を実装し、既存のuserドメインと統一感を保ちながら、commonパッケージを適切に活用した設計となっています。

## 実装日時
- 開始: 2025-01-21
- 完了: 2025-01-21
- 所要時間: 約1.5時間

## 実装コンポーネント

### 1. Value Objects

#### `value/active_flag.go`
- **役割**: 称号のアクティブ状態管理
- **特徴**: userドメインとの独立性を保持、bool型の型安全な実装
- **メソッド**: `String()`, `Bool()`, `IsValid()`

#### `value/current_flag.go`
- **役割**: 称号の現在表示状態管理
- **特徴**: 意味を明確化、現在表示中/非表示の2状態
- **メソッド**: `String()`, `Bool()`

#### `value/attendance_statistics.go`
- **役割**: 称号判定に必要な出席統計情報
- **特徴**: 他ドメインとの疎結合を実現、防御的プログラミング
- **メソッド**: `NewAttendanceStatistics()`, `IsValid()`

### 2. Entities

#### `entity/title.go`
- **役割**: 称号の基本情報管理
- **特徴**: 
  - `common.BaseEntity`を継承（ID、CreatedAt、UpdatedAt）
  - 8段階レベル管理（1-8）
  - 参加日数ベースの獲得条件判定
- **主要メソッド**: 
  - `NewTitle()`: エンティティ作成
  - `IsEligibleFor()`: 獲得可能性判定
  - `IsValidLevel()`: レベル妥当性チェック

#### `entity/title_achievement.go`
- **役割**: ユーザーの称号獲得実績管理
- **特徴**: 
  - `common.BaseEntity`を継承
  - 現在表示称号の排他制御
  - タイムスタンプ自動管理
- **主要メソッド**: 
  - `NewTitleAchievement()`: 実績作成
  - `SetAsCurrent()`: 現在表示設定
  - `UnsetCurrent()`: 表示解除
  - `IsCurrentTitle()`: 表示状態確認

### 3. Repositories

#### `repository/title_repository.go`
- **役割**: 称号データアクセスインターフェース
- **メソッド**: 
  - `GetAll()`: 全称号取得
  - `GetByID()`: ID指定取得
  - `GetByLevel()`: レベル指定取得
  - `GetActiveTitles()`: アクティブ称号取得

#### `repository/title_achievement_repository.go`
- **役割**: 称号獲得実績データアクセスインターフェース
- **メソッド**: 
  - `GetByUserID()`: ユーザー実績一覧
  - `GetCurrentByUserID()`: 現在表示称号
  - `GetByUserIDAndTitleID()`: 特定実績取得
  - `Create()`, `Update()`: CRUD操作
  - `SetCurrent()`: 現在表示設定（排他制御）

### 4. Services

#### `service/title_service.go`
- **役割**: 称号関連ビジネスロジック処理
- **主要機能**: 
  - 称号獲得可能性判定
  - 進捗計算（パーセンテージ、残り日数）
  - 最高獲得レベル算出
  - 表示変更妥当性チェック
- **主要メソッド**: 
  - `CheckEligibleTitles()`: 獲得可能称号判定
  - `GetNextTitle()`: 次獲得称号取得
  - `CalculateProgress()`: 進捗計算
  - `GetHighestAchievedLevel()`: 最高レベル取得
  - `ValidateCurrentTitleChange()`: 表示変更バリデーション
  - `GetRemainingDaysToNextTitle()`: 残り日数計算

### 5. Error定義

#### `errors.go`
- **役割**: ドメイン固有エラー定義
- **特徴**: `common.DomainError`を活用した構造化エラー
- **エラー種別**: 
  - Title errors: `ErrTitleNotFound`, `ErrTitleInactive`, `ErrInvalidTitleLevel`
  - Achievement errors: `ErrAchievementNotFound`, `ErrTitleNotAchieved`, `ErrTitleAlreadyCurrent`, `ErrAchievementAlreadyExists`

## 設計原則の遵守

### 1. DDD（Domain-Driven Design）
- ✅ **ユビキタス言語**: 称号、獲得、レベル等の業務用語を一貫使用
- ✅ **エンティティ**: IDによる一意性、ライフサイクル管理
- ✅ **値オブジェクト**: 不変性、等価性による値の表現
- ✅ **ドメインサービス**: 複数エンティティに跨るビジネスロジック

### 2. オニオンアーキテクチャ
- ✅ **Domain層の独立性**: 他ドメインへの依存を回避
- ✅ **依存方向**: 外部依存を排除、純粋なビジネスロジック
- ✅ **インターフェース分離**: Repository等の抽象化

### 3. 既存コードとの統一性
- ✅ **命名規則**: userドメインと統一したパターン
- ✅ **パッケージ構造**: `entity/value/repository/service`の標準構造
- ✅ **エラーハンドリング**: 日本語メッセージ、明確な分類

## commonパッケージの活用

### 活用コンポーネント
1. **`BaseEntity`**: ID、CreatedAt、UpdatedAtの共通化
2. **`UUID`**: 型安全なUUID使用
3. **`NewUUID()`**: 標準的なUUID生成
4. **`UpdateTimestamp()`**: 更新時刻の統一処理
5. **共通エラー**: `ErrNotFound`, `ErrAlreadyExists`の再利用
6. **`DomainError`**: 構造化エラーによる詳細情報管理

### 効果
- コード重複の削減
- 型安全性の向上
- エラーハンドリングの統一
- 保守性の向上

## 他ドメインとの疎結合

### 課題と解決
**課題**: 称号獲得判定には出席統計データが必要
**解決**: `AttendanceStatistics`値オブジェクトによる疎結合
- Application層で依存関係を調整
- titleドメインは統計データを引数で受け取り
- テスタビリティの向上

### 依存関係図
```
Application層:
  AttendanceService → 統計取得
  TitleService → 称号判定（統計データを引数）
  
Domain層:
  titleドメイン（独立）
```

## API要件への対応

### 実装済み機能
1. ✅ **称号一覧取得**: 8段階の称号情報提供
2. ✅ **称号詳細取得**: ID指定での詳細情報取得
3. ✅ **ユーザー実績管理**: 獲得履歴、現在表示称号
4. ✅ **進捗計算**: 次称号への進捗、残り日数
5. ✅ **表示変更**: 獲得済み称号からの選択、妥当性チェック

### 将来対応
- バッチ処理による自動獲得判定
- 通知システムとの連携
- ランキング機能

## テスト容易性

### 設計上の配慮
1. **純粋関数**: 外部依存のないビジネスロジック
2. **依存注入**: Repository、外部データの注入可能
3. **モック対応**: インターフェース設計による疎結合
4. **値オブジェクト**: 不変性による予測可能な動作

## パフォーマンス考慮

### 最適化ポイント
1. **メモリ効率**: 不要なポインタ使用を避ける設計
2. **計算最適化**: 進捗計算の効率的アルゴリズム
3. **データ構造**: Map使用による高速検索

## セキュリティ考慮

### 実装されたセキュリティ機能
1. **入力検証**: 不正レベル、負数の防御
2. **ビジネスルール**: 未獲得称号の表示変更防止
3. **データ整合性**: 排他制御による一貫性確保

## 今後の課題

### 技術的課題
1. **Infrastructure層**: GORM実装、マイグレーション
2. **Application層**: ユースケース、DTO実装
3. **Interface層**: REST API、認証連携

### 機能的課題
1. **バッチ処理**: 自動獲得判定の実装
2. **通知連携**: 称号獲得時の通知システム
3. **キャッシュ**: 称号データの高速アクセス

## 総括

titleドメインのDomain層実装は、DDDとオニオンアーキテクチャの原則に従い、既存コードとの統一性を保ちながら完成しました。commonパッケージを適切に活用し、他ドメインとの疎結合を実現することで、保守性とテスト容易性を確保しています。

次段階のInfrastructure層実装に向けて、堅牢な基盤が整備できました。

---

**実装者**: Claude Code  
**レビュー**: 実装完了  
**次のステップ**: BE-07-title-02 Infrastructure層実装