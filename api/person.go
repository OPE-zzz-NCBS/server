package api

import (
	"net/http"
	"github.com/OPENCBS/server/repo"
)

/*
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
*/

func GetPeople(w http.ResponseWriter, r *APIRequest) {
	repo := repo.NewPersonRepo(r.DbProvider)
	offset, limit := r.GetRange()
	people, err := repo.GetPeople(offset, limit)
	if err != nil {
		sendInternalServerError(w, err)
		return
	}
	sendJson(w, people)
}
