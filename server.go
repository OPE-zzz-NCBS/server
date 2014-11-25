package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/drone/routes"
	"github.com/OPENCBS/server/config"
	"github.com/OPENCBS/server/api"
	"github.com/OPENCBS/server/model"
)

func protected(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Access-Token")
		session := model.GetSession(token)
		if session == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			fn(w, r)
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Write([]byte("OPENCBS Server is running"))
}

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	bootstrap()

	// Set up routes
	mux := routes.New()
	mux.Get("/", home)
	mux.Get("/api/users", protected(api.GetUsers))
	mux.Get("/api/users/:id([0-9]+)", protected(api.GetUser))

	mux.Post("/api/sessions", api.AddSession)
	mux.Del("/api/sessions", protected(api.DeleteSession))

	mux.Get("/api/clients", protected(api.GetClients))
	mux.Get("/api/people/:id([0-9]+)", api.GetPerson)

	mux.Get("/api/lookupdata", api.GetLookupData)
	mux.Get("/api/activities", api.GetActivities)
	mux.Get("/api/branches", api.GetBranches)

	mux.Get("/api/cities", api.GetCities)

	http.Handle("/", mux)
	log.Println("OPENCBS Server")
	addr := fmt.Sprintf(":%d", conf.Server.Port)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

