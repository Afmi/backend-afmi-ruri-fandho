package controller

import (
	"errors"
	"fmt"
	"net/http"
	"user-test/helpers"
	models "user-test/models"
	"user-test/repository"
	service "user-test/services"
	"user-test/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ProductController struct {
	repository.SQLRepository
	service.ProductServices
}

func (ctrl *ProductController) GetData(ctx *gin.Context) {
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

	data, totalCount, totalPages, err := ctrl.GetProduct(request)
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

func (ctrl *ProductController) CreateData(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.CreateProduct(product)
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

func (ctrl *ProductController) UpdateData(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.UpdateProduct(product)
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
		Data:    &product,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (ctrl *ProductController) DeleteData(ctx *gin.Context) {
	id := ctx.Param("id")

	err := ctrl.DeleteProduct(id)
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
	}
	ctx.JSON(http.StatusOK, webResponse)
}
