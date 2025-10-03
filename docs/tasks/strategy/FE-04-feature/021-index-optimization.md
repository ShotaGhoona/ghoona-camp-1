# 命令書

feature層のindex.tsのカプセル化をentity層のカプセル化のレベルと合わせたい（userドメインのみ）
- サブディレクトリ（lib, api, model）の中にはindex.tsを入れない
- サブディレクトリ（lib, api, model）と同じ階層にindex.tsを入れる
- ドメインの中を外に出すためにindex.ts

# サブディレクトリ（lib, api, model）と同じ階層のindex.tsの例
// === API Functions ===
export * from './api/metadata-api';

// === Type Definitions ===
export type * from './model/metadata-types';

// === Validation Schemas & Type Guards ===
export * from './lib/validators';
export type * from './lib/validators';

// === Query Keys ===
export * from './lib/query-keys';

// === Mappers ===
export * from './lib/mappers';

# ドメインを外に出すやつ

// User Entity Re-exports

// === User Entity ===
export * from './user-entity';
export type * from './user-entity';

// === Metadata Entity ===
export * from './metadata-entity';
export type * from './metadata-entity';

// === Rivals Entity ===
export * from './rivals-entity';
export type * from './rivals-entity';

// === Social Links Entity ===
export * from './social-links-entity';
export type * from './social-links-entity';


# 現在のentity層これが完璧


yamashitashota@Macintosh-40743 ghoona-camp-1 % tr
ee frontend/src/entities
frontend/src/entities
└── user
    ├── index.ts
    ├── metadata-entity
    │   ├── api
    │   │   └── metadata-api.ts
    │   ├── index.ts
    │   ├── lib
    │   │   ├── mappers.ts
    │   │   ├── query-keys.ts
    │   │   └── validators.ts
    │   └── model
    │       └── metadata-types.ts
    ├── rivals-entity
    │   ├── api
    │   │   └── rivals-api.ts
    │   ├── index.ts
    │   ├── lib
    │   │   ├── mappers.ts
    │   │   ├── query-keys.ts
    │   │   └── validators.ts
    │   └── model
    │       └── rivals-types.ts
    ├── social-links-entity
    │   ├── api
    │   │   └── social-links-api.ts
    │   ├── index.ts
    │   ├── lib
    │   │   ├── mappers.ts
    │   │   ├── query-keys.ts
    │   │   └── validators.ts
    │   └── model
    │       └── social-links-types.ts
    └── user-entity
        ├── api
        │   └── user-api.ts
        ├── index.ts
        ├── lib
        │   ├── mappers.ts
        │   ├── query-keys.ts
        │   └── validators.ts
        └── model
            └── user-types.ts

# 現在のfeatures　これを変えたい

yamashitashota@Macintosh-40743 ghoona-camp-1 % tr
ee frontend/src/features
frontend/src/features
└── user
    ├── index.ts
    ├── metadata-feature
    │   ├── index.ts
    │   ├── metadata-get
    │   │   ├── index.ts
    │   │   ├── lib
    │   │   │   ├── index.ts
    │   │   │   └── use-metadata-get.ts
    │   │   └── ui
    │   └── metadata-update
    │       ├── index.ts
    │       ├── lib
    │       │   ├── index.ts
    │       │   └── use-metadata-update.ts
    │       ├── model
    │       │   └── form-types.ts
    │       └── ui
    ├── profile-feature
    │   ├── index.ts
    │   ├── profile-get
    │   │   ├── index.ts
    │   │   ├── lib
    │   │   │   ├── index.ts
    │   │   │   └── use-profile-get.ts
    │   │   └── ui
    │   ├── profile-update
    │   │   ├── index.ts
    │   │   ├── lib
    │   │   │   ├── index.ts
    │   │   │   └── use-profile-update.ts
    │   │   └── ui
    │   └── session-get
    │       ├── index.ts
    │       ├── lib
    │       │   ├── index.ts
    │       │   └── use-session-get.ts
    │       └── ui
    ├── rivals-feature
    │   ├── index.ts
    │   ├── rivals-create
    │   │   ├── index.ts
    │   │   ├── lib
    │   │   │   ├── index.ts
    │   │   │   └── use-rivals-create.ts
    │   │   └── ui
    │   ├── rivals-delete
    │   │   ├── index.ts
    │   │   ├── lib
    │   │   │   ├── index.ts
    │   │   │   └── use-rivals-delete.ts
    │   │   └── ui
    │   └── rivals-get
    │       ├── index.ts
    │       ├── lib
    │       │   ├── index.ts
    │       │   └── use-rivals-get.ts
    │       └── ui
    ├── social-links-feature
    │   ├── index.ts
    │   ├── links-create
    │   │   ├── index.ts
    │   │   ├── lib
    │   │   │   ├── index.ts
    │   │   │   └── use-links-create.ts
    │   │   └── ui
    │   ├── links-delete
    │   │   ├── index.ts
    │   │   ├── lib
    │   │   │   ├── index.ts
    │   │   │   └── use-links-delete.ts
    │   │   └── ui
    │   ├── links-get
    │   │   ├── index.ts
    │   │   ├── lib
    │   │   │   ├── index.ts
    │   │   │   └── use-links-get.ts
    │   │   └── ui
    │   └── links-update
    │       ├── index.ts
    │       ├── lib
    │       │   ├── index.ts
    │       │   └── use-links-update.ts
    │       └── ui
    └── user-feature
        ├── index.ts
        ├── user-detail-get
        │   ├── index.ts
        │   └── lib
        │       ├── index.ts
        │       └── use-user-detail.ts
        └── users-list
            ├── index.ts
            ├── lib
            │   ├── index.ts
            │   └── use-users-list.ts
            └── ui
