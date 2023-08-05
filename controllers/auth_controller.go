package controllers

import (
	"go-pos/models"
	"go-pos/services"
	"go-pos/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		authService: services.NewAuthService(db),
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := c.authService.Register(&user)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusCreated, "User created successfully", nil)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := c.authService.Login(user.Username, user.Password)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusOK, "Login successful", gin.H{"token": token})
}

func (c *AuthController) ValidateToken(ctx *gin.Context) {
	utils.SendSuccessResponse(ctx, http.StatusOK, "Valid token", nil)
}
