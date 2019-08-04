package userValidator

import "platform2.0-go-challenge/src/models"

type IUserValidator interface {
	ValidateUser(user models.User) error
	ValidateLoginCredentials(user models.User) error
}
