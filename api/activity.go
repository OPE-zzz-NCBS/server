package api

import (
	"net/http"
	"github.com/OPENCBS/server/iface"
)

func GetActivities(w http.ResponseWriter, r *http.Request) {
	repo := iface.NewActivityRepo()
	activities, err := repo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}
	
	sendJson(w, activities)
}

