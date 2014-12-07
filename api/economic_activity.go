package api

import (
	"net/http"
	"github.com/OPENCBS/server/repo"
)

func GetEconomicActivities(w http.ResponseWriter, r *APIRequest) {
	repo := repo.NewEconomicActivityRepo(r.DbProvider)
	economicActivities, err := repo.GetAll()
	if err != nil {
		sendInternalServerError(w, err)
		return
	}
	sendJson(w, economicActivities)
}

