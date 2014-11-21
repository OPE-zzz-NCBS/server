package api

import (
	"net/http"
	"github.com/OPENCBS/server/iface"
)

func GetBranches(w http.ResponseWriter, r *http.Request) {
	repo := iface.NewBranchRepo()
	branches, err := repo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}
	sendJson(w, branches)
}

