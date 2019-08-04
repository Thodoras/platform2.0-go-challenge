package assetController

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
	"platform2.0-go-challenge/src/servicelayer/services/assetService"

	"github.com/gorilla/mux"
)

type AssetController struct {
	assetService assetService.IAssetService
}

func NewAssetController(assetService assetService.IAssetService) *AssetController {
	return &AssetController{assetService: assetService}
}

func (a *AssetController) GetAllAssets(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	response, err := a.assetService.GetAllAssets(userID)
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

func (a *AssetController) GetAllAssetsPaginated(w http.ResponseWriter, r *http.Request) {
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

	response, err := a.assetService.GetAllAssetsPaginated(userID, limit, offset)
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

func (a *AssetController) AddAudience(w http.ResponseWriter, r *http.Request) {
	var audience models.Audience
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&audience)
	audience.UserID = userID
	id, err := a.assetService.AddAudience(audience)
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

func (a *AssetController) AddChart(w http.ResponseWriter, r *http.Request) {
	var chart models.Chart
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&chart)
	chart.UserID = userID
	id, err := a.assetService.AddChart(chart)
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

func (a *AssetController) AddInsight(w http.ResponseWriter, r *http.Request) {
	var insight models.Insight
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&insight)
	insight.UserID = userID
	id, err := a.assetService.AddInsight(insight)
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

func (a *AssetController) EditAudience(w http.ResponseWriter, r *http.Request) {
	var audience models.Audience
	json.NewDecoder(r.Body).Decode(&audience)
	rowsAffected, err := a.assetService.EditAudience(audience)
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

func (a *AssetController) EditChart(w http.ResponseWriter, r *http.Request) {
	var chart models.Chart
	json.NewDecoder(r.Body).Decode(&chart)
	rowsAffected, err := a.assetService.EditChart(chart)
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

func (a *AssetController) EditInsight(w http.ResponseWriter, r *http.Request) {
	var insight models.Insight
	json.NewDecoder(r.Body).Decode(&insight)
	rowsAffected, err := a.assetService.EditInsight(insight)
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

func (a *AssetController) DeleteAudience(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	rowsDeleted, err := a.assetService.DeleteAudience(id)
	if err != nil {
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, rowsDeleted)
}

func (a *AssetController) DeleteChart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	rowsDeleted, err := a.assetService.DeleteChart(id)
	if err != nil {
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, rowsDeleted)
}

func (a *AssetController) DeleteInsight(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseutils.SendError(w, http.StatusBadRequest, errors.New("Bad request"))
		log.Println(err)
		return
	}
	rowsDeleted, err := a.assetService.DeleteInsight(id)
	if err != nil {
		responseutils.SendError(w, http.StatusInternalServerError, errors.New("Server Error"))
		log.Println(err)
		return
	}
	responseutils.SendSuccess(w, rowsDeleted)
}
