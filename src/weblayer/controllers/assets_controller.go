package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/helpers/responseutils"
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/servicelayer/services"

	"github.com/gorilla/mux"
)

func GetAllAssets(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	response, err := services.GetAllAssets(userID, false)
	if err != nil {
		if err == sql.ErrNoRows {
			responseutils.SendError(w, http.StatusNotFound, errors.New("Not Found"))
			log.Println(err)
			return
		}
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, response)
}

func GetAllAssetsFavourite(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	response, err := services.GetAllAssets(userID, true)
	if err != nil {
		if err == sql.ErrNoRows {
			responseutils.SendError(w, http.StatusNotFound, errors.New("Not Found"))
			log.Println(err)
			return
		}
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, response)
}

func GetAllAssetsPaginated(w http.ResponseWriter, r *http.Request) {
	limit_key, ok := r.URL.Query()["limit"]
	if !ok || len(limit_key) < 1 {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println("url parameter error")
		return
	}

	limit, err := strconv.Atoi(limit_key[0])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}

	offset_key, ok := r.URL.Query()["offset"]
	if !ok || len(offset_key) < 1 {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println("url parameter error")
		return
	}

	offset, err := strconv.Atoi(offset_key[0])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}

	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}

	response, err := services.GetAllAssetsPaginated(userID, limit, offset, false)
	if err != nil {
		if err == sql.ErrNoRows {
			responseutils.SendError(w, http.StatusNotFound, errors.New("Not Found"))
			log.Println(err)
			return
		}
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, response)
}

func GetAllAssetsPaginatedFavourite(w http.ResponseWriter, r *http.Request) {
	limit_key, ok := r.URL.Query()["limit"]
	if !ok || len(limit_key) < 1 {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println("url parameter error")
		return
	}

	limit, err := strconv.Atoi(limit_key[0])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}

	offset_key, ok := r.URL.Query()["offset"]
	if !ok || len(offset_key) < 1 {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println("url parameter error")
		return
	}

	offset, err := strconv.Atoi(offset_key[0])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}

	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}

	response, err := services.GetAllAssetsPaginated(userID, limit, offset, true)
	if err != nil {
		if err == sql.ErrNoRows {
			responseutils.SendError(w, http.StatusNotFound, errors.New("Not Found"))
			log.Println(err)
			return
		}
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, response)
}

func AddAudience(w http.ResponseWriter, r *http.Request) {
	var audience models.Audience
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&audience)
	audience.UserID = userID
	id, err := services.AddAudience(audience)
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
	responseutils.SendSuccess(w, id)
}

func AddChart(w http.ResponseWriter, r *http.Request) {
	var chart models.Chart
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&chart)
	chart.UserID = userID
	id, err := services.AddChart(chart)
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
	responseutils.SendSuccess(w, id)
}

func AddInsight(w http.ResponseWriter, r *http.Request) {
	var insight models.Insight
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&insight)
	insight.UserID = userID
	id, err := services.AddInsight(insight)
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
	responseutils.SendSuccess(w, id)
}

func EditAudience(w http.ResponseWriter, r *http.Request) {
	var audience models.Audience
	json.NewDecoder(r.Body).Decode(&audience)
	rowsAffected, err := services.EditAudience(audience)
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
	responseutils.SendSuccess(w, rowsAffected)
}

func EditChart(w http.ResponseWriter, r *http.Request) {
	var chart models.Chart
	json.NewDecoder(r.Body).Decode(&chart)
	rowsAffected, err := services.EditChart(chart)
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
	responseutils.SendSuccess(w, rowsAffected)
}

func EditInsight(w http.ResponseWriter, r *http.Request) {
	var insight models.Insight
	json.NewDecoder(r.Body).Decode(&insight)
	rowsAffected, err := services.EditInsight(insight)
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
	responseutils.SendSuccess(w, rowsAffected)
}

func DeleteAudience(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	rowsDeleted, err := services.DeleteAudience(id)
	if err != nil {
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, rowsDeleted)
}

func DeleteChart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	rowsDeleted, err := services.DeleteChart(id)
	if err != nil {
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, rowsDeleted)
}

func DeleteInsight(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	rowsDeleted, err := services.DeleteInsight(id)
	if err != nil {
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, rowsDeleted)
}
