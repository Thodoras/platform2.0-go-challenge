package main

import (
	"github.com/gorilla/mux"
	"platform2.0-go-challenge/src/helpers/security"
)

func route(router *mux.Router) {
	router.HandleFunc("/assets/{user_id}", security.Authorize(AssetController.GetAllAssets)).Methods("GET")
	router.HandleFunc("/assets/paginated/{user_id}", security.Authorize(AssetController.GetAllAssetsPaginated)).Methods("GET")
	router.HandleFunc("/assets/audiences/{user_id}", security.Authorize(AssetController.AddAudience)).Methods("POST")
	router.HandleFunc("/assets/charts/{user_id}", security.Authorize(AssetController.AddChart)).Methods("POST")
	router.HandleFunc("/assets/insights/{user_id}", security.Authorize(AssetController.AddInsight)).Methods("POST")
	router.HandleFunc("/assets/audiences/{user_id}", security.Authorize(AssetController.EditAudience)).Methods("PUT")
	router.HandleFunc("/assets/charts/{user_id}", security.Authorize(AssetController.EditChart)).Methods("PUT")
	router.HandleFunc("/assets/insights/{user_id}", security.Authorize(AssetController.EditInsight)).Methods("PUT")
	router.HandleFunc("/assets/audiences/{user_id}/delete/{id}", security.Authorize(AssetController.DeleteAudience)).Methods("DELETE")
	router.HandleFunc("/assets/charts/{user_id}/delete/{id}", security.Authorize(AssetController.DeleteChart)).Methods("DELETE")
	router.HandleFunc("/assets/insights/{user_id}/delete/{id}", security.Authorize(AssetController.DeleteInsight)).Methods("DELETE")

	router.HandleFunc("/users/signup", UserController.SignUp).Methods("POST")
	router.HandleFunc("/users/login", UserController.Login).Methods("POST")
}
