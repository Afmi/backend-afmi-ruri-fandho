package utils

import (
	"net/http"
	"user-test/helpers"

	"github.com/gin-gonic/gin"
)

func StatusOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Data:    &data,
		Message: "Success",
	})
}

func StatusBadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, helpers.Response{
		Code:    http.StatusBadRequest,
		Status:  true,
		Message: message,
	})
}
func StatusNotFound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, helpers.Response{
		Code:    http.StatusNotFound,
		Status:  true,
		Message: message,
	})
}
func StatusBadGateway(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadGateway, helpers.Response{
		Code:    http.StatusBadGateway,
		Status:  true,
		Message: message,
	})
}
func StatusInternalServerError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, helpers.Response{
		Code:    http.StatusInternalServerError,
		Status:  true,
		Message: message,
	})
}
