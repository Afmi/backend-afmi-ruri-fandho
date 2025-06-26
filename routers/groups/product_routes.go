package group

import (
	"user-test/controller"
	"user-test/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controller.ProductController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.Use(middleware.JWTAuthMiddleware())
	groupRoutes.GET("/product/", ctrl.GetData)
	groupRoutes.POST("/product/", ctrl.CreateData)
	groupRoutes.PUT("/product/", ctrl.UpdateData)
	groupRoutes.DELETE("/product/:id", ctrl.DeleteData)
}
