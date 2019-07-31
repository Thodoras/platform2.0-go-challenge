package controllers

import "github.com/gorilla/mux"

// Route will map routes to the corresponding endpoints
func Route(router *mux.Router) {
	router.HandleFunc("/assets/{user_id}", GetAllAssets).Methods("GET")
	router.HandleFunc("/assets/audiences/{user_id}", AddAudience).Methods("POST")
	router.HandleFunc("/assets/charts/{user_id}", AddChart).Methods("POST")
	router.HandleFunc("/assets/insights/{user_id}", AddInsight).Methods("POST")
}
