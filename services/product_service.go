package service

import (
	"errors"
	"fmt"
	"log"
	"user-test/helpers"
	models "user-test/models"
	"user-test/repository"

	"gorm.io/gorm"
)

type ProductServices struct {
	repository.SQLRepository
}

func (srvc *ProductServices) GetProduct(request helpers.Payload) (interface{}, int, int, error) {
	var model []*models.Product
	data, totalCount, totalPages, err := srvc.Get(&model, request)
	if err != nil {
		log.Printf("Error retrieving from database: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, 0, fmt.Errorf("data not found")
		}
		return nil, 0, 0, fmt.Errorf("error retrieving data: %w", err)
	}

	return data, totalCount, totalPages, nil
}

func (srvs *ProductServices) CreateProduct(body models.Product) (interface{}, error) {

	productData := &models.Product{
		Id:          body.Id,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
	}

	data, err := srvs.Save(&productData, true)
	if err != nil {
		return nil, fmt.Errorf("error saving product: %w", err)
	}

	return data, nil
}

func (srvs *ProductServices) UpdateProduct(body models.Product) error {
	productData := &models.Product{
		Id:          body.Id,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
	}

	err := srvs.Updates(&productData, "id", body.Id)
	if err != nil {
		return fmt.Errorf("error data product: %w", err)
	}

	return nil
}

func (srvs *ProductServices) DeleteProduct(id string) error {
	var model models.Cart

	err := srvs.HardDelete(&model, "id", id)
	if err != nil {
		log.Printf("Error retrieving from database: %v", err)
		return fmt.Errorf("error deleting data: %w", err)
	}

	return nil
}
