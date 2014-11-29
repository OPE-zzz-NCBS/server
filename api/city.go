package api

import (
	"net/http"
	"github.com/OPENCBS/server/app"
	"github.com/OPENCBS/server/repo"
)

func GetCities(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	repo := repo.NewCityRepo(ctx.DbProvider)
	cities, err := repo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}
	sendJson(w, cities)
}

