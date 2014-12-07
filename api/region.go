package api

import (
	"net/http"
	"github.com/OPENCBS/server/repo"
)

func GetRegions(w http.ResponseWriter, r *APIRequest) {
	repo := repo.NewRegionRepo(r.DbProvider)
	regions, err := repo.GetAll()
	if err != nil {
		sendInternalServerError(w, err)
		return
	}
	sendJson(w, regions)
}
