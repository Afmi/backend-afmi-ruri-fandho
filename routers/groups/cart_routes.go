package group

import (
	"user-test/controller"
	"user-test/middleware"

	"github.com/gin-gonic/gin"
)

func CartRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controller.CartController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.Use(middleware.JWTAuthMiddleware())
	groupRoutes.POST("/cart/:id_user", ctrl.GetBasic)
	groupRoutes.POST("/cart/", ctrl.CreateData)
	groupRoutes.PUT("/cart/", ctrl.UpdateData)
	groupRoutes.DELETE("/cart/:id", ctrl.DeleteData)
}
