# Ghoona Camp Entity構成設計案

**参考**: eagle-ai プロジェクトのentity構成  
**対象**: FSDアーキテクチャのentities層設計

## 概要
Ghoona Campの朝活コミュニティアプリにおけるドメインエンティティの構成案。DB設計とAPI設計を基に、[domain]/[sub-entity]/[layer]の構成で設計する。

## DB設計ベースのドメイン分析
DB設計から抽出したドメインとサブエンティティ：

- **user**: ユーザー管理ドメイン
  - profile: 基本プロフィール
  - metadata: 詳細メタデータ
  - social-links: 外部リンク
  - rivals: ライバル関係

- **attendance**: 出席管理ドメイン
  - logs: 参加ログ
  - summaries: 日次サマリー
  - statistics: 統計情報
  - ranking: ランキング

- **goal**: 目標管理ドメイン
  - goal-entity: 目標本体

- **event**: イベント管理ドメイン
  - event-entity: イベント本体
  - participants: 参加者管理

- **title**: 称号管理ドメイン
  - title-entity: 称号マスター
  - achievements: 称号実績

- **notification**: 通知管理ドメイン
  - notification-entity: 通知本体
  - settings: 通知設定

## 🔄 Eagle-AI参考による構成見直し

**Eagle-AIからの学び:**
- APIエンドポイント階層 = Entity構造
- 関連リソースのドメイン統合
- UI固有entityの必要性

## 📊 API-Entity対応表

| Ghoona Camp API | Entity構造 | Eagle-AI類似パターン |
|---|---|---|
| `/users/{userId}/metadata` | `user/metadata` | `/users/{userId}/skills` → `user/skill` |
| `/users/{userId}/rivals` | `user/rivals` | `/users/{userId}/roles` → `user/role` |
| `/events/{eventId}/participants` | `event/participants` | `/projects/{projectId}/assignments` → `project/assignment` |
| `/users/{userId}/attendance/logs` | `attendance/logs` | `/availabilities` → `workload/availability` |
| `/ranking/monthly` | `attendance/ranking` | 独立API → 関連domainに統合 |

## ✅ 最終Entity構成案（User統合版）

### 1. User Domain
```
entities/user/
├── user-entity/         # /auth/me, /users/{userId} (認証・プロフィール統合)
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── metadata-entity/     # /users/{userId}/metadata
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── social-links-entity/ # /users/{userId}/social-links
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── rivals-entity/       # /users/{userId}/rivals
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── utils/
│   └── query-keys.ts
└── index.ts
```

### 2. Attendance Domain  
```
entities/attendance/
├── logs-entity/         # /users/{userId}/attendance/logs
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── summaries-entity/    # /users/{userId}/attendance/summaries
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── statistics-entity/   # /users/{userId}/attendance/statistics
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── ranking-entity/      # /ranking/monthly, /ranking/total, /ranking/streak
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── utils/
│   └── query-keys.ts
└── index.ts
```

### 3. Goal Domain
```
entities/goal/
├── goal-entity/         # /goals/{goalId}, /users/{userId}/goals
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── utils/
│   └── query-keys.ts
└── index.ts
```

### 4. Event Domain
```
entities/event/
├── event-entity/        # /events/{eventId}
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── participants-entity/ # /events/{eventId}/participants
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── utils/
│   └── query-keys.ts
└── index.ts
```

### 5. Title Domain
```
entities/title/
├── title-entity/        # /titles, /titles/{titleId}
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── achievements-entity/ # /users/{userId}/achievements
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── utils/
│   └── query-keys.ts
└── index.ts
```

### 6. Notification Domain
```
entities/notification/
├── notification-entity/  # /users/{userId}/notifications
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── settings-entity/     # /users/{userId}/notification-settings
│   ├── api/
│   ├── lib/
│   ├── model/
│   └── index.ts
├── utils/
│   └── query-keys.ts
└── index.ts
```

## 設計原則
1. **ドメイン境界の明確化**: 各entityは独立したドメインロジックを持つ
2. **型安全性**: TypeScript型定義の徹底
3. **再利用性**: 小さなUIコンポーネントで構成
4. **テスタビリティ**: ロジックとUIの分離

## 次のステップ
この構成案を基に、FE-01-setup-02でentitiesディレクトリを作成し、基本的なファイル構造を構築する。

---
*作成日: 2025-01-21*