package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"platform2.0-go-challenge/weblayer/controllers"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	router := mux.NewRouter()
	controllers.Route(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
