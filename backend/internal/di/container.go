package di

import (
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/application/usecase"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/domain/user/service"
	"ghoona-camp-backend/internal/infrastructure/clerk"
	"ghoona-camp-backend/internal/infrastructure/config"
	"ghoona-camp-backend/internal/infrastructure/discord"
	gormRepo "ghoona-camp-backend/internal/infrastructure/gorm/repository"
	"ghoona-camp-backend/internal/interface/controller"
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

	// Services/UseCases
	UserUseCase usecase.UserUseCase

	// Controllers
	UserController *controller.UserController

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

	// Initialize services
	c.initUserServices()

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

	// UseCase
	c.UserUseCase = usecase.NewUserUseCase(
		c.UserRepo,
		c.UserMetadataRepo,
		c.SocialLinkRepo,
		c.RivalRepo,
		userService,
		rivalService,
		validationService,
		c.TxManager,
	)
}

func (c *Container) initUserControllers() {
	c.UserController = controller.NewUserController(c.UserUseCase)
}

// TODO: 他のドメインで実装予定
// func (c *Container) initAttendanceComponents() {
//     c.AttendanceRepo = repository.NewAttendanceRepository(c.DB)
//     c.AttendanceUseCase = usecase.NewAttendanceUseCase(c.AttendanceRepo, c.TxManager)
//     c.AttendanceController = controller.NewAttendanceController(c.AttendanceUseCase)
// }