package routers

import (
	"net/http"
	routersGroup "user-test/routers/groups"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	apiVersion := "/api"
	routersGroup.AuthRoutes(route, apiVersion)
	routersGroup.UserRoutes(route, apiVersion)
	routersGroup.CartRoutes(route, apiVersion)
	routersGroup.ProductRoutes(route, apiVersion)
}
