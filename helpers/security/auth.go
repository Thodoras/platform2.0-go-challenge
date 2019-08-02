package security

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"platform2.0-go-challenge/models"
)

var encryptionKey = "someKey"

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
	})

	tokenString, err := token.SignedString([]byte(encryptionKey + strconv.Itoa(user.ID)))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
