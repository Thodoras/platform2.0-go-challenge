package security

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"platform2.0-go-challenge/src/helpers/responseutils"
	"platform2.0-go-challenge/src/models"
)

const unauthorizedAccessMessage = "unauthorized access."

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

func Authorize(endpoint http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["user_id"]

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return []byte(encryptionKey + id), nil
			})

			if err != nil {
				responseutils.SendError(w, http.StatusUnauthorized, errors.New(unauthorizedAccessMessage))
				return
			}

			if token.Valid {
				endpoint.ServeHTTP(w, r)
			} else {
				responseutils.SendError(w, http.StatusUnauthorized, errors.New(unauthorizedAccessMessage))
				return
			}
		} else {
			responseutils.SendError(w, http.StatusUnauthorized, errors.New(unauthorizedAccessMessage))
			return
		}
	})
}
