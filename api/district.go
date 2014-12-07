package api

import (
	"net/http"
	"github.com/OPENCBS/server/repo"
)

func GetDistricts(w http.ResponseWriter, r *APIRequest) {
	repo := repo.NewDistrictRepo(r.DbProvider)
	districts, err := repo.GetAll()
	if err != nil {
		sendInternalServerError(w, err)
		return
	}
	sendJson(w, districts)
}
