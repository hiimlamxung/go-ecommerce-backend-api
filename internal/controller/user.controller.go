package controller

import (
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/hiimlamxung/go-ecommerce-backend-api/internal/services"
	"github.com/hiimlamxung/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	userInfo := uc.userService.GetInfoUser(id)
	response.ResponseWithMessage(c, http.StatusOK, "Get user info successfully", userInfo)
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var createUserRequest CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	spew.Dump(createUserRequest)

	c.JSON(http.StatusOK, gin.H{
		"username": createUserRequest.Username,
		"password": createUserRequest.Password,
		"email":    createUserRequest.Email,
	})

}
