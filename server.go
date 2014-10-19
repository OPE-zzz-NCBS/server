package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/api"
)

func main() {
	bootstrap()

	db, err := iface.GetDb(nil)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Set up routes
	router := mux.NewRouter()
	router.HandleFunc("/api/users", api.GetUsers).Methods("GET")

	http.Handle("/", router)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

