package api

import (
	"encoding/json"
	"net/http"
	"github.com/OPENCBS/server/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []model.User{
		model.User{1, "pbastov", "Pavel", "Bastov"},
	}

	js, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

