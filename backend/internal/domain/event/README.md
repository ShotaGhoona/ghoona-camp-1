# Event Domain - イベント管理ドメイン

朝活コミュニティアプリのイベント管理機能を担当するドメイン層。

## Value Objects (値オブジェクト)

- **EventType**: イベント種別 (general, study, exercise, meditation, creative, business)
- **ParticipantStatus**: 参加状態 (registered, cancelled) 
- **RecurrencePattern**: 繰り返しパターン (none, daily, weekly, monthly)
- **EventStatus**: イベント進行状態 (upcoming, ongoing, completed, cancelled)
- **TimeSlot**: 開始・終了時刻ペア（時間前後関係バリデーション付き）

## Entities (エンティティ)

### Event
**型定義**: ID, CreatorID, Title, Description, EventType, 日時情報, MaxParticipants, 繰り返し設定, Discord連携, Active状態

**ファクトリ**: 
- `NewEvent()` - 基本フィールド設定とUUID生成

**ドメインメソッド**:
- **状態系**: `Status()`, `IsUpcoming()`, `IsOngoing()`, `IsCompleted()` - 時刻ベース状態判定
- **ライフサイクル**: `Cancel()`, `Reschedule()` - キャンセル・日時変更
- **定員管理**: `CanAcceptRegistration()`, `ChangeCapacity()` - 参加受付可否・定員変更
- **情報更新**: `UpdateBasicInfo()`, `AssignDiscord()`, `ClearDiscord()` - 基本情報・Discord設定

### EventParticipant
**型定義**: ID, EventID, UserID, Status, 作成・更新日時

**ファクトリ**:
- `NewEventParticipant()` - デフォルトregistered状態で作成

**ドメインメソッド**:
- **状態変更**: `Register()`, `Cancel()`, `UpdateStatus()` - 参加状態の一元管理
- **状態判定**: `IsRegistered()`, `IsCancelled()` - 読みやすい状態確認

## Repository Interfaces (リポジトリ)

### EventRepository (8メソッド)
- **基本CRUD**: FindByID, Create, Update, Delete
- **検索**: FindAll(フィルタ), FindByCreatorID, FindByDateRange
- **集計**: CountParticipants

### EventParticipantRepository (6メソッド)  
- **検索**: FindByEventID, FindByEventAndUser
- **操作**: Create, Update
- **集計**: CountRegisteredParticipants, ExistsByEventAndUser

## Errors (ドメインエラー)

- **イベント関連**: NotFound, AlreadyExists, Cancelled, Completed, Inactive
- **作成・更新**: InvalidTitle, TitleTooLong, InvalidEventType, InvalidTimeSlot, InvalidCapacity  
- **参加者**: ParticipantNotFound, AlreadyExists, EventFull, RegistrationClosed
- **権限**: Unauthorized, NotEventCreator

## 使用目的

API実装時のビジネスロジック保護・データ整合性確保・エラーハンドリング統一のためのドメイン知識集約層。