backend/internal
├── application
│   ├── dto
│   │   ├── approval_context
│   │   │   ├── document_type
│   │   │   │   ├── document_type_request.go
│   │   │   │   └── document_type_response.go
│   │   │   └── request
│   │   │       ├── request_request.go
│   │   │       └── request_response.go
│   │   ├── auth_context
│   │   │   ├── authority
│   │   │   │   └── authority_response.go
│   │   │   ├── role
│   │   │   │   ├── role_request.go
│   │   │   │   └── role_response.go
│   │   │   ├── skill
│   │   │   │   ├── skill_request.go
│   │   │   │   └── skill_response.go
│   │   │   └── user
│   │   │       ├── user_request.go
│   │   │       └── user_response.go
│   │   ├── auth_dto.go
│   │   ├── auth_request.go
│   │   ├── auth_response.go
│   │   ├── milestone_request.go
│   │   ├── milestone_response.go
│   │   ├── project_management
│   │   │   ├── assignment
│   │   │   │   ├── assignment_request.go
│   │   │   │   └── assignment_response.go
│   │   │   ├── comment
│   │   │   │   ├── comment_request.go
│   │   │   │   └── comment_response.go
│   │   │   ├── common
│   │   │   │   └── types.go
│   │   │   ├── evaluation
│   │   │   │   ├── evaluation_batch_response.go
│   │   │   │   ├── evaluation_request.go
│   │   │   │   └── evaluation_response.go
│   │   │   ├── milestone
│   │   │   │   ├── milestone_request.go
│   │   │   │   └── milestone_response.go
│   │   │   └── project
│   │   │       ├── project_request.go
│   │   │       └── project_response.go
│   │   └── workload_context
│   │       ├── availability
│   │       │   ├── availability_request.go
│   │       │   └── availability_response.go
│   │       └── demanded
│   │           ├── demanded_request.go
│   │           └── demanded_response.go
│   ├── transaction
│   │   └── manager.go
│   └── usecase
│       ├── auth_service.go
│       ├── authority_service.go
│       ├── document_type_service.go
│       ├── milestone_service.go
│       ├── project_service.go
│       ├── request_service.go
│       ├── role_service.go
│       ├── skill_service.go
│       ├── token_service.go
│       ├── user_service.go
│       ├── weekly_availability_service.go
│       └── weekly_demanded_service.go
├── di
│   └── container.go
├── domain
│   ├── auth
│   │   ├── entity
│   │   │   ├── user_role.go
│   │   │   ├── user_skill.go
│   │   │   └── user.go
│   │   ├── errors.go
│   │   ├── repository
│   │   │   ├── authority_repository.go
│   │   │   ├── role_repository.go
│   │   │   ├── skill_repository.go
│   │   │   └── user_repository.go
│   │   └── value
│   │       ├── authority.go
│   │       ├── role.go
│   │       ├── skill_level.go
│   │       └── skill.go
│   ├── common
│   │   └── errors.go
│   ├── project
│   │   ├── entity
│   │   │   ├── assignment.go
│   │   │   ├── comment.go
│   │   │   ├── project_milestone.go
│   │   │   ├── project.go
│   │   │   └── weekly_schedule_evaluation.go
│   │   ├── errors.go
│   │   ├── repository
│   │   │   ├── assignment_repository.go
│   │   │   ├── comment_repository.go
│   │   │   ├── milestone_repository.go
│   │   │   ├── project_category_repository.go
│   │   │   ├── project_milestone_repository.go
│   │   │   ├── project_repository.go
│   │   │   └── weekly_schedule_evaluation_repository.go
│   │   ├── service
│   │   │   └── project_service.go
│   │   └── value
│   │       ├── assignment_status.go
│   │       ├── evaluation_score.go
│   │       ├── milestone.go
│   │       ├── project_category.go
│   │       └── project_status.go
│   ├── request
│   │   ├── entity
│   │   │   └── request.go
│   │   ├── errors.go
│   │   ├── repository
│   │   │   ├── document_type_repository.go
│   │   │   └── request_repository.go
│   │   └── value
│   │       ├── document_type.go
│   │       └── request_status.go
│   └── workload
│       ├── entity
│       │   ├── weekly_availability.go
│       │   └── weekly_demanded.go
│       ├── errors.go
│       ├── repository
│       │   ├── weekly_availability_repository.go
│       │   └── weekly_demanded_repository.go
│       └── value
│           └── working_hours.go
├── infrastructure
│   ├── config
│   │   ├── app_config.go
│   │   └── server.go
│   ├── database
│   │   └── database.go
│   ├── gorm
│   │   ├── model
│   │   │   ├── assignment.go
│   │   │   ├── authority.go
│   │   │   ├── comment.go
│   │   │   ├── document_type.go
│   │   │   ├── milestone.go
│   │   │   ├── project_category.go
│   │   │   ├── project_milestone.go
│   │   │   ├── project.go
│   │   │   ├── request.go
│   │   │   ├── role.go
│   │   │   ├── skill.go
│   │   │   ├── user_role.go
│   │   │   ├── user_skill.go
│   │   │   ├── user.go
│   │   │   ├── weekly_availability.go
│   │   │   ├── weekly_demanded.go
│   │   │   └── weekly_schedule_evaluation.go
│   │   └── repository
│   │       ├── assignment_repository.go
│   │       ├── authority_repository.go
│   │       ├── base_repository.go
│   │       ├── comment_repository.go
│   │       ├── document_type_repository.go
│   │       ├── milestone_repository.go
│   │       ├── project_category_repository.go
│   │       ├── project_milestone_repository.go
│   │       ├── project_repository.go
│   │       ├── request_repository_ext.go
│   │       ├── request_repository.go
│   │       ├── role_repository.go
│   │       ├── skill_repository.go
│   │       ├── user_repository.go
│   │       ├── weekly_availability_repository.go
│   │       ├── weekly_demanded_repository.go
│   │       └── weekly_schedule_evaluation_repository.go
│   └── jwt
│       └── token_service.go
└── interface
    ├── controller
    │   ├── auth_controller.go
    │   ├── authority_controller.go
    │   ├── document_type_controller.go
    │   ├── evaluation_controller.go
    │   ├── helper.go
    │   ├── milestone_controller.go
    │   ├── project_controller.go
    │   ├── request_controller.go
    │   ├── role_controller.go
    │   ├── skill_controller.go
    │   ├── user_controller.go
    │   ├── weekly_availability_controller.go
    │   └── weekly_demanded_controller.go
    ├── middleware
    │   ├── auth.go
    │   ├── cors.go
    │   ├── error.go
    │   ├── logger.go
    │   └── permission.go
    └── router
        ├── auth_routes.go
        ├── master_routes.go
        ├── project_routes.go
        ├── request_routes.go
        ├── router.go
        ├── user_routes.go
        └── workload_routes.go

54 directories, 153 files