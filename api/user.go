package api

import (
	"encoding/json"
	"net/http"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var js []byte
	var repo iface.UserRepo
	var users []model.User

	db, err := iface.GetDb(r)
	if err != nil {
		goto Error
	}

	repo = iface.NewUserRepo()
	users, err = repo.FindAll(db)
	if err != nil {
		goto Error
	}

	js, err = json.Marshal(users)
	if err != nil {
		goto Error
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
	return

Error:
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

