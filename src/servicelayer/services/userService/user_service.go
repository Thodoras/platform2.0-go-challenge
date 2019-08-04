package userService

import (
	"golang.org/x/crypto/bcrypt"
	"platform2.0-go-challenge/src/datalayer/repositories/userRepository"
	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/helpers/security"
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/servicelayer/validators/userValidator"
	"platform2.0-go-challenge/src/weblayer/dtos"
)

type UserService struct {
	userRepository userRepository.IUserRepository
	userValidator  userValidator.IUserValidator
}

func NewUserService(
	userRepository userRepository.IUserRepository,
	userValidator userValidator.IUserValidator,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		userValidator:  userValidator,
	}
}

func (u *UserService) SignUp(user models.User) (int, error) {
	err := u.userValidator.ValidateUser(user)
	if err != nil {
		return 0, err
	}

	encryptedPassword, err := u.encrypt(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = encryptedPassword
	return u.userRepository.AddUser(user)
}

func (u *UserService) Login(user models.User) (*dtos.LoginResponse, error) {
	err := u.userValidator.ValidateLoginCredentials(user)
	if err != nil {
		return nil, err
	}

	dbUser, err := u.userRepository.GetUserByName(user.Name)
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

func (u *UserService) encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), err
}
