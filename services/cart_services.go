package service

import (
	"errors"
	"fmt"
	"log"
	models "user-test/models"
	"user-test/repository"

	"gorm.io/gorm"
)

type CartServices struct {
	repository.SQLRepository
}

func (srvs *CartServices) CreateCart(body models.Cart) (interface{}, error) {
	var (
		finalPrice   = body.Price
		freeShipping bool
		discount     string
	)

	// Cek apakah dapat bebas ongkir
	if body.Price > 15000 {
		freeShipping = true
	}

	// Cek apakah dapat diskon
	if body.Price > 50000 {
		discountRate := 0.10
		discount = "10%"
		finalPrice = body.Price - (body.Price * discountRate)
	}

	cardData := &models.Cart{
		Id:           body.Id,
		OwnerId:      body.OwnerId,
		ProductId:    body.ProductId,
		Price:        body.Price,
		TotalPrice:   finalPrice,
		FreeShipping: &freeShipping,
		Discount:     discount,
	}

	data, err := srvs.Save(&cardData, true)
	if err != nil {
		return nil, fmt.Errorf("error saving Cart: %w", err)
	}

	return data, nil
}

func (srvs *CartServices) UpdateCart(body models.Cart) error {
	cardData := models.Cart{
		Id:           body.Id,
		OwnerId:      body.OwnerId,
		ProductId:    body.ProductId,
		TotalPrice:   body.TotalPrice,
		FreeShipping: body.FreeShipping,
		Discount:     body.Discount,
	}

	err := srvs.Updates(&cardData, "id", body.Id)
	if err != nil {
		return fmt.Errorf("error saving Cart: %w", err)
	}

	return nil
}

func (srvs *CartServices) GetCartByUser(id string) (interface{}, error) {
	var model []*models.Cart

	data, err := srvs.GetList(&model, "id", id)
	if err != nil {
		log.Printf("Error retrieving from database: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("data not found")
		}
		return nil, fmt.Errorf("error retrieving data: %w", err)
	}

	cart, ok := data.([]*models.Cart)
	if !ok {
		return nil, errors.New("invalid type assertion for Cart")
	}

	return cart, nil
}

func (srvs *CartServices) DeleteCart(id string) error {
	var model models.Cart

	err := srvs.HardDelete(&model, "id", id)
	if err != nil {
		log.Printf("Error retrieving from database: %v", err)
		return fmt.Errorf("error deleting data: %w", err)
	}

	return nil
}
