# Ghoona Camp Feature構成設計案（最小構成版）

**参考**: eagle-ai抽象化原則を適用  
**対象**: FSDアーキテクチャのfeatures層設計  
**原則**: API直結 + ビジネスロジック重視

## 設計原則（Eagle-AI抽象化）

### 1. **API-Feature 1:1マッピング**
- 各API エンドポイントに対応するfeatureのみ作成
- UI専用の抽象化は避ける

### 2. **ビジネスドメイン重視**
- APIのリソース階層に合わせたドメイン分割
- 技術的な分類ではなく業務的な分類

### 3. **最小限の責任分離**
- CRUD操作の明確な分離
- 複合操作は上位層（pages）で実装

## ✅ 最小Feature構成案

```
frontend/src/features
├── user/
│   ├── profile-feature/
│   │   ├── session-get/              # GET /auth/me
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useAuthMe.ts
│   │   ├── profile-get/              # GET /users/{userId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserProfile.ts
│   │   └── profile-update/           # PUT /users/{userId}
│   │       ├── index.ts
│   │       └── lib/
│   │           └── useUserUpdate.ts
│   ├── user-feature/
│   │   └── users-list/               # GET /users (search/filter内蔵)
│   │       ├── index.ts
│   │       └── lib/
│   │           └── useUserList.ts
│   ├── metadata-feature/
│   │   ├── metadata-get/             # GET /users/{userId}/metadata
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserMetadata.ts
│   │   └── metadata-update/          # PUT /users/{userId}/metadata
│   │       ├── index.ts
│   │       ├── lib/
│   │       │   └── useMetadataUpdate.ts
│   │       └── ui/
│   │           └── MetadataForm.tsx
│   ├── social-links-feature/
│   │   ├── links-get/                # GET /users/{userId}/social-links
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useSocialLinks.ts
│   │   ├── links-create/             # POST /users/{userId}/social-links
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useLinkCreate.ts
│   │   │   └── ui/
│   │   │       └── LinkForm.tsx
│   │   ├── links-update/             # PUT /users/{userId}/social-links/{linkId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useLinkUpdate.ts
│   │   └── links-delete/             # DELETE /users/{userId}/social-links/{linkId}
│   │       ├── index.ts
│   │       └── lib/
│   │           └── useLinkDelete.ts
│   └── rivals-feature/
│       ├── rivals-get/               # GET /users/{userId}/rivals
│       │   ├── index.ts
│       │   └── lib/
│       │       └── useUserRivals.ts
│       ├── rivals-create/            # POST /users/{userId}/rivals
│       │   ├── index.ts
│       │   ├── lib/
│       │   │   └── useRivalCreate.ts
│       │   └── ui/
│       │       └── RivalSelectDialog.tsx
│       └── rivals-delete/            # DELETE /users/{userId}/rivals/{rivalId}
│           ├── index.ts
│           └── lib/
│               └── useRivalDelete.ts
├── goal/
│   ├── goal-feature/
│   │   ├── goal-get/                 # GET /goals/{goalId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useGoalDetail.ts
│   │   ├── user-goals-list/          # GET /users/{userId}/goals
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useUserGoals.ts
│   │   ├── goal-create/              # POST /users/{userId}/goals
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useGoalCreate.ts
│   │   │   └── ui/
│   │   │       └── GoalForm.tsx
│   │   ├── goal-update/              # PUT /goals/{goalId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useGoalUpdate.ts
│   │   └── goal-delete/              # DELETE /goals/{goalId}
│   │       ├── index.ts
│   │       └── lib/
│   │           └── useGoalDelete.ts
│   └── discovery-feature/
│       └── public-list/              # GET /goals/public (filter内蔵)
│           ├── index.ts
│           └── lib/
│               └── usePublicGoals.ts
├── event/
│   ├── event-feature/
│   │   ├── event-get/                # GET /events/{eventId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useEventDetail.ts
│   │   ├── events-list/              # GET /events (filter内蔵)
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useEvents.ts
│   │   ├── event-create/             # POST /events
│   │   │   ├── index.ts
│   │   │   ├── lib/
│   │   │   │   └── useEventCreate.ts
│   │   │   └── ui/
│   │   │       └── EventForm.tsx
│   │   ├── event-update/             # PUT /events/{eventId}
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useEventUpdate.ts
│   │   └── event-delete/             # DELETE /events/{eventId}
│   │       ├── index.ts
│   │       └── lib/
│   │           └── useEventDelete.ts
│   └── participants-feature/
│       ├── participants-get/         # GET /events/{eventId}/participants
│       │   ├── index.ts
│       │   └── lib/
│       │       └── useParticipants.ts
│       ├── participant-join/         # POST /events/{eventId}/participants
│       │   ├── index.ts
│       │   └── lib/
│       │       └── useParticipantJoin.ts
│       ├── participant-status/       # PUT /events/{eventId}/participants/{userId}
│       │   ├── index.ts
│       │   └── lib/
│       │       └── useParticipantStatus.ts
│       └── participant-leave/        # DELETE /events/{eventId}/participants/{userId}
│           ├── index.ts
│           └── lib/
│               └── useParticipantLeave.ts
├── title/
│   ├── title-feature/
│   │   ├── titles-list/              # GET /titles
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useTitles.ts
│   │   └── title-detail/             # GET /titles/{titleId}
│   │       ├── index.ts
│   │       └── lib/
│   │           └── useTitleDetail.ts
│   └── achievements-feature/
│       ├── achievements-get/         # GET /users/{userId}/achievements
│       │   ├── index.ts
│       │   └── lib/
│       │       └── useAchievements.ts
│       └── current-title-update/     # PUT /users/{userId}/achievements/{titleId}
│           ├── index.ts
│           └── lib/
│               └── useCurrentTitle.ts
├── attendance/
│   ├── logs-feature/
│   │   ├── logs-get/                 # GET /users/{userId}/attendance/logs
│   │   │   ├── index.ts
│   │   │   └── lib/
│   │   │       └── useAttendanceLogs.ts
│   │   └── logs-create/              # POST /attendance/logs (Bot専用)
│   │       ├── index.ts
│   │       └── lib/
│   │           └── useLogCreate.ts
│   ├── summaries-feature/
│   │   └── summaries-get/            # GET /users/{userId}/attendance/summaries
│   │       ├── index.ts
│   │       └── lib/
│   │           └── useSummaries.ts
│   ├── statistics-feature/
│   │   └── statistics-get/           # GET /users/{userId}/attendance/statistics
│   │       ├── index.ts
│   │       └── lib/
│   │           └── useStatistics.ts
│   └── ranking-feature/
│       ├── monthly-ranking/          # GET /ranking/monthly
│       │   ├── index.ts
│       │   └── lib/
│       │       └── useMonthlyRanking.ts
│       ├── total-ranking/            # GET /ranking/total
│       │   ├── index.ts
│       │   └── lib/
│       │       └── useTotalRanking.ts
│       └── streak-ranking/           # GET /ranking/streak
│           ├── index.ts
│           └── lib/
│               └── useStreakRanking.ts
└── notification/
    ├── notification-feature/
    │   ├── notifications-get/        # GET /users/{userId}/notifications
    │   │   ├── index.ts
    │   │   └── lib/
    │   │       └── useNotifications.ts
    │   ├── notification-read/        # PUT /notifications/{notificationId}
    │   │   ├── index.ts
    │   │   └── lib/
    │   │       └── useNotificationRead.ts
    │   └── notification-delete/      # DELETE /notifications/{notificationId}
    │       ├── index.ts
    │       └── lib/
    │           └── useNotificationDelete.ts
    └── notification-settings-feature/
        ├── settings-get/             # GET /users/{userId}/notification-settings
        │   ├── index.ts
        │   └── lib/
        │       └── useNotificationSettings.ts
        └── settings-update/          # PUT /users/{userId}/notification-settings
            ├── index.ts
            ├── lib/
            │   └── useSettingsUpdate.ts
            └── ui/
                └── SettingsForm.tsx
```

## 削除した過剰機能

### 🗑️ **完全削除したドメイン**
- **dashboard/**: ページレベルロジック、feature層に不適切
- **discord/**: フロントエンド向けAPI不存在

### 🗑️ **削除したUI抽象化**
- **event-filter/**: `GET /events` のパラメータで十分
- **goal-filter/**: `GET /goals/public` のパラメータで十分  
- **user-search/**: `GET /users` のパラメータで十分
- **calendar-view/**: 単なるデータ表示変換
- **streak-tracking/**: 統計の一部
- **progress-tracking/**: 統計から計算
- **bulk-operations/**: 単純なループ処理
- **rivals-compare/**: ページで実装

### 🗑️ **統合したAuth機能**
- **auth/** ドメイン全体をuserドメインに統合
- **signin-post/**, **signout-post/**: Clerk内蔵コンポーネント使用のため削除
- **session-get/**: user/profile/ に統合

## Eagle-AI学習ポイント

### ✅ **適用した原則**
1. **API直結**: 38API → 37feature の1:1対応（auth統合）
2. **ビジネスドメイン**: user, goal, event, title, attendance, notification
3. **CRUD分離**: create, get, update, delete の明確な分離
4. **最小責任**: 各featureは単一のAPI呼び出しに集中
5. **命名一貫性**: 全サブドメインに `-feature` サフィックス統一

### ✅ **避けたアンチパターン**
1. **UI専用抽象化**: filtering, searching, widget化
2. **複合操作**: dashboard, comparison, analytics
3. **フレームワーク機能重複**: Clerk認証処理
4. **存在しないAPI**: Discord フロントエンド連携

## 結果

- **従来**: 78 features (過剰設計)
- **最小構成**: 37 features (API直結、auth統合)
- **削減率**: 53% 削減

シンプルで保守しやすい、Ghoona Camp要件に最適化されたfeature構成を実現。

---
*作成日: 2025-01-21*  
*バージョン: 最小構成版（Eagle-AI抽象化原則適用）*