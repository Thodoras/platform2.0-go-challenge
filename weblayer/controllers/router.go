package controllers

import "github.com/gorilla/mux"

// Route will map routes to the corresponding endpoints
func Route(router *mux.Router) {
	router.HandleFunc("/assets/{id}", GetAllAssets).Methods("GET")
}
