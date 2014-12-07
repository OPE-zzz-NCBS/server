package api

import (
	"net/http"
	"github.com/OPENCBS/server/repo"
)

func GetCities(w http.ResponseWriter, r *APIRequest) {
	repo := repo.NewCityRepo(r.DbProvider)
	cities, err := repo.GetAll()
	if err != nil {
		sendInternalServerError(w, err)
		return
	}
	sendJson(w, cities)
}
