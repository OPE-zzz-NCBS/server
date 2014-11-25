package api

import (
	"net/http"
	"github.com/OPENCBS/server/iface"
)

func GetCities(w http.ResponseWriter, r *http.Request) {
	repo := iface.NewCityRepo()
	cities, err := repo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}
	sendJson(w, cities)
}

