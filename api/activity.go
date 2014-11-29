package api

import (
	"net/http"
	"github.com/OPENCBS/server/repo"
	"github.com/OPENCBS/server/app"
)

func GetActivities(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	repo := repo.NewActivityRepo(ctx.DbProvider)
	activities, err := repo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}
	sendJson(w, activities)
}

