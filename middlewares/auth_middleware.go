package middlewares

import (
	"go-pos/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) ValidateToken(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(401, gin.H{"error": "Authorization token missing"})
		ctx.Abort()
		return
	}

	token, err := utils.ValidateToken(tokenString)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		ctx.JSON(401, gin.H{"error": "Invalid token claims"})
		ctx.Abort()
		return
	}

	// Now you can access the user_id from the claims
	userID := claims.Subject
	ctx.Set("user_id", userID)
	ctx.Next()
}
