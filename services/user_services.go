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

type UserServices struct {
	repository.SQLRepository
}

func (srvc *UserServices) GetUser(request helpers.Payload) (interface{}, int, int, error) {
	var model models.User
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

func (srvs *UserServices) GetUserById(id string) (interface{}, error) {
	var model models.User
	// preloads = []string{""}

	data, err := srvs.GetOne(&model, "id", id)
	if err != nil {
		log.Printf("Error retrieving from database: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("data not found")
		}
		return nil, fmt.Errorf("error retrieving data: %w", err)
	}

	user, ok := data.(*models.User)
	if !ok {
		return nil, errors.New("invalid type assertion for Loanapplication")
	}

	return user, nil
}
