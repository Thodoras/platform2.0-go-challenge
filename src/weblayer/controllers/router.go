package controllers

import (
	"github.com/gorilla/mux"
	"platform2.0-go-challenge/src/helpers/security"
)

// Route will map routes to the corresponding endpoints
func Route(router *mux.Router) {
	router.HandleFunc("/assets/{user_id}", security.Authorize(GetAllAssets)).Methods("GET")
	router.HandleFunc("/assets/paginated/{user_id}", security.Authorize(GetAllAssetsPaginated)).Methods("GET")
	router.HandleFunc("/assets/audiences/{user_id}", security.Authorize(AddAudience)).Methods("POST")
	router.HandleFunc("/assets/charts/{user_id}", security.Authorize(AddChart)).Methods("POST")
	router.HandleFunc("/assets/insights/{user_id}", security.Authorize(AddInsight)).Methods("POST")
	router.HandleFunc("/assets/audiences/{user_id}", security.Authorize(EditAudience)).Methods("PUT")
	router.HandleFunc("/assets/charts/{user_id}", security.Authorize(EditChart)).Methods("PUT")
	router.HandleFunc("/assets/insights/{user_id}", security.Authorize(EditInsight)).Methods("PUT")
	router.HandleFunc("/assets/audiences/{user_id}/delete/{id}", security.Authorize(DeleteAudience)).Methods("DELETE")
	router.HandleFunc("/assets/charts/{user_id}/delete/{id}", security.Authorize(DeleteChart)).Methods("DELETE")
	router.HandleFunc("/assets/insights/{user_id}/delete/{id}", security.Authorize(DeleteInsight)).Methods("DELETE")

	router.HandleFunc("/users/signup", SignUp).Methods("POST")
	router.HandleFunc("/users/login", Login).Methods("POST")
}
