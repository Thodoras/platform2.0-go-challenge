package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"platform2.0-go-challenge/helpers/logutils"

	"github.com/gorilla/mux"
	"platform2.0-go-challenge/helpers/reponseutils"
	"platform2.0-go-challenge/models/assets"
	"platform2.0-go-challenge/servicelayer/services"
)

func GetAllAssets(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]
	fmt.Println("Hi!" + userID)
}

func AddAudience(w http.ResponseWriter, r *http.Request) {
	var audience assets.Audience
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: "Bad request"})
		log.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&audience)
	audience.UserID = userID
	id, err := services.AddAudience(audience)
	if err != nil {
		reponseutils.SendError(w, http.StatusInternalServerError, logutils.Error{Message: "Server Error"})
		log.Println(err)
		return
	}
	reponseutils.SendSuccess(w, id)
}

func AddChart(w http.ResponseWriter, r *http.Request) {
	var chart assets.Chart
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: "Bad request"})
		log.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&chart)
	chart.UserID = userID
	id, err := services.AddChart(chart)
	if err != nil {
		reponseutils.SendError(w, http.StatusInternalServerError, logutils.Error{Message: "Server Error"})
		log.Println(err)
		return
	}
	reponseutils.SendSuccess(w, id)
}

func AddInsight(w http.ResponseWriter, r *http.Request) {
	var insight assets.Insight
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: "Bad request"})
		log.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&insight)
	insight.UserID = userID
	id, err := services.AddInsight(insight)
	if err != nil {
		reponseutils.SendError(w, http.StatusInternalServerError, logutils.Error{Message: "Server Error"})
		log.Println(err)
		return
	}
	reponseutils.SendSuccess(w, id)
}
