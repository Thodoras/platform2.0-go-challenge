package userController

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/helpers/responseutils"
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/servicelayer/services/userService"
)

type UserController struct {
	userService userService.IUserService
}

func NewUserController(userService userService.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	id, err := u.userService.SignUp(user)
	if err != nil {
		if err == errorutils.UniqueConstrainViolation || err == errorutils.InvalidRequest {
			responseutils.SendError(w, http.StatusBadRequest, err)
			log.Println(err)
			return
		}
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}

	responseutils.SendSuccess(w, id)
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	response, err := u.userService.Login(user)
	if err != nil {
		if err == errorutils.InvalidRequest {
			responseutils.SendError(w, http.StatusBadRequest, err)
			log.Println(err)
			return
		}

		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}

	responseutils.SendSuccess(w, response)
}
