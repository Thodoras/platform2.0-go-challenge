package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"platform2.0-go-challenge/helpers/errorutils"

	"platform2.0-go-challenge/helpers/logutils"
	"platform2.0-go-challenge/helpers/reponseutils"
	"platform2.0-go-challenge/servicelayer/services"

	"platform2.0-go-challenge/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	id, err := services.SignUp(user)
	if err != nil {
		if err == errorutils.UniqueConstrainViolation || err == errorutils.InvalidRequest {
			reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: err.Error()})
			log.Println(err)
			return
		}
		reponseutils.SendError(w, http.StatusInternalServerError, logutils.Error{Message: "Server Error"})
		log.Println(err)
		return
	}

	reponseutils.SendSuccess(w, id)
}
