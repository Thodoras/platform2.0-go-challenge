package reponseutils

import (
	"encoding/json"
	"net/http"

	"platform2.0-go-challenge/helpers/logutils"
)

func SendError(w http.ResponseWriter, status int, err logutils.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
