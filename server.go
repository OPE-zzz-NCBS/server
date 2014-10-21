package main

import (
	"log"
	"net/http"
	"github.com/drone/routes"
	"github.com/OPENCBS/server/api"
	"github.com/OPENCBS/server/model"
)

func protected(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Authentication-Token")
		session := model.GetSession(token)
		if session == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			fn(w, r)
		}
	}
}

func main() {
	bootstrap()

	// Set up routes
	mux := routes.New()
	mux.Get("/api/users", protected(api.GetUsers))
	mux.Get("/api/users/:id([0-9]+)", api.GetUser)

	mux.Post("/api/sessions", api.AddSession)
	mux.Del("/api/sessions", protected(api.DeleteSession))

	http.Handle("/", mux)
	log.Println("Listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

