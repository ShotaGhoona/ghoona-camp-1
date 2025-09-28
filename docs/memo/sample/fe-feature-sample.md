frontend/src/features
├── assignment
│   ├── assignment-add
│   │   ├── index.ts
│   │   └── lib
│   │       ├── useAssignmentAdd.ts
│   │       └── useMultipleAssignmentAdd.ts
│   ├── assignment-delete
│   │   ├── index.ts
│   │   └── lib
│   │       ├── useAssignmentDelete.ts
│   │       └── useDeleteAssignment.ts
│   ├── assignment-update
│   │   ├── index.ts
│   │   └── lib
│   │       ├── useAssignmentStatus.ts
│   │       └── useUpdateAssignment.ts
│   └── index.ts
├── hr
│   ├── auth
│   │   ├── authority-filter
│   │   │   └── lib
│   │   │       └── use-authority-filter.ts
│   │   ├── authority-list
│   │   │   └── lib
│   │   │       └── use-authorities.ts
│   │   └── index.ts
│   ├── index.ts
│   ├── name-search
│   │   ├── index.ts
│   │   └── lib
│   │       └── use-name-search.ts
│   ├── role
│   │   ├── index.ts
│   │   ├── role-create
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       ├── index.ts
│   │   │       ├── useCreateRole.ts
│   │   │       └── useRoleCreate.ts
│   │   ├── role-delete
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       ├── index.ts
│   │   │       ├── useDeleteRole.ts
│   │   │       └── useRoleDelete.ts
│   │   └── role-get
│   │       ├── index.ts
│   │       └── lib
│   │           ├── index.ts
│   │           └── useGetRoles.ts
│   ├── skill
│   │   ├── index.ts
│   │   ├── skill-create
│   │   │   └── lib
│   │   │       ├── use-skill-create.ts
│   │   │       └── useSkillCreate.ts
│   │   ├── skill-delete
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       ├── index.ts
│   │   │       ├── useDeleteSkill.ts
│   │   │       └── useSkillDelete.ts
│   │   ├── skill-filter
│   │   │   └── lib
│   │   │       └── use-skill-filter.ts
│   │   └── skill-get
│   │       ├── index.ts
│   │       └── lib
│   │           ├── index.ts
│   │           └── useSkills.ts
│   ├── user
│   │   ├── index.ts
│   │   ├── user-create
│   │   │   └── lib
│   │   │       └── use-create-user.ts
│   │   ├── user-delete
│   │   │   └── lib
│   │   │       └── use-delete-user.ts
│   │   ├── user-detail-get
│   │   │   └── lib
│   │   │       └── use-user-detail.ts
│   │   ├── user-list
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       ├── use-user-list.ts
│   │   │       └── useLeaderUsers.ts
│   │   ├── user-roll
│   │   │   ├── add
│   │   │   │   └── lib
│   │   │   │       └── use-add-user-role.ts
│   │   │   ├── delete
│   │   │   │   └── lib
│   │   │   │       └── use-delete-user-role.ts
│   │   │   └── index.ts
│   │   ├── user-skill
│   │   │   ├── add
│   │   │   │   └── lib
│   │   │   │       └── use-add-user-skill.ts
│   │   │   ├── delete
│   │   │   │   └── lib
│   │   │   │       └── use-delete-user-skill.ts
│   │   │   ├── index.ts
│   │   │   └── update
│   │   │       └── lib
│   │   │           └── use-update-user-skill.ts
│   │   ├── user-update
│   │   │   └── lib
│   │   │       └── use-update-user.ts
│   │   └── utilization
│   │       └── lib
│   │           └── use-utilization-sort.ts
│   └── workload
│       ├── availability
│       │   └── lib
│       │       └── use-availabilities.ts
│       ├── demanded
│       │   ├── index.ts
│       │   ├── lib
│       │   │   ├── use-demanded.ts
│       │   │   └── useDemandedCreate.ts
│       │   └── ui
│       │       └── DemandedCreateDialog.tsx
│       └── index.ts
├── index.ts
├── login
│   ├── index.ts
│   └── lib
│       ├── index.ts
│       └── useLoginMutation.ts
├── project
│   ├── comment
│   │   ├── comment-create
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       └── useCommentCreate.ts
│   │   ├── comment-delete
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       └── useCommentDelete.ts
│   │   ├── comment-edit
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       └── useCommentEdit.ts
│   │   ├── comment-get
│   │   │   ├── index.ts
│   │   │   └── useCommentGet.ts
│   │   └── index.ts
│   ├── evaluation
│   │   ├── evaluation-create
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       ├── useCreateEvaluation.ts
│   │   │       └── useEvaluationCreate.ts
│   │   ├── evaluation-list
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       ├── useLatestEvaluations.ts
│   │   │       └── useProjectEvaluations.ts
│   │   └── index.ts
│   ├── index.ts
│   ├── milestone
│   │   ├── index.ts
│   │   ├── milestone-add
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       ├── useAddMilestone.ts
│   │   │       └── useMilestoneAdd.ts
│   │   ├── milestone-delete
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       ├── useDeleteMilestone.ts
│   │   │       └── useMilestoneDelete.ts
│   │   ├── milestone-edit
│   │   │   ├── index.ts
│   │   │   └── lib
│   │   │       ├── useMilestoneEdit.ts
│   │   │       └── useUpdateMilestone.ts
│   │   └── milestone-master
│   │       ├── index.ts
│   │       └── lib
│   │           ├── useCreateMilestoneMaster.ts
│   │           ├── useDeleteMilestoneMaster.ts
│   │           ├── useMilestone.ts
│   │           ├── useMilestones.ts
│   │           └── useUpdateMilestoneMaster.ts
│   ├── project-category
│   │   ├── index.ts
│   │   └── lib
│   │       └── useProjectCategories.ts
│   ├── project-category-navigate
│   │   ├── index.ts
│   │   └── lib
│   │       └── useProjectCategoryNavigate.ts
│   ├── project-create
│   │   ├── index.ts
│   │   └── lib
│   │       └── useCreateProject.ts
│   ├── project-delete
│   │   ├── index.ts
│   │   └── lib
│   │       └── useProjectDelete.ts
│   ├── project-edit
│   │   ├── index.ts
│   │   └── lib
│   │       └── useProjectEdit.ts
│   ├── project-filter
│   │   ├── index.ts
│   │   ├── lib
│   │   │   └── useProjectFilter.ts
│   │   └── ui
│   │       └── ProjectFilterBar.tsx
│   ├── project-list
│   │   ├── index.ts
│   │   └── lib
│   │       ├── useProject.ts
│   │       └── useProjectDetail.ts
│   └── project-navigate
│       ├── index.ts
│       └── lib
│           └── useProjectNavigate.ts
└── request
    ├── bulk-action
    │   ├── index.ts
    │   └── lib
    │       ├── bulk-processing-utils.ts
    │       └── index.ts
    ├── create
    │   ├── index.ts
    │   └── lib
    │       └── index.ts
    ├── document-type
    │   ├── document-type-create
    │   │   ├── index.ts
    │   │   └── lib
    │   │       ├── index.ts
    │   │       ├── useCreateDocumentType.ts
    │   │       └── useDocumentTypeCreate.ts
    │   ├── document-type-delete
    │   │   ├── index.ts
    │   │   └── lib
    │   │       ├── index.ts
    │   │       ├── useDeleteDocumentType.ts
    │   │       └── useDocumentTypeDelete.ts
    │   ├── document-type-get
    │   │   ├── index.ts
    │   │   └── lib
    │   │       ├── index.ts
    │   │       └── useGetDocumentTypes.ts
    │   ├── document-type-list
    │   │   ├── index.ts
    │   │   └── lib
    │   │       └── useDocumentTypes.ts
    │   └── index.ts
    ├── index.ts
    ├── list
    │   ├── index.ts
    │   └── lib
    │       └── index.ts
    ├── selection
    │   ├── index.ts
    │   ├── lib
    │   │   └── index.ts
    │   └── model
    │       └── index.ts
    └── update
        ├── index.ts
        └── lib
            └── index.ts
