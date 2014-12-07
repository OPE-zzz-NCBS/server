package api

import (
	"net/http"
	"github.com/OPENCBS/server/repo"
)

func GetBranches(w http.ResponseWriter, r *APIRequest) {
	repo := repo.NewBranchRepo(r.DbProvider)
	branches, err := repo.GetAll()
	if err != nil {
		sendInternalServerError(w, err)
		return
	}
	sendJson(w, branches)
}

