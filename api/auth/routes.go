package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.RouterGroup) {
	authService := AuthService{}
	authController := AuthController{service: authService}

	route.POST("/signup", authController.Signup)
	route.POST("/login", authController.Login)
	route.POST("/token/refresh", authController.RefreshToken)
	route.GET("/jwks", authController.GetJWKS)
}
