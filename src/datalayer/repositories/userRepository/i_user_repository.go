package userRepository

import "platform2.0-go-challenge/src/models"

type IUserRepository interface {
	GetUserByName(name string) (*models.User, error)
	AddUser(user models.User) (int, error)
}
