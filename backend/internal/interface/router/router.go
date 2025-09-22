package router

import (
	"github.com/gin-gonic/gin"
	"ghoona-camp-backend/internal/di"
	"ghoona-camp-backend/internal/interface/middleware"
)

// Router handles all route configurations
type Router struct {
	engine    *gin.Engine
	container *di.Container
}

// NewRouter creates a new router instance
func NewRouter(container *di.Container) *Router {
	engine := gin.Default()

	// Apply global middleware
	engine.Use(middleware.LoggerMiddleware())
	engine.Use(middleware.CORSMiddleware())
	engine.Use(middleware.SecurityHeadersMiddleware())
	engine.Use(middleware.ErrorHandlerMiddleware())

	return &Router{
		engine:    engine,
		container: container,
	}
}

// Setup configures all routes
func (r *Router) Setup() *gin.Engine {
	// Health check endpoint (simple string response)
	r.engine.GET("/health", r.healthCheck)

	// Metrics endpoint
	r.engine.GET("/metrics", r.metrics)

	// API v1 routes
	v1 := r.engine.Group("/api/v1")
	{
		v1.GET("/ping", r.ping)
		
		// TODO: BE-03-* で各ドメインのルート追加
		// r.setupUserRoutes(v1)
		// r.setupAttendanceRoutes(v1)
		// r.setupGoalRoutes(v1)
		// r.setupEventRoutes(v1)
		// r.setupTitleRoutes(v1)
		// r.setupNotificationRoutes(v1)
	}

	// 404/405 handlers
	r.engine.NoRoute(middleware.NotFoundHandler())
	r.engine.NoMethod(middleware.MethodNotAllowedHandler())

	return r.engine
}

// healthCheck handles health check requests
func (r *Router) healthCheck(c *gin.Context) {
	// TODO: BE-02-arch-02でDB接続チェックを追加
	response := gin.H{
		"status":      "ok",
		"service":     "ghoona-camp-backend",
		"version":     "1.0.0",
		"environment": r.container.Config.Env,
	}
	
	// TODO: データベース接続チェック
	// if err := r.container.DB.DB().Ping(); err != nil {
	//     response["status"] = "error"
	//     response["database"] = "disconnected"
	//     c.JSON(503, response)
	//     return
	// }
	// response["database"] = "connected"
	
	c.JSON(200, response)
}

// metrics handles metrics requests
func (r *Router) metrics(c *gin.Context) {
	c.JSON(200, gin.H{
		"metrics": "TODO: implement metrics in BE-05-*",
	})
}

// ping handles ping requests
func (r *Router) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// GetEngine returns the gin engine
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

// TODO: BE-03-user-01で実装
// func (r *Router) setupUserRoutes(v1 *gin.RouterGroup) {
//     auth := middleware.AuthMiddleware(r.container.ClerkService)
//     
//     users := v1.Group("/users")
//     users.Use(auth)
//     {
//         users.GET("", r.container.UserController.GetUsers)
//         users.GET("/:userId", r.container.UserController.GetUser)
//         users.POST("", r.container.UserController.CreateUser)
//         users.PUT("/:userId", r.container.UserController.UpdateUser)
//         users.DELETE("/:userId", r.container.UserController.DeleteUser)
//         
//         // User metadata
//         users.GET("/:userId/metadata", r.container.UserController.GetUserMetadata)
//         users.PUT("/:userId/metadata", r.container.UserController.UpdateUserMetadata)
//         
//         // Social links
//         users.GET("/:userId/social-links", r.container.UserController.GetSocialLinks)
//         users.POST("/:userId/social-links", r.container.UserController.CreateSocialLink)
//         users.PUT("/:userId/social-links/:linkId", r.container.UserController.UpdateSocialLink)
//         users.DELETE("/:userId/social-links/:linkId", r.container.UserController.DeleteSocialLink)
//         
//         // Rivals
//         users.GET("/:userId/rivals", r.container.UserController.GetRivals)
//         users.POST("/:userId/rivals", r.container.UserController.AddRival)
//         users.DELETE("/:userId/rivals/:rivalId", r.container.UserController.DeleteRival)
//     }
//     
//     // Auth routes
//     auth := v1.Group("/auth")
//     {
//         auth.GET("/me", middleware.AuthMiddleware(r.container.ClerkService), r.container.UserController.GetCurrentUser)
//     }
// }

// TODO: BE-03-attendance-01で実装
// func (r *Router) setupAttendanceRoutes(v1 *gin.RouterGroup) {
//     auth := middleware.AuthMiddleware(r.container.ClerkService)
//     
//     attendance := v1.Group("/attendance")
//     attendance.Use(auth)
//     {
//         attendance.GET("/logs", r.container.AttendanceController.GetAttendanceLogs)
//         attendance.POST("/logs", r.container.AttendanceController.RecordAttendance)
//         attendance.GET("/summary", r.container.AttendanceController.GetAttendanceSummary)
//         attendance.GET("/statistics", r.container.AttendanceController.GetAttendanceStatistics)
//         attendance.GET("/rankings", r.container.AttendanceController.GetAttendanceRankings)
//     }
// }

// TODO: BE-03-goal-01で実装
// func (r *Router) setupGoalRoutes(v1 *gin.RouterGroup) {
//     auth := middleware.AuthMiddleware(r.container.ClerkService)
//     optionalAuth := middleware.OptionalAuthMiddleware(r.container.ClerkService)
//     
//     goals := v1.Group("/goals")
//     {
//         goals.GET("/public", optionalAuth, r.container.GoalController.GetPublicGoals)
//         goals.GET("/:goalId", optionalAuth, r.container.GoalController.GetGoal)
//         
//         goals.Use(auth)
//         goals.POST("", r.container.GoalController.CreateGoal)
//         goals.PUT("/:goalId", r.container.GoalController.UpdateGoal)
//         goals.DELETE("/:goalId", r.container.GoalController.DeleteGoal)
//         
//         // Progress
//         goals.GET("/:goalId/progress", r.container.GoalController.GetGoalProgress)
//         goals.POST("/:goalId/progress", r.container.GoalController.RecordProgress)
//     }
//     
//     users := v1.Group("/users")
//     users.Use(auth)
//     {
//         users.GET("/:userId/goals", r.container.GoalController.GetUserGoals)
//         users.GET("/:userId/goals/statistics", r.container.GoalController.GetGoalStatistics)
//     }
// }

// TODO: BE-03-event-01で実装
// func (r *Router) setupEventRoutes(v1 *gin.RouterGroup) {
//     auth := middleware.AuthMiddleware(r.container.ClerkService)
//     
//     events := v1.Group("/events")
//     {
//         events.GET("", auth, r.container.EventController.GetEvents)
//         events.GET("/types", auth, r.container.EventController.GetEventTypes)
//         events.GET("/popular", auth, r.container.EventController.GetPopularEvents)
//         events.GET("/:eventId", auth, r.container.EventController.GetEvent)
//         
//         events.Use(auth)
//         events.POST("", r.container.EventController.CreateEvent)
//         events.PUT("/:eventId", r.container.EventController.UpdateEvent)
//         events.DELETE("/:eventId", r.container.EventController.DeleteEvent)
//         
//         // Participants
//         events.GET("/:eventId/participants", r.container.EventController.GetEventParticipants)
//         events.POST("/:eventId/participants", r.container.EventController.JoinEvent)
//         events.PUT("/:eventId/participants/:userId", r.container.EventController.UpdateParticipationStatus)
//         events.DELETE("/:eventId/participants/:userId", r.container.EventController.LeaveEvent)
//     }
// }

// TODO: BE-03-title-01で実装
// func (r *Router) setupTitleRoutes(v1 *gin.RouterGroup) {
//     auth := middleware.AuthMiddleware(r.container.ClerkService)
//     
//     titles := v1.Group("/titles")
//     titles.Use(auth)
//     {
//         titles.GET("", r.container.TitleController.GetTitles)
//         titles.GET("/:titleId", r.container.TitleController.GetTitle)
//     }
//     
//     users := v1.Group("/users")
//     users.Use(auth)
//     {
//         users.GET("/:userId/titles", r.container.TitleController.GetUserTitles)
//         users.POST("/:userId/titles/:titleId/claim", r.container.TitleController.ClaimTitle)
//     }
// }

// TODO: BE-03-notification-01で実装
// func (r *Router) setupNotificationRoutes(v1 *gin.RouterGroup) {
//     auth := middleware.AuthMiddleware(r.container.ClerkService)
//     
//     notifications := v1.Group("/notifications")
//     notifications.Use(auth)
//     {
//         notifications.GET("", r.container.NotificationController.GetNotifications)
//         notifications.PUT("/:notificationId/read", r.container.NotificationController.MarkAsRead)
//         notifications.PUT("/read-all", r.container.NotificationController.MarkAllAsRead)
//         notifications.GET("/unread-count", r.container.NotificationController.GetUnreadCount)
//     }
//     
//     users := v1.Group("/users")
//     users.Use(auth)
//     {
//         users.GET("/:userId/notification-settings", r.container.NotificationController.GetNotificationSettings)
//         users.PUT("/:userId/notification-settings", r.container.NotificationController.UpdateNotificationSettings)
//     }
// }