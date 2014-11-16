package api

import (
	"encoding/json"
	"net/http"
	"log"
	"strconv"
	"github.com/OPENCBS/server/factory"
	"github.com/OPENCBS/server/repo"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/util"
)

func fail(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := factory.GetDb(r)
	if err != nil {
		fail(w, err)
		return
	}

	sqlProvider := factory.GetSqlProvider(r)
	repo := repo.NewUserRepo(sqlProvider)
	offset := util.GetOffset(r)
	limit := util.GetLimit(r)
	items, err := repo.GetAll(db, offset, limit)
	if err != nil {
		fail(w, err)
		return
	}
	for _, user := range items {
		user.Href = util.GetUserUrl(r, user)
	}
	users := new(model.Users)
	users.Href = util.GetUsersUrl(r)
	users.Offset = offset
	users.Limit = limit
	users.Items = items

	js, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fail(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get(":id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		fail(w, err)
		return
	}

	db, err := factory.GetDb(r)
	if err != nil {
		fail(w, err)
		return
	}

	sqlProvider := factory.GetSqlProvider(r)
	repo := repo.NewUserRepo(sqlProvider)
	user, err := repo.GetById(db, id)
	if err != nil {
		fail(w, err)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	user.Href = util.GetUserUrl(r, user)

	js, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fail(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
}

