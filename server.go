package main

import (
	"log"
	"net/http"
	//"github.com/OPENCBS/server/iface"
	"github.com/drone/routes"
	"github.com/OPENCBS/server/api"
)

func main() {
	bootstrap()

	/*
	db, err := iface.GetDb(nil)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	*/

	// Set up routes
	mux := routes.New()
	mux.Get("/api/users", api.GetUsers)
	mux.Get("/api/users/:id([0-9]+)", api.GetUser)

	http.Handle("/", mux)
	log.Println("Listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

