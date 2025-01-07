package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pvfarooq/go-gin-crud/api/controllers"
)

func RegisterUserRoutes(router *gin.RouterGroup, controller *controllers.UserController) {
	router.GET("users", controller.FetchUsers)
	router.POST("users/register", controller.CreateUser)
	router.GET("users/:id", controller.FetchUser)
	router.PATCH("users/:id", controller.UpdateUser)
	router.DELETE("users/:id", controller.DeleteUser)
}
