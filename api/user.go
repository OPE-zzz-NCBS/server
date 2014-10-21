package api

import (
	"encoding/json"
	"net/http"
	"log"
	"strconv"
	"database/sql"
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

	js, err = json.MarshalIndent(users, "", "  ")
	if err != nil {
		goto Error
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
	return

Error:
	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var js []byte
	var repo iface.UserRepo
	var user *model.User
	var db *sql.DB

	idString := r.URL.Query().Get(":id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		goto Error
	}

	db, err = iface.GetDb(r)
	if err != nil {
		goto Error
	}

	repo = iface.NewUserRepo()
	user, err = repo.FindById(db, id)
	if err != nil {
		goto Error
	}

	js, err = json.MarshalIndent(user, "", "  ")
	if err != nil {
		goto Error
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
	return

Error:
	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

