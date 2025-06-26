package group

import (
	"user-test/controller"
	"user-test/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controller.UserController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.Use(middleware.JWTAuthMiddleware())
	groupRoutes.POST("/user/", ctrl.GetBasic)
	groupRoutes.GET("/user/:id", ctrl.GetById)
}
