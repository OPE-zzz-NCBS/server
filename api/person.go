package api

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/OPENCBS/server/repo"
	"github.com/OPENCBS/server/app"
)

func GetPerson(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)
	repo := repo.NewPersonRepo(ctx.DbProvider)
	person, err := repo.GetById(id)
	if err != nil {
		fail(w, err)
		return
	}
	if person == nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	sendJson(w, person)
}

