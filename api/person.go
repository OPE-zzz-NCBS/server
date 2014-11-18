package api

import (
	"net/http"
	"strconv"
	"github.com/OPENCBS/server/iface"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get(":id")
	id, _ := strconv.Atoi(idString)
	repo := iface.NewPersonRepo()
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

