package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pvfarooq/go-gin-crud/api/controllers"
	"github.com/pvfarooq/go-gin-crud/api/routes"
	"github.com/pvfarooq/go-gin-crud/config"
	"github.com/pvfarooq/go-gin-crud/database"
	"github.com/pvfarooq/go-gin-crud/repositories"
	"github.com/pvfarooq/go-gin-crud/services"
)

func main() {
	// Load configurations
	cfg := config.LoadConfig()

	// Initialize database
	db := database.InitDB(cfg)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userService)

	// Setup Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router, userController)

	// Start server (port loaded from config)
	router.Run()
}
