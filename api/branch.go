package api

import (
	"net/http"
	"github.com/OPENCBS/server/app"
	"github.com/OPENCBS/server/repo"
)

func GetBranches(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	repo := repo.NewBranchRepo(ctx.DbProvider)
	branches, err := repo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}
	sendJson(w, branches)
}

