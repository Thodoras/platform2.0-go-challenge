package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllAssets will get a user id and return all related assets.
func GetAllAssets(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	fmt.Println("Hi!" + userID)
}
