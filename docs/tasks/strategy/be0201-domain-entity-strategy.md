# Backend Domain Layer Implementation Strategy

## Overview
この戦略書では、Ghoona Campのドメイン層（Value Objects & Entities）の実装方針を定義します。オーバーエンジニアリングを避けつつ、要件を漏れなく実装することを目標とします。

## Architecture Strategy

### Directory Structure
```
backend/internal/domain/
├── attendance/
│   ├── entity/
│   │   ├── attendance_log.go       # 参加ログエンティティ
│   │   ├── attendance_summary.go   # 参加サマリーエンティティ
│   │   └── attendance_statistics.go # 参加統計エンティティ
│   └── vo/
│       ├── duration_minutes.go     # 参加時間(分)
│       ├── session_count.go        # セッション数
│       └── streak_days.go          # 連続参加日数
├── event/
│   ├── entity/
│   │   ├── event.go                # イベントエンティティ
│   │   └── event_participant.go    # イベント参加者エンティティ
│   └── vo/
│       ├── event_type.go           # イベントタイプ
│       ├── participant_status.go   # 参加者状態
│       ├── recurrence_pattern.go   # 繰り返しパターン
│       └── active_flag.go          # 有効/無効フラグ
├── goal/
│   ├── entity/
│   │   └── goal.go                 # 目標エンティティ
│   └── vo/
│       ├── active_flag.go          # 有効/無効フラグ
│       └── public_flag.go          # 公開フラグ
├── notification/
│   ├── entity/
│   │   ├── notification.go         # 通知エンティティ
│   │   └── notification_settings.go # 通知設定エンティティ
│   └── vo/
│       ├── notification_type.go    # 通知タイプ
│       └── read_flag.go            # 既読フラグ
├── title/
│   ├── entity/
│   │   ├── title.go                # 称号エンティティ
│   │   └── title_achievement.go    # 称号実績エンティティ
│   └── vo/
│       ├── title_level.go          # 称号レベル (1-8)
│       ├── active_flag.go          # 有効/無効フラグ
│       └── current_flag.go         # 現在設定フラグ
└── user/
    ├── entity/
    │   ├── user.go                 # ユーザーエンティティ
    │   ├── user_metadata.go        # ユーザーメタデータエンティティ
    │   ├── user_vision.go          # ユーザービジョンエンティティ
    │   ├── user_social_link.go     # ソーシャルリンクエンティティ
    │   └── user_rival.go           # ライバル関係エンティティ
    └── vo/
        ├── user_status.go          # ユーザー状態 (active/inactive)
        ├── platform.go             # プラットフォーム種別
        └── public_flag.go          # 公開フラグ
```

## Value Objects Implementation Strategy

### 1. Enum-like Value Objects

#### 1.1 UserStatus (user/vo/user_status.go)
```go
type UserStatus string

const (
    UserStatusActive   UserStatus = "active"
    UserStatusInactive UserStatus = "inactive"
)
```
**理由**: データベースのVARCHAR制約と一致、ビジネスルールで明確に定義済み

#### 1.2 Platform (user/vo/platform.go)
```go
type Platform string

const (
    PlatformTwitter    Platform = "twitter"
    PlatformInstagram  Platform = "instagram"
    PlatformGithub     Platform = "github"
    PlatformLinkedin   Platform = "linkedin"
    PlatformFacebook   Platform = "facebook"
    PlatformYoutube    Platform = "youtube"
    PlatformWebsite    Platform = "website"
    PlatformBlog       Platform = "blog"
    PlatformPortfolio  Platform = "portfolio"
    PlatformOther      Platform = "other"
)
```
**理由**: APIで明確に定義、UIでの選択肢制限が必要

#### 1.3 EventType (event/vo/event_type.go)
```go
type EventType string

const (
    EventTypeGeneral      EventType = "general"
    EventTypeStudy        EventType = "study"
    EventTypeExercise     EventType = "exercise"
    EventTypeMeditation   EventType = "meditation"
    EventTypeDiscussion   EventType = "discussion"
    EventTypePresentation EventType = "presentation"
    EventTypeWorkshop     EventType = "workshop"
)
```
**理由**: データベースとAPIで明確に定義、イベント分類の中核

#### 1.4 ParticipantStatus (event/vo/participant_status.go)
```go
type ParticipantStatus string

const (
    ParticipantStatusRegistered ParticipantStatus = "registered"
    ParticipantStatusCancelled  ParticipantStatus = "cancelled"
)
```
**理由**: 参加状態の制御、キャンセル後の再登録可能というビジネスルール

#### 1.5 RecurrencePattern (event/vo/recurrence_pattern.go)
```go
type RecurrencePattern string

const (
    RecurrencePatternDaily  RecurrencePattern = "daily"
    RecurrencePatternWeekly RecurrencePattern = "weekly"
)
```
**理由**: 定期開催イベントの繰り返しパターン制御

#### 1.6 NotificationType (notification/vo/notification_type.go)
```go
type NotificationType string

const (
    NotificationTypeAchievement NotificationType = "achievement"
    NotificationTypeReminder    NotificationType = "reminder"
    NotificationTypeRivalUpdate NotificationType = "rival_update"
    NotificationTypeEvent       NotificationType = "event"
)
```
**理由**: 通知の種別管理、ユーザーの通知設定制御

#### 1.7 TitleLevel (title/vo/title_level.go)
```go
type TitleLevel int

const (
    TitleLevel1 TitleLevel = 1  // まどろみ見習い - 1日
    TitleLevel2 TitleLevel = 2  // 早起き候補生 - 7日
    TitleLevel3 TitleLevel = 3  // 朝活探検家 - 30日
    TitleLevel4 TitleLevel = 4  // 朝の住人 - 100日
    TitleLevel5 TitleLevel = 5  // 夜明けの戦士 - 200日
    TitleLevel6 TitleLevel = 6  // 朝光の使者 - 365日
    TitleLevel7 TitleLevel = 7  // 暁の守護者 - 500日
    TitleLevel8 TitleLevel = 8  // 朝活の伝説 - 1000日
)
```
**理由**: 1-8の固定値、ビジネス上の意味が強く、称号システムの中核

### 2. Boolean Value Objects

#### 2.1 PublicFlag (user/vo/public_flag.go)
```go
type PublicFlag bool

const (
    PublicFlagPublic  PublicFlag = true
    PublicFlagPrivate PublicFlag = false
)
```
**理由**: デフォルト値の概念があり、APIでの明示的な制御が必要

#### 2.2 ActiveFlag (title/vo/active_flag.go)
```go
type ActiveFlag bool

const (
    ActiveFlagActive   ActiveFlag = true
    ActiveFlagInactive ActiveFlag = false
)
```
**理由**: エンティティの有効性制御、論理削除の概念

#### 2.3 CurrentFlag (title/vo/current_flag.go)
```go
type CurrentFlag bool

const (
    CurrentFlagCurrent    CurrentFlag = true
    CurrentFlagNotCurrent CurrentFlag = false
)
```
**理由**: ユーザーは1つの称号のみ現在設定可能というビジネスルール

#### 2.4 ReadFlag (notification/vo/read_flag.go)
```go
type ReadFlag bool

const (
    ReadFlagRead   ReadFlag = true
    ReadFlagUnread ReadFlag = false
)
```
**理由**: 通知の既読管理、未読通知数の計算等で使用

### 3. Numeric Value Objects

#### 3.1 DurationMinutes (attendance/vo/duration_minutes.go)
```go
type DurationMinutes int

const (
    MinDurationMinutes = 0
    MaxDurationMinutes = 1440 // 24時間
)
```
**理由**: 参加時間の制約、0以上1日以下の妥当性チェック

#### 3.2 SessionCount (attendance/vo/session_count.go)
```go
type SessionCount int

const (
    MinSessionCount = 0
    MaxSessionCount = 100 // 現実的な上限
)
```
**理由**: 1日のセッション数制約、異常値の検出

#### 3.3 StreakDays (attendance/vo/streak_days.go)
```go
type StreakDays int

const (
    MinStreakDays = 0
    MaxStreakDays = 10000 // 現実的な上限
)
```
**理由**: 連続参加日数の制約、称号システムとの連携

## Entity Implementation Strategy

### 1. User Domain Entities

#### 1.1 User Entity (user/entity/user.go)
**責務**: ユーザーの基本情報管理
**最低限メソッド**:
- `NewUser()` - ユーザー作成（バリデーション付き）

#### 1.2 UserMetadata Entity (user/entity/user_metadata.go)
**責務**: ユーザーのプロフィール詳細情報
**最低限メソッド**:
- `NewUserMetadata()` - メタデータ作成（バリデーション付き）

#### 1.3 UserSocialLink Entity (user/entity/user_social_link.go)
**責務**: 外部サービスとのリンク管理
**最低限メソッド**:
- `NewUserSocialLink()` - リンク作成（URL・Platform妥当性チェック付き）

#### 1.4 UserRival Entity (user/entity/user_rival.go)
**責務**: ライバル関係の管理
**最低限メソッド**:
- `NewUserRival()` - ライバル関係作成（自分自身不可チェック付き）
- **Note**: 最大3人制約はアプリケーション層で制御

### 2. Goal Domain Entities

#### 2.1 Goal Entity (goal/entity/goal.go)
**責務**: ユーザーの目標管理
**最低限メソッド**:
- `NewGoal()` - 目標作成（期間妥当性チェック: ended_at > started_at）

### 3. Event Domain Entities

#### 3.1 Event Entity (event/entity/event.go)
**責務**: 朝活イベントの管理
**最低限メソッド**:
- `NewEvent()` - イベント作成（時間制約チェック: start_time < end_time）

#### 3.2 EventParticipant Entity (event/entity/event_participant.go)
**責務**: イベント参加者の管理
**最低限メソッド**:
- `NewEventParticipant()` - 参加登録（重複チェック付き）

### 4. Attendance Domain Entities

#### 4.1 AttendanceLog Entity (attendance/entity/attendance_log.go)
**責務**: Discord参加ログの管理
**最低限メソッド**:
- `NewAttendanceLog()` - ログ作成（基本妥当性チェック付き）

#### 4.2 AttendanceSummary Entity (attendance/entity/attendance_summary.go)
**責務**: 日次参加サマリーの管理
**最低限メソッド**:
- `NewAttendanceSummary()` - サマリー作成（基本妥当性チェック付き）

#### 4.3 AttendanceStatistics Entity (attendance/entity/attendance_statistics.go)
**責務**: ユーザーの参加統計管理
**最低限メソッド**:
- `NewAttendanceStatistics()` - 統計作成（基本妥当性チェック付き）

### 5. Notification Domain Entities

#### 5.1 Notification Entity (notification/entity/notification.go)
**責務**: 個別通知の管理
**最低限メソッド**:
- `NewNotification()` - 通知作成（基本妥当性チェック付き）

#### 5.2 NotificationSettings Entity (notification/entity/notification_settings.go)
**責務**: ユーザーの通知設定管理
**最低限メソッド**:
- `NewNotificationSettings()` - 設定作成（デフォルト値設定付き）

### 6. Title Domain Entities

#### 6.1 Title Entity (title/entity/title.go)
**責務**: 称号マスター情報の管理
**最低限メソッド**:
- `NewTitle()` - 称号作成（レベル・必要日数妥当性チェック付き）

#### 6.2 TitleAchievement Entity (title/entity/title_achievement.go)
**責務**: ユーザーの称号獲得実績管理
**最低限メソッド**:
- `NewTitleAchievement()` - 実績作成（基本妥当性チェック付き）

## Implementation Rules

### 1. Value Object Rules
- **不変性**: 一度作成されたValue Objectは変更不可
- **値の等価性**: 同じ値を持つValue Objectは等価
- **バリデーション**: 作成時に値の妥当性を検証
- **GoのString型をベース**: stringやintを基本型として使用

### 2. Entity Rules (YAGNI適用)
- **一意性**: IDによる識別
- **最低限の責務**: 作成時のバリデーションのみ実装
- **メソッド方針**: 
  - 基本的にはコンストラクタ（`New*()`）のみ
  - エンティティ作成時に必要な不変条件チェックを含む
  - 更新・削除・その他のビジネスロジックはアプリケーション層で実装
- **Value Objectの活用**: 状態や分類にはValue Objectを使用

### 3. Business Rules Implementation (YAGNI適用)

#### エンティティレベル（作成時のみ）
- **自己参照制約**: UserRivalエンティティで自分自身不可チェック
- **称号レベル制約**: TitleLevelで1-8の範囲チェック
- **時間制約**: Eventエンティティでstart_time < end_timeチェック
- **期間制約**: Goalエンティティでended_at > started_atチェック

#### Value Objectレベル
- **公開設定**: PublicFlagで一貫した制御
- **アクティブ状態**: ActiveFlagで論理削除対応
- **参加状態**: ParticipantStatusでキャンセル後再登録可能
- **参加時間制約**: DurationMinutesで0-1440分の範囲チェック

#### アプリケーション層で実装
- **ライバル最大3人制約**: リポジトリで件数チェック
- **称号獲得条件**: 参加統計との連携
- **通知制御**: NotificationTypeと設定による送信制御
- **イベント定員制約**: 参加者数との比較
- **重複参加チェック**: 既存参加記録との照合

## Testing Strategy

### 1. Value Object Testing
- 有効値での作成テスト
- 無効値での作成失敗テスト
- 等価性テスト
- 文字列変換テスト

### 2. Entity Testing (YAGNI適用)
- エンティティ作成成功テスト
- 作成時不変条件違反テスト
- Value Objectとの連携テスト

## Implementation Priority

### Phase 1: Basic Value Objects
1. Common VOs: ActiveFlag, PublicFlag, ReadFlag
2. User VOs: UserStatus, Platform
3. Event VOs: EventType, ParticipantStatus, RecurrencePattern
4. Title VOs: TitleLevel, CurrentFlag
5. Notification VOs: NotificationType
6. Attendance VOs: DurationMinutes, SessionCount, StreakDays

### Phase 2: Core Domain Entities
1. User Entity (基盤となるため最優先)
2. UserMetadata Entity
3. Goal Entity
4. Title Entity

### Phase 3: Relationship Entities
1. UserSocialLink Entity
2. UserRival Entity
3. TitleAchievement Entity

### Phase 4: Activity Entities
1. Event Entity
2. EventParticipant Entity
3. Notification Entity
4. NotificationSettings Entity

### Phase 5: Analytics Entities (スコープ外だが将来実装)
1. AttendanceLog Entity
2. AttendanceSummary Entity
3. AttendanceStatistics Entity

## Future Considerations

### スコープ外だが将来的に考慮する項目
- 参加時間の最小値チェック (Discord連携時)
- タイムゾーンバリデーション
- より複雑な繰り返しパターン (月次、年次等)
- 通知配信チャンネルの拡張 (メール、プッシュ等)

### 拡張性の考慮
- 新しいプラットフォームの追加容易性
- 称号レベルの拡張可能性 (9レベル以上)
- 新しいイベントタイプの追加容易性
- 新しい通知タイプの追加容易性
- 新しいビジネスルールの追加容易性

## Domain Dependencies

### 依存関係の整理
```
User Domain (基盤)
├── Goal Domain (Userに依存)
├── Title Domain (Userに依存)
├── Event Domain (Userに依存)
├── Notification Domain (Userに依存)
└── Attendance Domain (User, Event, Titleに依存)
```

### Cross-Domain Business Rules
- **Title Achievement**: AttendanceStatisticsの参加日数に基づく
- **Rival Comparison**: AttendanceStatisticsとTitleAchievementの比較
- **Event Participation**: UserとEventの関係性
- **Notification Trigger**: Title獲得、Goal更新、Event参加等のドメイン連携

## Notes (YAGNI原則適用)
- **YAGNI原則**: 必要になったときに実装する
- **最低限実装**: エンティティはコンストラクタと基本的な不変条件チェックのみ
- **アプリケーション層重視**: 複雑なビジネスロジックは上位層で実装
- **拡張性**: 必要に応じて後からメソッドを追加可能な設計
- **データベーススキーマとの整合性を保つ**
- **API仕様との整合性を保つ**

### 実装方針の変更点
- エンティティのメソッドを大幅に削減
- 作成時のバリデーションに集中
- 複雑な制約チェックはアプリケーション層に委譲
- 必要になったタイミングでメソッドを追加