package services

import (
	"golang.org/x/crypto/bcrypt"
	"platform2.0-go-challenge/src/datalayer/repositories"
	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/helpers/security"
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/servicelayer/validators"
	"platform2.0-go-challenge/src/weblayer/dtos"
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

func Login(user models.User) (*dtos.LoginResponse, error) {
	err := validators.ValidateLoginCredentials(user)
	if err != nil {
		return nil, err
	}

	dbUser, err := repositories.GetUserByName(user.Name)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return nil, errorutils.NewInvalidRequest("Invalid credentials.")
	}

	token, err := security.GenerateJWT(*dbUser)
	if err != nil {
		return nil, nil
	}
	result := &dtos.LoginResponse{ID: dbUser.ID, Token: token}
	return result, nil
}

func encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), err
}
