package controllers

import (
	"booklist/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"platform2.0-go-challenge/helpers/errorutils"
	"platform2.0-go-challenge/models"

	"platform2.0-go-challenge/helpers/logutils"

	"github.com/gorilla/mux"
	"platform2.0-go-challenge/helpers/reponseutils"
	"platform2.0-go-challenge/servicelayer/services"
)

func GetAllAssets(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: "Bad request"})
		log.Println(err)
		return
	}
	response, err := services.GetAllAssets(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendError(w, http.StatusNotFound, utils.Error{Message: "Not Found"})
			utils.LogError(err)
			return
		}
		utils.SendError(w, http.StatusInternalServerError, utils.Error{Message: "Server Error"})
		utils.LogError(err)
		return
	}
	utils.SendSuccess(w, response)
}

func AddAudience(w http.ResponseWriter, r *http.Request) {
	var audience models.Audience
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
		if err == errorutils.InvalidRequest {
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

func AddChart(w http.ResponseWriter, r *http.Request) {
	var chart models.Chart
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
		if err == errorutils.InvalidRequest {
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

func AddInsight(w http.ResponseWriter, r *http.Request) {
	var insight models.Insight
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
		if err == errorutils.InvalidRequest {
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

func EditAudience(w http.ResponseWriter, r *http.Request) {
	var audience models.Audience
	json.NewDecoder(r.Body).Decode(&audience)
	rowsAffected, err := services.EditAudience(audience)
	if err != nil {
		if err == errorutils.InvalidRequest {
			reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: err.Error()})
			log.Println(err)
			return
		}
		utils.SendError(w, http.StatusInternalServerError, utils.Error{Message: "Server Error"})
		utils.LogError(err)
		return
	}
	utils.SendSuccess(w, rowsAffected)
}

func EditChart(w http.ResponseWriter, r *http.Request) {
	var chart models.Chart
	json.NewDecoder(r.Body).Decode(&chart)
	rowsAffected, err := services.EditChart(chart)
	if err != nil {
		if err == errorutils.InvalidRequest {
			reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: err.Error()})
			log.Println(err)
			return
		}
		utils.SendError(w, http.StatusInternalServerError, utils.Error{Message: "Server Error"})
		utils.LogError(err)
		return
	}
	utils.SendSuccess(w, rowsAffected)
}

func EditInsight(w http.ResponseWriter, r *http.Request) {
	var insight models.Insight
	json.NewDecoder(r.Body).Decode(&insight)
	rowsAffected, err := services.EditInsight(insight)
	if err != nil {
		if err == errorutils.InvalidRequest {
			reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: err.Error()})
			log.Println(err)
			return
		}
		utils.SendError(w, http.StatusInternalServerError, utils.Error{Message: "Server Error"})
		utils.LogError(err)
		return
	}
	utils.SendSuccess(w, rowsAffected)
}

func DeleteAudience(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: "Bad request"})
		log.Println(err)
		return
	}
	rowsDeleted, err := services.DeleteAudience(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, utils.Error{Message: "Server Error"})
		utils.LogError(err)
		return
	}
	utils.SendSuccess(w, rowsDeleted)
}

func DeleteChart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: "Bad request"})
		log.Println(err)
		return
	}
	rowsDeleted, err := services.DeleteChart(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, utils.Error{Message: "Server Error"})
		utils.LogError(err)
		return
	}
	utils.SendSuccess(w, rowsDeleted)
}

func DeleteInsight(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		reponseutils.SendError(w, http.StatusBadRequest, logutils.Error{Message: "Bad request"})
		log.Println(err)
		return
	}
	rowsDeleted, err := services.DeleteInsight(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, utils.Error{Message: "Server Error"})
		utils.LogError(err)
		return
	}
	utils.SendSuccess(w, rowsDeleted)
}
