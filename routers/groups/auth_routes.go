package group

import (
	"user-test/controller"
	"user-test/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controller.AuthController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.POST("/register", ctrl.Register)
	groupRoutes.POST("/login", ctrl.Login)
	groupRoutes.POST("/refresh", middleware.JWTAuthMiddleware(), ctrl.RefreshToken)
}
