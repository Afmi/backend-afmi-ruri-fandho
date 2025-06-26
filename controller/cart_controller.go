package controller

import (
	"errors"
	"fmt"
	"net/http"
	"user-test/helpers"
	models "user-test/models"
	service "user-test/services"
	"user-test/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartController struct {
	service.CartServices
}

func (ctrl *CartController) GetBasic(ctx *gin.Context) {
	id := ctx.Param("id_user")

	data, err := ctrl.GetCartByUser(id)
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

func (ctrl *CartController) CreateData(ctx *gin.Context) {
	var cart models.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.CreateCart(cart)
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

func (ctrl *CartController) UpdateData(ctx *gin.Context) {
	var cart models.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.UpdateCart(cart)
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
		Data:    &cart,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (ctrl *CartController) DeleteData(ctx *gin.Context) {
	id := ctx.Param("id")

	err := ctrl.DeleteCart(id)
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
