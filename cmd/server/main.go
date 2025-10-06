package main

import (
	"log"
	"time"

	"example/go_api_tutorial/internal/config"
	"example/go_api_tutorial/internal/database"
	"example/go_api_tutorial/internal/handler"
	"example/go_api_tutorial/internal/middleware"
	"example/go_api_tutorial/internal/repository/postgres"
	"example/go_api_tutorial/internal/service"
	"example/go_api_tutorial/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Run migrations
	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Seed initial data (for development)
	if err := database.SeedData(); err != nil {
		log.Fatal("Failed to seed data:", err)
	}

	// Initialize dependencies using dependency injection
	db := database.GetDB()
	
	// Initialize JWT manager
	expiresIn, _ := time.ParseDuration(cfg.JWT.ExpiresIn)
	jwtManager := utils.NewJWTManager(cfg.JWT.Secret, expiresIn)
	
	// Initialize repositories
	bookRepo := postgres.NewBookRepository(db)
	userRepo := postgres.NewUserRepository(db)
	
	// Initialize services
	bookService := service.NewBookService(bookRepo)
	userService := service.NewUserService(userRepo)
	
	// Initialize handlers
	bookHandler := handler.NewBookHandler(bookService)
	authHandler := handler.NewAuthHandler(userService, jwtManager)
	userHandler := handler.NewUserHandler(userService)

	// Initialize Gin router
	router := gin.Default()

	// Public routes (no authentication required)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Book Dictionary API is running"})
	})

	// Authentication routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)           // POST /auth/register
		authRoutes.POST("/login", authHandler.Login)                // POST /auth/login
		authRoutes.POST("/refresh", authHandler.RefreshToken)       // POST /auth/refresh
		
		// Protected auth routes (require authentication)
		protected := authRoutes.Group("", middleware.AuthMiddleware(jwtManager))
		{
			protected.GET("/profile", authHandler.GetProfile)           // GET /auth/profile
			protected.POST("/change-password", authHandler.ChangePassword) // POST /auth/change-password
		}
	}

	// Book routes (require authentication)
	bookRoutes := router.Group("/books", middleware.AuthMiddleware(jwtManager))
	{
		bookRoutes.GET("", bookHandler.GetBooks)                    // GET /books (all users)
		bookRoutes.GET("/:id", bookHandler.GetBookByID)             // GET /books/:id (all users)
		
		// Admin-only book routes
		adminBookRoutes := bookRoutes.Group("", middleware.AdminMiddleware())
		{
			adminBookRoutes.POST("", bookHandler.CreateBook)                 // POST /books (admin only)
			adminBookRoutes.PUT("/:id", bookHandler.UpdateBook)              // PUT /books/:id (admin only)
			adminBookRoutes.DELETE("/:id", bookHandler.DeleteBook)           // DELETE /books/:id (admin only)
			adminBookRoutes.PATCH("/:id/quantity", bookHandler.UpdateBookQuantity) // PATCH /books/:id/quantity (admin only)
		}
	}

	// User management routes (admin only)
	userRoutes := router.Group("/users", middleware.AuthMiddleware(jwtManager), middleware.AdminMiddleware())
	{
		userRoutes.GET("", userHandler.GetAllUsers)                 // GET /users (admin only)
		userRoutes.GET("/:id", userHandler.GetUserByID)             // GET /users/:id (admin only)
		userRoutes.PATCH("/:id/role", userHandler.UpdateUserRole)   // PATCH /users/:id/role (admin only)
	}

	// Start server
	log.Printf("Server starting on %s", cfg.GetServerAddress())
	log.Println("Available endpoints:")
	log.Println("  GET    /health")
	log.Println("  POST   /auth/register")
	log.Println("  POST   /auth/login")
	log.Println("  POST   /auth/refresh")
	log.Println("  GET    /auth/profile (auth required)")
	log.Println("  POST   /auth/change-password (auth required)")
	log.Println("  GET    /books (auth required)")
	log.Println("  GET    /books/:id (auth required)")
	log.Println("  POST   /books (admin only)")
	log.Println("  PUT    /books/:id (admin only)")
	log.Println("  DELETE /books/:id (admin only)")
	log.Println("  PATCH  /books/:id/quantity (admin only)")
	log.Println("  GET    /users (admin only)")
	log.Println("  GET    /users/:id (admin only)")
	log.Println("  PATCH  /users/:id/role (admin only)")
	
	if err := router.Run(cfg.GetServerAddress()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}