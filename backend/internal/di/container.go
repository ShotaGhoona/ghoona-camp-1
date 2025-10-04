package di

import (
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/application/transaction"
	userUsecase "ghoona-camp-backend/internal/application/usecase/user"
	titleUsecase "ghoona-camp-backend/internal/application/usecase/title"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/domain/user/service"
	titleRepository "ghoona-camp-backend/internal/domain/title/repository"
	titleService "ghoona-camp-backend/internal/domain/title/service"
	"ghoona-camp-backend/internal/infrastructure/clerk"
	"ghoona-camp-backend/internal/infrastructure/config"
	"ghoona-camp-backend/internal/infrastructure/discord"
	gormRepo "ghoona-camp-backend/internal/infrastructure/gorm/repository"
	userController "ghoona-camp-backend/internal/interface/controller/user"
)

// Container holds all dependencies
type Container struct {
	// Database
	DB *gorm.DB

	// Configuration
	Config *config.Config

	// Transaction Manager
	TxManager transaction.Manager

	// External Services
	ClerkService   *clerk.AuthService
	DiscordService *discord.WebhookService

	// BE-03-user-04で実装済み
	// Repositories - using domain interfaces
	UserRepo         repository.UserRepository
	UserMetadataRepo repository.UserMetadataRepository
	SocialLinkRepo   repository.UserSocialLinkRepository
	RivalRepo        repository.UserRivalRepository

	// Title Repositories
	TitleRepo            titleRepository.TitleRepository
	TitleAchievementRepo titleRepository.TitleAchievementRepository

	// Services/UseCases
	UserUseCase         userUsecase.UserUseCase
	UserMetadataUseCase userUsecase.UserMetadataUseCase
	UserSocialUseCase   userUsecase.UserSocialUseCase
	UserRivalUseCase    userUsecase.UserRivalUseCase

	// Title Services
	TitleService *titleService.TitleService

	// Title UseCases
	TitleUseCase            titleUsecase.TitleUseCase
	TitleAchievementUseCase titleUsecase.TitleAchievementUseCase
	TitleProgressUseCase    titleUsecase.TitleProgressUseCase

	// Controllers
	UserController         *userController.UserController
	UserMetadataController *userController.UserMetadataController
	UserSocialController   *userController.UserSocialController
	UserRivalController    *userController.UserRivalController

	// TODO: 他のドメインで追加予定
	// AttendanceRepo   attendanceRepo.AttendanceRepository
	// GoalRepo         goalRepo.GoalRepository
	// EventRepo        eventRepo.EventRepository
	// TitleRepo        titleRepo.TitleRepository
	// NotificationRepo notificationRepo.NotificationRepository
	// AttendanceUseCase   usecase.AttendanceUseCase
	// GoalUseCase         usecase.GoalUseCase
	// EventUseCase        usecase.EventUseCase
	// TitleUseCase        usecase.TitleUseCase
	// NotificationUseCase usecase.NotificationUseCase
	// AttendanceController   *controller.AttendanceController
	// GoalController         *controller.GoalController
	// EventController        *controller.EventController
	// TitleController        *controller.TitleController
	// NotificationController *controller.NotificationController
}

// NewContainer creates and initializes all dependencies
func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
	c := &Container{
		DB:     db,
		Config: cfg,
	}

	// Initialize transaction manager
	c.TxManager = transaction.NewManager(db)

	// Initialize external services
	c.initExternalServices()

	// BE-03-user-04で実装済み
	// Initialize repositories
	c.initUserRepositories()

	// BE-07-title-02で実装済み
	// Initialize title repositories
	c.initTitleRepositories()

	// Initialize services
	c.initUserServices()

	// Initialize title services
	c.initTitleServices()

	// Initialize title use cases
	c.initTitleUseCases()

	// Initialize controllers
	c.initUserControllers()

	// TODO: 他のドメインで実装予定
	// c.initAttendanceComponents()
	// c.initGoalComponents()
	// c.initEventComponents()
	// c.initTitleComponents()
	// c.initNotificationComponents()

	return c
}

func (c *Container) initExternalServices() {
	// TODO: BE-02-arch-03でClerk設定追加時に実装
	c.ClerkService = clerk.NewAuthService(c.Config)
	
	// TODO: BE-05-discord-*でDiscord設定追加時に実装
	c.DiscordService = discord.NewWebhookService(c.Config)
}

// BE-03-user-04で実装済み
func (c *Container) initUserRepositories() {
	c.UserRepo = gormRepo.NewUserRepository(c.DB)
	c.UserMetadataRepo = gormRepo.NewUserMetadataRepository(c.DB)
	c.SocialLinkRepo = gormRepo.NewUserSocialLinkRepository(c.DB)
	c.RivalRepo = gormRepo.NewUserRivalRepository(c.DB)
}

func (c *Container) initUserServices() {
	// Domain services
	userService := service.NewUserService()
	rivalService := service.NewRivalService(c.RivalRepo)
	validationService := service.NewUserValidationService(c.SocialLinkRepo)

	// UseCases
	c.UserUseCase = userUsecase.NewUserUseCase(
		c.UserRepo,
		c.UserMetadataRepo,
		c.SocialLinkRepo,
		c.RivalRepo,
		userService,
		validationService,
		c.TxManager,
	)
	
	c.UserMetadataUseCase = userUsecase.NewUserMetadataUseCase(
		c.UserRepo,
		c.UserMetadataRepo,
		validationService,
		c.TxManager,
	)
	
	c.UserSocialUseCase = userUsecase.NewUserSocialUseCase(
		c.UserRepo,
		c.SocialLinkRepo,
		userService,
		validationService,
		c.TxManager,
	)
	
	c.UserRivalUseCase = userUsecase.NewUserRivalUseCase(
		c.UserRepo,
		c.RivalRepo,
		rivalService,
		validationService,
		c.TxManager,
	)
}

func (c *Container) initUserControllers() {
	c.UserController = userController.NewUserController(c.UserUseCase)
	c.UserMetadataController = userController.NewUserMetadataController(c.UserMetadataUseCase, c.UserRepo)
	c.UserSocialController = userController.NewUserSocialController(c.UserSocialUseCase, c.UserRepo)
	c.UserRivalController = userController.NewUserRivalController(c.UserRivalUseCase, c.UserRepo)
}

// BE-07-title-02で実装済み
func (c *Container) initTitleRepositories() {
	c.TitleRepo = gormRepo.NewTitleRepository(c.DB)
	c.TitleAchievementRepo = gormRepo.NewTitleAchievementRepository(c.DB)
}

func (c *Container) initTitleServices() {
	// Domain services
	c.TitleService = titleService.NewTitleService()
}

func (c *Container) initTitleUseCases() {
	// Title validation service
	titleValidationService := titleService.NewTitleValidationService(
		c.TitleRepo,
		c.TitleAchievementRepo,
	)

	// Basic title operations
	c.TitleUseCase = titleUsecase.NewTitleUseCase(
		c.TitleRepo,
		titleValidationService,
	)

	// Title achievement operations
	c.TitleAchievementUseCase = titleUsecase.NewTitleAchievementUseCase(
		c.UserRepo,
		c.TitleRepo,
		c.TitleAchievementRepo,
		c.TitleService,
		titleValidationService,
		c.TxManager,
	)

	// Title progress operations (with nil attendance provider for now)
	c.TitleProgressUseCase = titleUsecase.NewTitleProgressUseCase(
		c.UserRepo,
		c.TitleRepo,
		c.TitleAchievementRepo,
		c.TitleService,
		nil, // 将来の出席サービス統合まではnil
	)
}

// TODO: 他のドメインで実装予定
// func (c *Container) initAttendanceComponents() {
//     c.AttendanceRepo = repository.NewAttendanceRepository(c.DB)
//     c.AttendanceUseCase = usecase.NewAttendanceUseCase(c.AttendanceRepo, c.TxManager)
//     c.AttendanceController = controller.NewAttendanceController(c.AttendanceUseCase)
// }