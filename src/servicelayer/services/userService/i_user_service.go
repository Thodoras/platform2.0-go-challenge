package userService

import (
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/weblayer/dtos"
)

type IUserService interface {
	SignUp(user models.User) (int, error)
	Login(user models.User) (*dtos.LoginResponse, error)
}
