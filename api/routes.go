package api

import (
	"github.com/gin-gonic/gin"
	"github.com/smartbot/auth/api/auth"
)

func RegisterRoutes() *gin.Engine {
	router := gin.New()
	authGroup := router.Group("/auth/api/v1")
	auth.RegisterRoutes(authGroup)

	return router

}
