package controller

import (
	"errors"
	"fmt"
	"net/http"
	"user-test/helpers"
	"user-test/repository"
	service "user-test/services"
	"user-test/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type UserController struct {
	repository.SQLRepository
	service.UserServices
}

func (ctrl *UserController) GetBasic(ctx *gin.Context) {
	var request helpers.Payload
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			utils.StatusBadRequest(ctx, validationErrors.Error())
		} else {
			utils.StatusBadRequest(ctx, "Invalid request payload")
		}
		return
	}

	data, totalCount, totalPages, err := ctrl.GetUser(request)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.StatusNotFound(ctx, "Data not found")
		} else {
			utils.StatusInternalServerError(ctx, fmt.Sprintf("Failed to retrieving data: %v", err))
		}
		return
	}

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
		Info: helpers.InfoPage{
			Page:       request.Page,
			Length:     request.Length,
			TotalPages: totalPages,
			TotalData:  totalCount,
		},
		Data: &data,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
func (ctrl *UserController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := ctrl.GetUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.StatusNotFound(ctx, "Data not found")
		} else {
			utils.StatusInternalServerError(ctx, fmt.Sprintf("Failed to retrieving data: %v", err))
		}
		return
	}

	webResponse := helpers.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
		Data:    &data,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
