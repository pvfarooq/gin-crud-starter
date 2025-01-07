package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pvfarooq/go-gin-crud/helpers"
	"github.com/pvfarooq/go-gin-crud/models"
	"github.com/pvfarooq/go-gin-crud/services"
)

type UserController struct {
	userService *services.UserService
}

type createUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var req createUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validation_errors, ok := err.(validator.ValidationErrors)

		if ok {
			// handle validation errors
			var errorMessages []string
			for _, e := range validation_errors {
				errorMessages = append(errorMessages, helpers.FormatValidationError(e))
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": errorMessages})
		} else {
			// handle generic errors
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	if err := ctrl.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (ctrl *UserController) FetchUsers(c *gin.Context) {
	users, err := ctrl.userService.FetchUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)

}

func (ctrl *UserController) FetchUser(c *gin.Context) {
	user, err := ctrl.userService.FetchUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	user, err := ctrl.userService.FetchUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		validation_errors, ok := err.(validator.ValidationErrors)

		if ok {
			// handle validation errors
			var errorMessages []string
			for _, e := range validation_errors {
				errorMessages = append(errorMessages, helpers.FormatValidationError(e))
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": errorMessages})
		} else {
			// handle generic errors
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	if err := ctrl.userService.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	user, err := ctrl.userService.FetchUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.userService.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
