package services

import (
	"golang.org/x/crypto/bcrypt"
	"platform2.0-go-challenge/datalayer/repositories"
	"platform2.0-go-challenge/models"
	"platform2.0-go-challenge/servicelayer/validators"
)

func SignUp(user models.User) (int, error) {
	err := validators.ValidateUser(user)
	if err != nil {
		return 0, err
	}

	encryptedPassword, err := encrypt(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = encryptedPassword
	return repositories.AddUser(user)
}

func encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), err
}
