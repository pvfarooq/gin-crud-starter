package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pvfarooq/go-gin-crud/api/controllers"
)

// RegisterRoutes sets up all API routes for the application.
// It creates a base API group with prefix "/api/v1/" and registers all route sub-groups.
//
// Parameters:
//   - router: A pointer to the gin.Engine instance that will handle the routes
//   - userController: A pointer to the UserController that handles user-related operations
func RegisterRoutes(router *gin.Engine, userController *controllers.UserController) {
	api := router.Group("/api/v1/")

	// Register routes here
	RegisterUserRoutes(api, userController)
}
