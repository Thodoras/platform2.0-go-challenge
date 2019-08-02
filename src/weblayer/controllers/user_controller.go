package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/helpers/responseutils"
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/servicelayer/services"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	id, err := services.SignUp(user)
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

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	token, err := services.Login(user)
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

	responseutils.SendSuccess(w, token)
}
