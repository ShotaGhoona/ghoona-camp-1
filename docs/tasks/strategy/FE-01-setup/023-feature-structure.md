# Ghoona Camp Feature構成設計案

**参考**: eagle-ai プロジェクトのfeature構成  
**対象**: FSDアーキテクチャのfeatures層設計  
**パターン**: `[domain]/[subdomain-feature]/xxx-get` 構成

## 概要
Ghoona Campの朝活コミュニティアプリにおけるfeature層の構成案。DB設計とAPI設計を基に、ドメイン→サブドメインfeature→操作の3層構造で設計する。

## 設計パターン
```
[domain]/[subdomain-feature]/[operation]-[action]
```

- **domain**: ユーザー、出席管理、目標、イベント等
- **subdomain-feature**: より具体的な機能領域  
- **operation**: get, list, create, update, delete
- **action**: 操作の詳細（get, detail, search等）

## ✅ 最終Feature構成案（Eagle-AIパターン適用）

```
frontend/src/features
├── auth/
│   ├── auth-session/
│   │   ├── session-get/           # GET /auth/me
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useAuthMe.ts
│   │   ├── signin-post/           # Clerk認証処理
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useAuthSignin.ts
│   │   ├── signout-post/          # Clerk認証解除
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useAuthSignout.ts
│   │   └── index.ts
│   └── index.ts
├── user/
│   ├── user-profile/
│   │   ├── profile-get/           # GET /users/{userId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserProfile.ts
│   │   ├── profile-update/        # PUT /users/{userId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserUpdate.ts
│   │   └── index.ts
│   ├── user-directory/
│   │   ├── users-list/            # GET /users
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserList.ts
│   │   ├── user-search/           # GET /users + フィルター
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useUserSearch.ts
│   │   │   └── ui/
│   │   │       └── UserSearchForm.tsx
│   │   └── index.ts
│   ├── user-metadata/
│   │   ├── metadata-get/          # GET /users/{userId}/metadata
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserMetadata.ts
│   │   ├── metadata-update/       # PUT /users/{userId}/metadata
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useMetadataUpdate.ts
│   │   │   └── ui/
│   │   │       └── MetadataEditForm.tsx
│   │   └── index.ts
│   ├── user-social-links/
│   │   ├── links-get/             # GET /users/{userId}/social-links
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useSocialLinks.ts
│   │   ├── links-create/          # POST /users/{userId}/social-links
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useSocialLinkCreate.ts
│   │   │   └── ui/
│   │   │       └── SocialLinkForm.tsx
│   │   ├── links-update/          # PUT /users/{userId}/social-links/{linkId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useSocialLinkUpdate.ts
│   │   ├── links-delete/          # DELETE /users/{userId}/social-links/{linkId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useSocialLinkDelete.ts
│   │   └── index.ts
│   ├── user-rivals/
│   │   ├── rivals-get/            # GET /users/{userId}/rivals
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserRivals.ts
│   │   ├── rivals-create/         # POST /users/{userId}/rivals
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useRivalCreate.ts
│   │   │   └── ui/
│   │   │       └── RivalSelectDialog.tsx
│   │   ├── rivals-delete/         # DELETE /users/{userId}/rivals/{rivalId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useRivalDelete.ts
│   │   ├── rivals-compare/        # 複数ユーザー統計比較（UI専用機能）
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useRivalComparison.ts
│   │   └── index.ts
│   └── index.ts
├── goal/
│   ├── goal-management/
│   │   ├── goal-get/              # GET /goals/{goalId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useGoalDetail.ts
│   │   ├── user-goals-list/       # GET /users/{userId}/goals
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserGoals.ts
│   │   ├── goal-create/           # POST /users/{userId}/goals
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useGoalCreate.ts
│   │   │   └── ui/
│   │   │       └── GoalCreateForm.tsx
│   │   ├── goal-update/           # PUT /goals/{goalId}
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useGoalUpdate.ts
│   │   │   └── ui/
│   │   │       └── GoalEditForm.tsx
│   │   ├── goal-delete/           # DELETE /goals/{goalId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useGoalDelete.ts
│   │   └── index.ts
│   ├── goal-discovery/
│   │   ├── public-list/           # GET /goals/public
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── usePublicGoals.ts
│   │   ├── goal-filter/           # 公開目標フィルター（UI専用機能）
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useGoalFilter.ts
│   │   │   └── ui/
│   │   │       └── GoalFilterBar.tsx
│   │   └── index.ts
│   └── index.ts
├── event/
│   ├── event-management/
│   │   ├── event-get/             # GET /events/{eventId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useEventDetail.ts
│   │   ├── events-list/           # GET /events
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useEvents.ts
│   │   ├── event-create/          # POST /events
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useEventCreate.ts
│   │   │   └── ui/
│   │   │       └── EventCreateForm.tsx
│   │   ├── event-update/          # PUT /events/{eventId}
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useEventUpdate.ts
│   │   │   └── ui/
│   │   │       └── EventEditForm.tsx
│   │   ├── event-delete/          # DELETE /events/{eventId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useEventDelete.ts
│   │   └── index.ts
│   ├── event-participation/
│   │   ├── participants-get/      # GET /events/{eventId}/participants
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useEventParticipants.ts
│   │   ├── participant-join/      # POST /events/{eventId}/participants
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useParticipantJoin.ts
│   │   ├── participant-status/    # PUT /events/{eventId}/participants/{userId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useParticipantStatus.ts
│   │   ├── participant-leave/     # DELETE /events/{eventId}/participants/{userId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useParticipantLeave.ts
│   │   └── index.ts
│   ├── event-filter/
│   │   ├── date-filter/           # 日付フィルター（UI専用機能）
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useEventDateFilter.ts
│   │   ├── type-filter/           # イベントタイプフィルター（UI専用機能）
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useEventTypeFilter.ts
│   │   │   └── ui/
│   │   │       └── EventFilterBar.tsx
│   │   └── index.ts
│   └── index.ts
├── title/
│   ├── title-system/
│   │   ├── titles-list/           # GET /titles
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useTitles.ts
│   │   ├── title-detail/          # GET /titles/{titleId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useTitleDetail.ts
│   │   └── index.ts
│   ├── title-achievements/
│   │   ├── achievements-get/      # GET /users/{userId}/achievements
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserAchievements.ts
│   │   ├── current-title-update/  # PUT /users/{userId}/achievements/{titleId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useCurrentTitleUpdate.ts
│   │   ├── progress-tracking/     # 称号進捗計算（UI専用機能）
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useTitleProgress.ts
│   │   └── index.ts
│   └── index.ts
├── attendance/
│   ├── attendance-logs/
│   │   ├── logs-get/              # GET /users/{userId}/attendance/logs
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useAttendanceLogs.ts
│   │   ├── logs-create/           # POST /attendance/logs (Discord Bot用)
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useAttendanceCreate.ts
│   │   └── index.ts
│   ├── attendance-summaries/
│   │   ├── summaries-get/         # GET /users/{userId}/attendance/summaries
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useAttendanceSummaries.ts
│   │   ├── calendar-view/         # カレンダー表示（UI専用機能）
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useCalendarData.ts
│   │   │   └── ui/
│   │   │       └── AttendanceCalendar.tsx
│   │   └── index.ts
│   ├── attendance-statistics/
│   │   ├── statistics-get/        # GET /users/{userId}/attendance/statistics
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useAttendanceStats.ts
│   │   ├── streak-tracking/       # 連続日数追跡（UI専用機能）
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useStreakTracking.ts
│   │   └── index.ts
│   ├── attendance-ranking/
│   │   ├── monthly-ranking/       # GET /ranking/monthly
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useMonthlyRanking.ts
│   │   ├── total-ranking/         # GET /ranking/total
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useTotalRanking.ts
│   │   ├── streak-ranking/        # GET /ranking/streak
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useStreakRanking.ts
│   │   └── index.ts
│   └── index.ts
├── notification/
│   ├── notification-management/
│   │   ├── notifications-get/     # GET /users/{userId}/notifications
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useNotifications.ts
│   │   ├── notification-read/     # PUT /notifications/{notificationId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useNotificationRead.ts
│   │   ├── notification-delete/   # DELETE /notifications/{notificationId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useNotificationDelete.ts
│   │   ├── bulk-operations/       # 一括既読・削除（UI専用機能）
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useNotificationBulk.ts
│   │   └── index.ts
│   ├── notification-settings/
│   │   ├── settings-get/          # GET /users/{userId}/notification-settings
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useNotificationSettings.ts
│   │   ├── settings-update/       # PUT /users/{userId}/notification-settings
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useSettingsUpdate.ts
│   │   │   └── ui/
│   │   │       └── NotificationSettingsForm.tsx
│   │   └── index.ts
│   └── index.ts
├── discord/
│   ├── discord-integration/
│   │   ├── integration-status/    # Discord連携状態確認（UI専用機能）
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useDiscordStatus.ts
│   │   ├── integration-setup/     # Discord連携設定（UI専用機能）
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useDiscordSetup.ts
│   │   │   └── ui/
│   │   │       └── DiscordSetupDialog.tsx
│   │   └── index.ts
│   ├── discord-attendance/
│   │   ├── webhook-handler/       # Discord Webhook処理（Bot用）
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useDiscordWebhook.ts
│   │   ├── real-time-status/      # リアルタイム参加状況（UI専用機能）
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useDiscordRealTime.ts
│   │   └── index.ts
│   └── index.ts
├── dashboard/
│   ├── dashboard-overview/
│   │   ├── personal-overview/     # 個人ダッシュボード（複数API統合）
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── usePersonalOverview.ts
│   │   ├── rival-comparison/      # ライバル比較表示（複数API統合）
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useRivalComparison.ts
│   │   │   └── ui/
│   │   │       └── RivalComparisonCard.tsx
│   │   └── index.ts
│   ├── dashboard-widgets/
│   │   ├── attendance-widget/     # 出席ウィジェット（UI専用機能）
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useAttendanceWidget.ts
│   │   │   └── ui/
│   │   │       └── AttendanceWidget.tsx
│   │   ├── goal-widget/           # 目標ウィジェット（UI専用機能）
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useGoalWidget.ts
│   │   │   └── ui/
│   │   │       └── GoalWidget.tsx
│   │   ├── title-widget/          # 称号ウィジェット（UI専用機能）
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useTitleWidget.ts
│   │   │   └── ui/
│   │   │       └── TitleWidget.tsx
│   │   └── index.ts
│   └── index.ts
└── index.ts
```

## 設計原則（Eagle-AIパターン適用）

### 1. API直結 vs UI専用機能の分離
- **API直結**: 実際のエンドポイントと1:1対応する機能
- **UI専用**: フィルター、比較、ウィジェットなど画面表示のためのロジック

### 2. 複合機能の抽象化
- **dashboard**: 複数APIを統合した画面レベルの機能
- **bulk-operations**: 複数API呼び出しをまとめた一括処理
- **comparison**: 複数データの比較ロジック

### 3. ビジネスロジック重視の構成
- **user-directory**: ユーザー一覧・検索
- **goal-discovery**: 公開目標の発見
- **event-participation**: イベント参加管理
- **title-achievements**: 称号取得・進捗管理

### 4. Ghoona Camp特化機能
- **朝活特化**: 6:00-6:30時間帯重視のイベント・出席管理
- **ライバル機能**: 最大3人比較、統計比較
- **称号システム**: 8段階進捗、称号選択
- **Discord連携**: リアルタイム出席、Webhook処理

### 5. 命名規則
- **単数形**: 単体データ操作（profile-get, goal-create）
- **複数形**: 一覧データ操作（users-list, events-list）
- **機能名**: UI専用機能（calendar-view, bulk-operations）

## 次のステップ
この構成案を基に、features層のディレクトリ構造を実装し、各機能の基本ファイルを作成する。

---
*更新日: 2025-01-21*
*バージョン: Eagle-AIパターン適用版*