frontend/src/entities
├── index.ts
├── project
│   ├── assignment
│   │   ├── api
│   │   │   ├── assignment-api.ts
│   │   │   └── index.ts
│   │   ├── index.ts
│   │   ├── lib
│   │   │   ├── assignment-validators.ts
│   │   │   └── index.ts
│   │   └── model
│   │       ├── constants.ts
│   │       ├── index.ts
│   │       └── types.ts
│   ├── comment
│   │   ├── api
│   │   │   ├── comment-api.ts
│   │   │   └── index.ts
│   │   ├── index.ts
│   │   ├── lib
│   │   │   ├── comment-validators.ts
│   │   │   └── index.ts
│   │   └── model
│   │       ├── index.ts
│   │       └── types.ts
│   ├── evaluation
│   │   ├── api
│   │   │   ├── evaluation-api.ts
│   │   │   └── index.ts
│   │   ├── index.ts
│   │   └── model
│   │       ├── constants.ts
│   │       ├── index.ts
│   │       ├── types.ts
│   │       └── utils.ts
│   ├── index.ts
│   ├── milestone
│   │   ├── api
│   │   │   ├── index.ts
│   │   │   └── milestone-api.ts
│   │   ├── index.ts
│   │   ├── lib
│   │   │   ├── index.ts
│   │   │   └── milestone-validators.ts
│   │   └── model
│   │       ├── index.ts
│   │       └── types.ts
│   ├── project-entity
│   │   ├── api
│   │   │   ├── index.ts
│   │   │   ├── project-api.ts
│   │   │   └── project-category-api.ts
│   │   ├── index.ts
│   │   ├── lib
│   │   │   ├── index.ts
│   │   │   ├── project-validators.ts
│   │   │   └── status-styles.ts
│   │   └── model
│   │       ├── constants.ts
│   │       ├── index.ts
│   │       └── types.ts
│   └── utils
│       ├── index.ts
│       └── query-keys.ts
├── request
│   ├── action
│   │   ├── index.ts
│   │   └── model
│   │       ├── button-styles.ts
│   │       ├── index.ts
│   │       ├── messages.ts
│   │       └── types.ts
│   ├── document-type
│   │   ├── api
│   │   │   ├── document-type-api.ts
│   │   │   └── index.ts
│   │   ├── index.ts
│   │   └── model
│   │       └── types.ts
│   ├── filter
│   │   ├── index.ts
│   │   └── model
│   │       ├── constants.ts
│   │       ├── index.ts
│   │       └── types.ts
│   ├── index.ts
│   └── request-entity
│       ├── api
│       │   ├── index.ts
│       │   └── request-api.ts
│       ├── index.ts
│       ├── lib
│       │   ├── index.ts
│       │   ├── request-validators.ts
│       │   ├── status-styles.ts
│       │   ├── status-validation.ts
│       │   └── utils.ts
│       └── model
│           ├── api-types.ts
│           ├── constants.ts
│           ├── index.ts
│           └── types.ts
├── user
│   ├── authority
│   │   ├── api
│   │   │   ├── authority-api.ts
│   │   │   └── index.ts
│   │   ├── index.ts
│   │   ├── lib
│   │   │   ├── authority-utils.ts
│   │   │   ├── index.ts
│   │   │   └── useAuthCheck.ts
│   │   ├── model
│   │   │   ├── constants.ts
│   │   │   ├── index.ts
│   │   │   └── types.ts
│   │   ├── ui
│   │   │   ├── index.ts
│   │   │   └── user-status-badge.tsx
│   │   └── utils
│   │       ├── index.ts
│   │       └── query-keys.ts
│   ├── index.ts
│   ├── role
│   │   ├── api
│   │   │   ├── index.ts
│   │   │   └── user-role-api.ts
│   │   ├── index.ts
│   │   ├── model
│   │   │   ├── index.ts
│   │   │   └── types.ts
│   │   └── utils
│   │       ├── index.ts
│   │       └── query-keys.ts
│   ├── skill
│   │   ├── api
│   │   │   ├── index.ts
│   │   │   └── user-skill-api.ts
│   │   ├── index.ts
│   │   ├── lib
│   │   │   ├── index.ts
│   │   │   └── validators.ts
│   │   ├── model
│   │   │   ├── constants.ts
│   │   │   ├── index.ts
│   │   │   └── types.ts
│   │   ├── ui
│   │   │   ├── index.ts
│   │   │   └── skill-rating-widget.tsx
│   │   └── utils
│   │       ├── index.ts
│   │       └── query-keys.ts
│   └── user-entity
│       ├── api
│       │   ├── index.ts
│       │   └── user-api.ts
│       ├── index.ts
│       ├── lib
│       │   ├── index.ts
│       │   └── validators.ts
│       ├── model
│       │   ├── index.ts
│       │   ├── selectors.ts
│       │   ├── slice.ts
│       │   └── types.ts
│       └── utils
│           ├── index.ts
│           └── query-keys.ts
├── utilization
│   ├── index.ts
│   ├── lib
│   │   └── get-utilization-status.ts
│   └── ui
│       └── UtilizationRateBadge.tsx
└── workload
    ├── availability
    │   ├── api
    │   │   ├── availability-api.ts
    │   │   └── index.ts
    │   ├── index.ts
    │   ├── lib
    │   │   ├── helpers.ts
    │   │   ├── index.ts
    │   │   └── validators.ts
    │   ├── model
    │   │   ├── index.ts
    │   │   └── types.ts
    │   └── utils
    │       ├── index.ts
    │       └── query-keys.ts
    ├── demanded
    │   ├── api
    │   │   ├── demanded-api.ts
    │   │   └── index.ts
    │   ├── index.ts
    │   ├── lib
    │   │   ├── index.ts
    │   │   ├── utils.ts
    │   │   └── validators.ts
    │   ├── model
    │   │   ├── index.ts
    │   │   └── types.ts
    │   └── utils
    │       ├── index.ts
    │       └── query-keys.ts
    ├── index.ts
    └── utils
        ├── index.ts
        └── transformers.ts






