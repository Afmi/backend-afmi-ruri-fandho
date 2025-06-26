package helpers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAuth(auth string) string {

	token, _ := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "invalid token claims"
	}
	user_id := fmt.Sprintf("%v", claims["user_id"])

	return user_id
}

func GetClaims(c *gin.Context) jwt.MapClaims {
	auth := c.GetHeader("Authorization")
	token, _ := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	return claims
}

func GetBranch(c *gin.Context) jwt.MapClaims {
	auth := c.GetHeader("Authorization")
	token, _ := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	return claims
}

func GetBranchId(branch jwt.MapClaims) string {
	branchData, ok := branch["branch"].(map[string]interface{})
	if !ok {

		return "branch ID not found"
	}
	branchId := fmt.Sprintf("%v", branchData["id"])
	return branchId
}
