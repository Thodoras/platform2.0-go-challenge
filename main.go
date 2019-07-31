package main

import (
	"log"
	"net/http"

	"platform2.0-go-challenge/datalayer/repositories"
	"platform2.0-go-challenge/helpers/drivers"

	"github.com/gorilla/mux"
	"platform2.0-go-challenge/weblayer/controllers"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	repositories.DB = drivers.ConnectPostgresDB()
}

func main() {
	router := mux.NewRouter()
	controllers.Route(router)
	log.Fatal(http.ListenAndServe(":8001", router))
}
