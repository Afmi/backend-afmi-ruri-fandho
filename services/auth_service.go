package service

import (
	"errors"
	"os"
	"time"
	models "user-test/models"
	"user-test/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthServices struct {
	repository.SQLRepository
}

func (svc *AuthServices) RegisterUser(user *models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	_, err = svc.Save(user, false)
	return err
}

func (svc *AuthServices) LoginUser(email, password string) (*models.User, error) {
	var user models.User
	_, err := svc.GetOne(&user, "email", email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

var refreshSecret = []byte("another-secret-key") // bisa juga dari .env

func (svc *AuthServices) GenerateAccessToken(userId string) (string, error) {
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})

	tokenString, err := access.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (svc *AuthServices) ValidateAndRefresh(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	userID := int(claims["user_id"].(float64))
	newAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})
	return newAccess.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (svc *AuthServices) GenerateAndSaveRefreshToken(userId string) (string, error) {
	// Buat refresh token
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	refreshTokenStr, err := refresh.SignedString(refreshSecret)
	if err != nil {
		return "", err
	}

	// Simpan ke database
	rt := &models.RefreshToken{
		Token:     refreshTokenStr,
		UserId:    userId,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	_, err = svc.Save(rt, false)
	if err != nil {
		return "", err
	}

	return refreshTokenStr, nil
}
