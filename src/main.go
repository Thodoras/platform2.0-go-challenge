package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	wireDependencies()
}

func main() {
	router := mux.NewRouter()
	route(router)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}
