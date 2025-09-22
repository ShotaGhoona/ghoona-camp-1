package di

import (
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/infrastructure/clerk"
	"ghoona-camp-backend/internal/infrastructure/config"
	"ghoona-camp-backend/internal/infrastructure/discord"
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

	// TODO: BE-03-* で追加予定
	// Repositories - using domain interfaces
	// UserRepo         userRepo.UserRepository
	// AttendanceRepo   attendanceRepo.AttendanceRepository
	// GoalRepo         goalRepo.GoalRepository
	// EventRepo        eventRepo.EventRepository
	// TitleRepo        titleRepo.TitleRepository
	// NotificationRepo notificationRepo.NotificationRepository

	// Services/UseCases
	// UserUseCase         usecase.UserUseCase
	// AttendanceUseCase   usecase.AttendanceUseCase
	// GoalUseCase         usecase.GoalUseCase
	// EventUseCase        usecase.EventUseCase
	// TitleUseCase        usecase.TitleUseCase
	// NotificationUseCase usecase.NotificationUseCase

	// Controllers
	// UserController         *controller.UserController
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

	// TODO: BE-03-* で実装予定
	// Initialize repositories
	// c.initRepositories()

	// Initialize services
	// c.initServices()

	// Initialize controllers
	// c.initControllers()

	return c
}

func (c *Container) initExternalServices() {
	// TODO: BE-02-arch-03でClerk設定追加時に実装
	c.ClerkService = clerk.NewAuthService(c.Config)
	
	// TODO: BE-05-discord-*でDiscord設定追加時に実装
	c.DiscordService = discord.NewWebhookService(c.Config)
}

// TODO: BE-03-* で実装予定
// func (c *Container) initRepositories() {
//     c.UserRepo = repository.NewUserRepository(c.DB)
//     c.AttendanceRepo = repository.NewAttendanceRepository(c.DB)
//     c.GoalRepo = repository.NewGoalRepository(c.DB)
//     c.EventRepo = repository.NewEventRepository(c.DB)
//     c.TitleRepo = repository.NewTitleRepository(c.DB)
//     c.NotificationRepo = repository.NewNotificationRepository(c.DB)
// }

// func (c *Container) initServices() {
//     c.UserUseCase = usecase.NewUserUseCase(c.UserRepo, c.TxManager)
//     c.AttendanceUseCase = usecase.NewAttendanceUseCase(c.AttendanceRepo, c.TxManager)
//     c.GoalUseCase = usecase.NewGoalUseCase(c.GoalRepo, c.TxManager)
//     c.EventUseCase = usecase.NewEventUseCase(c.EventRepo, c.TxManager)
//     c.TitleUseCase = usecase.NewTitleUseCase(c.TitleRepo, c.TxManager)
//     c.NotificationUseCase = usecase.NewNotificationUseCase(c.NotificationRepo, c.TxManager)
// }

// func (c *Container) initControllers() {
//     c.UserController = controller.NewUserController(c.UserUseCase)
//     c.AttendanceController = controller.NewAttendanceController(c.AttendanceUseCase)
//     c.GoalController = controller.NewGoalController(c.GoalUseCase)
//     c.EventController = controller.NewEventController(c.EventUseCase)
//     c.TitleController = controller.NewTitleController(c.TitleUseCase)
//     c.NotificationController = controller.NewNotificationController(c.NotificationUseCase)
// }